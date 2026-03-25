# APITicket

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

## Methods

### NewAPITicket

`func NewAPITicket(id string, urlId string, userId string, managedByTenantId string, assignedUserIds []string, subject string, createdAt string, state int32, fileCount int32, ) *APITicket`

NewAPITicket instantiates a new APITicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITicketWithDefaults

`func NewAPITicketWithDefaults() *APITicket`

NewAPITicketWithDefaults instantiates a new APITicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITicket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITicket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITicket) SetId(v string)`

SetId sets Id field to given value.


### GetUrlId

`func (o *APITicket) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *APITicket) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *APITicket) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUserId

`func (o *APITicket) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *APITicket) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *APITicket) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetManagedByTenantId

`func (o *APITicket) GetManagedByTenantId() string`

GetManagedByTenantId returns the ManagedByTenantId field if non-nil, zero value otherwise.

### GetManagedByTenantIdOk

`func (o *APITicket) GetManagedByTenantIdOk() (*string, bool)`

GetManagedByTenantIdOk returns a tuple with the ManagedByTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedByTenantId

`func (o *APITicket) SetManagedByTenantId(v string)`

SetManagedByTenantId sets ManagedByTenantId field to given value.


### GetAssignedUserIds

`func (o *APITicket) GetAssignedUserIds() []string`

GetAssignedUserIds returns the AssignedUserIds field if non-nil, zero value otherwise.

### GetAssignedUserIdsOk

`func (o *APITicket) GetAssignedUserIdsOk() (*[]string, bool)`

GetAssignedUserIdsOk returns a tuple with the AssignedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserIds

`func (o *APITicket) SetAssignedUserIds(v []string)`

SetAssignedUserIds sets AssignedUserIds field to given value.


### GetSubject

`func (o *APITicket) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *APITicket) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *APITicket) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCreatedAt

`func (o *APITicket) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APITicket) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APITicket) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetState

`func (o *APITicket) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *APITicket) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *APITicket) SetState(v int32)`

SetState sets State field to given value.


### GetFileCount

`func (o *APITicket) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *APITicket) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *APITicket) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


