# \ManagementAPIsPopulationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePopulation**](ManagementAPIsPopulationsApi.md#CreatePopulation) | **Post** /v1/environments/{envID}/populations | CREATE Population
[**DeletePopulation**](ManagementAPIsPopulationsApi.md#DeletePopulation) | **Delete** /v1/environments/{envID}/populations/{popID} | DELETE Population
[**ReadAllPopulations**](ManagementAPIsPopulationsApi.md#ReadAllPopulations) | **Get** /v1/environments/{envID}/populations | READ All Populations
[**ReadOnePopulation**](ManagementAPIsPopulationsApi.md#ReadOnePopulation) | **Get** /v1/environments/{envID}/populations/{popID} | READ One Population
[**UpdatePopulation**](ManagementAPIsPopulationsApi.md#UpdatePopulation) | **Put** /v1/environments/{envID}/populations/{popID} | UPDATE Population



## CreatePopulation

> Population CreatePopulation(ctx, envID).Population(population).ContentType(contentType).Execute()

CREATE Population



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
    population := *openapiclient.NewPopulation() // Population | 
    contentType := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPopulationsApi.CreatePopulation(context.Background(), envID).Population(population).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPopulationsApi.CreatePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePopulation`: Population
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsPopulationsApi.CreatePopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **population** | [**Population**](Population.md) |  | 
 **contentType** | **string** |  | 

### Return type

[**Population**](Population.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePopulation

> DeletePopulation(ctx, envID, popID).Execute()

DELETE Population



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
    popID := "popID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPopulationsApi.DeletePopulation(context.Background(), envID, popID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPopulationsApi.DeletePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**popID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePopulationRequest struct via the builder pattern


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


## ReadAllPopulations

> EntityArray ReadAllPopulations(ctx, envID).Limit(limit).Filter(filter).Execute()

READ All Populations



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
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)
    filter := "id eq "60971d3b-cc5a-4601-9c44-2be541f91bf1"" // string | Adding a SCIM filter for a population ID or population name to display only those resources associated with the specified population. Only the id and name parameters are supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPopulationsApi.ReadAllPopulations(context.Background(), envID).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPopulationsApi.ReadAllPopulations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllPopulations`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsPopulationsApi.ReadAllPopulations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllPopulationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **filter** | **string** | Adding a SCIM filter for a population ID or population name to display only those resources associated with the specified population. Only the id and name parameters are supported | 

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


## ReadOnePopulation

> Population ReadOnePopulation(ctx, envID, popID).Execute()

READ One Population



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
    popID := "popID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPopulationsApi.ReadOnePopulation(context.Background(), envID, popID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPopulationsApi.ReadOnePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOnePopulation`: Population
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsPopulationsApi.ReadOnePopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**popID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOnePopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Population**](Population.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePopulation

> UpdatePopulation(ctx, envID, popID).ContentType(contentType).Population(population).Execute()

UPDATE Population



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
    popID := "popID_example" // string | 
    contentType := "application/json" // string |  (optional)
    population := *openapiclient.NewPopulation() // Population |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsPopulationsApi.UpdatePopulation(context.Background(), envID, popID).ContentType(contentType).Population(population).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsPopulationsApi.UpdatePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**popID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **population** | [**Population**](Population.md) |  | 

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

