# UserSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewUserSearchResult

`func NewUserSearchResult(id string, name string, type_ string, ) *UserSearchResult`

NewUserSearchResult instantiates a new UserSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSearchResultWithDefaults

`func NewUserSearchResultWithDefaults() *UserSearchResult`

NewUserSearchResultWithDefaults instantiates a new UserSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSearchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSearchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSearchResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSearchResult) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *UserSearchResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserSearchResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserSearchResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserSearchResult) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *UserSearchResult) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *UserSearchResult) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *UserSearchResult) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *UserSearchResult) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *UserSearchResult) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *UserSearchResult) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetType

`func (o *UserSearchResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSearchResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSearchResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


