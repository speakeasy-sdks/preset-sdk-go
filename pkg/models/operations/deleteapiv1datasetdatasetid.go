// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteAPIV1DatasetDatasetIDRequest struct {
	DatasetID string `pathParam:"style=simple,explode=false,name=DatasetID"`
}

func (o *DeleteAPIV1DatasetDatasetIDRequest) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

type DeleteAPIV1DatasetDatasetIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteAPIV1DatasetDatasetIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAPIV1DatasetDatasetIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAPIV1DatasetDatasetIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
