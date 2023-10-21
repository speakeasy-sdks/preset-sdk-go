# SupersetAPIsOpenSourceGreaterThanAnnotationLayers
(*SupersetAPIsOpenSourceGreaterThanAnnotationLayers*)

## Overview

API to manage your Annotation Layers.

### Available Operations

* [DeleteAPIV1AnnotationLayer](#deleteapiv1annotationlayer) - Delete multiple Annotations Layers
* [DeleteAPIV1AnnotationLayerAnnotationLayerID](#deleteapiv1annotationlayerannotationlayerid) - Delete an Annotation Layer
* [DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotation](#deleteapiv1annotationlayerannotationlayeridannotation) - Delete multiple Annotations from an Annotation Layer
* [DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID](#deleteapiv1annotationlayerannotationlayeridannotationannotationid) - Delete an Annotation from an Annotation Layer
* [GetAPIV1AnnotationLayer](#getapiv1annotationlayer) - Get all Annotation Layers from a Workspace
* [GetAPIV1AnnotationLayerAnnotationLayerID](#getapiv1annotationlayerannotationlayerid) - Get an Annotation Layer
* [GetAPIV1AnnotationLayerAnnotationLayerIDAnnotation](#getapiv1annotationlayerannotationlayeridannotation) - Get all Annotations from an Annotation Layer
* [GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID](#getapiv1annotationlayerannotationlayeridannotationannotationid) - Get an Annotation from an Annotation Layer
* [PostAPIV1AnnotationLayer](#postapiv1annotationlayer) - Create an Annotation Layer
* [PostAPIV1AnnotationLayerAnnotationLayerIDAnnotation](#postapiv1annotationlayerannotationlayeridannotation) - Create an Annotation in an Annotation Layer
* [PutAPIV1AnnotationLayerAnnotationLayerID](#putapiv1annotationlayerannotationlayerid) - Update an Annotation Layer
* [PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID](#putapiv1annotationlayerannotationlayeridannotationannotationid) - Update an Annotation from an Annotation Layer

## DeleteAPIV1AnnotationLayer

Deletes multiple Annotations from an Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
- `{{AnnotationLayerIDs}` with comma separated Annotation Layer `ids` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.DeleteAPIV1AnnotationLayer(ctx, operations.DeleteAPIV1AnnotationLayerRequest{})
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
| `request`                                                                                                    | [operations.DeleteAPIV1AnnotationLayerRequest](../../models/operations/deleteapiv1annotationlayerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.DeleteAPIV1AnnotationLayerResponse](../../models/operations/deleteapiv1annotationlayerresponse.md), error**


## DeleteAPIV1AnnotationLayerAnnotationLayerID

Deletes an Annotation Layer through the API.

**Note:** You can only delete an Annotation Layer, after deleting **all its Annotations**.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.DeleteAPIV1AnnotationLayerAnnotationLayerID(ctx, operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDRequest{
        AnnotationLayerID: "string",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDRequest](../../models/operations/deleteapiv1annotationlayerannotationlayeridrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDResponse](../../models/operations/deleteapiv1annotationlayerannotationlayeridresponse.md), error**


## DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotation

Deletes multiple Annotations from an Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
- `{{AnnotationIDs}` with comma separated Annotation `ids` retrieved from the **Get all Annotations from an Annotation Layer** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotation(ctx, operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest{
        AnnotationLayerID: "string",
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

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `request`                                                                                                                                                          | [operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest](../../models/operations/deleteapiv1annotationlayerannotationlayeridannotationrequest.md) | :heavy_check_mark:                                                                                                                                                 | The request object to use for the request.                                                                                                                         |


### Response

**[*operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse](../../models/operations/deleteapiv1annotationlayerannotationlayeridannotationresponse.md), error**


## DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID

Deletes an Annotation from an Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
- `{{AnnotationID}}` with the Annotation `id` retrieved from the **Get all Annotations from an Annotation Layer** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID(ctx, operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest{
        AnnotationID: "string",
        AnnotationLayerID: "string",
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

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `request`                                                                                                                                                                                  | [operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest](../../models/operations/deleteapiv1annotationlayerannotationlayeridannotationannotationidrequest.md) | :heavy_check_mark:                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                 |


### Response

**[*operations.DeleteAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse](../../models/operations/deleteapiv1annotationlayerannotationlayeridannotationannotationidresponse.md), error**


## GetAPIV1AnnotationLayer

Gets all Annotation Layers from a Workspace.

Replace in the URL:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

Note that this endpoint returns 20 results by default. You can return up to 100 results at a time and use pagination by adding the following query parameters:

```
?q=(page_size:{{PageSize}},page:{{Page}})

```

Replace:

- `{{PageSize}}` with the desired size (min `1` max `100`).
- `{{Page}}` with the page number (useful when the total count > `{{PageSize}}` - min `0`).

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
        presetsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.GetAPIV1AnnotationLayer(ctx)
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

**[*operations.GetAPIV1AnnotationLayerResponse](../../models/operations/getapiv1annotationlayerresponse.md), error**


## GetAPIV1AnnotationLayerAnnotationLayerID

Gets a specific Annotation Layer from a Workspace.

Replace in the URL:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.GetAPIV1AnnotationLayerAnnotationLayerID(ctx, operations.GetAPIV1AnnotationLayerAnnotationLayerIDRequest{
        AnnotationLayerID: "string",
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
| `request`                                                                                                                                | [operations.GetAPIV1AnnotationLayerAnnotationLayerIDRequest](../../models/operations/getapiv1annotationlayerannotationlayeridrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.GetAPIV1AnnotationLayerAnnotationLayerIDResponse](../../models/operations/getapiv1annotationlayerannotationlayeridresponse.md), error**


## GetAPIV1AnnotationLayerAnnotationLayerIDAnnotation

Gets all Annotations from a specific Annotation Layer.

Replace in the URL:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get All Annotation Layers from a Workspace** endpoint.
    

Note that this endpoint returns 20 results by default. You can return up to 100 results at a time and use pagination by adding the following query parameters:

```
?q=(page_size:{{PageSize}},page:{{Page}})

```

Replace:

- `{{PageSize}}` with the desired size (min `1` max `100`).
- `{{Page}}` with the page number (useful when the total count > `{{PageSize}}` - min `0`).

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotation(ctx, operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest{
        AnnotationLayerID: "string",
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
| `request`                                                                                                                                                    | [operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest](../../models/operations/getapiv1annotationlayerannotationlayeridannotationrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse](../../models/operations/getapiv1annotationlayerannotationlayeridannotationresponse.md), error**


## GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID

Deletes an Annotation Layer through the API.

**Note:** You can only delete an Annotation Layer, after deleting **all its Annotations**.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
- `{{AnnotationID}}` with the Annotation `id` retrieved from the **Get all Annotations from an Annotation Layer** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID(ctx, operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest{
        AnnotationID: "string",
        AnnotationLayerID: "string",
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

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest](../../models/operations/getapiv1annotationlayerannotationlayeridannotationannotationidrequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |


### Response

**[*operations.GetAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse](../../models/operations/getapiv1annotationlayerannotationlayeridannotationannotationidresponse.md), error**


## PostAPIV1AnnotationLayer

Creates an Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

Replace in the body:

- `{{AnnotationLayerDescription}}` with a description for the Annotation Layer.
- `{{AnnotationLayerName}}` with the desired new name.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.PostAPIV1AnnotationLayer(ctx, operations.PostAPIV1AnnotationLayerRequest{
        RequestBody: &operations.PostAPIV1AnnotationLayerRequestBody{},
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostAPIV1AnnotationLayerRequest](../../models/operations/postapiv1annotationlayerrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PostAPIV1AnnotationLayerResponse](../../models/operations/postapiv1annotationlayerresponse.md), error**


## PostAPIV1AnnotationLayerAnnotationLayerIDAnnotation

Creates an Annotation on an existing Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
    

Replace in the body:

- `{{AnnotationEndDTTM}}` with the annotation's datetime end (`YYYY-MM-DD HH:MM`).
- `{{AnnotationLongDescription}}` with the annotation's description.
- `{{AnnotationTitle}}` with a name for it.
- `{{AnnotationStartDTTM}}` with the annotation's datetime start (`YYYY-MM-DD HH:MM`).

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotation(ctx, operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest{
        AnnotationLayerID: "string",
        RequestBody: &operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequestBody{},
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

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationRequest](../../models/operations/postapiv1annotationlayerannotationlayeridannotationrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.PostAPIV1AnnotationLayerAnnotationLayerIDAnnotationResponse](../../models/operations/postapiv1annotationlayerannotationlayeridannotationresponse.md), error**


## PutAPIV1AnnotationLayerAnnotationLayerID

Updates an Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
    

Replace in the body:

- `{{AnnotationLayerDescription}}` with a description for the Annotation Layer.
- `{{AnnotationLayerName}}` with the desired new name.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.PutAPIV1AnnotationLayerAnnotationLayerID(ctx, operations.PutAPIV1AnnotationLayerAnnotationLayerIDRequest{
        AnnotationLayerID: "string",
        RequestBody: &operations.PutAPIV1AnnotationLayerAnnotationLayerIDRequestBody{},
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
| `request`                                                                                                                                | [operations.PutAPIV1AnnotationLayerAnnotationLayerIDRequest](../../models/operations/putapiv1annotationlayerannotationlayeridrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.PutAPIV1AnnotationLayerAnnotationLayerIDResponse](../../models/operations/putapiv1annotationlayerannotationlayeridresponse.md), error**


## PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID

Updates an Annotation from an Annotation Layer through the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
- `{{AnnotationLayerID}}` with the Annotation Layer `id` retrieved from the **Get all Annotation Layers from a Workspace** endpoint.
    

Replace in the body:

- `{{AnnotationEndDTTM}}` with the annotation's datetime end (`YYYY-MM-DD HH:MM`).
- `{{AnnotationLongDescription}}` with the annotation's description.
- `{{AnnotationShortDescription}}` with a name for it.
- `{{AnnotationStartDTTM}}` with the annotation's datetime start (`YYYY-MM-DD HH:MM`).

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAnnotationLayers.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationID(ctx, operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest{
        AnnotationID: "string",
        AnnotationLayerID: "string",
        RequestBody: &operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequestBody{},
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

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDRequest](../../models/operations/putapiv1annotationlayerannotationlayeridannotationannotationidrequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |


### Response

**[*operations.PutAPIV1AnnotationLayerAnnotationLayerIDAnnotationAnnotationIDResponse](../../models/operations/putapiv1annotationlayerannotationlayeridannotationannotationidresponse.md), error**

