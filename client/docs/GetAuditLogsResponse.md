# GetAuditLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**AuditLogs** | [**[]PickTenantAuditLogTenantAuditLogKeys**](PickTenantAuditLogTenantAuditLogKeys.md) |  | 

## Methods

### NewGetAuditLogsResponse

`func NewGetAuditLogsResponse(status APIStatus, auditLogs []PickTenantAuditLogTenantAuditLogKeys, ) *GetAuditLogsResponse`

NewGetAuditLogsResponse instantiates a new GetAuditLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditLogsResponseWithDefaults

`func NewGetAuditLogsResponseWithDefaults() *GetAuditLogsResponse`

NewGetAuditLogsResponseWithDefaults instantiates a new GetAuditLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAuditLogsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAuditLogsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAuditLogsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetAuditLogs

`func (o *GetAuditLogsResponse) GetAuditLogs() []PickTenantAuditLogTenantAuditLogKeys`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *GetAuditLogsResponse) GetAuditLogsOk() (*[]PickTenantAuditLogTenantAuditLogKeys, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *GetAuditLogsResponse) SetAuditLogs(v []PickTenantAuditLogTenantAuditLogKeys)`

SetAuditLogs sets AuditLogs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


