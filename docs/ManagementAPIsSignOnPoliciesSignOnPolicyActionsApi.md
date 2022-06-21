# \ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPost**](ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.md#V1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPost) | **Post** /v1/environments/9ad15e9e-3ac6-43f7-a053-d46b87d6c4a7/signOnPolicies/{policyID}/actions | CREATE Sign-On Policy Action (AGREEMENT)
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDelete**](ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDelete) | **Delete** /v1/environments/{envID}/signOnPolicies/{policyID}/actions/{actionID} | DELETE Sign-On Policy Action
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGet**](ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGet) | **Get** /v1/environments/{envID}/signOnPolicies/{policyID}/actions/{actionID} | READ One Sign-On Policy Action
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPut**](ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPut) | **Put** /v1/environments/{envID}/signOnPolicies/{policyID}/actions/{actionID} | UPDATE Sign-On Policy Action
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGet**](ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGet) | **Get** /v1/environments/{envID}/signOnPolicies/{policyID}/actions | READ All Sign-On Policy Actions
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPost**](ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPost) | **Post** /v1/environments/{envID}/signOnPolicies/{policyID}/actions | CREATE Sign-On Policy Action (IDENTITY_PROVIDER)



## V1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPost

> V1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPost(ctx, policyID).Body(body).Execute()

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
    policyID := "policyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPost(context.Background(), policyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1Environments9ad15e9e3ac643f7A053D46b87d6c4a7SignOnPoliciesPolicyIDActionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDelete

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDelete(ctx, envID, policyID, actionID).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 
    actionID := "actionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDelete(context.Background(), envID, policyID, actionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**policyID** | **string** |  | 
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGet

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGet(ctx, envID, policyID, actionID).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 
    actionID := "actionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGet(context.Background(), envID, policyID, actionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**policyID** | **string** |  | 
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPut

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPut(ctx, envID, policyID, actionID).Body(body).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 
    actionID := "actionID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPut(context.Background(), envID, policyID, actionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**policyID** | **string** |  | 
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsActionIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGet

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGet(ctx, envID, policyID).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGet(context.Background(), envID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPost

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPost(ctx, envID, policyID).Body(body).Execute()

CREATE Sign-On Policy Action (IDENTITY_PROVIDER)



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
    policyID := "policyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPost(context.Background(), envID, policyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPolicyActionsApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDActionsPostRequest struct via the builder pattern


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

