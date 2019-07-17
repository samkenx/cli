package process

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/kr/pty"
)

type Process struct {
	c  *exec.Cmd
	wc chan struct{} // wait
}

func New(c *exec.Cmd, sin io.Reader, sout, serr io.Writer) (*Process, error) {
	we := func(err error) error {
		if err == nil {
			return nil
		}
		return fmt.Errorf("cannot construct new process: %s", err)
	}

	wc := make(chan struct{})

	p, t, err := pty.Open()
	if err != nil {
		return nil, we(err)
	}
	cleanup := func() {
		defer safeClose(wc)
		//_ = p.Close()
		//_ = t.Close()
	}

	if c.Stdout == nil {
		//c.Stdout = io.MultiWriter(t, sout)
		c.Stdout = t
	}
	if c.Stderr == nil {
		//c.Stderr = io.MultiWriter(t, serr)
		c.Stderr = t
	}
	if c.Stdin == nil {
		c.Stdin = t
	}
	if c.SysProcAttr == nil {
		c.SysProcAttr = &syscall.SysProcAttr{}
	}
	c.SysProcAttr.Setctty = true
	c.SysProcAttr.Setsid = true

	if err := c.Start(); err != nil {
		cleanup()
		return nil, we(err)
	}

	go func() {
		defer cleanup()
		c.Wait()
	}()

	go func() {
		_, _ = io.Copy(sout, p)
	}()

	go func() {
		_, err := io.Copy(p, sin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	x := Process{
		c:  c,
		wc: wc,
	}

	return &x, nil
}

func (p *Process) Wait() chan struct{} {
	return p.wc
}

func (p *Process) Interrupt() {
	p.c.Process.Signal(syscall.SIGINT)
}
