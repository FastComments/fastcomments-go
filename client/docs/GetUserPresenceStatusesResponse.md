# GetUserPresenceStatusesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 
**UserIdsOnline** | **map[string]bool** | Construct a type with a set of properties K of type T | 

## Methods

### NewGetUserPresenceStatusesResponse

`func NewGetUserPresenceStatusesResponse(status ImportedAPIStatusSUCCESS, userIdsOnline map[string]bool, ) *GetUserPresenceStatusesResponse`

NewGetUserPresenceStatusesResponse instantiates a new GetUserPresenceStatusesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserPresenceStatusesResponseWithDefaults

`func NewGetUserPresenceStatusesResponseWithDefaults() *GetUserPresenceStatusesResponse`

NewGetUserPresenceStatusesResponseWithDefaults instantiates a new GetUserPresenceStatusesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetUserPresenceStatusesResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserPresenceStatusesResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserPresenceStatusesResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.


### GetUserIdsOnline

`func (o *GetUserPresenceStatusesResponse) GetUserIdsOnline() map[string]bool`

GetUserIdsOnline returns the UserIdsOnline field if non-nil, zero value otherwise.

### GetUserIdsOnlineOk

`func (o *GetUserPresenceStatusesResponse) GetUserIdsOnlineOk() (*map[string]bool, bool)`

GetUserIdsOnlineOk returns a tuple with the UserIdsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdsOnline

`func (o *GetUserPresenceStatusesResponse) SetUserIdsOnline(v map[string]bool)`

SetUserIdsOnline sets UserIdsOnline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


