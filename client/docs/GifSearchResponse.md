# GifSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[][]GifSearchResponseImagesInnerInner**]([]GifSearchResponseImagesInnerInner.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGifSearchResponse

`func NewGifSearchResponse(images [][]GifSearchResponseImagesInnerInner, status APIStatus, ) *GifSearchResponse`

NewGifSearchResponse instantiates a new GifSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGifSearchResponseWithDefaults

`func NewGifSearchResponseWithDefaults() *GifSearchResponse`

NewGifSearchResponseWithDefaults instantiates a new GifSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *GifSearchResponse) GetImages() [][]GifSearchResponseImagesInnerInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GifSearchResponse) GetImagesOk() (*[][]GifSearchResponseImagesInnerInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GifSearchResponse) SetImages(v [][]GifSearchResponseImagesInnerInner)`

SetImages sets Images field to given value.


### GetStatus

`func (o *GifSearchResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GifSearchResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GifSearchResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


