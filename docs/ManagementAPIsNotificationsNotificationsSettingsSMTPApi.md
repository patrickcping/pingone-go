# \ManagementAPIsNotificationsNotificationsSettingsSMTPApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGet**](ManagementAPIsNotificationsNotificationsSettingsSMTPApi.md#V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGet) | **Get** /v1/environments/{envID}/notificationsSettings/emailDeliverySettings | READ Notifications Settings (SMTP)
[**V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPut**](ManagementAPIsNotificationsNotificationsSettingsSMTPApi.md#V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPut) | **Put** /v1/environments/{envID}/notificationsSettings/emailDeliverySettings | UPDATE Notifications Settings (SMTP)



## V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGet

> V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGet(ctx, envID).Authorization(authorization).Execute()

READ Notifications Settings (SMTP)



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
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGet(context.Background(), envID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPut

> V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPut(ctx, envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()

UPDATE Notifications Settings (SMTP)



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
    resp, r, err := api_client.ManagementAPIsNotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPut(context.Background(), envID).ContentType(contentType).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsNotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDNotificationsSettingsEmailDeliverySettingsPutRequest struct via the builder pattern


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

