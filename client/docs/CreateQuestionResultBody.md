# CreateQuestionResultBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlId** | **string** |  | 
**Value** | **float64** |  | 
**QuestionId** | **string** |  | 
**AnonUserId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**CommentId** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**[]MetaItem**](MetaItem.md) |  | [optional] 

## Methods

### NewCreateQuestionResultBody

`func NewCreateQuestionResultBody(urlId string, value float64, questionId string, ) *CreateQuestionResultBody`

NewCreateQuestionResultBody instantiates a new CreateQuestionResultBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestionResultBodyWithDefaults

`func NewCreateQuestionResultBodyWithDefaults() *CreateQuestionResultBody`

NewCreateQuestionResultBodyWithDefaults instantiates a new CreateQuestionResultBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlId

`func (o *CreateQuestionResultBody) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *CreateQuestionResultBody) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *CreateQuestionResultBody) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetValue

`func (o *CreateQuestionResultBody) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateQuestionResultBody) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateQuestionResultBody) SetValue(v float64)`

SetValue sets Value field to given value.


### GetQuestionId

`func (o *CreateQuestionResultBody) GetQuestionId() string`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *CreateQuestionResultBody) GetQuestionIdOk() (*string, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *CreateQuestionResultBody) SetQuestionId(v string)`

SetQuestionId sets QuestionId field to given value.


### GetAnonUserId

`func (o *CreateQuestionResultBody) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *CreateQuestionResultBody) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *CreateQuestionResultBody) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *CreateQuestionResultBody) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetUserId

`func (o *CreateQuestionResultBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateQuestionResultBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateQuestionResultBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateQuestionResultBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCommentId

`func (o *CreateQuestionResultBody) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *CreateQuestionResultBody) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *CreateQuestionResultBody) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *CreateQuestionResultBody) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetMeta

`func (o *CreateQuestionResultBody) GetMeta() []MetaItem`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateQuestionResultBody) GetMetaOk() (*[]MetaItem, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateQuestionResultBody) SetMeta(v []MetaItem)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateQuestionResultBody) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CreateQuestionResultBody) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CreateQuestionResultBody) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


