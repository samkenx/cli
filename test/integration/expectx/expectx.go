package expectx

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/ActiveState/cli/test/integration/expectx/process"
)

type Expectx struct {
	fail func(string, ...interface{})
	proc *process.Process
	done chan struct{}
	env  []string
	exec string
	sinr *io.PipeReader
	sinw *io.PipeWriter
	sout *io.PipeWriter
	serr *io.PipeWriter
	outc chan string
	errc chan string
}

func New(exec string, fail func(string, ...interface{})) *Expectx {
	done := make(chan struct{})
	sinr, sinw := io.Pipe()
	sout, outc := pipeChan(done, "sout")
	serr, errc := pipeChan(done, "serr")

	return &Expectx{
		fail: fail,
		exec: exec,
		done: done,
		sinr: sinr,
		sinw: sinw,
		sout: sout,
		serr: serr,
		outc: outc,
		errc: errc,
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

	cmd := exec.Command(e.exec, args...)
	p, err := process.New(cmd, e.sinr, e.sout, e.serr)
	if err != nil {
		panic(err)
	}
	e.proc = p
	//e.proc.SetEnv(e.env)
	//e.procEnded = make(chan bool)

	//stack := stacktrace.Get()
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
	case <-e.proc.Wait():
		return
	case <-time.After(t):
		e.fail("Timed out while waiting for process to finish", "stdout:\n---\n%s\n---\nstderr:\n---\n%s\n---\n", drain(e.outc), drain(e.errc))
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

	var sout, serr string
	var ok bool
	var err error
	select {
	case sout, ok = <-e.outc:
		if !ok {
			err = errors.New("stdout closed")
			break
		}
		if !value.MatchString(sout) {
			err = errors.New("stdout does not match")
		}
	case serr, ok = <-e.errc:
		if !ok {
			err = errors.New("stderr closed")
			break
		}
		if !value.MatchString(serr) {
			err = errors.New("stderr does not match")
		}
	case <-time.After(t):
		err = errors.New("timeout reached")
	}
	if err != nil {
		e.fail(
			"Could not meet expectation", "Expectation: '%s'\nError: %v\nstdout:\n---\n%s\n---\nstderr:\n---\n%s\n---\n",
			value.String(), err, sout, serr,
		)
	}
}

func (e *Expectx) Send(value string) {
	if _, err := e.sinw.Write([]byte(value + "\n")); err != nil {
		e.fail("Could not send data to stdin", "error: %v", err)
	}
	<-e.outc
}

func (e *Expectx) SendInterrupt() {
	e.proc.Interrupt()
}

func (e *Expectx) Close() {
	close(e.done)
	if e.proc == nil {
		e.fail("stop called without a spawned process")
	}
	// TODO: ?
}
