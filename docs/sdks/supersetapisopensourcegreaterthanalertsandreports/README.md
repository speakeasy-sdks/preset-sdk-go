# SupersetAPIsOpenSourceGreaterThanAlertsAndReports
(*.SupersetAPIsOpenSourceGreaterThanAlertsAndReports*)

## Overview

APIs to manage your Alerts & Reports.

### Available Operations

* [GetAPIV1Report](#getapiv1report) - Get all Reports from a Workspace
* [GetAPIV1ReportInfo](#getapiv1reportinfo) - Get Alerts & Reports API metadata Info
* [GetAPIV1ReportAlertIDORReportID](#getapiv1reportalertidorreportid) - Get an Alert/Report
* [PostAPIV1Report](#postapiv1report) - Create a Dashboard Alert
* [PutAPIV1ReportAlertIDORReportID](#putapiv1reportalertidorreportid) - Disable an Alert/Report

## GetAPIV1Report

Gets all Reports created on the Workspace.

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

Note that the response includes a `count` value, indicating the total count of Alerts. 100 alerts would be included on the response - if `count > 100`, you can access the remaining items by increasing the `page` value on the `q` parameter:

```
?q=(filters:!((col:type,opr:eq,value:Report)),page_size:{{PageSize}},page:{{Page}})

```

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAlertsAndReports.GetAPIV1Report(ctx, operations.GetAPIV1ReportRequest{})
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
| `request`                                                                            | [operations.GetAPIV1ReportRequest](../../models/operations/getapiv1reportrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetAPIV1ReportResponse](../../models/operations/getapiv1reportresponse.md), error**


## GetAPIV1ReportInfo

Gets metadata information about the Alerts & Reports API endpoints.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAlertsAndReports.GetAPIV1ReportInfo(ctx, operations.GetAPIV1ReportInfoRequest{})
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
| `request`                                                                                    | [operations.GetAPIV1ReportInfoRequest](../../models/operations/getapiv1reportinforequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetAPIV1ReportInfoResponse](../../models/operations/getapiv1reportinforesponse.md), error**


## GetAPIV1ReportAlertIDORReportID

Gets a specific Alert/Report from the Workspace.

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

- `{{AlertID OR ReportID}}` with the `id` retrieved from one of the endpoints below:
    - **Get all Alerts and Reports from a Workspace**
    - **Get all Alerts from a Workspace**
    - **Get all Reports from a Workspace**

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAlertsAndReports.GetAPIV1ReportAlertIDORReportID(ctx, operations.GetAPIV1ReportAlertIDORReportIDRequest{})
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
| `request`                                                                                                              | [operations.GetAPIV1ReportAlertIDORReportIDRequest](../../models/operations/getapiv1reportalertidorreportidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.GetAPIV1ReportAlertIDORReportIDResponse](../../models/operations/getapiv1reportalertidorreportidresponse.md), error**


## PostAPIV1Report

Creates a Dashboard Alert on the Workspace through the API.

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

- `{{ActiveBooleanFlag}}` with:
    - **true** to create it enabled.
    - **false** to create it disabled.
- `{{AlertName}}` with a name for your alert.
- `{{DashboardID}}` with the `id` of the desired Dashboard, retrieved using the **Get all Dashboards from a Workspace** endpoint.
- `{{AlertDescription}}` with a description for it.
- `{{CRONSchedule}}` with the desired frequency (in [cron](https://en.wikipedia.org/wiki/Cron)).
- `{{Timezone}}` with the desired timezone. The list of valid options can be retrieved using the **Get Alerts & Reports API metadata Info** endpoint.
- `{{ForceBooleanFlag}}` with:
    - **true** to ignore cache.
    - **false** to use cache if available.
- `{{DatabaseID}}` with the `id` of the database that should be used to execute the SQL validation. You can retrieve this `id` using the **Get all Database Connections from a Workspace** endpoint.
- `{{SQLCondition}}` with the SQL query that should be validated by the alert. **Note that the SQL query should return only one column**, for example `select count(\\\\\\*) from {{MyTable}}`.
- `{{ValidatorType}}` with:
    - `operator` when performing number comparisson.
    - `not null` to check if the SQL result is not null. When using `not null`, the `validator_config_json` should be empty:

``` json
"validator_type": "not null",
"validator_config_json": {}

```

- `{{Operator}}` with the operation that should be used to analyze the SQL result. Available options:
    - `==` to check if SQL result is equal the threshold value.
    - `<` to check if the SQL result is smaller than the t
    - the threshold value.
    - `>` to check if the SQL result is larger than the the threshold value.
    - `<=` to check if the SQL result is not larger than the the threshold value.
    - `>=` to check if the SQL result is not smaller than the the threshold value.
    - `!=` to check if the SQL result is different than the threshold value.
- `{{Threshold}}` with the condition value.
- For the `owners` field:
    - Replace `{{OwnerID}}` with the owner's account ID on the Workspace level (you can retrieve this ID using the **Get all possible Chart Owners** endpoint).
    - This field is an array, so multiple owners can be added, separated by comma.
- For the `recipients` field:
    - Replace `{{ReportType}}` with `Email` or `Slack`.
    - `{{TargetInfo}}` with the email address/Slack user handler.
    - This field is an array, so multiple recipient configuration can be added (comma separated).
- `{{LogRetention}}` with the retention period (in days). Default and max value is `90`_._
- `{{WorkingTimeout}}` with time out settings (in seconds). Default value is `3600`.
- `{{GracePeriod}}` with a grace period (in seconds). Default value is `14400`.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAlertsAndReports.PostAPIV1Report(ctx, operations.PostAPIV1ReportRequest{
        RequestBody: &operations.PostAPIV1ReportRequestBody{},
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PostAPIV1ReportRequest](../../models/operations/postapiv1reportrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PostAPIV1ReportResponse](../../models/operations/postapiv1reportresponse.md), error**


## PutAPIV1ReportAlertIDORReportID

Disables a specific Alert/Report from the Workspace.

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

- `{{AlertID OR ReportID}}` with the `id` retrieved from one of the endpoints below:
    - **Get all Alerts and Reports from a Workspace**
    - **Get all Alerts from a Workspace**
    - **Get all Reports from a Workspace**

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAlertsAndReports.PutAPIV1ReportAlertIDORReportID(ctx, operations.PutAPIV1ReportAlertIDORReportIDRequest{
        RequestBody: &operations.PutAPIV1ReportAlertIDORReportIDRequestBody{},
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
| `request`                                                                                                              | [operations.PutAPIV1ReportAlertIDORReportIDRequest](../../models/operations/putapiv1reportalertidorreportidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.PutAPIV1ReportAlertIDORReportIDResponse](../../models/operations/putapiv1reportalertidorreportidresponse.md), error**

