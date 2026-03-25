# SearchUsersSectionedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Sections** | [**[]UserSearchSectionResult**](UserSearchSectionResult.md) |  | 

## Methods

### NewSearchUsersSectionedResponse

`func NewSearchUsersSectionedResponse(status APIStatus, sections []UserSearchSectionResult, ) *SearchUsersSectionedResponse`

NewSearchUsersSectionedResponse instantiates a new SearchUsersSectionedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUsersSectionedResponseWithDefaults

`func NewSearchUsersSectionedResponseWithDefaults() *SearchUsersSectionedResponse`

NewSearchUsersSectionedResponseWithDefaults instantiates a new SearchUsersSectionedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SearchUsersSectionedResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchUsersSectionedResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchUsersSectionedResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetSections

`func (o *SearchUsersSectionedResponse) GetSections() []UserSearchSectionResult`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SearchUsersSectionedResponse) GetSectionsOk() (*[]UserSearchSectionResult, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SearchUsersSectionedResponse) SetSections(v []UserSearchSectionResult)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


