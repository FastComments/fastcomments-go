# GetSSOUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]APISSOUser**](APISSOUser.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetSSOUsersResponse

`func NewGetSSOUsersResponse(users []APISSOUser, status string, ) *GetSSOUsersResponse`

NewGetSSOUsersResponse instantiates a new GetSSOUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSSOUsersResponseWithDefaults

`func NewGetSSOUsersResponseWithDefaults() *GetSSOUsersResponse`

NewGetSSOUsersResponseWithDefaults instantiates a new GetSSOUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetSSOUsersResponse) GetUsers() []APISSOUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetSSOUsersResponse) GetUsersOk() (*[]APISSOUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetSSOUsersResponse) SetUsers(v []APISSOUser)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *GetSSOUsersResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSSOUsersResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSSOUsersResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


