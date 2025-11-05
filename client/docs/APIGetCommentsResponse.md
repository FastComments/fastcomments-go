# APIGetCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Comments** | [**[]PickFCommentAPICommentFieldsKeys**](PickFCommentAPICommentFieldsKeys.md) |  | 

## Methods

### NewAPIGetCommentsResponse

`func NewAPIGetCommentsResponse(status APIStatus, comments []PickFCommentAPICommentFieldsKeys, ) *APIGetCommentsResponse`

NewAPIGetCommentsResponse instantiates a new APIGetCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIGetCommentsResponseWithDefaults

`func NewAPIGetCommentsResponseWithDefaults() *APIGetCommentsResponse`

NewAPIGetCommentsResponseWithDefaults instantiates a new APIGetCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIGetCommentsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIGetCommentsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIGetCommentsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComments

`func (o *APIGetCommentsResponse) GetComments() []PickFCommentAPICommentFieldsKeys`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *APIGetCommentsResponse) GetCommentsOk() (*[]PickFCommentAPICommentFieldsKeys, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *APIGetCommentsResponse) SetComments(v []PickFCommentAPICommentFieldsKeys)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


