# CreateModeratorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Moderator** | [**Moderator**](Moderator.md) |  | 

## Methods

### NewCreateModeratorResponse

`func NewCreateModeratorResponse(status APIStatus, moderator Moderator, ) *CreateModeratorResponse`

NewCreateModeratorResponse instantiates a new CreateModeratorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateModeratorResponseWithDefaults

`func NewCreateModeratorResponseWithDefaults() *CreateModeratorResponse`

NewCreateModeratorResponseWithDefaults instantiates a new CreateModeratorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateModeratorResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateModeratorResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateModeratorResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetModerator

`func (o *CreateModeratorResponse) GetModerator() Moderator`

GetModerator returns the Moderator field if non-nil, zero value otherwise.

### GetModeratorOk

`func (o *CreateModeratorResponse) GetModeratorOk() (*Moderator, bool)`

GetModeratorOk returns a tuple with the Moderator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerator

`func (o *CreateModeratorResponse) SetModerator(v Moderator)`

SetModerator sets Moderator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


