# GetV2PageReactUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserNames** | **[]string** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetV2PageReactUsersResponse

`func NewGetV2PageReactUsersResponse(userNames []string, status APIStatus, ) *GetV2PageReactUsersResponse`

NewGetV2PageReactUsersResponse instantiates a new GetV2PageReactUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetV2PageReactUsersResponseWithDefaults

`func NewGetV2PageReactUsersResponseWithDefaults() *GetV2PageReactUsersResponse`

NewGetV2PageReactUsersResponseWithDefaults instantiates a new GetV2PageReactUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserNames

`func (o *GetV2PageReactUsersResponse) GetUserNames() []string`

GetUserNames returns the UserNames field if non-nil, zero value otherwise.

### GetUserNamesOk

`func (o *GetV2PageReactUsersResponse) GetUserNamesOk() (*[]string, bool)`

GetUserNamesOk returns a tuple with the UserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNames

`func (o *GetV2PageReactUsersResponse) SetUserNames(v []string)`

SetUserNames sets UserNames field to given value.


### GetStatus

`func (o *GetV2PageReactUsersResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetV2PageReactUsersResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetV2PageReactUsersResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


