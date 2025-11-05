# DeleteCommentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**DeleteCommentAction**](DeleteCommentAction.md) |  | 
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 

## Methods

### NewDeleteCommentResult

`func NewDeleteCommentResult(action DeleteCommentAction, status ImportedAPIStatusSUCCESS, ) *DeleteCommentResult`

NewDeleteCommentResult instantiates a new DeleteCommentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCommentResultWithDefaults

`func NewDeleteCommentResultWithDefaults() *DeleteCommentResult`

NewDeleteCommentResultWithDefaults instantiates a new DeleteCommentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeleteCommentResult) GetAction() DeleteCommentAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteCommentResult) GetActionOk() (*DeleteCommentAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteCommentResult) SetAction(v DeleteCommentAction)`

SetAction sets Action field to given value.


### GetStatus

`func (o *DeleteCommentResult) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteCommentResult) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteCommentResult) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


