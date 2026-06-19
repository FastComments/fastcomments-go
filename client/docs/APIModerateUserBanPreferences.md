# APIModerateUserBanPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldBanEmail** | **bool** |  | 
**ShouldBanByIP** | **bool** |  | 
**LastBanType** | **string** |  | 
**LastBanDuration** | **string** |  | 

## Methods

### NewAPIModerateUserBanPreferences

`func NewAPIModerateUserBanPreferences(shouldBanEmail bool, shouldBanByIP bool, lastBanType string, lastBanDuration string, ) *APIModerateUserBanPreferences`

NewAPIModerateUserBanPreferences instantiates a new APIModerateUserBanPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIModerateUserBanPreferencesWithDefaults

`func NewAPIModerateUserBanPreferencesWithDefaults() *APIModerateUserBanPreferences`

NewAPIModerateUserBanPreferencesWithDefaults instantiates a new APIModerateUserBanPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldBanEmail

`func (o *APIModerateUserBanPreferences) GetShouldBanEmail() bool`

GetShouldBanEmail returns the ShouldBanEmail field if non-nil, zero value otherwise.

### GetShouldBanEmailOk

`func (o *APIModerateUserBanPreferences) GetShouldBanEmailOk() (*bool, bool)`

GetShouldBanEmailOk returns a tuple with the ShouldBanEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBanEmail

`func (o *APIModerateUserBanPreferences) SetShouldBanEmail(v bool)`

SetShouldBanEmail sets ShouldBanEmail field to given value.


### GetShouldBanByIP

`func (o *APIModerateUserBanPreferences) GetShouldBanByIP() bool`

GetShouldBanByIP returns the ShouldBanByIP field if non-nil, zero value otherwise.

### GetShouldBanByIPOk

`func (o *APIModerateUserBanPreferences) GetShouldBanByIPOk() (*bool, bool)`

GetShouldBanByIPOk returns a tuple with the ShouldBanByIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBanByIP

`func (o *APIModerateUserBanPreferences) SetShouldBanByIP(v bool)`

SetShouldBanByIP sets ShouldBanByIP field to given value.


### GetLastBanType

`func (o *APIModerateUserBanPreferences) GetLastBanType() string`

GetLastBanType returns the LastBanType field if non-nil, zero value otherwise.

### GetLastBanTypeOk

`func (o *APIModerateUserBanPreferences) GetLastBanTypeOk() (*string, bool)`

GetLastBanTypeOk returns a tuple with the LastBanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBanType

`func (o *APIModerateUserBanPreferences) SetLastBanType(v string)`

SetLastBanType sets LastBanType field to given value.


### GetLastBanDuration

`func (o *APIModerateUserBanPreferences) GetLastBanDuration() string`

GetLastBanDuration returns the LastBanDuration field if non-nil, zero value otherwise.

### GetLastBanDurationOk

`func (o *APIModerateUserBanPreferences) GetLastBanDurationOk() (*string, bool)`

GetLastBanDurationOk returns a tuple with the LastBanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBanDuration

`func (o *APIModerateUserBanPreferences) SetLastBanDuration(v string)`

SetLastBanDuration sets LastBanDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


