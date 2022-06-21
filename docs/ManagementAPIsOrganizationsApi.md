# \ManagementAPIsOrganizationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrganizationsGet**](ManagementAPIsOrganizationsApi.md#V1OrganizationsGet) | **Get** /v1/organizations | READ All Organizations
[**V1OrganizationsOrgIDGet**](ManagementAPIsOrganizationsApi.md#V1OrganizationsOrgIDGet) | **Get** /v1/organizations/{orgID} | READ One Organization



## V1OrganizationsGet

> V1OrganizationsGet(ctx).Execute()

READ All Organizations



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsOrganizationsApi.V1OrganizationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsOrganizationsApi.V1OrganizationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsGetRequest struct via the builder pattern


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


## V1OrganizationsOrgIDGet

> V1OrganizationsOrgIDGet(ctx, orgID).Execute()

READ One Organization



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
    resp, r, err := apiClient.ManagementAPIsOrganizationsApi.V1OrganizationsOrgIDGet(context.Background(), orgID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsOrganizationsApi.V1OrganizationsOrgIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrgIDGetRequest struct via the builder pattern


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

