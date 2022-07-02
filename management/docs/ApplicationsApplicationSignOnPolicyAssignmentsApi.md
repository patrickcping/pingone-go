# \ApplicationsApplicationSignOnPolicyAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGet**](ApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGet) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments | READ All SOP Assignments
[**V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPost**](ApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPost) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments | CREATE SOP Assignment
[**V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDelete**](ApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDelete) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments/{SOPAssignmentID} | DELETE SOP Assignment
[**V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGet**](ApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGet) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments/{SOPAssignmentID} | READ One SOP Assignment
[**V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPut**](ApplicationsApplicationSignOnPolicyAssignmentsApi.md#V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPut) | **Put** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments/{SOPAssignmentID} | UPDATE SOP Assignment



## V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGet

> V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGet(ctx, environmentID, applicationID).Execute()

READ All SOP Assignments

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGet(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPost

> V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPost(ctx, environmentID, applicationID).Body(body).Execute()

CREATE SOP Assignment

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPost(context.Background(), environmentID, applicationID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDelete

> V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDelete(ctx, environmentID, applicationID, sOPAssignmentID).Execute()

DELETE SOP Assignment

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    sOPAssignmentID := "sOPAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDelete(context.Background(), environmentID, applicationID, sOPAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGet

> V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGet(ctx, environmentID, applicationID, sOPAssignmentID).Execute()

READ One SOP Assignment

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    sOPAssignmentID := "sOPAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGet(context.Background(), environmentID, applicationID, sOPAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPut

> V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPut(ctx, environmentID, applicationID, sOPAssignmentID).Body(body).Execute()

UPDATE SOP Assignment

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    sOPAssignmentID := "sOPAssignmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPut(context.Background(), environmentID, applicationID, sOPAssignmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationSignOnPolicyAssignmentsApi.V1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDApplicationsApplicationIDSignOnPolicyAssignmentsSOPAssignmentIDPutRequest struct via the builder pattern


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

