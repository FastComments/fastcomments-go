# APIGetCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Comment** | [**PickFCommentAPICommentFieldsKeys**](PickFCommentAPICommentFieldsKeys.md) |  | 

## Methods

### NewAPIGetCommentResponse

`func NewAPIGetCommentResponse(status APIStatus, comment PickFCommentAPICommentFieldsKeys, ) *APIGetCommentResponse`

NewAPIGetCommentResponse instantiates a new APIGetCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIGetCommentResponseWithDefaults

`func NewAPIGetCommentResponseWithDefaults() *APIGetCommentResponse`

NewAPIGetCommentResponseWithDefaults instantiates a new APIGetCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIGetCommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIGetCommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIGetCommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *APIGetCommentResponse) GetComment() PickFCommentAPICommentFieldsKeys`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *APIGetCommentResponse) GetCommentOk() (*PickFCommentAPICommentFieldsKeys, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *APIGetCommentResponse) SetComment(v PickFCommentAPICommentFieldsKeys)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


