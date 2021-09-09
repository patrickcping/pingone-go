# \ManagementAPIsRiskManagementRiskEvaluationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDRiskEvaluationsPost**](ManagementAPIsRiskManagementRiskEvaluationsApi.md#V1EnvironmentsEnvIDRiskEvaluationsPost) | **Post** /v1/environments/{envID}/riskEvaluations | CREATE Risk Evaluation
[**V1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPut**](ManagementAPIsRiskManagementRiskEvaluationsApi.md#V1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPut) | **Put** /v1/environments/{envID}/riskEvaluations/{riskID}/event | UPDATE Risk Evaluation
[**V1EnvironmentsEnvIDRiskEvaluationsRiskIDGet**](ManagementAPIsRiskManagementRiskEvaluationsApi.md#V1EnvironmentsEnvIDRiskEvaluationsRiskIDGet) | **Get** /v1/environments/{envID}/riskEvaluations/{riskID} | READ One Risk Evaluation



## V1EnvironmentsEnvIDRiskEvaluationsPost

> V1EnvironmentsEnvIDRiskEvaluationsPost(ctx, envID).ContentType(contentType).Body(body).Execute()

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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskEvaluationsApi.V1EnvironmentsEnvIDRiskEvaluationsPost(context.Background(), envID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskEvaluationsApi.V1EnvironmentsEnvIDRiskEvaluationsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskEvaluationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPut

> V1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPut(ctx, envID, riskID).ContentType(contentType).Body(body).Execute()

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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskEvaluationsApi.V1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPut(context.Background(), envID, riskID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskEvaluationsApi.V1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskEvaluationsRiskIDEventPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDRiskEvaluationsRiskIDGet

> V1EnvironmentsEnvIDRiskEvaluationsRiskIDGet(ctx, envID, riskID).Execute()

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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskEvaluationsApi.V1EnvironmentsEnvIDRiskEvaluationsRiskIDGet(context.Background(), envID, riskID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskEvaluationsApi.V1EnvironmentsEnvIDRiskEvaluationsRiskIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDRiskEvaluationsRiskIDGetRequest struct via the builder pattern


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

