// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1DashboardRequestBody struct {
}

type PostAPIV1DashboardRequest struct {
	Referer     *string                        `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PostAPIV1DashboardRequestBody `request:"mediaType=application/json"`
}

func (o *PostAPIV1DashboardRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PostAPIV1DashboardRequest) GetRequestBody() *PostAPIV1DashboardRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1DashboardResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAPIV1DashboardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1DashboardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1DashboardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
