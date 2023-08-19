// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1ChartDataRequestBody struct {
}

type PostAPIV1ChartDataRequest struct {
	RequestBody *PostAPIV1ChartDataRequestBody `request:"mediaType=application/json"`
	// Flag to force refresh data.
	Force *string `queryParam:"style=form,explode=true,name=force"`
	// ID of the Chart to be refreshed.
	SliceID *string `queryParam:"style=form,explode=true,name=slice_id"`
}

func (o *PostAPIV1ChartDataRequest) GetRequestBody() *PostAPIV1ChartDataRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *PostAPIV1ChartDataRequest) GetForce() *string {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *PostAPIV1ChartDataRequest) GetSliceID() *string {
	if o == nil {
		return nil
	}
	return o.SliceID
}

type PostAPIV1ChartDataResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostAPIV1ChartDataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1ChartDataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1ChartDataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
