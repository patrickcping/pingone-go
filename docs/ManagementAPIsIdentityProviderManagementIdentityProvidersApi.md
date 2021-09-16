# \ManagementAPIsIdentityProviderManagementIdentityProvidersApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDIdentityProvidersGet**](ManagementAPIsIdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvIDIdentityProvidersGet) | **Get** /v1/environments/{envID}/identityProviders | READ All Identity Providers
[**V1EnvironmentsEnvIDIdentityProvidersPost**](ManagementAPIsIdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvIDIdentityProvidersPost) | **Post** /v1/environments/{envID}/identityProviders | Discover OpenID Provider Metadata
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDDelete**](ManagementAPIsIdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDDelete) | **Delete** /v1/environments/{envID}/identityProviders/{providerID} | DELETE Identity Provider
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDGet**](ManagementAPIsIdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDGet) | **Get** /v1/environments/{envID}/identityProviders/{providerID} | READ One Identity Provider
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDPut**](ManagementAPIsIdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDPut) | **Put** /v1/environments/{envID}/identityProviders/{providerID} | UPDATE Identity Provider



## V1EnvironmentsEnvIDIdentityProvidersGet

> V1EnvironmentsEnvIDIdentityProvidersGet(ctx, envID).Execute()

READ All Identity Providers



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
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersPost

> V1EnvironmentsEnvIDIdentityProvidersPost(ctx, envID).ContentType(contentType).Body(body).Execute()

Discover OpenID Provider Metadata



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
    contentType := "application/vnd.pingidentity.openid-configuration.discover+json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersPost(context.Background(), envID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDDelete

> V1EnvironmentsEnvIDIdentityProvidersProviderIDDelete(ctx, envID, providerID).Execute()

DELETE Identity Provider



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
    providerID := "providerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDDelete(context.Background(), envID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDGet

> V1EnvironmentsEnvIDIdentityProvidersProviderIDGet(ctx, envID, providerID).Execute()

READ One Identity Provider



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
    providerID := "providerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDGet(context.Background(), envID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDPut

> V1EnvironmentsEnvIDIdentityProvidersProviderIDPut(ctx, envID, providerID).Body(body).Execute()

UPDATE Identity Provider



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
    providerID := "providerID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDPut(context.Background(), envID, providerID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

