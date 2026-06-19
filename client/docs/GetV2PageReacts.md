# GetV2PageReacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactedIds** | Pointer to **[]string** |  | [optional] 
**Counts** | Pointer to **map[string]float64** | Construct a type with a set of properties K of type T | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetV2PageReacts

`func NewGetV2PageReacts(status APIStatus, ) *GetV2PageReacts`

NewGetV2PageReacts instantiates a new GetV2PageReacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetV2PageReactsWithDefaults

`func NewGetV2PageReactsWithDefaults() *GetV2PageReacts`

NewGetV2PageReactsWithDefaults instantiates a new GetV2PageReacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactedIds

`func (o *GetV2PageReacts) GetReactedIds() []string`

GetReactedIds returns the ReactedIds field if non-nil, zero value otherwise.

### GetReactedIdsOk

`func (o *GetV2PageReacts) GetReactedIdsOk() (*[]string, bool)`

GetReactedIdsOk returns a tuple with the ReactedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactedIds

`func (o *GetV2PageReacts) SetReactedIds(v []string)`

SetReactedIds sets ReactedIds field to given value.

### HasReactedIds

`func (o *GetV2PageReacts) HasReactedIds() bool`

HasReactedIds returns a boolean if a field has been set.

### GetCounts

`func (o *GetV2PageReacts) GetCounts() map[string]float64`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetV2PageReacts) GetCountsOk() (*map[string]float64, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetV2PageReacts) SetCounts(v map[string]float64)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *GetV2PageReacts) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetStatus

`func (o *GetV2PageReacts) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetV2PageReacts) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetV2PageReacts) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


