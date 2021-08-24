# \ManagementAPIsApplicationsApplicationAttributeMappingApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDApplicationsAppIDAttributesGet**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#V1EnvironmentsEnvIDApplicationsAppIDAttributesGet) | **Get** /v1/environments/{envID}/applications/{appID}/attributes | READ All Application Attribute Mappings
[**V1EnvironmentsEnvIDApplicationsAppIDAttributesPost**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#V1EnvironmentsEnvIDApplicationsAppIDAttributesPost) | **Post** /v1/environments/{envID}/applications/{appID}/attributes | CREATE Application Attribute Mapping
[**V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDelete**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDelete) | **Delete** /v1/environments/{envID}/applications/{appID}/attributes/{samlAttrID} | DELETE Application Attribute Mapping
[**V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGet**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGet) | **Get** /v1/environments/{envID}/applications/{appID}/attributes/{samlAttrID} | READ One Application Attribute Mapping
[**V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPut**](ManagementAPIsApplicationsApplicationAttributeMappingApi.md#V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPut) | **Put** /v1/environments/{envID}/applications/{appID}/attributes/{samlAttrID} | UPDATE Application Attribute Mapping



## V1EnvironmentsEnvIDApplicationsAppIDAttributesGet

> V1EnvironmentsEnvIDApplicationsAppIDAttributesGet(ctx, envID, appID).Authorization(authorization).Execute()

READ All Application Attribute Mappings



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
    appID := "appID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesGet(context.Background(), envID, appID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDAttributesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDAttributesPost

> V1EnvironmentsEnvIDApplicationsAppIDAttributesPost(ctx, envID, appID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Application Attribute Mapping



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
    appID := "appID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesPost(context.Background(), envID, appID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDAttributesPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDelete

> V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDelete(ctx, envID, appID, samlAttrID).Authorization(authorization).Execute()

DELETE Application Attribute Mapping



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
    appID := "appID_example" // string | 
    samlAttrID := "samlAttrID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDelete(context.Background(), envID, appID, samlAttrID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**samlAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGet

> V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGet(ctx, envID, appID, samlAttrID).Authorization(authorization).Execute()

READ One Application Attribute Mapping



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
    appID := "appID_example" // string | 
    samlAttrID := "samlAttrID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGet(context.Background(), envID, appID, samlAttrID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**samlAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPut

> V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPut(ctx, envID, appID, samlAttrID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Application Attribute Mapping



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
    appID := "appID_example" // string | 
    samlAttrID := "samlAttrID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPut(context.Background(), envID, appID, samlAttrID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationAttributeMappingApi.V1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**samlAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDAttributesSamlAttrIDPutRequest struct via the builder pattern


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

