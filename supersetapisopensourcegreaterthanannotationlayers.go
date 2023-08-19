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

// supersetAPIsOpenSourceGreaterThanAnnotationLayers - API to manage your Annotation Layers.
type supersetAPIsOpenSourceGreaterThanAnnotationLayers struct {
	sdkConfiguration sdkConfiguration
}

func newSupersetAPIsOpenSourceGreaterThanAnnotationLayers(sdkConfig sdkConfiguration) *supersetAPIsOpenSourceGreaterThanAnnotationLayers {
	return &supersetAPIsOpenSourceGreaterThanAnnotationLayers{
		sdkConfiguration: sdkConfig,
	}
}

// DeleteAPIV1AnnotationLayer - Delete multiple Annotations Layers
// Deletes multiple Annotations from an Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
// - `{{AnnotationLayerIDs}` with comma separated Annotation Layer `ids` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) DeleteAPIV1AnnotationLayer(ctx context.Context, request operations.DeleteAPIV1AnnotationLayerRequest) (*operations.DeleteAPIV1AnnotationLayerResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/annotation_layer/"

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.DeleteAPIV1AnnotationLayerResponse{
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

// DeleteAPIV1AnnotationLayerAnnotationLayerID - Delete an Annotation Layer
// Deletes an Annotation Layer through the API.
//
// **Note:** You can only delete an Annotation Layer, after deleting **all its Annotations**.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) DeleteAPIV1AnnotationLayerAnnotationLayerID(ctx context.Context, request operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDRequest) (*operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDResponse{
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

// DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotation - Delete multiple Annotations from an Annotation Layer
// Deletes multiple Annotations from an Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
// - `{{AnnotationIDs}` with comma separated Annotation `ids` retrieved from the **Get all Annotations from an Annotation Layer** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotation(ctx context.Context, request operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest) (*operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}/annotation/", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse{
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

// DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID - Delete an Annotation from an Annotation Layer
// Deletes an Annotation from an Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
// - `{{AnnotationID}}` with the Annotation `id` retrieved from the **Get all Annotations from an Annotation Layer** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID(ctx context.Context, request operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest) (*operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}/annotation/{AnnotationID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse{
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

// GetAPIV1AnnotationLayer - Get all Annotation Layers from a Workspace
// Gets all Annotation Layers from a Workspace.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// Note that this endpoint returns 20 results by default. You can return up to 100 results at a time and use pagination by adding the following query parameters:
//
// ```
// ?q=(page_size:{{PageSize}},page:{{Page}})
//
// ```
//
// Replace:
//
// - `{{PageSize}}` with the desired size (min `1` max `100`).
// - `{{Page}}` with the page number (useful when the total count > `{{PageSize}}` - min `0`).
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) GetAPIV1AnnotationLayer(ctx context.Context) (*operations.GetAPIV1AnnotationLayerResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/annotation_layer/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.GetAPIV1AnnotationLayerResponse{
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

// GetAPIV1AnnotationLayerAnnotationLayerID - Get an Annotation Layer
// Gets a specific Annotation Layer from a Workspace.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) GetAPIV1AnnotationLayerAnnotationLayerID(ctx context.Context, request operations.GetAPIV1AnnotationLayerAnnotationLayerIDRequest) (*operations.GetAPIV1AnnotationLayerAnnotationLayerIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.GetAPIV1AnnotationLayerAnnotationLayerIDResponse{
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

// GetAPIV1AnnotationLayerAnnotationLayerIDAnnotation - Get all Annotations from an Annotation Layer
// Gets all Annotations from a specific Annotation Layer.
//
// Replace in the URL:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get All Annotation Layers from a Workspace** endpoint.
//
// Note that this endpoint returns 20 results by default. You can return up to 100 results at a time and use pagination by adding the following query parameters:
//
// ```
// ?q=(page_size:{{PageSize}},page:{{Page}})
//
// ```
//
// Replace:
//
// - `{{PageSize}}` with the desired size (min `1` max `100`).
// - `{{Page}}` with the page number (useful when the total count > `{{PageSize}}` - min `0`).
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) GetAPIV1AnnotationLayerAnnotationLayerIDAnnotation(ctx context.Context, request operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest) (*operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}/annotation", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse{
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

// GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID - Get an Annotation from an Annotation Layer
// Deletes an Annotation Layer through the API.
//
// **Note:** You can only delete an Annotation Layer, after deleting **all its Annotations**.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
// - `{{AnnotationID}}` with the Annotation `id` retrieved from the **Get all Annotations from an Annotation Layer** endpoint.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID(ctx context.Context, request operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest) (*operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}/annotation/{AnnotationID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse{
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

// PostAPIV1AnnotationLayer - Create an Annotation Layer
// Creates an Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// Replace in the body:
//
// - `{{AnnotationLayerDescription}}` with a description for the Annotation Layer.
// - `{{AnnotationLayerName}}` with the desired new name.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) PostAPIV1AnnotationLayer(ctx context.Context, request operations.PostAPIV1AnnotationLayerRequest) (*operations.PostAPIV1AnnotationLayerResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/annotation_layer/"

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

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.PostAPIV1AnnotationLayerResponse{
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

// PostAPIV1AnnotationLayerAnnotationLayerIDAnnotation - Create an Annotation in an Annotation Layer
// Creates an Annotation on an existing Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
//
// Replace in the body:
//
// - `{{AnnotationEndDTTM}}` with the annotation's datetime end (`YYYY-MM-DD HH:MM`).
// - `{{AnnotationLongDescription}}` with the annotation's description.
// - `{{AnnotationTitle}}` with a name for it.
// - `{{AnnotationStartDTTM}}` with the annotation's datetime start (`YYYY-MM-DD HH:MM`).
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) PostAPIV1AnnotationLayerAnnotationLayerIDAnnotation(ctx context.Context, request operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest) (*operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}/annotation", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse{
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

// PutAPIV1AnnotationLayerAnnotationLayerID - Update an Annotation Layer
// Updates an Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
//
// Replace in the body:
//
// - `{{AnnotationLayerDescription}}` with a description for the Annotation Layer.
// - `{{AnnotationLayerName}}` with the desired new name.
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) PutAPIV1AnnotationLayerAnnotationLayerID(ctx context.Context, request operations.PutAPIV1AnnotationLayerAnnotationLayerIDRequest) (*operations.PutAPIV1AnnotationLayerAnnotationLayerIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.PutAPIV1AnnotationLayerAnnotationLayerIDResponse{
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

// PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID - Update an Annotation from an Annotation Layer
// Updates an Annotation from an Annotation Layer through the API.
//
// Replace in the URL and on the `Referer` header:
//
// - `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
// - `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
//
// | **`region`** | **`WorkspaceRegion`** |
// | --- | --- |
// | us-east-1 | `us2a` |
// | us-west-2 | `us1a` |
// | eu-north-1 | `eu5a` |
// | ap-northeast-1 | `ap1a` |
//
// Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.
//
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
// - `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
//
// Replace in the body:
//
// - `{{AnnotationEndDTTM}}` with the annotation's datetime end (`YYYY-MM-DD HH:MM`).
// - `{{AnnotationLongDescription}}` with the annotation's description.
// - `{{AnnotationShortDescription}}` with a name for it.
// - `{{AnnotationStartDTTM}}` with the annotation's datetime start (`YYYY-MM-DD HH:MM`).
func (s *supersetAPIsOpenSourceGreaterThanAnnotationLayers) PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID(ctx context.Context, request operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest) (*operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v1/annotation_layer/{AnnotationLayerID}/annotation/{AnnotationID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse{
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
