# APITicketDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UrlId** | **string** |  | 
**UserId** | **string** |  | 
**ManagedByTenantId** | **string** |  | 
**AssignedUserIds** | **[]string** |  | 
**Subject** | **string** |  | 
**CreatedAt** | **string** |  | 
**State** | **int32** |  | 
**FileCount** | **int32** |  | 
**Files** | [**[]APITicketFile**](APITicketFile.md) |  | 
**ReopenedAt** | Pointer to **NullableString** |  | [optional] 
**ResolvedAt** | Pointer to **NullableString** |  | [optional] 
**AckAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAPITicketDetail

`func NewAPITicketDetail(id string, urlId string, userId string, managedByTenantId string, assignedUserIds []string, subject string, createdAt string, state int32, fileCount int32, files []APITicketFile, ) *APITicketDetail`

NewAPITicketDetail instantiates a new APITicketDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITicketDetailWithDefaults

`func NewAPITicketDetailWithDefaults() *APITicketDetail`

NewAPITicketDetailWithDefaults instantiates a new APITicketDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITicketDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITicketDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITicketDetail) SetId(v string)`

SetId sets Id field to given value.


### GetUrlId

`func (o *APITicketDetail) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *APITicketDetail) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *APITicketDetail) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUserId

`func (o *APITicketDetail) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APITicketDetail) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APITicketDetail) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetManagedByTenantId

`func (o *APITicketDetail) GetManagedByTenantId() string`

GetManagedByTenantId returns the ManagedByTenantId field if non-nil, zero value otherwise.

### GetManagedByTenantIdOk

`func (o *APITicketDetail) GetManagedByTenantIdOk() (*string, bool)`

GetManagedByTenantIdOk returns a tuple with the ManagedByTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedByTenantId

`func (o *APITicketDetail) SetManagedByTenantId(v string)`

SetManagedByTenantId sets ManagedByTenantId field to given value.


### GetAssignedUserIds

`func (o *APITicketDetail) GetAssignedUserIds() []string`

GetAssignedUserIds returns the AssignedUserIds field if non-nil, zero value otherwise.

### GetAssignedUserIdsOk

`func (o *APITicketDetail) GetAssignedUserIdsOk() (*[]string, bool)`

GetAssignedUserIdsOk returns a tuple with the AssignedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserIds

`func (o *APITicketDetail) SetAssignedUserIds(v []string)`

SetAssignedUserIds sets AssignedUserIds field to given value.


### GetSubject

`func (o *APITicketDetail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *APITicketDetail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *APITicketDetail) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCreatedAt

`func (o *APITicketDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APITicketDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APITicketDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetState

`func (o *APITicketDetail) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *APITicketDetail) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *APITicketDetail) SetState(v int32)`

SetState sets State field to given value.


### GetFileCount

`func (o *APITicketDetail) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *APITicketDetail) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *APITicketDetail) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.


### GetFiles

`func (o *APITicketDetail) GetFiles() []APITicketFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *APITicketDetail) GetFilesOk() (*[]APITicketFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *APITicketDetail) SetFiles(v []APITicketFile)`

SetFiles sets Files field to given value.


### GetReopenedAt

`func (o *APITicketDetail) GetReopenedAt() string`

GetReopenedAt returns the ReopenedAt field if non-nil, zero value otherwise.

### GetReopenedAtOk

`func (o *APITicketDetail) GetReopenedAtOk() (*string, bool)`

GetReopenedAtOk returns a tuple with the ReopenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenedAt

`func (o *APITicketDetail) SetReopenedAt(v string)`

SetReopenedAt sets ReopenedAt field to given value.

### HasReopenedAt

`func (o *APITicketDetail) HasReopenedAt() bool`

HasReopenedAt returns a boolean if a field has been set.

### SetReopenedAtNil

`func (o *APITicketDetail) SetReopenedAtNil(b bool)`

 SetReopenedAtNil sets the value for ReopenedAt to be an explicit nil

### UnsetReopenedAt
`func (o *APITicketDetail) UnsetReopenedAt()`

UnsetReopenedAt ensures that no value is present for ReopenedAt, not even an explicit nil
### GetResolvedAt

`func (o *APITicketDetail) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *APITicketDetail) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *APITicketDetail) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *APITicketDetail) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *APITicketDetail) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *APITicketDetail) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetAckAt

`func (o *APITicketDetail) GetAckAt() string`

GetAckAt returns the AckAt field if non-nil, zero value otherwise.

### GetAckAtOk

`func (o *APITicketDetail) GetAckAtOk() (*string, bool)`

GetAckAtOk returns a tuple with the AckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAt

`func (o *APITicketDetail) SetAckAt(v string)`

SetAckAt sets AckAt field to given value.

### HasAckAt

`func (o *APITicketDetail) HasAckAt() bool`

HasAckAt returns a boolean if a field has been set.

### SetAckAtNil

`func (o *APITicketDetail) SetAckAtNil(b bool)`

 SetAckAtNil sets the value for AckAt to be an explicit nil

### UnsetAckAt
`func (o *APITicketDetail) UnsetAckAt()`

UnsetAckAt ensures that no value is present for AckAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


