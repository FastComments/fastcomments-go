# ModerationAPICountCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Count** | **float64** |  | 

## Methods

### NewModerationAPICountCommentsResponse

`func NewModerationAPICountCommentsResponse(status APIStatus, count float64, ) *ModerationAPICountCommentsResponse`

NewModerationAPICountCommentsResponse instantiates a new ModerationAPICountCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPICountCommentsResponseWithDefaults

`func NewModerationAPICountCommentsResponseWithDefaults() *ModerationAPICountCommentsResponse`

NewModerationAPICountCommentsResponseWithDefaults instantiates a new ModerationAPICountCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModerationAPICountCommentsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationAPICountCommentsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationAPICountCommentsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCount

`func (o *ModerationAPICountCommentsResponse) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModerationAPICountCommentsResponse) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModerationAPICountCommentsResponse) SetCount(v float64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


