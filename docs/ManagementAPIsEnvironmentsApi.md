# \ManagementAPIsEnvironmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentActiveLicense**](ManagementAPIsEnvironmentsApi.md#CreateEnvironmentActiveLicense) | **Post** /v1/environments | CREATE Environment (Active License)
[**DeleteEnvironment**](ManagementAPIsEnvironmentsApi.md#DeleteEnvironment) | **Delete** /v1/environments/{envID} | DELETE Environment
[**ReadAllEnvironments**](ManagementAPIsEnvironmentsApi.md#ReadAllEnvironments) | **Get** /v1/environments | READ All Environments
[**ReadOneEnvironment**](ManagementAPIsEnvironmentsApi.md#ReadOneEnvironment) | **Get** /v1/environments/{envID} | READ One Environment
[**UpdateEnvironment**](ManagementAPIsEnvironmentsApi.md#UpdateEnvironment) | **Put** /v1/environments/{envID} | UPDATE Environment
[**UpdateEnvironmentType**](ManagementAPIsEnvironmentsApi.md#UpdateEnvironmentType) | **Put** /v1/environments/{envID}/type | UPDATE Environment Type



## CreateEnvironmentActiveLicense

> Environment CreateEnvironmentActiveLicense(ctx).ContentType(contentType).Environment(environment).Execute()

CREATE Environment (Active License)



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
    contentType := "application/json" // string |  (optional)
    environment := *openapiclient.NewEnvironment() // Environment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsEnvironmentsApi.CreateEnvironmentActiveLicense(context.Background()).ContentType(contentType).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsEnvironmentsApi.CreateEnvironmentActiveLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentActiveLicense`: Environment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsEnvironmentsApi.CreateEnvironmentActiveLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentActiveLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **environment** | [**Environment**](Environment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, envID).Execute()

DELETE Environment



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsEnvironmentsApi.DeleteEnvironment(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsEnvironmentsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## ReadAllEnvironments

> EntityArray ReadAllEnvironments(ctx).Limit(limit).Filter(filter).Execute()

READ All Environments



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
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)
    filter := "name sw "S" and license.id eq "34f0efac-21d9-4a17-8a35-196bb3362983"" // string | Adding a SCIM filter for an environment to display only those resources associated with the specified environment. 'sw', 'eq' and 'and' are supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsEnvironmentsApi.ReadAllEnvironments(context.Background()).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsEnvironmentsApi.ReadAllEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllEnvironments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsEnvironmentsApi.ReadAllEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadAllEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **filter** | **string** | Adding a SCIM filter for an environment to display only those resources associated with the specified environment. &#39;sw&#39;, &#39;eq&#39; and &#39;and&#39; are supported | 

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


## ReadOneEnvironment

> Environment ReadOneEnvironment(ctx, envID).Execute()

READ One Environment



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsEnvironmentsApi.ReadOneEnvironment(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsEnvironmentsApi.ReadOneEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsEnvironmentsApi.ReadOneEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> UpdateEnvironment(ctx, envID).ContentType(contentType).Environment(environment).Execute()

UPDATE Environment



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
    contentType := "application/json" // string |  (optional)
    environment := *openapiclient.NewEnvironment() // Environment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsEnvironmentsApi.UpdateEnvironment(context.Background(), envID).ContentType(contentType).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsEnvironmentsApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **environment** | [**Environment**](Environment.md) |  | 

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


## UpdateEnvironmentType

> UpdateEnvironmentType(ctx, envID).ContentType(contentType).InlineObject2(inlineObject2).Execute()

UPDATE Environment Type



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
    contentType := "application/json" // string |  (optional)
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsEnvironmentsApi.UpdateEnvironmentType(context.Background(), envID).ContentType(contentType).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsEnvironmentsApi.UpdateEnvironmentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

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

