# GetSSOUserByEmailAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**APISSOUser**](APISSOUser.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewGetSSOUserByEmailAPIResponse

`func NewGetSSOUserByEmailAPIResponse(status string, ) *GetSSOUserByEmailAPIResponse`

NewGetSSOUserByEmailAPIResponse instantiates a new GetSSOUserByEmailAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSSOUserByEmailAPIResponseWithDefaults

`func NewGetSSOUserByEmailAPIResponseWithDefaults() *GetSSOUserByEmailAPIResponse`

NewGetSSOUserByEmailAPIResponseWithDefaults instantiates a new GetSSOUserByEmailAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *GetSSOUserByEmailAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetSSOUserByEmailAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetSSOUserByEmailAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetSSOUserByEmailAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *GetSSOUserByEmailAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSSOUserByEmailAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSSOUserByEmailAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSSOUserByEmailAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUser

`func (o *GetSSOUserByEmailAPIResponse) GetUser() APISSOUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSSOUserByEmailAPIResponse) GetUserOk() (*APISSOUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSSOUserByEmailAPIResponse) SetUser(v APISSOUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetSSOUserByEmailAPIResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetStatus

`func (o *GetSSOUserByEmailAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSSOUserByEmailAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSSOUserByEmailAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


