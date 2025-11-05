# CreateFeedPostParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**ContentHTML** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**[]FeedPostMediaItem**](FeedPostMediaItem.md) |  | [optional] 
**Links** | Pointer to [**[]FeedPostLink**](FeedPostLink.md) |  | [optional] 
**FromUserId** | Pointer to **string** |  | [optional] 
**FromUserDisplayName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Meta** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewCreateFeedPostParams

`func NewCreateFeedPostParams() *CreateFeedPostParams`

NewCreateFeedPostParams instantiates a new CreateFeedPostParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeedPostParamsWithDefaults

`func NewCreateFeedPostParamsWithDefaults() *CreateFeedPostParams`

NewCreateFeedPostParamsWithDefaults instantiates a new CreateFeedPostParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateFeedPostParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFeedPostParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFeedPostParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateFeedPostParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContentHTML

`func (o *CreateFeedPostParams) GetContentHTML() string`

GetContentHTML returns the ContentHTML field if non-nil, zero value otherwise.

### GetContentHTMLOk

`func (o *CreateFeedPostParams) GetContentHTMLOk() (*string, bool)`

GetContentHTMLOk returns a tuple with the ContentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHTML

`func (o *CreateFeedPostParams) SetContentHTML(v string)`

SetContentHTML sets ContentHTML field to given value.

### HasContentHTML

`func (o *CreateFeedPostParams) HasContentHTML() bool`

HasContentHTML returns a boolean if a field has been set.

### GetMedia

`func (o *CreateFeedPostParams) GetMedia() []FeedPostMediaItem`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *CreateFeedPostParams) GetMediaOk() (*[]FeedPostMediaItem, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *CreateFeedPostParams) SetMedia(v []FeedPostMediaItem)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *CreateFeedPostParams) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetLinks

`func (o *CreateFeedPostParams) GetLinks() []FeedPostLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateFeedPostParams) GetLinksOk() (*[]FeedPostLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateFeedPostParams) SetLinks(v []FeedPostLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateFeedPostParams) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetFromUserId

`func (o *CreateFeedPostParams) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *CreateFeedPostParams) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *CreateFeedPostParams) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *CreateFeedPostParams) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetFromUserDisplayName

`func (o *CreateFeedPostParams) GetFromUserDisplayName() string`

GetFromUserDisplayName returns the FromUserDisplayName field if non-nil, zero value otherwise.

### GetFromUserDisplayNameOk

`func (o *CreateFeedPostParams) GetFromUserDisplayNameOk() (*string, bool)`

GetFromUserDisplayNameOk returns a tuple with the FromUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserDisplayName

`func (o *CreateFeedPostParams) SetFromUserDisplayName(v string)`

SetFromUserDisplayName sets FromUserDisplayName field to given value.

### HasFromUserDisplayName

`func (o *CreateFeedPostParams) HasFromUserDisplayName() bool`

HasFromUserDisplayName returns a boolean if a field has been set.

### GetTags

`func (o *CreateFeedPostParams) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateFeedPostParams) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateFeedPostParams) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateFeedPostParams) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMeta

`func (o *CreateFeedPostParams) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateFeedPostParams) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateFeedPostParams) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateFeedPostParams) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


