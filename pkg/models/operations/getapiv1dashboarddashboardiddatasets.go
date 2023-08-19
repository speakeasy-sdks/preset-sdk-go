// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1DashboardDashboardIDDatasetsRequest struct {
	DashboardID string `pathParam:"style=simple,explode=false,name=DashboardID"`
}

func (o *GetAPIV1DashboardDashboardIDDatasetsRequest) GetDashboardID() string {
	if o == nil {
		return ""
	}
	return o.DashboardID
}

type GetAPIV1DashboardDashboardIDDatasetsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetAPIV1DashboardDashboardIDDatasetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIV1DashboardDashboardIDDatasetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIV1DashboardDashboardIDDatasetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}