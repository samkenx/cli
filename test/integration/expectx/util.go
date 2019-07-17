package expectx

import (
	"bufio"
	"bytes"
	"io"
	"time"
)

func drain(c <-chan string) string {
	var s string
	for {
		select {
		case v, ok := <-c:
			if !ok {
				return s
			}
			s += v + "\n"
		case <-time.After(time.Millisecond * 100):
			return s
		}
	}
}

func pipeChan(done chan struct{}, name string) (*io.PipeWriter, chan string) {
	r, w := io.Pipe()
	c := make(chan string, 40960)
	bridge := make(chan string)

	go func() {
		defer close(c)
		defer w.Close()
		defer r.Close()
		defer close(bridge)

		sc := bufio.NewScanner(r)
		sc.Split(scanLines)
		for sc.Scan() {
			select {
			case c <- sc.Text():
			case <-done:
				return
			default:
				panic("exceeded buffer")
			}
		}
		if err := sc.Err(); err != nil {
			panic(err)
		}
	}()

	return w, c
}

func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// check for escape sequence
	if len(data) > 4 && bytes.HasSuffix(data, []byte{27, 91, 48, 109}) {
		return len(data), dropCR(data), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
