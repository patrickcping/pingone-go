# \SignOnPoliciesSignOnPolicyActionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDelete**](SignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDelete) | **Delete** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions/{actionID} | DELETE Sign-On Policy Action
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGet**](SignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGet) | **Get** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions/{actionID} | READ One Sign-On Policy Action
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPut**](SignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPut) | **Put** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions/{actionID} | UPDATE Sign-On Policy Action
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGet**](SignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGet) | **Get** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions | READ All Sign-On Policy Actions
[**V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPost**](SignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPost) | **Post** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions | CREATE Sign-On Policy Action (AGREEMENT)



## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDelete

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDelete(ctx, environmentID, policyID, actionID).Execute()

DELETE Sign-On Policy Action

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
    actionID := "actionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDelete(context.Background(), environmentID, policyID, actionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDelete``: %v\n", err)
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
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGet

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGet(ctx, environmentID, policyID, actionID).Execute()

READ One Sign-On Policy Action

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
    actionID := "actionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGet(context.Background(), environmentID, policyID, actionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGet``: %v\n", err)
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
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPut

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPut(ctx, environmentID, policyID, actionID).Body(body).Execute()

UPDATE Sign-On Policy Action

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
    actionID := "actionID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPut(context.Background(), environmentID, policyID, actionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPut``: %v\n", err)
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
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsActionIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGet

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGet(ctx, environmentID, policyID).Execute()

READ All Sign-On Policy Actions

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
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGet(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPost

> V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPost(ctx, environmentID, policyID).Body(body).Execute()

CREATE Sign-On Policy Action (AGREEMENT)

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
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPost(context.Background(), environmentID, policyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSignOnPoliciesPolicyIDActionsPostRequest struct via the builder pattern


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

