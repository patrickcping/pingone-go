# \ManagementAPIsRiskManagementRiskPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDRiskPolicySetsGet**](ManagementAPIsRiskManagementRiskPoliciesApi.md#V1EnvironmentsEnvIDRiskPolicySetsGet) | **Get** /v1/environments/{envID}/riskPolicySets | READ Risk Policy Sets
[**V1EnvironmentsEnvIDRiskPolicySetsPost**](ManagementAPIsRiskManagementRiskPoliciesApi.md#V1EnvironmentsEnvIDRiskPolicySetsPost) | **Post** /v1/environments/{envID}/riskPolicySets | CREATE Risk Policy Set
[**V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete**](ManagementAPIsRiskManagementRiskPoliciesApi.md#V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete) | **Delete** /v1/environments/{envID}/riskPolicySets/{riskPolicySetID} | DELETE Risk Policy Set 
[**V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet**](ManagementAPIsRiskManagementRiskPoliciesApi.md#V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet) | **Get** /v1/environments/{envID}/riskPolicySets/{riskPolicySetID} | READ One Risk Policy Set
[**V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut**](ManagementAPIsRiskManagementRiskPoliciesApi.md#V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut) | **Put** /v1/environments/{envID}/riskPolicySets/{riskPolicySetID} | UPDATE Risk Policy Set



## V1EnvironmentsEnvIDRiskPolicySetsGet

> V1EnvironmentsEnvIDRiskPolicySetsGet(ctx, envID).Execute()

READ Risk Policy Sets



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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPolicySetsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDRiskPolicySetsPost

> V1EnvironmentsEnvIDRiskPolicySetsPost(ctx, envID).Body(body).Execute()

CREATE Risk Policy Set



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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPolicySetsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete

> V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete(ctx, envID, riskPolicySetID).Execute()

DELETE Risk Policy Set 



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
    riskPolicySetID := "riskPolicySetID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete(context.Background(), envID, riskPolicySetID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet

> V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet(ctx, envID, riskPolicySetID).Execute()

READ One Risk Policy Set



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
    riskPolicySetID := "riskPolicySetID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet(context.Background(), envID, riskPolicySetID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut

> V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut(ctx, envID, riskPolicySetID).Body(body).Execute()

UPDATE Risk Policy Set



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
    riskPolicySetID := "riskPolicySetID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut(context.Background(), envID, riskPolicySetID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.V1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPolicySetsRiskPolicySetIDPutRequest struct via the builder pattern


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

