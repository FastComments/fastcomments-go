# ModerationAPIGetCommentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Translations** | **map[string]interface{}** |  | 
**Comments** | [**[]ModerationAPIComment**](ModerationAPIComment.md) |  | 
**ModerationFilter** | Pointer to [**ModerationFilter**](ModerationFilter.md) |  | [optional] 

## Methods

### NewModerationAPIGetCommentsResponse

`func NewModerationAPIGetCommentsResponse(status APIStatus, translations map[string]interface{}, comments []ModerationAPIComment, ) *ModerationAPIGetCommentsResponse`

NewModerationAPIGetCommentsResponse instantiates a new ModerationAPIGetCommentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPIGetCommentsResponseWithDefaults

`func NewModerationAPIGetCommentsResponseWithDefaults() *ModerationAPIGetCommentsResponse`

NewModerationAPIGetCommentsResponseWithDefaults instantiates a new ModerationAPIGetCommentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModerationAPIGetCommentsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationAPIGetCommentsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationAPIGetCommentsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTranslations

`func (o *ModerationAPIGetCommentsResponse) GetTranslations() map[string]interface{}`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *ModerationAPIGetCommentsResponse) GetTranslationsOk() (*map[string]interface{}, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *ModerationAPIGetCommentsResponse) SetTranslations(v map[string]interface{})`

SetTranslations sets Translations field to given value.


### GetComments

`func (o *ModerationAPIGetCommentsResponse) GetComments() []ModerationAPIComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModerationAPIGetCommentsResponse) GetCommentsOk() (*[]ModerationAPIComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModerationAPIGetCommentsResponse) SetComments(v []ModerationAPIComment)`

SetComments sets Comments field to given value.


### GetModerationFilter

`func (o *ModerationAPIGetCommentsResponse) GetModerationFilter() ModerationFilter`

GetModerationFilter returns the ModerationFilter field if non-nil, zero value otherwise.

### GetModerationFilterOk

`func (o *ModerationAPIGetCommentsResponse) GetModerationFilterOk() (*ModerationFilter, bool)`

GetModerationFilterOk returns a tuple with the ModerationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationFilter

`func (o *ModerationAPIGetCommentsResponse) SetModerationFilter(v ModerationFilter)`

SetModerationFilter sets ModerationFilter field to given value.

### HasModerationFilter

`func (o *ModerationAPIGetCommentsResponse) HasModerationFilter() bool`

HasModerationFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


