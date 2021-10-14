# \ManagementAPIsRiskManagementRiskEvaluationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskEvaluation**](ManagementAPIsRiskManagementRiskEvaluationsApi.md#CreateRiskEvaluation) | **Post** /v1/environments/{envID}/riskEvaluations | CREATE Risk Evaluation
[**ReadOneRiskEvaluation**](ManagementAPIsRiskManagementRiskEvaluationsApi.md#ReadOneRiskEvaluation) | **Get** /v1/environments/{envID}/riskEvaluations/{riskID} | READ One Risk Evaluation
[**UpdateRiskEvaluation**](ManagementAPIsRiskManagementRiskEvaluationsApi.md#UpdateRiskEvaluation) | **Put** /v1/environments/{envID}/riskEvaluations/{riskID}/event | UPDATE Risk Evaluation



## CreateRiskEvaluation

> RiskEvaluation CreateRiskEvaluation(ctx, envID).RiskEvaluation(riskEvaluation).Execute()

CREATE Risk Evaluation



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
    riskEvaluation := *openapiclient.NewRiskEvaluation(*openapiclient.NewRiskEvaluationEvent("Ip_example", *openapiclient.NewRiskEvaluationEventUser("Id_example", "Type_example"))) // RiskEvaluation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskEvaluationsApi.CreateRiskEvaluation(context.Background(), envID).RiskEvaluation(riskEvaluation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskEvaluationsApi.CreateRiskEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskEvaluation`: RiskEvaluation
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskEvaluationsApi.CreateRiskEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskEvaluation** | [**RiskEvaluation**](RiskEvaluation.md) |  | 

### Return type

[**RiskEvaluation**](RiskEvaluation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneRiskEvaluation

> RiskEvaluation ReadOneRiskEvaluation(ctx, envID, riskID).Execute()

READ One Risk Evaluation



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
    riskID := "riskID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskEvaluationsApi.ReadOneRiskEvaluation(context.Background(), envID, riskID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskEvaluationsApi.ReadOneRiskEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneRiskEvaluation`: RiskEvaluation
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskEvaluationsApi.ReadOneRiskEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRiskEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskEvaluation**](RiskEvaluation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskEvaluation

> RiskEvaluation UpdateRiskEvaluation(ctx, envID, riskID).RiskEvaluationEvent(riskEvaluationEvent).Execute()

UPDATE Risk Evaluation



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
    riskID := "riskID_example" // string | 
    riskEvaluationEvent := *openapiclient.NewRiskEvaluationEvent("Ip_example", *openapiclient.NewRiskEvaluationEventUser("Id_example", "Type_example")) // RiskEvaluationEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskEvaluationsApi.UpdateRiskEvaluation(context.Background(), envID, riskID).RiskEvaluationEvent(riskEvaluationEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskEvaluationsApi.UpdateRiskEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskEvaluation`: RiskEvaluation
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskEvaluationsApi.UpdateRiskEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **riskEvaluationEvent** | [**RiskEvaluationEvent**](RiskEvaluationEvent.md) |  | 

### Return type

[**RiskEvaluation**](RiskEvaluation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

