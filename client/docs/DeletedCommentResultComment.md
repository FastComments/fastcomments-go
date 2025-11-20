# DeletedCommentResultComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeleted** | Pointer to **bool** |  | [optional] 
**CommentHTML** | **string** |  | 
**CommenterName** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeletedCommentResultComment

`func NewDeletedCommentResultComment(commentHTML string, commenterName string, ) *DeletedCommentResultComment`

NewDeletedCommentResultComment instantiates a new DeletedCommentResultComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedCommentResultCommentWithDefaults

`func NewDeletedCommentResultCommentWithDefaults() *DeletedCommentResultComment`

NewDeletedCommentResultCommentWithDefaults instantiates a new DeletedCommentResultComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeleted

`func (o *DeletedCommentResultComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *DeletedCommentResultComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *DeletedCommentResultComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *DeletedCommentResultComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetCommentHTML

`func (o *DeletedCommentResultComment) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *DeletedCommentResultComment) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *DeletedCommentResultComment) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetCommenterName

`func (o *DeletedCommentResultComment) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *DeletedCommentResultComment) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *DeletedCommentResultComment) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetUserId

`func (o *DeletedCommentResultComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeletedCommentResultComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeletedCommentResultComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DeletedCommentResultComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *DeletedCommentResultComment) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *DeletedCommentResultComment) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


