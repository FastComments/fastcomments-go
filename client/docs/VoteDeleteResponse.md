# VoteDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VoteDeleteResponseStatus**](VoteDeleteResponseStatus.md) |  | 
**WasPendingVote** | Pointer to **bool** |  | [optional] 

## Methods

### NewVoteDeleteResponse

`func NewVoteDeleteResponse(status VoteDeleteResponseStatus, ) *VoteDeleteResponse`

NewVoteDeleteResponse instantiates a new VoteDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteDeleteResponseWithDefaults

`func NewVoteDeleteResponseWithDefaults() *VoteDeleteResponse`

NewVoteDeleteResponseWithDefaults instantiates a new VoteDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VoteDeleteResponse) GetStatus() VoteDeleteResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VoteDeleteResponse) GetStatusOk() (*VoteDeleteResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VoteDeleteResponse) SetStatus(v VoteDeleteResponseStatus)`

SetStatus sets Status field to given value.


### GetWasPendingVote

`func (o *VoteDeleteResponse) GetWasPendingVote() bool`

GetWasPendingVote returns the WasPendingVote field if non-nil, zero value otherwise.

### GetWasPendingVoteOk

`func (o *VoteDeleteResponse) GetWasPendingVoteOk() (*bool, bool)`

GetWasPendingVoteOk returns a tuple with the WasPendingVote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasPendingVote

`func (o *VoteDeleteResponse) SetWasPendingVote(v bool)`

SetWasPendingVote sets WasPendingVote field to given value.

### HasWasPendingVote

`func (o *VoteDeleteResponse) HasWasPendingVote() bool`

HasWasPendingVote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


