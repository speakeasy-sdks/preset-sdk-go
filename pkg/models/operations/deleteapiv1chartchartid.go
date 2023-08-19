// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAPIV1ChartChartIDRequest struct {
	ChartID string  `pathParam:"style=simple,explode=false,name=ChartID"`
	Referer *string `header:"style=simple,explode=false,name=Referer"`
}

func (o *DeleteAPIV1ChartChartIDRequest) GetChartID() string {
	if o == nil {
		return ""
	}
	return o.ChartID
}

func (o *DeleteAPIV1ChartChartIDRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

type DeleteAPIV1ChartChartIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteAPIV1ChartChartIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAPIV1ChartChartIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAPIV1ChartChartIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
