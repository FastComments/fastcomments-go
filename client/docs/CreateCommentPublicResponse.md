# CreateCommentPublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Comment** | [**PublicComment**](PublicComment.md) |  | 
**User** | [**NullableUserSessionInfo**](UserSessionInfo.md) |  | 
**ModuleData** | Pointer to **map[string]map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**UserIdWS** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewCreateCommentPublicResponse

`func NewCreateCommentPublicResponse(status APIStatus, comment PublicComment, user NullableUserSessionInfo, reason string, code string, ) *CreateCommentPublicResponse`

NewCreateCommentPublicResponse instantiates a new CreateCommentPublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentPublicResponseWithDefaults

`func NewCreateCommentPublicResponseWithDefaults() *CreateCommentPublicResponse`

NewCreateCommentPublicResponseWithDefaults instantiates a new CreateCommentPublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateCommentPublicResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCommentPublicResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCommentPublicResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *CreateCommentPublicResponse) GetComment() PublicComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateCommentPublicResponse) GetCommentOk() (*PublicComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateCommentPublicResponse) SetComment(v PublicComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *CreateCommentPublicResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateCommentPublicResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateCommentPublicResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *CreateCommentPublicResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CreateCommentPublicResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *CreateCommentPublicResponse) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *CreateCommentPublicResponse) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *CreateCommentPublicResponse) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *CreateCommentPublicResponse) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetUserIdWS

`func (o *CreateCommentPublicResponse) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *CreateCommentPublicResponse) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *CreateCommentPublicResponse) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *CreateCommentPublicResponse) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetReason

`func (o *CreateCommentPublicResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCommentPublicResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCommentPublicResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CreateCommentPublicResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCommentPublicResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCommentPublicResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CreateCommentPublicResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CreateCommentPublicResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CreateCommentPublicResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CreateCommentPublicResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CreateCommentPublicResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CreateCommentPublicResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CreateCommentPublicResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CreateCommentPublicResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CreateCommentPublicResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CreateCommentPublicResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CreateCommentPublicResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CreateCommentPublicResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CreateCommentPublicResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CreateCommentPublicResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CreateCommentPublicResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CreateCommentPublicResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CreateCommentPublicResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CreateCommentPublicResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CreateCommentPublicResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CreateCommentPublicResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


