# ChangeCommentPinStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentPositions** | [**map[string]RecordStringBeforeStringOrNullAfterStringOrNullValue**](RecordStringBeforeStringOrNullAfterStringOrNullValue.md) | Construct a type with a set of properties K of type T | 
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 

## Methods

### NewChangeCommentPinStatusResponse

`func NewChangeCommentPinStatusResponse(commentPositions map[string]RecordStringBeforeStringOrNullAfterStringOrNullValue, status ImportedAPIStatusSUCCESS, ) *ChangeCommentPinStatusResponse`

NewChangeCommentPinStatusResponse instantiates a new ChangeCommentPinStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeCommentPinStatusResponseWithDefaults

`func NewChangeCommentPinStatusResponseWithDefaults() *ChangeCommentPinStatusResponse`

NewChangeCommentPinStatusResponseWithDefaults instantiates a new ChangeCommentPinStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentPositions

`func (o *ChangeCommentPinStatusResponse) GetCommentPositions() map[string]RecordStringBeforeStringOrNullAfterStringOrNullValue`

GetCommentPositions returns the CommentPositions field if non-nil, zero value otherwise.

### GetCommentPositionsOk

`func (o *ChangeCommentPinStatusResponse) GetCommentPositionsOk() (*map[string]RecordStringBeforeStringOrNullAfterStringOrNullValue, bool)`

GetCommentPositionsOk returns a tuple with the CommentPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentPositions

`func (o *ChangeCommentPinStatusResponse) SetCommentPositions(v map[string]RecordStringBeforeStringOrNullAfterStringOrNullValue)`

SetCommentPositions sets CommentPositions field to given value.


### GetStatus

`func (o *ChangeCommentPinStatusResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeCommentPinStatusResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeCommentPinStatusResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


