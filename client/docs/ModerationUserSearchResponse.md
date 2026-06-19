# ModerationUserSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]ModerationUserSearchProjected**](ModerationUserSearchProjected.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationUserSearchResponse

`func NewModerationUserSearchResponse(users []ModerationUserSearchProjected, status APIStatus, ) *ModerationUserSearchResponse`

NewModerationUserSearchResponse instantiates a new ModerationUserSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationUserSearchResponseWithDefaults

`func NewModerationUserSearchResponseWithDefaults() *ModerationUserSearchResponse`

NewModerationUserSearchResponseWithDefaults instantiates a new ModerationUserSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ModerationUserSearchResponse) GetUsers() []ModerationUserSearchProjected`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModerationUserSearchResponse) GetUsersOk() (*[]ModerationUserSearchProjected, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModerationUserSearchResponse) SetUsers(v []ModerationUserSearchProjected)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *ModerationUserSearchResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationUserSearchResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationUserSearchResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


