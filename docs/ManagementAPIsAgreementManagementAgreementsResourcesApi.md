# \ManagementAPIsAgreementManagementAgreementsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDAgreementsAgreementIDDelete**](ManagementAPIsAgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDDelete) | **Delete** /v1/environments/{envID}/agreements/{agreementID} | DELETE Agreement
[**V1EnvironmentsEnvIDAgreementsAgreementIDGet**](ManagementAPIsAgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDGet) | **Get** /v1/environments/{envID}/agreements/{agreementID} | READ One Agreement
[**V1EnvironmentsEnvIDAgreementsAgreementIDPut**](ManagementAPIsAgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDPut) | **Put** /v1/environments/{envID}/agreements/{agreementID} | UPDATE Agreement
[**V1EnvironmentsEnvIDAgreementsGet**](ManagementAPIsAgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvIDAgreementsGet) | **Get** /v1/environments/{envID}/agreements | READ All Agreements
[**V1EnvironmentsEnvIDAgreementsPost**](ManagementAPIsAgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvIDAgreementsPost) | **Post** /v1/environments/{envID}/agreements | CREATE Agreement



## V1EnvironmentsEnvIDAgreementsAgreementIDDelete

> V1EnvironmentsEnvIDAgreementsAgreementIDDelete(ctx, envID, agreementID).Authorization(authorization).Execute()

DELETE Agreement



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
    agreementID := "agreementID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDDelete(context.Background(), envID, agreementID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDGet

> V1EnvironmentsEnvIDAgreementsAgreementIDGet(ctx, envID, agreementID).Authorization(authorization).Execute()

READ One Agreement



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
    agreementID := "agreementID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDGet(context.Background(), envID, agreementID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDPut

> V1EnvironmentsEnvIDAgreementsAgreementIDPut(ctx, envID, agreementID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Agreement



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
    agreementID := "agreementID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDPut(context.Background(), envID, agreementID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsGet

> V1EnvironmentsEnvIDAgreementsGet(ctx, envID).Authorization(authorization).Execute()

READ All Agreements



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
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsPost

> V1EnvironmentsEnvIDAgreementsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Agreement



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
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvIDAgreementsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsPostRequest struct via the builder pattern


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

