# GetBanUsersFromCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannedUsers** | [**[]APIBannedUserWithMultiMatchInfo**](APIBannedUserWithMultiMatchInfo.md) |  | 
**Code** | **string** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetBanUsersFromCommentResponse

`func NewGetBanUsersFromCommentResponse(bannedUsers []APIBannedUserWithMultiMatchInfo, code string, status APIStatus, reason string, ) *GetBanUsersFromCommentResponse`

NewGetBanUsersFromCommentResponse instantiates a new GetBanUsersFromCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBanUsersFromCommentResponseWithDefaults

`func NewGetBanUsersFromCommentResponseWithDefaults() *GetBanUsersFromCommentResponse`

NewGetBanUsersFromCommentResponseWithDefaults instantiates a new GetBanUsersFromCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannedUsers

`func (o *GetBanUsersFromCommentResponse) GetBannedUsers() []APIBannedUserWithMultiMatchInfo`

GetBannedUsers returns the BannedUsers field if non-nil, zero value otherwise.

### GetBannedUsersOk

`func (o *GetBanUsersFromCommentResponse) GetBannedUsersOk() (*[]APIBannedUserWithMultiMatchInfo, bool)`

GetBannedUsersOk returns a tuple with the BannedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUsers

`func (o *GetBanUsersFromCommentResponse) SetBannedUsers(v []APIBannedUserWithMultiMatchInfo)`

SetBannedUsers sets BannedUsers field to given value.


### GetCode

`func (o *GetBanUsersFromCommentResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetBanUsersFromCommentResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetBanUsersFromCommentResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetStatus

`func (o *GetBanUsersFromCommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBanUsersFromCommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBanUsersFromCommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetBanUsersFromCommentResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetBanUsersFromCommentResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetBanUsersFromCommentResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSecondaryCode

`func (o *GetBanUsersFromCommentResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetBanUsersFromCommentResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetBanUsersFromCommentResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetBanUsersFromCommentResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetBanUsersFromCommentResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetBanUsersFromCommentResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetBanUsersFromCommentResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetBanUsersFromCommentResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetBanUsersFromCommentResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetBanUsersFromCommentResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetBanUsersFromCommentResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetBanUsersFromCommentResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetBanUsersFromCommentResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetBanUsersFromCommentResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetBanUsersFromCommentResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetBanUsersFromCommentResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetBanUsersFromCommentResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetBanUsersFromCommentResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetBanUsersFromCommentResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetBanUsersFromCommentResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


