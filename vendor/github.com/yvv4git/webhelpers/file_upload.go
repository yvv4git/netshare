package webhelpers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// RequestUploadFile - function for upload file to server by http.
func RequestUploadFile(uri string, params map[string]string, paramName string, fileName string) (*http.Request, error) {
	// Open file.
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create bodyBuffer, writer...
	bodyBuffer := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuffer)
	part, err := writer.CreateFormFile(paramName, filepath.Base(fileName))
	if err != nil {
		return nil, err
	}

	// Add file as part of message.
	_, err = io.Copy(part, file)

	// Add keys as part of message.
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	// Close writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create request from body buffer.
	req, err := http.NewRequest("POST", uri, bodyBuffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
