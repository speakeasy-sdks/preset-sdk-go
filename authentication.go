// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package preset

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// authentication - API to authenticate and get a JWT token to interact with the Preset/Superset APIs.
type authentication struct {
	sdkConfiguration sdkConfiguration
}

func newAuthentication(sdkConfig sdkConfiguration) *authentication {
	return &authentication{
		sdkConfiguration: sdkConfig,
	}
}

// PostV1Auth - Get a JWT Token
// To interact with the Preset API, it's required to generate an API Key, that's used to generate a JWT token.
//
// 1. Generate an API Key on the Preset Manager UI. Refer to [our documentation](https://docs.preset.io/docs/the-preset-api).
// 2. Copy the API `token` and `secret`.
//
// Replace in the body:
//
// - `{{APIToken}}` with the `token` from the UI.
// - `{{APISecret}}` with the `secret` from the UI.
func (s *authentication) PostV1Auth(ctx context.Context, request operations.PostV1AuthRequest) (*operations.PostV1AuthResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/v1/auth/"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostV1AuthResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
