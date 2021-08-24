# \ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete**](ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.md#V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete) | **Delete** /v11/environments/{envID}/identityProviders/{providerID}/attributes/{idpAttrID} | DELETE Identity Provider Attribute
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet**](ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet) | **Get** /v1/environments/{envID}/identityProviders/{providerID}/attributes | READ All Identity Provider Attributes
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet**](ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet) | **Get** /v1/environments/{envID}/identityProviders/{providerID}/attributes/{idpAttrID} | READ One Identity Provider Attribute
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut**](ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut) | **Put** /v1/environments/{envID}/identityProviders/{providerID}/attributes/{idpAttrID} | UPDATE Identity Provider Attribute
[**V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost**](ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost) | **Post** /v1/environments/{envID}/identityProviders/{providerID}/attributes | CREATE Identity Provider Attribute (SAML)



## V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete

> V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete(ctx, envID, providerID, idpAttrID).Authorization(authorization).Execute()

DELETE Identity Provider Attribute



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
    providerID := "providerID_example" // string | 
    idpAttrID := "idpAttrID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete(context.Background(), envID, providerID, idpAttrID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 
**idpAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV11EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet

> V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet(ctx, envID, providerID).Authorization(authorization).Execute()

READ All Identity Provider Attributes



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
    providerID := "providerID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet(context.Background(), envID, providerID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet

> V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet(ctx, envID, providerID, idpAttrID).Authorization(authorization).Execute()

READ One Identity Provider Attribute



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
    providerID := "providerID_example" // string | 
    idpAttrID := "idpAttrID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet(context.Background(), envID, providerID, idpAttrID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 
**idpAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut

> V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut(ctx, envID, providerID, idpAttrID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Identity Provider Attribute



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
    providerID := "providerID_example" // string | 
    idpAttrID := "idpAttrID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut(context.Background(), envID, providerID, idpAttrID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 
**idpAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost

> V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost(ctx, envID, providerID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

CREATE Identity Provider Attribute (SAML)



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
    providerID := "providerID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost(context.Background(), envID, providerID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDIdentityProvidersProviderIDAttributesPostRequest struct via the builder pattern


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

