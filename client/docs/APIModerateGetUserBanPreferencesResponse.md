# APIModerateGetUserBanPreferencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preferences** | [**NullableAPIModerateUserBanPreferences**](APIModerateUserBanPreferences.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewAPIModerateGetUserBanPreferencesResponse

`func NewAPIModerateGetUserBanPreferencesResponse(preferences NullableAPIModerateUserBanPreferences, status APIStatus, ) *APIModerateGetUserBanPreferencesResponse`

NewAPIModerateGetUserBanPreferencesResponse instantiates a new APIModerateGetUserBanPreferencesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIModerateGetUserBanPreferencesResponseWithDefaults

`func NewAPIModerateGetUserBanPreferencesResponseWithDefaults() *APIModerateGetUserBanPreferencesResponse`

NewAPIModerateGetUserBanPreferencesResponseWithDefaults instantiates a new APIModerateGetUserBanPreferencesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferences

`func (o *APIModerateGetUserBanPreferencesResponse) GetPreferences() APIModerateUserBanPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *APIModerateGetUserBanPreferencesResponse) GetPreferencesOk() (*APIModerateUserBanPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *APIModerateGetUserBanPreferencesResponse) SetPreferences(v APIModerateUserBanPreferences)`

SetPreferences sets Preferences field to given value.


### SetPreferencesNil

`func (o *APIModerateGetUserBanPreferencesResponse) SetPreferencesNil(b bool)`

 SetPreferencesNil sets the value for Preferences to be an explicit nil

### UnsetPreferences
`func (o *APIModerateGetUserBanPreferencesResponse) UnsetPreferences()`

UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
### GetStatus

`func (o *APIModerateGetUserBanPreferencesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIModerateGetUserBanPreferencesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIModerateGetUserBanPreferencesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


