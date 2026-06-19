# SaveCommentsBulkResponse

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

### NewSaveCommentsBulkResponse

`func NewSaveCommentsBulkResponse(status APIStatus, comment APIComment, user NullableUserSessionInfo, reason string, code string, ) *SaveCommentsBulkResponse`

NewSaveCommentsBulkResponse instantiates a new SaveCommentsBulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCommentsBulkResponseWithDefaults

`func NewSaveCommentsBulkResponseWithDefaults() *SaveCommentsBulkResponse`

NewSaveCommentsBulkResponseWithDefaults instantiates a new SaveCommentsBulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SaveCommentsBulkResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaveCommentsBulkResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaveCommentsBulkResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *SaveCommentsBulkResponse) GetComment() APIComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SaveCommentsBulkResponse) GetCommentOk() (*APIComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SaveCommentsBulkResponse) SetComment(v APIComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *SaveCommentsBulkResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SaveCommentsBulkResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SaveCommentsBulkResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SaveCommentsBulkResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SaveCommentsBulkResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *SaveCommentsBulkResponse) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *SaveCommentsBulkResponse) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *SaveCommentsBulkResponse) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *SaveCommentsBulkResponse) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetReason

`func (o *SaveCommentsBulkResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SaveCommentsBulkResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SaveCommentsBulkResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *SaveCommentsBulkResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaveCommentsBulkResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaveCommentsBulkResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *SaveCommentsBulkResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *SaveCommentsBulkResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *SaveCommentsBulkResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *SaveCommentsBulkResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *SaveCommentsBulkResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *SaveCommentsBulkResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *SaveCommentsBulkResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *SaveCommentsBulkResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *SaveCommentsBulkResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *SaveCommentsBulkResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *SaveCommentsBulkResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *SaveCommentsBulkResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *SaveCommentsBulkResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *SaveCommentsBulkResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *SaveCommentsBulkResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *SaveCommentsBulkResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *SaveCommentsBulkResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *SaveCommentsBulkResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *SaveCommentsBulkResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *SaveCommentsBulkResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


