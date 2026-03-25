# UpdateSubscriptionAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Subscription** | Pointer to [**APIUserSubscription**](APIUserSubscription.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewUpdateSubscriptionAPIResponse

`func NewUpdateSubscriptionAPIResponse(status string, ) *UpdateSubscriptionAPIResponse`

NewUpdateSubscriptionAPIResponse instantiates a new UpdateSubscriptionAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionAPIResponseWithDefaults

`func NewUpdateSubscriptionAPIResponseWithDefaults() *UpdateSubscriptionAPIResponse`

NewUpdateSubscriptionAPIResponseWithDefaults instantiates a new UpdateSubscriptionAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *UpdateSubscriptionAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateSubscriptionAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateSubscriptionAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateSubscriptionAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *UpdateSubscriptionAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateSubscriptionAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateSubscriptionAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateSubscriptionAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSubscription

`func (o *UpdateSubscriptionAPIResponse) GetSubscription() APIUserSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UpdateSubscriptionAPIResponse) GetSubscriptionOk() (*APIUserSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UpdateSubscriptionAPIResponse) SetSubscription(v APIUserSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UpdateSubscriptionAPIResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateSubscriptionAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSubscriptionAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSubscriptionAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


