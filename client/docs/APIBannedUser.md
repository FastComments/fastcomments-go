# APIBannedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**IpHash** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**BannedByUserId** | **string** |  | 
**BannedCommentText** | **string** |  | 
**BanType** | **string** |  | 
**BannedUntil** | **NullableTime** |  | 
**HasEmailWildcard** | **bool** |  | 
**BanReason** | Pointer to **string** |  | [optional] 

## Methods

### NewAPIBannedUser

`func NewAPIBannedUser(id string, tenantId string, createdAt time.Time, bannedByUserId string, bannedCommentText string, banType string, bannedUntil NullableTime, hasEmailWildcard bool, ) *APIBannedUser`

NewAPIBannedUser instantiates a new APIBannedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIBannedUserWithDefaults

`func NewAPIBannedUserWithDefaults() *APIBannedUser`

NewAPIBannedUserWithDefaults instantiates a new APIBannedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIBannedUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIBannedUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIBannedUser) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *APIBannedUser) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *APIBannedUser) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *APIBannedUser) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUserId

`func (o *APIBannedUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIBannedUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIBannedUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIBannedUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *APIBannedUser) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *APIBannedUser) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEmail

`func (o *APIBannedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *APIBannedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *APIBannedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *APIBannedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *APIBannedUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *APIBannedUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUsername

`func (o *APIBannedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *APIBannedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *APIBannedUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *APIBannedUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *APIBannedUser) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *APIBannedUser) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetIpHash

`func (o *APIBannedUser) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *APIBannedUser) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *APIBannedUser) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.

### HasIpHash

`func (o *APIBannedUser) HasIpHash() bool`

HasIpHash returns a boolean if a field has been set.

### SetIpHashNil

`func (o *APIBannedUser) SetIpHashNil(b bool)`

 SetIpHashNil sets the value for IpHash to be an explicit nil

### UnsetIpHash
`func (o *APIBannedUser) UnsetIpHash()`

UnsetIpHash ensures that no value is present for IpHash, not even an explicit nil
### GetCreatedAt

`func (o *APIBannedUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIBannedUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIBannedUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetBannedByUserId

`func (o *APIBannedUser) GetBannedByUserId() string`

GetBannedByUserId returns the BannedByUserId field if non-nil, zero value otherwise.

### GetBannedByUserIdOk

`func (o *APIBannedUser) GetBannedByUserIdOk() (*string, bool)`

GetBannedByUserIdOk returns a tuple with the BannedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedByUserId

`func (o *APIBannedUser) SetBannedByUserId(v string)`

SetBannedByUserId sets BannedByUserId field to given value.


### GetBannedCommentText

`func (o *APIBannedUser) GetBannedCommentText() string`

GetBannedCommentText returns the BannedCommentText field if non-nil, zero value otherwise.

### GetBannedCommentTextOk

`func (o *APIBannedUser) GetBannedCommentTextOk() (*string, bool)`

GetBannedCommentTextOk returns a tuple with the BannedCommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedCommentText

`func (o *APIBannedUser) SetBannedCommentText(v string)`

SetBannedCommentText sets BannedCommentText field to given value.


### GetBanType

`func (o *APIBannedUser) GetBanType() string`

GetBanType returns the BanType field if non-nil, zero value otherwise.

### GetBanTypeOk

`func (o *APIBannedUser) GetBanTypeOk() (*string, bool)`

GetBanTypeOk returns a tuple with the BanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanType

`func (o *APIBannedUser) SetBanType(v string)`

SetBanType sets BanType field to given value.


### GetBannedUntil

`func (o *APIBannedUser) GetBannedUntil() time.Time`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *APIBannedUser) GetBannedUntilOk() (*time.Time, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *APIBannedUser) SetBannedUntil(v time.Time)`

SetBannedUntil sets BannedUntil field to given value.


### SetBannedUntilNil

`func (o *APIBannedUser) SetBannedUntilNil(b bool)`

 SetBannedUntilNil sets the value for BannedUntil to be an explicit nil

### UnsetBannedUntil
`func (o *APIBannedUser) UnsetBannedUntil()`

UnsetBannedUntil ensures that no value is present for BannedUntil, not even an explicit nil
### GetHasEmailWildcard

`func (o *APIBannedUser) GetHasEmailWildcard() bool`

GetHasEmailWildcard returns the HasEmailWildcard field if non-nil, zero value otherwise.

### GetHasEmailWildcardOk

`func (o *APIBannedUser) GetHasEmailWildcardOk() (*bool, bool)`

GetHasEmailWildcardOk returns a tuple with the HasEmailWildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmailWildcard

`func (o *APIBannedUser) SetHasEmailWildcard(v bool)`

SetHasEmailWildcard sets HasEmailWildcard field to given value.


### GetBanReason

`func (o *APIBannedUser) GetBanReason() string`

GetBanReason returns the BanReason field if non-nil, zero value otherwise.

### GetBanReasonOk

`func (o *APIBannedUser) GetBanReasonOk() (*string, bool)`

GetBanReasonOk returns a tuple with the BanReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanReason

`func (o *APIBannedUser) SetBanReason(v string)`

SetBanReason sets BanReason field to given value.

### HasBanReason

`func (o *APIBannedUser) HasBanReason() bool`

HasBanReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


