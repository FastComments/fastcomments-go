# GetFeedPostsPublic200Response

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

### NewGetFeedPostsPublic200Response

`func NewGetFeedPostsPublic200Response(status APIStatus, feedPosts []FeedPost, reason string, code string, ) *GetFeedPostsPublic200Response`

NewGetFeedPostsPublic200Response instantiates a new GetFeedPostsPublic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeedPostsPublic200ResponseWithDefaults

`func NewGetFeedPostsPublic200ResponseWithDefaults() *GetFeedPostsPublic200Response`

NewGetFeedPostsPublic200ResponseWithDefaults instantiates a new GetFeedPostsPublic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMyReacts

`func (o *GetFeedPostsPublic200Response) GetMyReacts() map[string]map[string]bool`

GetMyReacts returns the MyReacts field if non-nil, zero value otherwise.

### GetMyReactsOk

`func (o *GetFeedPostsPublic200Response) GetMyReactsOk() (*map[string]map[string]bool, bool)`

GetMyReactsOk returns a tuple with the MyReacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyReacts

`func (o *GetFeedPostsPublic200Response) SetMyReacts(v map[string]map[string]bool)`

SetMyReacts sets MyReacts field to given value.

### HasMyReacts

`func (o *GetFeedPostsPublic200Response) HasMyReacts() bool`

HasMyReacts returns a boolean if a field has been set.

### GetStatus

`func (o *GetFeedPostsPublic200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFeedPostsPublic200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFeedPostsPublic200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetFeedPosts

`func (o *GetFeedPostsPublic200Response) GetFeedPosts() []FeedPost`

GetFeedPosts returns the FeedPosts field if non-nil, zero value otherwise.

### GetFeedPostsOk

`func (o *GetFeedPostsPublic200Response) GetFeedPostsOk() (*[]FeedPost, bool)`

GetFeedPostsOk returns a tuple with the FeedPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPosts

`func (o *GetFeedPostsPublic200Response) SetFeedPosts(v []FeedPost)`

SetFeedPosts sets FeedPosts field to given value.


### GetUser

`func (o *GetFeedPostsPublic200Response) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetFeedPostsPublic200Response) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetFeedPostsPublic200Response) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.

### HasUser

`func (o *GetFeedPostsPublic200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GetFeedPostsPublic200Response) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetFeedPostsPublic200Response) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdWS

`func (o *GetFeedPostsPublic200Response) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetFeedPostsPublic200Response) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetFeedPostsPublic200Response) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.

### HasUrlIdWS

`func (o *GetFeedPostsPublic200Response) HasUrlIdWS() bool`

HasUrlIdWS returns a boolean if a field has been set.

### GetUserIdWS

`func (o *GetFeedPostsPublic200Response) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *GetFeedPostsPublic200Response) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *GetFeedPostsPublic200Response) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *GetFeedPostsPublic200Response) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetTenantIdWS

`func (o *GetFeedPostsPublic200Response) GetTenantIdWS() string`

GetTenantIdWS returns the TenantIdWS field if non-nil, zero value otherwise.

### GetTenantIdWSOk

`func (o *GetFeedPostsPublic200Response) GetTenantIdWSOk() (*string, bool)`

GetTenantIdWSOk returns a tuple with the TenantIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdWS

`func (o *GetFeedPostsPublic200Response) SetTenantIdWS(v string)`

SetTenantIdWS sets TenantIdWS field to given value.

### HasTenantIdWS

`func (o *GetFeedPostsPublic200Response) HasTenantIdWS() bool`

HasTenantIdWS returns a boolean if a field has been set.

### GetReason

`func (o *GetFeedPostsPublic200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetFeedPostsPublic200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetFeedPostsPublic200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetFeedPostsPublic200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFeedPostsPublic200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFeedPostsPublic200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetFeedPostsPublic200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetFeedPostsPublic200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetFeedPostsPublic200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetFeedPostsPublic200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetFeedPostsPublic200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetFeedPostsPublic200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetFeedPostsPublic200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetFeedPostsPublic200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetFeedPostsPublic200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetFeedPostsPublic200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetFeedPostsPublic200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetFeedPostsPublic200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetFeedPostsPublic200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetFeedPostsPublic200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetFeedPostsPublic200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetFeedPostsPublic200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetFeedPostsPublic200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetFeedPostsPublic200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetFeedPostsPublic200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetFeedPostsPublic200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


