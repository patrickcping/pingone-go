# \AuthenticationAPIsExternalAuthenticationApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvIDRpAuthenticateGet**](AuthenticationAPIsExternalAuthenticationApi.md#V1EnvIDRpAuthenticateGet) | **Get** /v1/{envID}/rp/authenticate | READ External Authentication Initialization
[**V1EnvIDRpCallbackProviderTypeGet**](AuthenticationAPIsExternalAuthenticationApi.md#V1EnvIDRpCallbackProviderTypeGet) | **Get** /v1/{envID}/rp/callback/{providerType} | READ External Authentication Callback
[**V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet**](AuthenticationAPIsExternalAuthenticationApi.md#V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet) | **Get** /v1/environments/{envID}/externalAuthentications/{extAuthID} | READ External Authentication Status



## V1EnvIDRpAuthenticateGet

> V1EnvIDRpAuthenticateGet(ctx, envID).ProviderId(providerId).FlowId(flowId).Execute()

READ External Authentication Initialization



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
    providerId := "{{idpID}}" // string |  (optional)
    flowId := "{{flowID}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsExternalAuthenticationApi.V1EnvIDRpAuthenticateGet(context.Background(), envID).ProviderId(providerId).FlowId(flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsExternalAuthenticationApi.V1EnvIDRpAuthenticateGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDRpAuthenticateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerId** | **string** |  | 
 **flowId** | **string** |  | 

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


## V1EnvIDRpCallbackProviderTypeGet

> V1EnvIDRpCallbackProviderTypeGet(ctx, envID, providerType).Code(code).State(state).Execute()

READ External Authentication Callback



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
    providerType := "providerType_example" // string | 
    code := "{{externalAuthCode}}" // string |  (optional)
    state := "{{externalProviderState}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsExternalAuthenticationApi.V1EnvIDRpCallbackProviderTypeGet(context.Background(), envID, providerType).Code(code).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsExternalAuthenticationApi.V1EnvIDRpCallbackProviderTypeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvIDRpCallbackProviderTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **code** | **string** |  | 
 **state** | **string** |  | 

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


## V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet

> V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet(ctx, envID, extAuthID).Cookie(cookie).Execute()

READ External Authentication Status



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
    extAuthID := "extAuthID_example" // string | 
    cookie := "ST={{sessionToken}}" // string | Use this for localhost (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsExternalAuthenticationApi.V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet(context.Background(), envID, extAuthID).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsExternalAuthenticationApi.V1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**extAuthID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDExternalAuthenticationsExtAuthIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cookie** | **string** | Use this for localhost | 

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

