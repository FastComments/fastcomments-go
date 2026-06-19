# ModerationAPIChildCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | [**[]ModerationAPIComment**](ModerationAPIComment.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationAPIChildCommentsResponse

`func NewModerationAPIChildCommentsResponse(comments []ModerationAPIComment, status APIStatus, ) *ModerationAPIChildCommentsResponse`

NewModerationAPIChildCommentsResponse instantiates a new ModerationAPIChildCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPIChildCommentsResponseWithDefaults

`func NewModerationAPIChildCommentsResponseWithDefaults() *ModerationAPIChildCommentsResponse`

NewModerationAPIChildCommentsResponseWithDefaults instantiates a new ModerationAPIChildCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *ModerationAPIChildCommentsResponse) GetComments() []ModerationAPIComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModerationAPIChildCommentsResponse) GetCommentsOk() (*[]ModerationAPIComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModerationAPIChildCommentsResponse) SetComments(v []ModerationAPIComment)`

SetComments sets Comments field to given value.


### GetStatus

`func (o *ModerationAPIChildCommentsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationAPIChildCommentsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationAPIChildCommentsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


