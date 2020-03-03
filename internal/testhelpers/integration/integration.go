package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/pborman/ansi"
	"github.com/phayes/permbits"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/pkg/expect"
	"github.com/ActiveState/cli/pkg/projectfile"
)

var (
	PersistentUsername = "cli-integration-tests"
	PersistentPassword = "test-cli-integration"

	defaultTimeout = 10 * time.Second
	authnTimeout   = 40 * time.Second
)

// Suite is our integration test suite
type Suite struct {
	suite.Suite
	executable string
	env        []string
	wd         *string
}

// SetupTest sets up an integration test suite for testing the state tool executable
func (s *Suite) SetupTest() {
	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}

	root := environment.GetRootPathUnsafe()
	executable := filepath.Join(root, "build/"+constants.CommandName+exe)

	s.wd = nil

	if !fileutils.FileExists(executable) {
		s.FailNow("Integration tests require you to have built a state tool binary. Please run `state run build`.")
	}

	configDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)
	cacheDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)
	binDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)

	fmt.Println("Configdir: " + configDir)
	fmt.Println("Cachedir: " + cacheDir)
	fmt.Println("Bindir: " + binDir)

	s.executable = filepath.Join(binDir, constants.CommandName+exe)
	fail := fileutils.CopyFile(executable, s.executable)
	s.Require().NoError(fail.ToError())

	permissions, _ := permbits.Stat(s.executable)
	permissions.SetUserExecute(true)
	err = permbits.Chmod(s.executable, permissions)
	s.Require().NoError(err)

	s.ClearEnv()
	s.AppendEnv(os.Environ())
	s.AppendEnv([]string{
		"ACTIVESTATE_CLI_CONFIGDIR=" + configDir,
		"ACTIVESTATE_CLI_CACHEDIR=" + cacheDir,
		"ACTIVESTATE_CLI_DISABLE_UPDATES=true",
		"ACTIVESTATE_CLI_DISABLE_RUNTIME=true",
		"ACTIVESTATE_PROJECT=",
	})
}

// PrepareTemporaryWorkingDirectory prepares a temporary working directory to run the tests in
// It returns the directory name a clean-up function
func (s *Suite) PrepareTemporaryWorkingDirectory(prefix string) (tempDir string, cleanup func()) {

	tempDir, err := ioutil.TempDir("", prefix)
	s.Require().NoError(err)
	err = os.RemoveAll(tempDir)
	s.Require().NoError(err)
	err = os.MkdirAll(tempDir, 0770)
	s.Require().NoError(err)
	dir, err := filepath.EvalSymlinks(tempDir)
	s.Require().NoError(err)
	s.SetWd(dir)

	return tempDir, func() {
		_ = os.RemoveAll(dir)
		if tempDir != dir {
			_ = os.RemoveAll(tempDir)
		}
	}
}

// PrepareActiveStateYAML creates a projectfile.Project instance from the
// provided contents and saves the output to an as.y file within the named
// directory.
func (s *Suite) PrepareActiveStateYAML(dir, contents string) {
	msg := "cannot setup activestate.yaml file"

	contents = strings.TrimSpace(contents)
	projectFile := &projectfile.Project{}

	err := yaml.Unmarshal([]byte(contents), projectFile)
	s.Require().NoError(err, msg)

	projectFile.SetPath(filepath.Join(dir, "activestate.yaml"))
	fail := projectFile.Save()
	s.Require().NoError(fail.ToError(), msg)
}

func (s *Suite) PrepareFile(path, contents string) {
	errMsg := fmt.Sprintf("cannot setup file %q", path)

	contents = strings.TrimSpace(contents)

	err := os.MkdirAll(filepath.Dir(path), 0770)
	s.Require().NoError(err, errMsg)

	bs := append([]byte(contents), '\n')

	err = ioutil.WriteFile(path, bs, 0660)
	s.Require().NoError(err, errMsg)
}

// Executable returns the path to the executable under test (state tool)
func (s *Suite) Executable() string {
	return s.executable
}

// ClearEnv removes all environment variables
func (s *Suite) ClearEnv() {
	s.env = []string{}
}

// AppendEnv appends new environment variable settings
func (s *Suite) AppendEnv(env []string) {
	s.env = append(s.env, env...)
}

// SetWd specifies a working directory for the spawned processes.
// Use this method if you rely on running the test executable in a clean directory.
// By default all tests are run in `os.TempDir()`.
// SetWd returns a function that unsets the working directory. Use this if
// you do not want other tests to use the set directory.
func (s *Suite) SetWd(dir string) {
	s.wd = &dir
}

// Spawn executes the state tool executable under test in a pseudo-terminal
func (s *Suite) Spawn(args ...string) *ConsoleProcess {
	return s.SpawnCustom(s.executable, args...)
}

type ConsoleProcess struct {
	stateCh chan error
	console *expect.Console
	cmd     *exec.Cmd
	suite   suite.Suite
	ctx     context.Context
	cancel  func()
}

// SpawnCustom executes an executable in a pseudo-terminal for integration tests
func (s *Suite) SpawnCustom(executable string, args ...string) *ConsoleProcess {
	var wd string
	if s.wd == nil {
		wd = fileutils.TempDirUnsafe()
	} else {
		wd = *s.wd
	}

	cmd := exec.Command(executable, args...)
	cmd.Dir = wd
	cmd.Env = s.env

	// Create the process in a new process group.
	// This makes the behavior more consistent, as it isolates the signal handling from
	// the parent processes, which are dependent on the test environment.
	cmd.SysProcAttr = osutils.SysProcAttrForNewProcessGroup()
	fmt.Printf("Spawning '%s' from %s\n", osutils.CmdString(cmd), wd)

	var err error
	console, err := expect.NewConsole(
		expect.WithDefaultTimeout(defaultTimeout),
		expect.WithReadBufferMutation(ansi.Strip),
	)
	s.Require().NoError(err)

	err = console.Pty.StartProcessInTerminal(cmd)
	s.Require().NoError(err)

	ctx, cancel := context.WithCancel(context.Background())

	cp := &ConsoleProcess{
		stateCh: make(chan error),
		suite:   s.Suite,
		console: console,
		cmd:     cmd,
		ctx:     ctx,
		cancel:  cancel,
	}

	go func() {
		defer close(cp.stateCh)
		err := cmd.Wait()

		console.Close()

		fmt.Printf("send err to channel: %v\n", err)
		select {
		case cp.stateCh <- err:
		case <-cp.ctx.Done():
		}

		fmt.Printf("done sending err to channel: %v\n", err)
	}()

	return cp
}

func (cp *ConsoleProcess) Close() error {
	fmt.Println("closing channel")
	cp.cancel()
	if cp.cmd.ProcessState.Exited() {
		return nil
	}
	err := cp.cmd.Process.Kill()
	if err == nil {
		return nil
	}
	return cp.cmd.Process.Signal(syscall.SIGTERM)
}

// UnsyncedOutput returns the current Terminal snapshot.
// However the goroutine that creates this output is separate from this
// function so any output is not synced
func (cp *ConsoleProcess) UnsyncedOutput() string {
	return cp.console.Pty.State.String()
}

// ExpectRe listens to the terminal output and returns once the expected regular expression is matched or
// a timeout occurs
// Default timeout is 10 seconds
func (cp *ConsoleProcess) ExpectRe(value string, timeout ...time.Duration) {
	opts := []expect.ExpectOpt{expect.RegexpPattern(value)}
	if len(timeout) > 0 {
		opts = append(opts, expect.WithTimeout(timeout[0]))
	}
	_, err := cp.console.Expect(opts...)
	if err != nil {
		cp.suite.FailNow(
			"Could not meet expectation",
			"Expectation: '%s'\nError: %v\n---\nTerminal snapshot:\n%s\n---\n",
			value, err, cp.UnsyncedOutput())
	}
}

// TerminalSnapshot returns a snapshot of the terminal output
func (cp *ConsoleProcess) TerminalSnapshot() string {
	return cp.console.Pty.State.String()
}

// Expect listens to the terminal output and returns once the expected value is found or
// a timeout occurs
// Default timeout is 10 seconds
func (cp *ConsoleProcess) Expect(value string, timeout ...time.Duration) {
	opts := []expect.ExpectOpt{expect.String(value)}
	if len(timeout) > 0 {
		opts = append(opts, expect.WithTimeout(timeout[0]))
	}
	_, err := cp.console.Expect(opts...)
	if err != nil {
		cp.suite.FailNow(
			"Could not meet expectation",
			"Expectation: '%s'\nError: %v\n---\nTerminal snapshot:\n%s\n---\n",
			value, err, cp.UnsyncedOutput())
	}
}

// WaitForInput returns once a shell prompt is active on the terminal
// Default timeout is 10 seconds
func (cp *ConsoleProcess) WaitForInput(timeout ...time.Duration) {
	usr, err := user.Current()
	cp.suite.Require().NoError(err)

	msg := "echo wait_ready_$HOME"
	if runtime.GOOS == "windows" {
		msg = "echo wait_ready_%USERPROFILE%"
	}

	cp.SendLine(msg)
	cp.Expect("wait_ready_"+usr.HomeDir, timeout...)
}

// SendLine sends a new line to the terminal, as if a user typed it
func (cp *ConsoleProcess) SendLine(value string) {
	_, err := cp.console.SendLine(value)
	if err != nil {
		cp.suite.FailNow("Could not send data to terminal", "error: %v", err)
	}
}

// Send sends a string to the terminal as if a user typed it
func (cp *ConsoleProcess) Send(value string) {
	_, err := cp.console.Send(value)
	if err != nil {
		cp.suite.FailNow("Could not send data to terminal", "error: %v", err)
	}
}

// Signal sends an arbitrary signal to the running process
func (cp *ConsoleProcess) Signal(sig os.Signal) error {
	return cp.cmd.Process.Signal(sig)
}

// SendCtrlC tries to emulate what would happen in an interactive shell, when the user presses Ctrl-C
func (cp *ConsoleProcess) SendCtrlC() {
	cp.Send(string([]byte{0x03})) // 0x03 is ASCI character for ^C
}

// Quit sends an interrupt signal to the tested process
func (cp *ConsoleProcess) Quit() error {
	return cp.cmd.Process.Signal(os.Interrupt)
}

// Stop sends an interrupt signal for the tested process and fails if no process has been started yet.
func (cp *ConsoleProcess) Stop() error {
	if cp.cmd == nil || cp.cmd.Process == nil {
		cp.suite.FailNow("stop called without a spawned process")
	}
	return cp.Quit()
}

// LoginAsPersistentUser is a common test case after which an integration test user should be logged in to the platform
func (s *Suite) LoginAsPersistentUser() {
	cp := s.Spawn("auth", "--username", PersistentUsername, "--password", PersistentPassword)
	defer cp.Close()
	fmt.Println("1")
	cp.Expect("successfully authenticated", authnTimeout)
	fmt.Println("2")
	cp.ExpectExitCode(0)
	fmt.Println("3")
}

// ExpectExitCode waits for the program under test to terminate, and checks that the returned exit code meets expectations
func (cp *ConsoleProcess) ExpectExitCode(exitCode int, timeout ...time.Duration) {
	ps, err := cp.Wait(timeout...)
	if err != nil {
		cp.suite.FailNow(
			"Error waiting for process:",
			"\n%v\n---\nTerminal snapshot:\n%s\n---\n",
			err, cp.TerminalSnapshot())
	}
	if ps.ExitCode() != exitCode {
		cp.suite.FailNow(
			"Process terminated with unexpected exit code\n",
			"Expected: %d, got %d\n---\nTerminal snapshot:\n%s\n---\n",
			exitCode, ps.ExitCode(), cp.TerminalSnapshot())
	}
}

// ExpectNotExitCode waits for the program under test to terminate, and checks that the returned exit code is not the value provide
func (cp *ConsoleProcess) ExpectNotExitCode(exitCode int, timeout ...time.Duration) {
	ps, err := cp.Wait(timeout...)
	if err != nil {
		cp.suite.FailNow(
			"Error waiting for process:",
			"\n%v\n---\nTerminal snapshot:\n%s\n---\n",
			err, cp.TerminalSnapshot())
	}
	if ps.ExitCode() == exitCode {
		cp.suite.FailNow(
			"Process terminated with unexpected exit code\n",
			"Expected anything except: %d, got %d\n---\nTerminal snapshot:\n%s\n---\n",
			exitCode, ps.ExitCode(), cp.TerminalSnapshot())
	}
}

// Wait waits for the tested process to finish and returns its state including ExitCode
func (cp *ConsoleProcess) Wait(timeout ...time.Duration) (state *os.ProcessState, err error) {
	if cp.cmd == nil || cp.cmd.Process == nil {
		return
	}

	t := defaultTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}

	fmt.Printf("waiting for EOF\n")
	// TODO: This might need to be different for Windows, I think that Windows sends a different error message when we close the pseudo-terminal...
	_, err = cp.console.Expect(expect.PTSClosed, expect.EOF, expect.WithTimeout(t))
	fmt.Printf("EOF received: %v\n", err)

	if err != nil /* && err is timeout (?) */ {
		fmt.Println("killing process")
		err = cp.cmd.Process.Kill()
		if err != nil {
			// Don't know what else to do otherwise, honestly...
			panic(err)
		}
	}

	fmt.Printf("Waiting for stateCh")
	select {
	case pErr := <-cp.stateCh:
		fmt.Printf("got error: %v\n", pErr)
		return cp.cmd.ProcessState, pErr
	case <-cp.ctx.Done():
		return nil, fmt.Errorf("context canceled")
	}
}

// UnsyncedTrimSpaceOutput displays the terminal output a user would see
// however the goroutine that creates this output is separate from this
// function so any output is not synced
func (cp *ConsoleProcess) UnsyncedTrimSpaceOutput() string {
	// When the PTY reaches 80 characters it continues output on a new line.
	// On Windows this means both a carriage return and a new line. Windows
	// also picks up any spaces at the end of the console output, hence all
	// the cleaning we must do here.
	newlineRe := regexp.MustCompile(`\r?\n`)
	return newlineRe.ReplaceAllString(strings.TrimSpace(cp.UnsyncedOutput()), "")
}

func (s *Suite) CreateNewUser() string {
	uid, err := uuid.NewRandom()
	s.Require().NoError(err)

	username := fmt.Sprintf("user-%s", uid.String()[0:8])
	password := username
	email := fmt.Sprintf("%s@test.tld", username)

	cp := s.Spawn("auth", "signup")
	defer cp.Close()
	cp.Expect("username:")
	cp.SendLine(username)
	cp.Expect("password:")
	cp.SendLine(password)
	cp.Expect("again:")
	cp.SendLine(password)
	cp.Expect("name:")
	cp.SendLine(username)
	cp.Expect("email:")
	cp.SendLine(email)
	cp.Expect("account has been registered", authnTimeout)
	cp.ExpectExitCode(0)

	return username
}
