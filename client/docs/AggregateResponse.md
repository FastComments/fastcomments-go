# AggregateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Data** | [**[]AggregationItem**](AggregationItem.md) |  | 
**Stats** | Pointer to [**AggregationResponseStats**](AggregationResponseStats.md) |  | [optional] 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**ValidResourceNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAggregateResponse

`func NewAggregateResponse(status APIStatus, data []AggregationItem, reason string, code string, ) *AggregateResponse`

NewAggregateResponse instantiates a new AggregateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateResponseWithDefaults

`func NewAggregateResponseWithDefaults() *AggregateResponse`

NewAggregateResponseWithDefaults instantiates a new AggregateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AggregateResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregateResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregateResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetData

`func (o *AggregateResponse) GetData() []AggregationItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AggregateResponse) GetDataOk() (*[]AggregationItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AggregateResponse) SetData(v []AggregationItem)`

SetData sets Data field to given value.


### GetStats

`func (o *AggregateResponse) GetStats() AggregationResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AggregateResponse) GetStatsOk() (*AggregationResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AggregateResponse) SetStats(v AggregationResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AggregateResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetReason

`func (o *AggregateResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AggregateResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AggregateResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *AggregateResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AggregateResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AggregateResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetValidResourceNames

`func (o *AggregateResponse) GetValidResourceNames() []string`

GetValidResourceNames returns the ValidResourceNames field if non-nil, zero value otherwise.

### GetValidResourceNamesOk

`func (o *AggregateResponse) GetValidResourceNamesOk() (*[]string, bool)`

GetValidResourceNamesOk returns a tuple with the ValidResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResourceNames

`func (o *AggregateResponse) SetValidResourceNames(v []string)`

SetValidResourceNames sets ValidResourceNames field to given value.

### HasValidResourceNames

`func (o *AggregateResponse) HasValidResourceNames() bool`

HasValidResourceNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


