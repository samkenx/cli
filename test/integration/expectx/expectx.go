package expectx

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/ActiveState/cli/internal/osutils/stacktrace"
	"github.com/ActiveState/cli/test/integration/expectx/process"
)

type Expectx struct {
	fail      func(string, ...interface{})
	proc      *process.Process
	procEnded chan bool
	env       []string
	exec      string
}

func New(exec string, fail func(string, ...interface{})) *Expectx {
	return &Expectx{
		fail: fail,
	}
}

func (e *Expectx) ClearEnv() {
	e.env = []string{}
}

func (e *Expectx) AppendEnv(env []string) {
	e.env = append(e.env, env...)
}

func (e *Expectx) Spawn(args ...string) {
	wd, _ := os.Getwd()
	commandLine := fmt.Sprintf("%s %s", e.exec, strings.Join(args, " "))
	fmt.Printf("Spawning '%s' from %s\n", commandLine, wd)
	e.proc = process.New(e.exec, args...)
	e.proc.SetEnv(e.env)
	e.procEnded = make(chan bool)

	stack := stacktrace.Get()

	go func() {
		time.Sleep(10 * time.Millisecond) // Ensure we don't start receiving output before the expectx rule has been set
		err := e.proc.Run()
		if err != nil {
			e.fail("Error while running process", "error: %v, stdout:\n---\n%s\n---\nstderr:\n---\n%s\n---\nstack:\n%s\n",
				err, e.proc.Stdout(), e.proc.Stderr(), stack.String())
		}
		e.procEnded <- true
	}()
}

func (e *Expectx) WaitForInput(timeout ...time.Duration) {
	usr, err := user.Current()
	if err != nil {
		e.fail("unexpected error: %s", err)
	}

	msg := "echo wait_ready_$HOME"
	if runtime.GOOS == "windows" {
		msg = "echo wait_ready_%USERPROFILE%"
	}

	e.Send(msg)
	e.Expect("wait_ready_"+usr.HomeDir, timeout...)
}

func (e *Expectx) Wait(timeout ...time.Duration) {
	t := 10 * time.Second
	if len(timeout) > 0 {
		t = timeout[0]
	}

	select {
	case <-e.procEnded:
		return
	case <-time.After(t):
		e.fail("Timed out while waiting for process to finish", "stdout:\n---\n%s\n---\nstderr:\n---\n%s\n---\n", e.proc.Stdout(), e.proc.Stderr())
	}
}

func (e *Expectx) Expect(value string, timeout ...time.Duration) {
	rx, err := regexp.Compile(regexp.QuoteMeta(value))
	if err != nil {
		e.fail("Value is not valid regex", "value: %s", regexp.QuoteMeta(value))
	}
	e.ExpectRe(rx, timeout...)
}

func (e *Expectx) ExpectExact(value string, timeout ...time.Duration) {
	rx, err := regexp.Compile("^" + regexp.QuoteMeta(value) + "$")
	if err != nil {
		e.fail("Value is not valid regex", "value: %s", regexp.QuoteMeta(value))
	}
	e.ExpectRe(rx, timeout...)
}

func (e *Expectx) ExpectRe(value *regexp.Regexp, timeout ...time.Duration) {
	t := 10 * time.Second
	if len(timeout) > 0 {
		t = timeout[0]
	}

	out := ""
	err := e.Timeout(func(stop chan bool) {
		e.proc.OnOutput(func(output []byte) {
			if value.MatchString(string(output)) {
				stop <- true
			}
			out = out + string(output)
		})
	}, t)
	if err != nil {
		e.fail("Could not meet expectation", "Expectation: '%s'\nError: %v\nstdout:\n---\n%s\n---\nstderr:\n---\n%s\n---\ncombined:\n---\n%s\n---\n",
			value.String(), err, e.proc.Stdout(), e.proc.Stderr(), e.proc.CombinedOutput())
	}
}

func (e *Expectx) Send(value string) {
	// Since we're not running a TTY emulator we need little workarounds like this to ensure stdin is ready
	time.Sleep(100 * time.Millisecond)

	err := e.proc.Write(value + "\n")
	if err != nil {
		e.fail("Could not send data to stdin", "error: %v", err)
	}
}

func (e *Expectx) SendQuit() {
	e.proc.Quit()
}

func (e *Expectx) Stop() {
	if e.proc == nil {
		e.fail("stop called without a spawned process")
	}
}

func (e *Expectx) Timeout(f func(stop chan bool), t time.Duration) error {
	stop := make(chan bool)
	go func() {
		f(stop)
	}()

	select {
	case <-stop:
		return nil
	case <-e.procEnded:
		return errors.New("Process ended")
	case <-time.After(t):
		return errors.New("Timeout reached")
	}
}
