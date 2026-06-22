# GetCommentVoteUserNamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**VoteUserNames** | **[]string** |  | 
**HasMore** | **bool** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetCommentVoteUserNamesResponse

`func NewGetCommentVoteUserNamesResponse(status APIStatus, voteUserNames []string, hasMore bool, reason string, code string, ) *GetCommentVoteUserNamesResponse`

NewGetCommentVoteUserNamesResponse instantiates a new GetCommentVoteUserNamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentVoteUserNamesResponseWithDefaults

`func NewGetCommentVoteUserNamesResponseWithDefaults() *GetCommentVoteUserNamesResponse`

NewGetCommentVoteUserNamesResponseWithDefaults instantiates a new GetCommentVoteUserNamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCommentVoteUserNamesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentVoteUserNamesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentVoteUserNamesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetVoteUserNames

`func (o *GetCommentVoteUserNamesResponse) GetVoteUserNames() []string`

GetVoteUserNames returns the VoteUserNames field if non-nil, zero value otherwise.

### GetVoteUserNamesOk

`func (o *GetCommentVoteUserNamesResponse) GetVoteUserNamesOk() (*[]string, bool)`

GetVoteUserNamesOk returns a tuple with the VoteUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteUserNames

`func (o *GetCommentVoteUserNamesResponse) SetVoteUserNames(v []string)`

SetVoteUserNames sets VoteUserNames field to given value.


### GetHasMore

`func (o *GetCommentVoteUserNamesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCommentVoteUserNamesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCommentVoteUserNamesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetReason

`func (o *GetCommentVoteUserNamesResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetCommentVoteUserNamesResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetCommentVoteUserNamesResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetCommentVoteUserNamesResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCommentVoteUserNamesResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCommentVoteUserNamesResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetCommentVoteUserNamesResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetCommentVoteUserNamesResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetCommentVoteUserNamesResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetCommentVoteUserNamesResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetCommentVoteUserNamesResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetCommentVoteUserNamesResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetCommentVoteUserNamesResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetCommentVoteUserNamesResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetCommentVoteUserNamesResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetCommentVoteUserNamesResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetCommentVoteUserNamesResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetCommentVoteUserNamesResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetCommentVoteUserNamesResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetCommentVoteUserNamesResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetCommentVoteUserNamesResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetCommentVoteUserNamesResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetCommentVoteUserNamesResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetCommentVoteUserNamesResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetCommentVoteUserNamesResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetCommentVoteUserNamesResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


