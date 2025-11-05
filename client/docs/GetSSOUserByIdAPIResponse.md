# GetSSOUserByIdAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**APISSOUser**](APISSOUser.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewGetSSOUserByIdAPIResponse

`func NewGetSSOUserByIdAPIResponse(status string, ) *GetSSOUserByIdAPIResponse`

NewGetSSOUserByIdAPIResponse instantiates a new GetSSOUserByIdAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSSOUserByIdAPIResponseWithDefaults

`func NewGetSSOUserByIdAPIResponseWithDefaults() *GetSSOUserByIdAPIResponse`

NewGetSSOUserByIdAPIResponseWithDefaults instantiates a new GetSSOUserByIdAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *GetSSOUserByIdAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetSSOUserByIdAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetSSOUserByIdAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetSSOUserByIdAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *GetSSOUserByIdAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSSOUserByIdAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSSOUserByIdAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSSOUserByIdAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUser

`func (o *GetSSOUserByIdAPIResponse) GetUser() APISSOUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSSOUserByIdAPIResponse) GetUserOk() (*APISSOUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSSOUserByIdAPIResponse) SetUser(v APISSOUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetSSOUserByIdAPIResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetStatus

`func (o *GetSSOUserByIdAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSSOUserByIdAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSSOUserByIdAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


