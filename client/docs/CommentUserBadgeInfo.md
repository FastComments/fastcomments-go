# CommentUserBadgeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **int32** |  | 
**Description** | **string** |  | 
**DisplayLabel** | Pointer to **NullableString** |  | [optional] 
**DisplaySrc** | Pointer to **NullableString** |  | [optional] 
**BackgroundColor** | Pointer to **NullableString** |  | [optional] 
**BorderColor** | Pointer to **NullableString** |  | [optional] 
**TextColor** | Pointer to **NullableString** |  | [optional] 
**CssClass** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCommentUserBadgeInfo

`func NewCommentUserBadgeInfo(id string, type_ int32, description string, ) *CommentUserBadgeInfo`

NewCommentUserBadgeInfo instantiates a new CommentUserBadgeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentUserBadgeInfoWithDefaults

`func NewCommentUserBadgeInfoWithDefaults() *CommentUserBadgeInfo`

NewCommentUserBadgeInfoWithDefaults instantiates a new CommentUserBadgeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentUserBadgeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentUserBadgeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentUserBadgeInfo) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CommentUserBadgeInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommentUserBadgeInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommentUserBadgeInfo) SetType(v int32)`

SetType sets Type field to given value.


### GetDescription

`func (o *CommentUserBadgeInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommentUserBadgeInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommentUserBadgeInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayLabel

`func (o *CommentUserBadgeInfo) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *CommentUserBadgeInfo) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *CommentUserBadgeInfo) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *CommentUserBadgeInfo) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *CommentUserBadgeInfo) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *CommentUserBadgeInfo) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetDisplaySrc

`func (o *CommentUserBadgeInfo) GetDisplaySrc() string`

GetDisplaySrc returns the DisplaySrc field if non-nil, zero value otherwise.

### GetDisplaySrcOk

`func (o *CommentUserBadgeInfo) GetDisplaySrcOk() (*string, bool)`

GetDisplaySrcOk returns a tuple with the DisplaySrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySrc

`func (o *CommentUserBadgeInfo) SetDisplaySrc(v string)`

SetDisplaySrc sets DisplaySrc field to given value.

### HasDisplaySrc

`func (o *CommentUserBadgeInfo) HasDisplaySrc() bool`

HasDisplaySrc returns a boolean if a field has been set.

### SetDisplaySrcNil

`func (o *CommentUserBadgeInfo) SetDisplaySrcNil(b bool)`

 SetDisplaySrcNil sets the value for DisplaySrc to be an explicit nil

### UnsetDisplaySrc
`func (o *CommentUserBadgeInfo) UnsetDisplaySrc()`

UnsetDisplaySrc ensures that no value is present for DisplaySrc, not even an explicit nil
### GetBackgroundColor

`func (o *CommentUserBadgeInfo) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *CommentUserBadgeInfo) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *CommentUserBadgeInfo) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *CommentUserBadgeInfo) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *CommentUserBadgeInfo) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *CommentUserBadgeInfo) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBorderColor

`func (o *CommentUserBadgeInfo) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *CommentUserBadgeInfo) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *CommentUserBadgeInfo) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *CommentUserBadgeInfo) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### SetBorderColorNil

`func (o *CommentUserBadgeInfo) SetBorderColorNil(b bool)`

 SetBorderColorNil sets the value for BorderColor to be an explicit nil

### UnsetBorderColor
`func (o *CommentUserBadgeInfo) UnsetBorderColor()`

UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
### GetTextColor

`func (o *CommentUserBadgeInfo) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *CommentUserBadgeInfo) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *CommentUserBadgeInfo) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *CommentUserBadgeInfo) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### SetTextColorNil

`func (o *CommentUserBadgeInfo) SetTextColorNil(b bool)`

 SetTextColorNil sets the value for TextColor to be an explicit nil

### UnsetTextColor
`func (o *CommentUserBadgeInfo) UnsetTextColor()`

UnsetTextColor ensures that no value is present for TextColor, not even an explicit nil
### GetCssClass

`func (o *CommentUserBadgeInfo) GetCssClass() string`

GetCssClass returns the CssClass field if non-nil, zero value otherwise.

### GetCssClassOk

`func (o *CommentUserBadgeInfo) GetCssClassOk() (*string, bool)`

GetCssClassOk returns a tuple with the CssClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssClass

`func (o *CommentUserBadgeInfo) SetCssClass(v string)`

SetCssClass sets CssClass field to given value.

### HasCssClass

`func (o *CommentUserBadgeInfo) HasCssClass() bool`

HasCssClass returns a boolean if a field has been set.

### SetCssClassNil

`func (o *CommentUserBadgeInfo) SetCssClassNil(b bool)`

 SetCssClassNil sets the value for CssClass to be an explicit nil

### UnsetCssClass
`func (o *CommentUserBadgeInfo) UnsetCssClass()`

UnsetCssClass ensures that no value is present for CssClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


