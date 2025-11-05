# FeedPostStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reacts** | Pointer to **map[string]int32** |  | [optional] 
**CommentCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewFeedPostStats

`func NewFeedPostStats() *FeedPostStats`

NewFeedPostStats instantiates a new FeedPostStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedPostStatsWithDefaults

`func NewFeedPostStatsWithDefaults() *FeedPostStats`

NewFeedPostStatsWithDefaults instantiates a new FeedPostStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReacts

`func (o *FeedPostStats) GetReacts() map[string]int32`

GetReacts returns the Reacts field if non-nil, zero value otherwise.

### GetReactsOk

`func (o *FeedPostStats) GetReactsOk() (*map[string]int32, bool)`

GetReactsOk returns a tuple with the Reacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReacts

`func (o *FeedPostStats) SetReacts(v map[string]int32)`

SetReacts sets Reacts field to given value.

### HasReacts

`func (o *FeedPostStats) HasReacts() bool`

HasReacts returns a boolean if a field has been set.

### GetCommentCount

`func (o *FeedPostStats) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *FeedPostStats) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *FeedPostStats) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *FeedPostStats) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


