# \ManagementAPIsCertificateManagementApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDCertificatesCertIDApplicationsGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDCertificatesCertIDApplicationsGet) | **Get** /v1/environments/{envID}/certificates/{certID}/applications | GET Certificate Applications
[**V1EnvironmentsEnvIDCertificatesCertIDDelete**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDCertificatesCertIDDelete) | **Delete** /v1/environments/{envID}/certificates/{certID} | DELETE Certificate
[**V1EnvironmentsEnvIDCertificatesCertIDGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDCertificatesCertIDGet) | **Get** /v1/environments/{envID}/certificates/{certID} | GET Certificate
[**V1EnvironmentsEnvIDCertificatesGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDCertificatesGet) | **Get** /v1/environments/{envID}/certificates | GET Certificates
[**V1EnvironmentsEnvIDCertificatesPost**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDCertificatesPost) | **Post** /v1/environments/{envID}/certificates | CREATE Certificate with PKCS7 or PEM File
[**V1EnvironmentsEnvIDDecryptionsPost**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDDecryptionsPost) | **Post** /v1/environments/{envID}/decryptions | DECRYPT Data
[**V1EnvironmentsEnvIDEncryptionsPost**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDEncryptionsPost) | **Post** /v1/environments/{envID}/encryptions | ENCRYPT Data
[**V1EnvironmentsEnvIDKeysGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysGet) | **Get** /v1/environments/{envID}/keys | GET Keys
[**V1EnvironmentsEnvIDKeysKeyIDApplicationsGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysKeyIDApplicationsGet) | **Get** /v1/environments/{envID}/keys/{keyID}/applications | GET Key Applications
[**V1EnvironmentsEnvIDKeysKeyIDCsrGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysKeyIDCsrGet) | **Get** /v1/environments/{envID}/keys/{keyID}/csr | Export a certificate signing request (CSR)
[**V1EnvironmentsEnvIDKeysKeyIDCsrPut**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysKeyIDCsrPut) | **Put** /v1/environments/{envID}/keys/{keyID}/csr | Import Certificate Authority (CA) Response to a CSR
[**V1EnvironmentsEnvIDKeysKeyIDDelete**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysKeyIDDelete) | **Delete** /v1/environments/{envID}/keys/{keyID} | DELETE Key
[**V1EnvironmentsEnvIDKeysKeyIDGet**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysKeyIDGet) | **Get** /v1/environments/{envID}/keys/{keyID} | EXPORT Public Key (X509 PEM)
[**V1EnvironmentsEnvIDKeysKeyIDPut**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysKeyIDPut) | **Put** /v1/environments/{envID}/keys/{keyID} | UPDATE Key
[**V1EnvironmentsEnvIDKeysPost**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDKeysPost) | **Post** /v1/environments/{envID}/keys | CREATE Key with PKCS12 File
[**V1EnvironmentsEnvIDSigningsPost**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDSigningsPost) | **Post** /v1/environments/{envID}/signings | SIGN Data
[**V1EnvironmentsEnvIDVerificationsPost**](ManagementAPIsCertificateManagementApi.md#V1EnvironmentsEnvIDVerificationsPost) | **Post** /v1/environments/{envID}/verifications | VERIFY Signed Data



## V1EnvironmentsEnvIDCertificatesCertIDApplicationsGet

> V1EnvironmentsEnvIDCertificatesCertIDApplicationsGet(ctx, envID, certID).Authorization(authorization).Execute()

GET Certificate Applications



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
    certID := "certID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesCertIDApplicationsGet(context.Background(), envID, certID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesCertIDApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCertificatesCertIDApplicationsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCertificatesCertIDDelete

> V1EnvironmentsEnvIDCertificatesCertIDDelete(ctx, envID, certID).Authorization(authorization).Execute()

DELETE Certificate



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
    certID := "certID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesCertIDDelete(context.Background(), envID, certID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesCertIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCertificatesCertIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCertificatesCertIDGet

> V1EnvironmentsEnvIDCertificatesCertIDGet(ctx, envID, certID).Authorization(authorization).Execute()

GET Certificate



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
    certID := "certID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesCertIDGet(context.Background(), envID, certID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesCertIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCertificatesCertIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCertificatesGet

> V1EnvironmentsEnvIDCertificatesGet(ctx, envID).Authorization(authorization).Execute()

GET Certificates



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
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCertificatesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDCertificatesPost

> V1EnvironmentsEnvIDCertificatesPost(ctx, envID).ContentType(contentType).Authorization(authorization).UsageType(usageType).File(file).Execute()

CREATE Certificate with PKCS7 or PEM File



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
    contentType := "application/x-pkcs7-certificates" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    usageType := "usageType_example" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).UsageType(usageType).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDCertificatesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDCertificatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **authorization** | **string** |  | 
 **usageType** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDDecryptionsPost

> V1EnvironmentsEnvIDDecryptionsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

DECRYPT Data



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
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDDecryptionsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDDecryptionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDDecryptionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDEncryptionsPost

> V1EnvironmentsEnvIDEncryptionsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

ENCRYPT Data



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
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDEncryptionsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDEncryptionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDEncryptionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDKeysGet

> V1EnvironmentsEnvIDKeysGet(ctx, envID).Authorization(authorization).Execute()

GET Keys



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
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDKeysKeyIDApplicationsGet

> V1EnvironmentsEnvIDKeysKeyIDApplicationsGet(ctx, envID, keyID).Authorization(authorization).Execute()

GET Key Applications



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
    keyID := "keyID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDApplicationsGet(context.Background(), envID, keyID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysKeyIDApplicationsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDKeysKeyIDCsrGet

> V1EnvironmentsEnvIDKeysKeyIDCsrGet(ctx, envID, keyID).Authorization(authorization).ContentType(contentType).Execute()

Export a certificate signing request (CSR)



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
    keyID := "keyID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    contentType := "application/pkcs10" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDCsrGet(context.Background(), envID, keyID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDCsrGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysKeyIDCsrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvIDKeysKeyIDCsrPut

> V1EnvironmentsEnvIDKeysKeyIDCsrPut(ctx, envID, keyID).ContentType(contentType).Authorization(authorization).Execute()

Import Certificate Authority (CA) Response to a CSR



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
    keyID := "keyID_example" // string | 
    contentType := "application/x-pem-file" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDCsrPut(context.Background(), envID, keyID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDCsrPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysKeyIDCsrPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1EnvironmentsEnvIDKeysKeyIDDelete

> V1EnvironmentsEnvIDKeysKeyIDDelete(ctx, envID, keyID).Authorization(authorization).Execute()

DELETE Key



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
    keyID := "keyID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDDelete(context.Background(), envID, keyID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysKeyIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDKeysKeyIDGet

> V1EnvironmentsEnvIDKeysKeyIDGet(ctx, envID, keyID).ContentType(contentType).Authorization(authorization).Execute()

EXPORT Public Key (X509 PEM)



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
    keyID := "keyID_example" // string | 
    contentType := "application/x-pkcs7-certificates" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDGet(context.Background(), envID, keyID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysKeyIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1EnvironmentsEnvIDKeysKeyIDPut

> V1EnvironmentsEnvIDKeysKeyIDPut(ctx, envID, keyID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Key



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
    keyID := "keyID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDPut(context.Background(), envID, keyID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysKeyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysKeyIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDKeysPost

> V1EnvironmentsEnvIDKeysPost(ctx, envID).ContentType(contentType).Authorization(authorization).UsageType(usageType).File(file).Execute()

CREATE Key with PKCS12 File



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
    contentType := "application/x-pkcs12-certificates" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    usageType := "usageType_example" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).UsageType(usageType).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDKeysPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **authorization** | **string** |  | 
 **usageType** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDSigningsPost

> V1EnvironmentsEnvIDSigningsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

SIGN Data



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
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDSigningsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDSigningsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDSigningsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDVerificationsPost

> V1EnvironmentsEnvIDVerificationsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

VERIFY Signed Data



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
    resp, r, err := api_client.ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDVerificationsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsCertificateManagementApi.V1EnvironmentsEnvIDVerificationsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDVerificationsPostRequest struct via the builder pattern


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

