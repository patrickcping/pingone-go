# \ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost**](ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost) | **Post** /v1/environments/{envID}/propagation/storeMetadata/Aquera | Identity Propagation Store Metadata (Aquera)
[**V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost**](ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost) | **Post** /v1/environments/{envID}/propagation/storeMetadata/SalesforceContacts | Identity Propagation Store Metadata (SalesforceContacts)
[**V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost**](ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost) | **Post** /v1/environments/{envID}/propagation/storeMetadata/Salesforce | Identity Propagation Store Metadata (Salesforce)
[**V1EnvironmentsEnvIDPropagationStoreMetadataScimPost**](ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvIDPropagationStoreMetadataScimPost) | **Post** /v1/environments/{envID}/propagation/storeMetadata/scim | Identity Propagation Store Metadata (SCIM)



## V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost

> V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

Identity Propagation Store Metadata (Aquera)



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
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataAqueraPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationStoreMetadataAqueraPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost

> V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

Identity Propagation Store Metadata (SalesforceContacts)



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
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforceContactsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost

> V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

Identity Propagation Store Metadata (Salesforce)



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
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationStoreMetadataSalesforcePostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationStoreMetadataScimPost

> V1EnvironmentsEnvIDPropagationStoreMetadataScimPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

Identity Propagation Store Metadata (SCIM)



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
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataScimPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvIDPropagationStoreMetadataScimPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationStoreMetadataScimPostRequest struct via the builder pattern


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

