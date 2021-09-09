# \ManagementAPIsNotificationsPhoneDeliverySettingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete**](ManagementAPIsNotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete) | **Delete** /v1/environments/{envID}/notificationsSettings/emailDeliverySettings | DELETE Phone Delivery Settings
[**V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet**](ManagementAPIsNotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet) | **Get** /v1/environments/{envID}/notificationsSettings/phoneDeliverySettings | READ All Phone Delivery Settings
[**V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet**](ManagementAPIsNotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet) | **Get** /v1/environments/{envID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsId} | READ One Phone Delivery Settings
[**V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut**](ManagementAPIsNotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut) | **Put** /v1/environments/{envID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsId} | UPDATE Phone Delivery Settings
[**V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost**](ManagementAPIsNotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost) | **Post** /v1/environments/{envID}/notificationsSettings/phoneDeliverySettings | CREATE Phone Delivery Settings (Syniverse)



## V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete

> V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete(ctx, envID).Execute()

DELETE Phone Delivery Settings



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
    resp, r, err := api_client.ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet

> V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet(ctx, envID).Execute()

READ All Phone Delivery Settings



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
    resp, r, err := api_client.ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet

> V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet(ctx, envID, phoneDeliverySettingsId).Execute()

READ One Phone Delivery Settings



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
    phoneDeliverySettingsId := "phoneDeliverySettingsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet(context.Background(), envID, phoneDeliverySettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**phoneDeliverySettingsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut

> V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut(ctx, envID, phoneDeliverySettingsId).ContentType(contentType).Body(body).Execute()

UPDATE Phone Delivery Settings



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
    phoneDeliverySettingsId := "phoneDeliverySettingsId_example" // string | 
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut(context.Background(), envID, phoneDeliverySettingsId).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**phoneDeliverySettingsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost

> V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost(ctx, envID).ContentType(contentType).Body(body).Execute()

CREATE Phone Delivery Settings (Syniverse)



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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost(context.Background(), envID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsPhoneDeliverySettingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
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

