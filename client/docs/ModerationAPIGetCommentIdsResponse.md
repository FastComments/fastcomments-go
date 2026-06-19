# ModerationAPIGetCommentIdsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
**HasMore** | **bool** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationAPIGetCommentIdsResponse

`func NewModerationAPIGetCommentIdsResponse(ids []string, hasMore bool, status APIStatus, ) *ModerationAPIGetCommentIdsResponse`

NewModerationAPIGetCommentIdsResponse instantiates a new ModerationAPIGetCommentIdsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPIGetCommentIdsResponseWithDefaults

`func NewModerationAPIGetCommentIdsResponseWithDefaults() *ModerationAPIGetCommentIdsResponse`

NewModerationAPIGetCommentIdsResponseWithDefaults instantiates a new ModerationAPIGetCommentIdsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ModerationAPIGetCommentIdsResponse) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ModerationAPIGetCommentIdsResponse) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ModerationAPIGetCommentIdsResponse) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetHasMore

`func (o *ModerationAPIGetCommentIdsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ModerationAPIGetCommentIdsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ModerationAPIGetCommentIdsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetStatus

`func (o *ModerationAPIGetCommentIdsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationAPIGetCommentIdsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationAPIGetCommentIdsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


