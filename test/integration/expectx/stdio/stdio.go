package stdio

type StdReader struct {
	onRead func(data []byte)
}

func NewStdReader() *StdReader {
	return &StdReader{}
}

func (w *StdReader) OnRead(cb func(data []byte)) {
	w.onRead = cb
}

func (w *StdReader) Read(p []byte) (n int, err error) {
	w.onRead(p)
	return len(p), nil
}

type StdWriter struct {
	onWrite func(data []byte)
}

func NewStdWriter() *StdWriter {
	return &StdWriter{}
}

func (w *StdWriter) OnWrite(cb func(data []byte)) {
	w.onWrite = cb
}

func (w *StdWriter) Write(p []byte) (n int, err error) {
	w.onWrite(p)
	return len(p), nil
}
