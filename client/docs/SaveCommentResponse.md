# SaveCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Comment** | [**FComment**](FComment.md) |  | 
**User** | [**NullableUserSessionInfo**](UserSessionInfo.md) |  | 
**ModuleData** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewSaveCommentResponse

`func NewSaveCommentResponse(status APIStatus, comment FComment, user NullableUserSessionInfo, ) *SaveCommentResponse`

NewSaveCommentResponse instantiates a new SaveCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCommentResponseWithDefaults

`func NewSaveCommentResponseWithDefaults() *SaveCommentResponse`

NewSaveCommentResponseWithDefaults instantiates a new SaveCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SaveCommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaveCommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaveCommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *SaveCommentResponse) GetComment() FComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SaveCommentResponse) GetCommentOk() (*FComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SaveCommentResponse) SetComment(v FComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *SaveCommentResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SaveCommentResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SaveCommentResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SaveCommentResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SaveCommentResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *SaveCommentResponse) GetModuleData() map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *SaveCommentResponse) GetModuleDataOk() (*map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *SaveCommentResponse) SetModuleData(v map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *SaveCommentResponse) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


