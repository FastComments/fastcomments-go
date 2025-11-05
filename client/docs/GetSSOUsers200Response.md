# GetSSOUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]APISSOUser**](APISSOUser.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetSSOUsers200Response

`func NewGetSSOUsers200Response(users []APISSOUser, status string, ) *GetSSOUsers200Response`

NewGetSSOUsers200Response instantiates a new GetSSOUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSSOUsers200ResponseWithDefaults

`func NewGetSSOUsers200ResponseWithDefaults() *GetSSOUsers200Response`

NewGetSSOUsers200ResponseWithDefaults instantiates a new GetSSOUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetSSOUsers200Response) GetUsers() []APISSOUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetSSOUsers200Response) GetUsersOk() (*[]APISSOUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetSSOUsers200Response) SetUsers(v []APISSOUser)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *GetSSOUsers200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSSOUsers200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSSOUsers200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


