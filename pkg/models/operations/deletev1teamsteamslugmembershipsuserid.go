// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteV1TeamsTeamSlugMembershipsUserIDRequest struct {
	TeamSlug string `pathParam:"style=simple,explode=false,name=TeamSlug"`
	UserID   string `pathParam:"style=simple,explode=false,name=UserID"`
}

func (o *DeleteV1TeamsTeamSlugMembershipsUserIDRequest) GetTeamSlug() string {
	if o == nil {
		return ""
	}
	return o.TeamSlug
}

func (o *DeleteV1TeamsTeamSlugMembershipsUserIDRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type DeleteV1TeamsTeamSlugMembershipsUserIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteV1TeamsTeamSlugMembershipsUserIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteV1TeamsTeamSlugMembershipsUserIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteV1TeamsTeamSlugMembershipsUserIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
