# CreateCommentPublic200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusFAILED**](ImportedAPIStatusFAILED.md) |  | 
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

### NewCreateCommentPublic200Response

`func NewCreateCommentPublic200Response(status ImportedAPIStatusFAILED, comment PublicComment, user NullableUserSessionInfo, reason string, code string, ) *CreateCommentPublic200Response`

NewCreateCommentPublic200Response instantiates a new CreateCommentPublic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentPublic200ResponseWithDefaults

`func NewCreateCommentPublic200ResponseWithDefaults() *CreateCommentPublic200Response`

NewCreateCommentPublic200ResponseWithDefaults instantiates a new CreateCommentPublic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateCommentPublic200Response) GetStatus() ImportedAPIStatusFAILED`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCommentPublic200Response) GetStatusOk() (*ImportedAPIStatusFAILED, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCommentPublic200Response) SetStatus(v ImportedAPIStatusFAILED)`

SetStatus sets Status field to given value.


### GetComment

`func (o *CreateCommentPublic200Response) GetComment() PublicComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateCommentPublic200Response) GetCommentOk() (*PublicComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateCommentPublic200Response) SetComment(v PublicComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *CreateCommentPublic200Response) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateCommentPublic200Response) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateCommentPublic200Response) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *CreateCommentPublic200Response) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CreateCommentPublic200Response) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *CreateCommentPublic200Response) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *CreateCommentPublic200Response) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *CreateCommentPublic200Response) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *CreateCommentPublic200Response) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetUserIdWS

`func (o *CreateCommentPublic200Response) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *CreateCommentPublic200Response) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *CreateCommentPublic200Response) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *CreateCommentPublic200Response) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetReason

`func (o *CreateCommentPublic200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCommentPublic200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCommentPublic200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *CreateCommentPublic200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCommentPublic200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCommentPublic200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *CreateCommentPublic200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *CreateCommentPublic200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *CreateCommentPublic200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *CreateCommentPublic200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *CreateCommentPublic200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *CreateCommentPublic200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *CreateCommentPublic200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *CreateCommentPublic200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *CreateCommentPublic200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *CreateCommentPublic200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *CreateCommentPublic200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *CreateCommentPublic200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *CreateCommentPublic200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *CreateCommentPublic200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *CreateCommentPublic200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *CreateCommentPublic200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *CreateCommentPublic200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *CreateCommentPublic200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *CreateCommentPublic200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *CreateCommentPublic200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


