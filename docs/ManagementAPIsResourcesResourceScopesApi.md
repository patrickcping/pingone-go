# \ManagementAPIsResourcesResourceScopesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceScope**](ManagementAPIsResourcesResourceScopesApi.md#CreateResourceScope) | **Post** /v1/environments/{envID}/resources/{resourceID}/scopes | CREATE PingOne access control scope
[**DeleteResourceScope**](ManagementAPIsResourcesResourceScopesApi.md#DeleteResourceScope) | **Delete** /v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID} | DELETE Scope
[**ReadAllResourceScopes**](ManagementAPIsResourcesResourceScopesApi.md#ReadAllResourceScopes) | **Get** /v1/environments/{envID}/resources/{resourceID}/scopes | READ All Scopes (Resource)
[**ReadOneResourceScope**](ManagementAPIsResourcesResourceScopesApi.md#ReadOneResourceScope) | **Get** /v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID} | READ One Scope
[**UpdateResourceScope**](ManagementAPIsResourcesResourceScopesApi.md#UpdateResourceScope) | **Put** /v1/environments/{envID}/resources/{resourceID}/scopes/{scopeID} | UPDATE PingOne access control scope



## CreateResourceScope

> ResourceScope CreateResourceScope(ctx, envID, resourceID).ResourceScope(resourceScope).Execute()

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
    resourceScope := *openapiclient.NewResourceScope("Name_example") // ResourceScope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsResourcesResourceScopesApi.CreateResourceScope(context.Background(), envID, resourceID).ResourceScope(resourceScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.CreateResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceScope`: ResourceScope
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsResourcesResourceScopesApi.CreateResourceScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceScope** | [**ResourceScope**](ResourceScope.md) |  | 

### Return type

[**ResourceScope**](ResourceScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceScope

> DeleteResourceScope(ctx, envID, resourceID, scopeID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsResourcesResourceScopesApi.DeleteResourceScope(context.Background(), envID, resourceID, scopeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.DeleteResourceScope``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteResourceScopeRequest struct via the builder pattern


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


## ReadAllResourceScopes

> EntityArray ReadAllResourceScopes(ctx, envID, resourceID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsResourcesResourceScopesApi.ReadAllResourceScopes(context.Background(), envID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.ReadAllResourceScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllResourceScopes`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsResourcesResourceScopesApi.ReadAllResourceScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllResourceScopesRequest struct via the builder pattern


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


## ReadOneResourceScope

> ResourceScope ReadOneResourceScope(ctx, envID, resourceID, scopeID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsResourcesResourceScopesApi.ReadOneResourceScope(context.Background(), envID, resourceID, scopeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.ReadOneResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneResourceScope`: ResourceScope
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsResourcesResourceScopesApi.ReadOneResourceScope`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadOneResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceScope**](ResourceScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceScope

> ResourceScope UpdateResourceScope(ctx, envID, resourceID, scopeID).ResourceScope(resourceScope).Execute()

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
    resourceScope := *openapiclient.NewResourceScope("Name_example") // ResourceScope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsResourcesResourceScopesApi.UpdateResourceScope(context.Background(), envID, resourceID, scopeID).ResourceScope(resourceScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceScopesApi.UpdateResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourceScope`: ResourceScope
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsResourcesResourceScopesApi.UpdateResourceScope`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceScope** | [**ResourceScope**](ResourceScope.md) |  | 

### Return type

[**ResourceScope**](ResourceScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

