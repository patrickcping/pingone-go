# \ManagementAPIsRiskManagementRiskAdvancedPredictorsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDRiskPredictorsGet**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#V1EnvironmentsEnvIDRiskPredictorsGet) | **Get** /v1/environments/{envID}/riskPredictors | READ All Risk Predictors
[**V1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGet**](ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.md#V1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGet) | **Get** /v1/environments/{envID}/riskPredictors/{riskPredictorID} | READ One Risk Predictor



## V1EnvironmentsEnvIDRiskPredictorsGet

> V1EnvironmentsEnvIDRiskPredictorsGet(ctx, envID).Execute()

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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.V1EnvironmentsEnvIDRiskPredictorsGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.V1EnvironmentsEnvIDRiskPredictorsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPredictorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGet

> V1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGet(ctx, envID, riskPredictorID).Execute()

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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.V1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGet(context.Background(), envID, riskPredictorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskAdvancedPredictorsApi.V1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskPredictorsRiskPredictorIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

