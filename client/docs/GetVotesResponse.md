# GetVotesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**AppliedAuthorizedVotes** | [**[]PublicVote**](PublicVote.md) |  | 
**AppliedAnonymousVotes** | [**[]PublicVote**](PublicVote.md) |  | 
**PendingVotes** | [**[]PublicVote**](PublicVote.md) |  | 

## Methods

### NewGetVotesResponse

`func NewGetVotesResponse(status APIStatus, appliedAuthorizedVotes []PublicVote, appliedAnonymousVotes []PublicVote, pendingVotes []PublicVote, ) *GetVotesResponse`

NewGetVotesResponse instantiates a new GetVotesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVotesResponseWithDefaults

`func NewGetVotesResponseWithDefaults() *GetVotesResponse`

NewGetVotesResponseWithDefaults instantiates a new GetVotesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetVotesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVotesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVotesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetAppliedAuthorizedVotes

`func (o *GetVotesResponse) GetAppliedAuthorizedVotes() []PublicVote`

GetAppliedAuthorizedVotes returns the AppliedAuthorizedVotes field if non-nil, zero value otherwise.

### GetAppliedAuthorizedVotesOk

`func (o *GetVotesResponse) GetAppliedAuthorizedVotesOk() (*[]PublicVote, bool)`

GetAppliedAuthorizedVotesOk returns a tuple with the AppliedAuthorizedVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAuthorizedVotes

`func (o *GetVotesResponse) SetAppliedAuthorizedVotes(v []PublicVote)`

SetAppliedAuthorizedVotes sets AppliedAuthorizedVotes field to given value.


### GetAppliedAnonymousVotes

`func (o *GetVotesResponse) GetAppliedAnonymousVotes() []PublicVote`

GetAppliedAnonymousVotes returns the AppliedAnonymousVotes field if non-nil, zero value otherwise.

### GetAppliedAnonymousVotesOk

`func (o *GetVotesResponse) GetAppliedAnonymousVotesOk() (*[]PublicVote, bool)`

GetAppliedAnonymousVotesOk returns a tuple with the AppliedAnonymousVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAnonymousVotes

`func (o *GetVotesResponse) SetAppliedAnonymousVotes(v []PublicVote)`

SetAppliedAnonymousVotes sets AppliedAnonymousVotes field to given value.


### GetPendingVotes

`func (o *GetVotesResponse) GetPendingVotes() []PublicVote`

GetPendingVotes returns the PendingVotes field if non-nil, zero value otherwise.

### GetPendingVotesOk

`func (o *GetVotesResponse) GetPendingVotesOk() (*[]PublicVote, bool)`

GetPendingVotesOk returns a tuple with the PendingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingVotes

`func (o *GetVotesResponse) SetPendingVotes(v []PublicVote)`

SetPendingVotes sets PendingVotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


