# PresetManagerAPIsGreaterThanEmbedded
(*.PresetManagerAPIsGreaterThanEmbedded*)

## Overview

APIs associated with the Embedded functionality.

### Available Operations

* [PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestToken](#postapiv1teamsteamslugworkspacesworkspaceslugguesttoken) - Create a new Guest Token

## PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestToken

Creates a new Guest Token to be used with Embedded.

Replace in the URL:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{TeamSlug}}` with the `name` retrieved through the API (using the **Get Preset Teams** endpoint). Alternatively, access the team administration through the UI, and get the `{{TeamSlug}}` from the URL -> `https://manage.app.preset.io/app/teams/{{TeamSlug}}/members`.
    

For instructions on how to populate the body, refer to [our documentation](https://preset-io.github.io/embedded-beta-docs/v1).

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
    res, err := s.PresetManagerAPIsGreaterThanEmbedded.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestToken(ctx, operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenRequest{
        RequestBody: &operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenRequestBody{},
        TeamSlug: "string",
        WorkspaceSlug: "string",
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

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenRequest](../../models/operations/postapiv1teamsteamslugworkspacesworkspaceslugguesttokenrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.PostAPIV1TeamsTeamSlugWorkspacesWorkspaceSlugGuestTokenResponse](../../models/operations/postapiv1teamsteamslugworkspacesworkspaceslugguesttokenresponse.md), error**

