# PreBanSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Usernames** | **[]string** |  | 
**Count** | **float64** |  | 

## Methods

### NewPreBanSummary

`func NewPreBanSummary(status APIStatus, usernames []string, count float64, ) *PreBanSummary`

NewPreBanSummary instantiates a new PreBanSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreBanSummaryWithDefaults

`func NewPreBanSummaryWithDefaults() *PreBanSummary`

NewPreBanSummaryWithDefaults instantiates a new PreBanSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PreBanSummary) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreBanSummary) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreBanSummary) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUsernames

`func (o *PreBanSummary) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *PreBanSummary) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *PreBanSummary) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.


### GetCount

`func (o *PreBanSummary) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PreBanSummary) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PreBanSummary) SetCount(v float64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


