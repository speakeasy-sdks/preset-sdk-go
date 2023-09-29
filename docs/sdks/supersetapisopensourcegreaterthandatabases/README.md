# SupersetAPIsOpenSourceGreaterThanDatabases
(*SupersetAPIsOpenSourceGreaterThanDatabases*)

## Overview

APIs to manage your database connections.

### Available Operations

* [CreateDatabaseUsingSSH](#createdatabaseusingssh) - Create a Database Connection using SSH
* [GetAPIV1Database](#getapiv1database) - Get all Database Connections from a Workspace
* [GetAPIV1DatabaseExport](#getapiv1databaseexport) - Export Database Connections
* [GetAPIV1DatabaseDatabaseID](#getapiv1databasedatabaseid) - Get a Database Connection
* [GetAPIV1DatabaseDatabaseIDConnection](#getapiv1databasedatabaseidconnection) - Get a Database Connection Parameters
* [PostAPIV1Database](#postapiv1database) - Create a Database Connection
* [PostAPIV1DatabaseImport](#postapiv1databaseimport) - Import a Database Connection
* [PutAPIV1DatabaseDatabaseID](#putapiv1databasedatabaseid) - Update a Database Connection

## CreateDatabaseUsingSSH

Creates a new database connection on the Workspace using SSH.

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

Replace on the body:

- `{{DBConnectionName}}` with the desired DB connection display name.
- `{{SQLAlchemyURI}}` with the [connection string](https://docs.preset.io/docs/uri-connection-strings) to the database.
- For the SSH configuration:
    - Use either a `username` and `password`, or `username`, `private_key` and `private_key_password` combination to authenticate to the SSH server, removing the un-used fields.
    - Replace `{{SSHServerAddress}}` with the SSH server host address.
- `boolean` with either `true` or `false` to enable/disable the connection settings.
- `{{ExtraConfiguration}}` _(optional)_ with any addifional configuration (like **Engine Parameters**). For example:
    

```
{\"engine_params\": {\"connect_args\": {\"http_path\": \"/sql/1.0/warehouses/********\"}}}

```

**Note:** You need to escape quotes (`"`) and other special characters using `\` on all body fields.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.CreateDatabaseUsingSSH(ctx, operations.CreateDatabaseUsingSSHRequest{
        Referer: presetsdkgo.String("methodical hemorrhage"),
        RequestBody: presetsdkgo.String("\"{\n    \\"database_name\\": \\"{{DBConnectionName}}\\",\n    \\"configuration_method\\": \\"sqlalchemy_form\\",\n    \\"sqlalchemy_uri\\": \\"{{SQLAlchemyURI}}\\",\n    \\"ssh_tunnel\\": {\n        \\"username\\": \\"{{SSHUsername}}\\",\n        \\"password\\": \\"{{SSHServerPassword}}\\",\n        \\"private_key\\": \\"{{PrivateKey}}\\",\n        \\"private_key_password\\": \\"{{PrivateKeyPassword}}\\",\n        \\"server_port\\": 22,\n        \\"server_address\\": \\"{{SSHServerAddress}}\\"\n    },\n    \\"allow_csv_upload\\": \\"boolean\\",\n    \\"allow_ctas\\": \\"boolean\\",\n    \\"allow_cvas\\": \\"boolean\\",\n    \\"allow_dml\\": \\"boolean\\",\n    \\"allow_multi_schema_metadata_fetch\\": \\"boolean\\",\n    \\"allow_run_async\\": \\"boolean\\",\n    \\"cache_timeout\\": 0,\n    \\"expose_in_sqllab\\": \\"boolean\\",\n    \\"impersonate_user\\": \\"boolean\\",\n    \\"extra\\": \\"{{ExtraConfiguration}}\\"\n}\""),
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
| `request`                                                                                            | [operations.CreateDatabaseUsingSSHRequest](../../models/operations/createdatabaseusingsshrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateDatabaseUsingSSHResponse](../../models/operations/createdatabaseusingsshresponse.md), error**


## GetAPIV1Database

Gets the databases connected to your Workspace.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.GetAPIV1Database(ctx, operations.GetAPIV1DatabaseRequest{})
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
| `request`                                                                                | [operations.GetAPIV1DatabaseRequest](../../models/operations/getapiv1databaserequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetAPIV1DatabaseResponse](../../models/operations/getapiv1databaseresponse.md), error**


## GetAPIV1DatabaseExport

Exports Database Connections from the Workspace.

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

*   `{{DatabaseIDs}` with comma separated DB `ids` retrieved from the **Get all Database Connections from a Workspace** endpoint.
    

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
        presetsdkgo.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.GetAPIV1DatabaseExport(ctx, operations.GetAPIV1DatabaseExportRequest{
        Q: presetsdkgo.String("Mercury Mercury"),
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
| `request`                                                                                            | [operations.GetAPIV1DatabaseExportRequest](../../models/operations/getapiv1databaseexportrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetAPIV1DatabaseExportResponse](../../models/operations/getapiv1databaseexportresponse.md), error**


## GetAPIV1DatabaseDatabaseID

Get a specific Database Connection from a Workspace.

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

- `{{DatabaseID}}` with the Database Connection `id` retrieved from the **Get All Database Connections from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.GetAPIV1DatabaseDatabaseID(ctx, operations.GetAPIV1DatabaseDatabaseIDRequest{
        DatabaseID: "quirkily neural Outdoors",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetAPIV1DatabaseDatabaseIDRequest](../../models/operations/getapiv1databasedatabaseidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetAPIV1DatabaseDatabaseIDResponse](../../models/operations/getapiv1databasedatabaseidresponse.md), error**


## GetAPIV1DatabaseDatabaseIDConnection

###### _Requires admin permission._

Get connection parameters from a specific Database Connection from a Workspace.

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

- `{{DatabaseID}}` with the Database Connection `id` retrieved from the **Get All Database Connections from a Workspace** endpoint.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.GetAPIV1DatabaseDatabaseIDConnection(ctx, operations.GetAPIV1DatabaseDatabaseIDConnectionRequest{
        DatabaseID: "Regional safely",
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
| `request`                                                                                                                        | [operations.GetAPIV1DatabaseDatabaseIDConnectionRequest](../../models/operations/getapiv1databasedatabaseidconnectionrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetAPIV1DatabaseDatabaseIDConnectionResponse](../../models/operations/getapiv1databasedatabaseidconnectionresponse.md), error**


## PostAPIV1Database

Creates a new database connection on the Workspace.

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

Replace on the body:

- `boolean` with either `true` or `false` to enable/disable the connection settings.
- `{{DBConnectionName}}` with the desired DB connection display name.
- `{{SQLAlchemyURI}}` with the [connection string](https://docs.preset.io/docs/uri-connection-strings) to the database.
- `{{ExtraConfiguration}}` _(optional)_ with any addifional configuration (like **Engine Parameters**). For example:
    

```
{\"engine_params\": {\"connect_args\": {\"http_path\": \"/sql/1.0/warehouses/********\"}}}

```

**Note:** You need to escape quotes (`"`) and other special characters using `\` on all body fields.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.PostAPIV1Database(ctx, operations.PostAPIV1DatabaseRequest{
        Referer: presetsdkgo.String("Convertible"),
        RequestBody: &operations.PostAPIV1DatabaseRequestBody{},
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PostAPIV1DatabaseRequest](../../models/operations/postapiv1databaserequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PostAPIV1DatabaseResponse](../../models/operations/postapiv1databaseresponse.md), error**


## PostAPIV1DatabaseImport

Imports a Database Connection.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.PostAPIV1DatabaseImport(ctx, operations.PostAPIV1DatabaseImportRequest{
        Referer: presetsdkgo.String("Account"),
        RequestBody: &operations.PostAPIV1DatabaseImportRequestBody{
            FormData: &operations.PostAPIV1DatabaseImportRequestBodyFormData{
                Content: []byte("mf(22;Sr]|"),
                FormData: "teal Facilitator",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostAPIV1DatabaseImportRequest](../../models/operations/postapiv1databaseimportrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostAPIV1DatabaseImportResponse](../../models/operations/postapiv1databaseimportresponse.md), error**


## PutAPIV1DatabaseDatabaseID

Updates an existing Database Connection.

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

- `{{DatabaseID}}` with the Database Connection `id` retrieved from the **Get All Database Connections from a Workspace** endpoint.
    

In the body:

Include the keys you would like to update. Refer to the **Create a Database Connection** payload to check supported values.

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
    res, err := s.SupersetAPIsOpenSourceGreaterThanDatabases.PutAPIV1DatabaseDatabaseID(ctx, operations.PutAPIV1DatabaseDatabaseIDRequest{
        DatabaseID: "male",
        RequestBody: presetsdkgo.String("\"{\n    //include the keys you would like to modify\n    \\"database_name\\": \\"{{DBConnectionName}}\\"\n}\""),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PutAPIV1DatabaseDatabaseIDRequest](../../models/operations/putapiv1databasedatabaseidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PutAPIV1DatabaseDatabaseIDResponse](../../models/operations/putapiv1databasedatabaseidresponse.md), error**

