# QuestionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**AnonUserId** | **string** |  | 
**UserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Value** | **int32** |  | 
**CommentId** | Pointer to **NullableString** |  | [optional] 
**QuestionId** | **string** |  | 
**Meta** | Pointer to [**[]MetaItem**](MetaItem.md) |  | [optional] 
**IpHash** | **string** |  | 

## Methods

### NewQuestionResult

`func NewQuestionResult(id string, tenantId string, urlId string, anonUserId string, userId string, createdAt time.Time, value int32, questionId string, ipHash string, ) *QuestionResult`

NewQuestionResult instantiates a new QuestionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionResultWithDefaults

`func NewQuestionResultWithDefaults() *QuestionResult`

NewQuestionResultWithDefaults instantiates a new QuestionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionResult) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuestionResult) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuestionResult) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuestionResult) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *QuestionResult) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *QuestionResult) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *QuestionResult) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetAnonUserId

`func (o *QuestionResult) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *QuestionResult) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *QuestionResult) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.


### GetUserId

`func (o *QuestionResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *QuestionResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *QuestionResult) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *QuestionResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuestionResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuestionResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetValue

`func (o *QuestionResult) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QuestionResult) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QuestionResult) SetValue(v int32)`

SetValue sets Value field to given value.


### GetCommentId

`func (o *QuestionResult) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *QuestionResult) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *QuestionResult) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *QuestionResult) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### SetCommentIdNil

`func (o *QuestionResult) SetCommentIdNil(b bool)`

 SetCommentIdNil sets the value for CommentId to be an explicit nil

### UnsetCommentId
`func (o *QuestionResult) UnsetCommentId()`

UnsetCommentId ensures that no value is present for CommentId, not even an explicit nil
### GetQuestionId

`func (o *QuestionResult) GetQuestionId() string`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *QuestionResult) GetQuestionIdOk() (*string, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *QuestionResult) SetQuestionId(v string)`

SetQuestionId sets QuestionId field to given value.


### GetMeta

`func (o *QuestionResult) GetMeta() []MetaItem`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *QuestionResult) GetMetaOk() (*[]MetaItem, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *QuestionResult) SetMeta(v []MetaItem)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *QuestionResult) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *QuestionResult) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *QuestionResult) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetIpHash

`func (o *QuestionResult) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *QuestionResult) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *QuestionResult) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


