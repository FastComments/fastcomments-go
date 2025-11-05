# GetSubscriptionsAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**[]APIUserSubscription**](APIUserSubscription.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewGetSubscriptionsAPIResponse

`func NewGetSubscriptionsAPIResponse(status string, ) *GetSubscriptionsAPIResponse`

NewGetSubscriptionsAPIResponse instantiates a new GetSubscriptionsAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubscriptionsAPIResponseWithDefaults

`func NewGetSubscriptionsAPIResponseWithDefaults() *GetSubscriptionsAPIResponse`

NewGetSubscriptionsAPIResponseWithDefaults instantiates a new GetSubscriptionsAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *GetSubscriptionsAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetSubscriptionsAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetSubscriptionsAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetSubscriptionsAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *GetSubscriptionsAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSubscriptionsAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSubscriptionsAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSubscriptionsAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSubscriptions

`func (o *GetSubscriptionsAPIResponse) GetSubscriptions() []APIUserSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *GetSubscriptionsAPIResponse) GetSubscriptionsOk() (*[]APIUserSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *GetSubscriptionsAPIResponse) SetSubscriptions(v []APIUserSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *GetSubscriptionsAPIResponse) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetStatus

`func (o *GetSubscriptionsAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSubscriptionsAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSubscriptionsAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


