# RenderEmailTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Html** | **string** |  | 

## Methods

### NewRenderEmailTemplateResponse

`func NewRenderEmailTemplateResponse(status APIStatus, html string, ) *RenderEmailTemplateResponse`

NewRenderEmailTemplateResponse instantiates a new RenderEmailTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderEmailTemplateResponseWithDefaults

`func NewRenderEmailTemplateResponseWithDefaults() *RenderEmailTemplateResponse`

NewRenderEmailTemplateResponseWithDefaults instantiates a new RenderEmailTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RenderEmailTemplateResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RenderEmailTemplateResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RenderEmailTemplateResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetHtml

`func (o *RenderEmailTemplateResponse) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *RenderEmailTemplateResponse) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *RenderEmailTemplateResponse) SetHtml(v string)`

SetHtml sets Html field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


