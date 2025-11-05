# SpamRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** |  | 
**CommentContains** | Pointer to **string** |  | [optional] 

## Methods

### NewSpamRule

`func NewSpamRule(actions []string, ) *SpamRule`

NewSpamRule instantiates a new SpamRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpamRuleWithDefaults

`func NewSpamRuleWithDefaults() *SpamRule`

NewSpamRuleWithDefaults instantiates a new SpamRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *SpamRule) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *SpamRule) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *SpamRule) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetCommentContains

`func (o *SpamRule) GetCommentContains() string`

GetCommentContains returns the CommentContains field if non-nil, zero value otherwise.

### GetCommentContainsOk

`func (o *SpamRule) GetCommentContainsOk() (*string, bool)`

GetCommentContainsOk returns a tuple with the CommentContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentContains

`func (o *SpamRule) SetCommentContains(v string)`

SetCommentContains sets CommentContains field to given value.

### HasCommentContains

`func (o *SpamRule) HasCommentContains() bool`

HasCommentContains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


