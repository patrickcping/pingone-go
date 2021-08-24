# \ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDPropagationMappingMappingIDDelete**](ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvIDPropagationMappingMappingIDDelete) | **Delete** /v1/environments/{envID}/propagation/mapping/{mappingID} | DELETE Mapping
[**V1EnvironmentsEnvIDPropagationMappingsMappingIDGet**](ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvIDPropagationMappingsMappingIDGet) | **Get** /v1/environments/{envID}/propagation/mappings/{mappingID} | READ One Mapping
[**V1EnvironmentsEnvIDPropagationMappingsMappingIDPut**](ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvIDPropagationMappingsMappingIDPut) | **Put** /v1/environments/{envID}/propagation/mappings/{mappingID} | UPDATE Mapping
[**V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGet**](ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGet) | **Get** /v1/environments/{envID}/propagation/rules/{ruleID}/mappings | READ One Rule  Mapping
[**V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPost**](ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPost) | **Post** /v1/environments/{envID}/propagation/rules/{ruleID}/mappings | CREATE Rule Mapping



## V1EnvironmentsEnvIDPropagationMappingMappingIDDelete

> V1EnvironmentsEnvIDPropagationMappingMappingIDDelete(ctx, envID, mappingID).Accept(accept).Authorization(authorization).Execute()

DELETE Mapping



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
    mappingID := "mappingID_example" // string | 
    accept := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationMappingMappingIDDelete(context.Background(), envID, mappingID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationMappingMappingIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**mappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationMappingMappingIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
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


## V1EnvironmentsEnvIDPropagationMappingsMappingIDGet

> V1EnvironmentsEnvIDPropagationMappingsMappingIDGet(ctx, envID, mappingID).Accept(accept).Authorization(authorization).Execute()

READ One Mapping



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
    mappingID := "mappingID_example" // string | 
    accept := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationMappingsMappingIDGet(context.Background(), envID, mappingID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationMappingsMappingIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**mappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationMappingsMappingIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
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


## V1EnvironmentsEnvIDPropagationMappingsMappingIDPut

> V1EnvironmentsEnvIDPropagationMappingsMappingIDPut(ctx, envID, mappingID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Mapping



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
    mappingID := "mappingID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationMappingsMappingIDPut(context.Background(), envID, mappingID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationMappingsMappingIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**mappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationMappingsMappingIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGet

> V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGet(ctx, envID, ruleID).Accept(accept).ContentType(contentType).Authorization(authorization).Execute()

READ One Rule  Mapping



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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGet(context.Background(), envID, ruleID).Accept(accept).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesRuleIDMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPost

> V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPost(ctx, envID, ruleID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Rule Mapping



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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPost(context.Background(), envID, ruleID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRulesRuleIDMappingsPostRequest struct via the builder pattern


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

