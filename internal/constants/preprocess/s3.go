package preprocess

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type versionFile struct {
	Version string
}

func startSession() (*session.Session, error) {
	// Enable loading shared config file
	os.Setenv("aws_SDK_LOAD_CONFIG", "1")

	// Specify profile to load for the session's config
	return session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config:  aws.Config{Region: aws.String("ca-central-1")},
	})
}

func getVersionString(branchName string) (string, error) {
	session, err := startSession()
	if err != nil {
		return "", err
	}

	downloader := s3manager.NewDownloader(session)

	var buffer []byte
	atBuffer := aws.NewWriteAtBuffer(buffer)

	params := &s3.GetObjectInput{
		Bucket: aws.String("cli-update"),
		// Key:    aws.String(fmt.Sprintf("update/state/versions/%s/version.json", branchName)),
		// TEMPORARY: Use current PRs version file
		Key: aws.String("update/state/versions/sync-version-increments-171012882/version.json"),
	}

	_, err = downloader.Download(atBuffer, params)
	if err != nil {
		return "", err
	}

	version := &versionFile{}
	err = json.Unmarshal(atBuffer.Bytes(), version)
	if err != nil {
		return "", err
	}

	return version.Version, nil
}
