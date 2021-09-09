# \ManagementAPIsAlertingApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDAlertChannelsAlertChannelIDPut**](ManagementAPIsAlertingApi.md#V1EnvironmentsEnvIDAlertChannelsAlertChannelIDPut) | **Put** /v1/environments/{envID}/alertChannels/{alertChannelID} | UPDATE Alert Channel
[**V1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDelete**](ManagementAPIsAlertingApi.md#V1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDelete) | **Delete** /v1/environments/{envID}/alertChannels/{alertChannelsID} | DELETE Alert Channel
[**V1EnvironmentsEnvIDAlertChannelsGet**](ManagementAPIsAlertingApi.md#V1EnvironmentsEnvIDAlertChannelsGet) | **Get** /v1/environments/{envID}/alertChannels | READ All Alert Channels per Environment
[**V1EnvironmentsEnvIDAlertChannelsPost**](ManagementAPIsAlertingApi.md#V1EnvironmentsEnvIDAlertChannelsPost) | **Post** /v1/environments/{envID}/alertChannels | CREATE Alert Channel (Email)



## V1EnvironmentsEnvIDAlertChannelsAlertChannelIDPut

> V1EnvironmentsEnvIDAlertChannelsAlertChannelIDPut(ctx, envID, alertChannelID).Body(body).Execute()

UPDATE Alert Channel



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
    alertChannelID := "alertChannelID_example" // string | 
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsAlertChannelIDPut(context.Background(), envID, alertChannelID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsAlertChannelIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**alertChannelID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAlertChannelsAlertChannelIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDelete

> V1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDelete(ctx, envID, alertChannelsID).Execute()

DELETE Alert Channel



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
    alertChannelsID := "alertChannelsID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDelete(context.Background(), envID, alertChannelsID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**alertChannelsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAlertChannelsAlertChannelsIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAlertChannelsGet

> V1EnvironmentsEnvIDAlertChannelsGet(ctx, envID).Execute()

READ All Alert Channels per Environment



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
    resp, r, err := api_client.ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAlertChannelsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDAlertChannelsPost

> V1EnvironmentsEnvIDAlertChannelsPost(ctx, envID).Body(body).Execute()

CREATE Alert Channel (Email)



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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsPost(context.Background(), envID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsAlertingApi.V1EnvironmentsEnvIDAlertChannelsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDAlertChannelsPostRequest struct via the builder pattern


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

