package process

import (
	"os"
)

func safeClose(c chan struct{}) {
	select {
	case <-c:
	default:
		close(c)
	}
}

func safeCloseString(c chan string) {
	select {
	case _, ok := <-c:
		if ok {
			close(c)
		}
	default:
		close(c)
	}
}

type outSplit struct {
	f *os.File
	c chan string
	d chan struct{}
}

func newOutSplit(done chan struct{}, f *os.File, c chan string) *outSplit {
	s := outSplit{
		f: f,
		c: c,
		d: done,
	}

	return &s
}

func (s *outSplit) Write(p []byte) (int, error) {
	go func() {
		select {
		case s.c <- string(p):
		case <-s.d:
			return
		default:
			panic("exceeded buffer")
		}
	}()
	return s.f.Write(p)
}

func (s *outSplit) Read(p []byte) (int, error) {
	go func() {
		select {
		case s.c <- string(p):
		case <-s.d:
			return
		default:
			panic("exceeded buffer")
		}
	}()
	return s.f.Read(p)
}
