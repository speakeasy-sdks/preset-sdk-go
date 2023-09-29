// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIV1DatabaseDatabaseIDRequest struct {
	DatabaseID  string  `pathParam:"style=simple,explode=false,name=DatabaseID"`
	RequestBody *string `request:"mediaType=*/*"`
}

func (o *PutAPIV1DatabaseDatabaseIDRequest) GetDatabaseID() string {
	if o == nil {
		return ""
	}
	return o.DatabaseID
}

func (o *PutAPIV1DatabaseDatabaseIDRequest) GetRequestBody() *string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PutAPIV1DatabaseDatabaseIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutAPIV1DatabaseDatabaseIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPIV1DatabaseDatabaseIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPIV1DatabaseDatabaseIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
