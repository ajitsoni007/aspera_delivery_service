package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUploader(t *testing.T) {

	e := echo.New()
	testCases := UploaderTestCases

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/upload", strings.NewReader(tc.Form.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := Uploader(c)

		assert.NoError(t, err)

		assert.Equal(t, tc.ExpectedStatus, rec.Code)

		assert.Equal(t, tc.ExpectedResponseBody, rec.Body.String())

	}

}

// package handlers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUploader(t *testing.T) {
// 	e := echo.New()

// 	// Test Case 1: Successful upload with password authentication
// 	req1 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("username=user&password=pass&host=example.com&port=22&auth_type=password&source_pathS3URI=s3://example-bucket/source-file&destination_path=/destination-path"))
// 	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
// 	rec1 := httptest.NewRecorder()
// 	c1 := e.NewContext(req1, rec1)

// 	_,err := returnUploadEventInfo(c1)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec1.Code)

// 	// Test Case 2: Successful upload with private key authentication
// 	req2 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("username=user&keyS3URI=s3://example-bucket/key-file&host=example.com&port=22&auth_type=privatekey&source_pathS3URI=s3://example-bucket/source-file&destination_path=/destination-path"))
// 	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
// 	rec2 := httptest.NewRecorder()
// 	c2 := e.NewContext(req2, rec2)

// 	_,err = returnUploadEventInfo(c2)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec2.Code)

// 	// Test Case 3: Failed download of source file
// 	req3 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("username=user&password=pass&host=example.com&port=22&auth_type=password&source_pathS3URI=s3://example-bucket/invalid-path&destination_path=/destination-path"))
// 	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
// 	rec3 := httptest.NewRecorder()
// 	c3 := e.NewContext(req3, rec3)

// 	_,err = returnUploadEventInfo(c3)
// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, rec3.Code)
// }

// // func TestUploader(t *testing.T) {
// //     // Create a new Echo instance
// //     e := echo.New()

// //     // Create a new HTTP request to test the uploader handler
// //     req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
// //     req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

// //     // Create a new HTTP response recorder
// //     rec := httptest.NewRecorder()

// //     // Create a new Echo context from the HTTP request and response recorder
// //     c := e.NewContext(req, rec)

// //     // Call the uploader handler function
// //     if err := Uploader(c); err != nil {
// //         t.Errorf("error calling Uploader handler: %v", err)
// //         return
// //     }

// //     // Check the status code of the HTTP response
// //     if rec.Code != http.StatusOK {
// //         t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
// //         return
// //     }

// //     // Check the response body of the HTTP response
// //     expectedBody := `{"fileName":"","sourcePath":"","destinationPath":"","outputResult":""}`
// //     if rec.Body.String() != expectedBody {
// //         t.Errorf("expected response body %q, but got %q", expectedBody, rec.Body.String())
// //         return
// //     }
// // }

// // func TestReturnUploadEventInfo(t *testing.T) {
// //     // Create a new Echo instance
// //     e := echo.New()

// //     // Create a new HTTP request to test the returnUploadEventInfo function
// //     req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
// //     req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
// //     q := req.URL.Query()
// //     q.Add("username", "tesasp2657331_1")
// //     q.Add("host", "aspera.kaltura.com")
// //     q.Add("port", "22")
// //     q.Add("auth_type", "password")
// //     q.Add("destination_path", "/")
// // 	q.Add("password", "2qcttTOp")
// //     q.Add("source_pathS3URI", "s3://aspera-delivery-temp/test.txt")
// //     req.URL.RawQuery = q.Encode()

// //     // Create a new Echo context from the HTTP request
// //     c := e.NewContext(req, nil)

// //     // Call the returnUploadEventInfo function
// //     resp, err := returnUploadEventInfo(c)
// //     if err != nil {
// //         t.Errorf("error calling returnUploadEventInfo function: %v", err)
// //         return
// //     }

// //     // Check the response fields of the UploadEventResponse
// //     if resp.FileName == "" {
// //         t.Error("expected non-empty file name, but got empty string")
// //         return
// //     }
// //     if resp.SourcePath == "" {
// //         t.Error("expected non-empty source path, but got empty string")
// //         return
// //     }
// //     if resp.DestinationPath != "/" {
// //         t.Errorf("expected destination path %q, but got %q", "/tmp/", resp.DestinationPath)
// //         return
// //     }
// //     if resp.OutputResult == "" {
// //         t.Error("expected non-empty output result, but got empty string")
// //         return
// //     }
// // }
