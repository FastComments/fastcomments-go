# APIBanUserChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBannedUserId** | Pointer to **string** |  | [optional] 
**UpdatedBannedUserId** | Pointer to **string** |  | [optional] 
**DeletedBannedUsers** | Pointer to [**[]APIBannedUser**](APIBannedUser.md) |  | [optional] 
**ChangedValuesBefore** | Pointer to [**APIBanUserChangedValues**](APIBanUserChangedValues.md) |  | [optional] 

## Methods

### NewAPIBanUserChangeLog

`func NewAPIBanUserChangeLog() *APIBanUserChangeLog`

NewAPIBanUserChangeLog instantiates a new APIBanUserChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIBanUserChangeLogWithDefaults

`func NewAPIBanUserChangeLogWithDefaults() *APIBanUserChangeLog`

NewAPIBanUserChangeLogWithDefaults instantiates a new APIBanUserChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBannedUserId

`func (o *APIBanUserChangeLog) GetCreatedBannedUserId() string`

GetCreatedBannedUserId returns the CreatedBannedUserId field if non-nil, zero value otherwise.

### GetCreatedBannedUserIdOk

`func (o *APIBanUserChangeLog) GetCreatedBannedUserIdOk() (*string, bool)`

GetCreatedBannedUserIdOk returns a tuple with the CreatedBannedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBannedUserId

`func (o *APIBanUserChangeLog) SetCreatedBannedUserId(v string)`

SetCreatedBannedUserId sets CreatedBannedUserId field to given value.

### HasCreatedBannedUserId

`func (o *APIBanUserChangeLog) HasCreatedBannedUserId() bool`

HasCreatedBannedUserId returns a boolean if a field has been set.

### GetUpdatedBannedUserId

`func (o *APIBanUserChangeLog) GetUpdatedBannedUserId() string`

GetUpdatedBannedUserId returns the UpdatedBannedUserId field if non-nil, zero value otherwise.

### GetUpdatedBannedUserIdOk

`func (o *APIBanUserChangeLog) GetUpdatedBannedUserIdOk() (*string, bool)`

GetUpdatedBannedUserIdOk returns a tuple with the UpdatedBannedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBannedUserId

`func (o *APIBanUserChangeLog) SetUpdatedBannedUserId(v string)`

SetUpdatedBannedUserId sets UpdatedBannedUserId field to given value.

### HasUpdatedBannedUserId

`func (o *APIBanUserChangeLog) HasUpdatedBannedUserId() bool`

HasUpdatedBannedUserId returns a boolean if a field has been set.

### GetDeletedBannedUsers

`func (o *APIBanUserChangeLog) GetDeletedBannedUsers() []APIBannedUser`

GetDeletedBannedUsers returns the DeletedBannedUsers field if non-nil, zero value otherwise.

### GetDeletedBannedUsersOk

`func (o *APIBanUserChangeLog) GetDeletedBannedUsersOk() (*[]APIBannedUser, bool)`

GetDeletedBannedUsersOk returns a tuple with the DeletedBannedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBannedUsers

`func (o *APIBanUserChangeLog) SetDeletedBannedUsers(v []APIBannedUser)`

SetDeletedBannedUsers sets DeletedBannedUsers field to given value.

### HasDeletedBannedUsers

`func (o *APIBanUserChangeLog) HasDeletedBannedUsers() bool`

HasDeletedBannedUsers returns a boolean if a field has been set.

### GetChangedValuesBefore

`func (o *APIBanUserChangeLog) GetChangedValuesBefore() APIBanUserChangedValues`

GetChangedValuesBefore returns the ChangedValuesBefore field if non-nil, zero value otherwise.

### GetChangedValuesBeforeOk

`func (o *APIBanUserChangeLog) GetChangedValuesBeforeOk() (*APIBanUserChangedValues, bool)`

GetChangedValuesBeforeOk returns a tuple with the ChangedValuesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedValuesBefore

`func (o *APIBanUserChangeLog) SetChangedValuesBefore(v APIBanUserChangedValues)`

SetChangedValuesBefore sets ChangedValuesBefore field to given value.

### HasChangedValuesBefore

`func (o *APIBanUserChangeLog) HasChangedValuesBefore() bool`

HasChangedValuesBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


