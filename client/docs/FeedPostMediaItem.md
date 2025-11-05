# FeedPostMediaItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**LinkUrl** | Pointer to **string** |  | [optional] 
**Sizes** | [**[]FeedPostMediaItemAsset**](FeedPostMediaItemAsset.md) |  | 

## Methods

### NewFeedPostMediaItem

`func NewFeedPostMediaItem(sizes []FeedPostMediaItemAsset, ) *FeedPostMediaItem`

NewFeedPostMediaItem instantiates a new FeedPostMediaItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedPostMediaItemWithDefaults

`func NewFeedPostMediaItemWithDefaults() *FeedPostMediaItem`

NewFeedPostMediaItemWithDefaults instantiates a new FeedPostMediaItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FeedPostMediaItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeedPostMediaItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FeedPostMediaItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FeedPostMediaItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLinkUrl

`func (o *FeedPostMediaItem) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *FeedPostMediaItem) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *FeedPostMediaItem) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *FeedPostMediaItem) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetSizes

`func (o *FeedPostMediaItem) GetSizes() []FeedPostMediaItemAsset`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *FeedPostMediaItem) GetSizesOk() (*[]FeedPostMediaItemAsset, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *FeedPostMediaItem) SetSizes(v []FeedPostMediaItemAsset)`

SetSizes sets Sizes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


