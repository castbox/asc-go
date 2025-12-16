/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
)

// ErrMissingChunkBounds happens when the UploadOperation object is missing an offset or length used to mark
// what bytes in the Reader will be uploaded.
var ErrMissingChunkBounds = errors.New("could not establish bounds of upload operation")

// ErrMissingUploadDestination happens when the UploadOperation object is missing a URL or HTTP method.
var ErrMissingUploadDestination = errors.New("could not establish destination of upload operation")

// UploadOperation defines model for UploadOperation.
//
// https://developer.apple.com/documentation/appstoreconnectapi/uploadoperation
// https://developer.apple.com/documentation/appstoreconnectapi/uploading_assets_to_app_store_connect
type UploadOperation struct {
	Length         *int                    `json:"length,omitempty"`
	Method         *string                 `json:"method,omitempty"`
	Offset         *int                    `json:"offset,omitempty"`
	RequestHeaders []UploadOperationHeader `json:"requestHeaders,omitempty"`
	URL            *string                 `json:"url,omitempty"`
}

// UploadOperationHeader defines model for UploadOperationHeader.
//
// https://developer.apple.com/documentation/appstoreconnectapi/uploadoperationheader
type UploadOperationHeader struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// UploadOperationError pairs a failed operation and its associated error so it
// can be retried later.
type UploadOperationError struct {
	Operation UploadOperation
	Err       error
}

func (e UploadOperationError) Error() string {
	return e.Err.Error()
}

// chunk returns the bytes in the file from the given offset and with the given length.
func (op *UploadOperation) chunk(f io.ReadSeeker) (*bytes.Buffer, error) {
	if op.Offset == nil || op.Length == nil {
		return nil, ErrMissingChunkBounds
	}

	_, err := f.Seek(int64(*op.Offset), 0)
	if err != nil {
		return nil, err
	}

	data := make([]byte, *op.Length)

	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(data), nil
}

// request creates a new http.request instance from the given UploadOperation and buffer.
func (op *UploadOperation) request(ctx context.Context, data *bytes.Buffer) (*http.Request, error) {
	if op.Method == nil || op.URL == nil {
		return nil, ErrMissingUploadDestination
	}

	// Create request with the data bytes directly
	bodyBytes := data.Bytes()
	req, err := http.NewRequestWithContext(ctx, *op.Method, *op.URL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	// Set Content-Length explicitly - required for CDN uploads
	req.ContentLength = int64(len(bodyBytes))

	// Set headers from upload operation
	if op.RequestHeaders != nil {
		for _, h := range op.RequestHeaders {
			if h.Name == nil || h.Value == nil {
				continue
			}

			req.Header.Add(*h.Name, *h.Value)
		}
	}

	return req, nil
}

// Upload takes a file path and concurrently uploads each part of the file to App Store Connect.
func (c *Client) Upload(ctx context.Context, ops []UploadOperation, file io.ReadSeeker) error {
	var wg sync.WaitGroup

	// Use buffered channel to avoid blocking
	errs := make(chan UploadOperationError, len(ops))

	// Create a single HTTP client for all chunk uploads to reuse connection pool
	// Use a plain HTTP client for CDN uploads - no auth headers needed
	// The CDN URL already contains signed parameters
	plainClient := &http.Client{}

	for i, operation := range ops {
		chunk, err := operation.chunk(file)
		if err != nil {
			errs <- UploadOperationError{
				Operation: operation,
				Err:       err,
			}

			continue
		}

		wg.Add(1)

		go c.uploadChunk(ctx, ops[i], chunk, plainClient, errs, &wg)
	}

	// Wait for all uploads to complete, then close the channel
	go func() {
		wg.Wait()
		close(errs)
	}()

	// Collect all errors
	var firstErr error
	for err := range errs {
		if firstErr == nil {
			firstErr = err
		}
	}

	return firstErr
}

func (c *Client) uploadChunk(ctx context.Context, op UploadOperation, chunk *bytes.Buffer, client *http.Client, errs chan<- UploadOperationError, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := op.request(ctx, chunk)
	if err != nil {
		errs <- UploadOperationError{
			Operation: op,
			Err:       err,
		}

		return
	}

	resp, err := client.Do(req)
	if err != nil {
		errs <- UploadOperationError{
			Operation: op,
			Err:       err,
		}
		return
	}
	defer resp.Body.Close()

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		errs <- UploadOperationError{
			Operation: op,
			Err:       errors.New("upload failed with status " + resp.Status + ": " + string(body)),
		}
		return
	}
}
