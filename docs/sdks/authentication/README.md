# Authentication

## Overview

API to authenticate and get a JWT token to interact with the Preset/Superset APIs.

### Available Operations

* [PostV1Auth](#postv1auth) - Get a JWT Token

## PostV1Auth

To interact with the Preset API, it's required to generate an API Key, that's used to generate a JWT token.

1. Generate an API Key on the Preset Manager UI. Refer to [our documentation](https://docs.preset.io/docs/the-preset-api).
2. Copy the API `token` and `secret`.
    

Replace in the body:

- `{{APIToken}}` with the `token` from the UI.
- `{{APISecret}}` with the `secret` from the UI.

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

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.PostV1AuthRequest](../../models/operations/postv1authrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.PostV1AuthResponse](../../models/operations/postv1authresponse.md), error**

