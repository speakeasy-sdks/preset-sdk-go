// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1DatasetRequestBody struct {
}

type PostAPIV1DatasetRequest struct {
	RequestBody *PostAPIV1DatasetRequestBody `request:"mediaType=application/json"`
}

func (o *PostAPIV1DatasetRequest) GetRequestBody() *PostAPIV1DatasetRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1DatasetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIV1DatasetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1DatasetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1DatasetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
