# \ManagementAPIsIntegrationCatalogApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDIntegrationsGet**](ManagementAPIsIntegrationCatalogApi.md#V1EnvironmentsEnvIDIntegrationsGet) | **Get** /v1/environments/{envID}/integrations | READ Integration Metadata
[**V1EnvironmentsEnvIDIntegrationsIntegrationIDGet**](ManagementAPIsIntegrationCatalogApi.md#V1EnvironmentsEnvIDIntegrationsIntegrationIDGet) | **Get** /v1/environments/{envID}/integrations/{integrationID} | READ One Integration Metadata
[**V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGet**](ManagementAPIsIntegrationCatalogApi.md#V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGet) | **Get** /v1/environments/{envID}/integrations/{integrationID}/versions | READ Integration Version Metadata
[**V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet**](ManagementAPIsIntegrationCatalogApi.md#V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet) | **Get** /v1/environments/{envID}/integrations/{integrationID}/versions/{integrationVersionID}/asset | READ Integration Version Asset Download
[**V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet**](ManagementAPIsIntegrationCatalogApi.md#V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet) | **Get** /v1/environments/{envID}/integrations/{integrationID}/versions/{integrationVersionID} | READ One Integration Version Metadata



## V1EnvironmentsEnvIDIntegrationsGet

> V1EnvironmentsEnvIDIntegrationsGet(ctx, envID).Execute()

READ Integration Metadata



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
    resp, r, err := api_client.ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIntegrationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1EnvironmentsEnvIDIntegrationsIntegrationIDGet

> V1EnvironmentsEnvIDIntegrationsIntegrationIDGet(ctx, envID, integrationID).Execute()

READ One Integration Metadata



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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDGet(context.Background(), envID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIntegrationsIntegrationIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGet

> V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGet(ctx, envID, integrationID).Execute()

READ Integration Version Metadata



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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGet(context.Background(), envID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet

> V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet(ctx, envID, integrationID, integrationVersionID).Execute()

READ Integration Version Asset Download



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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet(context.Background(), envID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet

> V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet(ctx, envID, integrationID, integrationVersionID).Execute()

READ One Integration Version Metadata



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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet(context.Background(), envID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIntegrationCatalogApi.V1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

