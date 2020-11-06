# \ChannelAdaptorApi

All URIs are relative to *http://localhost:58100/app/wechatmp/adaptor/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelEvent**](ChannelAdaptorApi.md#ChannelEvent) | **Post** /channelEvent | 渠道事件
[**ContactDetail**](ChannelAdaptorApi.md#ContactDetail) | **Post** /contactDetail | 获取单个联系人详情
[**ListContact**](ChannelAdaptorApi.md#ListContact) | **Post** /listContact | 获取运营号所有联系人信息
[**ProfileDetail**](ChannelAdaptorApi.md#ProfileDetail) | **Post** /profileDetail | 获取运营号详情
[**SendCommand**](ChannelAdaptorApi.md#SendCommand) | **Post** /sendCommand | 通用命令
[**SendMessage**](ChannelAdaptorApi.md#SendMessage) | **Post** /sendMessage | 往给定channel发送消息
[**UpdateContact**](ChannelAdaptorApi.md#UpdateContact) | **Post** /updateContact | 更新联系人信息
[**UpdateProfile**](ChannelAdaptorApi.md#UpdateProfile) | **Post** /updateProfile | 更新运营号信息



## ChannelEvent

> InlineResponse2004 ChannelEvent(ctx, body)

渠道事件

渠道事件

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AdaptorApiEvent**](AdaptorApiEvent.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactDetail

> InlineResponse2002 ContactDetail(ctx, body)

获取单个联系人详情

获取单个联系人详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ContactDetailReq**](ContactDetailReq.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContact

> InlineResponse2001 ListContact(ctx, body)

获取运营号所有联系人信息

获取运营号所有联系人信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ListContactReq**](ListContactReq.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfileDetail

> InlineResponse2003 ProfileDetail(ctx, body)

获取运营号详情

获取运营号详情

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ProfileDetailReq**](ProfileDetailReq.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCommand

> InlineResponse2005 SendCommand(ctx, body)

通用命令

通用命令

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SendCommandReq**](SendCommandReq.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> InlineResponse200 SendMessage(ctx, body)

往给定channel发送消息

往给定channel发送消息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SendMessageReq**](SendMessageReq.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> ApiResponse UpdateContact(ctx, body)

更新联系人信息

更新联系人信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdateContactReq**](UpdateContactReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> ApiResponse UpdateProfile(ctx, body)

更新运营号信息

更新运营号信息

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UpdateProfileReq**](UpdateProfileReq.md)|  | 

### Return type

[**ApiResponse**](APIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

