# SaveCommentResponseOptimized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 
**Comment** | [**PublicComment**](PublicComment.md) |  | 
**User** | [**NullableUserSessionInfo**](UserSessionInfo.md) |  | 
**ModuleData** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewSaveCommentResponseOptimized

`func NewSaveCommentResponseOptimized(status ImportedAPIStatusSUCCESS, comment PublicComment, user NullableUserSessionInfo, ) *SaveCommentResponseOptimized`

NewSaveCommentResponseOptimized instantiates a new SaveCommentResponseOptimized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCommentResponseOptimizedWithDefaults

`func NewSaveCommentResponseOptimizedWithDefaults() *SaveCommentResponseOptimized`

NewSaveCommentResponseOptimizedWithDefaults instantiates a new SaveCommentResponseOptimized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SaveCommentResponseOptimized) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaveCommentResponseOptimized) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaveCommentResponseOptimized) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.


### GetComment

`func (o *SaveCommentResponseOptimized) GetComment() PublicComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SaveCommentResponseOptimized) GetCommentOk() (*PublicComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SaveCommentResponseOptimized) SetComment(v PublicComment)`

SetComment sets Comment field to given value.


### GetUser

`func (o *SaveCommentResponseOptimized) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SaveCommentResponseOptimized) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SaveCommentResponseOptimized) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SaveCommentResponseOptimized) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SaveCommentResponseOptimized) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetModuleData

`func (o *SaveCommentResponseOptimized) GetModuleData() map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *SaveCommentResponseOptimized) GetModuleDataOk() (*map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *SaveCommentResponseOptimized) SetModuleData(v map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *SaveCommentResponseOptimized) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


