# GetUserInternalProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | [**GetUserInternalProfileResponseProfile**](GetUserInternalProfileResponseProfile.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetUserInternalProfileResponse

`func NewGetUserInternalProfileResponse(profile GetUserInternalProfileResponseProfile, status APIStatus, ) *GetUserInternalProfileResponse`

NewGetUserInternalProfileResponse instantiates a new GetUserInternalProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInternalProfileResponseWithDefaults

`func NewGetUserInternalProfileResponseWithDefaults() *GetUserInternalProfileResponse`

NewGetUserInternalProfileResponseWithDefaults instantiates a new GetUserInternalProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *GetUserInternalProfileResponse) GetProfile() GetUserInternalProfileResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetUserInternalProfileResponse) GetProfileOk() (*GetUserInternalProfileResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetUserInternalProfileResponse) SetProfile(v GetUserInternalProfileResponseProfile)`

SetProfile sets Profile field to given value.


### GetStatus

`func (o *GetUserInternalProfileResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUserInternalProfileResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUserInternalProfileResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


