# UpdateUserNotificationPageSubscriptionStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**MatchedCount** | **int64** |  | 
**ModifiedCount** | **int64** |  | 
**Note** | **string** |  | 

## Methods

### NewUpdateUserNotificationPageSubscriptionStatusResponse

`func NewUpdateUserNotificationPageSubscriptionStatusResponse(status APIStatus, matchedCount int64, modifiedCount int64, note string, ) *UpdateUserNotificationPageSubscriptionStatusResponse`

NewUpdateUserNotificationPageSubscriptionStatusResponse instantiates a new UpdateUserNotificationPageSubscriptionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserNotificationPageSubscriptionStatusResponseWithDefaults

`func NewUpdateUserNotificationPageSubscriptionStatusResponseWithDefaults() *UpdateUserNotificationPageSubscriptionStatusResponse`

NewUpdateUserNotificationPageSubscriptionStatusResponseWithDefaults instantiates a new UpdateUserNotificationPageSubscriptionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetMatchedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetMatchedCount() int64`

GetMatchedCount returns the MatchedCount field if non-nil, zero value otherwise.

### GetMatchedCountOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetMatchedCountOk() (*int64, bool)`

GetMatchedCountOk returns a tuple with the MatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetMatchedCount(v int64)`

SetMatchedCount sets MatchedCount field to given value.


### GetModifiedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetModifiedCount() int64`

GetModifiedCount returns the ModifiedCount field if non-nil, zero value otherwise.

### GetModifiedCountOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetModifiedCountOk() (*int64, bool)`

GetModifiedCountOk returns a tuple with the ModifiedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedCount

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetModifiedCount(v int64)`

SetModifiedCount sets ModifiedCount field to given value.


### GetNote

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateUserNotificationPageSubscriptionStatusResponse) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


