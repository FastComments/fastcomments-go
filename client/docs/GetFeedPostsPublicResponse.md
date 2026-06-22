# GetFeedPostsPublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MyReacts** | Pointer to **map[string]map[string]bool** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**FeedPosts** | [**[]FeedPost**](FeedPost.md) |  | 
**User** | Pointer to [**NullableUserSessionInfo**](UserSessionInfo.md) |  | [optional] 
**UrlIdWS** | Pointer to **string** |  | [optional] 
**UserIdWS** | Pointer to **string** |  | [optional] 
**TenantIdWS** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetFeedPostsPublicResponse

`func NewGetFeedPostsPublicResponse(status APIStatus, feedPosts []FeedPost, reason string, code string, ) *GetFeedPostsPublicResponse`

NewGetFeedPostsPublicResponse instantiates a new GetFeedPostsPublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeedPostsPublicResponseWithDefaults

`func NewGetFeedPostsPublicResponseWithDefaults() *GetFeedPostsPublicResponse`

NewGetFeedPostsPublicResponseWithDefaults instantiates a new GetFeedPostsPublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMyReacts

`func (o *GetFeedPostsPublicResponse) GetMyReacts() map[string]map[string]bool`

GetMyReacts returns the MyReacts field if non-nil, zero value otherwise.

### GetMyReactsOk

`func (o *GetFeedPostsPublicResponse) GetMyReactsOk() (*map[string]map[string]bool, bool)`

GetMyReactsOk returns a tuple with the MyReacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyReacts

`func (o *GetFeedPostsPublicResponse) SetMyReacts(v map[string]map[string]bool)`

SetMyReacts sets MyReacts field to given value.

### HasMyReacts

`func (o *GetFeedPostsPublicResponse) HasMyReacts() bool`

HasMyReacts returns a boolean if a field has been set.

### GetStatus

`func (o *GetFeedPostsPublicResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFeedPostsPublicResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFeedPostsPublicResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetFeedPosts

`func (o *GetFeedPostsPublicResponse) GetFeedPosts() []FeedPost`

GetFeedPosts returns the FeedPosts field if non-nil, zero value otherwise.

### GetFeedPostsOk

`func (o *GetFeedPostsPublicResponse) GetFeedPostsOk() (*[]FeedPost, bool)`

GetFeedPostsOk returns a tuple with the FeedPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPosts

`func (o *GetFeedPostsPublicResponse) SetFeedPosts(v []FeedPost)`

SetFeedPosts sets FeedPosts field to given value.


### GetUser

`func (o *GetFeedPostsPublicResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetFeedPostsPublicResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetFeedPostsPublicResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *GetFeedPostsPublicResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GetFeedPostsPublicResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetFeedPostsPublicResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdWS

`func (o *GetFeedPostsPublicResponse) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetFeedPostsPublicResponse) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetFeedPostsPublicResponse) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.

### HasUrlIdWS

`func (o *GetFeedPostsPublicResponse) HasUrlIdWS() bool`

HasUrlIdWS returns a boolean if a field has been set.

### GetUserIdWS

`func (o *GetFeedPostsPublicResponse) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *GetFeedPostsPublicResponse) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *GetFeedPostsPublicResponse) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *GetFeedPostsPublicResponse) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetTenantIdWS

`func (o *GetFeedPostsPublicResponse) GetTenantIdWS() string`

GetTenantIdWS returns the TenantIdWS field if non-nil, zero value otherwise.

### GetTenantIdWSOk

`func (o *GetFeedPostsPublicResponse) GetTenantIdWSOk() (*string, bool)`

GetTenantIdWSOk returns a tuple with the TenantIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdWS

`func (o *GetFeedPostsPublicResponse) SetTenantIdWS(v string)`

SetTenantIdWS sets TenantIdWS field to given value.

### HasTenantIdWS

`func (o *GetFeedPostsPublicResponse) HasTenantIdWS() bool`

HasTenantIdWS returns a boolean if a field has been set.

### GetReason

`func (o *GetFeedPostsPublicResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetFeedPostsPublicResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetFeedPostsPublicResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetFeedPostsPublicResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFeedPostsPublicResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFeedPostsPublicResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetFeedPostsPublicResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetFeedPostsPublicResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetFeedPostsPublicResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetFeedPostsPublicResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetFeedPostsPublicResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetFeedPostsPublicResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetFeedPostsPublicResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetFeedPostsPublicResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetFeedPostsPublicResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetFeedPostsPublicResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetFeedPostsPublicResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetFeedPostsPublicResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetFeedPostsPublicResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetFeedPostsPublicResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetFeedPostsPublicResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetFeedPostsPublicResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetFeedPostsPublicResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetFeedPostsPublicResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetFeedPostsPublicResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetFeedPostsPublicResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


