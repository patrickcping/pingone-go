# \ManagementAPIsResourcesResourceScopesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDResourcesResourceIDScopesGet**](ManagementAPIsResourcesResourceScopesApi.md#V1EnvironmentsEnvIDResourcesResourceIDScopesGet) | **Get** /v1/environments/{envID}/resources/{resourceID}/scopes | READ All Scopes (Resource)
[**V1EnvironmentsEnvIDResourcesResourceIDScopesPost**](ManagementAPIsResourcesResourceScopesApi.md#V1EnvironmentsEnvIDResourcesResourceIDScopesPost) | **Post** /v1/environments/{envID}/resources/{resourceID}/scopes | CREATE PingOne access control scope
[**V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete**](ManagementAPIsResourcesResourceScopesApi.md#V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete) | **Delete** /v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID} | DELETE Scope
[**V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet**](ManagementAPIsResourcesResourceScopesApi.md#V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet) | **Get** /v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID} | READ One Scope
[**V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut**](ManagementAPIsResourcesResourceScopesApi.md#V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut) | **Put** /v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID} | UPDATE PingOne access control scope



## V1EnvironmentsEnvIDResourcesResourceIDScopesGet

> V1EnvironmentsEnvIDResourcesResourceIDScopesGet(ctx, envID, resourceID).Authorization(authorization).Execute()

READ All Scopes (Resource)



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
    resourceID := "resourceID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesGet(context.Background(), envID, resourceID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDScopesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDResourcesResourceIDScopesPost

> V1EnvironmentsEnvIDResourcesResourceIDScopesPost(ctx, envID, resourceID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE PingOne access control scope



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
    resourceID := "resourceID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesPost(context.Background(), envID, resourceID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDScopesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **authorization** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete

> V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete(ctx, envID, resourceID, scopeID).Authorization(authorization).Execute()

DELETE Scope



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
    resourceID := "resourceID_example" // string | 
    scopeID := "scopeID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete(context.Background(), envID, resourceID, scopeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 
**scopeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet

> V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet(ctx, envID, resourceID, scopeID).Authorization(authorization).Execute()

READ One Scope



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
    resourceID := "resourceID_example" // string | 
    scopeID := "scopeID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet(context.Background(), envID, resourceID, scopeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 
**scopeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut

> V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut(ctx, envID, resourceID, scopeID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE PingOne access control scope



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
    resourceID := "resourceID_example" // string | 
    scopeID := "scopeID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut(context.Background(), envID, resourceID, scopeID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.V1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 
**scopeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDScopesScopeIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
 **authorization** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

