# ReactFeedPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**ReactType** | **string** |  | 
**IsUndo** | **bool** |  | 

## Methods

### NewReactFeedPostResponse

`func NewReactFeedPostResponse(status APIStatus, reactType string, isUndo bool, ) *ReactFeedPostResponse`

NewReactFeedPostResponse instantiates a new ReactFeedPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactFeedPostResponseWithDefaults

`func NewReactFeedPostResponseWithDefaults() *ReactFeedPostResponse`

NewReactFeedPostResponseWithDefaults instantiates a new ReactFeedPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReactFeedPostResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReactFeedPostResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReactFeedPostResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReactType

`func (o *ReactFeedPostResponse) GetReactType() string`

GetReactType returns the ReactType field if non-nil, zero value otherwise.

### GetReactTypeOk

`func (o *ReactFeedPostResponse) GetReactTypeOk() (*string, bool)`

GetReactTypeOk returns a tuple with the ReactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactType

`func (o *ReactFeedPostResponse) SetReactType(v string)`

SetReactType sets ReactType field to given value.


### GetIsUndo

`func (o *ReactFeedPostResponse) GetIsUndo() bool`

GetIsUndo returns the IsUndo field if non-nil, zero value otherwise.

### GetIsUndoOk

`func (o *ReactFeedPostResponse) GetIsUndoOk() (*bool, bool)`

GetIsUndoOk returns a tuple with the IsUndo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUndo

`func (o *ReactFeedPostResponse) SetIsUndo(v bool)`

SetIsUndo sets IsUndo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


