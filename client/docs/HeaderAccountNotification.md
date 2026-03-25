# HeaderAccountNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Message** | **string** |  | 
**MessagesByLocale** | **map[string]string** | Construct a type with a set of properties K of type T | 
**Dates** | **map[string]string** | Construct a type with a set of properties K of type T | 
**Severity** | **string** |  | 
**LinkUrl** | **NullableString** |  | 
**LinkText** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewHeaderAccountNotification

`func NewHeaderAccountNotification(id string, title string, message string, messagesByLocale map[string]string, dates map[string]string, severity string, linkUrl NullableString, linkText NullableString, createdAt time.Time, ) *HeaderAccountNotification`

NewHeaderAccountNotification instantiates a new HeaderAccountNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderAccountNotificationWithDefaults

`func NewHeaderAccountNotificationWithDefaults() *HeaderAccountNotification`

NewHeaderAccountNotificationWithDefaults instantiates a new HeaderAccountNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeaderAccountNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeaderAccountNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeaderAccountNotification) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *HeaderAccountNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HeaderAccountNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HeaderAccountNotification) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *HeaderAccountNotification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HeaderAccountNotification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HeaderAccountNotification) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessagesByLocale

`func (o *HeaderAccountNotification) GetMessagesByLocale() map[string]string`

GetMessagesByLocale returns the MessagesByLocale field if non-nil, zero value otherwise.

### GetMessagesByLocaleOk

`func (o *HeaderAccountNotification) GetMessagesByLocaleOk() (*map[string]string, bool)`

GetMessagesByLocaleOk returns a tuple with the MessagesByLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesByLocale

`func (o *HeaderAccountNotification) SetMessagesByLocale(v map[string]string)`

SetMessagesByLocale sets MessagesByLocale field to given value.


### SetMessagesByLocaleNil

`func (o *HeaderAccountNotification) SetMessagesByLocaleNil(b bool)`

 SetMessagesByLocaleNil sets the value for MessagesByLocale to be an explicit nil

### UnsetMessagesByLocale
`func (o *HeaderAccountNotification) UnsetMessagesByLocale()`

UnsetMessagesByLocale ensures that no value is present for MessagesByLocale, not even an explicit nil
### GetDates

`func (o *HeaderAccountNotification) GetDates() map[string]string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *HeaderAccountNotification) GetDatesOk() (*map[string]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *HeaderAccountNotification) SetDates(v map[string]string)`

SetDates sets Dates field to given value.


### SetDatesNil

`func (o *HeaderAccountNotification) SetDatesNil(b bool)`

 SetDatesNil sets the value for Dates to be an explicit nil

### UnsetDates
`func (o *HeaderAccountNotification) UnsetDates()`

UnsetDates ensures that no value is present for Dates, not even an explicit nil
### GetSeverity

`func (o *HeaderAccountNotification) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *HeaderAccountNotification) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *HeaderAccountNotification) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetLinkUrl

`func (o *HeaderAccountNotification) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *HeaderAccountNotification) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *HeaderAccountNotification) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.


### SetLinkUrlNil

`func (o *HeaderAccountNotification) SetLinkUrlNil(b bool)`

 SetLinkUrlNil sets the value for LinkUrl to be an explicit nil

### UnsetLinkUrl
`func (o *HeaderAccountNotification) UnsetLinkUrl()`

UnsetLinkUrl ensures that no value is present for LinkUrl, not even an explicit nil
### GetLinkText

`func (o *HeaderAccountNotification) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *HeaderAccountNotification) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *HeaderAccountNotification) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.


### SetLinkTextNil

`func (o *HeaderAccountNotification) SetLinkTextNil(b bool)`

 SetLinkTextNil sets the value for LinkText to be an explicit nil

### UnsetLinkText
`func (o *HeaderAccountNotification) UnsetLinkText()`

UnsetLinkText ensures that no value is present for LinkText, not even an explicit nil
### GetCreatedAt

`func (o *HeaderAccountNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HeaderAccountNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HeaderAccountNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


