<!-- Start SDK Example Usage -->


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
<!-- End SDK Example Usage -->