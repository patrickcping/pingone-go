# \ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDelete**](ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDelete) | **Delete** /v1/environments/:envID/agreements/:agreeID/languages/:langID/revisions/:revisionID | DELETE Revision
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet**](ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet) | **Get** /v1/environments/{envID}/agreements/{agreementID}/languages/{languageID}/revisions | READ All Revisions
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost**](ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost) | **Post** /v1/environments/{envID}/agreements/{agreementID}/languages/{languageID}/revisions | CREATE Revision
[**V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet**](ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet) | **Get** /v1/environments/{envID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | READ One Revision



## V1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDelete

> V1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDelete(ctx).Execute()

DELETE Revision



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreeIDLanguagesLangIDRevisionsRevisionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet(ctx, envID, agreementID, languageID).Execute()

READ All Revisions



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
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet(context.Background(), envID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost(ctx, envID, agreementID, languageID).ContentType(contentType).Body(body).Execute()

CREATE Revision



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
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost(context.Background(), envID, agreementID, languageID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet

> V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet(ctx, envID, agreementID, languageID, revisionID).Execute()

READ One Revision



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
    revisionID := "revisionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet(context.Background(), envID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet``: %v\n", err)
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
**revisionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGetRequest struct via the builder pattern


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

