// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1DashboardInfoRequest struct {
}

type GetAPIV1DashboardInfoResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetAPIV1DashboardInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIV1DashboardInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIV1DashboardInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}