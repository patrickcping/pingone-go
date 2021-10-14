# \ManagementAPIsRiskManagementRiskAdvancedPredictorsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskPredictor**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#CreateRiskPredictor) | **Post** /v1/environments/{envID}/riskPredictors | CREATE Risk Predictor
[**DeleteRiskAdvancedPredictor**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#DeleteRiskAdvancedPredictor) | **Delete** /v1/environments/{envID}/riskPredictors/{riskPredictorID} | DELETE Risk Advanced Predictor
[**ReadAllRiskPredictors**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#ReadAllRiskPredictors) | **Get** /v1/environments/{envID}/riskPredictors | READ All Risk Predictors
[**ReadOneRiskPredictor**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#ReadOneRiskPredictor) | **Get** /v1/environments/{envID}/riskPredictors/{riskPredictorID} | READ One Risk Predictor
[**UpdateRiskPredictor**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#UpdateRiskPredictor) | **Put** /v1/environments/{envID}/riskPredictors/{riskPredictorID} | UPDATE Risk Predictor



## CreateRiskPredictor

> RiskPredictor CreateRiskPredictor(ctx, envID).RiskPredictor(riskPredictor).Execute()

CREATE Risk Predictor



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
    riskPredictor := *openapiclient.NewRiskPredictor("Name_example", "CompactName_example", "Type_example") // RiskPredictor |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.CreateRiskPredictor(context.Background(), envID).RiskPredictor(riskPredictor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.CreateRiskPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskPredictor`: RiskPredictor
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.CreateRiskPredictor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskPredictorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskPredictor** | [**RiskPredictor**](RiskPredictor.md) |  | 

### Return type

[**RiskPredictor**](RiskPredictor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskAdvancedPredictor

> DeleteRiskAdvancedPredictor(ctx, envID, riskPredictorID).Execute()

DELETE Risk Advanced Predictor



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
    riskPredictorID := "riskPredictorID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.DeleteRiskAdvancedPredictor(context.Background(), envID, riskPredictorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.DeleteRiskAdvancedPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPredictorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskAdvancedPredictorRequest struct via the builder pattern


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


## ReadAllRiskPredictors

> EntityArray ReadAllRiskPredictors(ctx, envID).Execute()

READ All Risk Predictors



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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.ReadAllRiskPredictors(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.ReadAllRiskPredictors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllRiskPredictors`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.ReadAllRiskPredictors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllRiskPredictorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneRiskPredictor

> RiskPredictor ReadOneRiskPredictor(ctx, envID, riskPredictorID).Execute()

READ One Risk Predictor



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
    riskPredictorID := "riskPredictorID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.ReadOneRiskPredictor(context.Background(), envID, riskPredictorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.ReadOneRiskPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneRiskPredictor`: RiskPredictor
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.ReadOneRiskPredictor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPredictorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRiskPredictorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskPredictor**](RiskPredictor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskPredictor

> RiskPredictor UpdateRiskPredictor(ctx, envID, riskPredictorID).RiskPredictor(riskPredictor).Execute()

UPDATE Risk Predictor



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
    riskPredictorID := "riskPredictorID_example" // string | 
    riskPredictor := *openapiclient.NewRiskPredictor("Name_example", "CompactName_example", "Type_example") // RiskPredictor |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.UpdateRiskPredictor(context.Background(), envID, riskPredictorID).RiskPredictor(riskPredictor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.UpdateRiskPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskPredictor`: RiskPredictor
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.UpdateRiskPredictor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPredictorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskPredictorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **riskPredictor** | [**RiskPredictor**](RiskPredictor.md) |  | 

### Return type

[**RiskPredictor**](RiskPredictor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

