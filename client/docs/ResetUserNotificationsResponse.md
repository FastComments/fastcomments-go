# ResetUserNotificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewResetUserNotificationsResponse

`func NewResetUserNotificationsResponse(status ImportedAPIStatusSUCCESS, ) *ResetUserNotificationsResponse`

NewResetUserNotificationsResponse instantiates a new ResetUserNotificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetUserNotificationsResponseWithDefaults

`func NewResetUserNotificationsResponseWithDefaults() *ResetUserNotificationsResponse`

NewResetUserNotificationsResponseWithDefaults instantiates a new ResetUserNotificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResetUserNotificationsResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResetUserNotificationsResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResetUserNotificationsResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.


### GetCode

`func (o *ResetUserNotificationsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResetUserNotificationsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResetUserNotificationsResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResetUserNotificationsResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


