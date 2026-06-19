# GifSearchInternalError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGifSearchInternalError

`func NewGifSearchInternalError(code string, status APIStatus, ) *GifSearchInternalError`

NewGifSearchInternalError instantiates a new GifSearchInternalError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGifSearchInternalErrorWithDefaults

`func NewGifSearchInternalErrorWithDefaults() *GifSearchInternalError`

NewGifSearchInternalErrorWithDefaults instantiates a new GifSearchInternalError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GifSearchInternalError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GifSearchInternalError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GifSearchInternalError) SetCode(v string)`

SetCode sets Code field to given value.


### GetStatus

`func (o *GifSearchInternalError) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GifSearchInternalError) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GifSearchInternalError) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


