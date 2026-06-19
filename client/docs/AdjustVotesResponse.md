# AdjustVotesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**NewCommentVotes** | **int32** |  | 

## Methods

### NewAdjustVotesResponse

`func NewAdjustVotesResponse(status string, newCommentVotes int32, ) *AdjustVotesResponse`

NewAdjustVotesResponse instantiates a new AdjustVotesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustVotesResponseWithDefaults

`func NewAdjustVotesResponseWithDefaults() *AdjustVotesResponse`

NewAdjustVotesResponseWithDefaults instantiates a new AdjustVotesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdjustVotesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdjustVotesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdjustVotesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNewCommentVotes

`func (o *AdjustVotesResponse) GetNewCommentVotes() int32`

GetNewCommentVotes returns the NewCommentVotes field if non-nil, zero value otherwise.

### GetNewCommentVotesOk

`func (o *AdjustVotesResponse) GetNewCommentVotesOk() (*int32, bool)`

GetNewCommentVotesOk returns a tuple with the NewCommentVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCommentVotes

`func (o *AdjustVotesResponse) SetNewCommentVotes(v int32)`

SetNewCommentVotes sets NewCommentVotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


