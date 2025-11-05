# CommentUserMentionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Tag** | **string** |  | 
**RawTag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to **bool** |  | [optional] 

## Methods

### NewCommentUserMentionInfo

`func NewCommentUserMentionInfo(id string, tag string, ) *CommentUserMentionInfo`

NewCommentUserMentionInfo instantiates a new CommentUserMentionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentUserMentionInfoWithDefaults

`func NewCommentUserMentionInfoWithDefaults() *CommentUserMentionInfo`

NewCommentUserMentionInfoWithDefaults instantiates a new CommentUserMentionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentUserMentionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentUserMentionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentUserMentionInfo) SetId(v string)`

SetId sets Id field to given value.


### GetTag

`func (o *CommentUserMentionInfo) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CommentUserMentionInfo) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CommentUserMentionInfo) SetTag(v string)`

SetTag sets Tag field to given value.


### GetRawTag

`func (o *CommentUserMentionInfo) GetRawTag() string`

GetRawTag returns the RawTag field if non-nil, zero value otherwise.

### GetRawTagOk

`func (o *CommentUserMentionInfo) GetRawTagOk() (*string, bool)`

GetRawTagOk returns a tuple with the RawTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTag

`func (o *CommentUserMentionInfo) SetRawTag(v string)`

SetRawTag sets RawTag field to given value.

### HasRawTag

`func (o *CommentUserMentionInfo) HasRawTag() bool`

HasRawTag returns a boolean if a field has been set.

### GetType

`func (o *CommentUserMentionInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommentUserMentionInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommentUserMentionInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommentUserMentionInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSent

`func (o *CommentUserMentionInfo) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *CommentUserMentionInfo) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *CommentUserMentionInfo) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *CommentUserMentionInfo) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


