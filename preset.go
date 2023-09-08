// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package preset

import (
	"fmt"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.app.preset.io",
	"https://manage.app.preset.io",
	"https://{{workspaceslug}}.{{workspaceregion}}.app.preset.io",
	"http://{{workspaceslug}}.{{workspaceregion}}.app.preset.io",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Preset API: Welcome to the Preset API Collection.
//
// ## Overview
//
// The Preset REST API is a powerful feature that provides access to core functional aspects of both Preset Manager and Superset.
//
// The API supports the following areas of Preset:
//
// *   User and Team Management
// *   Workspace Management
// *   Connection and Data Management
// *   Visualization Management
// *   Permissions Management \[Beta\]
//
// This documentation lists the most relevant endpoints across all of the above functional areas of Preset.
//
// ## Authentication
//
// ### Generate an API Key
//
// To generate an API key, from the Preset Manager screen (after log-in), hover your cursor over the initials icon and, in the sub-menu, select Manage User Settings.
//
// ![](https://i.ibb.co/D1xHc92/api1.png)
//
// In the *API Keys* section, select **\+ Generate New API Key**
//
// ![](https://i.ibb.co/LRWp7HC/api2.png)
//
// The *Generate a New API Key* panel appears.
//
// In the **Key Title** field, enter a name for the new API key.
//
// In the **Key Description** field, enter a brief descripton of the API key.
//
// Select **Submit**.
//
// ![](https://i.ibb.co/cC0H4mY/api3.png)
//
// The **Token** field will automatically populate with a generated token.
//
// Likewise, the **Secret** field will automatically populate with a secret.
//
// ![](https://i.ibb.co/8smp5pZ/api5.png)
//
// *Reminder: Safeguard the Secret**Please take a moment to select the Copy icon and then safely store it.*
//
// When ready, select **OK**.
//
// ![](https://i.ibb.co/LdNDGNp/api6.png)
//
// The newly-created API key appears in the *API Keys* section.
//
// By default, the API key will be activated — to deactivate, toggle the **Active** slider to the off position.
//
// To delete an API key, select the trash bin icon.
//
// ### Using the Postman Collection
//
// All requests on this collection inherit the **token** specified on the **Preset API** collection.
//
// To authenticate:
//
// 1.  Click on the **Preset API Collection**.
// 2.  Navigate to the **Variables** tab.
// 3.  Provide your **API Token** on the `APITokenName` current value.
// 4.  Provide the **API Token Secret** on the `APITokenSecret` current value.
//
// These would be used to generate a **JWT Token** that's stored as a **Global Variable**.
//
// There's a script defined on this collection, that is always executed before any request. The script basically checks if there's a valid **JWT Token** and tries to generate one/refresh it if needed.
//
// * * *
//
// ### Manually using the API
//
// Use the **Get a JWT Token** request to generate a `JWT Token.`
type Preset struct {
	// API to authenticate and get a JWT token to interact with the Preset/Superset APIs.
	Authentication *authentication
	// APIs associated with the Embedded functionality.
	PresetManagerAPIsGreaterThanEmbedded *presetManagerAPIsGreaterThanEmbedded
	// APIs to manage permissions on the Workspace level.
	//
	// Note that all Permission APIs require **Team Admin** permission.
	PresetManagerAPIsGreaterThanPermissions *presetManagerAPIsGreaterThanPermissions
	// APIs to manage your Preset team.
	PresetManagerAPIsGreaterThanTeams *presetManagerAPIsGreaterThanTeams
	// APIs to manage your Workspaces.
	PresetManagerAPIsGreaterThanWorkspaces *presetManagerAPIsGreaterThanWorkspaces
	// APIs to manage your Alerts & Reports.
	SupersetAPIsOpenSourceGreaterThanAlertsAndReports *supersetAPIsOpenSourceGreaterThanAlertsAndReports
	// API to manage your Annotation Layers.
	SupersetAPIsOpenSourceGreaterThanAnnotationLayers *supersetAPIsOpenSourceGreaterThanAnnotationLayers
	// APIs to export/import an `assets` ZIP file from the Workspace, which includes all:
	//
	// *   databases.
	// *   datasets.
	// *   charts.
	// *   saved queries.
	SupersetAPIsOpenSourceGreaterThanAssets *supersetAPIsOpenSourceGreaterThanAssets
	// APIs to manage Charts on your Workspace.
	SupersetAPIsOpenSourceGreaterThanCharts *supersetAPIsOpenSourceGreaterThanCharts
	// APIs to manage your Dashboards.
	SupersetAPIsOpenSourceGreaterThanDashboards *supersetAPIsOpenSourceGreaterThanDashboards
	// APIs to manage your database connections.
	SupersetAPIsOpenSourceGreaterThanDatabases *supersetAPIsOpenSourceGreaterThanDatabases
	// APIs to manage your datasets.
	SupersetAPIsOpenSourceGreaterThanDatasets *supersetAPIsOpenSourceGreaterThanDatasets
	SupersetAPIsOpenSourceGreaterThanQueries  *supersetAPIsOpenSourceGreaterThanQueries
	SupersetAPIsOpenSourceGreaterThanSQLLab   *supersetAPIsOpenSourceGreaterThanSQLLab

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Preset)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Preset) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Preset) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Preset) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Preset) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Preset) {
		sdk.sdkConfiguration.Security = &security
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Preset {
	sdk := &Preset{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.3.0",
			GenVersion:        "2.107.0",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Authentication = newAuthentication(sdk.sdkConfiguration)

	sdk.PresetManagerAPIsGreaterThanEmbedded = newPresetManagerAPIsGreaterThanEmbedded(sdk.sdkConfiguration)

	sdk.PresetManagerAPIsGreaterThanPermissions = newPresetManagerAPIsGreaterThanPermissions(sdk.sdkConfiguration)

	sdk.PresetManagerAPIsGreaterThanTeams = newPresetManagerAPIsGreaterThanTeams(sdk.sdkConfiguration)

	sdk.PresetManagerAPIsGreaterThanWorkspaces = newPresetManagerAPIsGreaterThanWorkspaces(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanAlertsAndReports = newSupersetAPIsOpenSourceGreaterThanAlertsAndReports(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanAnnotationLayers = newSupersetAPIsOpenSourceGreaterThanAnnotationLayers(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanAssets = newSupersetAPIsOpenSourceGreaterThanAssets(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanCharts = newSupersetAPIsOpenSourceGreaterThanCharts(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanDashboards = newSupersetAPIsOpenSourceGreaterThanDashboards(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanDatabases = newSupersetAPIsOpenSourceGreaterThanDatabases(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanDatasets = newSupersetAPIsOpenSourceGreaterThanDatasets(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanQueries = newSupersetAPIsOpenSourceGreaterThanQueries(sdk.sdkConfiguration)

	sdk.SupersetAPIsOpenSourceGreaterThanSQLLab = newSupersetAPIsOpenSourceGreaterThanSQLLab(sdk.sdkConfiguration)

	return sdk
}
