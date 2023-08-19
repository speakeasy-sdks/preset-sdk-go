// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1DatabaseRequestBody struct {
}

type PostAPIV1DatabaseRequest struct {
	Referer     *string                       `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PostAPIV1DatabaseRequestBody `request:"mediaType=application/json"`
}

func (o *PostAPIV1DatabaseRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PostAPIV1DatabaseRequest) GetRequestBody() *PostAPIV1DatabaseRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1DatabaseResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostAPIV1DatabaseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1DatabaseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1DatabaseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
