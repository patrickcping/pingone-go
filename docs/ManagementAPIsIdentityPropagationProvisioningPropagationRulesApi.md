# \ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDPropagationPlansPlanIDRulesGet**](ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.md#V1EnvironmentsEnvIDPropagationPlansPlanIDRulesGet) | **Get** /v1/environments/{envID}/propagation/plans/{planID}/rules | READ One Plan&#39;s Rules
[**V1EnvironmentsEnvIDPropagationRulesGet**](ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.md#V1EnvironmentsEnvIDPropagationRulesGet) | **Get** /v1/environments/{envID}/propagation/rules | READ All Rules
[**V1EnvironmentsEnvIDPropagationRulesPost**](ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.md#V1EnvironmentsEnvIDPropagationRulesPost) | **Post** /v1/environments/{envID}/propagation/rules | CREATE Rule
[**V1EnvironmentsEnvIDPropagationRulesRuleIDDelete**](ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.md#V1EnvironmentsEnvIDPropagationRulesRuleIDDelete) | **Delete** /v1/environments/{envID}/propagation/rules/{ruleID} | DELETE Rule
[**V1EnvironmentsEnvIDPropagationRulesRuleIDGet**](ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.md#V1EnvironmentsEnvIDPropagationRulesRuleIDGet) | **Get** /v1/environments/{envID}/propagation/rules/{ruleID} | READ One Rule
[**V1EnvironmentsEnvIDPropagationRulesStoreIDPut**](ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.md#V1EnvironmentsEnvIDPropagationRulesStoreIDPut) | **Put** /v1/environments/{envID}/propagation/rules/{storeID} | UPDATE Rule



## V1EnvironmentsEnvIDPropagationPlansPlanIDRulesGet

> V1EnvironmentsEnvIDPropagationPlansPlanIDRulesGet(ctx, envID, planID).Accept(accept).ContentType(contentType).Execute()

READ One Plan's Rules



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
    contentType := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationPlansPlanIDRulesGet(context.Background(), envID, planID).Accept(accept).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationPlansPlanIDRulesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationPlansPlanIDRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvIDPropagationRulesGet

> V1EnvironmentsEnvIDPropagationRulesGet(ctx, envID).Accept(accept).Authorization(authorization).Execute()

READ All Rules



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
    authorization := "authorization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesGet(context.Background(), envID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvIDPropagationRulesPost

> V1EnvironmentsEnvIDPropagationRulesPost(ctx, envID).Body(body).Execute()

CREATE Rule



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
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationRulesRuleIDDelete

> V1EnvironmentsEnvIDPropagationRulesRuleIDDelete(ctx, envID, ruleID).Accept(accept).Execute()

DELETE Rule



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
    ruleID := "ruleID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesRuleIDDelete(context.Background(), envID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesRuleIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**ruleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesRuleIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationRulesRuleIDGet

> V1EnvironmentsEnvIDPropagationRulesRuleIDGet(ctx, envID, ruleID).Accept(accept).Execute()

READ One Rule



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
    ruleID := "ruleID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesRuleIDGet(context.Background(), envID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesRuleIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**ruleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesRuleIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationRulesStoreIDPut

> V1EnvironmentsEnvIDPropagationRulesStoreIDPut(ctx, envID, storeID).Body(body).Execute()

UPDATE Rule



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
    storeID := "storeID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesStoreIDPut(context.Background(), envID, storeID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRulesApi.V1EnvironmentsEnvIDPropagationRulesStoreIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesStoreIDPutRequest struct via the builder pattern


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

