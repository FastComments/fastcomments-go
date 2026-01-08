# GetModeratorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Moderator** | [**Moderator**](Moderator.md) |  | 

## Methods

### NewGetModeratorResponse

`func NewGetModeratorResponse(status APIStatus, moderator Moderator, ) *GetModeratorResponse`

NewGetModeratorResponse instantiates a new GetModeratorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetModeratorResponseWithDefaults

`func NewGetModeratorResponseWithDefaults() *GetModeratorResponse`

NewGetModeratorResponseWithDefaults instantiates a new GetModeratorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetModeratorResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetModeratorResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetModeratorResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetModerator

`func (o *GetModeratorResponse) GetModerator() Moderator`

GetModerator returns the Moderator field if non-nil, zero value otherwise.

### GetModeratorOk

`func (o *GetModeratorResponse) GetModeratorOk() (*Moderator, bool)`

GetModeratorOk returns a tuple with the Moderator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerator

`func (o *GetModeratorResponse) SetModerator(v Moderator)`

SetModerator sets Moderator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


