# Moderator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Name** | **NullableString** |  | 
**UserId** | **NullableString** |  | 
**AcceptedInvite** | **bool** |  | 
**Email** | **NullableString** |  | 
**MarkReviewedCount** | **float64** |  | 
**DeletedCount** | **float64** |  | 
**MarkedSpamCount** | **float64** |  | 
**MarkedNotSpamCount** | **float64** |  | 
**ApprovedCount** | **float64** |  | 
**UnApprovedCount** | **float64** |  | 
**EditedCount** | **float64** |  | 
**BannedCount** | **float64** |  | 
**UnFlaggedCount** | **float64** |  | 
**VerificationId** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**ModerationGroupIds** | **[]string** |  | 
**IsEmailSuppressed** | Pointer to **bool** |  | [optional] 

## Methods

### NewModerator

`func NewModerator(id string, tenantId string, name NullableString, userId NullableString, acceptedInvite bool, email NullableString, markReviewedCount float64, deletedCount float64, markedSpamCount float64, markedNotSpamCount float64, approvedCount float64, unApprovedCount float64, editedCount float64, bannedCount float64, unFlaggedCount float64, verificationId NullableString, createdAt time.Time, moderationGroupIds []string, ) *Moderator`

NewModerator instantiates a new Moderator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModeratorWithDefaults

`func NewModeratorWithDefaults() *Moderator`

NewModeratorWithDefaults instantiates a new Moderator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Moderator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Moderator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Moderator) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Moderator) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Moderator) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Moderator) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *Moderator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Moderator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Moderator) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Moderator) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Moderator) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserId

`func (o *Moderator) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Moderator) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Moderator) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *Moderator) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *Moderator) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAcceptedInvite

`func (o *Moderator) GetAcceptedInvite() bool`

GetAcceptedInvite returns the AcceptedInvite field if non-nil, zero value otherwise.

### GetAcceptedInviteOk

`func (o *Moderator) GetAcceptedInviteOk() (*bool, bool)`

GetAcceptedInviteOk returns a tuple with the AcceptedInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedInvite

`func (o *Moderator) SetAcceptedInvite(v bool)`

SetAcceptedInvite sets AcceptedInvite field to given value.


### GetEmail

`func (o *Moderator) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Moderator) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Moderator) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *Moderator) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Moderator) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMarkReviewedCount

`func (o *Moderator) GetMarkReviewedCount() float64`

GetMarkReviewedCount returns the MarkReviewedCount field if non-nil, zero value otherwise.

### GetMarkReviewedCountOk

`func (o *Moderator) GetMarkReviewedCountOk() (*float64, bool)`

GetMarkReviewedCountOk returns a tuple with the MarkReviewedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkReviewedCount

`func (o *Moderator) SetMarkReviewedCount(v float64)`

SetMarkReviewedCount sets MarkReviewedCount field to given value.


### GetDeletedCount

`func (o *Moderator) GetDeletedCount() float64`

GetDeletedCount returns the DeletedCount field if non-nil, zero value otherwise.

### GetDeletedCountOk

`func (o *Moderator) GetDeletedCountOk() (*float64, bool)`

GetDeletedCountOk returns a tuple with the DeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedCount

`func (o *Moderator) SetDeletedCount(v float64)`

SetDeletedCount sets DeletedCount field to given value.


### GetMarkedSpamCount

`func (o *Moderator) GetMarkedSpamCount() float64`

GetMarkedSpamCount returns the MarkedSpamCount field if non-nil, zero value otherwise.

### GetMarkedSpamCountOk

`func (o *Moderator) GetMarkedSpamCountOk() (*float64, bool)`

GetMarkedSpamCountOk returns a tuple with the MarkedSpamCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedSpamCount

`func (o *Moderator) SetMarkedSpamCount(v float64)`

SetMarkedSpamCount sets MarkedSpamCount field to given value.


### GetMarkedNotSpamCount

`func (o *Moderator) GetMarkedNotSpamCount() float64`

GetMarkedNotSpamCount returns the MarkedNotSpamCount field if non-nil, zero value otherwise.

### GetMarkedNotSpamCountOk

`func (o *Moderator) GetMarkedNotSpamCountOk() (*float64, bool)`

GetMarkedNotSpamCountOk returns a tuple with the MarkedNotSpamCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedNotSpamCount

`func (o *Moderator) SetMarkedNotSpamCount(v float64)`

SetMarkedNotSpamCount sets MarkedNotSpamCount field to given value.


### GetApprovedCount

`func (o *Moderator) GetApprovedCount() float64`

GetApprovedCount returns the ApprovedCount field if non-nil, zero value otherwise.

### GetApprovedCountOk

`func (o *Moderator) GetApprovedCountOk() (*float64, bool)`

GetApprovedCountOk returns a tuple with the ApprovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedCount

`func (o *Moderator) SetApprovedCount(v float64)`

SetApprovedCount sets ApprovedCount field to given value.


### GetUnApprovedCount

`func (o *Moderator) GetUnApprovedCount() float64`

GetUnApprovedCount returns the UnApprovedCount field if non-nil, zero value otherwise.

### GetUnApprovedCountOk

`func (o *Moderator) GetUnApprovedCountOk() (*float64, bool)`

GetUnApprovedCountOk returns a tuple with the UnApprovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnApprovedCount

`func (o *Moderator) SetUnApprovedCount(v float64)`

SetUnApprovedCount sets UnApprovedCount field to given value.


### GetEditedCount

`func (o *Moderator) GetEditedCount() float64`

GetEditedCount returns the EditedCount field if non-nil, zero value otherwise.

### GetEditedCountOk

`func (o *Moderator) GetEditedCountOk() (*float64, bool)`

GetEditedCountOk returns a tuple with the EditedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedCount

`func (o *Moderator) SetEditedCount(v float64)`

SetEditedCount sets EditedCount field to given value.


### GetBannedCount

`func (o *Moderator) GetBannedCount() float64`

GetBannedCount returns the BannedCount field if non-nil, zero value otherwise.

### GetBannedCountOk

`func (o *Moderator) GetBannedCountOk() (*float64, bool)`

GetBannedCountOk returns a tuple with the BannedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedCount

`func (o *Moderator) SetBannedCount(v float64)`

SetBannedCount sets BannedCount field to given value.


### GetUnFlaggedCount

`func (o *Moderator) GetUnFlaggedCount() float64`

GetUnFlaggedCount returns the UnFlaggedCount field if non-nil, zero value otherwise.

### GetUnFlaggedCountOk

`func (o *Moderator) GetUnFlaggedCountOk() (*float64, bool)`

GetUnFlaggedCountOk returns a tuple with the UnFlaggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnFlaggedCount

`func (o *Moderator) SetUnFlaggedCount(v float64)`

SetUnFlaggedCount sets UnFlaggedCount field to given value.


### GetVerificationId

`func (o *Moderator) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *Moderator) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *Moderator) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.


### SetVerificationIdNil

`func (o *Moderator) SetVerificationIdNil(b bool)`

 SetVerificationIdNil sets the value for VerificationId to be an explicit nil

### UnsetVerificationId
`func (o *Moderator) UnsetVerificationId()`

UnsetVerificationId ensures that no value is present for VerificationId, not even an explicit nil
### GetCreatedAt

`func (o *Moderator) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Moderator) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Moderator) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModerationGroupIds

`func (o *Moderator) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *Moderator) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *Moderator) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.


### SetModerationGroupIdsNil

`func (o *Moderator) SetModerationGroupIdsNil(b bool)`

 SetModerationGroupIdsNil sets the value for ModerationGroupIds to be an explicit nil

### UnsetModerationGroupIds
`func (o *Moderator) UnsetModerationGroupIds()`

UnsetModerationGroupIds ensures that no value is present for ModerationGroupIds, not even an explicit nil
### GetIsEmailSuppressed

`func (o *Moderator) GetIsEmailSuppressed() bool`

GetIsEmailSuppressed returns the IsEmailSuppressed field if non-nil, zero value otherwise.

### GetIsEmailSuppressedOk

`func (o *Moderator) GetIsEmailSuppressedOk() (*bool, bool)`

GetIsEmailSuppressedOk returns a tuple with the IsEmailSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailSuppressed

`func (o *Moderator) SetIsEmailSuppressed(v bool)`

SetIsEmailSuppressed sets IsEmailSuppressed field to given value.

### HasIsEmailSuppressed

`func (o *Moderator) HasIsEmailSuppressed() bool`

HasIsEmailSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


