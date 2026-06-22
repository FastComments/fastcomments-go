# GetVotesForUserResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**AppliedAuthorizedVotes** | [**[]PublicVote**](PublicVote.md) |  | 
**AppliedAnonymousVotes** | [**[]PublicVote**](PublicVote.md) |  | 
**PendingVotes** | [**[]PublicVote**](PublicVote.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetVotesForUserResponse1

`func NewGetVotesForUserResponse1(status APIStatus, appliedAuthorizedVotes []PublicVote, appliedAnonymousVotes []PublicVote, pendingVotes []PublicVote, reason string, code string, ) *GetVotesForUserResponse1`

NewGetVotesForUserResponse1 instantiates a new GetVotesForUserResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVotesForUserResponse1WithDefaults

`func NewGetVotesForUserResponse1WithDefaults() *GetVotesForUserResponse1`

NewGetVotesForUserResponse1WithDefaults instantiates a new GetVotesForUserResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetVotesForUserResponse1) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVotesForUserResponse1) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVotesForUserResponse1) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetAppliedAuthorizedVotes

`func (o *GetVotesForUserResponse1) GetAppliedAuthorizedVotes() []PublicVote`

GetAppliedAuthorizedVotes returns the AppliedAuthorizedVotes field if non-nil, zero value otherwise.

### GetAppliedAuthorizedVotesOk

`func (o *GetVotesForUserResponse1) GetAppliedAuthorizedVotesOk() (*[]PublicVote, bool)`

GetAppliedAuthorizedVotesOk returns a tuple with the AppliedAuthorizedVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAuthorizedVotes

`func (o *GetVotesForUserResponse1) SetAppliedAuthorizedVotes(v []PublicVote)`

SetAppliedAuthorizedVotes sets AppliedAuthorizedVotes field to given value.


### GetAppliedAnonymousVotes

`func (o *GetVotesForUserResponse1) GetAppliedAnonymousVotes() []PublicVote`

GetAppliedAnonymousVotes returns the AppliedAnonymousVotes field if non-nil, zero value otherwise.

### GetAppliedAnonymousVotesOk

`func (o *GetVotesForUserResponse1) GetAppliedAnonymousVotesOk() (*[]PublicVote, bool)`

GetAppliedAnonymousVotesOk returns a tuple with the AppliedAnonymousVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAnonymousVotes

`func (o *GetVotesForUserResponse1) SetAppliedAnonymousVotes(v []PublicVote)`

SetAppliedAnonymousVotes sets AppliedAnonymousVotes field to given value.


### GetPendingVotes

`func (o *GetVotesForUserResponse1) GetPendingVotes() []PublicVote`

GetPendingVotes returns the PendingVotes field if non-nil, zero value otherwise.

### GetPendingVotesOk

`func (o *GetVotesForUserResponse1) GetPendingVotesOk() (*[]PublicVote, bool)`

GetPendingVotesOk returns a tuple with the PendingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingVotes

`func (o *GetVotesForUserResponse1) SetPendingVotes(v []PublicVote)`

SetPendingVotes sets PendingVotes field to given value.


### GetReason

`func (o *GetVotesForUserResponse1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetVotesForUserResponse1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetVotesForUserResponse1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetVotesForUserResponse1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetVotesForUserResponse1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetVotesForUserResponse1) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetVotesForUserResponse1) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetVotesForUserResponse1) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetVotesForUserResponse1) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetVotesForUserResponse1) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetVotesForUserResponse1) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetVotesForUserResponse1) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetVotesForUserResponse1) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetVotesForUserResponse1) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetVotesForUserResponse1) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetVotesForUserResponse1) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetVotesForUserResponse1) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetVotesForUserResponse1) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetVotesForUserResponse1) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetVotesForUserResponse1) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetVotesForUserResponse1) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetVotesForUserResponse1) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetVotesForUserResponse1) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetVotesForUserResponse1) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetVotesForUserResponse1) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetVotesForUserResponse1) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


