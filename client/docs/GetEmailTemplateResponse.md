# GetEmailTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**EmailTemplate** | [**CustomEmailTemplate**](CustomEmailTemplate.md) |  | 

## Methods

### NewGetEmailTemplateResponse

`func NewGetEmailTemplateResponse(status APIStatus, emailTemplate CustomEmailTemplate, ) *GetEmailTemplateResponse`

NewGetEmailTemplateResponse instantiates a new GetEmailTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailTemplateResponseWithDefaults

`func NewGetEmailTemplateResponseWithDefaults() *GetEmailTemplateResponse`

NewGetEmailTemplateResponseWithDefaults instantiates a new GetEmailTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEmailTemplateResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEmailTemplateResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEmailTemplateResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetEmailTemplate

`func (o *GetEmailTemplateResponse) GetEmailTemplate() CustomEmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *GetEmailTemplateResponse) GetEmailTemplateOk() (*CustomEmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *GetEmailTemplateResponse) SetEmailTemplate(v CustomEmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


