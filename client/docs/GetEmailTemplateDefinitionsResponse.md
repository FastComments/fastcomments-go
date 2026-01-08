# GetEmailTemplateDefinitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Definitions** | [**[]EmailTemplateDefinition**](EmailTemplateDefinition.md) |  | 

## Methods

### NewGetEmailTemplateDefinitionsResponse

`func NewGetEmailTemplateDefinitionsResponse(status APIStatus, definitions []EmailTemplateDefinition, ) *GetEmailTemplateDefinitionsResponse`

NewGetEmailTemplateDefinitionsResponse instantiates a new GetEmailTemplateDefinitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailTemplateDefinitionsResponseWithDefaults

`func NewGetEmailTemplateDefinitionsResponseWithDefaults() *GetEmailTemplateDefinitionsResponse`

NewGetEmailTemplateDefinitionsResponseWithDefaults instantiates a new GetEmailTemplateDefinitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEmailTemplateDefinitionsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEmailTemplateDefinitionsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEmailTemplateDefinitionsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetDefinitions

`func (o *GetEmailTemplateDefinitionsResponse) GetDefinitions() []EmailTemplateDefinition`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *GetEmailTemplateDefinitionsResponse) GetDefinitionsOk() (*[]EmailTemplateDefinition, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *GetEmailTemplateDefinitionsResponse) SetDefinitions(v []EmailTemplateDefinition)`

SetDefinitions sets Definitions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


