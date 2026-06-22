# SaveCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Comment** | [**APIComment**](APIComment.md) |  | 
**User** | [**NullableUserSessionInfo**](UserSessionInfo.md) |  | 
**ModuleData** | Pointer to **map[string]map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewSaveCommentResponse

`func NewSaveCommentResponse(status APIStatus, comment APIComment, user NullableUserSessionInfo, reason string, code string, ) *SaveCommentResponse`

NewSaveCommentResponse instantiates a new SaveCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCommentResponseWithDefaults

`func NewSaveCommentResponseWithDefaults() *SaveCommentResponse`

NewSaveCommentResponseWithDefaults instantiates a new SaveCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SaveCommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaveCommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaveCommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *SaveCommentResponse) GetComment() APIComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SaveCommentResponse) GetCommentOk() (*APIComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SaveCommentResponse) SetComment(v APIComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *SaveCommentResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SaveCommentResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SaveCommentResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SaveCommentResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SaveCommentResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *SaveCommentResponse) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *SaveCommentResponse) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *SaveCommentResponse) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *SaveCommentResponse) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetReason

`func (o *SaveCommentResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SaveCommentResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SaveCommentResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *SaveCommentResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaveCommentResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaveCommentResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *SaveCommentResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *SaveCommentResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *SaveCommentResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *SaveCommentResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *SaveCommentResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *SaveCommentResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *SaveCommentResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *SaveCommentResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *SaveCommentResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *SaveCommentResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *SaveCommentResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *SaveCommentResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *SaveCommentResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *SaveCommentResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *SaveCommentResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *SaveCommentResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *SaveCommentResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *SaveCommentResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *SaveCommentResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *SaveCommentResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


