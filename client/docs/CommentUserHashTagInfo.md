# CommentUserHashTagInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Tag** | **string** |  | 
**Url** | **NullableString** |  | 
**Retain** | Pointer to **bool** |  | [optional] 

## Methods

### NewCommentUserHashTagInfo

`func NewCommentUserHashTagInfo(id string, tag string, url NullableString, ) *CommentUserHashTagInfo`

NewCommentUserHashTagInfo instantiates a new CommentUserHashTagInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentUserHashTagInfoWithDefaults

`func NewCommentUserHashTagInfoWithDefaults() *CommentUserHashTagInfo`

NewCommentUserHashTagInfoWithDefaults instantiates a new CommentUserHashTagInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentUserHashTagInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentUserHashTagInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentUserHashTagInfo) SetId(v string)`

SetId sets Id field to given value.


### GetTag

`func (o *CommentUserHashTagInfo) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CommentUserHashTagInfo) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CommentUserHashTagInfo) SetTag(v string)`

SetTag sets Tag field to given value.


### GetUrl

`func (o *CommentUserHashTagInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommentUserHashTagInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommentUserHashTagInfo) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *CommentUserHashTagInfo) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CommentUserHashTagInfo) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRetain

`func (o *CommentUserHashTagInfo) GetRetain() bool`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *CommentUserHashTagInfo) GetRetainOk() (*bool, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *CommentUserHashTagInfo) SetRetain(v bool)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *CommentUserHashTagInfo) HasRetain() bool`

HasRetain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


