# SupersetAPIsOpenSourceGreaterThanSQLLab
(*.SupersetAPIsOpenSourceGreaterThanSQLLab*)

### Available Operations

* [PostAPIV1SqllabExecute](#postapiv1sqllabexecute) - Execute a SQL Query

## PostAPIV1SqllabExecute

Executes a SQL query through the API.

Replace in the URL and in the `Referer` header:

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

- `{{DBID}}` with the `id` retrieved from the **Get all Database Connections from a Workspace** endpoint.
- `{{SQLQuery}}` with the desired SQL query to be executed. Don't forget to escape quotes with `\`. For example:
    

``` json
{
    "database_id": 1,
    "sql": "SELECT * FROM \"Vehicle Sales\" limit 7" 
  }

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanSQLLab.PostAPIV1SqllabExecute(ctx, operations.PostAPIV1SqllabExecuteRequest{
        RequestBody: &operations.PostAPIV1SqllabExecuteRequestBody{},
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
| `request`                                                                                            | [operations.PostAPIV1SqllabExecuteRequest](../../models/operations/postapiv1sqllabexecuterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PostAPIV1SqllabExecuteResponse](../../models/operations/postapiv1sqllabexecuteresponse.md), error**

