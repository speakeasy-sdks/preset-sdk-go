# PresetManagerAPIsGreaterThanWorkspaces
(*PresetManagerAPIsGreaterThanWorkspaces*)

## Overview

APIs to manage your Workspaces.

### Available Operations

* [GetV1TeamsTeamSlugWorkspaces](#getv1teamsteamslugworkspaces) - Get Workspaces from a Team
* [GetV1TeamsTeamSlugWorkspacesWorkspaceIDMemberships](#getv1teamsteamslugworkspacesworkspaceidmemberships) - Get Workspace Users and Roles
* [PostV1TeamsTeamSlugWorkspaces](#postv1teamsteamslugworkspaces) - Create Workspace for a Team
* [PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembership](#putv1teamsteamslugworkspacesworkspaceidmembership) - Change Workspace Role

## GetV1TeamsTeamSlugWorkspaces

Gets all Workspaces from the Team.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.

### Example Usage

```go
package main

import(
	"context"
	"log"
	presetsdkgo "github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := presetsdkgo.New(
        presetsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanWorkspaces.GetV1TeamsTeamSlugWorkspaces(ctx, operations.GetV1TeamsTeamSlugWorkspacesRequest{
        TeamSlug: "string",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetV1TeamsTeamSlugWorkspacesRequest](../../models/operations/getv1teamsteamslugworkspacesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetV1TeamsTeamSlugWorkspacesResponse](../../models/operations/getv1teamsteamslugworkspacesresponse.md), error**


## GetV1TeamsTeamSlugWorkspacesWorkspaceIDMemberships

Gets Users and their Workspace Roles from the Workspace.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

*   `{{WorkspaceID}}` with the `id` retrieved through the API with the **Get Workspaces from a Team** endpoint.

### Example Usage

```go
package main

import(
	"context"
	"log"
	presetsdkgo "github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := presetsdkgo.New(
        presetsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanWorkspaces.GetV1TeamsTeamSlugWorkspacesWorkspaceIDMemberships(ctx, operations.GetV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipsRequest{
        TeamSlug: "string",
        WorkspaceID: "string",
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.GetV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipsRequest](../../models/operations/getv1teamsteamslugworkspacesworkspaceidmembershipsrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipsResponse](../../models/operations/getv1teamsteamslugworkspacesworkspaceidmembershipsresponse.md), error**


## PostV1TeamsTeamSlugWorkspaces

###### *Requires admin permission.*

Creates a new Workspace on the Team.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Replace in the Body:

*   `{{NewWorkspaceTitle}}` by the title of the New Workspace you want to create.
*   `{{NewWorkspaceRegion}}` with the region that the Workspace should be created - refer to below table:
    

| Location | `NewWorkspaceRegion` |
| --- | --- |
| **US East Cost** | `us-east-1` |
| **US West Coast** | `us-west-2` |
| **Europe** | `eu-north-1` |
| **Asia-Pacific** | `ap-northeast-1` |

### Example Usage

```go
package main

import(
	"context"
	"log"
	presetsdkgo "github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := presetsdkgo.New(
        presetsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanWorkspaces.PostV1TeamsTeamSlugWorkspaces(ctx, operations.PostV1TeamsTeamSlugWorkspacesRequest{
        RequestBody: &operations.PostV1TeamsTeamSlugWorkspacesRequestBody{},
        TeamSlug: "string",
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
| `request`                                                                                                          | [operations.PostV1TeamsTeamSlugWorkspacesRequest](../../models/operations/postv1teamsteamslugworkspacesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PostV1TeamsTeamSlugWorkspacesResponse](../../models/operations/postv1teamsteamslugworkspacesresponse.md), error**


## PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembership

Changes a user's Workspace Role.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

- `{{WorkspaceID}}` with the `id` retrieved through the API with the **Get Workspaces from a Team** endpoint.
    

Replace in the body:

- `{{UserID}}`with the `id` retrieved using the **Get Team Members** endpoint.
    

- `{{RoleIdentifier}}`with one of the the `role_identifier` below, wrapped in double quotes:
    

| **Workspace Role** | **role_identifier** |
| --- | --- |
| Workspace Admin | "Admin" |
| Primary Contributor | "PresetAlpha" |
| Secondary Contributor | "PresetBeta" |
| Limited Contributor | "PresetGamma" |
| Viewer | "PresetReportsOnly" |
| Dashboard Viewer | "PresetDashboardsOnly" |
| No Access | "PresetNoAccess" |

Each Role Identifier corresponds to a specific Workspace Role. As a reminder, the access restrictions for each Workspace Role can be found [in our documentation](https://docs.preset.io/docs/data-access-roles-at-preset).

### Example Usage

```go
package main

import(
	"context"
	"log"
	presetsdkgo "github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := presetsdkgo.New(
        presetsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PresetManagerAPIsGreaterThanWorkspaces.PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembership(ctx, operations.PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipRequest{
        RequestBody: &operations.PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipRequestBody{},
        TeamSlug: "string",
        WorkspaceID: "string",
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

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `request`                                                                                                                                                  | [operations.PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipRequest](../../models/operations/putv1teamsteamslugworkspacesworkspaceidmembershiprequest.md) | :heavy_check_mark:                                                                                                                                         | The request object to use for the request.                                                                                                                 |


### Response

**[*operations.PutV1TeamsTeamSlugWorkspacesWorkspaceIDMembershipResponse](../../models/operations/putv1teamsteamslugworkspacesworkspaceidmembershipresponse.md), error**

