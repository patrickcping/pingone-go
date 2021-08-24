# \AuthenticationAPIsSAML20Api

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvIDSaml20IdpSloGet**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20IdpSloGet) | **Get** /{envID}/saml20/idp/slo | SAML SLO Using GET
[**EnvIDSaml20IdpSloPost**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20IdpSloPost) | **Post** /{envID}/saml20/idp/slo | SAML SLO Using POST
[**EnvIDSaml20IdpSsoGet**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20IdpSsoGet) | **Get** /{envID}/saml20/idp/sso | SAML SSO Using GET
[**EnvIDSaml20IdpSsoPost**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20IdpSsoPost) | **Post** /{envID}/saml20/idp/sso | SAML SSO Using POST
[**EnvIDSaml20IdpStartssoGet**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20IdpStartssoGet) | **Get** /{envID}/saml20/idp/startsso | Identity Provider Initiated SSO
[**EnvIDSaml20MetadataAppIDGet**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20MetadataAppIDGet) | **Get** /{envID}/saml20/metadata/{appID} | READ SAML Metadata
[**EnvIDSaml20SpAcsPost**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20SpAcsPost) | **Post** /{envID}/saml20/sp/acs | SAML ACS Endpoint for Identity Provider Initiated Inbound SSO
[**EnvIDSaml20SpMetadataIdpIDGet**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20SpMetadataIdpIDGet) | **Get** /{envID}/saml20/sp/metadata/{idpID} | READ SAML Service Provider Metadata
[**EnvIDSaml20SpSsoGet**](AuthenticationAPIsSAML20Api.md#EnvIDSaml20SpSsoGet) | **Get** /{envID}/saml20/sp/sso | Service Provider Initiated Inbound SSO



## EnvIDSaml20IdpSloGet

> EnvIDSaml20IdpSloGet(ctx, envID).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()

SAML SLO Using GET



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
    sAMLRequest := "{{SAMLRequest}}" // string |  (optional)
    relayState := "testing..." // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSloGet(context.Background(), envID).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSloGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20IdpSloGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sAMLRequest** | **string** |  | 
 **relayState** | **string** |  | 

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


## EnvIDSaml20IdpSloPost

> EnvIDSaml20IdpSloPost(ctx, envID).ContentType(contentType).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()

SAML SLO Using POST



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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    sAMLRequest := "sAMLRequest_example" // string |  (optional)
    relayState := "relayState_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSloPost(context.Background(), envID).ContentType(contentType).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSloPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20IdpSloPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **sAMLRequest** | **string** |  | 
 **relayState** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvIDSaml20IdpSsoGet

> EnvIDSaml20IdpSsoGet(ctx, envID).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()

SAML SSO Using GET



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
    sAMLRequest := "{{SAMLRequest}}" // string |  (optional)
    relayState := "token" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSsoGet(context.Background(), envID).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSsoGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20IdpSsoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sAMLRequest** | **string** |  | 
 **relayState** | **string** |  | 

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


## EnvIDSaml20IdpSsoPost

> EnvIDSaml20IdpSsoPost(ctx, envID).ContentType(contentType).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()

SAML SSO Using POST



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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    sAMLRequest := "sAMLRequest_example" // string |  (optional)
    relayState := "relayState_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSsoPost(context.Background(), envID).ContentType(contentType).SAMLRequest(sAMLRequest).RelayState(relayState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20IdpSsoPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20IdpSsoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **sAMLRequest** | **string** |  | 
 **relayState** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvIDSaml20IdpStartssoGet

> EnvIDSaml20IdpStartssoGet(ctx, envID).ContentType(contentType).SpEntityId(spEntityId).ApplicationUrl(applicationUrl).Execute()

Identity Provider Initiated SSO



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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    spEntityId := "{{spEntityIdValue}}" // string |  (optional)
    applicationUrl := "{{AppUrl}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20IdpStartssoGet(context.Background(), envID).ContentType(contentType).SpEntityId(spEntityId).ApplicationUrl(applicationUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20IdpStartssoGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20IdpStartssoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **spEntityId** | **string** |  | 
 **applicationUrl** | **string** |  | 

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


## EnvIDSaml20MetadataAppIDGet

> EnvIDSaml20MetadataAppIDGet(ctx, envID, appID).Execute()

READ SAML Metadata



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20MetadataAppIDGet(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20MetadataAppIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20MetadataAppIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## EnvIDSaml20SpAcsPost

> EnvIDSaml20SpAcsPost(ctx, envID).ContentType(contentType).RelayState(relayState).Execute()

SAML ACS Endpoint for Identity Provider Initiated Inbound SSO



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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    relayState := "relayState_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20SpAcsPost(context.Background(), envID).ContentType(contentType).RelayState(relayState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20SpAcsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20SpAcsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **relayState** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvIDSaml20SpMetadataIdpIDGet

> EnvIDSaml20SpMetadataIdpIDGet(ctx, envID, idpID).Execute()

READ SAML Service Provider Metadata



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
    idpID := "idpID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20SpMetadataIdpIDGet(context.Background(), envID, idpID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20SpMetadataIdpIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**idpID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvIDSaml20SpMetadataIdpIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## EnvIDSaml20SpSsoGet

> EnvIDSaml20SpSsoGet(ctx, envID).IdpId(idpId).FlowId(flowId).Execute()

Service Provider Initiated Inbound SSO



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
    idpId := "{{extIdpID}}" // string |  (optional)
    flowId := "{{flowID}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsSAML20Api.EnvIDSaml20SpSsoGet(context.Background(), envID).IdpId(idpId).FlowId(flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsSAML20Api.EnvIDSaml20SpSsoGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvIDSaml20SpSsoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpId** | **string** |  | 
 **flowId** | **string** |  | 

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

