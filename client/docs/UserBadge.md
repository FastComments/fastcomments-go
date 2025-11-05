# UserBadge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**BadgeId** | **string** |  | 
**FromTenantId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Type** | **int32** |  | 
**Threshold** | **int64** |  | 
**Description** | **string** |  | 
**DisplayLabel** | **string** |  | 
**DisplaySrc** | Pointer to **NullableString** |  | [optional] 
**BackgroundColor** | Pointer to **NullableString** |  | [optional] 
**BorderColor** | Pointer to **NullableString** |  | [optional] 
**TextColor** | Pointer to **NullableString** |  | [optional] 
**CssClass** | Pointer to **NullableString** |  | [optional] 
**VeteranUserThresholdMillis** | **int64** |  | 
**DisplayedOnComments** | **bool** |  | 
**ReceivedAt** | **time.Time** |  | 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserBadge

`func NewUserBadge(id string, userId string, badgeId string, fromTenantId string, createdAt time.Time, type_ int32, threshold int64, description string, displayLabel string, veteranUserThresholdMillis int64, displayedOnComments bool, receivedAt time.Time, ) *UserBadge`

NewUserBadge instantiates a new UserBadge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBadgeWithDefaults

`func NewUserBadgeWithDefaults() *UserBadge`

NewUserBadgeWithDefaults instantiates a new UserBadge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBadge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBadge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBadge) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *UserBadge) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserBadge) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserBadge) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBadgeId

`func (o *UserBadge) GetBadgeId() string`

GetBadgeId returns the BadgeId field if non-nil, zero value otherwise.

### GetBadgeIdOk

`func (o *UserBadge) GetBadgeIdOk() (*string, bool)`

GetBadgeIdOk returns a tuple with the BadgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeId

`func (o *UserBadge) SetBadgeId(v string)`

SetBadgeId sets BadgeId field to given value.


### GetFromTenantId

`func (o *UserBadge) GetFromTenantId() string`

GetFromTenantId returns the FromTenantId field if non-nil, zero value otherwise.

### GetFromTenantIdOk

`func (o *UserBadge) GetFromTenantIdOk() (*string, bool)`

GetFromTenantIdOk returns a tuple with the FromTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTenantId

`func (o *UserBadge) SetFromTenantId(v string)`

SetFromTenantId sets FromTenantId field to given value.


### GetCreatedAt

`func (o *UserBadge) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserBadge) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserBadge) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetType

`func (o *UserBadge) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserBadge) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserBadge) SetType(v int32)`

SetType sets Type field to given value.


### GetThreshold

`func (o *UserBadge) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UserBadge) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UserBadge) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.


### GetDescription

`func (o *UserBadge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserBadge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserBadge) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayLabel

`func (o *UserBadge) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *UserBadge) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *UserBadge) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.


### GetDisplaySrc

`func (o *UserBadge) GetDisplaySrc() string`

GetDisplaySrc returns the DisplaySrc field if non-nil, zero value otherwise.

### GetDisplaySrcOk

`func (o *UserBadge) GetDisplaySrcOk() (*string, bool)`

GetDisplaySrcOk returns a tuple with the DisplaySrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySrc

`func (o *UserBadge) SetDisplaySrc(v string)`

SetDisplaySrc sets DisplaySrc field to given value.

### HasDisplaySrc

`func (o *UserBadge) HasDisplaySrc() bool`

HasDisplaySrc returns a boolean if a field has been set.

### SetDisplaySrcNil

`func (o *UserBadge) SetDisplaySrcNil(b bool)`

 SetDisplaySrcNil sets the value for DisplaySrc to be an explicit nil

### UnsetDisplaySrc
`func (o *UserBadge) UnsetDisplaySrc()`

UnsetDisplaySrc ensures that no value is present for DisplaySrc, not even an explicit nil
### GetBackgroundColor

`func (o *UserBadge) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *UserBadge) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *UserBadge) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *UserBadge) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *UserBadge) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *UserBadge) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBorderColor

`func (o *UserBadge) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *UserBadge) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *UserBadge) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *UserBadge) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### SetBorderColorNil

`func (o *UserBadge) SetBorderColorNil(b bool)`

 SetBorderColorNil sets the value for BorderColor to be an explicit nil

### UnsetBorderColor
`func (o *UserBadge) UnsetBorderColor()`

UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
### GetTextColor

`func (o *UserBadge) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *UserBadge) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *UserBadge) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *UserBadge) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### SetTextColorNil

`func (o *UserBadge) SetTextColorNil(b bool)`

 SetTextColorNil sets the value for TextColor to be an explicit nil

### UnsetTextColor
`func (o *UserBadge) UnsetTextColor()`

UnsetTextColor ensures that no value is present for TextColor, not even an explicit nil
### GetCssClass

`func (o *UserBadge) GetCssClass() string`

GetCssClass returns the CssClass field if non-nil, zero value otherwise.

### GetCssClassOk

`func (o *UserBadge) GetCssClassOk() (*string, bool)`

GetCssClassOk returns a tuple with the CssClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssClass

`func (o *UserBadge) SetCssClass(v string)`

SetCssClass sets CssClass field to given value.

### HasCssClass

`func (o *UserBadge) HasCssClass() bool`

HasCssClass returns a boolean if a field has been set.

### SetCssClassNil

`func (o *UserBadge) SetCssClassNil(b bool)`

 SetCssClassNil sets the value for CssClass to be an explicit nil

### UnsetCssClass
`func (o *UserBadge) UnsetCssClass()`

UnsetCssClass ensures that no value is present for CssClass, not even an explicit nil
### GetVeteranUserThresholdMillis

`func (o *UserBadge) GetVeteranUserThresholdMillis() int64`

GetVeteranUserThresholdMillis returns the VeteranUserThresholdMillis field if non-nil, zero value otherwise.

### GetVeteranUserThresholdMillisOk

`func (o *UserBadge) GetVeteranUserThresholdMillisOk() (*int64, bool)`

GetVeteranUserThresholdMillisOk returns a tuple with the VeteranUserThresholdMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVeteranUserThresholdMillis

`func (o *UserBadge) SetVeteranUserThresholdMillis(v int64)`

SetVeteranUserThresholdMillis sets VeteranUserThresholdMillis field to given value.


### GetDisplayedOnComments

`func (o *UserBadge) GetDisplayedOnComments() bool`

GetDisplayedOnComments returns the DisplayedOnComments field if non-nil, zero value otherwise.

### GetDisplayedOnCommentsOk

`func (o *UserBadge) GetDisplayedOnCommentsOk() (*bool, bool)`

GetDisplayedOnCommentsOk returns a tuple with the DisplayedOnComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedOnComments

`func (o *UserBadge) SetDisplayedOnComments(v bool)`

SetDisplayedOnComments sets DisplayedOnComments field to given value.


### GetReceivedAt

`func (o *UserBadge) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *UserBadge) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *UserBadge) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetOrder

`func (o *UserBadge) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UserBadge) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UserBadge) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UserBadge) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


