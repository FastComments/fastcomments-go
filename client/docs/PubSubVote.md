# PubSubVote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**UrlIdRaw** | **string** |  | 
**CommentId** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Direction** | **int32** |  | 
**CreatedAt** | **int64** |  | 
**VerificationId** | **NullableString** |  | 

## Methods

### NewPubSubVote

`func NewPubSubVote(id string, tenantId string, urlId string, urlIdRaw string, commentId string, direction int32, createdAt int64, verificationId NullableString, ) *PubSubVote`

NewPubSubVote instantiates a new PubSubVote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubSubVoteWithDefaults

`func NewPubSubVoteWithDefaults() *PubSubVote`

NewPubSubVoteWithDefaults instantiates a new PubSubVote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PubSubVote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PubSubVote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PubSubVote) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *PubSubVote) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PubSubVote) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PubSubVote) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *PubSubVote) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *PubSubVote) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *PubSubVote) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrlIdRaw

`func (o *PubSubVote) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *PubSubVote) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *PubSubVote) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.


### GetCommentId

`func (o *PubSubVote) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *PubSubVote) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *PubSubVote) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.


### GetUserId

`func (o *PubSubVote) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PubSubVote) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PubSubVote) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PubSubVote) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *PubSubVote) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *PubSubVote) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDirection

`func (o *PubSubVote) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PubSubVote) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PubSubVote) SetDirection(v int32)`

SetDirection sets Direction field to given value.


### GetCreatedAt

`func (o *PubSubVote) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PubSubVote) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PubSubVote) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetVerificationId

`func (o *PubSubVote) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *PubSubVote) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *PubSubVote) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### SetVerificationIdNil

`func (o *PubSubVote) SetVerificationIdNil(b bool)`

 SetVerificationIdNil sets the value for VerificationId to be an explicit nil

### UnsetVerificationId
`func (o *PubSubVote) UnsetVerificationId()`

UnsetVerificationId ensures that no value is present for VerificationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


