# FindCommentsByRangeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | [**NullableFComment**](FComment.md) |  | 
**Result** | [**QuestionResult**](QuestionResult.md) |  | 

## Methods

### NewFindCommentsByRangeItem

`func NewFindCommentsByRangeItem(comment NullableFComment, result QuestionResult, ) *FindCommentsByRangeItem`

NewFindCommentsByRangeItem instantiates a new FindCommentsByRangeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindCommentsByRangeItemWithDefaults

`func NewFindCommentsByRangeItemWithDefaults() *FindCommentsByRangeItem`

NewFindCommentsByRangeItemWithDefaults instantiates a new FindCommentsByRangeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *FindCommentsByRangeItem) GetComment() FComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FindCommentsByRangeItem) GetCommentOk() (*FComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FindCommentsByRangeItem) SetComment(v FComment)`

SetComment sets Comment field to given value.


### SetCommentNil

`func (o *FindCommentsByRangeItem) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *FindCommentsByRangeItem) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetResult

`func (o *FindCommentsByRangeItem) GetResult() QuestionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FindCommentsByRangeItem) GetResultOk() (*QuestionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FindCommentsByRangeItem) SetResult(v QuestionResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


