# APIBannedUserWithMultiMatchInfo

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
**Matches** | [**[]BannedUserMatch**](BannedUserMatch.md) |  | 

## Methods

### NewAPIBannedUserWithMultiMatchInfo

`func NewAPIBannedUserWithMultiMatchInfo(id string, banType string, bannedUntil NullableTime, hasEmailWildcard bool, matches []BannedUserMatch, ) *APIBannedUserWithMultiMatchInfo`

NewAPIBannedUserWithMultiMatchInfo instantiates a new APIBannedUserWithMultiMatchInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIBannedUserWithMultiMatchInfoWithDefaults

`func NewAPIBannedUserWithMultiMatchInfoWithDefaults() *APIBannedUserWithMultiMatchInfo`

NewAPIBannedUserWithMultiMatchInfoWithDefaults instantiates a new APIBannedUserWithMultiMatchInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIBannedUserWithMultiMatchInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIBannedUserWithMultiMatchInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIBannedUserWithMultiMatchInfo) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *APIBannedUserWithMultiMatchInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIBannedUserWithMultiMatchInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIBannedUserWithMultiMatchInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIBannedUserWithMultiMatchInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *APIBannedUserWithMultiMatchInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *APIBannedUserWithMultiMatchInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetBanType

`func (o *APIBannedUserWithMultiMatchInfo) GetBanType() string`

GetBanType returns the BanType field if non-nil, zero value otherwise.

### GetBanTypeOk

`func (o *APIBannedUserWithMultiMatchInfo) GetBanTypeOk() (*string, bool)`

GetBanTypeOk returns a tuple with the BanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanType

`func (o *APIBannedUserWithMultiMatchInfo) SetBanType(v string)`

SetBanType sets BanType field to given value.


### GetEmail

`func (o *APIBannedUserWithMultiMatchInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *APIBannedUserWithMultiMatchInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *APIBannedUserWithMultiMatchInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *APIBannedUserWithMultiMatchInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *APIBannedUserWithMultiMatchInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *APIBannedUserWithMultiMatchInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIpHash

`func (o *APIBannedUserWithMultiMatchInfo) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *APIBannedUserWithMultiMatchInfo) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *APIBannedUserWithMultiMatchInfo) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.

### HasIpHash

`func (o *APIBannedUserWithMultiMatchInfo) HasIpHash() bool`

HasIpHash returns a boolean if a field has been set.

### SetIpHashNil

`func (o *APIBannedUserWithMultiMatchInfo) SetIpHashNil(b bool)`

 SetIpHashNil sets the value for IpHash to be an explicit nil

### UnsetIpHash
`func (o *APIBannedUserWithMultiMatchInfo) UnsetIpHash()`

UnsetIpHash ensures that no value is present for IpHash, not even an explicit nil
### GetBannedUntil

`func (o *APIBannedUserWithMultiMatchInfo) GetBannedUntil() time.Time`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *APIBannedUserWithMultiMatchInfo) GetBannedUntilOk() (*time.Time, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *APIBannedUserWithMultiMatchInfo) SetBannedUntil(v time.Time)`

SetBannedUntil sets BannedUntil field to given value.


### SetBannedUntilNil

`func (o *APIBannedUserWithMultiMatchInfo) SetBannedUntilNil(b bool)`

 SetBannedUntilNil sets the value for BannedUntil to be an explicit nil

### UnsetBannedUntil
`func (o *APIBannedUserWithMultiMatchInfo) UnsetBannedUntil()`

UnsetBannedUntil ensures that no value is present for BannedUntil, not even an explicit nil
### GetHasEmailWildcard

`func (o *APIBannedUserWithMultiMatchInfo) GetHasEmailWildcard() bool`

GetHasEmailWildcard returns the HasEmailWildcard field if non-nil, zero value otherwise.

### GetHasEmailWildcardOk

`func (o *APIBannedUserWithMultiMatchInfo) GetHasEmailWildcardOk() (*bool, bool)`

GetHasEmailWildcardOk returns a tuple with the HasEmailWildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmailWildcard

`func (o *APIBannedUserWithMultiMatchInfo) SetHasEmailWildcard(v bool)`

SetHasEmailWildcard sets HasEmailWildcard field to given value.


### GetBanReason

`func (o *APIBannedUserWithMultiMatchInfo) GetBanReason() string`

GetBanReason returns the BanReason field if non-nil, zero value otherwise.

### GetBanReasonOk

`func (o *APIBannedUserWithMultiMatchInfo) GetBanReasonOk() (*string, bool)`

GetBanReasonOk returns a tuple with the BanReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanReason

`func (o *APIBannedUserWithMultiMatchInfo) SetBanReason(v string)`

SetBanReason sets BanReason field to given value.

### HasBanReason

`func (o *APIBannedUserWithMultiMatchInfo) HasBanReason() bool`

HasBanReason returns a boolean if a field has been set.

### GetMatches

`func (o *APIBannedUserWithMultiMatchInfo) GetMatches() []BannedUserMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *APIBannedUserWithMultiMatchInfo) GetMatchesOk() (*[]BannedUserMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *APIBannedUserWithMultiMatchInfo) SetMatches(v []BannedUserMatch)`

SetMatches sets Matches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


