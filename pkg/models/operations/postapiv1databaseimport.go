// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostAPIV1DatabaseImportRequestBodyFormData struct {
	Content  []byte `multipartForm:"content"`
	FormData string `multipartForm:"name=formData"`
}

func (o *PostAPIV1DatabaseImportRequestBodyFormData) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *PostAPIV1DatabaseImportRequestBodyFormData) GetFormData() string {
	if o == nil {
		return ""
	}
	return o.FormData
}

type PostAPIV1DatabaseImportRequestBody struct {
	FormData  *PostAPIV1DatabaseImportRequestBodyFormData `multipartForm:"file"`
	Overwrite *bool                                       `multipartForm:"name=overwrite"`
	Passwords *string                                     `multipartForm:"name=passwords"`
}

func (o *PostAPIV1DatabaseImportRequestBody) GetFormData() *PostAPIV1DatabaseImportRequestBodyFormData {
	if o == nil {
		return nil
	}
	return o.FormData
}

func (o *PostAPIV1DatabaseImportRequestBody) GetOverwrite() *bool {
	if o == nil {
		return nil
	}
	return o.Overwrite
}

func (o *PostAPIV1DatabaseImportRequestBody) GetPasswords() *string {
	if o == nil {
		return nil
	}
	return o.Passwords
}

type PostAPIV1DatabaseImportRequest struct {
	Referer     *string                             `header:"style=simple,explode=false,name=Referer"`
	RequestBody *PostAPIV1DatabaseImportRequestBody `request:"mediaType=multipart/form-data"`
}

func (o *PostAPIV1DatabaseImportRequest) GetReferer() *string {
	if o == nil {
		return nil
	}
	return o.Referer
}

func (o *PostAPIV1DatabaseImportRequest) GetRequestBody() *PostAPIV1DatabaseImportRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostAPIV1DatabaseImportResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostAPIV1DatabaseImportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPIV1DatabaseImportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPIV1DatabaseImportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
