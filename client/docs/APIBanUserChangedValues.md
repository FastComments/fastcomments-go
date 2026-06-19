# APIBanUserChangedValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**IpHash** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**BannedByUserId** | Pointer to **string** |  | [optional] 
**BannedCommentText** | Pointer to **string** |  | [optional] 
**BanType** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **NullableTime** |  | [optional] 
**HasEmailWildcard** | Pointer to **bool** |  | [optional] 
**BanReason** | Pointer to **string** |  | [optional] 

## Methods

### NewAPIBanUserChangedValues

`func NewAPIBanUserChangedValues() *APIBanUserChangedValues`

NewAPIBanUserChangedValues instantiates a new APIBanUserChangedValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIBanUserChangedValuesWithDefaults

`func NewAPIBanUserChangedValuesWithDefaults() *APIBanUserChangedValues`

NewAPIBanUserChangedValuesWithDefaults instantiates a new APIBanUserChangedValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIBanUserChangedValues) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIBanUserChangedValues) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIBanUserChangedValues) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIBanUserChangedValues) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *APIBanUserChangedValues) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *APIBanUserChangedValues) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *APIBanUserChangedValues) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *APIBanUserChangedValues) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserId

`func (o *APIBanUserChangedValues) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIBanUserChangedValues) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIBanUserChangedValues) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIBanUserChangedValues) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *APIBanUserChangedValues) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *APIBanUserChangedValues) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEmail

`func (o *APIBanUserChangedValues) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *APIBanUserChangedValues) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *APIBanUserChangedValues) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *APIBanUserChangedValues) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *APIBanUserChangedValues) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *APIBanUserChangedValues) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUsername

`func (o *APIBanUserChangedValues) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *APIBanUserChangedValues) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *APIBanUserChangedValues) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *APIBanUserChangedValues) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *APIBanUserChangedValues) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *APIBanUserChangedValues) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetIpHash

`func (o *APIBanUserChangedValues) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *APIBanUserChangedValues) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *APIBanUserChangedValues) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.

### HasIpHash

`func (o *APIBanUserChangedValues) HasIpHash() bool`

HasIpHash returns a boolean if a field has been set.

### SetIpHashNil

`func (o *APIBanUserChangedValues) SetIpHashNil(b bool)`

 SetIpHashNil sets the value for IpHash to be an explicit nil

### UnsetIpHash
`func (o *APIBanUserChangedValues) UnsetIpHash()`

UnsetIpHash ensures that no value is present for IpHash, not even an explicit nil
### GetCreatedAt

`func (o *APIBanUserChangedValues) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIBanUserChangedValues) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIBanUserChangedValues) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *APIBanUserChangedValues) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetBannedByUserId

`func (o *APIBanUserChangedValues) GetBannedByUserId() string`

GetBannedByUserId returns the BannedByUserId field if non-nil, zero value otherwise.

### GetBannedByUserIdOk

`func (o *APIBanUserChangedValues) GetBannedByUserIdOk() (*string, bool)`

GetBannedByUserIdOk returns a tuple with the BannedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedByUserId

`func (o *APIBanUserChangedValues) SetBannedByUserId(v string)`

SetBannedByUserId sets BannedByUserId field to given value.

### HasBannedByUserId

`func (o *APIBanUserChangedValues) HasBannedByUserId() bool`

HasBannedByUserId returns a boolean if a field has been set.

### GetBannedCommentText

`func (o *APIBanUserChangedValues) GetBannedCommentText() string`

GetBannedCommentText returns the BannedCommentText field if non-nil, zero value otherwise.

### GetBannedCommentTextOk

`func (o *APIBanUserChangedValues) GetBannedCommentTextOk() (*string, bool)`

GetBannedCommentTextOk returns a tuple with the BannedCommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedCommentText

`func (o *APIBanUserChangedValues) SetBannedCommentText(v string)`

SetBannedCommentText sets BannedCommentText field to given value.

### HasBannedCommentText

`func (o *APIBanUserChangedValues) HasBannedCommentText() bool`

HasBannedCommentText returns a boolean if a field has been set.

### GetBanType

`func (o *APIBanUserChangedValues) GetBanType() string`

GetBanType returns the BanType field if non-nil, zero value otherwise.

### GetBanTypeOk

`func (o *APIBanUserChangedValues) GetBanTypeOk() (*string, bool)`

GetBanTypeOk returns a tuple with the BanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanType

`func (o *APIBanUserChangedValues) SetBanType(v string)`

SetBanType sets BanType field to given value.

### HasBanType

`func (o *APIBanUserChangedValues) HasBanType() bool`

HasBanType returns a boolean if a field has been set.

### GetBannedUntil

`func (o *APIBanUserChangedValues) GetBannedUntil() time.Time`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *APIBanUserChangedValues) GetBannedUntilOk() (*time.Time, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *APIBanUserChangedValues) SetBannedUntil(v time.Time)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *APIBanUserChangedValues) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### SetBannedUntilNil

`func (o *APIBanUserChangedValues) SetBannedUntilNil(b bool)`

 SetBannedUntilNil sets the value for BannedUntil to be an explicit nil

### UnsetBannedUntil
`func (o *APIBanUserChangedValues) UnsetBannedUntil()`

UnsetBannedUntil ensures that no value is present for BannedUntil, not even an explicit nil
### GetHasEmailWildcard

`func (o *APIBanUserChangedValues) GetHasEmailWildcard() bool`

GetHasEmailWildcard returns the HasEmailWildcard field if non-nil, zero value otherwise.

### GetHasEmailWildcardOk

`func (o *APIBanUserChangedValues) GetHasEmailWildcardOk() (*bool, bool)`

GetHasEmailWildcardOk returns a tuple with the HasEmailWildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmailWildcard

`func (o *APIBanUserChangedValues) SetHasEmailWildcard(v bool)`

SetHasEmailWildcard sets HasEmailWildcard field to given value.

### HasHasEmailWildcard

`func (o *APIBanUserChangedValues) HasHasEmailWildcard() bool`

HasHasEmailWildcard returns a boolean if a field has been set.

### GetBanReason

`func (o *APIBanUserChangedValues) GetBanReason() string`

GetBanReason returns the BanReason field if non-nil, zero value otherwise.

### GetBanReasonOk

`func (o *APIBanUserChangedValues) GetBanReasonOk() (*string, bool)`

GetBanReasonOk returns a tuple with the BanReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanReason

`func (o *APIBanUserChangedValues) SetBanReason(v string)`

SetBanReason sets BanReason field to given value.

### HasBanReason

`func (o *APIBanUserChangedValues) HasBanReason() bool`

HasBanReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


