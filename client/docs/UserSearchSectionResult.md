# UserSearchSectionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Section** | [**UserSearchSection**](UserSearchSection.md) |  | 
**Users** | [**[]UserSearchResult**](UserSearchResult.md) |  | 

## Methods

### NewUserSearchSectionResult

`func NewUserSearchSectionResult(section UserSearchSection, users []UserSearchResult, ) *UserSearchSectionResult`

NewUserSearchSectionResult instantiates a new UserSearchSectionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSearchSectionResultWithDefaults

`func NewUserSearchSectionResultWithDefaults() *UserSearchSectionResult`

NewUserSearchSectionResultWithDefaults instantiates a new UserSearchSectionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSection

`func (o *UserSearchSectionResult) GetSection() UserSearchSection`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *UserSearchSectionResult) GetSectionOk() (*UserSearchSection, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *UserSearchSectionResult) SetSection(v UserSearchSection)`

SetSection sets Section field to given value.


### GetUsers

`func (o *UserSearchSectionResult) GetUsers() []UserSearchResult`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserSearchSectionResult) GetUsersOk() (*[]UserSearchResult, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserSearchSectionResult) SetUsers(v []UserSearchResult)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


