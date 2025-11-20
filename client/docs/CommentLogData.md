# CommentLogData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearContent** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**Phrase** | Pointer to **string** |  | [optional] 
**BadWord** | Pointer to **string** |  | [optional] 
**Word** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**TenantBadgeId** | Pointer to **string** |  | [optional] 
**BadgeId** | Pointer to **string** |  | [optional] 
**WasLoggedIn** | Pointer to **bool** |  | [optional] 
**FoundUser** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Engine** | Pointer to **string** |  | [optional] 
**EngineResponse** | Pointer to **string** |  | [optional] 
**EngineTokens** | Pointer to **float64** |  | [optional] 
**TrustFactor** | Pointer to **float64** |  | [optional] 
**Rule** | Pointer to [**SpamRule**](SpamRule.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Subscribers** | Pointer to **float64** |  | [optional] 
**NotificationCount** | Pointer to **float64** |  | [optional] 
**VotesBefore** | Pointer to **NullableFloat64** |  | [optional] 
**VotesUpBefore** | Pointer to **NullableFloat64** |  | [optional] 
**VotesDownBefore** | Pointer to **NullableFloat64** |  | [optional] 
**VotesAfter** | Pointer to **NullableFloat64** |  | [optional] 
**VotesUpAfter** | Pointer to **NullableFloat64** |  | [optional] 
**VotesDownAfter** | Pointer to **NullableFloat64** |  | [optional] 
**RepeatAction** | Pointer to [**RepeatCommentHandlingAction**](RepeatCommentHandlingAction.md) |  | [optional] 
**Reason** | Pointer to [**RepeatCommentCheckIgnoredReason**](RepeatCommentCheckIgnoredReason.md) |  | [optional] 
**OtherData** | Pointer to **interface{}** |  | [optional] 
**SpamBefore** | Pointer to **bool** |  | [optional] 
**SpamAfter** | Pointer to **bool** |  | [optional] 
**PermanentFlag** | Pointer to **string** |  | [optional] 
**ApprovedBefore** | Pointer to **bool** |  | [optional] 
**ApprovedAfter** | Pointer to **bool** |  | [optional] 
**ReviewedBefore** | Pointer to **bool** |  | [optional] 
**ReviewedAfter** | Pointer to **bool** |  | [optional] 
**TextBefore** | Pointer to **string** |  | [optional] 
**TextAfter** | Pointer to **string** |  | [optional] 
**ExpireBefore** | Pointer to **NullableTime** |  | [optional] 
**ExpireAfter** | Pointer to **NullableTime** |  | [optional] 
**FlagCountBefore** | Pointer to **NullableFloat64** |  | [optional] 
**TrustFactorBefore** | Pointer to **float64** |  | [optional] 
**TrustFactorAfter** | Pointer to **float64** |  | [optional] 
**ReferencedCommentId** | Pointer to **string** |  | [optional] 

## Methods

### NewCommentLogData

`func NewCommentLogData() *CommentLogData`

NewCommentLogData instantiates a new CommentLogData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentLogDataWithDefaults

`func NewCommentLogDataWithDefaults() *CommentLogData`

NewCommentLogDataWithDefaults instantiates a new CommentLogData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearContent

`func (o *CommentLogData) GetClearContent() bool`

GetClearContent returns the ClearContent field if non-nil, zero value otherwise.

### GetClearContentOk

`func (o *CommentLogData) GetClearContentOk() (*bool, bool)`

GetClearContentOk returns a tuple with the ClearContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearContent

`func (o *CommentLogData) SetClearContent(v bool)`

SetClearContent sets ClearContent field to given value.

### HasClearContent

`func (o *CommentLogData) HasClearContent() bool`

HasClearContent returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *CommentLogData) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *CommentLogData) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *CommentLogData) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *CommentLogData) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetPhrase

`func (o *CommentLogData) GetPhrase() string`

GetPhrase returns the Phrase field if non-nil, zero value otherwise.

### GetPhraseOk

`func (o *CommentLogData) GetPhraseOk() (*string, bool)`

GetPhraseOk returns a tuple with the Phrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhrase

`func (o *CommentLogData) SetPhrase(v string)`

SetPhrase sets Phrase field to given value.

### HasPhrase

`func (o *CommentLogData) HasPhrase() bool`

HasPhrase returns a boolean if a field has been set.

### GetBadWord

`func (o *CommentLogData) GetBadWord() string`

GetBadWord returns the BadWord field if non-nil, zero value otherwise.

### GetBadWordOk

`func (o *CommentLogData) GetBadWordOk() (*string, bool)`

GetBadWordOk returns a tuple with the BadWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadWord

`func (o *CommentLogData) SetBadWord(v string)`

SetBadWord sets BadWord field to given value.

### HasBadWord

`func (o *CommentLogData) HasBadWord() bool`

HasBadWord returns a boolean if a field has been set.

### GetWord

`func (o *CommentLogData) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *CommentLogData) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *CommentLogData) SetWord(v string)`

SetWord sets Word field to given value.

### HasWord

`func (o *CommentLogData) HasWord() bool`

HasWord returns a boolean if a field has been set.

### GetLocale

`func (o *CommentLogData) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CommentLogData) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CommentLogData) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CommentLogData) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTenantBadgeId

`func (o *CommentLogData) GetTenantBadgeId() string`

GetTenantBadgeId returns the TenantBadgeId field if non-nil, zero value otherwise.

### GetTenantBadgeIdOk

`func (o *CommentLogData) GetTenantBadgeIdOk() (*string, bool)`

GetTenantBadgeIdOk returns a tuple with the TenantBadgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantBadgeId

`func (o *CommentLogData) SetTenantBadgeId(v string)`

SetTenantBadgeId sets TenantBadgeId field to given value.

### HasTenantBadgeId

`func (o *CommentLogData) HasTenantBadgeId() bool`

HasTenantBadgeId returns a boolean if a field has been set.

### GetBadgeId

`func (o *CommentLogData) GetBadgeId() string`

GetBadgeId returns the BadgeId field if non-nil, zero value otherwise.

### GetBadgeIdOk

`func (o *CommentLogData) GetBadgeIdOk() (*string, bool)`

GetBadgeIdOk returns a tuple with the BadgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeId

`func (o *CommentLogData) SetBadgeId(v string)`

SetBadgeId sets BadgeId field to given value.

### HasBadgeId

`func (o *CommentLogData) HasBadgeId() bool`

HasBadgeId returns a boolean if a field has been set.

### GetWasLoggedIn

`func (o *CommentLogData) GetWasLoggedIn() bool`

GetWasLoggedIn returns the WasLoggedIn field if non-nil, zero value otherwise.

### GetWasLoggedInOk

`func (o *CommentLogData) GetWasLoggedInOk() (*bool, bool)`

GetWasLoggedInOk returns a tuple with the WasLoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasLoggedIn

`func (o *CommentLogData) SetWasLoggedIn(v bool)`

SetWasLoggedIn sets WasLoggedIn field to given value.

### HasWasLoggedIn

`func (o *CommentLogData) HasWasLoggedIn() bool`

HasWasLoggedIn returns a boolean if a field has been set.

### GetFoundUser

`func (o *CommentLogData) GetFoundUser() bool`

GetFoundUser returns the FoundUser field if non-nil, zero value otherwise.

### GetFoundUserOk

`func (o *CommentLogData) GetFoundUserOk() (*bool, bool)`

GetFoundUserOk returns a tuple with the FoundUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundUser

`func (o *CommentLogData) SetFoundUser(v bool)`

SetFoundUser sets FoundUser field to given value.

### HasFoundUser

`func (o *CommentLogData) HasFoundUser() bool`

HasFoundUser returns a boolean if a field has been set.

### GetVerified

`func (o *CommentLogData) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *CommentLogData) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *CommentLogData) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *CommentLogData) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetEngine

`func (o *CommentLogData) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *CommentLogData) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *CommentLogData) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *CommentLogData) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetEngineResponse

`func (o *CommentLogData) GetEngineResponse() string`

GetEngineResponse returns the EngineResponse field if non-nil, zero value otherwise.

### GetEngineResponseOk

`func (o *CommentLogData) GetEngineResponseOk() (*string, bool)`

GetEngineResponseOk returns a tuple with the EngineResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineResponse

`func (o *CommentLogData) SetEngineResponse(v string)`

SetEngineResponse sets EngineResponse field to given value.

### HasEngineResponse

`func (o *CommentLogData) HasEngineResponse() bool`

HasEngineResponse returns a boolean if a field has been set.

### GetEngineTokens

`func (o *CommentLogData) GetEngineTokens() float64`

GetEngineTokens returns the EngineTokens field if non-nil, zero value otherwise.

### GetEngineTokensOk

`func (o *CommentLogData) GetEngineTokensOk() (*float64, bool)`

GetEngineTokensOk returns a tuple with the EngineTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineTokens

`func (o *CommentLogData) SetEngineTokens(v float64)`

SetEngineTokens sets EngineTokens field to given value.

### HasEngineTokens

`func (o *CommentLogData) HasEngineTokens() bool`

HasEngineTokens returns a boolean if a field has been set.

### GetTrustFactor

`func (o *CommentLogData) GetTrustFactor() float64`

GetTrustFactor returns the TrustFactor field if non-nil, zero value otherwise.

### GetTrustFactorOk

`func (o *CommentLogData) GetTrustFactorOk() (*float64, bool)`

GetTrustFactorOk returns a tuple with the TrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFactor

`func (o *CommentLogData) SetTrustFactor(v float64)`

SetTrustFactor sets TrustFactor field to given value.

### HasTrustFactor

`func (o *CommentLogData) HasTrustFactor() bool`

HasTrustFactor returns a boolean if a field has been set.

### GetRule

`func (o *CommentLogData) GetRule() SpamRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CommentLogData) GetRuleOk() (*SpamRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CommentLogData) SetRule(v SpamRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *CommentLogData) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetUserId

`func (o *CommentLogData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommentLogData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommentLogData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommentLogData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSubscribers

`func (o *CommentLogData) GetSubscribers() float64`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *CommentLogData) GetSubscribersOk() (*float64, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *CommentLogData) SetSubscribers(v float64)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *CommentLogData) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetNotificationCount

`func (o *CommentLogData) GetNotificationCount() float64`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *CommentLogData) GetNotificationCountOk() (*float64, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *CommentLogData) SetNotificationCount(v float64)`

SetNotificationCount sets NotificationCount field to given value.

### HasNotificationCount

`func (o *CommentLogData) HasNotificationCount() bool`

HasNotificationCount returns a boolean if a field has been set.

### GetVotesBefore

`func (o *CommentLogData) GetVotesBefore() float64`

GetVotesBefore returns the VotesBefore field if non-nil, zero value otherwise.

### GetVotesBeforeOk

`func (o *CommentLogData) GetVotesBeforeOk() (*float64, bool)`

GetVotesBeforeOk returns a tuple with the VotesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesBefore

`func (o *CommentLogData) SetVotesBefore(v float64)`

SetVotesBefore sets VotesBefore field to given value.

### HasVotesBefore

`func (o *CommentLogData) HasVotesBefore() bool`

HasVotesBefore returns a boolean if a field has been set.

### SetVotesBeforeNil

`func (o *CommentLogData) SetVotesBeforeNil(b bool)`

 SetVotesBeforeNil sets the value for VotesBefore to be an explicit nil

### UnsetVotesBefore
`func (o *CommentLogData) UnsetVotesBefore()`

UnsetVotesBefore ensures that no value is present for VotesBefore, not even an explicit nil
### GetVotesUpBefore

`func (o *CommentLogData) GetVotesUpBefore() float64`

GetVotesUpBefore returns the VotesUpBefore field if non-nil, zero value otherwise.

### GetVotesUpBeforeOk

`func (o *CommentLogData) GetVotesUpBeforeOk() (*float64, bool)`

GetVotesUpBeforeOk returns a tuple with the VotesUpBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUpBefore

`func (o *CommentLogData) SetVotesUpBefore(v float64)`

SetVotesUpBefore sets VotesUpBefore field to given value.

### HasVotesUpBefore

`func (o *CommentLogData) HasVotesUpBefore() bool`

HasVotesUpBefore returns a boolean if a field has been set.

### SetVotesUpBeforeNil

`func (o *CommentLogData) SetVotesUpBeforeNil(b bool)`

 SetVotesUpBeforeNil sets the value for VotesUpBefore to be an explicit nil

### UnsetVotesUpBefore
`func (o *CommentLogData) UnsetVotesUpBefore()`

UnsetVotesUpBefore ensures that no value is present for VotesUpBefore, not even an explicit nil
### GetVotesDownBefore

`func (o *CommentLogData) GetVotesDownBefore() float64`

GetVotesDownBefore returns the VotesDownBefore field if non-nil, zero value otherwise.

### GetVotesDownBeforeOk

`func (o *CommentLogData) GetVotesDownBeforeOk() (*float64, bool)`

GetVotesDownBeforeOk returns a tuple with the VotesDownBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDownBefore

`func (o *CommentLogData) SetVotesDownBefore(v float64)`

SetVotesDownBefore sets VotesDownBefore field to given value.

### HasVotesDownBefore

`func (o *CommentLogData) HasVotesDownBefore() bool`

HasVotesDownBefore returns a boolean if a field has been set.

### SetVotesDownBeforeNil

`func (o *CommentLogData) SetVotesDownBeforeNil(b bool)`

 SetVotesDownBeforeNil sets the value for VotesDownBefore to be an explicit nil

### UnsetVotesDownBefore
`func (o *CommentLogData) UnsetVotesDownBefore()`

UnsetVotesDownBefore ensures that no value is present for VotesDownBefore, not even an explicit nil
### GetVotesAfter

`func (o *CommentLogData) GetVotesAfter() float64`

GetVotesAfter returns the VotesAfter field if non-nil, zero value otherwise.

### GetVotesAfterOk

`func (o *CommentLogData) GetVotesAfterOk() (*float64, bool)`

GetVotesAfterOk returns a tuple with the VotesAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesAfter

`func (o *CommentLogData) SetVotesAfter(v float64)`

SetVotesAfter sets VotesAfter field to given value.

### HasVotesAfter

`func (o *CommentLogData) HasVotesAfter() bool`

HasVotesAfter returns a boolean if a field has been set.

### SetVotesAfterNil

`func (o *CommentLogData) SetVotesAfterNil(b bool)`

 SetVotesAfterNil sets the value for VotesAfter to be an explicit nil

### UnsetVotesAfter
`func (o *CommentLogData) UnsetVotesAfter()`

UnsetVotesAfter ensures that no value is present for VotesAfter, not even an explicit nil
### GetVotesUpAfter

`func (o *CommentLogData) GetVotesUpAfter() float64`

GetVotesUpAfter returns the VotesUpAfter field if non-nil, zero value otherwise.

### GetVotesUpAfterOk

`func (o *CommentLogData) GetVotesUpAfterOk() (*float64, bool)`

GetVotesUpAfterOk returns a tuple with the VotesUpAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUpAfter

`func (o *CommentLogData) SetVotesUpAfter(v float64)`

SetVotesUpAfter sets VotesUpAfter field to given value.

### HasVotesUpAfter

`func (o *CommentLogData) HasVotesUpAfter() bool`

HasVotesUpAfter returns a boolean if a field has been set.

### SetVotesUpAfterNil

`func (o *CommentLogData) SetVotesUpAfterNil(b bool)`

 SetVotesUpAfterNil sets the value for VotesUpAfter to be an explicit nil

### UnsetVotesUpAfter
`func (o *CommentLogData) UnsetVotesUpAfter()`

UnsetVotesUpAfter ensures that no value is present for VotesUpAfter, not even an explicit nil
### GetVotesDownAfter

`func (o *CommentLogData) GetVotesDownAfter() float64`

GetVotesDownAfter returns the VotesDownAfter field if non-nil, zero value otherwise.

### GetVotesDownAfterOk

`func (o *CommentLogData) GetVotesDownAfterOk() (*float64, bool)`

GetVotesDownAfterOk returns a tuple with the VotesDownAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDownAfter

`func (o *CommentLogData) SetVotesDownAfter(v float64)`

SetVotesDownAfter sets VotesDownAfter field to given value.

### HasVotesDownAfter

`func (o *CommentLogData) HasVotesDownAfter() bool`

HasVotesDownAfter returns a boolean if a field has been set.

### SetVotesDownAfterNil

`func (o *CommentLogData) SetVotesDownAfterNil(b bool)`

 SetVotesDownAfterNil sets the value for VotesDownAfter to be an explicit nil

### UnsetVotesDownAfter
`func (o *CommentLogData) UnsetVotesDownAfter()`

UnsetVotesDownAfter ensures that no value is present for VotesDownAfter, not even an explicit nil
### GetRepeatAction

`func (o *CommentLogData) GetRepeatAction() RepeatCommentHandlingAction`

GetRepeatAction returns the RepeatAction field if non-nil, zero value otherwise.

### GetRepeatActionOk

`func (o *CommentLogData) GetRepeatActionOk() (*RepeatCommentHandlingAction, bool)`

GetRepeatActionOk returns a tuple with the RepeatAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatAction

`func (o *CommentLogData) SetRepeatAction(v RepeatCommentHandlingAction)`

SetRepeatAction sets RepeatAction field to given value.

### HasRepeatAction

`func (o *CommentLogData) HasRepeatAction() bool`

HasRepeatAction returns a boolean if a field has been set.

### GetReason

`func (o *CommentLogData) GetReason() RepeatCommentCheckIgnoredReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CommentLogData) GetReasonOk() (*RepeatCommentCheckIgnoredReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CommentLogData) SetReason(v RepeatCommentCheckIgnoredReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CommentLogData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetOtherData

`func (o *CommentLogData) GetOtherData() interface{}`

GetOtherData returns the OtherData field if non-nil, zero value otherwise.

### GetOtherDataOk

`func (o *CommentLogData) GetOtherDataOk() (*interface{}, bool)`

GetOtherDataOk returns a tuple with the OtherData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherData

`func (o *CommentLogData) SetOtherData(v interface{})`

SetOtherData sets OtherData field to given value.

### HasOtherData

`func (o *CommentLogData) HasOtherData() bool`

HasOtherData returns a boolean if a field has been set.

### SetOtherDataNil

`func (o *CommentLogData) SetOtherDataNil(b bool)`

 SetOtherDataNil sets the value for OtherData to be an explicit nil

### UnsetOtherData
`func (o *CommentLogData) UnsetOtherData()`

UnsetOtherData ensures that no value is present for OtherData, not even an explicit nil
### GetSpamBefore

`func (o *CommentLogData) GetSpamBefore() bool`

GetSpamBefore returns the SpamBefore field if non-nil, zero value otherwise.

### GetSpamBeforeOk

`func (o *CommentLogData) GetSpamBeforeOk() (*bool, bool)`

GetSpamBeforeOk returns a tuple with the SpamBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamBefore

`func (o *CommentLogData) SetSpamBefore(v bool)`

SetSpamBefore sets SpamBefore field to given value.

### HasSpamBefore

`func (o *CommentLogData) HasSpamBefore() bool`

HasSpamBefore returns a boolean if a field has been set.

### GetSpamAfter

`func (o *CommentLogData) GetSpamAfter() bool`

GetSpamAfter returns the SpamAfter field if non-nil, zero value otherwise.

### GetSpamAfterOk

`func (o *CommentLogData) GetSpamAfterOk() (*bool, bool)`

GetSpamAfterOk returns a tuple with the SpamAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamAfter

`func (o *CommentLogData) SetSpamAfter(v bool)`

SetSpamAfter sets SpamAfter field to given value.

### HasSpamAfter

`func (o *CommentLogData) HasSpamAfter() bool`

HasSpamAfter returns a boolean if a field has been set.

### GetPermanentFlag

`func (o *CommentLogData) GetPermanentFlag() string`

GetPermanentFlag returns the PermanentFlag field if non-nil, zero value otherwise.

### GetPermanentFlagOk

`func (o *CommentLogData) GetPermanentFlagOk() (*string, bool)`

GetPermanentFlagOk returns a tuple with the PermanentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentFlag

`func (o *CommentLogData) SetPermanentFlag(v string)`

SetPermanentFlag sets PermanentFlag field to given value.

### HasPermanentFlag

`func (o *CommentLogData) HasPermanentFlag() bool`

HasPermanentFlag returns a boolean if a field has been set.

### GetApprovedBefore

`func (o *CommentLogData) GetApprovedBefore() bool`

GetApprovedBefore returns the ApprovedBefore field if non-nil, zero value otherwise.

### GetApprovedBeforeOk

`func (o *CommentLogData) GetApprovedBeforeOk() (*bool, bool)`

GetApprovedBeforeOk returns a tuple with the ApprovedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBefore

`func (o *CommentLogData) SetApprovedBefore(v bool)`

SetApprovedBefore sets ApprovedBefore field to given value.

### HasApprovedBefore

`func (o *CommentLogData) HasApprovedBefore() bool`

HasApprovedBefore returns a boolean if a field has been set.

### GetApprovedAfter

`func (o *CommentLogData) GetApprovedAfter() bool`

GetApprovedAfter returns the ApprovedAfter field if non-nil, zero value otherwise.

### GetApprovedAfterOk

`func (o *CommentLogData) GetApprovedAfterOk() (*bool, bool)`

GetApprovedAfterOk returns a tuple with the ApprovedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAfter

`func (o *CommentLogData) SetApprovedAfter(v bool)`

SetApprovedAfter sets ApprovedAfter field to given value.

### HasApprovedAfter

`func (o *CommentLogData) HasApprovedAfter() bool`

HasApprovedAfter returns a boolean if a field has been set.

### GetReviewedBefore

`func (o *CommentLogData) GetReviewedBefore() bool`

GetReviewedBefore returns the ReviewedBefore field if non-nil, zero value otherwise.

### GetReviewedBeforeOk

`func (o *CommentLogData) GetReviewedBeforeOk() (*bool, bool)`

GetReviewedBeforeOk returns a tuple with the ReviewedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBefore

`func (o *CommentLogData) SetReviewedBefore(v bool)`

SetReviewedBefore sets ReviewedBefore field to given value.

### HasReviewedBefore

`func (o *CommentLogData) HasReviewedBefore() bool`

HasReviewedBefore returns a boolean if a field has been set.

### GetReviewedAfter

`func (o *CommentLogData) GetReviewedAfter() bool`

GetReviewedAfter returns the ReviewedAfter field if non-nil, zero value otherwise.

### GetReviewedAfterOk

`func (o *CommentLogData) GetReviewedAfterOk() (*bool, bool)`

GetReviewedAfterOk returns a tuple with the ReviewedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAfter

`func (o *CommentLogData) SetReviewedAfter(v bool)`

SetReviewedAfter sets ReviewedAfter field to given value.

### HasReviewedAfter

`func (o *CommentLogData) HasReviewedAfter() bool`

HasReviewedAfter returns a boolean if a field has been set.

### GetTextBefore

`func (o *CommentLogData) GetTextBefore() string`

GetTextBefore returns the TextBefore field if non-nil, zero value otherwise.

### GetTextBeforeOk

`func (o *CommentLogData) GetTextBeforeOk() (*string, bool)`

GetTextBeforeOk returns a tuple with the TextBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBefore

`func (o *CommentLogData) SetTextBefore(v string)`

SetTextBefore sets TextBefore field to given value.

### HasTextBefore

`func (o *CommentLogData) HasTextBefore() bool`

HasTextBefore returns a boolean if a field has been set.

### GetTextAfter

`func (o *CommentLogData) GetTextAfter() string`

GetTextAfter returns the TextAfter field if non-nil, zero value otherwise.

### GetTextAfterOk

`func (o *CommentLogData) GetTextAfterOk() (*string, bool)`

GetTextAfterOk returns a tuple with the TextAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAfter

`func (o *CommentLogData) SetTextAfter(v string)`

SetTextAfter sets TextAfter field to given value.

### HasTextAfter

`func (o *CommentLogData) HasTextAfter() bool`

HasTextAfter returns a boolean if a field has been set.

### GetExpireBefore

`func (o *CommentLogData) GetExpireBefore() time.Time`

GetExpireBefore returns the ExpireBefore field if non-nil, zero value otherwise.

### GetExpireBeforeOk

`func (o *CommentLogData) GetExpireBeforeOk() (*time.Time, bool)`

GetExpireBeforeOk returns a tuple with the ExpireBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireBefore

`func (o *CommentLogData) SetExpireBefore(v time.Time)`

SetExpireBefore sets ExpireBefore field to given value.

### HasExpireBefore

`func (o *CommentLogData) HasExpireBefore() bool`

HasExpireBefore returns a boolean if a field has been set.

### SetExpireBeforeNil

`func (o *CommentLogData) SetExpireBeforeNil(b bool)`

 SetExpireBeforeNil sets the value for ExpireBefore to be an explicit nil

### UnsetExpireBefore
`func (o *CommentLogData) UnsetExpireBefore()`

UnsetExpireBefore ensures that no value is present for ExpireBefore, not even an explicit nil
### GetExpireAfter

`func (o *CommentLogData) GetExpireAfter() time.Time`

GetExpireAfter returns the ExpireAfter field if non-nil, zero value otherwise.

### GetExpireAfterOk

`func (o *CommentLogData) GetExpireAfterOk() (*time.Time, bool)`

GetExpireAfterOk returns a tuple with the ExpireAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfter

`func (o *CommentLogData) SetExpireAfter(v time.Time)`

SetExpireAfter sets ExpireAfter field to given value.

### HasExpireAfter

`func (o *CommentLogData) HasExpireAfter() bool`

HasExpireAfter returns a boolean if a field has been set.

### SetExpireAfterNil

`func (o *CommentLogData) SetExpireAfterNil(b bool)`

 SetExpireAfterNil sets the value for ExpireAfter to be an explicit nil

### UnsetExpireAfter
`func (o *CommentLogData) UnsetExpireAfter()`

UnsetExpireAfter ensures that no value is present for ExpireAfter, not even an explicit nil
### GetFlagCountBefore

`func (o *CommentLogData) GetFlagCountBefore() float64`

GetFlagCountBefore returns the FlagCountBefore field if non-nil, zero value otherwise.

### GetFlagCountBeforeOk

`func (o *CommentLogData) GetFlagCountBeforeOk() (*float64, bool)`

GetFlagCountBeforeOk returns a tuple with the FlagCountBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCountBefore

`func (o *CommentLogData) SetFlagCountBefore(v float64)`

SetFlagCountBefore sets FlagCountBefore field to given value.

### HasFlagCountBefore

`func (o *CommentLogData) HasFlagCountBefore() bool`

HasFlagCountBefore returns a boolean if a field has been set.

### SetFlagCountBeforeNil

`func (o *CommentLogData) SetFlagCountBeforeNil(b bool)`

 SetFlagCountBeforeNil sets the value for FlagCountBefore to be an explicit nil

### UnsetFlagCountBefore
`func (o *CommentLogData) UnsetFlagCountBefore()`

UnsetFlagCountBefore ensures that no value is present for FlagCountBefore, not even an explicit nil
### GetTrustFactorBefore

`func (o *CommentLogData) GetTrustFactorBefore() float64`

GetTrustFactorBefore returns the TrustFactorBefore field if non-nil, zero value otherwise.

### GetTrustFactorBeforeOk

`func (o *CommentLogData) GetTrustFactorBeforeOk() (*float64, bool)`

GetTrustFactorBeforeOk returns a tuple with the TrustFactorBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFactorBefore

`func (o *CommentLogData) SetTrustFactorBefore(v float64)`

SetTrustFactorBefore sets TrustFactorBefore field to given value.

### HasTrustFactorBefore

`func (o *CommentLogData) HasTrustFactorBefore() bool`

HasTrustFactorBefore returns a boolean if a field has been set.

### GetTrustFactorAfter

`func (o *CommentLogData) GetTrustFactorAfter() float64`

GetTrustFactorAfter returns the TrustFactorAfter field if non-nil, zero value otherwise.

### GetTrustFactorAfterOk

`func (o *CommentLogData) GetTrustFactorAfterOk() (*float64, bool)`

GetTrustFactorAfterOk returns a tuple with the TrustFactorAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFactorAfter

`func (o *CommentLogData) SetTrustFactorAfter(v float64)`

SetTrustFactorAfter sets TrustFactorAfter field to given value.

### HasTrustFactorAfter

`func (o *CommentLogData) HasTrustFactorAfter() bool`

HasTrustFactorAfter returns a boolean if a field has been set.

### GetReferencedCommentId

`func (o *CommentLogData) GetReferencedCommentId() string`

GetReferencedCommentId returns the ReferencedCommentId field if non-nil, zero value otherwise.

### GetReferencedCommentIdOk

`func (o *CommentLogData) GetReferencedCommentIdOk() (*string, bool)`

GetReferencedCommentIdOk returns a tuple with the ReferencedCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedCommentId

`func (o *CommentLogData) SetReferencedCommentId(v string)`

SetReferencedCommentId sets ReferencedCommentId field to given value.

### HasReferencedCommentId

`func (o *CommentLogData) HasReferencedCommentId() bool`

HasReferencedCommentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


