# LiveEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LiveEventType**](LiveEventType.md) |  | 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Ts** | Pointer to **int64** |  | [optional] 
**BroadcastId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Notification** | Pointer to [**UserNotification**](UserNotification.md) |  | [optional] 
**Vote** | Pointer to [**PubSubVote**](PubSubVote.md) |  | [optional] 
**Comment** | Pointer to [**PubSubComment**](PubSubComment.md) |  | [optional] 
**FeedPost** | Pointer to [**FeedPost**](FeedPost.md) |  | [optional] 
**ExtraInfo** | Pointer to [**LiveEventExtraInfo**](LiveEventExtraInfo.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**Uj** | Pointer to **[]string** |  | [optional] 
**Ul** | Pointer to **[]string** |  | [optional] 
**Changes** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewLiveEvent

`func NewLiveEvent(type_ LiveEventType, ) *LiveEvent`

NewLiveEvent instantiates a new LiveEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveEventWithDefaults

`func NewLiveEventWithDefaults() *LiveEvent`

NewLiveEventWithDefaults instantiates a new LiveEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LiveEvent) GetType() LiveEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LiveEvent) GetTypeOk() (*LiveEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LiveEvent) SetType(v LiveEventType)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *LiveEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LiveEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LiveEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LiveEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTs

`func (o *LiveEvent) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *LiveEvent) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *LiveEvent) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *LiveEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetBroadcastId

`func (o *LiveEvent) GetBroadcastId() string`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *LiveEvent) GetBroadcastIdOk() (*string, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *LiveEvent) SetBroadcastId(v string)`

SetBroadcastId sets BroadcastId field to given value.

### HasBroadcastId

`func (o *LiveEvent) HasBroadcastId() bool`

HasBroadcastId returns a boolean if a field has been set.

### GetUserId

`func (o *LiveEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LiveEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LiveEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LiveEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetBadges

`func (o *LiveEvent) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *LiveEvent) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *LiveEvent) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *LiveEvent) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetNotification

`func (o *LiveEvent) GetNotification() UserNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *LiveEvent) GetNotificationOk() (*UserNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *LiveEvent) SetNotification(v UserNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *LiveEvent) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetVote

`func (o *LiveEvent) GetVote() PubSubVote`

GetVote returns the Vote field if non-nil, zero value otherwise.

### GetVoteOk

`func (o *LiveEvent) GetVoteOk() (*PubSubVote, bool)`

GetVoteOk returns a tuple with the Vote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVote

`func (o *LiveEvent) SetVote(v PubSubVote)`

SetVote sets Vote field to given value.

### HasVote

`func (o *LiveEvent) HasVote() bool`

HasVote returns a boolean if a field has been set.

### GetComment

`func (o *LiveEvent) GetComment() PubSubComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LiveEvent) GetCommentOk() (*PubSubComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LiveEvent) SetComment(v PubSubComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LiveEvent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFeedPost

`func (o *LiveEvent) GetFeedPost() FeedPost`

GetFeedPost returns the FeedPost field if non-nil, zero value otherwise.

### GetFeedPostOk

`func (o *LiveEvent) GetFeedPostOk() (*FeedPost, bool)`

GetFeedPostOk returns a tuple with the FeedPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPost

`func (o *LiveEvent) SetFeedPost(v FeedPost)`

SetFeedPost sets FeedPost field to given value.

### HasFeedPost

`func (o *LiveEvent) HasFeedPost() bool`

HasFeedPost returns a boolean if a field has been set.

### GetExtraInfo

`func (o *LiveEvent) GetExtraInfo() LiveEventExtraInfo`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *LiveEvent) GetExtraInfoOk() (*LiveEventExtraInfo, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *LiveEvent) SetExtraInfo(v LiveEventExtraInfo)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *LiveEvent) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetConfig

`func (o *LiveEvent) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LiveEvent) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LiveEvent) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LiveEvent) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetIsClosed

`func (o *LiveEvent) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *LiveEvent) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *LiveEvent) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *LiveEvent) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetUj

`func (o *LiveEvent) GetUj() []string`

GetUj returns the Uj field if non-nil, zero value otherwise.

### GetUjOk

`func (o *LiveEvent) GetUjOk() (*[]string, bool)`

GetUjOk returns a tuple with the Uj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUj

`func (o *LiveEvent) SetUj(v []string)`

SetUj sets Uj field to given value.

### HasUj

`func (o *LiveEvent) HasUj() bool`

HasUj returns a boolean if a field has been set.

### GetUl

`func (o *LiveEvent) GetUl() []string`

GetUl returns the Ul field if non-nil, zero value otherwise.

### GetUlOk

`func (o *LiveEvent) GetUlOk() (*[]string, bool)`

GetUlOk returns a tuple with the Ul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUl

`func (o *LiveEvent) SetUl(v []string)`

SetUl sets Ul field to given value.

### HasUl

`func (o *LiveEvent) HasUl() bool`

HasUl returns a boolean if a field has been set.

### GetChanges

`func (o *LiveEvent) GetChanges() map[string]int32`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *LiveEvent) GetChangesOk() (*map[string]int32, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *LiveEvent) SetChanges(v map[string]int32)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *LiveEvent) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


