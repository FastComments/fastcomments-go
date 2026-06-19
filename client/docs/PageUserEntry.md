# PageUserEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrivate** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Id** | **string** |  | 

## Methods

### NewPageUserEntry

`func NewPageUserEntry(displayName string, id string, ) *PageUserEntry`

NewPageUserEntry instantiates a new PageUserEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageUserEntryWithDefaults

`func NewPageUserEntryWithDefaults() *PageUserEntry`

NewPageUserEntryWithDefaults instantiates a new PageUserEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrivate

`func (o *PageUserEntry) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *PageUserEntry) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *PageUserEntry) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *PageUserEntry) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *PageUserEntry) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PageUserEntry) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PageUserEntry) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PageUserEntry) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetDisplayName

`func (o *PageUserEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PageUserEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PageUserEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *PageUserEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageUserEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageUserEntry) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


