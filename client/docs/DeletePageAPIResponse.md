# DeletePageAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewDeletePageAPIResponse

`func NewDeletePageAPIResponse(status string, ) *DeletePageAPIResponse`

NewDeletePageAPIResponse instantiates a new DeletePageAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePageAPIResponseWithDefaults

`func NewDeletePageAPIResponseWithDefaults() *DeletePageAPIResponse`

NewDeletePageAPIResponseWithDefaults instantiates a new DeletePageAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *DeletePageAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeletePageAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeletePageAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeletePageAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *DeletePageAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeletePageAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeletePageAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DeletePageAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *DeletePageAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeletePageAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeletePageAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


