# APICommentCommonBannedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**BanType** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**IpHash** | Pointer to **NullableString** |  | [optional] 
**BannedUntil** | **NullableTime** |  | 
**HasEmailWildcard** | **bool** |  | 
**BanReason** | Pointer to **string** |  | [optional] 

## Methods

### NewAPICommentCommonBannedUser

`func NewAPICommentCommonBannedUser(id string, banType string, bannedUntil NullableTime, hasEmailWildcard bool, ) *APICommentCommonBannedUser`

NewAPICommentCommonBannedUser instantiates a new APICommentCommonBannedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICommentCommonBannedUserWithDefaults

`func NewAPICommentCommonBannedUserWithDefaults() *APICommentCommonBannedUser`

NewAPICommentCommonBannedUserWithDefaults instantiates a new APICommentCommonBannedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APICommentCommonBannedUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APICommentCommonBannedUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APICommentCommonBannedUser) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *APICommentCommonBannedUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APICommentCommonBannedUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APICommentCommonBannedUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APICommentCommonBannedUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *APICommentCommonBannedUser) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *APICommentCommonBannedUser) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetBanType

`func (o *APICommentCommonBannedUser) GetBanType() string`

GetBanType returns the BanType field if non-nil, zero value otherwise.

### GetBanTypeOk

`func (o *APICommentCommonBannedUser) GetBanTypeOk() (*string, bool)`

GetBanTypeOk returns a tuple with the BanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanType

`func (o *APICommentCommonBannedUser) SetBanType(v string)`

SetBanType sets BanType field to given value.


### GetEmail

`func (o *APICommentCommonBannedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *APICommentCommonBannedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *APICommentCommonBannedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *APICommentCommonBannedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *APICommentCommonBannedUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *APICommentCommonBannedUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIpHash

`func (o *APICommentCommonBannedUser) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *APICommentCommonBannedUser) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *APICommentCommonBannedUser) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.

### HasIpHash

`func (o *APICommentCommonBannedUser) HasIpHash() bool`

HasIpHash returns a boolean if a field has been set.

### SetIpHashNil

`func (o *APICommentCommonBannedUser) SetIpHashNil(b bool)`

 SetIpHashNil sets the value for IpHash to be an explicit nil

### UnsetIpHash
`func (o *APICommentCommonBannedUser) UnsetIpHash()`

UnsetIpHash ensures that no value is present for IpHash, not even an explicit nil
### GetBannedUntil

`func (o *APICommentCommonBannedUser) GetBannedUntil() time.Time`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *APICommentCommonBannedUser) GetBannedUntilOk() (*time.Time, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *APICommentCommonBannedUser) SetBannedUntil(v time.Time)`

SetBannedUntil sets BannedUntil field to given value.


### SetBannedUntilNil

`func (o *APICommentCommonBannedUser) SetBannedUntilNil(b bool)`

 SetBannedUntilNil sets the value for BannedUntil to be an explicit nil

### UnsetBannedUntil
`func (o *APICommentCommonBannedUser) UnsetBannedUntil()`

UnsetBannedUntil ensures that no value is present for BannedUntil, not even an explicit nil
### GetHasEmailWildcard

`func (o *APICommentCommonBannedUser) GetHasEmailWildcard() bool`

GetHasEmailWildcard returns the HasEmailWildcard field if non-nil, zero value otherwise.

### GetHasEmailWildcardOk

`func (o *APICommentCommonBannedUser) GetHasEmailWildcardOk() (*bool, bool)`

GetHasEmailWildcardOk returns a tuple with the HasEmailWildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmailWildcard

`func (o *APICommentCommonBannedUser) SetHasEmailWildcard(v bool)`

SetHasEmailWildcard sets HasEmailWildcard field to given value.


### GetBanReason

`func (o *APICommentCommonBannedUser) GetBanReason() string`

GetBanReason returns the BanReason field if non-nil, zero value otherwise.

### GetBanReasonOk

`func (o *APICommentCommonBannedUser) GetBanReasonOk() (*string, bool)`

GetBanReasonOk returns a tuple with the BanReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanReason

`func (o *APICommentCommonBannedUser) SetBanReason(v string)`

SetBanReason sets BanReason field to given value.

### HasBanReason

`func (o *APICommentCommonBannedUser) HasBanReason() bool`

HasBanReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


