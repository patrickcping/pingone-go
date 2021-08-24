# \ManagementAPIsSignOnPoliciesSignOnPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDSignOnPoliciesGet**](ManagementAPIsSignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvIDSignOnPoliciesGet) | **Get** /v1/environments/{envID}/signOnPolicies | READ All Sign On Policies
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDDelete**](ManagementAPIsSignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDDelete) | **Delete** /v1/environments/{envID}/signOnPolicies/{policyID} | DELETE Sign On Policy
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDGet**](ManagementAPIsSignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDGet) | **Get** /v1/environments/{envID}/signOnPolicies/{policyID} | READ One Sign On Policy
[**V1EnvironmentsEnvIDSignOnPoliciesPolicyIDPut**](ManagementAPIsSignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvIDSignOnPoliciesPolicyIDPut) | **Put** /v1/environments/{envID}/signOnPolicies/{policyID} | UPDATE Sign On Policy
[**V1EnvironmentsEnvIDSignOnPoliciesPost**](ManagementAPIsSignOnPoliciesSignOnPoliciesApi.md#V1EnvironmentsEnvIDSignOnPoliciesPost) | **Post** /v1/environments/{envID}/signOnPolicies | CREATE Sign On Policy



## V1EnvironmentsEnvIDSignOnPoliciesGet

> V1EnvironmentsEnvIDSignOnPoliciesGet(ctx, envID).Authorization(authorization).Execute()

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
    envID := "envID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDDelete

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDDelete(ctx, envID, policyID).Authorization(authorization).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDDelete(context.Background(), envID, policyID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDGet

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDGet(ctx, envID, policyID).Authorization(authorization).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDGet(context.Background(), envID, policyID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPolicyIDPut

> V1EnvironmentsEnvIDSignOnPoliciesPolicyIDPut(ctx, envID, policyID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

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
    envID := "envID_example" // string | 
    policyID := "policyID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDPut(context.Background(), envID, policyID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPolicyIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPolicyIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDSignOnPoliciesPost

> V1EnvironmentsEnvIDSignOnPoliciesPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

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
    envID := "envID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsSignOnPoliciesSignOnPoliciesApi.V1EnvironmentsEnvIDSignOnPoliciesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSignOnPoliciesPostRequest struct via the builder pattern


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

