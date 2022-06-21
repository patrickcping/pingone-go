# \ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDPropagationPlansGet**](ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.md#V1EnvironmentsEnvIDPropagationPlansGet) | **Get** /v1/environments/{envID}/propagation/plans | READ All Plans
[**V1EnvironmentsEnvIDPropagationPlansPlanIDDelete**](ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.md#V1EnvironmentsEnvIDPropagationPlansPlanIDDelete) | **Delete** /v1/environments/{envID}/propagation/plans/{planID} | DELETE Plan
[**V1EnvironmentsEnvIDPropagationPlansPlanIDGet**](ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.md#V1EnvironmentsEnvIDPropagationPlansPlanIDGet) | **Get** /v1/environments/{envID}/propagation/plans/{planID} | READ One Plan
[**V1EnvironmentsEnvIDPropagationPlansPlanIDPut**](ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.md#V1EnvironmentsEnvIDPropagationPlansPlanIDPut) | **Put** /v1/environments/{envID}/propagation/plans/{planID} | UPDATE Plan
[**V1EnvironmentsEnvIDPropagationPlansPost**](ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.md#V1EnvironmentsEnvIDPropagationPlansPost) | **Post** /v1/environments/{envID}/propagation/plans | CREATE Plan



## V1EnvironmentsEnvIDPropagationPlansGet

> V1EnvironmentsEnvIDPropagationPlansGet(ctx, envID).Accept(accept).Execute()

READ All Plans



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
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansGet(context.Background(), envID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationPlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 

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


## V1EnvironmentsEnvIDPropagationPlansPlanIDDelete

> V1EnvironmentsEnvIDPropagationPlansPlanIDDelete(ctx, envID, planID).Accept(accept).Execute()

DELETE Plan



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
    planID := "planID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPlanIDDelete(context.Background(), envID, planID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPlanIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**planID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationPlansPlanIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

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


## V1EnvironmentsEnvIDPropagationPlansPlanIDGet

> V1EnvironmentsEnvIDPropagationPlansPlanIDGet(ctx, envID, planID).Accept(accept).Execute()

READ One Plan



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
    planID := "planID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPlanIDGet(context.Background(), envID, planID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPlanIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**planID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationPlansPlanIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

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


## V1EnvironmentsEnvIDPropagationPlansPlanIDPut

> V1EnvironmentsEnvIDPropagationPlansPlanIDPut(ctx, envID, planID).Body(body).Execute()

UPDATE Plan



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
    planID := "planID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPlanIDPut(context.Background(), envID, planID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPlanIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**planID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationPlansPlanIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationPlansPost

> V1EnvironmentsEnvIDPropagationPlansPost(ctx, envID).Body(body).Execute()

CREATE Plan



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationPlansApi.V1EnvironmentsEnvIDPropagationPlansPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationPlansPostRequest struct via the builder pattern


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

