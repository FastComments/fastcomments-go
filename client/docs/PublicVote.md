# PublicVote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UrlId** | **string** |  | 
**CommentId** | **string** |  | 
**UserId** | **string** |  | 
**Direction** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPublicVote

`func NewPublicVote(id string, urlId string, commentId string, userId string, direction string, createdAt time.Time, ) *PublicVote`

NewPublicVote instantiates a new PublicVote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicVoteWithDefaults

`func NewPublicVoteWithDefaults() *PublicVote`

NewPublicVoteWithDefaults instantiates a new PublicVote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicVote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicVote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicVote) SetId(v string)`

SetId sets Id field to given value.


### GetUrlId

`func (o *PublicVote) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *PublicVote) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *PublicVote) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetCommentId

`func (o *PublicVote) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *PublicVote) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *PublicVote) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.


### GetUserId

`func (o *PublicVote) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicVote) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicVote) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetDirection

`func (o *PublicVote) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PublicVote) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PublicVote) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetCreatedAt

`func (o *PublicVote) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicVote) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicVote) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


