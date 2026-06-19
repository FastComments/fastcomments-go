# ModerationSiteSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sites** | [**[]ModerationSiteSearchProjected**](ModerationSiteSearchProjected.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationSiteSearchResponse

`func NewModerationSiteSearchResponse(sites []ModerationSiteSearchProjected, status APIStatus, ) *ModerationSiteSearchResponse`

NewModerationSiteSearchResponse instantiates a new ModerationSiteSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationSiteSearchResponseWithDefaults

`func NewModerationSiteSearchResponseWithDefaults() *ModerationSiteSearchResponse`

NewModerationSiteSearchResponseWithDefaults instantiates a new ModerationSiteSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSites

`func (o *ModerationSiteSearchResponse) GetSites() []ModerationSiteSearchProjected`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ModerationSiteSearchResponse) GetSitesOk() (*[]ModerationSiteSearchProjected, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ModerationSiteSearchResponse) SetSites(v []ModerationSiteSearchProjected)`

SetSites sets Sites field to given value.


### GetStatus

`func (o *ModerationSiteSearchResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationSiteSearchResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationSiteSearchResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


