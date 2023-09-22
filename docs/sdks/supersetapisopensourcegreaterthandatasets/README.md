# SupersetAPIsOpenSourceGreaterThanDatasets

## Overview

APIs to manage your datasets.

### Available Operations

* [DeleteAPIV1DatasetDatasetID](#deleteapiv1datasetdatasetid) - Delete a Dataset
* [GetAPIV1Dataset](#getapiv1dataset) - Get all Datasets from a Workspace
* [GetAPIV1DatasetExport](#getapiv1datasetexport) - Export Datasets
* [GetAPIV1DatasetRelatedOwners](#getapiv1datasetrelatedowners) - Get all possible Dataset Owners
* [GetAPIV1DatasetDatasetID](#getapiv1datasetdatasetid) - Get a Dataset
* [PostAPIV1Dataset](#postapiv1dataset) - Create a Virtual Dataset
* [PostAPIV1DatasetImport](#postapiv1datasetimport) - Import a Dataset
* [PutAPIV1DatasetDatasetID](#putapiv1datasetdatasetid) - Update a Virtual Dataset
* [PutAPIV1DatasetDatasetIDRefresh](#putapiv1datasetdatasetidrefresh) - Refresh a Dataset

## DeleteAPIV1DatasetDatasetID

Deletes a Dataset.

Replace in the URL:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.DeleteAPIV1DatasetDatasetID(ctx, operations.DeleteAPIV1DatasetDatasetIDRequest{
        DatasetID: "omnis",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteAPIV1DatasetDatasetIDRequest](../../models/operations/deleteapiv1datasetdatasetidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteAPIV1DatasetDatasetIDResponse](../../models/operations/deleteapiv1datasetdatasetidresponse.md), error**


## GetAPIV1Dataset

Gets all Datasets available on the Worpkspace.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.GetAPIV1Dataset(ctx, operations.GetAPIV1DatasetRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetAPIV1DatasetRequest](../../models/operations/getapiv1datasetrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetAPIV1DatasetResponse](../../models/operations/getapiv1datasetresponse.md), error**


## GetAPIV1DatasetExport

Exports Datasets from the Workspace.

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

- `{{DatasetIDs}}` with comma separated `ids` retrieved from the **Get all Datasets from a Workspace** endpoint.
    

_**Tip:**_ If used in Postman, select `Save Response` and `Save to a File` to get the zip export.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.GetAPIV1DatasetExport(ctx, operations.GetAPIV1DatasetExportRequest{
        Q: presetsdkgo.String("nemo"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetAPIV1DatasetExportRequest](../../models/operations/getapiv1datasetexportrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetAPIV1DatasetExportResponse](../../models/operations/getapiv1datasetexportresponse.md), error**


## GetAPIV1DatasetRelatedOwners

Gets all possible Dataset owners on the Workspace.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.GetAPIV1DatasetRelatedOwners(ctx, operations.GetAPIV1DatasetRelatedOwnersRequest{})
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
| `request`                                                                                                        | [operations.GetAPIV1DatasetRelatedOwnersRequest](../../models/operations/getapiv1datasetrelatedownersrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetAPIV1DatasetRelatedOwnersResponse](../../models/operations/getapiv1datasetrelatedownersresponse.md), error**


## GetAPIV1DatasetDatasetID

Get a specific Dataset from a Workspace.

Replace in the URL:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

- `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.GetAPIV1DatasetDatasetID(ctx, operations.GetAPIV1DatasetDatasetIDRequest{
        DatasetID: "minima",
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
| `request`                                                                                                | [operations.GetAPIV1DatasetDatasetIDRequest](../../models/operations/getapiv1datasetdatasetidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetAPIV1DatasetDatasetIDResponse](../../models/operations/getapiv1datasetdatasetidresponse.md), error**


## PostAPIV1Dataset

Creates a new Virtual Dataset.

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

Replace in the body:

- `{{DBID}}` with the database `id` retrieved using the **Get Databases** endoint.
- `{{SchemaName}}` with the desired schema.
- `{{TableName}}` with the desired table.
- `{{DatasetSQL}}` with the SQL query to power your dataset.
- `{{OwnerID}}` with the `id` for the desired owner account(s) retrieved from the **Get all possible Dataset Owners** endpoint. Use a comma to separate multiple IDs (for example, `[2,5]`).

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.PostAPIV1Dataset(ctx, operations.PostAPIV1DatasetRequest{
        RequestBody: &operations.PostAPIV1DatasetRequestBody{},
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PostAPIV1DatasetRequest](../../models/operations/postapiv1datasetrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostAPIV1DatasetResponse](../../models/operations/postapiv1datasetresponse.md), error**


## PostAPIV1DatasetImport

Imports a Dataset.

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

In the body:

- Select your Dataset ZIP file as a value for the `formData`.
- For the `passwords` field:
    - Replace `{{DatabaseYAMLFile}}` by the database YAML file name. You can find it in your export file, under the `databases` folder.
    - Replace `{{DatabasePassword}}` by the DB password.
- For the `overwrite` field:
    - If the DB Connection already exists on the destination Workspace, set it as `true` to overwrite it.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.PostAPIV1DatasetImport(ctx, operations.PostAPIV1DatasetImportRequest{
        Referer: presetsdkgo.String("excepturi"),
        RequestBody: &operations.PostAPIV1DatasetImportRequestBody{
            FormData: &operations.PostAPIV1DatasetImportRequestBodyFormData{
                Content: []byte("accusantium"),
                FormData: "iure",
            },
            Overwrite: presetsdkgo.Bool(true),
            Passwords: presetsdkgo.String("{"databases/{{DatabaseYAMLFile}}": "{{DatabasePassword}}"}"),
        },
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PostAPIV1DatasetImportRequest](../../models/operations/postapiv1datasetimportrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PostAPIV1DatasetImportResponse](../../models/operations/postapiv1datasetimportresponse.md), error**


## PutAPIV1DatasetDatasetID

Updates a Virtual Dataset.

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

- `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.
    

Replace in the body:

- `{{DBID}}` with the database `id` retrieved using the **Get Databases** endoint.
- `{{SchemaName}}` with the desired schema.
- `{{TableName}}` with the desired table.
- `{{DatasetSQL}}` with the SQL query to power your dataset.
- `{{OwnerID}}` with the `id` for the desired owner account(s) retrieved from the **Get all possible Dataset Owners** endpoint. Use a comma to separate multiple IDs (for example, `[2,5]`).

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.PutAPIV1DatasetDatasetID(ctx, operations.PutAPIV1DatasetDatasetIDRequest{
        DatasetID: "culpa",
        Referer: presetsdkgo.String("doloribus"),
        RequestBody: &operations.PutAPIV1DatasetDatasetIDRequestBody{},
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
| `request`                                                                                                | [operations.PutAPIV1DatasetDatasetIDRequest](../../models/operations/putapiv1datasetdatasetidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutAPIV1DatasetDatasetIDResponse](../../models/operations/putapiv1datasetdatasetidresponse.md), error**


## PutAPIV1DatasetDatasetIDRefresh

Refreshes and updates columns of a Dataset.

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

- `{{DatasetID}}` with the Dataset `id` retrieved from the **Get All Datasets from a Workspace** endpoint.

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatasets.PutAPIV1DatasetDatasetIDRefresh(ctx, operations.PutAPIV1DatasetDatasetIDRefreshRequest{
        DatasetID: "sapiente",
        Referer: presetsdkgo.String("architecto"),
        RequestBody: &operations.PutAPIV1DatasetDatasetIDRefreshRequestBody{},
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PutAPIV1DatasetDatasetIDRefreshRequest](../../models/operations/putapiv1datasetdatasetidrefreshrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PutAPIV1DatasetDatasetIDRefreshResponse](../../models/operations/putapiv1datasetdatasetidrefreshresponse.md), error**

