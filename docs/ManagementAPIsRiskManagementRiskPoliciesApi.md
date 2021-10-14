# \ManagementAPIsRiskManagementRiskPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskPolicySet**](ManagementAPIsRiskManagementRiskPoliciesApi.md#CreateRiskPolicySet) | **Post** /v1/environments/{envID}/riskPolicySets | CREATE Risk Policy Set
[**DeleteRiskPolicySet**](ManagementAPIsRiskManagementRiskPoliciesApi.md#DeleteRiskPolicySet) | **Delete** /v1/environments/{envID}/riskPolicySets/{riskPolicySetID} | DELETE Risk Policy Set 
[**ReadOneRiskPolicySet**](ManagementAPIsRiskManagementRiskPoliciesApi.md#ReadOneRiskPolicySet) | **Get** /v1/environments/{envID}/riskPolicySets/{riskPolicySetID} | READ One Risk Policy Set
[**ReadRiskPolicySets**](ManagementAPIsRiskManagementRiskPoliciesApi.md#ReadRiskPolicySets) | **Get** /v1/environments/{envID}/riskPolicySets | READ Risk Policy Sets
[**UpdateRiskPolicySet**](ManagementAPIsRiskManagementRiskPoliciesApi.md#UpdateRiskPolicySet) | **Put** /v1/environments/{envID}/riskPolicySets/{riskPolicySetID} | UPDATE Risk Policy Set



## CreateRiskPolicySet

> RiskPolicySet CreateRiskPolicySet(ctx, envID).RiskPolicySet(riskPolicySet).Execute()

CREATE Risk Policy Set



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
    riskPolicySet := *openapiclient.NewRiskPolicySet("Name_example") // RiskPolicySet |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.CreateRiskPolicySet(context.Background(), envID).RiskPolicySet(riskPolicySet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.CreateRiskPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskPolicySet`: RiskPolicySet
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskPoliciesApi.CreateRiskPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskPolicySet** | [**RiskPolicySet**](RiskPolicySet.md) |  | 

### Return type

[**RiskPolicySet**](RiskPolicySet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskPolicySet

> DeleteRiskPolicySet(ctx, envID, riskPolicySetID).Execute()

DELETE Risk Policy Set 



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
    riskPolicySetID := "riskPolicySetID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.DeleteRiskPolicySet(context.Background(), envID, riskPolicySetID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.DeleteRiskPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskPolicySetRequest struct via the builder pattern


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


## ReadOneRiskPolicySet

> RiskPolicySet ReadOneRiskPolicySet(ctx, envID, riskPolicySetID).Execute()

READ One Risk Policy Set



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
    riskPolicySetID := "riskPolicySetID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.ReadOneRiskPolicySet(context.Background(), envID, riskPolicySetID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.ReadOneRiskPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneRiskPolicySet`: RiskPolicySet
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskPoliciesApi.ReadOneRiskPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRiskPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskPolicySet**](RiskPolicySet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRiskPolicySets

> EntityArray ReadRiskPolicySets(ctx, envID).Execute()

READ Risk Policy Sets



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
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.ReadRiskPolicySets(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.ReadRiskPolicySets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRiskPolicySets`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskPoliciesApi.ReadRiskPolicySets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRiskPolicySetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskPolicySet

> RiskPolicySet UpdateRiskPolicySet(ctx, envID, riskPolicySetID).RiskPolicySet(riskPolicySet).Execute()

UPDATE Risk Policy Set



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
    riskPolicySetID := "riskPolicySetID_example" // string | 
    riskPolicySet := *openapiclient.NewRiskPolicySet("Name_example") // RiskPolicySet |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsRiskManagementRiskPoliciesApi.UpdateRiskPolicySet(context.Background(), envID, riskPolicySetID).RiskPolicySet(riskPolicySet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsRiskManagementRiskPoliciesApi.UpdateRiskPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskPolicySet`: RiskPolicySet
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsRiskManagementRiskPoliciesApi.UpdateRiskPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **riskPolicySet** | [**RiskPolicySet**](RiskPolicySet.md) |  | 

### Return type

[**RiskPolicySet**](RiskPolicySet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

