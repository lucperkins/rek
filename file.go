package rek

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const defaultFileFieldName = "file"

type file struct {
	FieldName string
	Filepath  string
	Params    map[string]string
}

func (f *file) build() *file {
	if f.FieldName == "" {
		f.FieldName = defaultFileFieldName
	}

	return f
}

func buildMultipartBody(opts *options) (io.Reader, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	f := opts.file.build()

	data, err := os.Open(f.Filepath)
	if err != nil {
		return nil, "", err
	}

	part, err := writer.CreateFormFile(f.FieldName, filepath.Base(f.Filepath))
	if err != nil {
		return nil, "", err
	}

	if _, err := io.Copy(part, data); err != nil {
		return nil, "", err
	}

	if f.Params != nil {
		for k, v := range f.Params {
			if err := writer.WriteField(k, v); err != nil {
				return nil, "", err
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}
