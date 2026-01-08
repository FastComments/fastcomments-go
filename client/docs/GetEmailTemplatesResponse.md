# GetEmailTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**EmailTemplates** | [**[]CustomEmailTemplate**](CustomEmailTemplate.md) |  | 

## Methods

### NewGetEmailTemplatesResponse

`func NewGetEmailTemplatesResponse(status APIStatus, emailTemplates []CustomEmailTemplate, ) *GetEmailTemplatesResponse`

NewGetEmailTemplatesResponse instantiates a new GetEmailTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailTemplatesResponseWithDefaults

`func NewGetEmailTemplatesResponseWithDefaults() *GetEmailTemplatesResponse`

NewGetEmailTemplatesResponseWithDefaults instantiates a new GetEmailTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEmailTemplatesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEmailTemplatesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEmailTemplatesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetEmailTemplates

`func (o *GetEmailTemplatesResponse) GetEmailTemplates() []CustomEmailTemplate`

GetEmailTemplates returns the EmailTemplates field if non-nil, zero value otherwise.

### GetEmailTemplatesOk

`func (o *GetEmailTemplatesResponse) GetEmailTemplatesOk() (*[]CustomEmailTemplate, bool)`

GetEmailTemplatesOk returns a tuple with the EmailTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplates

`func (o *GetEmailTemplatesResponse) SetEmailTemplates(v []CustomEmailTemplate)`

SetEmailTemplates sets EmailTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


