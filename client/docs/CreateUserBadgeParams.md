# CreateUserBadgeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**BadgeId** | **string** |  | 
**DisplayedOnComments** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateUserBadgeParams

`func NewCreateUserBadgeParams(userId string, badgeId string, ) *CreateUserBadgeParams`

NewCreateUserBadgeParams instantiates a new CreateUserBadgeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserBadgeParamsWithDefaults

`func NewCreateUserBadgeParamsWithDefaults() *CreateUserBadgeParams`

NewCreateUserBadgeParamsWithDefaults instantiates a new CreateUserBadgeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateUserBadgeParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateUserBadgeParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateUserBadgeParams) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBadgeId

`func (o *CreateUserBadgeParams) GetBadgeId() string`

GetBadgeId returns the BadgeId field if non-nil, zero value otherwise.

### GetBadgeIdOk

`func (o *CreateUserBadgeParams) GetBadgeIdOk() (*string, bool)`

GetBadgeIdOk returns a tuple with the BadgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeId

`func (o *CreateUserBadgeParams) SetBadgeId(v string)`

SetBadgeId sets BadgeId field to given value.


### GetDisplayedOnComments

`func (o *CreateUserBadgeParams) GetDisplayedOnComments() bool`

GetDisplayedOnComments returns the DisplayedOnComments field if non-nil, zero value otherwise.

### GetDisplayedOnCommentsOk

`func (o *CreateUserBadgeParams) GetDisplayedOnCommentsOk() (*bool, bool)`

GetDisplayedOnCommentsOk returns a tuple with the DisplayedOnComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedOnComments

`func (o *CreateUserBadgeParams) SetDisplayedOnComments(v bool)`

SetDisplayedOnComments sets DisplayedOnComments field to given value.

### HasDisplayedOnComments

`func (o *CreateUserBadgeParams) HasDisplayedOnComments() bool`

HasDisplayedOnComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


