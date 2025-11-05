# UploadImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Url** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**[]MediaAsset**](MediaAsset.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewUploadImageResponse

`func NewUploadImageResponse(status APIStatus, ) *UploadImageResponse`

NewUploadImageResponse instantiates a new UploadImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadImageResponseWithDefaults

`func NewUploadImageResponseWithDefaults() *UploadImageResponse`

NewUploadImageResponseWithDefaults instantiates a new UploadImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UploadImageResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadImageResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadImageResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *UploadImageResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadImageResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadImageResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadImageResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMedia

`func (o *UploadImageResponse) GetMedia() []MediaAsset`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *UploadImageResponse) GetMediaOk() (*[]MediaAsset, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *UploadImageResponse) SetMedia(v []MediaAsset)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *UploadImageResponse) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetReason

`func (o *UploadImageResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UploadImageResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UploadImageResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UploadImageResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *UploadImageResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UploadImageResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UploadImageResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UploadImageResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


