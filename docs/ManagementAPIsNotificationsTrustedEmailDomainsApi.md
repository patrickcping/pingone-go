# \ManagementAPIsNotificationsTrustedEmailDomainsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete) | **Delete** /v1/environments/{envID}/emailDomains/{emailDomainId} | DELETE Trusted Email Domain
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet) | **Get** /v1/environments/{envID}/emailDomains/{emailDomainId}/dkim | READ Trusted Email Domain DKIM Status
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet) | **Get** /v1/environments/{envID}/emailDomains/{emailDomainId} | READ One Trusted Email Domain
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet) | **Get** /v1/environments/{envID}/emailDomains/{emailDomainId}/ownership | READ Trusted Email Domain Ownership Status
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet) | **Get** /v1/environments/{envID}/emailDomains/{emailDomainId}/spf | READ Trusted Email Domain SPF Status
[**V1EnvironmentsEnvIDEmailDomainsGet**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsGet) | **Get** /v1/environments/{envID}/emailDomains | READ All Trusted Email Domains
[**V1EnvironmentsEnvIDEmailDomainsPost**](ManagementAPIsNotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvIDEmailDomainsPost) | **Post** /v1/environments/{envID}/emailDomains | CREATE Trusted Email Domain



## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete(ctx, envID, emailDomainId).Execute()

DELETE Trusted Email Domain



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
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete(context.Background(), envID, emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet(ctx, envID, emailDomainId).Execute()

READ Trusted Email Domain DKIM Status



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
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet(context.Background(), envID, emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdDkimGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet(ctx, envID, emailDomainId).Execute()

READ One Trusted Email Domain



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
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet(context.Background(), envID, emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet(ctx, envID, emailDomainId).Execute()

READ Trusted Email Domain Ownership Status



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
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet(context.Background(), envID, emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdOwnershipGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet(ctx, envID, emailDomainId).Execute()

READ Trusted Email Domain SPF Status



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
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet(context.Background(), envID, emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdSpfGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsGet

> V1EnvironmentsEnvIDEmailDomainsGet(ctx, envID).Execute()

READ All Trusted Email Domains



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
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsPost

> V1EnvironmentsEnvIDEmailDomainsPost(ctx, envID).Body(body).Execute()

CREATE Trusted Email Domain



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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvIDEmailDomainsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsPostRequest struct via the builder pattern


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

