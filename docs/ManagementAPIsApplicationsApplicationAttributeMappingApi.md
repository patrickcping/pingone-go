# \ManagementAPIsApplicationsApplicationAttributeMappingApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationAttributeMapping**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#CreateApplicationAttributeMapping) | **Post** /v1/environments/{envID}/applications/{appID}/attributes | CREATE Application Attribute Mapping
[**DeleteApplicationAttributeMapping**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#DeleteApplicationAttributeMapping) | **Delete** /v1/environments/{envID}/applications/{appID}/attributes/{attrMappingID} | DELETE Application Attribute Mapping
[**ReadAllApplicationAttributeMappings**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#ReadAllApplicationAttributeMappings) | **Get** /v1/environments/{envID}/applications/{appID}/attributes | READ All Application Attribute Mappings
[**ReadOneApplicationAttributeMapping**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#ReadOneApplicationAttributeMapping) | **Get** /v1/environments/{envID}/applications/{appID}/attributes/{attrMappingID} | READ One Application Attribute Mapping
[**UpdateApplicationAttributeMapping**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#UpdateApplicationAttributeMapping) | **Put** /v1/environments/{envID}/applications/{appID}/attributes/{attrMappingID} | UPDATE Application Attribute Mapping



## CreateApplicationAttributeMapping

> ApplicationAttributeMapping CreateApplicationAttributeMapping(ctx, envID, appID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()

CREATE Application Attribute Mapping



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
    applicationAttributeMapping := *openapiclient.NewApplicationAttributeMapping("Name_example", false, "Value_example") // ApplicationAttributeMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationAttributeMappingApi.CreateApplicationAttributeMapping(context.Background(), envID, appID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.CreateApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationAttributeMapping`: ApplicationAttributeMapping
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationAttributeMappingApi.CreateApplicationAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationAttributeMapping** | [**ApplicationAttributeMapping**](ApplicationAttributeMapping.md) |  | 

### Return type

[**ApplicationAttributeMapping**](ApplicationAttributeMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationAttributeMapping

> DeleteApplicationAttributeMapping(ctx, envID, appID, attrMappingID).Execute()

DELETE Application Attribute Mapping



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
    attrMappingID := "attrMappingID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationAttributeMappingApi.DeleteApplicationAttributeMapping(context.Background(), envID, appID, attrMappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.DeleteApplicationAttributeMapping``: %v\n", err)
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
**attrMappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationAttributeMappingRequest struct via the builder pattern


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


## ReadAllApplicationAttributeMappings

> EntityArray ReadAllApplicationAttributeMappings(ctx, envID, appID).Execute()

READ All Application Attribute Mappings



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplicationAttributeMappings`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationAttributeMappingsRequest struct via the builder pattern


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


## ReadOneApplicationAttributeMapping

> ApplicationAttributeMapping ReadOneApplicationAttributeMapping(ctx, envID, appID, attrMappingID).Execute()

READ One Application Attribute Mapping



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
    attrMappingID := "attrMappingID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationAttributeMappingApi.ReadOneApplicationAttributeMapping(context.Background(), envID, appID, attrMappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.ReadOneApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationAttributeMapping`: ApplicationAttributeMapping
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationAttributeMappingApi.ReadOneApplicationAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**attrMappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationAttributeMapping**](ApplicationAttributeMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationAttributeMapping

> ApplicationAttributeMapping UpdateApplicationAttributeMapping(ctx, envID, appID, attrMappingID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()

UPDATE Application Attribute Mapping



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
    attrMappingID := "attrMappingID_example" // string | 
    applicationAttributeMapping := *openapiclient.NewApplicationAttributeMapping("Name_example", false, "Value_example") // ApplicationAttributeMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationAttributeMappingApi.UpdateApplicationAttributeMapping(context.Background(), envID, appID, attrMappingID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.UpdateApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationAttributeMapping`: ApplicationAttributeMapping
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationAttributeMappingApi.UpdateApplicationAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**attrMappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationAttributeMapping** | [**ApplicationAttributeMapping**](ApplicationAttributeMapping.md) |  | 

### Return type

[**ApplicationAttributeMapping**](ApplicationAttributeMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

