# GetVotesResponse1

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

### NewGetVotesResponse1

`func NewGetVotesResponse1(status APIStatus, appliedAuthorizedVotes []PublicVote, appliedAnonymousVotes []PublicVote, pendingVotes []PublicVote, reason string, code string, ) *GetVotesResponse1`

NewGetVotesResponse1 instantiates a new GetVotesResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVotesResponse1WithDefaults

`func NewGetVotesResponse1WithDefaults() *GetVotesResponse1`

NewGetVotesResponse1WithDefaults instantiates a new GetVotesResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetVotesResponse1) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVotesResponse1) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVotesResponse1) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetAppliedAuthorizedVotes

`func (o *GetVotesResponse1) GetAppliedAuthorizedVotes() []PublicVote`

GetAppliedAuthorizedVotes returns the AppliedAuthorizedVotes field if non-nil, zero value otherwise.

### GetAppliedAuthorizedVotesOk

`func (o *GetVotesResponse1) GetAppliedAuthorizedVotesOk() (*[]PublicVote, bool)`

GetAppliedAuthorizedVotesOk returns a tuple with the AppliedAuthorizedVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAuthorizedVotes

`func (o *GetVotesResponse1) SetAppliedAuthorizedVotes(v []PublicVote)`

SetAppliedAuthorizedVotes sets AppliedAuthorizedVotes field to given value.


### GetAppliedAnonymousVotes

`func (o *GetVotesResponse1) GetAppliedAnonymousVotes() []PublicVote`

GetAppliedAnonymousVotes returns the AppliedAnonymousVotes field if non-nil, zero value otherwise.

### GetAppliedAnonymousVotesOk

`func (o *GetVotesResponse1) GetAppliedAnonymousVotesOk() (*[]PublicVote, bool)`

GetAppliedAnonymousVotesOk returns a tuple with the AppliedAnonymousVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAnonymousVotes

`func (o *GetVotesResponse1) SetAppliedAnonymousVotes(v []PublicVote)`

SetAppliedAnonymousVotes sets AppliedAnonymousVotes field to given value.


### GetPendingVotes

`func (o *GetVotesResponse1) GetPendingVotes() []PublicVote`

GetPendingVotes returns the PendingVotes field if non-nil, zero value otherwise.

### GetPendingVotesOk

`func (o *GetVotesResponse1) GetPendingVotesOk() (*[]PublicVote, bool)`

GetPendingVotesOk returns a tuple with the PendingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingVotes

`func (o *GetVotesResponse1) SetPendingVotes(v []PublicVote)`

SetPendingVotes sets PendingVotes field to given value.


### GetReason

`func (o *GetVotesResponse1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetVotesResponse1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetVotesResponse1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetVotesResponse1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetVotesResponse1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetVotesResponse1) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetVotesResponse1) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetVotesResponse1) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetVotesResponse1) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetVotesResponse1) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetVotesResponse1) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetVotesResponse1) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetVotesResponse1) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetVotesResponse1) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetVotesResponse1) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetVotesResponse1) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetVotesResponse1) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetVotesResponse1) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetVotesResponse1) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetVotesResponse1) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetVotesResponse1) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetVotesResponse1) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetVotesResponse1) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetVotesResponse1) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetVotesResponse1) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetVotesResponse1) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


