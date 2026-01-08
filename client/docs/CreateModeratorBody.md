# CreateModeratorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateModeratorBody

`func NewCreateModeratorBody(name string, email string, ) *CreateModeratorBody`

NewCreateModeratorBody instantiates a new CreateModeratorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateModeratorBodyWithDefaults

`func NewCreateModeratorBodyWithDefaults() *CreateModeratorBody`

NewCreateModeratorBodyWithDefaults instantiates a new CreateModeratorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateModeratorBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateModeratorBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateModeratorBody) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateModeratorBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateModeratorBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateModeratorBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUserId

`func (o *CreateModeratorBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateModeratorBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateModeratorBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateModeratorBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *CreateModeratorBody) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *CreateModeratorBody) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *CreateModeratorBody) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *CreateModeratorBody) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


