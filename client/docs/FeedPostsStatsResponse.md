# FeedPostsStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Stats** | [**map[string]FeedPostStats**](FeedPostStats.md) |  | 

## Methods

### NewFeedPostsStatsResponse

`func NewFeedPostsStatsResponse(status APIStatus, stats map[string]FeedPostStats, ) *FeedPostsStatsResponse`

NewFeedPostsStatsResponse instantiates a new FeedPostsStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedPostsStatsResponseWithDefaults

`func NewFeedPostsStatsResponseWithDefaults() *FeedPostsStatsResponse`

NewFeedPostsStatsResponseWithDefaults instantiates a new FeedPostsStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FeedPostsStatsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeedPostsStatsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeedPostsStatsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetStats

`func (o *FeedPostsStatsResponse) GetStats() map[string]FeedPostStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *FeedPostsStatsResponse) GetStatsOk() (*map[string]FeedPostStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *FeedPostsStatsResponse) SetStats(v map[string]FeedPostStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


