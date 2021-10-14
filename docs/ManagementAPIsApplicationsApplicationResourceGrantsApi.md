# \ManagementAPIsApplicationsApplicationResourceGrantsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationGrant**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#CreateApplicationGrant) | **Post** /v1/environments/{envID}/applications/{appID}/grants | CREATE Grant
[**DeleteApplicationGrant**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#DeleteApplicationGrant) | **Delete** /v1/environments/{envID}/applications/{appID}/grants/{grantID} | DELETE Grant
[**ReadAllApplicationGrants**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#ReadAllApplicationGrants) | **Get** /v1/environments/{envID}/applications/{appID}/grants | READ All Grants for an Application
[**ReadOneApplicationGrant**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#ReadOneApplicationGrant) | **Get** /v1/environments/{envID}/applications/{appID}/grants/{grantID} | READ One Grant for an Application
[**UpdateApplicationGrant**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#UpdateApplicationGrant) | **Put** /v1/environments/{envID}/applications/{appID}/grants/{grantID} | UPDATE Grant



## CreateApplicationGrant

> ApplicationResourceGrant CreateApplicationGrant(ctx, envID, appID).ApplicationResourceGrant(applicationResourceGrant).Execute()

CREATE Grant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    applicationResourceGrant := *openapiclient.NewApplicationResourceGrant(*openapiclient.NewApplicationResourceGrantResource("Id_example"), []openapiclient.ApplicationResourceGrantScopes{*openapiclient.NewApplicationResourceGrantScopes("Id_example")}) // ApplicationResourceGrant |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.CreateApplicationGrant(context.Background(), envID, appID).ApplicationResourceGrant(applicationResourceGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.CreateApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationGrant`: ApplicationResourceGrant
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationResourceGrantsApi.CreateApplicationGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationResourceGrant** | [**ApplicationResourceGrant**](ApplicationResourceGrant.md) |  | 

### Return type

[**ApplicationResourceGrant**](ApplicationResourceGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationGrant

> DeleteApplicationGrant(ctx, envID, appID, grantID).Execute()

DELETE Grant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    grantID := "grantID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.DeleteApplicationGrant(context.Background(), envID, appID, grantID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.DeleteApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**grantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllApplicationGrants

> EntityArray ReadAllApplicationGrants(ctx, envID, appID).Execute()

READ All Grants for an Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.ReadAllApplicationGrants(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.ReadAllApplicationGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplicationGrants`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationResourceGrantsApi.ReadAllApplicationGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneApplicationGrant

> ApplicationResourceGrant ReadOneApplicationGrant(ctx, envID, appID, grantID).Execute()

READ One Grant for an Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    grantID := "grantID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.ReadOneApplicationGrant(context.Background(), envID, appID, grantID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.ReadOneApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationGrant`: ApplicationResourceGrant
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationResourceGrantsApi.ReadOneApplicationGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**grantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationResourceGrant**](ApplicationResourceGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationGrant

> ApplicationResourceGrant UpdateApplicationGrant(ctx, envID, appID, grantID).ApplicationResourceGrant(applicationResourceGrant).Execute()

UPDATE Grant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    grantID := "grantID_example" // string | 
    applicationResourceGrant := *openapiclient.NewApplicationResourceGrant(*openapiclient.NewApplicationResourceGrantResource("Id_example"), []openapiclient.ApplicationResourceGrantScopes{*openapiclient.NewApplicationResourceGrantScopes("Id_example")}) // ApplicationResourceGrant |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.UpdateApplicationGrant(context.Background(), envID, appID, grantID).ApplicationResourceGrant(applicationResourceGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.UpdateApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationGrant`: ApplicationResourceGrant
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationResourceGrantsApi.UpdateApplicationGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**grantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationResourceGrant** | [**ApplicationResourceGrant**](ApplicationResourceGrant.md) |  | 

### Return type

[**ApplicationResourceGrant**](ApplicationResourceGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

