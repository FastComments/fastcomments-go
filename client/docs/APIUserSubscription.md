# APIUserSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**PageTitle** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UrlId** | **string** |  | 
**AnonUserId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewAPIUserSubscription

`func NewAPIUserSubscription(createdAt time.Time, urlId string, id string, ) *APIUserSubscription`

NewAPIUserSubscription instantiates a new APIUserSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIUserSubscriptionWithDefaults

`func NewAPIUserSubscriptionWithDefaults() *APIUserSubscription`

NewAPIUserSubscriptionWithDefaults instantiates a new APIUserSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *APIUserSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIUserSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIUserSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPageTitle

`func (o *APIUserSubscription) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *APIUserSubscription) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *APIUserSubscription) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *APIUserSubscription) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetUrl

`func (o *APIUserSubscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APIUserSubscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APIUserSubscription) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *APIUserSubscription) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlId

`func (o *APIUserSubscription) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *APIUserSubscription) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *APIUserSubscription) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetAnonUserId

`func (o *APIUserSubscription) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *APIUserSubscription) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *APIUserSubscription) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *APIUserSubscription) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetUserId

`func (o *APIUserSubscription) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APIUserSubscription) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APIUserSubscription) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *APIUserSubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetId

`func (o *APIUserSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIUserSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIUserSubscription) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


