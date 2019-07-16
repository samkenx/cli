package integration_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/test/integration/expectx"
	tsuite "github.com/stretchr/testify/suite"
)

var persistentUsername = "cli-integration-tests"
var persistentPassword = "test-cli-integration"

type Suite struct {
	tsuite.Suite
	*expectx.Expectx
}

func (s *Suite) SetupTest() {
	root := environment.GetRootPathUnsafe()
	failNow := func(format string, args ...interface{}) {
		s.FailNow(format, args...)
	}
	s.Expectx = expectx.New(filepath.Join(root, "build/state"), failNow)

	configDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)
	cacheDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)

	fmt.Println("Configdir: " + configDir)
	fmt.Println("Cachedir: " + cacheDir)

	s.ClearEnv()
	s.AppendEnv(os.Environ())
	s.AppendEnv([]string{
		"ACTIVESTATE_CLI_CONFIGDIR=" + configDir,
		"ACTIVESTATE_CLI_CACHEDIR=" + cacheDir,
		"ACTIVESTATE_CLI_DISABLE_UPDATES=true",
		"ACTIVESTATE_CLI_DISABLE_RUNTIME=true",
		"SHELL=bash",
		"VERBOSE=true",
	})
}

func (s *Suite) LoginAsPersistentUser() {
	s.Spawn("auth", "--username", persistentUsername, "--password", persistentPassword)
	s.Expect("succesfully authenticated")
	s.Wait()
}
