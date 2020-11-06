# Go API client for channeladaptor

uim channel adaptor api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./channeladaptor"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:58100/app/wechatmp/adaptor/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ChannelAdaptorApi* | [**ChannelEvent**](docs/ChannelAdaptorApi.md#channelevent) | **Post** /channelEvent | 渠道事件
*ChannelAdaptorApi* | [**ContactDetail**](docs/ChannelAdaptorApi.md#contactdetail) | **Post** /contactDetail | 获取单个联系人详情
*ChannelAdaptorApi* | [**ListContact**](docs/ChannelAdaptorApi.md#listcontact) | **Post** /listContact | 获取运营号所有联系人信息
*ChannelAdaptorApi* | [**ProfileDetail**](docs/ChannelAdaptorApi.md#profiledetail) | **Post** /profileDetail | 获取运营号详情
*ChannelAdaptorApi* | [**SendCommand**](docs/ChannelAdaptorApi.md#sendcommand) | **Post** /sendCommand | 通用命令
*ChannelAdaptorApi* | [**SendMessage**](docs/ChannelAdaptorApi.md#sendmessage) | **Post** /sendMessage | 往给定channel发送消息
*ChannelAdaptorApi* | [**UpdateContact**](docs/ChannelAdaptorApi.md#updatecontact) | **Post** /updateContact | 更新联系人信息
*ChannelAdaptorApi* | [**UpdateProfile**](docs/ChannelAdaptorApi.md#updateprofile) | **Post** /updateProfile | 更新运营号信息


## Documentation For Models

 - [AdaptorApiEvent](docs/AdaptorApiEvent.md)
 - [AdaptorCallbackEvent](docs/AdaptorCallbackEvent.md)
 - [AddContactEvent](docs/AddContactEvent.md)
 - [ApiResponse](docs/ApiResponse.md)
 - [Contact](docs/Contact.md)
 - [ContactDetailReq](docs/ContactDetailReq.md)
 - [ContactMessageEvent](docs/ContactMessageEvent.md)
 - [DelContactEvent](docs/DelContactEvent.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [ListContactReq](docs/ListContactReq.md)
 - [Profile](docs/Profile.md)
 - [ProfileDetailReq](docs/ProfileDetailReq.md)
 - [SendCommandReq](docs/SendCommandReq.md)
 - [SendCommandRsp](docs/SendCommandRsp.md)
 - [SendMessageReq](docs/SendMessageReq.md)
 - [UpdateContactReq](docs/UpdateContactReq.md)
 - [UpdateProfileReq](docs/UpdateProfileReq.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



