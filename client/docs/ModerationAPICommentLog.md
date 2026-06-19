# ModerationAPICommentLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**Username** | Pointer to **string** |  | [optional] 
**ActionName** | **string** |  | 
**MessageHTML** | **string** |  | 

## Methods

### NewModerationAPICommentLog

`func NewModerationAPICommentLog(date time.Time, actionName string, messageHTML string, ) *ModerationAPICommentLog`

NewModerationAPICommentLog instantiates a new ModerationAPICommentLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPICommentLogWithDefaults

`func NewModerationAPICommentLogWithDefaults() *ModerationAPICommentLog`

NewModerationAPICommentLogWithDefaults instantiates a new ModerationAPICommentLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ModerationAPICommentLog) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ModerationAPICommentLog) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ModerationAPICommentLog) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetUsername

`func (o *ModerationAPICommentLog) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModerationAPICommentLog) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModerationAPICommentLog) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModerationAPICommentLog) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetActionName

`func (o *ModerationAPICommentLog) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *ModerationAPICommentLog) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *ModerationAPICommentLog) SetActionName(v string)`

SetActionName sets ActionName field to given value.


### GetMessageHTML

`func (o *ModerationAPICommentLog) GetMessageHTML() string`

GetMessageHTML returns the MessageHTML field if non-nil, zero value otherwise.

### GetMessageHTMLOk

`func (o *ModerationAPICommentLog) GetMessageHTMLOk() (*string, bool)`

GetMessageHTMLOk returns a tuple with the MessageHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHTML

`func (o *ModerationAPICommentLog) SetMessageHTML(v string)`

SetMessageHTML sets MessageHTML field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


