# SearchUsersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Sections** | [**[]UserSearchSectionResult**](UserSearchSectionResult.md) |  | 
**Users** | [**[]UserSearchResult**](UserSearchResult.md) |  | 

## Methods

### NewSearchUsersResult

`func NewSearchUsersResult(status APIStatus, sections []UserSearchSectionResult, users []UserSearchResult, ) *SearchUsersResult`

NewSearchUsersResult instantiates a new SearchUsersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUsersResultWithDefaults

`func NewSearchUsersResultWithDefaults() *SearchUsersResult`

NewSearchUsersResultWithDefaults instantiates a new SearchUsersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SearchUsersResult) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchUsersResult) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchUsersResult) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetSections

`func (o *SearchUsersResult) GetSections() []UserSearchSectionResult`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SearchUsersResult) GetSectionsOk() (*[]UserSearchSectionResult, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SearchUsersResult) SetSections(v []UserSearchSectionResult)`

SetSections sets Sections field to given value.


### GetUsers

`func (o *SearchUsersResult) GetUsers() []UserSearchResult`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SearchUsersResult) GetUsersOk() (*[]UserSearchResult, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SearchUsersResult) SetUsers(v []UserSearchResult)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


