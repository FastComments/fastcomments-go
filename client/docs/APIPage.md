# APIPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsClosed** | Pointer to **bool** |  | [optional] 
**AccessibleByGroupIds** | Pointer to **[]string** |  | [optional] 
**RootCommentCount** | **int64** |  | 
**CommentCount** | **int64** |  | 
**CreatedAt** | **time.Time** |  | 
**Title** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**UrlId** | **string** |  | 
**Id** | **string** |  | 

## Methods

### NewAPIPage

`func NewAPIPage(rootCommentCount int64, commentCount int64, createdAt time.Time, title string, urlId string, id string, ) *APIPage`

NewAPIPage instantiates a new APIPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIPageWithDefaults

`func NewAPIPageWithDefaults() *APIPage`

NewAPIPageWithDefaults instantiates a new APIPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsClosed

`func (o *APIPage) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *APIPage) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *APIPage) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *APIPage) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetAccessibleByGroupIds

`func (o *APIPage) GetAccessibleByGroupIds() []string`

GetAccessibleByGroupIds returns the AccessibleByGroupIds field if non-nil, zero value otherwise.

### GetAccessibleByGroupIdsOk

`func (o *APIPage) GetAccessibleByGroupIdsOk() (*[]string, bool)`

GetAccessibleByGroupIdsOk returns a tuple with the AccessibleByGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleByGroupIds

`func (o *APIPage) SetAccessibleByGroupIds(v []string)`

SetAccessibleByGroupIds sets AccessibleByGroupIds field to given value.

### HasAccessibleByGroupIds

`func (o *APIPage) HasAccessibleByGroupIds() bool`

HasAccessibleByGroupIds returns a boolean if a field has been set.

### GetRootCommentCount

`func (o *APIPage) GetRootCommentCount() int64`

GetRootCommentCount returns the RootCommentCount field if non-nil, zero value otherwise.

### GetRootCommentCountOk

`func (o *APIPage) GetRootCommentCountOk() (*int64, bool)`

GetRootCommentCountOk returns a tuple with the RootCommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCommentCount

`func (o *APIPage) SetRootCommentCount(v int64)`

SetRootCommentCount sets RootCommentCount field to given value.


### GetCommentCount

`func (o *APIPage) GetCommentCount() int64`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *APIPage) GetCommentCountOk() (*int64, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *APIPage) SetCommentCount(v int64)`

SetCommentCount sets CommentCount field to given value.


### GetCreatedAt

`func (o *APIPage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIPage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIPage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTitle

`func (o *APIPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *APIPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *APIPage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *APIPage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APIPage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APIPage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *APIPage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlId

`func (o *APIPage) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *APIPage) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *APIPage) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetId

`func (o *APIPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIPage) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


