# GetEmailTemplateRenderErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**RenderErrors** | [**[]EmailTemplateRenderErrorResponse**](EmailTemplateRenderErrorResponse.md) |  | 

## Methods

### NewGetEmailTemplateRenderErrorsResponse

`func NewGetEmailTemplateRenderErrorsResponse(status APIStatus, renderErrors []EmailTemplateRenderErrorResponse, ) *GetEmailTemplateRenderErrorsResponse`

NewGetEmailTemplateRenderErrorsResponse instantiates a new GetEmailTemplateRenderErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailTemplateRenderErrorsResponseWithDefaults

`func NewGetEmailTemplateRenderErrorsResponseWithDefaults() *GetEmailTemplateRenderErrorsResponse`

NewGetEmailTemplateRenderErrorsResponseWithDefaults instantiates a new GetEmailTemplateRenderErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEmailTemplateRenderErrorsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEmailTemplateRenderErrorsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEmailTemplateRenderErrorsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetRenderErrors

`func (o *GetEmailTemplateRenderErrorsResponse) GetRenderErrors() []EmailTemplateRenderErrorResponse`

GetRenderErrors returns the RenderErrors field if non-nil, zero value otherwise.

### GetRenderErrorsOk

`func (o *GetEmailTemplateRenderErrorsResponse) GetRenderErrorsOk() (*[]EmailTemplateRenderErrorResponse, bool)`

GetRenderErrorsOk returns a tuple with the RenderErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderErrors

`func (o *GetEmailTemplateRenderErrorsResponse) SetRenderErrors(v []EmailTemplateRenderErrorResponse)`

SetRenderErrors sets RenderErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


