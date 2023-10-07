# SupersetAPIsOpenSourceGreaterThanDashboards
(*SupersetAPIsOpenSourceGreaterThanDashboards*)

## Overview

APIs to manage your Dashboards.

### Available Operations

* [DeleteAPIV1DashboardDashboardID](#deleteapiv1dashboarddashboardid) - Delete Dashboard
* [GetAPIV1Dashboard](#getapiv1dashboard) - Get all Dashboards From a Workspace
* [GetAPIV1DashboardInfo](#getapiv1dashboardinfo) - Get Dashboard Info
* [GetAPIV1DashboardExport](#getapiv1dashboardexport) - Export Dashboards
* [GetAPIV1DashboardDashboardID](#getapiv1dashboarddashboardid) - Get a Dashboard
* [GetAPIV1DashboardDashboardIDCharts](#getapiv1dashboarddashboardidcharts) - Get Charts from a Dashboard
* [GetAPIV1DashboardDashboardIDDatasets](#getapiv1dashboarddashboardiddatasets) - Get Datasets from a Dashboard
* [PostAPIV1Dashboard](#postapiv1dashboard) - Create a Dashboard
* [PostAPIV1DashboardImport](#postapiv1dashboardimport) - Import a Dashboard
* [PostAPIV1DashboardDashboardIDPermalink](#postapiv1dashboarddashboardidpermalink) - Create a Permalink to a Dashboard
* [PutAPIV1DashboardDashboardID](#putapiv1dashboarddashboardid) - Update a Dashboard

## DeleteAPIV1DashboardDashboardID

Deletes a Dashboard.

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

- `{{DashboardID}}` with the `id` of the desired Dashboard. You can get the `id` using the **Get all Dashboards From a Workspace** endpoint. Alternatively, you can get from the Dashboard URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/dashboard/{{DashboardID}}`.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.DeleteAPIV1DashboardDashboardID(ctx, operations.DeleteAPIV1DashboardDashboardIDRequest{
        DashboardID: "maiores Bentley",
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
| `request`                                                                                                              | [operations.DeleteAPIV1DashboardDashboardIDRequest](../../models/operations/deleteapiv1dashboarddashboardidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.DeleteAPIV1DashboardDashboardIDResponse](../../models/operations/deleteapiv1dashboarddashboardidresponse.md), error**


## GetAPIV1Dashboard

Gets all Dashboards from a Workspace.

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
	"github.com/speakeasy-sdks/preset-sdk-go/pkg/models/operations"
)

func main() {
    s := presetsdkgo.New(
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.GetAPIV1Dashboard(ctx, operations.GetAPIV1DashboardRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAPIV1DashboardRequest](../../models/operations/getapiv1dashboardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetAPIV1DashboardResponse](../../models/operations/getapiv1dashboardresponse.md), error**


## GetAPIV1DashboardInfo

Gets Dashboard info.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.GetAPIV1DashboardInfo(ctx, operations.GetAPIV1DashboardInfoRequest{})
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
| `request`                                                                                          | [operations.GetAPIV1DashboardInfoRequest](../../models/operations/getapiv1dashboardinforequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetAPIV1DashboardInfoResponse](../../models/operations/getapiv1dashboardinforesponse.md), error**


## GetAPIV1DashboardExport

Exports a ZIP file from a Dashboard.

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

- `{{DashboardIDs}}` with the `id`s of the Dashboard(s) you want to export (separated by comma). You can get the `id`s using the **Get all Dashboards From a Workspace** endpoint. Alternatively, you can get from the Dashboard URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/dashboard/{{DashboardID}}`.
    

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.GetAPIV1DashboardExport(ctx, operations.GetAPIV1DashboardExportRequest{})
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
| `request`                                                                                              | [operations.GetAPIV1DashboardExportRequest](../../models/operations/getapiv1dashboardexportrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetAPIV1DashboardExportResponse](../../models/operations/getapiv1dashboardexportresponse.md), error**


## GetAPIV1DashboardDashboardID

Get a specific Dashboard from a Workspace.

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

- `{{DashboardID}}` with the `id` of the desired Dashboard. You can get the `id` using the **Get all Dashboards From a Workspace** endpoint. Alternatively, you can get from the Dashboard URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/dashboard/{{DashboardID}}`.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.GetAPIV1DashboardDashboardID(ctx, operations.GetAPIV1DashboardDashboardIDRequest{
        DashboardID: "North",
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
| `request`                                                                                                        | [operations.GetAPIV1DashboardDashboardIDRequest](../../models/operations/getapiv1dashboarddashboardidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetAPIV1DashboardDashboardIDResponse](../../models/operations/getapiv1dashboarddashboardidresponse.md), error**


## GetAPIV1DashboardDashboardIDCharts

Gets all Charts associated with a Dashboard.

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

- `{{DashboardID}}` with the `id`s of the desired Dashboard. You can get the `id`s using the **Get all Dashboards From a Workspace** endpoint. Alternatively, you can get from the Dashboard URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/dashboard/{{DashboardID}}`.
    

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.GetAPIV1DashboardDashboardIDCharts(ctx, operations.GetAPIV1DashboardDashboardIDChartsRequest{
        DashboardID: "loiter partnerships Gasoline",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetAPIV1DashboardDashboardIDChartsRequest](../../models/operations/getapiv1dashboarddashboardidchartsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetAPIV1DashboardDashboardIDChartsResponse](../../models/operations/getapiv1dashboarddashboardidchartsresponse.md), error**


## GetAPIV1DashboardDashboardIDDatasets

Gets all Datasets associated with a Dashboard.

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

- `{{DashboardID}}` with the `id`s of the desired Dashboard. You can get the `id`s using the **Get all Dashboards From a Workspace** endpoint. Alternatively, you can get from the Dashboard URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/dashboard/{{DashboardID}}`.
    

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.GetAPIV1DashboardDashboardIDDatasets(ctx, operations.GetAPIV1DashboardDashboardIDDatasetsRequest{
        DashboardID: "furthermore Computers Communications",
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
| `request`                                                                                                                        | [operations.GetAPIV1DashboardDashboardIDDatasetsRequest](../../models/operations/getapiv1dashboarddashboardiddatasetsrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetAPIV1DashboardDashboardIDDatasetsResponse](../../models/operations/getapiv1dashboarddashboardiddatasetsresponse.md), error**


## PostAPIV1Dashboard

Creates a Dashboard using the API.

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

- `certification_details` (optional) by the details of certification
- `certified_by` (optional) by the certifier
- `css` (optional) add any css to the dashboard in a string here

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.PostAPIV1Dashboard(ctx, operations.PostAPIV1DashboardRequest{
        RequestBody: &operations.PostAPIV1DashboardRequestBody{},
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
| `request`                                                                                    | [operations.PostAPIV1DashboardRequest](../../models/operations/postapiv1dashboardrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostAPIV1DashboardResponse](../../models/operations/postapiv1dashboardresponse.md), error**


## PostAPIV1DashboardImport

Imports a Dashboard via the API.

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

- Select your Dashboard ZIP file as a value for the `formData`.
- For the `passwords` field:
    - If the Database used by the Dashboard doesn't exist on the destination Workspace yet:
        - Replace `{{DatabaseYAMLFile}}` by the Database YAML file name. You can find it in your Dashboard export file, under the `databases` folder.
        - Replace `{{DatabasePassword}}` by your DB password.
    - If the Database already exists on the destination Workspace, you can remove this field from the body.
- For the `overwrite` field:
    - If the Dashboard already exists on the destination Workspace, set it as `true` to overwrite it.
    - If the Dashboard doesn't exist in there yet, you can remove this field from the body.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.PostAPIV1DashboardImport(ctx, operations.PostAPIV1DashboardImportRequest{
        RequestBody: &operations.PostAPIV1DashboardImportRequestBody{
            FormData: &operations.PostAPIV1DashboardImportRequestBodyFormData{
                Content: []byte("#J8]!0YaF/"),
                FormData: "bus fooey",
            },
            Overwrite: presetsdkgo.Bool(true),
            Passwords: presetsdkgo.String("{\"databases/{{DatabaseYAMLFile}}\": \"{{DatabasePassword}}\"}"),
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostAPIV1DashboardImportRequest](../../models/operations/postapiv1dashboardimportrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PostAPIV1DashboardImportResponse](../../models/operations/postapiv1dashboardimportresponse.md), error**


## PostAPIV1DashboardDashboardIDPermalink

Creates a permalink to a Dashboard (with applied filters) using the API.

Replace in the URL and on the `Referer` header:

- `{{WorkspaceSlug}}` with the `name` retrieved through the API (using the **Get Workspaces from a Team** endpoint).
- `{{WorkspaceRegion}}` corresponding to the `region` retrieved from the **Get Workspaces from a Team** endpoint. Refer to below table:
    

| **`region`** | **`WorkspaceRegion`** |
| --- | --- |
| us-east-1 | `us2a` |
| us-west-2 | `us1a` |
| eu-north-1 | `eu5a` |
| ap-northeast-1 | `ap1a` |

Alternatively, access the Workspace through the UI, and get the `{{WorkspaceSlug}}` and `{{WorkspaceRegion}}` from the URL -> `https://{{WorkspaceSlug}}.{{Region}}.preset.io/superset/welcome/`.

Replace in the body:

- `{{FilterID}}` with the ID of the filter you want to modify. You can retrieve it either via the **Dashboard JSON Metadata**, or via the **Get a Dashboard** endpoint.
- `{{Column}}` with the column that is used on the filter.
- `{{Operator}}` with the filtering operation to be applied. Available options:
    - `IN`
    - `NOT IN`
- `{{Value}}` with the value to be applied on the filter.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.PostAPIV1DashboardDashboardIDPermalink(ctx, operations.PostAPIV1DashboardDashboardIDPermalinkRequest{
        DashboardID: "Northeast male Division",
        RequestBody: &operations.PostAPIV1DashboardDashboardIDPermalinkRequestBody{},
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
| `request`                                                                                                                            | [operations.PostAPIV1DashboardDashboardIDPermalinkRequest](../../models/operations/postapiv1dashboarddashboardidpermalinkrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.PostAPIV1DashboardDashboardIDPermalinkResponse](../../models/operations/postapiv1dashboarddashboardidpermalinkresponse.md), error**


## PutAPIV1DashboardDashboardID

Update a Dashboard

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDashboards.PutAPIV1DashboardDashboardID(ctx, operations.PutAPIV1DashboardDashboardIDRequest{
        DashboardID: "counterbalance",
        RequestBody: presetsdkgo.String("\"{ \\"position_json\\":\\"{\\\\"CHART-1NOOLm5YPs\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-1NOOLm5YPs\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 20, \\\\"height\\\\": 25, \\\\"sliceName\\\\": \\\\"Total Items Sold\\\\", \\\\"sliceNameOverride\\\\": \\\\"Total Products Sold\\\\", \\\\"uuid\\\\": \\\\"c3d643cd-fd6f-4659-a5b7-59402487a8d0\\\\", \\\\"width\\\\": 2}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-Tyv02UA_6W\\\\", \\\\"COLUMN-8Rp54B6ikC\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-AYpv8gFi_q\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-AYpv8gFi_q\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 43, \\\\"height\\\\": 91, \\\\"sliceName\\\\": \\\\"Number of Deals (for each Combination)\\\\", \\\\"uuid\\\\": \\\\"bd20fc69-dd51-46c1-99b5-09e37a434bf1\\\\", \\\\"width\\\\": 3}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-0l1WcDzW3\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-KKT9BsnUst\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-KKT9BsnUst\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 33, \\\\"height\\\\": 59, \\\\"sliceName\\\\": \\\\"Quarterly Sales (By Product Line)\\\\", \\\\"sliceNameOverride\\\\": \\\\"Quarterly Revenue (By Product Line)\\\\", \\\\"uuid\\\\": \\\\"db9609e4-9b78-4a32-87a7-4d9e19d51cd8\\\\", \\\\"width\\\\": 7}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-oAtmu5grZ\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-OJ9aWDmn1q\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-OJ9aWDmn1q\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 36, \\\\"height\\\\": 91, \\\\"sliceName\\\\": \\\\"Proportion of Revenue by Product Line\\\\", \\\\"sliceNameOverride\\\\": \\\\"Proportion of Monthly Revenue by Product Line\\\\", \\\\"uuid\\\\": \\\\"08aff161-f60c-4cb3-a225-dc9b1140d2e3\\\\", \\\\"width\\\\": 6}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-0l1WcDzW3\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-YFg-9wHE7s\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-YFg-9wHE7s\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 24, \\\\"height\\\\": 63, \\\\"sliceName\\\\": \\\\"Seasonality of Revenue (per Product Line)\\\\", \\\\"uuid\\\\": \\\\"cf0da099-b3ab-4d94-ab62-cf353ac3c611\\\\", \\\\"width\\\\": 6}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-E7MDSGfnm\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-_LMKI0D3tj\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-_LMKI0D3tj\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 30, \\\\"height\\\\": 62, \\\\"sliceName\\\\": \\\\"Revenue by Deal SIze\\\\", \\\\"sliceNameOverride\\\\": \\\\"Monthly Revenue by Deal SIze\\\\", \\\\"uuid\\\\": \\\\"f065a533-2e13-42b9-bd19-801a21700dff\\\\", \\\\"width\\\\": 6}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-E7MDSGfnm\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-id4RGv80N-\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-id4RGv80N-\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 19, \\\\"height\\\\": 40, \\\\"sliceName\\\\": \\\\"Total Items Sold (By Product Line)\\\\", \\\\"sliceNameOverride\\\\": \\\\"Total Products Sold (By Product Line)\\\\", \\\\"uuid\\\\": \\\\"b8b7ca30-6291-44b0-bc64-ba42e2892b86\\\\", \\\\"width\\\\": 2}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-oAtmu5grZ\\\\", \\\\"COLUMN-G6_2DvG8aK\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-iyvXMcqHt9\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-iyvXMcqHt9\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 8, \\\\"height\\\\": 39, \\\\"sliceName\\\\": \\\\"Filter\\\\", \\\\"uuid\\\\": \\\\"a5689df7-98fc-7c51-602c-ebd92dc3ec70\\\\", \\\\"width\\\\": 2}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-0l1WcDzW3\\\\", \\\\"COLUMN-jlNWyWCfTC\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-j24u8ve41b\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-j24u8ve41b\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 41, \\\\"height\\\\": 59, \\\\"sliceName\\\\": \\\\"Overall Sales (By Product Line)\\\\", \\\\"sliceNameOverride\\\\": \\\\"Total Revenue (By Product Line)\\\\", \\\\"uuid\\\\": \\\\"09c497e0-f442-1121-c9e7-671e37750424\\\\", \\\\"width\\\\": 3}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-oAtmu5grZ\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-lFanAaYKBK\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-lFanAaYKBK\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 18, \\\\"height\\\\": 26, \\\\"sliceName\\\\": \\\\"Total Revenue\\\\", \\\\"uuid\\\\": \\\\"7b12a243-88e0-4dc5-ac33-9a840bb0ac5a\\\\", \\\\"width\\\\": 3}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-Tyv02UA_6W\\\\", \\\\"COLUMN-8Rp54B6ikC\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"CHART-vomBOiI7U9\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"CHART-vomBOiI7U9\\\\", \\\\"meta\\\\": {\\\\"chartId\\\\": 34, \\\\"height\\\\": 53, \\\\"sliceName\\\\": \\\\"Quarterly Sales\\\\", \\\\"sliceNameOverride\\\\": \\\\"Quarterly Revenue\\\\", \\\\"uuid\\\\": \\\\"692aca26-a526-85db-c94c-411c91cc1077\\\\", \\\\"width\\\\": 7}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-Tyv02UA_6W\\\\"], \\\\"type\\\\": \\\\"CHART\\\\"}, \\\\"COLUMN-8Rp54B6ikC\\\\": {\\\\"children\\\\": [\\\\"CHART-lFanAaYKBK\\\\", \\\\"CHART-1NOOLm5YPs\\\\"], \\\\"id\\\\": \\\\"COLUMN-8Rp54B6ikC\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\", \\\\"width\\\\": 2}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-Tyv02UA_6W\\\\"], \\\\"type\\\\": \\\\"COLUMN\\\\"}, \\\\"COLUMN-G6_2DvG8aK\\\\": {\\\\"children\\\\": [\\\\"CHART-id4RGv80N-\\\\"], \\\\"id\\\\": \\\\"COLUMN-G6_2DvG8aK\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\", \\\\"width\\\\": 2}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-oAtmu5grZ\\\\"], \\\\"type\\\\": \\\\"COLUMN\\\\"}, \\\\"COLUMN-jlNWyWCfTC\\\\": {\\\\"children\\\\": [\\\\"MARKDOWN-HrzsMmvGQo\\\\", \\\\"CHART-iyvXMcqHt9\\\\"], \\\\"id\\\\": \\\\"COLUMN-jlNWyWCfTC\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\", \\\\"width\\\\": 3}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-0l1WcDzW3\\\\"], \\\\"type\\\\": \\\\"COLUMN\\\\"}, \\\\"DASHBOARD_VERSION_KEY\\\\": \\\\"v2\\\\", \\\\"GRID_ID\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"GRID_ID\\\\", \\\\"parents\\\\": [\\\\"ROOT_ID\\\\"], \\\\"type\\\\": \\\\"GRID\\\\"}, \\\\"HEADER_ID\\\\": {\\\\"id\\\\": \\\\"HEADER_ID\\\\", \\\\"meta\\\\": {\\\\"text\\\\": \\\\"Sales Dashboard\\\\"}, \\\\"type\\\\": \\\\"HEADER\\\\"}, \\\\"MARKDOWN--AtDSWnapE\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"MARKDOWN--AtDSWnapE\\\\", \\\\"meta\\\\": {\\\\"code\\\\": \\\\"# \\�\\� Vehicle Sales Dashboard \\�\\�\\\\n\\\\nThis example dashboard provides insight into the business operations of vehicle seller. The dataset powering this dashboard can be found [here on Kaggle](https://www.kaggle.com/kyanyoga/sample-sales-data).\\\\n\\\\n### Timeline\\\\n\\\\nThe dataset contains data on all orders from the 2003 and 2004 fiscal years, and some orders from 2005.\\\\n\\\\n### Products Sold\\\\n\\\\nThis shop mainly sells the following products:\\\\n\\\\n- \\�\\� Classic Cars\\\\n- \\�\\�\\️ Vintage Cars\\\\n- \\�\\�\\️ Motorcycles\\\\n- \\�\\� Trucks & Buses \\�\\�\\\\n- \\�\\�\\️ Planes\\\\n- \\�\\� Ships\\\\n- \\�\\� Trains\\\\", \\\\"height\\\\": 53, \\\\"width\\\\": 3}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"ROW-Tyv02UA_6W\\\\"], \\\\"type\\\\": \\\\"MARKDOWN\\\\"}, \\\\"MARKDOWN-HrzsMmvGQo\\\\": {\\\\"children\\\\": [], \\\\"id\\\\": \\\\"MARKDOWN-HrzsMmvGQo\\\\", \\\\"meta\\\\": {\\\\"code\\\\": \\\\"# \\�\\� Filter Box\\\\n\\\\nDashboard filters are a powerful way to enable teams to dive deeper into their business operations data. This filter box helps focus the charts along the following variables:\\\\n\\\\n- Time Range: Focus in on a specific time period (e.g. a holiday or quarter)\\\\n- Product Line: Choose 1 or more product lines to see relevant sales data\\\\n- Deal Size: Zoom in on small, medium, and / or large sales deals\\\\n\\\\nThe filter box below \\�\\� is configured to only apply to the charts in this tab (**Exploratory**). You can customize the charts that this filter box applies to by:\\\\n\\\\n- entering Edit mode in this dashboard\\\\n- selecting the `...` in the top right corner\\\\n- selecting the **Set filter mapping** button\\\\", \\\\"height\\\\": 50, \\\\"width\\\\": 3}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\", \\\\"ROW-0l1WcDzW3\\\\", \\\\"COLUMN-jlNWyWCfTC\\\\"], \\\\"type\\\\": \\\\"MARKDOWN\\\\"}, \\\\"ROOT_ID\\\\": {\\\\"children\\\\": [\\\\"TABS-e5Ruro0cjP\\\\"], \\\\"id\\\\": \\\\"ROOT_ID\\\\", \\\\"type\\\\": \\\\"ROOT\\\\"}, \\\\"ROW-0l1WcDzW3\\\\": {\\\\"children\\\\": [\\\\"COLUMN-jlNWyWCfTC\\\\", \\\\"CHART-OJ9aWDmn1q\\\\", \\\\"CHART-AYpv8gFi_q\\\\"], \\\\"id\\\\": \\\\"ROW-0l1WcDzW3\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\"}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\"], \\\\"type\\\\": \\\\"ROW\\\\"}, \\\\"ROW-E7MDSGfnm\\\\": {\\\\"children\\\\": [\\\\"CHART-YFg-9wHE7s\\\\", \\\\"CHART-_LMKI0D3tj\\\\"], \\\\"id\\\\": \\\\"ROW-E7MDSGfnm\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\"}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-4fthLQmdX\\\\"], \\\\"type\\\\": \\\\"ROW\\\\"}, \\\\"ROW-Tyv02UA_6W\\\\": {\\\\"children\\\\": [\\\\"COLUMN-8Rp54B6ikC\\\\", \\\\"CHART-vomBOiI7U9\\\\", \\\\"MARKDOWN--AtDSWnapE\\\\"], \\\\"id\\\\": \\\\"ROW-Tyv02UA_6W\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\"}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\"], \\\\"type\\\\": \\\\"ROW\\\\"}, \\\\"ROW-oAtmu5grZ\\\\": {\\\\"children\\\\": [\\\\"COLUMN-G6_2DvG8aK\\\\", \\\\"CHART-KKT9BsnUst\\\\", \\\\"CHART-j24u8ve41b\\\\"], \\\\"id\\\\": \\\\"ROW-oAtmu5grZ\\\\", \\\\"meta\\\\": {\\\\"background\\\\": \\\\"BACKGROUND_TRANSPARENT\\\\"}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\", \\\\"TAB-d-E0Zc1cTH\\\\"], \\\\"type\\\\": \\\\"ROW\\\\"}, \\\\"TAB-4fthLQmdX\\\\": {\\\\"children\\\\": [\\\\"ROW-0l1WcDzW3\\\\", \\\\"ROW-E7MDSGfnm\\\\"], \\\\"id\\\\": \\\\"TAB-4fthLQmdX\\\\", \\\\"meta\\\\": {\\\\"text\\\\": \\\\"\\�\\� Exploratory\\\\"}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\"], \\\\"type\\\\": \\\\"TAB\\\\"}, \\\\"TAB-d-E0Zc1cTH\\\\": {\\\\"children\\\\": [\\\\"ROW-Tyv02UA_6W\\\\", \\\\"ROW-oAtmu5grZ\\\\"], \\\\"id\\\\": \\\\"TAB-d-E0Zc1cTH\\\\", \\\\"meta\\\\": {\\\\"text\\\\": \\\\"\\�\\� Sales Overview\\\\"}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\", \\\\"TABS-e5Ruro0cjP\\\\"], \\\\"type\\\\": \\\\"TAB\\\\"}, \\\\"TABS-e5Ruro0cjP\\\\": {\\\\"children\\\\": [\\\\"TAB-d-E0Zc1cTH\\\\", \\\\"TAB-4fthLQmdX\\\\"], \\\\"id\\\\": \\\\"TABS-e5Ruro0cjP\\\\", \\\\"meta\\\\": {}, \\\\"parents\\\\": [\\\\"ROOT_ID\\\\"], \\\\"type\\\\": \\\\"TABS\\\\"}}\\"\n\n    }\""),
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
| `request`                                                                                                        | [operations.PutAPIV1DashboardDashboardIDRequest](../../models/operations/putapiv1dashboarddashboardidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PutAPIV1DashboardDashboardIDResponse](../../models/operations/putapiv1dashboarddashboardidresponse.md), error**

