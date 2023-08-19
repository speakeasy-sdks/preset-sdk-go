# PresetManagerAPIsGreaterThanTeams

## Overview

APIs to manage your Preset team.

### Available Operations

* [DeleteV1TeamsTeamSlugInvitesInviteID](#deletev1teamsteamsluginvitesinviteid) - Delete Pending Invite
* [DeleteV1TeamsTeamSlugMembershipsUserID](#deletev1teamsteamslugmembershipsuserid) - Delete a Team Member
* [GetV1Teams](#getv1teams) - Get Preset Teams
* [GetV1TeamsTeamSlugInvites](#getv1teamsteamsluginvites) - Get Pending Team Invites
* [GetV1TeamsTeamSlugMemberships](#getv1teamsteamslugmemberships) - Get Team Members
* [PatchV1TeamsTeamSlugMembershipsUserID](#patchv1teamsteamslugmembershipsuserid) - Change User Role
* [PostV1TeamsTeamSlugInvites](#postv1teamsteamsluginvites) - Create a Team Invite
* [PostV1TeamsTeamSlugInvitesMany](#postv1teamsteamsluginvitesmany) - Create Multiple Team Invites
* [PostV1TeamsTeamSlugInvitesResendInviteID](#postv1teamsteamsluginvitesresendinviteid) - Resend Invite
* [PutV1TeamsTeamSlug](#putv1teamsteamslug) - Update Team Title

## DeleteV1TeamsTeamSlugInvitesInviteID

###### *Requires admin permission.*

Deletes a Team invitation.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

*   `{{InviteID}}`with the `id` retrieved through the API using the **Get Pending Team Invites** endpoint.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.DeleteV1TeamsTeamSlugInvitesInviteID(ctx, operations.DeleteV1TeamsTeamSlugInvitesInviteIDRequest{
        InviteID: "debitis",
        TeamSlug: "ipsa",
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.DeleteV1TeamsTeamSlugInvitesInviteIDRequest](../../models/operations/deletev1teamsteamsluginvitesinviteidrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.DeleteV1TeamsTeamSlugInvitesInviteIDResponse](../../models/operations/deletev1teamsteamsluginvitesinviteidresponse.md), error**


## DeleteV1TeamsTeamSlugMembershipsUserID

###### _Requires admin permission._

Deletes a Team member.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
- `{{UserID}}`with the `id` retrieved using the **Get Team Members** endpoint.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.DeleteV1TeamsTeamSlugMembershipsUserID(ctx, operations.DeleteV1TeamsTeamSlugMembershipsUserIDRequest{
        TeamSlug: "delectus",
        UserID: "tempora",
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
| `request`                                                                                                                            | [operations.DeleteV1TeamsTeamSlugMembershipsUserIDRequest](../../models/operations/deletev1teamsteamslugmembershipsuseridrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.DeleteV1TeamsTeamSlugMembershipsUserIDResponse](../../models/operations/deletev1teamsteamslugmembershipsuseridresponse.md), error**


## GetV1Teams

Retrieves all Preset teams the user has access to.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.GetV1Teams(ctx, operations.GetV1TeamsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetV1TeamsRequest](../../models/operations/getv1teamsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetV1TeamsResponse](../../models/operations/getv1teamsresponse.md), error**


## GetV1TeamsTeamSlugInvites

###### *Requires admin permission.*

Gets all pending team invitations.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.GetV1TeamsTeamSlugInvites(ctx, operations.GetV1TeamsTeamSlugInvitesRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetV1TeamsTeamSlugInvitesRequest](../../models/operations/getv1teamsteamsluginvitesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetV1TeamsTeamSlugInvitesResponse](../../models/operations/getv1teamsteamsluginvitesresponse.md), error**


## GetV1TeamsTeamSlugMemberships

###### _Requires admin permission._

Gets all members of the Team.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.GetV1TeamsTeamSlugMemberships(ctx, operations.GetV1TeamsTeamSlugMembershipsRequest{
        TeamSlug: "molestiae",
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
| `request`                                                                                                          | [operations.GetV1TeamsTeamSlugMembershipsRequest](../../models/operations/getv1teamsteamslugmembershipsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetV1TeamsTeamSlugMembershipsResponse](../../models/operations/getv1teamsteamslugmembershipsresponse.md), error**


## PatchV1TeamsTeamSlugMembershipsUserID

###### _Requires admin permission._

Changes a user's Team role.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
- `{{UserID}}`with the `id` retrieved using the **Get Team Members** endpoint.
    

Replace in the Body:

- `{{RoleID}}` with the desired role:
    - Use **`2`** for **User**.
    - Use **`1`** for **Admin**.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.PatchV1TeamsTeamSlugMembershipsUserID(ctx, operations.PatchV1TeamsTeamSlugMembershipsUserIDRequest{
        RequestBody: &operations.PatchV1TeamsTeamSlugMembershipsUserIDRequestBody{},
        TeamSlug: "minus",
        UserID: "placeat",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PatchV1TeamsTeamSlugMembershipsUserIDRequest](../../models/operations/patchv1teamsteamslugmembershipsuseridrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.PatchV1TeamsTeamSlugMembershipsUserIDResponse](../../models/operations/patchv1teamsteamslugmembershipsuseridresponse.md), error**


## PostV1TeamsTeamSlugInvites

###### *Requires admin permission.*

Creates a team invitation, sent via email.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Replace in the Body:

*   `{{RoleID}}` with the desired role:
    *   Use **`2`** for **User**.
    *   Use **`1`** for **Admin**.
*   `{{Email}}` with the user's email address.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.PostV1TeamsTeamSlugInvites(ctx, operations.PostV1TeamsTeamSlugInvitesRequest{
        RequestBody: &operations.PostV1TeamsTeamSlugInvitesRequestBody{},
        TeamSlug: "voluptatum",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostV1TeamsTeamSlugInvitesRequest](../../models/operations/postv1teamsteamsluginvitesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostV1TeamsTeamSlugInvitesResponse](../../models/operations/postv1teamsteamsluginvitesresponse.md), error**


## PostV1TeamsTeamSlugInvitesMany

###### *Requires admin permission.*

Creates multiple team invitations.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Replace in the body:

*   `{{RoleID}}` with the desired role:
    *   Use **`2`** for **User**.
    *   Use **`1`** for **Admin**.
*   `{{Email}}, {{Email2}}...` with the users' email addresses.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.PostV1TeamsTeamSlugInvitesMany(ctx, operations.PostV1TeamsTeamSlugInvitesManyRequest{
        RequestBody: &operations.PostV1TeamsTeamSlugInvitesManyRequestBody{},
        TeamSlug: "iusto",
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
| `request`                                                                                                            | [operations.PostV1TeamsTeamSlugInvitesManyRequest](../../models/operations/postv1teamsteamsluginvitesmanyrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostV1TeamsTeamSlugInvitesManyResponse](../../models/operations/postv1teamsteamsluginvitesmanyresponse.md), error**


## PostV1TeamsTeamSlugInvitesResendInviteID

Resends a pending Team invitation.

Replace in the URL:

*   `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`
    

*   `{{InviteID}}`with the `id` retrieved through the API using the **Get Pending Team Invites** endpoint.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.PostV1TeamsTeamSlugInvitesResendInviteID(ctx, operations.PostV1TeamsTeamSlugInvitesResendInviteIDRequest{
        InviteID: "excepturi",
        TeamSlug: "nisi",
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

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.PostV1TeamsTeamSlugInvitesResendInviteIDRequest](../../models/operations/postv1teamsteamsluginvitesresendinviteidrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.PostV1TeamsTeamSlugInvitesResendInviteIDResponse](../../models/operations/postv1teamsteamsluginvitesresendinviteidresponse.md), error**


## PutV1TeamsTeamSlug

###### _Requires admin permission._

Updates the Team title.

Replace in the URL:

- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

Replace in the Body:

- `{{NewTeamTitle}}` by the new title you want to set.

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
    res, err := s.PresetManagerAPIsGreaterThanTeams.PutV1TeamsTeamSlug(ctx, operations.PutV1TeamsTeamSlugRequest{
        RequestBody: &operations.PutV1TeamsTeamSlugRequestBody{},
        TeamSlug: "recusandae",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PutV1TeamsTeamSlugRequest](../../models/operations/putv1teamsteamslugrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PutV1TeamsTeamSlugResponse](../../models/operations/putv1teamsteamslugresponse.md), error**

