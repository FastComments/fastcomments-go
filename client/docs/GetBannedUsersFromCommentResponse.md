# GetBannedUsersFromCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannedUsers** | [**[]APIBannedUserWithMultiMatchInfo**](APIBannedUserWithMultiMatchInfo.md) |  | 
**Code** | Pointer to **string** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetBannedUsersFromCommentResponse

`func NewGetBannedUsersFromCommentResponse(bannedUsers []APIBannedUserWithMultiMatchInfo, status APIStatus, ) *GetBannedUsersFromCommentResponse`

NewGetBannedUsersFromCommentResponse instantiates a new GetBannedUsersFromCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBannedUsersFromCommentResponseWithDefaults

`func NewGetBannedUsersFromCommentResponseWithDefaults() *GetBannedUsersFromCommentResponse`

NewGetBannedUsersFromCommentResponseWithDefaults instantiates a new GetBannedUsersFromCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannedUsers

`func (o *GetBannedUsersFromCommentResponse) GetBannedUsers() []APIBannedUserWithMultiMatchInfo`

GetBannedUsers returns the BannedUsers field if non-nil, zero value otherwise.

### GetBannedUsersOk

`func (o *GetBannedUsersFromCommentResponse) GetBannedUsersOk() (*[]APIBannedUserWithMultiMatchInfo, bool)`

GetBannedUsersOk returns a tuple with the BannedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUsers

`func (o *GetBannedUsersFromCommentResponse) SetBannedUsers(v []APIBannedUserWithMultiMatchInfo)`

SetBannedUsers sets BannedUsers field to given value.


### GetCode

`func (o *GetBannedUsersFromCommentResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetBannedUsersFromCommentResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetBannedUsersFromCommentResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetBannedUsersFromCommentResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetBannedUsersFromCommentResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBannedUsersFromCommentResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBannedUsersFromCommentResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


