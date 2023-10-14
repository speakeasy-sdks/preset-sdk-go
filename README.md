# github.com/speakeasy-sdks/preset-sdk-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/preset-sdk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	presetsdkgo "github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := presetsdkgo.New(
		presetsdkgo.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Authentication.PostV1Auth(ctx, operations.PostV1AuthRequest{
		RequestBody: &operations.PostV1AuthRequestBody{},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Authentication](docs/sdks/authentication/README.md)

* [PostV1Auth](docs/sdks/authentication/README.md#postv1auth) - Get a JWT Token

### [PresetManagerAPIsGreaterThanEmbedded](docs/sdks/presetmanagerapisgreaterthanembedded/README.md)

* [PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestToken](docs/sdks/presetmanagerapisgreaterthanembedded/README.md#postapiv1teamsteamslugworkspacesworkspaceslugguesttoken) - Create a new Guest Token

### [PresetManagerAPIsGreaterThanPermissions](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md)

* [DeleteV1TeamsTeamSlugPermissionsPermissionName](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#deletev1teamsteamslugpermissionspermissionname) - Delete Row Level Security
* [DeleteV1TeamsTeamSlugPermissionsPermissionNameGrantees](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#deletev1teamsteamslugpermissionspermissionnamegrantees) - Delete Grantee from existing Permission
* [GetV1TeamsTeamSlugPermissions](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#getv1teamsteamslugpermissions) - Get Permissions
* [GetV1TeamsTeamSlugPermissionsResources](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#getv1teamsteamslugpermissionsresources) - Get Resources
* [PatchV1TeamsTeamSlugPermissionsPermissionName](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#patchv1teamsteamslugpermissionspermissionname) - Update Row Level Security
* [PostV1TeamsTeamSlugPermissions](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#postv1teamsteamslugpermissions) - Create Row Level Security
* [PostV1TeamsTeamSlugPermissionsPermissionNameGrantees](docs/sdks/presetmanagerapisgreaterthanpermissions/README.md#postv1teamsteamslugpermissionspermissionnamegrantees) - Add Grantees to existing Permission

### [PresetManagerAPIsGreaterThanTeams](docs/sdks/presetmanagerapisgreaterthanteams/README.md)

* [DeleteV1TeamsTeamSlugInvitesInviteID](docs/sdks/presetmanagerapisgreaterthanteams/README.md#deletev1teamsteamsluginvitesinviteid) - Delete Pending Invite
* [DeleteV1TeamsTeamSlugMembershipsUserID](docs/sdks/presetmanagerapisgreaterthanteams/README.md#deletev1teamsteamslugmembershipsuserid) - Delete a Team Member
* [GetV1Teams](docs/sdks/presetmanagerapisgreaterthanteams/README.md#getv1teams) - Get Preset Teams
* [GetV1TeamsTeamSlugInvites](docs/sdks/presetmanagerapisgreaterthanteams/README.md#getv1teamsteamsluginvites) - Get Pending Team Invites
* [GetV1TeamsTeamSlugMemberships](docs/sdks/presetmanagerapisgreaterthanteams/README.md#getv1teamsteamslugmemberships) - Get Team Members
* [PatchV1TeamsTeamSlugMembershipsUserID](docs/sdks/presetmanagerapisgreaterthanteams/README.md#patchv1teamsteamslugmembershipsuserid) - Change User Role
* [PostV1TeamsTeamSlugInvites](docs/sdks/presetmanagerapisgreaterthanteams/README.md#postv1teamsteamsluginvites) - Create a Team Invite
* [PostV1TeamsTeamSlugInvitesMany](docs/sdks/presetmanagerapisgreaterthanteams/README.md#postv1teamsteamsluginvitesmany) - Create Multiple Team Invites
* [PostV1TeamsTeamSlugInvitesResendInviteID](docs/sdks/presetmanagerapisgreaterthanteams/README.md#postv1teamsteamsluginvitesresendinviteid) - Resend Invite
* [PutV1TeamsTeamSlug](docs/sdks/presetmanagerapisgreaterthanteams/README.md#putv1teamsteamslug) - Update Team Title

### [PresetManagerAPIsGreaterThanWorkspaces](docs/sdks/presetmanagerapisgreaterthanworkspaces/README.md)

* [GetV1TeamsTeamSlugWorkspaces](docs/sdks/presetmanagerapisgreaterthanworkspaces/README.md#getv1teamsteamslugworkspaces) - Get Workspaces from a Team
* [GetV1TeamsTeamSlugWorkspacesWorkspaceIDMemberships](docs/sdks/presetmanagerapisgreaterthanworkspaces/README.md#getv1teamsteamslugworkspacesworkspaceidmemberships) - Get Workspace Users and Roles
* [PostV1TeamsTeamSlugWorkspaces](docs/sdks/presetmanagerapisgreaterthanworkspaces/README.md#postv1teamsteamslugworkspaces) - Create Workspace for a Team
* [PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembership](docs/sdks/presetmanagerapisgreaterthanworkspaces/README.md#putv1teamsteamslugworkspacesworkspaceidmembership) - Change Workspace Role

### [SupersetAPIsOpenSourceGreaterThanAlertsAndReports](docs/sdks/supersetapisopensourcegreaterthanalertsandreports/README.md)

* [GetAPIV1Report](docs/sdks/supersetapisopensourcegreaterthanalertsandreports/README.md#getapiv1report) - Get all Reports from a Workspace
* [GetAPIV1ReportInfo](docs/sdks/supersetapisopensourcegreaterthanalertsandreports/README.md#getapiv1reportinfo) - Get Alerts & Reports API metadata Info
* [GetAPIV1ReportAlertIDORReportID](docs/sdks/supersetapisopensourcegreaterthanalertsandreports/README.md#getapiv1reportalertidorreportid) - Get an Alert/Report
* [PostAPIV1Report](docs/sdks/supersetapisopensourcegreaterthanalertsandreports/README.md#postapiv1report) - Create a Dashboard Alert
* [PutAPIV1ReportAlertIDORReportID](docs/sdks/supersetapisopensourcegreaterthanalertsandreports/README.md#putapiv1reportalertidorreportid) - Disable an Alert/Report

### [SupersetAPIsOpenSourceGreaterThanAnnotationLayers](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md)

* [DeleteAPIV1AnnotationLayer](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#deleteapiv1annotationlayer) - Delete multiple Annotations Layers
* [DeleteAPIV1AnnotationLayerAnnotationLayerID](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#deleteapiv1annotationlayerannotationlayerid) - Delete an Annotation Layer
* [DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotation](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#deleteapiv1annotationlayerannotationlayeridannotation) - Delete multiple Annotations from an Annotation Layer
* [DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#deleteapiv1annotationlayerannotationlayeridannotationannotationid) - Delete an Annotation from an Annotation Layer
* [GetAPIV1AnnotationLayer](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#getapiv1annotationlayer) - Get all Annotation Layers from a Workspace
* [GetAPIV1AnnotationLayerAnnotationLayerID](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#getapiv1annotationlayerannotationlayerid) - Get an Annotation Layer
* [GetAPIV1AnnotationLayerAnnotationLayerIDAnnotation](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#getapiv1annotationlayerannotationlayeridannotation) - Get all Annotations from an Annotation Layer
* [GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#getapiv1annotationlayerannotationlayeridannotationannotationid) - Get an Annotation from an Annotation Layer
* [PostAPIV1AnnotationLayer](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#postapiv1annotationlayer) - Create an Annotation Layer
* [PostAPIV1AnnotationLayerAnnotationLayerIDAnnotation](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#postapiv1annotationlayerannotationlayeridannotation) - Create an Annotation in an Annotation Layer
* [PutAPIV1AnnotationLayerAnnotationLayerID](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#putapiv1annotationlayerannotationlayerid) - Update an Annotation Layer
* [PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID](docs/sdks/supersetapisopensourcegreaterthanannotationlayers/README.md#putapiv1annotationlayerannotationlayeridannotationannotationid) - Update an Annotation from an Annotation Layer

### [SupersetAPIsOpenSourceGreaterThanAssets](docs/sdks/supersetapisopensourcegreaterthanassets/README.md)

* [GetAPIV1AssetsExport](docs/sdks/supersetapisopensourcegreaterthanassets/README.md#getapiv1assetsexport) - Export Assets
* [PostAPIV1AssetsImport](docs/sdks/supersetapisopensourcegreaterthanassets/README.md#postapiv1assetsimport) - Import Assets

### [SupersetAPIsOpenSourceGreaterThanCharts](docs/sdks/supersetapisopensourcegreaterthancharts/README.md)

* [DeleteAPIV1ChartChartID](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#deleteapiv1chartchartid) - Delete a Chart
* [GetAPIV1Chart](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#getapiv1chart) - Get all Charts from a Workspace
* [GetAPIV1ChartExport](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#getapiv1chartexport) - Export Charts
* [GetAPIV1ChartRelatedOwners](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#getapiv1chartrelatedowners) - Get all possible Chart Owners
* [GetAPIV1ChartChartID](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#getapiv1chartchartid) - Get a Chart
* [GetAPIV1ChartChartIDCacheScreenshot](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#getapiv1chartchartidcachescreenshot) - Get Chart Screenshot
* [GetAPIV1ChartChartIDData](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#getapiv1chartchartiddata) - Get Chart's Data
* [PostAPIV1Chart](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#postapiv1chart) - Create a Chart
* [PostAPIV1ChartData](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#postapiv1chartdata) - Refresh a Chart
* [PostAPIV1ChartImport](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#postapiv1chartimport) - Import a Chart
* [PutAPIV1ChartChartID](docs/sdks/supersetapisopensourcegreaterthancharts/README.md#putapiv1chartchartid) - Update a Chart

### [SupersetAPIsOpenSourceGreaterThanDashboards](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md)

* [DeleteAPIV1DashboardDashboardID](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#deleteapiv1dashboarddashboardid) - Delete Dashboard
* [GetAPIV1Dashboard](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#getapiv1dashboard) - Get all Dashboards From a Workspace
* [GetAPIV1DashboardInfo](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#getapiv1dashboardinfo) - Get Dashboard Info
* [GetAPIV1DashboardExport](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#getapiv1dashboardexport) - Export Dashboards
* [GetAPIV1DashboardDashboardID](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#getapiv1dashboarddashboardid) - Get a Dashboard
* [GetAPIV1DashboardDashboardIDCharts](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#getapiv1dashboarddashboardidcharts) - Get Charts from a Dashboard
* [GetAPIV1DashboardDashboardIDDatasets](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#getapiv1dashboarddashboardiddatasets) - Get Datasets from a Dashboard
* [PostAPIV1Dashboard](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#postapiv1dashboard) - Create a Dashboard
* [PostAPIV1DashboardImport](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#postapiv1dashboardimport) - Import a Dashboard
* [PostAPIV1DashboardDashboardIDPermalink](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#postapiv1dashboarddashboardidpermalink) - Create a Permalink to a Dashboard
* [PutAPIV1DashboardDashboardID](docs/sdks/supersetapisopensourcegreaterthandashboards/README.md#putapiv1dashboarddashboardid) - Update a Dashboard

### [SupersetAPIsOpenSourceGreaterThanDatabases](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md)

* [CreateDatabaseUsingSSH](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#createdatabaseusingssh) - Create a Database Connection using SSH
* [GetAPIV1Database](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#getapiv1database) - Get all Database Connections from a Workspace
* [GetAPIV1DatabaseExport](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#getapiv1databaseexport) - Export Database Connections
* [GetAPIV1DatabaseDatabaseID](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#getapiv1databasedatabaseid) - Get a Database Connection
* [GetAPIV1DatabaseDatabaseIDConnection](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#getapiv1databasedatabaseidconnection) - Get a Database Connection Parameters
* [PostAPIV1Database](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#postapiv1database) - Create a Database Connection
* [PostAPIV1DatabaseImport](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#postapiv1databaseimport) - Import a Database Connection
* [PutAPIV1DatabaseDatabaseID](docs/sdks/supersetapisopensourcegreaterthandatabases/README.md#putapiv1databasedatabaseid) - Update a Database Connection

### [SupersetAPIsOpenSourceGreaterThanDatasets](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md)

* [DeleteAPIV1DatasetDatasetID](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#deleteapiv1datasetdatasetid) - Delete a Dataset
* [GetAPIV1Dataset](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#getapiv1dataset) - Get all Datasets from a Workspace
* [GetAPIV1DatasetExport](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#getapiv1datasetexport) - Export Datasets
* [GetAPIV1DatasetRelatedOwners](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#getapiv1datasetrelatedowners) - Get all possible Dataset Owners
* [GetAPIV1DatasetDatasetID](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#getapiv1datasetdatasetid) - Get a Dataset
* [PostAPIV1Dataset](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#postapiv1dataset) - Create a Virtual Dataset
* [PostAPIV1DatasetImport](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#postapiv1datasetimport) - Import a Dataset
* [PutAPIV1DatasetDatasetID](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#putapiv1datasetdatasetid) - Update a Virtual Dataset
* [PutAPIV1DatasetDatasetIDRefresh](docs/sdks/supersetapisopensourcegreaterthandatasets/README.md#putapiv1datasetdatasetidrefresh) - Refresh a Dataset

### [SupersetAPIsOpenSourceGreaterThanQueries](docs/sdks/supersetapisopensourcegreaterthanqueries/README.md)

* [GetAPIV1Query](docs/sdks/supersetapisopensourcegreaterthanqueries/README.md#getapiv1query) - Get All Workspace Queries

### [SupersetAPIsOpenSourceGreaterThanSQLLab](docs/sdks/supersetapisopensourcegreaterthansqllab/README.md)

* [PostAPIV1SqllabExecute](docs/sdks/supersetapisopensourcegreaterthansqllab/README.md#postapiv1sqllabexecute) - Execute a SQL Query
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
