# \ManagementAPIsLicensesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrganizationsOrgIDLicensesGet**](ManagementAPIsLicensesApi.md#V1OrganizationsOrgIDLicensesGet) | **Get** /v1/organizations/{orgID}/licenses | READ All Licenses
[**V1OrganizationsOrgIDLicensesLicenseIDGet**](ManagementAPIsLicensesApi.md#V1OrganizationsOrgIDLicensesLicenseIDGet) | **Get** /v1/organizations/{orgID}/licenses/{licenseID} | READ One License
[**V1OrganizationsOrgIDLicensesLicenseIDNameGet**](ManagementAPIsLicensesApi.md#V1OrganizationsOrgIDLicensesLicenseIDNameGet) | **Get** /v1/organizations/{orgID}/licenses/{licenseID}/name | READ One License Name
[**V1OrganizationsOrgIDLicensesLicenseIDNamePut**](ManagementAPIsLicensesApi.md#V1OrganizationsOrgIDLicensesLicenseIDNamePut) | **Put** /v1/organizations/{orgID}/licenses/{licenseID}/name | Update One License Name



## V1OrganizationsOrgIDLicensesGet

> V1OrganizationsOrgIDLicensesGet(ctx, orgID).Execute()

READ All Licenses



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
    orgID := "orgID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesGet(context.Background(), orgID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIDLicensesGetRequest struct via the builder pattern


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


## V1OrganizationsOrgIDLicensesLicenseIDGet

> V1OrganizationsOrgIDLicensesLicenseIDGet(ctx, orgID, licenseID).Execute()

READ One License



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
    orgID := "orgID_example" // string | 
    licenseID := "licenseID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesLicenseIDGet(context.Background(), orgID, licenseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesLicenseIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIDLicensesLicenseIDGetRequest struct via the builder pattern


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


## V1OrganizationsOrgIDLicensesLicenseIDNameGet

> V1OrganizationsOrgIDLicensesLicenseIDNameGet(ctx, orgID, licenseID).Execute()

READ One License Name



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
    orgID := "orgID_example" // string | 
    licenseID := "licenseID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesLicenseIDNameGet(context.Background(), orgID, licenseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesLicenseIDNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIDLicensesLicenseIDNameGetRequest struct via the builder pattern


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


## V1OrganizationsOrgIDLicensesLicenseIDNamePut

> V1OrganizationsOrgIDLicensesLicenseIDNamePut(ctx, orgID, licenseID).Body(body).Execute()

Update One License Name



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
    orgID := "orgID_example" // string | 
    licenseID := "licenseID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesLicenseIDNamePut(context.Background(), orgID, licenseID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsLicensesApi.V1OrganizationsOrgIDLicensesLicenseIDNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIDLicensesLicenseIDNamePutRequest struct via the builder pattern


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

