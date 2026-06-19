# BannedUserMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedOn** | [**BannedUserMatchType**](BannedUserMatchType.md) |  | 
**MatchedOnValue** | [**NullableBannedUserMatchMatchedOnValue**](BannedUserMatchMatchedOnValue.md) |  | 

## Methods

### NewBannedUserMatch

`func NewBannedUserMatch(matchedOn BannedUserMatchType, matchedOnValue NullableBannedUserMatchMatchedOnValue, ) *BannedUserMatch`

NewBannedUserMatch instantiates a new BannedUserMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannedUserMatchWithDefaults

`func NewBannedUserMatchWithDefaults() *BannedUserMatch`

NewBannedUserMatchWithDefaults instantiates a new BannedUserMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedOn

`func (o *BannedUserMatch) GetMatchedOn() BannedUserMatchType`

GetMatchedOn returns the MatchedOn field if non-nil, zero value otherwise.

### GetMatchedOnOk

`func (o *BannedUserMatch) GetMatchedOnOk() (*BannedUserMatchType, bool)`

GetMatchedOnOk returns a tuple with the MatchedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedOn

`func (o *BannedUserMatch) SetMatchedOn(v BannedUserMatchType)`

SetMatchedOn sets MatchedOn field to given value.


### GetMatchedOnValue

`func (o *BannedUserMatch) GetMatchedOnValue() BannedUserMatchMatchedOnValue`

GetMatchedOnValue returns the MatchedOnValue field if non-nil, zero value otherwise.

### GetMatchedOnValueOk

`func (o *BannedUserMatch) GetMatchedOnValueOk() (*BannedUserMatchMatchedOnValue, bool)`

GetMatchedOnValueOk returns a tuple with the MatchedOnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedOnValue

`func (o *BannedUserMatch) SetMatchedOnValue(v BannedUserMatchMatchedOnValue)`

SetMatchedOnValue sets MatchedOnValue field to given value.


### SetMatchedOnValueNil

`func (o *BannedUserMatch) SetMatchedOnValueNil(b bool)`

 SetMatchedOnValueNil sets the value for MatchedOnValue to be an explicit nil

### UnsetMatchedOnValue
`func (o *BannedUserMatch) UnsetMatchedOnValue()`

UnsetMatchedOnValue ensures that no value is present for MatchedOnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


