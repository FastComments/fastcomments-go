# PendingCommentToSyncOutbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CommentId** | **string** |  | 
**Comment** | Pointer to [**FComment**](FComment.md) |  | [optional] 
**ExternalId** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**TenantId** | **string** |  | 
**AttemptCount** | **float64** |  | 
**NextAttemptAt** | **time.Time** |  | 
**EventType** | **float64** |  | 
**Type** | **float64** |  | 
**Domain** | **string** |  | 
**LastError** | **map[string]interface{}** |  | 
**WebhookId** | Pointer to **string** |  | [optional] 

## Methods

### NewPendingCommentToSyncOutbound

`func NewPendingCommentToSyncOutbound(id string, commentId string, externalId NullableString, createdAt time.Time, tenantId string, attemptCount float64, nextAttemptAt time.Time, eventType float64, type_ float64, domain string, lastError map[string]interface{}, ) *PendingCommentToSyncOutbound`

NewPendingCommentToSyncOutbound instantiates a new PendingCommentToSyncOutbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingCommentToSyncOutboundWithDefaults

`func NewPendingCommentToSyncOutboundWithDefaults() *PendingCommentToSyncOutbound`

NewPendingCommentToSyncOutboundWithDefaults instantiates a new PendingCommentToSyncOutbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PendingCommentToSyncOutbound) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PendingCommentToSyncOutbound) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PendingCommentToSyncOutbound) SetId(v string)`

SetId sets Id field to given value.


### GetCommentId

`func (o *PendingCommentToSyncOutbound) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *PendingCommentToSyncOutbound) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *PendingCommentToSyncOutbound) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.


### GetComment

`func (o *PendingCommentToSyncOutbound) GetComment() FComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PendingCommentToSyncOutbound) GetCommentOk() (*FComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PendingCommentToSyncOutbound) SetComment(v FComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PendingCommentToSyncOutbound) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalId

`func (o *PendingCommentToSyncOutbound) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PendingCommentToSyncOutbound) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PendingCommentToSyncOutbound) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *PendingCommentToSyncOutbound) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *PendingCommentToSyncOutbound) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCreatedAt

`func (o *PendingCommentToSyncOutbound) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PendingCommentToSyncOutbound) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PendingCommentToSyncOutbound) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTenantId

`func (o *PendingCommentToSyncOutbound) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PendingCommentToSyncOutbound) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PendingCommentToSyncOutbound) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAttemptCount

`func (o *PendingCommentToSyncOutbound) GetAttemptCount() float64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *PendingCommentToSyncOutbound) GetAttemptCountOk() (*float64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *PendingCommentToSyncOutbound) SetAttemptCount(v float64)`

SetAttemptCount sets AttemptCount field to given value.


### GetNextAttemptAt

`func (o *PendingCommentToSyncOutbound) GetNextAttemptAt() time.Time`

GetNextAttemptAt returns the NextAttemptAt field if non-nil, zero value otherwise.

### GetNextAttemptAtOk

`func (o *PendingCommentToSyncOutbound) GetNextAttemptAtOk() (*time.Time, bool)`

GetNextAttemptAtOk returns a tuple with the NextAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttemptAt

`func (o *PendingCommentToSyncOutbound) SetNextAttemptAt(v time.Time)`

SetNextAttemptAt sets NextAttemptAt field to given value.


### GetEventType

`func (o *PendingCommentToSyncOutbound) GetEventType() float64`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PendingCommentToSyncOutbound) GetEventTypeOk() (*float64, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PendingCommentToSyncOutbound) SetEventType(v float64)`

SetEventType sets EventType field to given value.


### GetType

`func (o *PendingCommentToSyncOutbound) GetType() float64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingCommentToSyncOutbound) GetTypeOk() (*float64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingCommentToSyncOutbound) SetType(v float64)`

SetType sets Type field to given value.


### GetDomain

`func (o *PendingCommentToSyncOutbound) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PendingCommentToSyncOutbound) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PendingCommentToSyncOutbound) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetLastError

`func (o *PendingCommentToSyncOutbound) GetLastError() map[string]interface{}`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *PendingCommentToSyncOutbound) GetLastErrorOk() (*map[string]interface{}, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *PendingCommentToSyncOutbound) SetLastError(v map[string]interface{})`

SetLastError sets LastError field to given value.


### GetWebhookId

`func (o *PendingCommentToSyncOutbound) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *PendingCommentToSyncOutbound) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *PendingCommentToSyncOutbound) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *PendingCommentToSyncOutbound) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


