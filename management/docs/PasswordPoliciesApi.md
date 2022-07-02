# \PasswordPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDPasswordPoliciesGet**](PasswordPoliciesApi.md#V1EnvironmentsEnvironmentIDPasswordPoliciesGet) | **Get** /v1/environments/{environmentID}/passwordPolicies | READ All Password Policies
[**V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGet**](PasswordPoliciesApi.md#V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGet) | **Get** /v1/environments/{environmentID}/passwordPolicies/{passwordPolicyID} | READ One Password Policy
[**V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPut**](PasswordPoliciesApi.md#V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPut) | **Put** /v1/environments/{environmentID}/passwordPolicies/{passwordPolicyID} | UPDATE Password Policy



## V1EnvironmentsEnvironmentIDPasswordPoliciesGet

> V1EnvironmentsEnvironmentIDPasswordPoliciesGet(ctx, environmentID).Execute()

READ All Password Policies

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
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.V1EnvironmentsEnvironmentIDPasswordPoliciesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.V1EnvironmentsEnvironmentIDPasswordPoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPasswordPoliciesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGet

> V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGet(ctx, environmentID, passwordPolicyID).Execute()

READ One Password Policy

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
    environmentID := "environmentID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGet(context.Background(), environmentID, passwordPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPut

> V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPut(ctx, environmentID, passwordPolicyID).Body(body).Execute()

UPDATE Password Policy

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
    environmentID := "environmentID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPut(context.Background(), environmentID, passwordPolicyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.V1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPasswordPoliciesPasswordPolicyIDPutRequest struct via the builder pattern


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

