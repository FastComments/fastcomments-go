# FeedPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**FromUserId** | Pointer to **string** |  | [optional] 
**FromUserDisplayName** | Pointer to **NullableString** |  | [optional] 
**FromUserAvatar** | Pointer to **NullableString** |  | [optional] 
**FromIpHash** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Weight** | Pointer to **float64** |  | [optional] 
**Meta** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**ContentHTML** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**[]FeedPostMediaItem**](FeedPostMediaItem.md) |  | [optional] 
**Links** | Pointer to [**[]FeedPostLink**](FeedPostLink.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Reacts** | Pointer to **map[string]int32** |  | [optional] 
**CommentCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFeedPost

`func NewFeedPost(id string, tenantId string, createdAt time.Time, ) *FeedPost`

NewFeedPost instantiates a new FeedPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedPostWithDefaults

`func NewFeedPostWithDefaults() *FeedPost`

NewFeedPostWithDefaults instantiates a new FeedPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeedPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedPost) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *FeedPost) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FeedPost) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FeedPost) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTitle

`func (o *FeedPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeedPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FeedPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FeedPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFromUserId

`func (o *FeedPost) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *FeedPost) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *FeedPost) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *FeedPost) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetFromUserDisplayName

`func (o *FeedPost) GetFromUserDisplayName() string`

GetFromUserDisplayName returns the FromUserDisplayName field if non-nil, zero value otherwise.

### GetFromUserDisplayNameOk

`func (o *FeedPost) GetFromUserDisplayNameOk() (*string, bool)`

GetFromUserDisplayNameOk returns a tuple with the FromUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserDisplayName

`func (o *FeedPost) SetFromUserDisplayName(v string)`

SetFromUserDisplayName sets FromUserDisplayName field to given value.

### HasFromUserDisplayName

`func (o *FeedPost) HasFromUserDisplayName() bool`

HasFromUserDisplayName returns a boolean if a field has been set.

### SetFromUserDisplayNameNil

`func (o *FeedPost) SetFromUserDisplayNameNil(b bool)`

 SetFromUserDisplayNameNil sets the value for FromUserDisplayName to be an explicit nil

### UnsetFromUserDisplayName
`func (o *FeedPost) UnsetFromUserDisplayName()`

UnsetFromUserDisplayName ensures that no value is present for FromUserDisplayName, not even an explicit nil
### GetFromUserAvatar

`func (o *FeedPost) GetFromUserAvatar() string`

GetFromUserAvatar returns the FromUserAvatar field if non-nil, zero value otherwise.

### GetFromUserAvatarOk

`func (o *FeedPost) GetFromUserAvatarOk() (*string, bool)`

GetFromUserAvatarOk returns a tuple with the FromUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserAvatar

`func (o *FeedPost) SetFromUserAvatar(v string)`

SetFromUserAvatar sets FromUserAvatar field to given value.

### HasFromUserAvatar

`func (o *FeedPost) HasFromUserAvatar() bool`

HasFromUserAvatar returns a boolean if a field has been set.

### SetFromUserAvatarNil

`func (o *FeedPost) SetFromUserAvatarNil(b bool)`

 SetFromUserAvatarNil sets the value for FromUserAvatar to be an explicit nil

### UnsetFromUserAvatar
`func (o *FeedPost) UnsetFromUserAvatar()`

UnsetFromUserAvatar ensures that no value is present for FromUserAvatar, not even an explicit nil
### GetFromIpHash

`func (o *FeedPost) GetFromIpHash() string`

GetFromIpHash returns the FromIpHash field if non-nil, zero value otherwise.

### GetFromIpHashOk

`func (o *FeedPost) GetFromIpHashOk() (*string, bool)`

GetFromIpHashOk returns a tuple with the FromIpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIpHash

`func (o *FeedPost) SetFromIpHash(v string)`

SetFromIpHash sets FromIpHash field to given value.

### HasFromIpHash

`func (o *FeedPost) HasFromIpHash() bool`

HasFromIpHash returns a boolean if a field has been set.

### GetTags

`func (o *FeedPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeedPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeedPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeedPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWeight

`func (o *FeedPost) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FeedPost) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FeedPost) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *FeedPost) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMeta

`func (o *FeedPost) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FeedPost) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FeedPost) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FeedPost) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetContentHTML

`func (o *FeedPost) GetContentHTML() string`

GetContentHTML returns the ContentHTML field if non-nil, zero value otherwise.

### GetContentHTMLOk

`func (o *FeedPost) GetContentHTMLOk() (*string, bool)`

GetContentHTMLOk returns a tuple with the ContentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHTML

`func (o *FeedPost) SetContentHTML(v string)`

SetContentHTML sets ContentHTML field to given value.

### HasContentHTML

`func (o *FeedPost) HasContentHTML() bool`

HasContentHTML returns a boolean if a field has been set.

### GetMedia

`func (o *FeedPost) GetMedia() []FeedPostMediaItem`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *FeedPost) GetMediaOk() (*[]FeedPostMediaItem, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *FeedPost) SetMedia(v []FeedPostMediaItem)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *FeedPost) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetLinks

`func (o *FeedPost) GetLinks() []FeedPostLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeedPost) GetLinksOk() (*[]FeedPostLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeedPost) SetLinks(v []FeedPostLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FeedPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeedPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReacts

`func (o *FeedPost) GetReacts() map[string]int32`

GetReacts returns the Reacts field if non-nil, zero value otherwise.

### GetReactsOk

`func (o *FeedPost) GetReactsOk() (*map[string]int32, bool)`

GetReactsOk returns a tuple with the Reacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReacts

`func (o *FeedPost) SetReacts(v map[string]int32)`

SetReacts sets Reacts field to given value.

### HasReacts

`func (o *FeedPost) HasReacts() bool`

HasReacts returns a boolean if a field has been set.

### GetCommentCount

`func (o *FeedPost) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *FeedPost) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *FeedPost) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *FeedPost) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### SetCommentCountNil

`func (o *FeedPost) SetCommentCountNil(b bool)`

 SetCommentCountNil sets the value for CommentCount to be an explicit nil

### UnsetCommentCount
`func (o *FeedPost) UnsetCommentCount()`

UnsetCommentCount ensures that no value is present for CommentCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


