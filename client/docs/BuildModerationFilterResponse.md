# BuildModerationFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**ModerationFilter** | [**ModerationFilter**](ModerationFilter.md) |  | 

## Methods

### NewBuildModerationFilterResponse

`func NewBuildModerationFilterResponse(status string, moderationFilter ModerationFilter, ) *BuildModerationFilterResponse`

NewBuildModerationFilterResponse instantiates a new BuildModerationFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildModerationFilterResponseWithDefaults

`func NewBuildModerationFilterResponseWithDefaults() *BuildModerationFilterResponse`

NewBuildModerationFilterResponseWithDefaults instantiates a new BuildModerationFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BuildModerationFilterResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildModerationFilterResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildModerationFilterResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetModerationFilter

`func (o *BuildModerationFilterResponse) GetModerationFilter() ModerationFilter`

GetModerationFilter returns the ModerationFilter field if non-nil, zero value otherwise.

### GetModerationFilterOk

`func (o *BuildModerationFilterResponse) GetModerationFilterOk() (*ModerationFilter, bool)`

GetModerationFilterOk returns a tuple with the ModerationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationFilter

`func (o *BuildModerationFilterResponse) SetModerationFilter(v ModerationFilter)`

SetModerationFilter sets ModerationFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


