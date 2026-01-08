# GetVotes200Response

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

### NewGetVotes200Response

`func NewGetVotes200Response(status APIStatus, appliedAuthorizedVotes []PublicVote, appliedAnonymousVotes []PublicVote, pendingVotes []PublicVote, reason string, code string, ) *GetVotes200Response`

NewGetVotes200Response instantiates a new GetVotes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVotes200ResponseWithDefaults

`func NewGetVotes200ResponseWithDefaults() *GetVotes200Response`

NewGetVotes200ResponseWithDefaults instantiates a new GetVotes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetVotes200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVotes200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVotes200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetAppliedAuthorizedVotes

`func (o *GetVotes200Response) GetAppliedAuthorizedVotes() []PublicVote`

GetAppliedAuthorizedVotes returns the AppliedAuthorizedVotes field if non-nil, zero value otherwise.

### GetAppliedAuthorizedVotesOk

`func (o *GetVotes200Response) GetAppliedAuthorizedVotesOk() (*[]PublicVote, bool)`

GetAppliedAuthorizedVotesOk returns a tuple with the AppliedAuthorizedVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAuthorizedVotes

`func (o *GetVotes200Response) SetAppliedAuthorizedVotes(v []PublicVote)`

SetAppliedAuthorizedVotes sets AppliedAuthorizedVotes field to given value.


### GetAppliedAnonymousVotes

`func (o *GetVotes200Response) GetAppliedAnonymousVotes() []PublicVote`

GetAppliedAnonymousVotes returns the AppliedAnonymousVotes field if non-nil, zero value otherwise.

### GetAppliedAnonymousVotesOk

`func (o *GetVotes200Response) GetAppliedAnonymousVotesOk() (*[]PublicVote, bool)`

GetAppliedAnonymousVotesOk returns a tuple with the AppliedAnonymousVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAnonymousVotes

`func (o *GetVotes200Response) SetAppliedAnonymousVotes(v []PublicVote)`

SetAppliedAnonymousVotes sets AppliedAnonymousVotes field to given value.


### GetPendingVotes

`func (o *GetVotes200Response) GetPendingVotes() []PublicVote`

GetPendingVotes returns the PendingVotes field if non-nil, zero value otherwise.

### GetPendingVotesOk

`func (o *GetVotes200Response) GetPendingVotesOk() (*[]PublicVote, bool)`

GetPendingVotesOk returns a tuple with the PendingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingVotes

`func (o *GetVotes200Response) SetPendingVotes(v []PublicVote)`

SetPendingVotes sets PendingVotes field to given value.


### GetReason

`func (o *GetVotes200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetVotes200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetVotes200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetVotes200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetVotes200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetVotes200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetVotes200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetVotes200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetVotes200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetVotes200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetVotes200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetVotes200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetVotes200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetVotes200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetVotes200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetVotes200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetVotes200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetVotes200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetVotes200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetVotes200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetVotes200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetVotes200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetVotes200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetVotes200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetVotes200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetVotes200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


