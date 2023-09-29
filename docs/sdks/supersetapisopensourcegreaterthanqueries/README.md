# SupersetAPIsOpenSourceGreaterThanQueries
(*SupersetAPIsOpenSourceGreaterThanQueries*)

### Available Operations

* [GetAPIV1Query](#getapiv1query) - Get All Workspace Queries

## GetAPIV1Query

Get All Workspace Queries

### Example Usage

```go
package main

import(
	"context"
	"log"
	presetsdkgo "github.com/speakeasy-sdks/preset-sdk-go"
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/shared"
)

func main() {
    s := presetsdkgo.New(
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanQueries.GetAPIV1Query(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAPIV1QueryResponse](../../models/operations/getapiv1queryresponse.md), error**

