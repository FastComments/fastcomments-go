# UserReactsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reacts** | **map[string]map[string]bool** |  | 

## Methods

### NewUserReactsResponse

`func NewUserReactsResponse(status APIStatus, reacts map[string]map[string]bool, ) *UserReactsResponse`

NewUserReactsResponse instantiates a new UserReactsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReactsResponseWithDefaults

`func NewUserReactsResponseWithDefaults() *UserReactsResponse`

NewUserReactsResponseWithDefaults instantiates a new UserReactsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserReactsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserReactsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserReactsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReacts

`func (o *UserReactsResponse) GetReacts() map[string]map[string]bool`

GetReacts returns the Reacts field if non-nil, zero value otherwise.

### GetReactsOk

`func (o *UserReactsResponse) GetReactsOk() (*map[string]map[string]bool, bool)`

GetReactsOk returns a tuple with the Reacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReacts

`func (o *UserReactsResponse) SetReacts(v map[string]map[string]bool)`

SetReacts sets Reacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


