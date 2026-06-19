# RemoveUserBadgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewRemoveUserBadgeResponse

`func NewRemoveUserBadgeResponse(status APIStatus, ) *RemoveUserBadgeResponse`

NewRemoveUserBadgeResponse instantiates a new RemoveUserBadgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveUserBadgeResponseWithDefaults

`func NewRemoveUserBadgeResponseWithDefaults() *RemoveUserBadgeResponse`

NewRemoveUserBadgeResponseWithDefaults instantiates a new RemoveUserBadgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *RemoveUserBadgeResponse) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *RemoveUserBadgeResponse) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *RemoveUserBadgeResponse) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *RemoveUserBadgeResponse) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetStatus

`func (o *RemoveUserBadgeResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoveUserBadgeResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoveUserBadgeResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


