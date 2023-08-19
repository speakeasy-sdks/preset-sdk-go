# SupersetAPIsOpenSourceGreaterThanCharts

## Overview

APIs to manage Charts on your Workspace.

### Available Operations

* [DeleteAPIV1ChartChartID](#deleteapiv1chartchartid) - Delete a Chart
* [GetAPIV1Chart](#getapiv1chart) - Get all Charts from a Workspace
* [GetAPIV1ChartExport](#getapiv1chartexport) - Export Charts
* [GetAPIV1ChartRelatedOwners](#getapiv1chartrelatedowners) - Get all possible Chart Owners
* [GetAPIV1ChartChartID](#getapiv1chartchartid) - Get a Chart
* [GetAPIV1ChartChartIDCacheScreenshot](#getapiv1chartchartidcachescreenshot) - Get Chart Screenshot
* [GetAPIV1ChartChartIDData](#getapiv1chartchartiddata) - Get Chart's Data
* [PostAPIV1Chart](#postapiv1chart) - Create a Chart
* [PostAPIV1ChartData](#postapiv1chartdata) - Refresh a Chart
* [PostAPIV1ChartImport](#postapiv1chartimport) - Import a Chart
* [PutAPIV1ChartChartID](#putapiv1chartchartid) - Update a Chart

## DeleteAPIV1ChartChartID

Deletes a Chart.

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

- `{{ChartID}}` with the chart `id` retrieved from the **Get all Charts from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.DeleteAPIV1ChartChartID(ctx, operations.DeleteAPIV1ChartChartIDRequest{
        ChartID: "modi",
        Referer: preset.String("qui"),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteAPIV1ChartChartIDRequest](../../models/operations/deleteapiv1chartchartidrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.DeleteAPIV1ChartChartIDResponse](../../models/operations/deleteapiv1chartchartidresponse.md), error**


## GetAPIV1Chart

Gets all Charts from a Workspace.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.GetAPIV1Chart(ctx, operations.GetAPIV1ChartRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAPIV1ChartRequest](../../models/operations/getapiv1chartrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetAPIV1ChartResponse](../../models/operations/getapiv1chartresponse.md), error**


## GetAPIV1ChartExport

Exports Charts.

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

- `{{ChartIDs}` with comma separated chart `ids` retrieved from the **Get all Charts from a Workspace** endpoint.
    

_**Tip:**_ If used in Postman, select `Save Response` and `Save to a File` to get the zip export.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.GetAPIV1ChartExport(ctx, operations.GetAPIV1ChartExportRequest{
        Referer: preset.String("impedit"),
        Q: preset.String("cum"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAPIV1ChartExportRequest](../../models/operations/getapiv1chartexportrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetAPIV1ChartExportResponse](../../models/operations/getapiv1chartexportresponse.md), error**


## GetAPIV1ChartRelatedOwners

Gets all possible Chart owners on the Workspace.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.GetAPIV1ChartRelatedOwners(ctx, operations.GetAPIV1ChartRelatedOwnersRequest{})
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
| `request`                                                                                                    | [operations.GetAPIV1ChartRelatedOwnersRequest](../../models/operations/getapiv1chartrelatedownersrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetAPIV1ChartRelatedOwnersResponse](../../models/operations/getapiv1chartrelatedownersresponse.md), error**


## GetAPIV1ChartChartID

Gets a specific Chart from a Workspace.

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

- `{{ChartID}}` with the chart `id` retrieved from the  
    **Get all Charts from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.GetAPIV1ChartChartID(ctx, operations.GetAPIV1ChartChartIDRequest{
        ChartID: "esse",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAPIV1ChartChartIDRequest](../../models/operations/getapiv1chartchartidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetAPIV1ChartChartIDResponse](../../models/operations/getapiv1chartchartidresponse.md), error**


## GetAPIV1ChartChartIDCacheScreenshot

Retrieves a URL that can be used to download a Chart as an image.

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

- `{{ChartID}}` with the `id` retrieved from the **Get all Charts from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.GetAPIV1ChartChartIDCacheScreenshot(ctx, operations.GetAPIV1ChartChartIDCacheScreenshotRequest{
        ChartID: "ipsum",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetAPIV1ChartChartIDCacheScreenshotRequest](../../models/operations/getapiv1chartchartidcachescreenshotrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.GetAPIV1ChartChartIDCacheScreenshotResponse](../../models/operations/getapiv1chartchartidcachescreenshotresponse.md), error**


## GetAPIV1ChartChartIDData

Gets Chart's data from a Workspace.

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

- `{{ChartID}}` with the chart `id` retrieved from the **Get all Charts from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.GetAPIV1ChartChartIDData(ctx, operations.GetAPIV1ChartChartIDDataRequest{
        ChartID: "excepturi",
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
| `request`                                                                                                | [operations.GetAPIV1ChartChartIDDataRequest](../../models/operations/getapiv1chartchartiddatarequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetAPIV1ChartChartIDDataResponse](../../models/operations/getapiv1chartchartiddataresponse.md), error**


## PostAPIV1Chart

Creates a Chart through the API.

Replace in the URL:

*   `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
*   `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

Replace in the body:

*   `cache_timeout` (optional) by the chart timeout in seconds.
*   `certification_details` (optional) by the details of certification.
*   `certified_by` (optional) by the certifier.
*   `dashboards` by a comma separated list (csl) of dashboard ids the chart should be added to (can be empty).
*   `description` (optional) by your description.
*   `is_managed_externally` by true or false.
*   `owners` by a csl of owner ids.
*   `params` by a string of all parameters need to define your chart. You can get an example of this either by checking the network tab by creating a chart in the UI or by calling the **Get a Chart** endpoint.
*   `slice_name` by the name you want to give to your chart.
*   `viz_type` by the visualization type. You can find this with the **Get a Chart** endpoint for a similar chart.
*   `datasource_id` by the id of the dataset powering your chart.
*   `datasource_type` by the type of the underlying dataset.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.PostAPIV1Chart(ctx, operations.PostAPIV1ChartRequest{
        RequestBody: &operations.PostAPIV1ChartRequestBody{},
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostAPIV1ChartRequest](../../models/operations/postapiv1chartrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PostAPIV1ChartResponse](../../models/operations/postapiv1chartresponse.md), error**


## PostAPIV1ChartData

Refreshes the Chart data.

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

- `{{ChartID}}` with the `id` retrieved from the **Get all Charts from a Workspace** endpoint.
    

The body payload can vary depending on the **visualization type** used and also the **chart configuration**, The easiest way to get the accurate payload is retrieving the `query_context` data from the **Get a Chart** endpoint, and then set `force` to `true` both in the top level and also inside `form_data`.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.PostAPIV1ChartData(ctx, operations.PostAPIV1ChartDataRequest{
        RequestBody: &operations.PostAPIV1ChartDataRequestBody{},
        Force: preset.String("aspernatur"),
        SliceID: preset.String("perferendis"),
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
| `request`                                                                                    | [operations.PostAPIV1ChartDataRequest](../../models/operations/postapiv1chartdatarequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostAPIV1ChartDataResponse](../../models/operations/postapiv1chartdataresponse.md), error**


## PostAPIV1ChartImport

Imports a Chart via the API.

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

- Select your Chart ZIP file as a value for the `formData`.
- For the `passwords` field:
    - If the Database used by the Chart doesn't exist on the destination Workspace yet:
        - Replace `{{DatabaseYAMLFile}}` by the database YAML file name. You can find it in your Chart export file, under the `databases` folder.
        - Replace `{{DatabasePassword}}` by your DB password.
    - If the Database already exists on the destination Workspace, you can remove this field from the body.
- For the `overwrite` field:
    - If the Chart already exists on the destination Workspace, set it as `true` to overwrite it.
    - If the Chart doesn't exist in there yet, you can remove this field from the body.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.PostAPIV1ChartImport(ctx, operations.PostAPIV1ChartImportRequest{
        Referer: preset.String("ad"),
        RequestBody: &operations.PostAPIV1ChartImportRequestBody{
            FormData: &operations.PostAPIV1ChartImportRequestBodyFormData{
                Content: []byte("natus"),
                FormData: "sed",
            },
            Overwrite: preset.Bool(true),
            Passwords: preset.String("{"databases/{{DatabaseYAMLFile}}": "{{DatabasePassword}}"}"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PostAPIV1ChartImportRequest](../../models/operations/postapiv1chartimportrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PostAPIV1ChartImportResponse](../../models/operations/postapiv1chartimportresponse.md), error**


## PutAPIV1ChartChartID

Updates a Chart via the API.

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

- `{{ChartID}}` with the chart `id` retrieved from the  
    **Get all Charts from a Workspace** endpoint.
    

Replace in the body:

- `cache_timeout` (optional) by the chart timeout in seconds.
- `certification_details` (optional) by the details of certification.
- `certified_by` (optional) by the certifier.
- `dashboards` by a comma separated list (csl) of dashboard ids the chart should be added to (can be empty).
- `description` (optional) by your description.
- `is_managed_externally` by true or false.
- `owners` by a csl of owner ids.
- `params` by a string of all parameters need to define your chart. You can get an example of this either by checking the network tab by creating a chart in the UI or by calling the **Get a Chart** endpoint.
- `slice_name` by the name you want to give to your chart.
- `viz_type` by the visualization type. You can find this with the **Get a Chart** endpoint for a similar chart.
- `datasource_id` by the id of the dataset powering your chart.
- `datasource_type` by the type of the underlying dataset.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanCharts.PutAPIV1ChartChartID(ctx, operations.PutAPIV1ChartChartIDRequest{
        ChartID: "iste",
        Referer: preset.String("dolor"),
        RequestBody: &operations.PutAPIV1ChartChartIDRequestBody{},
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PutAPIV1ChartChartIDRequest](../../models/operations/putapiv1chartchartidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PutAPIV1ChartChartIDResponse](../../models/operations/putapiv1chartchartidresponse.md), error**

