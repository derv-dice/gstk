package iof

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/jszwec/csvutil"
)

var (
	ErrWriterCSVNotInited = errors.New("WriterCSV not initialized. It's important to use NewWriterCSV()")
)

type WriterCSV struct {
	init       bool
	structType any

	out *os.File
	enc *csvutil.Encoder
	wr  *csv.Writer
}

func NewWriterCSV(filename string, structType any, headers bool) (wrCSV *WriterCSV, err error) {
	wrCSV = new(WriterCSV)
	if wrCSV.out, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return
	}

	wrCSV.structType = structType
	wrCSV.wr = csv.NewWriter(wrCSV.out)
	wrCSV.enc = csvutil.NewEncoder(wrCSV.wr)
	wrCSV.init = true

	if headers {
		if err = wrCSV.writeHeaders(); err != nil {
			return
		}
	}

	return
}

func (w *WriterCSV) writeHeaders() (err error) {
	if !w.init {
		return ErrWriterCSVNotInited
	}

	if err = w.enc.EncodeHeader(w.structType); err != nil {
		return
	}

	w.wr.Flush()
	return
}

func (w *WriterCSV) WriteRow(rows ...any) (err error) {
	if !w.init {
		return ErrWriterCSVNotInited
	}

	for i := range rows {
		if err = w.enc.Encode(rows[i]); err != nil {
			return
		}
	}

	w.wr.Flush()
	return
}

func (w *WriterCSV) Close() (err error) {
	if !w.init {
		return ErrWriterCSVNotInited
	}

	return w.out.Close()
}

type ReaderCSV struct{}

func NewReaderCSV() (r *ReaderCSV) {
	return &ReaderCSV{}
}

func (r *ReaderCSV) ReadCSV(filename string, dest any) (err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	return csvutil.Unmarshal(data, dest)
}

func MarshalCSV(v any) ([]byte, error) {
	return csvutil.Marshal(v)
}

func UnmarshalCSV(data []byte, dst any) error {
	return csvutil.Unmarshal(data, dst)
}
