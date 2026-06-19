# PageUsersOfflineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAfterUserId** | **NullableString** |  | 
**NextAfterName** | **NullableString** |  | 
**Users** | [**[]PageUserEntry**](PageUserEntry.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewPageUsersOfflineResponse

`func NewPageUsersOfflineResponse(nextAfterUserId NullableString, nextAfterName NullableString, users []PageUserEntry, status APIStatus, ) *PageUsersOfflineResponse`

NewPageUsersOfflineResponse instantiates a new PageUsersOfflineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageUsersOfflineResponseWithDefaults

`func NewPageUsersOfflineResponseWithDefaults() *PageUsersOfflineResponse`

NewPageUsersOfflineResponseWithDefaults instantiates a new PageUsersOfflineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAfterUserId

`func (o *PageUsersOfflineResponse) GetNextAfterUserId() string`

GetNextAfterUserId returns the NextAfterUserId field if non-nil, zero value otherwise.

### GetNextAfterUserIdOk

`func (o *PageUsersOfflineResponse) GetNextAfterUserIdOk() (*string, bool)`

GetNextAfterUserIdOk returns a tuple with the NextAfterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterUserId

`func (o *PageUsersOfflineResponse) SetNextAfterUserId(v string)`

SetNextAfterUserId sets NextAfterUserId field to given value.


### SetNextAfterUserIdNil

`func (o *PageUsersOfflineResponse) SetNextAfterUserIdNil(b bool)`

 SetNextAfterUserIdNil sets the value for NextAfterUserId to be an explicit nil

### UnsetNextAfterUserId
`func (o *PageUsersOfflineResponse) UnsetNextAfterUserId()`

UnsetNextAfterUserId ensures that no value is present for NextAfterUserId, not even an explicit nil
### GetNextAfterName

`func (o *PageUsersOfflineResponse) GetNextAfterName() string`

GetNextAfterName returns the NextAfterName field if non-nil, zero value otherwise.

### GetNextAfterNameOk

`func (o *PageUsersOfflineResponse) GetNextAfterNameOk() (*string, bool)`

GetNextAfterNameOk returns a tuple with the NextAfterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterName

`func (o *PageUsersOfflineResponse) SetNextAfterName(v string)`

SetNextAfterName sets NextAfterName field to given value.


### SetNextAfterNameNil

`func (o *PageUsersOfflineResponse) SetNextAfterNameNil(b bool)`

 SetNextAfterNameNil sets the value for NextAfterName to be an explicit nil

### UnsetNextAfterName
`func (o *PageUsersOfflineResponse) UnsetNextAfterName()`

UnsetNextAfterName ensures that no value is present for NextAfterName, not even an explicit nil
### GetUsers

`func (o *PageUsersOfflineResponse) GetUsers() []PageUserEntry`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PageUsersOfflineResponse) GetUsersOk() (*[]PageUserEntry, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PageUsersOfflineResponse) SetUsers(v []PageUserEntry)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *PageUsersOfflineResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageUsersOfflineResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageUsersOfflineResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


