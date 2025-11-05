# PutSSOUserAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**NullableAPISSOUser**](APISSOUser.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewPutSSOUserAPIResponse

`func NewPutSSOUserAPIResponse(status string, ) *PutSSOUserAPIResponse`

NewPutSSOUserAPIResponse instantiates a new PutSSOUserAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutSSOUserAPIResponseWithDefaults

`func NewPutSSOUserAPIResponseWithDefaults() *PutSSOUserAPIResponse`

NewPutSSOUserAPIResponseWithDefaults instantiates a new PutSSOUserAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *PutSSOUserAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PutSSOUserAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PutSSOUserAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PutSSOUserAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *PutSSOUserAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PutSSOUserAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PutSSOUserAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PutSSOUserAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUser

`func (o *PutSSOUserAPIResponse) GetUser() APISSOUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PutSSOUserAPIResponse) GetUserOk() (*APISSOUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PutSSOUserAPIResponse) SetUser(v APISSOUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PutSSOUserAPIResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PutSSOUserAPIResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PutSSOUserAPIResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetStatus

`func (o *PutSSOUserAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutSSOUserAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutSSOUserAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


