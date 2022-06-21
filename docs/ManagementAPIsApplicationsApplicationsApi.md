# \ManagementAPIsApplicationsApplicationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ManagementAPIsApplicationsApplicationsApi.md#CreateApplication) | **Post** /v1/environments/{envID}/applications | CREATE Application
[**DeleteApplication**](ManagementAPIsApplicationsApplicationsApi.md#DeleteApplication) | **Delete** /v1/environments/{envID}/applications/{appID} | DELETE Application
[**ReadAllApplications**](ManagementAPIsApplicationsApplicationsApi.md#ReadAllApplications) | **Get** /v1/environments/{envID}/applications | READ All Applications
[**ReadOneApplication**](ManagementAPIsApplicationsApplicationsApi.md#ReadOneApplication) | **Get** /v1/environments/{envID}/applications/{appID} | READ One Application
[**UpdateApplication**](ManagementAPIsApplicationsApplicationsApi.md#UpdateApplication) | **Put** /v1/environments/{envID}/applications/{appID} | UPDATE Application



## CreateApplication

> CreateApplication201Response CreateApplication(ctx, envID).CreateApplicationRequest(createApplicationRequest).Execute()

CREATE Application



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
    createApplicationRequest := openapiclient.createApplication_request{ApplicationOIDC: openapiclient.NewApplicationOIDC(false, "Name_example", "Protocol_example", "Type_example", []string{"GrantTypes_example"}, "TokenEndpointAuthMethod_example")} // CreateApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationsApi.CreateApplication(context.Background(), envID).CreateApplicationRequest(createApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationsApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: CreateApplication201Response
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationsApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApplicationRequest** | [**CreateApplicationRequest**](CreateApplicationRequest.md) |  | 

### Return type

[**CreateApplication201Response**](CreateApplication201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, envID, appID).Execute()

DELETE Application



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
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationsApi.DeleteApplication(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationsApi.DeleteApplication``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## ReadAllApplications

> EntityArray ReadAllApplications(ctx, envID).Execute()

READ All Applications



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationsApi.ReadAllApplications(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationsApi.ReadAllApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplications`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationsApi.ReadAllApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationsRequest struct via the builder pattern


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


## ReadOneApplication

> CreateApplication201Response ReadOneApplication(ctx, envID, appID).Execute()

READ One Application



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
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationsApi.ReadOneApplication(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationsApi.ReadOneApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplication`: CreateApplication201Response
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationsApi.ReadOneApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateApplication201Response**](CreateApplication201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> CreateApplication201Response UpdateApplication(ctx, envID, appID).UpdateApplicationRequest(updateApplicationRequest).Execute()

UPDATE Application



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
    updateApplicationRequest := openapiclient.updateApplication_request{ApplicationOIDC: openapiclient.NewApplicationOIDC(false, "Name_example", "Protocol_example", "Type_example", []string{"GrantTypes_example"}, "TokenEndpointAuthMethod_example")} // UpdateApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsApplicationsApplicationsApi.UpdateApplication(context.Background(), envID, appID).UpdateApplicationRequest(updateApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationsApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: CreateApplication201Response
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsApplicationsApplicationsApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateApplicationRequest** | [**UpdateApplicationRequest**](UpdateApplicationRequest.md) |  | 

### Return type

[**CreateApplication201Response**](CreateApplication201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

