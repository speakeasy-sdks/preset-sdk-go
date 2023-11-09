# SupersetAPIsOpenSourceGreaterThanAssets
(*SupersetAPIsOpenSourceGreaterThanAssets*)

## Overview

APIs to export/import an `assets` ZIP file from the Workspace, which includes all:

*   databases.
*   datasets.
*   charts.
*   saved queries.

### Available Operations

* [GetAPIV1AssetsExport](#getapiv1assetsexport) - Export Assets
* [PostAPIV1AssetsImport](#postapiv1assetsimport) - Import Assets

## GetAPIV1AssetsExport

Generates and export a ZIP file from the Workspace containing all:

*   Databases
*   Datasets
*   Charts
*   Dashboards
*   Saved Queries
    

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

***Tip:*** If used in Postman, select `Save Response` and `Save to a File` to get the zip export.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAssets.GetAPIV1AssetsExport(ctx, operations.GetAPIV1AssetsExportRequest{})
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
| `request`                                                                                            | [operations.GetAPIV1AssetsExportRequest](../../pkg/models/operations/getapiv1assetsexportrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetAPIV1AssetsExportResponse](../../pkg/models/operations/getapiv1assetsexportresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostAPIV1AssetsImport

Imports an `assets` ZIP file.

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

*   `{{DatabaseYAMLFile}}` by the database YAML File you can find in your chart export under the folder `Databases`.
*   `{{DatabasePassword}}` by your database password
*   Chose your Chart Export Zip file as a value for the `formData`.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanAssets.PostAPIV1AssetsImport(ctx, operations.PostAPIV1AssetsImportRequest{
        RequestBody: &operations.PostAPIV1AssetsImportRequestBody{
            Bundle: &operations.Bundle{
                Content: []byte("0x6a47A4beED"),
                FileName: "convertible_female.mpe",
            },
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostAPIV1AssetsImportRequest](../../pkg/models/operations/postapiv1assetsimportrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostAPIV1AssetsImportResponse](../../pkg/models/operations/postapiv1assetsimportresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
