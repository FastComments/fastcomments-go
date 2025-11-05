# CheckBlockedCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentStatuses** | **map[string]bool** | Construct a type with a set of properties K of type T | 
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 

## Methods

### NewCheckBlockedCommentsResponse

`func NewCheckBlockedCommentsResponse(commentStatuses map[string]bool, status ImportedAPIStatusSUCCESS, ) *CheckBlockedCommentsResponse`

NewCheckBlockedCommentsResponse instantiates a new CheckBlockedCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckBlockedCommentsResponseWithDefaults

`func NewCheckBlockedCommentsResponseWithDefaults() *CheckBlockedCommentsResponse`

NewCheckBlockedCommentsResponseWithDefaults instantiates a new CheckBlockedCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentStatuses

`func (o *CheckBlockedCommentsResponse) GetCommentStatuses() map[string]bool`

GetCommentStatuses returns the CommentStatuses field if non-nil, zero value otherwise.

### GetCommentStatusesOk

`func (o *CheckBlockedCommentsResponse) GetCommentStatusesOk() (*map[string]bool, bool)`

GetCommentStatusesOk returns a tuple with the CommentStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentStatuses

`func (o *CheckBlockedCommentsResponse) SetCommentStatuses(v map[string]bool)`

SetCommentStatuses sets CommentStatuses field to given value.


### GetStatus

`func (o *CheckBlockedCommentsResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckBlockedCommentsResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckBlockedCommentsResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


