# \SignOnPoliciesSignOnPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDSignOnPoliciesGet**](SignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesGet) | **Get** /v1/environments/{environmentID}/signOnPolicies | READ All Sign On Policies
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDelete**](SignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDelete) | **Delete** /v1/environments/{environmentID}/signOnPolicies/{policyID} | DELETE Sign On Policy
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGet**](SignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGet) | **Get** /v1/environments/{environmentID}/signOnPolicies/{policyID} | READ One Sign On Policy
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPut**](SignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPut) | **Put** /v1/environments/{environmentID}/signOnPolicies/{policyID} | UPDATE Sign On Policy
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPost**](SignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPost) | **Post** /v1/environments/{environmentID}/signOnPolicies | CREATE Sign On Policy



## V1EnvironmentsEnvironmentIDSignOnPoliciesGet

> V1EnvironmentsEnvironmentIDSignOnPoliciesGet(ctx, environmentID).Execute()

READ All Sign On Policies

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
    resp, r, err := apiClient.SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDelete

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDelete(ctx, environmentID, policyID).Execute()

DELETE Sign On Policy

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
    policyID := "policyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDelete(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGet

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGet(ctx, environmentID, policyID).Execute()

READ One Sign On Policy

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
    policyID := "policyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGet(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPut

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPut(ctx, environmentID, policyID).Body(body).Execute()

UPDATE Sign On Policy

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
    policyID := "policyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPut(context.Background(), environmentID, policyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPost

> V1EnvironmentsEnvironmentIDSignOnPoliciesPost(ctx, environmentID).Body(body).Execute()

CREATE Sign On Policy

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPostRequest struct via the builder pattern


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

