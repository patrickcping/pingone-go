# \ManagementAPIsCustomDomainsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDCustomDomainsDomIDDelete**](ManagementAPIsCustomDomainsApi.md#V1EnvironmentsEnvIDCustomDomainsDomIDDelete) | **Delete** /v1/environments/{envID}/customDomains/{domID} | DELETE Domain
[**V1EnvironmentsEnvIDCustomDomainsDomIDGet**](ManagementAPIsCustomDomainsApi.md#V1EnvironmentsEnvIDCustomDomainsDomIDGet) | **Get** /v1/environments/{envID}/customDomains/{domID} | READ One Domain
[**V1EnvironmentsEnvIDCustomDomainsDomIDPost**](ManagementAPIsCustomDomainsApi.md#V1EnvironmentsEnvIDCustomDomainsDomIDPost) | **Post** /v1/environments/{envID}/customDomains/{domID} | Import Certificate
[**V1EnvironmentsEnvIDCustomDomainsGet**](ManagementAPIsCustomDomainsApi.md#V1EnvironmentsEnvIDCustomDomainsGet) | **Get** /v1/environments/{envID}/customDomains | READ All Domains
[**V1EnvironmentsEnvIDCustomDomainsPost**](ManagementAPIsCustomDomainsApi.md#V1EnvironmentsEnvIDCustomDomainsPost) | **Post** /v1/environments/{envID}/customDomains | CREATE Domain



## V1EnvironmentsEnvIDCustomDomainsDomIDDelete

> V1EnvironmentsEnvIDCustomDomainsDomIDDelete(ctx, envID, domID).Authorization(authorization).Execute()

DELETE Domain



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
    domID := "domID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsDomIDDelete(context.Background(), envID, domID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsDomIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**domID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCustomDomainsDomIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCustomDomainsDomIDGet

> V1EnvironmentsEnvIDCustomDomainsDomIDGet(ctx, envID, domID).Authorization(authorization).Execute()

READ One Domain



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
    domID := "domID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsDomIDGet(context.Background(), envID, domID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsDomIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**domID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCustomDomainsDomIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCustomDomainsDomIDPost

> V1EnvironmentsEnvIDCustomDomainsDomIDPost(ctx, envID, domID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

Import Certificate



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
    domID := "domID_example" // string | 
    contentType := "application/vnd.pingidentity.certificate.import+json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsDomIDPost(context.Background(), envID, domID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsDomIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**domID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCustomDomainsDomIDPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCustomDomainsGet

> V1EnvironmentsEnvIDCustomDomainsGet(ctx, envID).Authorization(authorization).Execute()

READ All Domains



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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCustomDomainsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCustomDomainsPost

> V1EnvironmentsEnvIDCustomDomainsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Domain



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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCustomDomainsApi.V1EnvironmentsEnvIDCustomDomainsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCustomDomainsPostRequest struct via the builder pattern


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

