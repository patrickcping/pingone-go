# \ManagementAPIsApplicationsApplicationResourceGrantsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDApplicationsAppIDGrantsGet**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#V1EnvironmentsEnvIDApplicationsAppIDGrantsGet) | **Get** /v1/environments/{envID}/applications/{appID}/grants | READ All Grants for an Application
[**V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDelete**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDelete) | **Delete** /v1/environments/{envID}/applications/{appID}/grants/{grantID} | DELETE Grant
[**V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGet**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGet) | **Get** /v1/environments/{envID}/applications/{appID}/grants/{grantID} | READ One Grant for an Application
[**V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPut**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPut) | **Put** /v1/environments/{envID}/applications/{appID}/grants/{grantID} | UPDATE Grant
[**V1EnvironmentsEnvIDApplicationsAppIDGrantsPost**](ManagementAPIsApplicationsApplicationResourceGrantsApi.md#V1EnvironmentsEnvIDApplicationsAppIDGrantsPost) | **Post** /v1/environments/{envID}/applications/{appID}/grants | CREATE Grant



## V1EnvironmentsEnvIDApplicationsAppIDGrantsGet

> V1EnvironmentsEnvIDApplicationsAppIDGrantsGet(ctx, envID, appID).Execute()

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
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGet(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGet``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDGrantsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDelete

> V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDelete(ctx, envID, appID, grantID).Execute()

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
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDelete(context.Background(), envID, appID, grantID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGet

> V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGet(ctx, envID, appID, grantID).Execute()

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
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGet(context.Background(), envID, appID, grantID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPut

> V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPut(ctx, envID, appID, grantID).ContentType(contentType).Body(body).Execute()

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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPut(context.Background(), envID, appID, grantID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDGrantsGrantIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDApplicationsAppIDGrantsPost

> V1EnvironmentsEnvIDApplicationsAppIDGrantsPost(ctx, envID, appID).ContentType(contentType).Body(body).Execute()

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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsPost(context.Background(), envID, appID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationResourceGrantsApi.V1EnvironmentsEnvIDApplicationsAppIDGrantsPost``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDGrantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

