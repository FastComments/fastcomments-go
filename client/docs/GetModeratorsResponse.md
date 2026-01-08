# GetModeratorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Moderators** | [**[]Moderator**](Moderator.md) |  | 

## Methods

### NewGetModeratorsResponse

`func NewGetModeratorsResponse(status APIStatus, moderators []Moderator, ) *GetModeratorsResponse`

NewGetModeratorsResponse instantiates a new GetModeratorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetModeratorsResponseWithDefaults

`func NewGetModeratorsResponseWithDefaults() *GetModeratorsResponse`

NewGetModeratorsResponseWithDefaults instantiates a new GetModeratorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetModeratorsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetModeratorsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetModeratorsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetModerators

`func (o *GetModeratorsResponse) GetModerators() []Moderator`

GetModerators returns the Moderators field if non-nil, zero value otherwise.

### GetModeratorsOk

`func (o *GetModeratorsResponse) GetModeratorsOk() (*[]Moderator, bool)`

GetModeratorsOk returns a tuple with the Moderators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerators

`func (o *GetModeratorsResponse) SetModerators(v []Moderator)`

SetModerators sets Moderators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


