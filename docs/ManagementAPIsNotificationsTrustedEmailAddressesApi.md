# \ManagementAPIsNotificationsTrustedEmailAddressesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGet**](ManagementAPIsNotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGet) | **Get** /v1/environments/{envID}/emailDomains/{emailDomainID}/trustedEmails | READ All Trusted Email Addresses
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPost**](ManagementAPIsNotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPost) | **Post** /v1/environments/{envID}/emailDomains/{emailDomainId}/trustedEmails | CREATE Trusted Email Address
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDelete**](ManagementAPIsNotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDelete) | **Delete** /v1/environments/{envID}/emailDomains/{emailDomainId}/trustedEmails/{trustedEmailId} | DELETE Trusted Email Address
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGet**](ManagementAPIsNotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGet) | **Get** /v1/environments/{envID}/emailDomains/{emailDomainId}/trustedEmails/{trustedEmailId} | READ One Trusted Email Address
[**V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPost**](ManagementAPIsNotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPost) | **Post** /v1/environments/{envID}/emailDomains/{emailDomainId}/trustedEmails/{trustedEmailId} | Resend Verification Code To Email



## V1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGet

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGet(ctx, envID, emailDomainID).Authorization(authorization).Execute()

READ All Trusted Email Addresses



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
    emailDomainID := "emailDomainID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGet(context.Background(), envID, emailDomainID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIDTrustedEmailsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPost

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPost(ctx, envID, emailDomainId).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Trusted Email Address



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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPost(context.Background(), envID, emailDomainId).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDelete

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDelete(ctx, envID, emailDomainId, trustedEmailId).Authorization(authorization).Execute()

DELETE Trusted Email Address



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
    trustedEmailId := "trustedEmailId_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDelete(context.Background(), envID, emailDomainId, trustedEmailId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDelete``: %v\n", err)
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
**trustedEmailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGet

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGet(ctx, envID, emailDomainId, trustedEmailId).Authorization(authorization).Execute()

READ One Trusted Email Address



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
    trustedEmailId := "trustedEmailId_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGet(context.Background(), envID, emailDomainId, trustedEmailId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGet``: %v\n", err)
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
**trustedEmailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPost

> V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPost(ctx, envID, emailDomainId, trustedEmailId).ContentType(contentType).Authorization(authorization).Execute()

Resend Verification Code To Email



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
    trustedEmailId := "trustedEmailId_example" // string | 
    contentType := "application/vnd.pingidentity.trustedEmail.sendVerificationCode+json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPost(context.Background(), envID, emailDomainId, trustedEmailId).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPost``: %v\n", err)
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
**trustedEmailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEmailDomainsEmailDomainIdTrustedEmailsTrustedEmailIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
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

