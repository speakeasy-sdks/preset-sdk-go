// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package presetsdkgo

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

type supersetAPIsOpenSourceGreaterThanSQLLab struct {
	sdkConfiguration sdkConfiguration
}

func newSupersetAPIsOpenSourceGreaterThanSQLLab(sdkConfig sdkConfiguration) *supersetAPIsOpenSourceGreaterThanSQLLab {
	return &supersetAPIsOpenSourceGreaterThanSQLLab{
		sdkConfiguration: sdkConfig,
	}
}

// PostAPIV1SqllabExecute - Execute a SQL Query
// Executes a SQL query through the API.
//
// Replace in the URL and in the `Referer` header:
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
// - `{{DBID}}` with the `id` retrieved from the **Get all Database Connections from a Workspace** endpoint.
// - `{{SQLQuery}}` with the desired SQL query to be executed. Don't forget to escape quotes with `\`. For example:
//
// ``` json
//
//	{
//	    "database_id": 1,
//	    "sql": "SELECT * FROM \"Vehicle Sales\" limit 7"
//	  }
//
// ```
func (s *supersetAPIsOpenSourceGreaterThanSQLLab) PostAPIV1SqllabExecute(ctx context.Context, request operations.PostAPIV1SqllabExecuteRequest) (*operations.PostAPIV1SqllabExecuteResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/sqllab/execute/"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIV1SqllabExecuteResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
