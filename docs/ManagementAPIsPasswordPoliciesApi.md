# \ManagementAPIsPasswordPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDPasswordPoliciesGet**](ManagementAPIsPasswordPoliciesApi.md#V1EnvironmentsEnvIDPasswordPoliciesGet) | **Get** /v1/environments/{envID}/passwordPolicies | READ All Password Policies
[**V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGet**](ManagementAPIsPasswordPoliciesApi.md#V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGet) | **Get** /v1/environments/{envID}/passwordPolicies/{passwordPolicyID} | READ One Password Policy
[**V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPut**](ManagementAPIsPasswordPoliciesApi.md#V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPut) | **Put** /v1/environments/{envID}/passwordPolicies/{passwordPolicyID} | UPDATE Password Policy



## V1EnvironmentsEnvIDPasswordPoliciesGet

> V1EnvironmentsEnvIDPasswordPoliciesGet(ctx, envID).Authorization(authorization).Execute()

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
    envID := "envID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPasswordPoliciesApi.V1EnvironmentsEnvIDPasswordPoliciesGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPasswordPoliciesApi.V1EnvironmentsEnvIDPasswordPoliciesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPasswordPoliciesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGet

> V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGet(ctx, envID, passwordPolicyID).Authorization(authorization).Execute()

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
    envID := "envID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPasswordPoliciesApi.V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGet(context.Background(), envID, passwordPolicyID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPasswordPoliciesApi.V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPut

> V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPut(ctx, envID, passwordPolicyID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

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
    envID := "envID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPasswordPoliciesApi.V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPut(context.Background(), envID, passwordPolicyID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPasswordPoliciesApi.V1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPasswordPoliciesPasswordPolicyIDPutRequest struct via the builder pattern


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

