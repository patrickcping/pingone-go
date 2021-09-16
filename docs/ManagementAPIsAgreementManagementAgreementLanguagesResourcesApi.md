# \ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet**](ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet) | **Get** /v1/environments/{envID}/agreements/{agreementID}/languages | READ All Languages
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete**](ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete) | **Delete** /v1/environments/{envID}/agreements/{agreementID}/languages/{languageID} | DELETE Language
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet**](ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet) | **Get** /v1/environments/{envID}/agreements/{agreementID}/languages/{languageID} | READ One Language
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut**](ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut) | **Put** /v1/environments/{envID}/agreements/{agreementID}/languages/{languageID} | UPDATE Language
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost**](ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost) | **Post** /v1/environments/{envID}/agreements/{agreementID}/languages | CREATE Language



## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet(ctx, envID, agreementID).Execute()

READ All Languages



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet(context.Background(), envID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete(ctx, envID, agreementID, languageID).Execute()

DELETE Language



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
    languageID := "languageID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete(context.Background(), envID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDelete``: %v\n", err)
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
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet(ctx, envID, agreementID, languageID).Execute()

READ One Language



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
    languageID := "languageID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet(context.Background(), envID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGet``: %v\n", err)
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
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut(ctx, envID, agreementID, languageID).Body(body).Execute()

UPDATE Language



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
    languageID := "languageID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut(context.Background(), envID, agreementID, languageID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPut``: %v\n", err)
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
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost(ctx, envID, agreementID).Body(body).Execute()

CREATE Language



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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost(context.Background(), envID, agreementID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesPostRequest struct via the builder pattern


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

