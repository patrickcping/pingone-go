# \AuthenticationAPIsOpenIDConnectOAuth2Api

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvIDAsAuthorizeGet**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsAuthorizeGet) | **Get** /v1/{envID}/as/authorize | Authorize (Transaction Approval) 
[**V1EnvIDAsAuthorizePost**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsAuthorizePost) | **Post** /v1/{envID}/as/authorize | Authorize (implicit)
[**V1EnvIDAsIntrospectPost**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsIntrospectPost) | **Post** /v1/{envID}/as/introspect | Token Introspection (Refresh Token)
[**V1EnvIDAsJwksGet**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsJwksGet) | **Get** /v1/{envID}/as/jwks | READ JWKS
[**V1EnvIDAsResumeGet**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsResumeGet) | **Get** /v1/{envID}/as/resume | Resume
[**V1EnvIDAsRevokePost**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsRevokePost) | **Post** /v1/{envID}/as/revoke | Token Revocation
[**V1EnvIDAsSignoffGet**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsSignoffGet) | **Get** /v1/{envID}/as/signoff | Signoff
[**V1EnvIDAsTokenPost**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsTokenPost) | **Post** /v1/{envID}/as/token | Token Exchange (Gateway Credential) 
[**V1EnvIDAsUserinfoGet**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsUserinfoGet) | **Get** /v1/{envID}/as/userinfo | Userinfo
[**V1EnvIDAsUserinfoPost**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsUserinfoPost) | **Post** /v1/{envID}/as/userinfo | Userinfo
[**V1EnvIDAsWellKnownOpenidConfigurationGet**](AuthenticationAPIsOpenIDConnectOAuth2Api.md#V1EnvIDAsWellKnownOpenidConfigurationGet) | **Get** /v1/{envID}/as/.well-known/openid-configuration | Discovery OpenID Configuration



## V1EnvIDAsAuthorizeGet

> V1EnvIDAsAuthorizeGet(ctx, envID).ResponseType(responseType).ClientId(clientId).ResponseMode(responseMode).Scope(scope).State(state).Request(request).Execute()

Authorize (Transaction Approval) 



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
    responseType := "token id_token" // string |  (optional)
    clientId := "6c796712-0f16-4062-815a-e0a92f4a2143" // string |  (optional)
    responseMode := "pi.flow" // string |  (optional)
    scope := "openid" // string |  (optional)
    state := "{string}" // string |  (optional)
    request := "{requestString}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsAuthorizeGet(context.Background(), envID).ResponseType(responseType).ClientId(clientId).ResponseMode(responseMode).Scope(scope).State(state).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsAuthorizeGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsAuthorizeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **responseType** | **string** |  | 
 **clientId** | **string** |  | 
 **responseMode** | **string** |  | 
 **scope** | **string** |  | 
 **state** | **string** |  | 
 **request** | **string** |  | 

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


## V1EnvIDAsAuthorizePost

> V1EnvIDAsAuthorizePost(ctx, envID).ContentType(contentType).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).Nonce(nonce).State(state).Prompt(prompt).MaxAge(maxAge).AcrValues(acrValues).Execute()

Authorize (implicit)



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
    responseType := "responseType_example" // string | Required (optional)
    clientId := "clientId_example" // string | Required (optional)
    redirectUri := "redirectUri_example" // string | Required (optional)
    scope := "scope_example" // string | Required (optional)
    nonce := int32(56) // int32 | Required (optional)
    state := "state_example" // string | Recommended (optional)
    prompt := "prompt_example" // string | Optional ( none | login ) (optional)
    maxAge := int32(56) // int32 | Optional - uses seconds (optional)
    acrValues := "acrValues_example" // string | Optional - use Sign-on Policy names (space-delimited) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsAuthorizePost(context.Background(), envID).ContentType(contentType).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).Nonce(nonce).State(state).Prompt(prompt).MaxAge(maxAge).AcrValues(acrValues).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsAuthorizePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsAuthorizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **responseType** | **string** | Required | 
 **clientId** | **string** | Required | 
 **redirectUri** | **string** | Required | 
 **scope** | **string** | Required | 
 **nonce** | **int32** | Required | 
 **state** | **string** | Recommended | 
 **prompt** | **string** | Optional ( none | login ) | 
 **maxAge** | **int32** | Optional - uses seconds | 
 **acrValues** | **string** | Optional - use Sign-on Policy names (space-delimited) | 

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


## V1EnvIDAsIntrospectPost

> V1EnvIDAsIntrospectPost(ctx, envID).ContentType(contentType).Token(token).ClientId(clientId).ClientSecret(clientSecret).Execute()

Token Introspection (Refresh Token)



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
    token := "token_example" // string | Required (optional)
    clientId := "clientId_example" // string | (For CLIENT_SECRET_POST; if included with Basic Auth, must match Header value) (optional)
    clientSecret := "clientSecret_example" // string | (For for CLIENT_SECRET_POST) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsIntrospectPost(context.Background(), envID).ContentType(contentType).Token(token).ClientId(clientId).ClientSecret(clientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsIntrospectPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsIntrospectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **token** | **string** | Required | 
 **clientId** | **string** | (For CLIENT_SECRET_POST; if included with Basic Auth, must match Header value) | 
 **clientSecret** | **string** | (For for CLIENT_SECRET_POST) | 

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


## V1EnvIDAsJwksGet

> V1EnvIDAsJwksGet(ctx, envID).Execute()

READ JWKS



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsJwksGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsJwksGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsJwksGetRequest struct via the builder pattern


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


## V1EnvIDAsResumeGet

> V1EnvIDAsResumeGet(ctx, envID).Cookie(cookie).FlowId(flowId).Execute()

Resume



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
    cookie := "{{sessionToken}}" // string |  (optional)
    flowId := "{{flowID}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsResumeGet(context.Background(), envID).Cookie(cookie).FlowId(flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsResumeGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsResumeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **string** |  | 
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


## V1EnvIDAsRevokePost

> V1EnvIDAsRevokePost(ctx, envID).ContentType(contentType).Token(token).Execute()

Token Revocation



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
    token := "token_example" // string | Required (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsRevokePost(context.Background(), envID).ContentType(contentType).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsRevokePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsRevokePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **token** | **string** | Required | 

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


## V1EnvIDAsSignoffGet

> V1EnvIDAsSignoffGet(ctx, envID).Cookie(cookie).IdTokenHint(idTokenHint).Execute()

Signoff



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
    cookie := "{{sessionToken}}" // string |  (optional)
    idTokenHint := "{{idToken}}" // string | Required (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsSignoffGet(context.Background(), envID).Cookie(cookie).IdTokenHint(idTokenHint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsSignoffGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsSignoffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **string** |  | 
 **idTokenHint** | **string** | Required | 

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


## V1EnvIDAsTokenPost

> V1EnvIDAsTokenPost(ctx, envID).ContentType(contentType).GrantType(grantType).SubjectTokenType(subjectTokenType).SubjectToken(subjectToken).Execute()

Token Exchange (Gateway Credential) 



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
    grantType := "grantType_example" // string | Required (optional)
    subjectTokenType := "subjectTokenType_example" // string | (For CLIENT_SECRET_POST; if included with Basic Auth, must match Header value) (optional)
    subjectToken := "subjectToken_example" // string | (For for CLIENT_SECRET_POST) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsTokenPost(context.Background(), envID).ContentType(contentType).GrantType(grantType).SubjectTokenType(subjectTokenType).SubjectToken(subjectToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsTokenPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **grantType** | **string** | Required | 
 **subjectTokenType** | **string** | (For CLIENT_SECRET_POST; if included with Basic Auth, must match Header value) | 
 **subjectToken** | **string** | (For for CLIENT_SECRET_POST) | 

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


## V1EnvIDAsUserinfoGet

> V1EnvIDAsUserinfoGet(ctx, envID).Execute()

Userinfo



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsUserinfoGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsUserinfoGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsUserinfoGetRequest struct via the builder pattern


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


## V1EnvIDAsUserinfoPost

> V1EnvIDAsUserinfoPost(ctx, envID).Execute()

Userinfo



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsUserinfoPost(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsUserinfoPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsUserinfoPostRequest struct via the builder pattern


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


## V1EnvIDAsWellKnownOpenidConfigurationGet

> V1EnvIDAsWellKnownOpenidConfigurationGet(ctx, envID).Execute()

Discovery OpenID Configuration



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsWellKnownOpenidConfigurationGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsOpenIDConnectOAuth2Api.V1EnvIDAsWellKnownOpenidConfigurationGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvIDAsWellKnownOpenidConfigurationGetRequest struct via the builder pattern


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

