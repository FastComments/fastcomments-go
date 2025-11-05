# UpdateFeedPostParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**ContentHTML** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**[]FeedPostMediaItem**](FeedPostMediaItem.md) |  | [optional] 
**Links** | Pointer to [**[]FeedPostLink**](FeedPostLink.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Meta** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewUpdateFeedPostParams

`func NewUpdateFeedPostParams() *UpdateFeedPostParams`

NewUpdateFeedPostParams instantiates a new UpdateFeedPostParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeedPostParamsWithDefaults

`func NewUpdateFeedPostParamsWithDefaults() *UpdateFeedPostParams`

NewUpdateFeedPostParamsWithDefaults instantiates a new UpdateFeedPostParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateFeedPostParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateFeedPostParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateFeedPostParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateFeedPostParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContentHTML

`func (o *UpdateFeedPostParams) GetContentHTML() string`

GetContentHTML returns the ContentHTML field if non-nil, zero value otherwise.

### GetContentHTMLOk

`func (o *UpdateFeedPostParams) GetContentHTMLOk() (*string, bool)`

GetContentHTMLOk returns a tuple with the ContentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHTML

`func (o *UpdateFeedPostParams) SetContentHTML(v string)`

SetContentHTML sets ContentHTML field to given value.

### HasContentHTML

`func (o *UpdateFeedPostParams) HasContentHTML() bool`

HasContentHTML returns a boolean if a field has been set.

### GetMedia

`func (o *UpdateFeedPostParams) GetMedia() []FeedPostMediaItem`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *UpdateFeedPostParams) GetMediaOk() (*[]FeedPostMediaItem, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *UpdateFeedPostParams) SetMedia(v []FeedPostMediaItem)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *UpdateFeedPostParams) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetLinks

`func (o *UpdateFeedPostParams) GetLinks() []FeedPostLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateFeedPostParams) GetLinksOk() (*[]FeedPostLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateFeedPostParams) SetLinks(v []FeedPostLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateFeedPostParams) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTags

`func (o *UpdateFeedPostParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateFeedPostParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateFeedPostParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateFeedPostParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateFeedPostParams) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateFeedPostParams) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateFeedPostParams) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateFeedPostParams) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


