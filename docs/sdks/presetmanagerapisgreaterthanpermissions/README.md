# PresetManagerAPIsGreaterThanPermissions

## Overview

APIs to manage permissions on the Workspace level.

Note that all Permission APIs require **Team Admin** permission.

### Available Operations

* [DeleteV1TeamsTeamSlugPermissionsPermissionName](#deletev1teamsteamslugpermissionspermissionname) - Delete Row Level Security
* [DeleteV1TeamsTeamSlugPermissionsPermissionNameGrantees](#deletev1teamsteamslugpermissionspermissionnamegrantees) - Delete Grantee from existing Permission
* [GetV1TeamsTeamSlugPermissions](#getv1teamsteamslugpermissions) - Get Permissions
* [GetV1TeamsTeamSlugPermissionsResources](#getv1teamsteamslugpermissionsresources) - Get Resources
* [PatchV1TeamsTeamSlugPermissionsPermissionName](#patchv1teamsteamslugpermissionspermissionname) - Update Row Level Security
* [PostV1TeamsTeamSlugPermissions](#postv1teamsteamslugpermissions) - Create Row Level Security
* [PostV1TeamsTeamSlugPermissionsPermissionNameGrantees](#postv1teamsteamslugpermissionspermissionnamegrantees) - Add Grantees to existing Permission

## DeleteV1TeamsTeamSlugPermissionsPermissionName

###### _Requires Team Admin permission._

Deletes a Row Level Security (RLS).

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
- `{{PermissionName}}` with the `name` retrieved via the **Get Permissions** API.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.DeleteV1TeamsTeamSlugPermissionsPermissionName(ctx, operations.DeleteV1TeamsTeamSlugPermissionsPermissionNameRequest{
        PermissionName: "distinctio",
        TeamSlug: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.DeleteV1TeamsTeamSlugPermissionsPermissionNameRequest](../../models/operations/deletev1teamsteamslugpermissionspermissionnamerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.DeleteV1TeamsTeamSlugPermissionsPermissionNameResponse](../../models/operations/deletev1teamsteamslugpermissionspermissionnameresponse.md), error**


## DeleteV1TeamsTeamSlugPermissionsPermissionNameGrantees

###### _Requires Team Admin permission._

Removes Grantee(s) from an existing Permission (DAR or RLS).

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
- `{{PermissionName}}` with the `name` retrieved via the **Get Permissions** API.
    

Replace in the body:

- `{{GranteeType}}` with:
    - `USER` to add users to DARs.
    - `ROLE` to add DAR to RLS.
- `{{GranteeIdentifier}}` with:
    - `username` when adding users to DAR. You can get a list of `usernames` with the **Get Team Members** API.
    - `DAR Name` when adding DAR to RLS. You can get a list of availabe DARs using the **Get Permissions** endpoint. Note that you would use the **display name** (without `dar:`).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.DeleteV1TeamsTeamSlugPermissionsPermissionNameGrantees(ctx, operations.DeleteV1TeamsTeamSlugPermissionsPermissionNameGranteesRequest{
        PermissionName: "unde",
        TeamSlug: "nulla",
        Referer: preset.String("corrupti"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                            | Type                                                                                                                                                                 | Required                                                                                                                                                             | Description                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                | :heavy_check_mark:                                                                                                                                                   | The context to use for the request.                                                                                                                                  |
| `request`                                                                                                                                                            | [operations.DeleteV1TeamsTeamSlugPermissionsPermissionNameGranteesRequest](../../models/operations/deletev1teamsteamslugpermissionspermissionnamegranteesrequest.md) | :heavy_check_mark:                                                                                                                                                   | The request object to use for the request.                                                                                                                           |


### Response

**[*operations.DeleteV1TeamsTeamSlugPermissionsPermissionNameGranteesResponse](../../models/operations/deletev1teamsteamslugpermissionspermissionnamegranteesresponse.md), error**


## GetV1TeamsTeamSlugPermissions

###### _Requires Team Admin permission._

List all permissions from a team.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Note that the query parameters are optional and very useful to retrieve the desired data:

- `workspace_name` can be used to filter for permissions only applied to a particular Workspace.
- `permission_type` can be used to filter for only `data_access_role`, or `row_level_security`.
- `grantee_identifier` can be used to filter for all permissions included in a DAR (`dar:{{DAR NAME}}`), or all permissions applied to a user (`username`).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.GetV1TeamsTeamSlugPermissions(ctx, operations.GetV1TeamsTeamSlugPermissionsRequest{
        TeamSlug: "illum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetV1TeamsTeamSlugPermissionsRequest](../../models/operations/getv1teamsteamslugpermissionsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetV1TeamsTeamSlugPermissionsResponse](../../models/operations/getv1teamsteamslugpermissionsresponse.md), error**


## GetV1TeamsTeamSlugPermissionsResources

###### _Requires Team Admin permission._

List all resources from a team.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Note that the query parameters are optional and very useful to retrieve the desired data:

- `workspace_name` can be used to filter for permissions only applied to a particular Workspace.
- `resource_type` can be used to filter for:
    - `database` to list databases
    - `database_schema` to list schemas
    - `datasource` to list DAR permission reference of a dataset (not the dataset itself)
    - `dataset` to list datasets
    - `data_access_role` to list DARs
    - `row_level_security` to list RLSs

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.GetV1TeamsTeamSlugPermissionsResources(ctx, operations.GetV1TeamsTeamSlugPermissionsResourcesRequest{
        TeamSlug: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetV1TeamsTeamSlugPermissionsResourcesRequest](../../models/operations/getv1teamsteamslugpermissionsresourcesrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetV1TeamsTeamSlugPermissionsResourcesResponse](../../models/operations/getv1teamsteamslugpermissionsresourcesresponse.md), error**


## PatchV1TeamsTeamSlugPermissionsPermissionName

###### _Requires Team Admin permission._

Updates an existing Row Level Security (RLS).

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
- `{{PermissionName}}` with the `name` retrieved via the **Get Permissions** API.
    

Note that the payload has to be complete - you can't remove the `grantees`.

Replace in the body:

- `{{DARName1}}` with the DAR name you want to associated with this RLS (also for `{{DARName2}}` and so on). You can get a list of availabe DARs using the **Get Permissions** endpoint. Note that you would use the **display name** (without `dar:`).
- `{{RLSName}}` with the name of your RLS.
- `{{RLSClause}}` with the SQL syntax for the filter to be applied.
- `{{RLSFilterType}}` with `Regular` or `Base`.
- `{{RLSGroupKey}}` with the desired group key.
- `{{ResourceName}}` with the datasource name. You can get a list of available options using the **Get Resources** endpoint (filtering for `&resource_type=datasource`).
- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
    

_Please note:_

- _**Grantee list for RLS permissions create/update is limited to 100 users in total (API will respond with 400 if over the limit)**_

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.PatchV1TeamsTeamSlugPermissionsPermissionName(ctx, operations.PatchV1TeamsTeamSlugPermissionsPermissionNameRequest{
        PermissionName: "error",
        RequestBody: &operations.PatchV1TeamsTeamSlugPermissionsPermissionNameRequestBody{},
        TeamSlug: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.PatchV1TeamsTeamSlugPermissionsPermissionNameRequest](../../models/operations/patchv1teamsteamslugpermissionspermissionnamerequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.PatchV1TeamsTeamSlugPermissionsPermissionNameResponse](../../models/operations/patchv1teamsteamslugpermissionspermissionnameresponse.md), error**


## PostV1TeamsTeamSlugPermissions

###### _Requires Team Admin permission._

Creates a Row Level Security (RLS).

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Replace in the body:

- `{{DARName1}}` with the DAR name you want to associated with this RLS (also for `{{DARName2}}` and so on). You can get a list of availabe DARs using the **Get Permissions** endpoint. Note that you would use the **display name** (without `dar:`).
- `{{RLSName}}` with the name of your RLS.
- `{{RLSClause}}` with the SQL syntax for the filter to be applied.
- `{{RLSFilterType}}` with `Regular` or `Base`.
- `{{RLSGroupKey}}` with the desired group key.
- `{{ResourceName}}` with the datasource name. You can get a list of available options using the **Get Resources** endpoint (filtering for `&resource_type=datasource`).
- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
    

Please note:

- _**Grantee list for RLS permissions create/update is limited to 100 users in total (API will respond with 400 if over the limit)**_

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.PostV1TeamsTeamSlugPermissions(ctx, operations.PostV1TeamsTeamSlugPermissionsRequest{
        RequestBody: &operations.PostV1TeamsTeamSlugPermissionsRequestBody{},
        TeamSlug: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostV1TeamsTeamSlugPermissionsRequest](../../models/operations/postv1teamsteamslugpermissionsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostV1TeamsTeamSlugPermissionsResponse](../../models/operations/postv1teamsteamslugpermissionsresponse.md), error**


## PostV1TeamsTeamSlugPermissionsPermissionNameGrantees

###### _Requires Team Admin permission._

Adds Grantee(s) to an existing Permission (DAR or RLS).

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
- `{{PermissionName}}` with the `name` retrieved via the **Get Permissions** API.
    

Replace in the body:

- `{{GranteeType}}` with:
    - `USER` to add existing user to DARs.
    - `INVITEE` to add pending user to DARs
    - `ROLE` to add DAR to RLS.
- `{{GranteeIdentifier}}` with:
    - `username` when adding users to DAR. You can get a list of `usernames` with the **Get Team Members** API. For pending users you can use their email.
    - `DAR Name` when adding DAR to RLS. You can get a list of availabe DARs using the **Get Permissions** endpoint. Note that you would use the **display name** (without `dar:`).

Please note:

- _**Max number of grantees to be appended per request is 100.**_
- _**Duplicate grantees will be ignored.**_

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := preset.New(
        preset.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanPermissions.PostV1TeamsTeamSlugPermissionsPermissionNameGrantees(ctx, operations.PostV1TeamsTeamSlugPermissionsPermissionNameGranteesRequest{
        PermissionName: "iure",
        RequestBody: &operations.PostV1TeamsTeamSlugPermissionsPermissionNameGranteesRequestBody{},
        TeamSlug: "magnam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.PostV1TeamsTeamSlugPermissionsPermissionNameGranteesRequest](../../models/operations/postv1teamsteamslugpermissionspermissionnamegranteesrequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.PostV1TeamsTeamSlugPermissionsPermissionNameGranteesResponse](../../models/operations/postv1teamsteamslugpermissionspermissionnamegranteesresponse.md), error**
