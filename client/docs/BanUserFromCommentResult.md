# BanUserFromCommentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Changelog** | Pointer to [**APIBanUserChangeLog**](APIBanUserChangeLog.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewBanUserFromCommentResult

`func NewBanUserFromCommentResult(status string, ) *BanUserFromCommentResult`

NewBanUserFromCommentResult instantiates a new BanUserFromCommentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanUserFromCommentResultWithDefaults

`func NewBanUserFromCommentResultWithDefaults() *BanUserFromCommentResult`

NewBanUserFromCommentResultWithDefaults instantiates a new BanUserFromCommentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BanUserFromCommentResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BanUserFromCommentResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BanUserFromCommentResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChangelog

`func (o *BanUserFromCommentResult) GetChangelog() APIBanUserChangeLog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *BanUserFromCommentResult) GetChangelogOk() (*APIBanUserChangeLog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *BanUserFromCommentResult) SetChangelog(v APIBanUserChangeLog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *BanUserFromCommentResult) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetCode

`func (o *BanUserFromCommentResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BanUserFromCommentResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BanUserFromCommentResult) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BanUserFromCommentResult) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *BanUserFromCommentResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BanUserFromCommentResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BanUserFromCommentResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BanUserFromCommentResult) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


