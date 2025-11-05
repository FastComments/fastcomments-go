# CommentTextUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** |  | 
**Mentions** | Pointer to [**[]CommentUserMentionInfo**](CommentUserMentionInfo.md) |  | [optional] 
**HashTags** | Pointer to [**[]CommentUserHashTagInfo**](CommentUserHashTagInfo.md) |  | [optional] 

## Methods

### NewCommentTextUpdateRequest

`func NewCommentTextUpdateRequest(comment string, ) *CommentTextUpdateRequest`

NewCommentTextUpdateRequest instantiates a new CommentTextUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentTextUpdateRequestWithDefaults

`func NewCommentTextUpdateRequestWithDefaults() *CommentTextUpdateRequest`

NewCommentTextUpdateRequestWithDefaults instantiates a new CommentTextUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CommentTextUpdateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CommentTextUpdateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CommentTextUpdateRequest) SetComment(v string)`

SetComment sets Comment field to given value.


### GetMentions

`func (o *CommentTextUpdateRequest) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *CommentTextUpdateRequest) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *CommentTextUpdateRequest) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *CommentTextUpdateRequest) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetHashTags

`func (o *CommentTextUpdateRequest) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *CommentTextUpdateRequest) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *CommentTextUpdateRequest) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *CommentTextUpdateRequest) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


