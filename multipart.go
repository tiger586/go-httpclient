package httpclient

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

// NewMultipart 建立 multipart body
// fields: form fields, files: key -> file path
func NewMultipart(fields map[string]string, files map[string]string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range fields {
		_ = writer.WriteField(k, v)
	}

	for key, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, "", err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(key, path)
		if err != nil {
			return nil, "", err
		}
		_, _ = io.Copy(part, file)
	}

	writer.Close()
	return body, writer.FormDataContentType(), nil
}
