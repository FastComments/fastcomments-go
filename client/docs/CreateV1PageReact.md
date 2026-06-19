# CreateV1PageReact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewCreateV1PageReact

`func NewCreateV1PageReact(status APIStatus, ) *CreateV1PageReact`

NewCreateV1PageReact instantiates a new CreateV1PageReact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateV1PageReactWithDefaults

`func NewCreateV1PageReactWithDefaults() *CreateV1PageReact`

NewCreateV1PageReactWithDefaults instantiates a new CreateV1PageReact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateV1PageReact) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateV1PageReact) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateV1PageReact) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateV1PageReact) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *CreateV1PageReact) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateV1PageReact) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateV1PageReact) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


