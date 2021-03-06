package runtime

import (
	"net/url"
	"path"
	"path/filepath"

	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/internal/analytics"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/download"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/progress"
	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/buildlogstream"
	"github.com/ActiveState/cli/pkg/platform/api/headchef"
	"github.com/ActiveState/cli/pkg/platform/api/headchef/headchef_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
)

// ensure that Downloader implements the Download interface
var _ Downloader = &Download{}

// InstallerTestsSubstr is used to exclude test artifacts, we don't care about them
const InstallerTestsSubstr = "-tests."

var (
	// FailNoCommit indicates a failure due to there not being a commit
	FailNoCommit = failures.Type("runtime.fail.nocommit")

	// FailNoArtifacts indicates a failure due to the project not containing any artifacts
	FailNoArtifacts = failures.Type("runtime.fail.noartifacts")

	// FailNoValidArtifact indicates a failure due to the project not containing any valid artifacts
	FailNoValidArtifact = failures.Type("runtime.fail.novalidartifact")

	// FailBuildFailed indicates a failure due to the build failing
	FailBuildFailed = failures.Type("runtime.fail.buildfailed", failures.FailUser)

	// FailBuildInProgress indicates a failure due to the build being in progress
	FailBuildInProgress = failures.Type("runtime.fail.buildinprogress", failures.FailUser)

	// FailBuildBadResponse indicates a failure due to the build req/resp malfunctioning
	FailBuildBadResponse = failures.Type("runtime.fail.buildbadresponse")

	// FailBuildErrResponse indicates a failure due to the build req/resp returning an error
	FailBuildErrResponse = failures.Type("runtime.fail.builderrresponse")

	// FailArtifactInvalidURL indicates a failure due to an artifact having an invalid URL
	FailArtifactInvalidURL = failures.Type("runtime.fail.invalidurl")
)

// HeadChefArtifact is a convenient type alias cause swagger generates some really shitty code
type HeadChefArtifact = headchef_models.Artifact

// FetchArtifactsResult stores the information needed by the installer to
// install and assemble a runtime environment.
// This information is extracted from a build request response in the
// FetchArtifacts() method
type FetchArtifactsResult struct {
	BuildEngine BuildEngine
	Artifacts   []*HeadChefArtifact
	RecipeID    strfmt.UUID
}

// DownloadDirectoryProvider provides download directories for individual artifacts
type DownloadDirectoryProvider interface {

	// DownloadDirectory returns the download path for a given artifact
	DownloadDirectory(artf *HeadChefArtifact) (string, *failures.Failure)
}

// Downloader defines the behavior required to be a runtime downloader.
type Downloader interface {
	// Download will attempt to download some runtime locally and return back the filename of
	// the downloaded archive or a Failure.
	Download(artifacts []*HeadChefArtifact, d DownloadDirectoryProvider, progress *progress.Progress) (files map[string]*HeadChefArtifact, fail *failures.Failure)

	// FetchArtifacts will fetch artifact
	FetchArtifacts() (*FetchArtifactsResult, *failures.Failure)
}

// Download is the main struct for orchestrating the download of all the artifacts belonging to a runtime
type Download struct {
	runtime *Runtime
	orgID   string
	private bool
}

// NewDownload creates a new RuntimeDownload using all custom args
func NewDownload(runtime *Runtime) Downloader {
	return &Download{
		runtime: runtime,
	}
}

// fetchRecipe juggles API's to get the build request that can be sent to the head-chef
func (r *Download) fetchRecipeID() (strfmt.UUID, *failures.Failure) {
	commitID := r.runtime.commitID
	if commitID == "" {
		return "", FailNoCommit.New(locale.T("err_no_commit"))
	}

	recipeID, fail := model.FetchRecipeIDForCommitAndPlatform(commitID, r.runtime.owner, r.runtime.projectName, r.orgID, r.private, model.HostPlatform)
	if fail != nil {
		return "", fail
	}

	return *recipeID, nil
}

// FetchArtifacts will retrieve artifact information from the head-chef (eg language installers)
// The first return argument specifies whether we are dealing with an alternative build
func (r *Download) FetchArtifacts() (*FetchArtifactsResult, *failures.Failure) {
	orgID := strfmt.UUID(constants.ValidZeroUUID)
	projectID := strfmt.UUID(constants.ValidZeroUUID)

	if r.runtime.owner != "" && r.runtime.projectName != "" {
		platProject, fail := model.FetchProjectByName(r.runtime.owner, r.runtime.projectName)
		if fail != nil {
			return nil, fail
		}

		projectID = platProject.ProjectID
		orgID = platProject.OrganizationID
		r.orgID = platProject.OrganizationID.String()
		r.private = platProject.Private
	}

	recipeID, fail := r.fetchRecipeID()
	if fail != nil {
		return nil, fail
	}

	return r.fetchArtifacts(r.runtime.commitID, recipeID, orgID, projectID)
}

func (r *Download) fetchArtifacts(commitID, recipeID, orgID, projectID strfmt.UUID) (*FetchArtifactsResult, *failures.Failure) {
	result := &FetchArtifactsResult{}

	buildAnnotations := headchef.BuildAnnotations{
		CommitID:     commitID.String(),
		Project:      r.runtime.projectName,
		Organization: r.runtime.owner,
	}

	logging.Debug("sending request to head-chef")
	buildRequest, fail := headchef.NewBuildRequest(recipeID, orgID, projectID, buildAnnotations)
	if fail != nil {
		return result, fail
	}
	buildStatus := headchef.InitClient().RequestBuild(buildRequest)

	for {
		select {
		case resp := <-buildStatus.Completed:
			if len(resp.Artifacts) == 0 {
				return result, FailNoArtifacts.New(locale.T("err_no_artifacts"))
			}

			result.BuildEngine = BuildEngineFromResponse(resp)
			if result.BuildEngine == UnknownEngine {
				return result, FailRuntimeUnknownEngine.New("installer_err_engine_unknown")
			}

			if resp.RecipeID == nil {
				return result, FailBuildBadResponse.New(locale.T("err_corrupted_build_request_response"))
			}
			result.RecipeID = *resp.RecipeID
			result.Artifacts = resp.Artifacts
			logging.Debug("request engine=%v, recipeID=%s", result.BuildEngine, result.RecipeID.String())

			return result, nil

		case msg := <-buildStatus.Failed:
			logging.Debug("BuildFailed: %s", msg)
			return result, FailBuildFailed.New(locale.Tr("build_status_failed", r.projectURL(), msg))

		case resp := <-buildStatus.Started:
			logging.Debug("BuildStarted")
			namespaced := project.Namespaced{
				Owner:   r.runtime.owner,
				Project: r.runtime.projectName,
			}
			analytics.EventWithLabel(
				analytics.CatBuild, analytics.ActBuildProject, namespaced.String(),
			)

			// For non-alternate builds we do not support in-progress builds
			engine := BuildEngineFromResponse(resp)
			if engine != Alternative && engine != Hybrid {
				return result, FailBuildInProgress.New(locale.Tr("build_status_in_progress", r.projectURL()))
			}

			if err := r.waitForArtifacts(recipeID); err != nil {
				return nil, failures.FailMisc.Wrap(err, locale.Tl("err_wait_artifacts", "Error happened while waiting for packages"))
			}
			return r.fetchArtifacts(commitID, recipeID, orgID, projectID)

		case fail := <-buildStatus.RunFail:
			logging.Debug("Failure: %v", fail)

			switch {
			case fail.Type.Matches(headchef.FailBuildReqErrorResp):
				l10n := locale.Tr("build_status_unknown_error", fail.Error(), r.projectURL())
				return result, FailBuildErrResponse.New(l10n)
			default:
				l10n := locale.Tr("build_status_unknown", r.projectURL())
				return result, FailBuildBadResponse.New(l10n)
			}
		}
	}
}

func (r *Download) waitForArtifacts(recipeID strfmt.UUID) error {
	logstream := buildlogstream.NewRequest(recipeID, r.runtime.msgHandler)
	if err := logstream.Wait(); err != nil {
		return locale.WrapError(err, "err_wait_artifacts_logstream", "Error happened while waiting for builds to complete")
	}

	return nil
}

func (r *Download) projectURL() string {
	url := api.GetServiceURL(api.ServiceHeadChef)
	url.Path = path.Join(r.runtime.owner, r.runtime.projectName)
	return url.String()
}

// Download is the main function used to kick off the runtime download
func (r *Download) Download(artifacts []*HeadChefArtifact, dp DownloadDirectoryProvider, progress *progress.Progress) (files map[string]*HeadChefArtifact, fail *failures.Failure) {
	files = map[string]*HeadChefArtifact{}
	entries := []*download.Entry{}

	for _, artf := range artifacts {
		artifactURL, err := url.Parse(artf.URI.String())
		if err != nil {
			return files, FailArtifactInvalidURL.New(locale.T("err_artifact_invalid_url"))
		}
		u, fail := model.SignS3URL(artifactURL)
		if fail != nil {
			return files, fail
		}

		// Ideally we'd be passing authentication down the chain somehow, but for now this would require way too much
		// additional plumbing, so we're going to use the global version until the higher level architecture is refactored
		auth := authentication.Get()
		uid := "00000000-0000-0000-0000-000000000000"
		if auth.Authenticated() {
			uid = auth.UserID().String()
		}

		q := u.Query()
		q.Set("x-uuid", uid) // x-uuid is used so our analytics can filter out activator traffic

		// Disabled for now as `x-` seems to interact with signing negatively
		// And adding it to the URL to be signed just drops it from the resulting URL
		// u.RawQuery = q.Encode()

		targetDir, fail := dp.DownloadDirectory(artf)
		if fail != nil {
			return files, fail
		}

		targetPath := filepath.Join(targetDir, filepath.Base(u.Path))
		entries = append(entries, &download.Entry{
			Path:     targetPath,
			Download: u.String(),
		})
		files[targetPath] = artf
	}

	downloader := download.New(entries, 1, progress)
	return files, downloader.Download()
}
