# UpdatableCommentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlId** | Pointer to **string** |  | [optional] 
**UrlIdRaw** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**PageTitle** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**CommenterEmail** | Pointer to **NullableString** |  | [optional] 
**CommenterName** | Pointer to **string** |  | [optional] 
**CommenterLink** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CommentHTML** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **NullableFloat64** |  | [optional] 
**LocalDateString** | Pointer to **NullableString** |  | [optional] 
**LocalDateHours** | Pointer to **NullableInt32** |  | [optional] 
**Votes** | Pointer to **NullableInt32** |  | [optional] 
**VotesUp** | Pointer to **NullableInt32** |  | [optional] 
**VotesDown** | Pointer to **NullableInt32** |  | [optional] 
**ExpireAt** | Pointer to **NullableTime** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**VerifiedDate** | Pointer to **NullableTime** |  | [optional] 
**NotificationSentForParent** | Pointer to **bool** |  | [optional] 
**NotificationSentForParentTenant** | Pointer to **bool** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalParentId** | Pointer to **NullableString** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **NullableBool** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**FlagCount** | Pointer to **NullableInt32** |  | [optional] 
**DisplayLabel** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to [**NullableFCommentMeta**](FCommentMeta.md) |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdatableCommentParams

`func NewUpdatableCommentParams() *UpdatableCommentParams`

NewUpdatableCommentParams instantiates a new UpdatableCommentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatableCommentParamsWithDefaults

`func NewUpdatableCommentParamsWithDefaults() *UpdatableCommentParams`

NewUpdatableCommentParamsWithDefaults instantiates a new UpdatableCommentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlId

`func (o *UpdatableCommentParams) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *UpdatableCommentParams) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *UpdatableCommentParams) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.

### HasUrlId

`func (o *UpdatableCommentParams) HasUrlId() bool`

HasUrlId returns a boolean if a field has been set.

### GetUrlIdRaw

`func (o *UpdatableCommentParams) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *UpdatableCommentParams) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *UpdatableCommentParams) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.

### HasUrlIdRaw

`func (o *UpdatableCommentParams) HasUrlIdRaw() bool`

HasUrlIdRaw returns a boolean if a field has been set.

### GetUrl

`func (o *UpdatableCommentParams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdatableCommentParams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdatableCommentParams) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdatableCommentParams) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPageTitle

`func (o *UpdatableCommentParams) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *UpdatableCommentParams) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *UpdatableCommentParams) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *UpdatableCommentParams) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *UpdatableCommentParams) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *UpdatableCommentParams) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetUserId

`func (o *UpdatableCommentParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdatableCommentParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdatableCommentParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UpdatableCommentParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UpdatableCommentParams) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UpdatableCommentParams) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetCommenterEmail

`func (o *UpdatableCommentParams) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *UpdatableCommentParams) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *UpdatableCommentParams) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *UpdatableCommentParams) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *UpdatableCommentParams) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *UpdatableCommentParams) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetCommenterName

`func (o *UpdatableCommentParams) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *UpdatableCommentParams) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *UpdatableCommentParams) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.

### HasCommenterName

`func (o *UpdatableCommentParams) HasCommenterName() bool`

HasCommenterName returns a boolean if a field has been set.

### GetCommenterLink

`func (o *UpdatableCommentParams) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *UpdatableCommentParams) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *UpdatableCommentParams) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *UpdatableCommentParams) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *UpdatableCommentParams) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *UpdatableCommentParams) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetComment

`func (o *UpdatableCommentParams) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdatableCommentParams) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdatableCommentParams) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdatableCommentParams) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommentHTML

`func (o *UpdatableCommentParams) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *UpdatableCommentParams) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *UpdatableCommentParams) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.

### HasCommentHTML

`func (o *UpdatableCommentParams) HasCommentHTML() bool`

HasCommentHTML returns a boolean if a field has been set.

### GetParentId

`func (o *UpdatableCommentParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdatableCommentParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdatableCommentParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdatableCommentParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *UpdatableCommentParams) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *UpdatableCommentParams) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetDate

`func (o *UpdatableCommentParams) GetDate() float64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UpdatableCommentParams) GetDateOk() (*float64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UpdatableCommentParams) SetDate(v float64)`

SetDate sets Date field to given value.

### HasDate

`func (o *UpdatableCommentParams) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *UpdatableCommentParams) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *UpdatableCommentParams) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetLocalDateString

`func (o *UpdatableCommentParams) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *UpdatableCommentParams) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *UpdatableCommentParams) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *UpdatableCommentParams) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### SetLocalDateStringNil

`func (o *UpdatableCommentParams) SetLocalDateStringNil(b bool)`

 SetLocalDateStringNil sets the value for LocalDateString to be an explicit nil

### UnsetLocalDateString
`func (o *UpdatableCommentParams) UnsetLocalDateString()`

UnsetLocalDateString ensures that no value is present for LocalDateString, not even an explicit nil
### GetLocalDateHours

`func (o *UpdatableCommentParams) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *UpdatableCommentParams) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *UpdatableCommentParams) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *UpdatableCommentParams) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### SetLocalDateHoursNil

`func (o *UpdatableCommentParams) SetLocalDateHoursNil(b bool)`

 SetLocalDateHoursNil sets the value for LocalDateHours to be an explicit nil

### UnsetLocalDateHours
`func (o *UpdatableCommentParams) UnsetLocalDateHours()`

UnsetLocalDateHours ensures that no value is present for LocalDateHours, not even an explicit nil
### GetVotes

`func (o *UpdatableCommentParams) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *UpdatableCommentParams) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *UpdatableCommentParams) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *UpdatableCommentParams) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### SetVotesNil

`func (o *UpdatableCommentParams) SetVotesNil(b bool)`

 SetVotesNil sets the value for Votes to be an explicit nil

### UnsetVotes
`func (o *UpdatableCommentParams) UnsetVotes()`

UnsetVotes ensures that no value is present for Votes, not even an explicit nil
### GetVotesUp

`func (o *UpdatableCommentParams) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *UpdatableCommentParams) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *UpdatableCommentParams) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *UpdatableCommentParams) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### SetVotesUpNil

`func (o *UpdatableCommentParams) SetVotesUpNil(b bool)`

 SetVotesUpNil sets the value for VotesUp to be an explicit nil

### UnsetVotesUp
`func (o *UpdatableCommentParams) UnsetVotesUp()`

UnsetVotesUp ensures that no value is present for VotesUp, not even an explicit nil
### GetVotesDown

`func (o *UpdatableCommentParams) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *UpdatableCommentParams) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *UpdatableCommentParams) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *UpdatableCommentParams) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### SetVotesDownNil

`func (o *UpdatableCommentParams) SetVotesDownNil(b bool)`

 SetVotesDownNil sets the value for VotesDown to be an explicit nil

### UnsetVotesDown
`func (o *UpdatableCommentParams) UnsetVotesDown()`

UnsetVotesDown ensures that no value is present for VotesDown, not even an explicit nil
### GetExpireAt

`func (o *UpdatableCommentParams) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *UpdatableCommentParams) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *UpdatableCommentParams) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *UpdatableCommentParams) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *UpdatableCommentParams) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *UpdatableCommentParams) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetVerified

`func (o *UpdatableCommentParams) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UpdatableCommentParams) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UpdatableCommentParams) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UpdatableCommentParams) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifiedDate

`func (o *UpdatableCommentParams) GetVerifiedDate() time.Time`

GetVerifiedDate returns the VerifiedDate field if non-nil, zero value otherwise.

### GetVerifiedDateOk

`func (o *UpdatableCommentParams) GetVerifiedDateOk() (*time.Time, bool)`

GetVerifiedDateOk returns a tuple with the VerifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDate

`func (o *UpdatableCommentParams) SetVerifiedDate(v time.Time)`

SetVerifiedDate sets VerifiedDate field to given value.

### HasVerifiedDate

`func (o *UpdatableCommentParams) HasVerifiedDate() bool`

HasVerifiedDate returns a boolean if a field has been set.

### SetVerifiedDateNil

`func (o *UpdatableCommentParams) SetVerifiedDateNil(b bool)`

 SetVerifiedDateNil sets the value for VerifiedDate to be an explicit nil

### UnsetVerifiedDate
`func (o *UpdatableCommentParams) UnsetVerifiedDate()`

UnsetVerifiedDate ensures that no value is present for VerifiedDate, not even an explicit nil
### GetNotificationSentForParent

`func (o *UpdatableCommentParams) GetNotificationSentForParent() bool`

GetNotificationSentForParent returns the NotificationSentForParent field if non-nil, zero value otherwise.

### GetNotificationSentForParentOk

`func (o *UpdatableCommentParams) GetNotificationSentForParentOk() (*bool, bool)`

GetNotificationSentForParentOk returns a tuple with the NotificationSentForParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParent

`func (o *UpdatableCommentParams) SetNotificationSentForParent(v bool)`

SetNotificationSentForParent sets NotificationSentForParent field to given value.

### HasNotificationSentForParent

`func (o *UpdatableCommentParams) HasNotificationSentForParent() bool`

HasNotificationSentForParent returns a boolean if a field has been set.

### GetNotificationSentForParentTenant

`func (o *UpdatableCommentParams) GetNotificationSentForParentTenant() bool`

GetNotificationSentForParentTenant returns the NotificationSentForParentTenant field if non-nil, zero value otherwise.

### GetNotificationSentForParentTenantOk

`func (o *UpdatableCommentParams) GetNotificationSentForParentTenantOk() (*bool, bool)`

GetNotificationSentForParentTenantOk returns a tuple with the NotificationSentForParentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParentTenant

`func (o *UpdatableCommentParams) SetNotificationSentForParentTenant(v bool)`

SetNotificationSentForParentTenant sets NotificationSentForParentTenant field to given value.

### HasNotificationSentForParentTenant

`func (o *UpdatableCommentParams) HasNotificationSentForParentTenant() bool`

HasNotificationSentForParentTenant returns a boolean if a field has been set.

### GetReviewed

`func (o *UpdatableCommentParams) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *UpdatableCommentParams) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *UpdatableCommentParams) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *UpdatableCommentParams) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdatableCommentParams) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdatableCommentParams) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdatableCommentParams) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdatableCommentParams) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalParentId

`func (o *UpdatableCommentParams) GetExternalParentId() string`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *UpdatableCommentParams) GetExternalParentIdOk() (*string, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *UpdatableCommentParams) SetExternalParentId(v string)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *UpdatableCommentParams) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### SetExternalParentIdNil

`func (o *UpdatableCommentParams) SetExternalParentIdNil(b bool)`

 SetExternalParentIdNil sets the value for ExternalParentId to be an explicit nil

### UnsetExternalParentId
`func (o *UpdatableCommentParams) UnsetExternalParentId()`

UnsetExternalParentId ensures that no value is present for ExternalParentId, not even an explicit nil
### GetAvatarSrc

`func (o *UpdatableCommentParams) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *UpdatableCommentParams) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *UpdatableCommentParams) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *UpdatableCommentParams) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *UpdatableCommentParams) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *UpdatableCommentParams) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetIsSpam

`func (o *UpdatableCommentParams) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *UpdatableCommentParams) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *UpdatableCommentParams) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *UpdatableCommentParams) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetApproved

`func (o *UpdatableCommentParams) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *UpdatableCommentParams) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *UpdatableCommentParams) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *UpdatableCommentParams) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UpdatableCommentParams) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UpdatableCommentParams) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UpdatableCommentParams) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UpdatableCommentParams) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *UpdatableCommentParams) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *UpdatableCommentParams) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *UpdatableCommentParams) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *UpdatableCommentParams) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *UpdatableCommentParams) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *UpdatableCommentParams) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *UpdatableCommentParams) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *UpdatableCommentParams) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *UpdatableCommentParams) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *UpdatableCommentParams) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *UpdatableCommentParams) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *UpdatableCommentParams) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *UpdatableCommentParams) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *UpdatableCommentParams) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *UpdatableCommentParams) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *UpdatableCommentParams) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### SetIsPinnedNil

`func (o *UpdatableCommentParams) SetIsPinnedNil(b bool)`

 SetIsPinnedNil sets the value for IsPinned to be an explicit nil

### UnsetIsPinned
`func (o *UpdatableCommentParams) UnsetIsPinned()`

UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
### GetIsLocked

`func (o *UpdatableCommentParams) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UpdatableCommentParams) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UpdatableCommentParams) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UpdatableCommentParams) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *UpdatableCommentParams) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *UpdatableCommentParams) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetFlagCount

`func (o *UpdatableCommentParams) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *UpdatableCommentParams) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *UpdatableCommentParams) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *UpdatableCommentParams) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### SetFlagCountNil

`func (o *UpdatableCommentParams) SetFlagCountNil(b bool)`

 SetFlagCountNil sets the value for FlagCount to be an explicit nil

### UnsetFlagCount
`func (o *UpdatableCommentParams) UnsetFlagCount()`

UnsetFlagCount ensures that no value is present for FlagCount, not even an explicit nil
### GetDisplayLabel

`func (o *UpdatableCommentParams) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *UpdatableCommentParams) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *UpdatableCommentParams) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *UpdatableCommentParams) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *UpdatableCommentParams) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *UpdatableCommentParams) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetMeta

`func (o *UpdatableCommentParams) GetMeta() FCommentMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdatableCommentParams) GetMetaOk() (*FCommentMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdatableCommentParams) SetMeta(v FCommentMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdatableCommentParams) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *UpdatableCommentParams) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *UpdatableCommentParams) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetModerationGroupIds

`func (o *UpdatableCommentParams) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *UpdatableCommentParams) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *UpdatableCommentParams) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *UpdatableCommentParams) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### SetModerationGroupIdsNil

`func (o *UpdatableCommentParams) SetModerationGroupIdsNil(b bool)`

 SetModerationGroupIdsNil sets the value for ModerationGroupIds to be an explicit nil

### UnsetModerationGroupIds
`func (o *UpdatableCommentParams) UnsetModerationGroupIds()`

UnsetModerationGroupIds ensures that no value is present for ModerationGroupIds, not even an explicit nil
### GetFeedbackIds

`func (o *UpdatableCommentParams) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *UpdatableCommentParams) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *UpdatableCommentParams) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *UpdatableCommentParams) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


