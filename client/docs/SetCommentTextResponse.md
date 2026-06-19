# SetCommentTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewCommentTextHTML** | **string** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewSetCommentTextResponse

`func NewSetCommentTextResponse(newCommentTextHTML string, status APIStatus, ) *SetCommentTextResponse`

NewSetCommentTextResponse instantiates a new SetCommentTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCommentTextResponseWithDefaults

`func NewSetCommentTextResponseWithDefaults() *SetCommentTextResponse`

NewSetCommentTextResponseWithDefaults instantiates a new SetCommentTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewCommentTextHTML

`func (o *SetCommentTextResponse) GetNewCommentTextHTML() string`

GetNewCommentTextHTML returns the NewCommentTextHTML field if non-nil, zero value otherwise.

### GetNewCommentTextHTMLOk

`func (o *SetCommentTextResponse) GetNewCommentTextHTMLOk() (*string, bool)`

GetNewCommentTextHTMLOk returns a tuple with the NewCommentTextHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCommentTextHTML

`func (o *SetCommentTextResponse) SetNewCommentTextHTML(v string)`

SetNewCommentTextHTML sets NewCommentTextHTML field to given value.


### GetStatus

`func (o *SetCommentTextResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetCommentTextResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetCommentTextResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


