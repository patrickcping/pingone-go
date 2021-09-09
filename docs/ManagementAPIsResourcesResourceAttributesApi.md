# \ManagementAPIsResourcesResourceAttributesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDResourcesResourceIDAttributesGet**](ManagementAPIsResourcesResourceAttributesApi.md#V1EnvironmentsEnvIDResourcesResourceIDAttributesGet) | **Get** /v1/environments/{envID}/resources/{resourceID}/attributes | READ All Resource Attributes
[**V1EnvironmentsEnvIDResourcesResourceIDAttributesPost**](ManagementAPIsResourcesResourceAttributesApi.md#V1EnvironmentsEnvIDResourcesResourceIDAttributesPost) | **Post** /v1/environments/{envID}/resources/{resourceID}/attributes | CREATE Resource Attribute
[**V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete**](ManagementAPIsResourcesResourceAttributesApi.md#V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete) | **Delete** /v1/environments/{envID}/resources/{resourceID}/attributes/{resourceAttrID} | DELETE Resource Attribute
[**V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet**](ManagementAPIsResourcesResourceAttributesApi.md#V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet) | **Get** /v1/environments/{envID}/resources/{resourceID}/attributes/{resourceAttrID} | READ One Resource Attribute
[**V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut**](ManagementAPIsResourcesResourceAttributesApi.md#V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut) | **Put** /v1/environments/{envID}/resources/{resourceID}/attributes/{resourceAttrID} | UPDATE Resource Attribute



## V1EnvironmentsEnvIDResourcesResourceIDAttributesGet

> V1EnvironmentsEnvIDResourcesResourceIDAttributesGet(ctx, envID, resourceID).Execute()

READ All Resource Attributes



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesGet(context.Background(), envID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDAttributesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDResourcesResourceIDAttributesPost

> V1EnvironmentsEnvIDResourcesResourceIDAttributesPost(ctx, envID, resourceID).ContentType(contentType).Body(body).Execute()

CREATE Resource Attribute



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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesPost(context.Background(), envID, resourceID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDAttributesPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete

> V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete(ctx, envID, resourceID, resourceAttrID).Execute()

DELETE Resource Attribute



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
    resourceAttrID := "resourceAttrID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete(context.Background(), envID, resourceID, resourceAttrID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDelete``: %v\n", err)
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
**resourceAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet

> V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet(ctx, envID, resourceID, resourceAttrID).Execute()

READ One Resource Attribute



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
    resourceAttrID := "resourceAttrID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet(context.Background(), envID, resourceID, resourceAttrID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGet``: %v\n", err)
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
**resourceAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut

> V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut(ctx, envID, resourceID, resourceAttrID).ContentType(contentType).Body(body).Execute()

UPDATE Resource Attribute



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
    resourceAttrID := "resourceAttrID_example" // string | 
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut(context.Background(), envID, resourceID, resourceAttrID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsResourcesResourceAttributesApi.V1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPut``: %v\n", err)
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
**resourceAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDResourcesResourceIDAttributesResourceAttrIDPutRequest struct via the builder pattern


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

