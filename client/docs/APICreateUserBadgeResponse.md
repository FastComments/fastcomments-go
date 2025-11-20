# APICreateUserBadgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**UserBadge** | [**UserBadge**](UserBadge.md) |  | 
**Notes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAPICreateUserBadgeResponse

`func NewAPICreateUserBadgeResponse(status APIStatus, userBadge UserBadge, ) *APICreateUserBadgeResponse`

NewAPICreateUserBadgeResponse instantiates a new APICreateUserBadgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICreateUserBadgeResponseWithDefaults

`func NewAPICreateUserBadgeResponseWithDefaults() *APICreateUserBadgeResponse`

NewAPICreateUserBadgeResponseWithDefaults instantiates a new APICreateUserBadgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APICreateUserBadgeResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APICreateUserBadgeResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APICreateUserBadgeResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUserBadge

`func (o *APICreateUserBadgeResponse) GetUserBadge() UserBadge`

GetUserBadge returns the UserBadge field if non-nil, zero value otherwise.

### GetUserBadgeOk

`func (o *APICreateUserBadgeResponse) GetUserBadgeOk() (*UserBadge, bool)`

GetUserBadgeOk returns a tuple with the UserBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBadge

`func (o *APICreateUserBadgeResponse) SetUserBadge(v UserBadge)`

SetUserBadge sets UserBadge field to given value.


### GetNotes

`func (o *APICreateUserBadgeResponse) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *APICreateUserBadgeResponse) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *APICreateUserBadgeResponse) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *APICreateUserBadgeResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


