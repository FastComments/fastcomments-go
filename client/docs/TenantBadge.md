# TenantBadge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CreatedByUserId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Enabled** | **bool** |  | 
**UrlId** | Pointer to **NullableString** |  | [optional] 
**Type** | **float64** |  | 
**Threshold** | **float64** |  | 
**Uses** | **float64** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**DisplayLabel** | **string** |  | 
**DisplaySrc** | **NullableString** |  | 
**BackgroundColor** | **NullableString** |  | 
**BorderColor** | **NullableString** |  | 
**TextColor** | **NullableString** |  | 
**CssClass** | Pointer to **NullableString** |  | [optional] 
**VeteranUserThresholdMillis** | Pointer to **NullableFloat64** |  | [optional] 
**IsAwaitingReprocess** | **bool** |  | 
**IsAwaitingDeletion** | **bool** |  | 
**ReplacesBadgeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantBadge

`func NewTenantBadge(id string, tenantId string, createdByUserId string, createdAt time.Time, enabled bool, type_ float64, threshold float64, uses float64, name string, description string, displayLabel string, displaySrc NullableString, backgroundColor NullableString, borderColor NullableString, textColor NullableString, isAwaitingReprocess bool, isAwaitingDeletion bool, ) *TenantBadge`

NewTenantBadge instantiates a new TenantBadge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantBadgeWithDefaults

`func NewTenantBadgeWithDefaults() *TenantBadge`

NewTenantBadgeWithDefaults instantiates a new TenantBadge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantBadge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantBadge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantBadge) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *TenantBadge) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantBadge) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantBadge) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCreatedByUserId

`func (o *TenantBadge) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *TenantBadge) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *TenantBadge) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetCreatedAt

`func (o *TenantBadge) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantBadge) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantBadge) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *TenantBadge) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TenantBadge) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TenantBadge) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUrlId

`func (o *TenantBadge) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *TenantBadge) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *TenantBadge) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.

### HasUrlId

`func (o *TenantBadge) HasUrlId() bool`

HasUrlId returns a boolean if a field has been set.

### SetUrlIdNil

`func (o *TenantBadge) SetUrlIdNil(b bool)`

 SetUrlIdNil sets the value for UrlId to be an explicit nil

### UnsetUrlId
`func (o *TenantBadge) UnsetUrlId()`

UnsetUrlId ensures that no value is present for UrlId, not even an explicit nil
### GetType

`func (o *TenantBadge) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TenantBadge) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TenantBadge) SetType(v float64)`

SetType sets Type field to given value.


### GetThreshold

`func (o *TenantBadge) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TenantBadge) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TenantBadge) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetUses

`func (o *TenantBadge) GetUses() float64`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *TenantBadge) GetUsesOk() (*float64, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *TenantBadge) SetUses(v float64)`

SetUses sets Uses field to given value.


### GetName

`func (o *TenantBadge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantBadge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantBadge) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TenantBadge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantBadge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantBadge) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayLabel

`func (o *TenantBadge) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *TenantBadge) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *TenantBadge) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.


### GetDisplaySrc

`func (o *TenantBadge) GetDisplaySrc() string`

GetDisplaySrc returns the DisplaySrc field if non-nil, zero value otherwise.

### GetDisplaySrcOk

`func (o *TenantBadge) GetDisplaySrcOk() (*string, bool)`

GetDisplaySrcOk returns a tuple with the DisplaySrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySrc

`func (o *TenantBadge) SetDisplaySrc(v string)`

SetDisplaySrc sets DisplaySrc field to given value.


### SetDisplaySrcNil

`func (o *TenantBadge) SetDisplaySrcNil(b bool)`

 SetDisplaySrcNil sets the value for DisplaySrc to be an explicit nil

### UnsetDisplaySrc
`func (o *TenantBadge) UnsetDisplaySrc()`

UnsetDisplaySrc ensures that no value is present for DisplaySrc, not even an explicit nil
### GetBackgroundColor

`func (o *TenantBadge) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *TenantBadge) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *TenantBadge) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.


### SetBackgroundColorNil

`func (o *TenantBadge) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *TenantBadge) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBorderColor

`func (o *TenantBadge) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *TenantBadge) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *TenantBadge) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.


### SetBorderColorNil

`func (o *TenantBadge) SetBorderColorNil(b bool)`

 SetBorderColorNil sets the value for BorderColor to be an explicit nil

### UnsetBorderColor
`func (o *TenantBadge) UnsetBorderColor()`

UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
### GetTextColor

`func (o *TenantBadge) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *TenantBadge) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *TenantBadge) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.


### SetTextColorNil

`func (o *TenantBadge) SetTextColorNil(b bool)`

 SetTextColorNil sets the value for TextColor to be an explicit nil

### UnsetTextColor
`func (o *TenantBadge) UnsetTextColor()`

UnsetTextColor ensures that no value is present for TextColor, not even an explicit nil
### GetCssClass

`func (o *TenantBadge) GetCssClass() string`

GetCssClass returns the CssClass field if non-nil, zero value otherwise.

### GetCssClassOk

`func (o *TenantBadge) GetCssClassOk() (*string, bool)`

GetCssClassOk returns a tuple with the CssClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssClass

`func (o *TenantBadge) SetCssClass(v string)`

SetCssClass sets CssClass field to given value.

### HasCssClass

`func (o *TenantBadge) HasCssClass() bool`

HasCssClass returns a boolean if a field has been set.

### SetCssClassNil

`func (o *TenantBadge) SetCssClassNil(b bool)`

 SetCssClassNil sets the value for CssClass to be an explicit nil

### UnsetCssClass
`func (o *TenantBadge) UnsetCssClass()`

UnsetCssClass ensures that no value is present for CssClass, not even an explicit nil
### GetVeteranUserThresholdMillis

`func (o *TenantBadge) GetVeteranUserThresholdMillis() float64`

GetVeteranUserThresholdMillis returns the VeteranUserThresholdMillis field if non-nil, zero value otherwise.

### GetVeteranUserThresholdMillisOk

`func (o *TenantBadge) GetVeteranUserThresholdMillisOk() (*float64, bool)`

GetVeteranUserThresholdMillisOk returns a tuple with the VeteranUserThresholdMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVeteranUserThresholdMillis

`func (o *TenantBadge) SetVeteranUserThresholdMillis(v float64)`

SetVeteranUserThresholdMillis sets VeteranUserThresholdMillis field to given value.

### HasVeteranUserThresholdMillis

`func (o *TenantBadge) HasVeteranUserThresholdMillis() bool`

HasVeteranUserThresholdMillis returns a boolean if a field has been set.

### SetVeteranUserThresholdMillisNil

`func (o *TenantBadge) SetVeteranUserThresholdMillisNil(b bool)`

 SetVeteranUserThresholdMillisNil sets the value for VeteranUserThresholdMillis to be an explicit nil

### UnsetVeteranUserThresholdMillis
`func (o *TenantBadge) UnsetVeteranUserThresholdMillis()`

UnsetVeteranUserThresholdMillis ensures that no value is present for VeteranUserThresholdMillis, not even an explicit nil
### GetIsAwaitingReprocess

`func (o *TenantBadge) GetIsAwaitingReprocess() bool`

GetIsAwaitingReprocess returns the IsAwaitingReprocess field if non-nil, zero value otherwise.

### GetIsAwaitingReprocessOk

`func (o *TenantBadge) GetIsAwaitingReprocessOk() (*bool, bool)`

GetIsAwaitingReprocessOk returns a tuple with the IsAwaitingReprocess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAwaitingReprocess

`func (o *TenantBadge) SetIsAwaitingReprocess(v bool)`

SetIsAwaitingReprocess sets IsAwaitingReprocess field to given value.


### GetIsAwaitingDeletion

`func (o *TenantBadge) GetIsAwaitingDeletion() bool`

GetIsAwaitingDeletion returns the IsAwaitingDeletion field if non-nil, zero value otherwise.

### GetIsAwaitingDeletionOk

`func (o *TenantBadge) GetIsAwaitingDeletionOk() (*bool, bool)`

GetIsAwaitingDeletionOk returns a tuple with the IsAwaitingDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAwaitingDeletion

`func (o *TenantBadge) SetIsAwaitingDeletion(v bool)`

SetIsAwaitingDeletion sets IsAwaitingDeletion field to given value.


### GetReplacesBadgeId

`func (o *TenantBadge) GetReplacesBadgeId() string`

GetReplacesBadgeId returns the ReplacesBadgeId field if non-nil, zero value otherwise.

### GetReplacesBadgeIdOk

`func (o *TenantBadge) GetReplacesBadgeIdOk() (*string, bool)`

GetReplacesBadgeIdOk returns a tuple with the ReplacesBadgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacesBadgeId

`func (o *TenantBadge) SetReplacesBadgeId(v string)`

SetReplacesBadgeId sets ReplacesBadgeId field to given value.

### HasReplacesBadgeId

`func (o *TenantBadge) HasReplacesBadgeId() bool`

HasReplacesBadgeId returns a boolean if a field has been set.

### SetReplacesBadgeIdNil

`func (o *TenantBadge) SetReplacesBadgeIdNil(b bool)`

 SetReplacesBadgeIdNil sets the value for ReplacesBadgeId to be an explicit nil

### UnsetReplacesBadgeId
`func (o *TenantBadge) UnsetReplacesBadgeId()`

UnsetReplacesBadgeId ensures that no value is present for ReplacesBadgeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


