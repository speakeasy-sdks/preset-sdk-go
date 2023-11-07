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
	res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.DeleteAPIV1AnnotationLayer(ctx, operations.DeleteAPIV1AnnotationLayerRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->