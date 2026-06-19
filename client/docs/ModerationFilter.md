# ModerationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviewed** | Pointer to **bool** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**IsBannedUser** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**FlagCountGt** | Pointer to **float64** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UrlId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**CommentTextSearch** | Pointer to **[]string** | Text search terms. Each term is matched case-insensitively against the comment text. A term wrapped in quotes means exact phrase match. | [optional] 
**ExactCommentText** | Pointer to **string** | Set by the &#x60;exact&#x3D;\&quot;...\&quot;&#x60; search syntax. The comment text must equal this value exactly (case-sensitive, full-string), as opposed to the substring matching of commentTextSearch. | [optional] 

## Methods

### NewModerationFilter

`func NewModerationFilter() *ModerationFilter`

NewModerationFilter instantiates a new ModerationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationFilterWithDefaults

`func NewModerationFilterWithDefaults() *ModerationFilter`

NewModerationFilterWithDefaults instantiates a new ModerationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewed

`func (o *ModerationFilter) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *ModerationFilter) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *ModerationFilter) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *ModerationFilter) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetApproved

`func (o *ModerationFilter) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ModerationFilter) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ModerationFilter) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ModerationFilter) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetIsSpam

`func (o *ModerationFilter) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *ModerationFilter) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *ModerationFilter) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *ModerationFilter) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetIsBannedUser

`func (o *ModerationFilter) GetIsBannedUser() bool`

GetIsBannedUser returns the IsBannedUser field if non-nil, zero value otherwise.

### GetIsBannedUserOk

`func (o *ModerationFilter) GetIsBannedUserOk() (*bool, bool)`

GetIsBannedUserOk returns a tuple with the IsBannedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBannedUser

`func (o *ModerationFilter) SetIsBannedUser(v bool)`

SetIsBannedUser sets IsBannedUser field to given value.

### HasIsBannedUser

`func (o *ModerationFilter) HasIsBannedUser() bool`

HasIsBannedUser returns a boolean if a field has been set.

### GetIsLocked

`func (o *ModerationFilter) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ModerationFilter) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ModerationFilter) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ModerationFilter) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetFlagCountGt

`func (o *ModerationFilter) GetFlagCountGt() float64`

GetFlagCountGt returns the FlagCountGt field if non-nil, zero value otherwise.

### GetFlagCountGtOk

`func (o *ModerationFilter) GetFlagCountGtOk() (*float64, bool)`

GetFlagCountGtOk returns a tuple with the FlagCountGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCountGt

`func (o *ModerationFilter) SetFlagCountGt(v float64)`

SetFlagCountGt sets FlagCountGt field to given value.

### HasFlagCountGt

`func (o *ModerationFilter) HasFlagCountGt() bool`

HasFlagCountGt returns a boolean if a field has been set.

### GetUserId

`func (o *ModerationFilter) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModerationFilter) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModerationFilter) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModerationFilter) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUrlId

`func (o *ModerationFilter) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *ModerationFilter) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *ModerationFilter) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.

### HasUrlId

`func (o *ModerationFilter) HasUrlId() bool`

HasUrlId returns a boolean if a field has been set.

### GetDomain

`func (o *ModerationFilter) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ModerationFilter) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ModerationFilter) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ModerationFilter) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *ModerationFilter) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *ModerationFilter) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *ModerationFilter) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *ModerationFilter) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### GetCommentTextSearch

`func (o *ModerationFilter) GetCommentTextSearch() []string`

GetCommentTextSearch returns the CommentTextSearch field if non-nil, zero value otherwise.

### GetCommentTextSearchOk

`func (o *ModerationFilter) GetCommentTextSearchOk() (*[]string, bool)`

GetCommentTextSearchOk returns a tuple with the CommentTextSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentTextSearch

`func (o *ModerationFilter) SetCommentTextSearch(v []string)`

SetCommentTextSearch sets CommentTextSearch field to given value.

### HasCommentTextSearch

`func (o *ModerationFilter) HasCommentTextSearch() bool`

HasCommentTextSearch returns a boolean if a field has been set.

### GetExactCommentText

`func (o *ModerationFilter) GetExactCommentText() string`

GetExactCommentText returns the ExactCommentText field if non-nil, zero value otherwise.

### GetExactCommentTextOk

`func (o *ModerationFilter) GetExactCommentTextOk() (*string, bool)`

GetExactCommentTextOk returns a tuple with the ExactCommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactCommentText

`func (o *ModerationFilter) SetExactCommentText(v string)`

SetExactCommentText sets ExactCommentText field to given value.

### HasExactCommentText

`func (o *ModerationFilter) HasExactCommentText() bool`

HasExactCommentText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


