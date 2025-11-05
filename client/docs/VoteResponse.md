# VoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VoteResponseStatus**](VoteResponseStatus.md) |  | 
**VoteId** | Pointer to **string** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**VoteResponseUser**](VoteResponseUser.md) |  | [optional] 
**EditKey** | Pointer to **string** |  | [optional] 

## Methods

### NewVoteResponse

`func NewVoteResponse(status VoteResponseStatus, ) *VoteResponse`

NewVoteResponse instantiates a new VoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteResponseWithDefaults

`func NewVoteResponseWithDefaults() *VoteResponse`

NewVoteResponseWithDefaults instantiates a new VoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VoteResponse) GetStatus() VoteResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoteResponse) GetStatusOk() (*VoteResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoteResponse) SetStatus(v VoteResponseStatus)`

SetStatus sets Status field to given value.


### GetVoteId

`func (o *VoteResponse) GetVoteId() string`

GetVoteId returns the VoteId field if non-nil, zero value otherwise.

### GetVoteIdOk

`func (o *VoteResponse) GetVoteIdOk() (*string, bool)`

GetVoteIdOk returns a tuple with the VoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteId

`func (o *VoteResponse) SetVoteId(v string)`

SetVoteId sets VoteId field to given value.

### HasVoteId

`func (o *VoteResponse) HasVoteId() bool`

HasVoteId returns a boolean if a field has been set.

### GetIsVerified

`func (o *VoteResponse) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *VoteResponse) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *VoteResponse) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *VoteResponse) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetUser

`func (o *VoteResponse) GetUser() VoteResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VoteResponse) GetUserOk() (*VoteResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VoteResponse) SetUser(v VoteResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *VoteResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEditKey

`func (o *VoteResponse) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *VoteResponse) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *VoteResponse) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *VoteResponse) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


