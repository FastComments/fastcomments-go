# SearchUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 
**Users** | [**[]UserSearchResult**](UserSearchResult.md) |  | 

## Methods

### NewSearchUsersResponse

`func NewSearchUsersResponse(status ImportedAPIStatusSUCCESS, users []UserSearchResult, ) *SearchUsersResponse`

NewSearchUsersResponse instantiates a new SearchUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUsersResponseWithDefaults

`func NewSearchUsersResponseWithDefaults() *SearchUsersResponse`

NewSearchUsersResponseWithDefaults instantiates a new SearchUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SearchUsersResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchUsersResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchUsersResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.


### GetUsers

`func (o *SearchUsersResponse) GetUsers() []UserSearchResult`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SearchUsersResponse) GetUsersOk() (*[]UserSearchResult, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SearchUsersResponse) SetUsers(v []UserSearchResult)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


