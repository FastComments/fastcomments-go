# UserNotificationCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Count** | **float64** |  | 
**CreatedAt** | **time.Time** |  | 
**ExpireAt** | **time.Time** |  | 

## Methods

### NewUserNotificationCount

`func NewUserNotificationCount(id string, count float64, createdAt time.Time, expireAt time.Time, ) *UserNotificationCount`

NewUserNotificationCount instantiates a new UserNotificationCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationCountWithDefaults

`func NewUserNotificationCountWithDefaults() *UserNotificationCount`

NewUserNotificationCountWithDefaults instantiates a new UserNotificationCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserNotificationCount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNotificationCount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNotificationCount) SetId(v string)`

SetId sets Id field to given value.


### GetCount

`func (o *UserNotificationCount) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserNotificationCount) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserNotificationCount) SetCount(v float64)`

SetCount sets Count field to given value.


### GetCreatedAt

`func (o *UserNotificationCount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserNotificationCount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserNotificationCount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpireAt

`func (o *UserNotificationCount) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *UserNotificationCount) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *UserNotificationCount) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


