// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1DashboardDashboardIDRequest struct {
	DashboardID string `pathParam:"style=simple,explode=false,name=DashboardID"`
}

func (o *GetAPIV1DashboardDashboardIDRequest) GetDashboardID() string {
	if o == nil {
		return ""
	}
	return o.DashboardID
}

type GetAPIV1DashboardDashboardIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPIV1DashboardDashboardIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIV1DashboardDashboardIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIV1DashboardDashboardIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
