# Preset SDK


## Overview

Preset API: Welcome to the Preset API Collection.

## Overview

The Preset REST API is a powerful feature that provides access to core functional aspects of both Preset Manager and Superset.

The API supports the following areas of Preset:

*   User and Team Management
*   Workspace Management
*   Connection and Data Management
*   Visualization Management
*   Permissions Management \[Beta\]
    

This documentation lists the most relevant endpoints across all of the above functional areas of Preset.

## Authentication

### Generate an API Key

To generate an API key, from the Preset Manager screen (after log-in), hover your cursor over the initials icon and, in the sub-menu, select Manage User Settings.

![](https://i.ibb.co/D1xHc92/api1.png)

In the *API Keys* section, select **\+ Generate New API Key**

![](https://i.ibb.co/LRWp7HC/api2.png)

The *Generate a New API Key* panel appears.

In the **Key Title** field, enter a name for the new API key.

In the **Key Description** field, enter a brief descripton of the API key.

Select **Submit**.

![](https://i.ibb.co/cC0H4mY/api3.png)

The **Token** field will automatically populate with a generated token.

Likewise, the **Secret** field will automatically populate with a secret.

![](https://i.ibb.co/8smp5pZ/api5.png)

*Reminder: Safeguard the Secret**Please take a moment to select the Copy icon and then safely store it.*

When ready, select **OK**.

![](https://i.ibb.co/LdNDGNp/api6.png)

The newly-created API key appears in the *API Keys* section.

By default, the API key will be activated — to deactivate, toggle the **Active** slider to the off position.

To delete an API key, select the trash bin icon.

### Using the Postman Collection

All requests on this collection inherit the **token** specified on the **Preset API** collection.

To authenticate:

1.  Click on the **Preset API Collection**.
2.  Navigate to the **Variables** tab.
3.  Provide your **API Token** on the `APITokenName` current value.
4.  Provide the **API Token Secret** on the `APITokenSecret` current value.
    

These would be used to generate a **JWT Token** that's stored as a **Global Variable**.

There's a script defined on this collection, that is always executed before any request. The script basically checks if there's a valid **JWT Token** and tries to generate one/refresh it if needed.

* * *

### Manually using the API

Use the **Get a JWT Token** request to generate a `JWT Token.`

### Available Operations

