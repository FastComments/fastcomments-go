# AwardUserBadgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **[]string** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewAwardUserBadgeResponse

`func NewAwardUserBadgeResponse(status APIStatus, ) *AwardUserBadgeResponse`

NewAwardUserBadgeResponse instantiates a new AwardUserBadgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardUserBadgeResponseWithDefaults

`func NewAwardUserBadgeResponseWithDefaults() *AwardUserBadgeResponse`

NewAwardUserBadgeResponseWithDefaults instantiates a new AwardUserBadgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *AwardUserBadgeResponse) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AwardUserBadgeResponse) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AwardUserBadgeResponse) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AwardUserBadgeResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetBadges

`func (o *AwardUserBadgeResponse) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *AwardUserBadgeResponse) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *AwardUserBadgeResponse) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *AwardUserBadgeResponse) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetStatus

`func (o *AwardUserBadgeResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwardUserBadgeResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwardUserBadgeResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


