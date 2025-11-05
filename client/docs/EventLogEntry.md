# EventLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**BroadcastId** | **string** |  | 
**Data** | **string** |  | 

## Methods

### NewEventLogEntry

`func NewEventLogEntry(id string, createdAt time.Time, tenantId string, urlId string, broadcastId string, data string, ) *EventLogEntry`

NewEventLogEntry instantiates a new EventLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogEntryWithDefaults

`func NewEventLogEntryWithDefaults() *EventLogEntry`

NewEventLogEntryWithDefaults instantiates a new EventLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EventLogEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventLogEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventLogEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTenantId

`func (o *EventLogEntry) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EventLogEntry) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EventLogEntry) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *EventLogEntry) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *EventLogEntry) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *EventLogEntry) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetBroadcastId

`func (o *EventLogEntry) GetBroadcastId() string`

GetBroadcastId returns the BroadcastId field if non-nil, zero value otherwise.

### GetBroadcastIdOk

`func (o *EventLogEntry) GetBroadcastIdOk() (*string, bool)`

GetBroadcastIdOk returns a tuple with the BroadcastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastId

`func (o *EventLogEntry) SetBroadcastId(v string)`

SetBroadcastId sets BroadcastId field to given value.


### GetData

`func (o *EventLogEntry) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventLogEntry) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventLogEntry) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


