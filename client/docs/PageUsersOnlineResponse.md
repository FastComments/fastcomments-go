# PageUsersOnlineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAfterUserId** | **NullableString** |  | 
**NextAfterName** | **NullableString** |  | 
**TotalCount** | **float64** |  | 
**AnonCount** | **float64** |  | 
**Users** | [**[]PageUserEntry**](PageUserEntry.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewPageUsersOnlineResponse

`func NewPageUsersOnlineResponse(nextAfterUserId NullableString, nextAfterName NullableString, totalCount float64, anonCount float64, users []PageUserEntry, status APIStatus, ) *PageUsersOnlineResponse`

NewPageUsersOnlineResponse instantiates a new PageUsersOnlineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageUsersOnlineResponseWithDefaults

`func NewPageUsersOnlineResponseWithDefaults() *PageUsersOnlineResponse`

NewPageUsersOnlineResponseWithDefaults instantiates a new PageUsersOnlineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAfterUserId

`func (o *PageUsersOnlineResponse) GetNextAfterUserId() string`

GetNextAfterUserId returns the NextAfterUserId field if non-nil, zero value otherwise.

### GetNextAfterUserIdOk

`func (o *PageUsersOnlineResponse) GetNextAfterUserIdOk() (*string, bool)`

GetNextAfterUserIdOk returns a tuple with the NextAfterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterUserId

`func (o *PageUsersOnlineResponse) SetNextAfterUserId(v string)`

SetNextAfterUserId sets NextAfterUserId field to given value.


### SetNextAfterUserIdNil

`func (o *PageUsersOnlineResponse) SetNextAfterUserIdNil(b bool)`

 SetNextAfterUserIdNil sets the value for NextAfterUserId to be an explicit nil

### UnsetNextAfterUserId
`func (o *PageUsersOnlineResponse) UnsetNextAfterUserId()`

UnsetNextAfterUserId ensures that no value is present for NextAfterUserId, not even an explicit nil
### GetNextAfterName

`func (o *PageUsersOnlineResponse) GetNextAfterName() string`

GetNextAfterName returns the NextAfterName field if non-nil, zero value otherwise.

### GetNextAfterNameOk

`func (o *PageUsersOnlineResponse) GetNextAfterNameOk() (*string, bool)`

GetNextAfterNameOk returns a tuple with the NextAfterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAfterName

`func (o *PageUsersOnlineResponse) SetNextAfterName(v string)`

SetNextAfterName sets NextAfterName field to given value.


### SetNextAfterNameNil

`func (o *PageUsersOnlineResponse) SetNextAfterNameNil(b bool)`

 SetNextAfterNameNil sets the value for NextAfterName to be an explicit nil

### UnsetNextAfterName
`func (o *PageUsersOnlineResponse) UnsetNextAfterName()`

UnsetNextAfterName ensures that no value is present for NextAfterName, not even an explicit nil
### GetTotalCount

`func (o *PageUsersOnlineResponse) GetTotalCount() float64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PageUsersOnlineResponse) GetTotalCountOk() (*float64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PageUsersOnlineResponse) SetTotalCount(v float64)`

SetTotalCount sets TotalCount field to given value.


### GetAnonCount

`func (o *PageUsersOnlineResponse) GetAnonCount() float64`

GetAnonCount returns the AnonCount field if non-nil, zero value otherwise.

### GetAnonCountOk

`func (o *PageUsersOnlineResponse) GetAnonCountOk() (*float64, bool)`

GetAnonCountOk returns a tuple with the AnonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonCount

`func (o *PageUsersOnlineResponse) SetAnonCount(v float64)`

SetAnonCount sets AnonCount field to given value.


### GetUsers

`func (o *PageUsersOnlineResponse) GetUsers() []PageUserEntry`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PageUsersOnlineResponse) GetUsersOk() (*[]PageUserEntry, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PageUsersOnlineResponse) SetUsers(v []PageUserEntry)`

SetUsers sets Users field to given value.


### GetStatus

`func (o *PageUsersOnlineResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageUsersOnlineResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageUsersOnlineResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


