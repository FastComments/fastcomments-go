# GetGifsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[][]GifSearchResponseImagesInnerInner**]([]GifSearchResponseImagesInnerInner.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Code** | **string** |  | 

## Methods

### NewGetGifsSearchResponse

`func NewGetGifsSearchResponse(images [][]GifSearchResponseImagesInnerInner, status APIStatus, code string, ) *GetGifsSearchResponse`

NewGetGifsSearchResponse instantiates a new GetGifsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGifsSearchResponseWithDefaults

`func NewGetGifsSearchResponseWithDefaults() *GetGifsSearchResponse`

NewGetGifsSearchResponseWithDefaults instantiates a new GetGifsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *GetGifsSearchResponse) GetImages() [][]GifSearchResponseImagesInnerInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetGifsSearchResponse) GetImagesOk() (*[][]GifSearchResponseImagesInnerInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetGifsSearchResponse) SetImages(v [][]GifSearchResponseImagesInnerInner)`

SetImages sets Images field to given value.


### GetStatus

`func (o *GetGifsSearchResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGifsSearchResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGifsSearchResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCode

`func (o *GetGifsSearchResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetGifsSearchResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetGifsSearchResponse) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


