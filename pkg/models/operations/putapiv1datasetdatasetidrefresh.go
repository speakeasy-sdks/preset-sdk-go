// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIV1DatasetDatasetIDRefreshRequestBody struct {
}

type PutAPIV1DatasetDatasetIDRefreshRequest struct {
	DatasetID   string                                      `pathParam:"style=simple,explode=false,name=DatasetID"`
	Referer     *string                                     `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PutAPIV1DatasetDatasetIDRefreshRequestBody `request:"mediaType=application/json"`
}

func (o *PutAPIV1DatasetDatasetIDRefreshRequest) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

func (o *PutAPIV1DatasetDatasetIDRefreshRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PutAPIV1DatasetDatasetIDRefreshRequest) GetRequestBody() *PutAPIV1DatasetDatasetIDRefreshRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PutAPIV1DatasetDatasetIDRefreshResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *PutAPIV1DatasetDatasetIDRefreshResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAPIV1DatasetDatasetIDRefreshResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAPIV1DatasetDatasetIDRefreshResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}