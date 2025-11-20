# SaveCommentsResponseWithPresence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Comment** | [**PublicComment**](PublicComment.md) |  | 
**User** | [**NullableUserSessionInfo**](UserSessionInfo.md) |  | 
**ModuleData** | Pointer to **map[string]map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**UserIdWS** | Pointer to **string** |  | [optional] 

## Methods

### NewSaveCommentsResponseWithPresence

`func NewSaveCommentsResponseWithPresence(status APIStatus, comment PublicComment, user NullableUserSessionInfo, ) *SaveCommentsResponseWithPresence`

NewSaveCommentsResponseWithPresence instantiates a new SaveCommentsResponseWithPresence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCommentsResponseWithPresenceWithDefaults

`func NewSaveCommentsResponseWithPresenceWithDefaults() *SaveCommentsResponseWithPresence`

NewSaveCommentsResponseWithPresenceWithDefaults instantiates a new SaveCommentsResponseWithPresence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SaveCommentsResponseWithPresence) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaveCommentsResponseWithPresence) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaveCommentsResponseWithPresence) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *SaveCommentsResponseWithPresence) GetComment() PublicComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SaveCommentsResponseWithPresence) GetCommentOk() (*PublicComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SaveCommentsResponseWithPresence) SetComment(v PublicComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *SaveCommentsResponseWithPresence) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SaveCommentsResponseWithPresence) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SaveCommentsResponseWithPresence) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SaveCommentsResponseWithPresence) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SaveCommentsResponseWithPresence) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *SaveCommentsResponseWithPresence) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *SaveCommentsResponseWithPresence) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *SaveCommentsResponseWithPresence) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *SaveCommentsResponseWithPresence) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetUserIdWS

`func (o *SaveCommentsResponseWithPresence) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *SaveCommentsResponseWithPresence) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *SaveCommentsResponseWithPresence) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *SaveCommentsResponseWithPresence) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


