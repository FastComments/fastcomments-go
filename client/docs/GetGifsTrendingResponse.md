# GetGifsTrendingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[][]GifSearchResponseImagesInnerInner**]([]GifSearchResponseImagesInnerInner.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Code** | **string** |  | 

## Methods

### NewGetGifsTrendingResponse

`func NewGetGifsTrendingResponse(images [][]GifSearchResponseImagesInnerInner, status APIStatus, code string, ) *GetGifsTrendingResponse`

NewGetGifsTrendingResponse instantiates a new GetGifsTrendingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGifsTrendingResponseWithDefaults

`func NewGetGifsTrendingResponseWithDefaults() *GetGifsTrendingResponse`

NewGetGifsTrendingResponseWithDefaults instantiates a new GetGifsTrendingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *GetGifsTrendingResponse) GetImages() [][]GifSearchResponseImagesInnerInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetGifsTrendingResponse) GetImagesOk() (*[][]GifSearchResponseImagesInnerInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetGifsTrendingResponse) SetImages(v [][]GifSearchResponseImagesInnerInner)`

SetImages sets Images field to given value.


### GetStatus

`func (o *GetGifsTrendingResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGifsTrendingResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGifsTrendingResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCode

`func (o *GetGifsTrendingResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetGifsTrendingResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetGifsTrendingResponse) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


