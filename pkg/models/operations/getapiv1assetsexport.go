// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPIV1AssetsExportRequest struct {
}

type GetAPIV1AssetsExportResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetAPIV1AssetsExportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIV1AssetsExportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIV1AssetsExportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}