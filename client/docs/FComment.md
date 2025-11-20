# FComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**UrlIdRaw** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**PageTitle** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**AnonUserId** | Pointer to **NullableString** |  | [optional] 
**CommenterEmail** | Pointer to **NullableString** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **NullableString** |  | [optional] 
**Comment** | **string** |  | 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Date** | **NullableTime** |  | 
**LocalDateString** | Pointer to **NullableString** |  | [optional] 
**LocalDateHours** | Pointer to **NullableInt32** |  | [optional] 
**Votes** | Pointer to **NullableInt32** |  | [optional] 
**VotesUp** | Pointer to **NullableInt32** |  | [optional] 
**VotesDown** | Pointer to **NullableInt32** |  | [optional] 
**ExpireAt** | Pointer to **NullableTime** |  | [optional] 
**Verified** | **bool** |  | 
**VerifiedDate** | Pointer to **NullableTime** |  | [optional] 
**VerificationId** | Pointer to **NullableString** |  | [optional] 
**NotificationSentForParent** | Pointer to **bool** |  | [optional] 
**NotificationSentForParentTenant** | Pointer to **bool** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**Imported** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalParentId** | Pointer to **NullableString** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**PermNotSpam** | Pointer to **bool** |  | [optional] 
**AiDeterminedSpam** | Pointer to **bool** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**PageNumber** | Pointer to **NullableInt32** |  | [optional] 
**PageNumberOF** | Pointer to **NullableInt32** |  | [optional] 
**PageNumberNF** | Pointer to **NullableInt32** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HasCode** | Pointer to **bool** |  | [optional] 
**Approved** | **bool** |  | 
**Locale** | **NullableString** |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsBannedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **NullableBool** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**FlagCount** | Pointer to **NullableInt32** |  | [optional] 
**Rating** | Pointer to **NullableFloat64** |  | [optional] 
**DisplayLabel** | Pointer to **NullableString** |  | [optional] 
**FromProductId** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**NullableFCommentMeta**](FCommentMeta.md) |  | [optional] 
**IpHash** | Pointer to **string** |  | [optional] 
**Mentions** | Pointer to [**[]CommentUserMentionInfo**](CommentUserMentionInfo.md) |  | [optional] 
**HashTags** | Pointer to [**[]CommentUserHashTagInfo**](CommentUserHashTagInfo.md) |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**VeteranBadgeProcessed** | Pointer to **string** |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**DidProcessBadges** | Pointer to **bool** |  | [optional] 
**FromOfflineRestore** | Pointer to **bool** |  | [optional] 
**AutoplayJobId** | Pointer to **string** |  | [optional] 
**AutoplayDelayMS** | Pointer to **int64** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**Logs** | Pointer to [**[]CommentLogEntry**](CommentLogEntry.md) |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**ViewCount** | Pointer to **NullableInt64** |  | [optional] 
**RequiresVerification** | Pointer to **bool** |  | [optional] 
**EditKey** | Pointer to **string** |  | [optional] 

## Methods

### NewFComment

`func NewFComment(id string, tenantId string, urlId string, url string, commenterName string, comment string, commentHTML string, date NullableTime, verified bool, approved bool, locale NullableString, ) *FComment`

NewFComment instantiates a new FComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCommentWithDefaults

`func NewFCommentWithDefaults() *FComment`

NewFCommentWithDefaults instantiates a new FComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FComment) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *FComment) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FComment) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FComment) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *FComment) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *FComment) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *FComment) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrlIdRaw

`func (o *FComment) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *FComment) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *FComment) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.

### HasUrlIdRaw

`func (o *FComment) HasUrlIdRaw() bool`

HasUrlIdRaw returns a boolean if a field has been set.

### GetUrl

`func (o *FComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *FComment) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *FComment) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *FComment) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *FComment) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *FComment) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *FComment) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetUserId

`func (o *FComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *FComment) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *FComment) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAnonUserId

`func (o *FComment) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *FComment) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *FComment) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *FComment) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *FComment) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *FComment) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetCommenterEmail

`func (o *FComment) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *FComment) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *FComment) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *FComment) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *FComment) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *FComment) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetCommenterName

`func (o *FComment) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *FComment) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *FComment) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *FComment) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *FComment) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *FComment) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *FComment) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *FComment) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *FComment) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetComment

`func (o *FComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FComment) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *FComment) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *FComment) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *FComment) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *FComment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FComment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FComment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *FComment) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *FComment) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetDate

`func (o *FComment) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FComment) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FComment) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *FComment) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *FComment) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetLocalDateString

`func (o *FComment) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *FComment) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *FComment) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *FComment) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### SetLocalDateStringNil

`func (o *FComment) SetLocalDateStringNil(b bool)`

 SetLocalDateStringNil sets the value for LocalDateString to be an explicit nil

### UnsetLocalDateString
`func (o *FComment) UnsetLocalDateString()`

UnsetLocalDateString ensures that no value is present for LocalDateString, not even an explicit nil
### GetLocalDateHours

`func (o *FComment) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *FComment) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *FComment) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *FComment) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### SetLocalDateHoursNil

`func (o *FComment) SetLocalDateHoursNil(b bool)`

 SetLocalDateHoursNil sets the value for LocalDateHours to be an explicit nil

### UnsetLocalDateHours
`func (o *FComment) UnsetLocalDateHours()`

UnsetLocalDateHours ensures that no value is present for LocalDateHours, not even an explicit nil
### GetVotes

`func (o *FComment) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *FComment) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *FComment) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *FComment) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### SetVotesNil

`func (o *FComment) SetVotesNil(b bool)`

 SetVotesNil sets the value for Votes to be an explicit nil

### UnsetVotes
`func (o *FComment) UnsetVotes()`

UnsetVotes ensures that no value is present for Votes, not even an explicit nil
### GetVotesUp

`func (o *FComment) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *FComment) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *FComment) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *FComment) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### SetVotesUpNil

`func (o *FComment) SetVotesUpNil(b bool)`

 SetVotesUpNil sets the value for VotesUp to be an explicit nil

### UnsetVotesUp
`func (o *FComment) UnsetVotesUp()`

UnsetVotesUp ensures that no value is present for VotesUp, not even an explicit nil
### GetVotesDown

`func (o *FComment) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *FComment) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *FComment) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *FComment) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### SetVotesDownNil

`func (o *FComment) SetVotesDownNil(b bool)`

 SetVotesDownNil sets the value for VotesDown to be an explicit nil

### UnsetVotesDown
`func (o *FComment) UnsetVotesDown()`

UnsetVotesDown ensures that no value is present for VotesDown, not even an explicit nil
### GetExpireAt

`func (o *FComment) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *FComment) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *FComment) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *FComment) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *FComment) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *FComment) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetVerified

`func (o *FComment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *FComment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *FComment) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedDate

`func (o *FComment) GetVerifiedDate() time.Time`

GetVerifiedDate returns the VerifiedDate field if non-nil, zero value otherwise.

### GetVerifiedDateOk

`func (o *FComment) GetVerifiedDateOk() (*time.Time, bool)`

GetVerifiedDateOk returns a tuple with the VerifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDate

`func (o *FComment) SetVerifiedDate(v time.Time)`

SetVerifiedDate sets VerifiedDate field to given value.

### HasVerifiedDate

`func (o *FComment) HasVerifiedDate() bool`

HasVerifiedDate returns a boolean if a field has been set.

### SetVerifiedDateNil

`func (o *FComment) SetVerifiedDateNil(b bool)`

 SetVerifiedDateNil sets the value for VerifiedDate to be an explicit nil

### UnsetVerifiedDate
`func (o *FComment) UnsetVerifiedDate()`

UnsetVerifiedDate ensures that no value is present for VerifiedDate, not even an explicit nil
### GetVerificationId

`func (o *FComment) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *FComment) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *FComment) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *FComment) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.

### SetVerificationIdNil

`func (o *FComment) SetVerificationIdNil(b bool)`

 SetVerificationIdNil sets the value for VerificationId to be an explicit nil

### UnsetVerificationId
`func (o *FComment) UnsetVerificationId()`

UnsetVerificationId ensures that no value is present for VerificationId, not even an explicit nil
### GetNotificationSentForParent

`func (o *FComment) GetNotificationSentForParent() bool`

GetNotificationSentForParent returns the NotificationSentForParent field if non-nil, zero value otherwise.

### GetNotificationSentForParentOk

`func (o *FComment) GetNotificationSentForParentOk() (*bool, bool)`

GetNotificationSentForParentOk returns a tuple with the NotificationSentForParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParent

`func (o *FComment) SetNotificationSentForParent(v bool)`

SetNotificationSentForParent sets NotificationSentForParent field to given value.

### HasNotificationSentForParent

`func (o *FComment) HasNotificationSentForParent() bool`

HasNotificationSentForParent returns a boolean if a field has been set.

### GetNotificationSentForParentTenant

`func (o *FComment) GetNotificationSentForParentTenant() bool`

GetNotificationSentForParentTenant returns the NotificationSentForParentTenant field if non-nil, zero value otherwise.

### GetNotificationSentForParentTenantOk

`func (o *FComment) GetNotificationSentForParentTenantOk() (*bool, bool)`

GetNotificationSentForParentTenantOk returns a tuple with the NotificationSentForParentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParentTenant

`func (o *FComment) SetNotificationSentForParentTenant(v bool)`

SetNotificationSentForParentTenant sets NotificationSentForParentTenant field to given value.

### HasNotificationSentForParentTenant

`func (o *FComment) HasNotificationSentForParentTenant() bool`

HasNotificationSentForParentTenant returns a boolean if a field has been set.

### GetReviewed

`func (o *FComment) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *FComment) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *FComment) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *FComment) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetImported

`func (o *FComment) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *FComment) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *FComment) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *FComment) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetExternalId

`func (o *FComment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FComment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FComment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FComment) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalParentId

`func (o *FComment) GetExternalParentId() string`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *FComment) GetExternalParentIdOk() (*string, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *FComment) SetExternalParentId(v string)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *FComment) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### SetExternalParentIdNil

`func (o *FComment) SetExternalParentIdNil(b bool)`

 SetExternalParentIdNil sets the value for ExternalParentId to be an explicit nil

### UnsetExternalParentId
`func (o *FComment) UnsetExternalParentId()`

UnsetExternalParentId ensures that no value is present for ExternalParentId, not even an explicit nil
### GetAvatarSrc

`func (o *FComment) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *FComment) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *FComment) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *FComment) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *FComment) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *FComment) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetIsSpam

`func (o *FComment) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *FComment) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *FComment) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *FComment) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetPermNotSpam

`func (o *FComment) GetPermNotSpam() bool`

GetPermNotSpam returns the PermNotSpam field if non-nil, zero value otherwise.

### GetPermNotSpamOk

`func (o *FComment) GetPermNotSpamOk() (*bool, bool)`

GetPermNotSpamOk returns a tuple with the PermNotSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermNotSpam

`func (o *FComment) SetPermNotSpam(v bool)`

SetPermNotSpam sets PermNotSpam field to given value.

### HasPermNotSpam

`func (o *FComment) HasPermNotSpam() bool`

HasPermNotSpam returns a boolean if a field has been set.

### GetAiDeterminedSpam

`func (o *FComment) GetAiDeterminedSpam() bool`

GetAiDeterminedSpam returns the AiDeterminedSpam field if non-nil, zero value otherwise.

### GetAiDeterminedSpamOk

`func (o *FComment) GetAiDeterminedSpamOk() (*bool, bool)`

GetAiDeterminedSpamOk returns a tuple with the AiDeterminedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDeterminedSpam

`func (o *FComment) SetAiDeterminedSpam(v bool)`

SetAiDeterminedSpam sets AiDeterminedSpam field to given value.

### HasAiDeterminedSpam

`func (o *FComment) HasAiDeterminedSpam() bool`

HasAiDeterminedSpam returns a boolean if a field has been set.

### GetHasImages

`func (o *FComment) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *FComment) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *FComment) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *FComment) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetPageNumber

`func (o *FComment) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *FComment) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *FComment) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *FComment) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### SetPageNumberNil

`func (o *FComment) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *FComment) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageNumberOF

`func (o *FComment) GetPageNumberOF() int32`

GetPageNumberOF returns the PageNumberOF field if non-nil, zero value otherwise.

### GetPageNumberOFOk

`func (o *FComment) GetPageNumberOFOk() (*int32, bool)`

GetPageNumberOFOk returns a tuple with the PageNumberOF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumberOF

`func (o *FComment) SetPageNumberOF(v int32)`

SetPageNumberOF sets PageNumberOF field to given value.

### HasPageNumberOF

`func (o *FComment) HasPageNumberOF() bool`

HasPageNumberOF returns a boolean if a field has been set.

### SetPageNumberOFNil

`func (o *FComment) SetPageNumberOFNil(b bool)`

 SetPageNumberOFNil sets the value for PageNumberOF to be an explicit nil

### UnsetPageNumberOF
`func (o *FComment) UnsetPageNumberOF()`

UnsetPageNumberOF ensures that no value is present for PageNumberOF, not even an explicit nil
### GetPageNumberNF

`func (o *FComment) GetPageNumberNF() int32`

GetPageNumberNF returns the PageNumberNF field if non-nil, zero value otherwise.

### GetPageNumberNFOk

`func (o *FComment) GetPageNumberNFOk() (*int32, bool)`

GetPageNumberNFOk returns a tuple with the PageNumberNF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumberNF

`func (o *FComment) SetPageNumberNF(v int32)`

SetPageNumberNF sets PageNumberNF field to given value.

### HasPageNumberNF

`func (o *FComment) HasPageNumberNF() bool`

HasPageNumberNF returns a boolean if a field has been set.

### SetPageNumberNFNil

`func (o *FComment) SetPageNumberNFNil(b bool)`

 SetPageNumberNFNil sets the value for PageNumberNF to be an explicit nil

### UnsetPageNumberNF
`func (o *FComment) UnsetPageNumberNF()`

UnsetPageNumberNF ensures that no value is present for PageNumberNF, not even an explicit nil
### GetHasLinks

`func (o *FComment) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *FComment) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *FComment) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *FComment) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasCode

`func (o *FComment) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *FComment) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *FComment) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *FComment) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetApproved

`func (o *FComment) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *FComment) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *FComment) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetLocale

`func (o *FComment) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *FComment) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *FComment) SetLocale(v string)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *FComment) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *FComment) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetIsDeleted

`func (o *FComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *FComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *FComment) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *FComment) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *FComment) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *FComment) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsBannedUser

`func (o *FComment) GetIsBannedUser() bool`

GetIsBannedUser returns the IsBannedUser field if non-nil, zero value otherwise.

### GetIsBannedUserOk

`func (o *FComment) GetIsBannedUserOk() (*bool, bool)`

GetIsBannedUserOk returns a tuple with the IsBannedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBannedUser

`func (o *FComment) SetIsBannedUser(v bool)`

SetIsBannedUser sets IsBannedUser field to given value.

### HasIsBannedUser

`func (o *FComment) HasIsBannedUser() bool`

HasIsBannedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *FComment) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *FComment) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *FComment) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *FComment) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *FComment) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *FComment) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *FComment) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *FComment) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *FComment) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *FComment) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *FComment) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *FComment) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### SetIsPinnedNil

`func (o *FComment) SetIsPinnedNil(b bool)`

 SetIsPinnedNil sets the value for IsPinned to be an explicit nil

### UnsetIsPinned
`func (o *FComment) UnsetIsPinned()`

UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
### GetIsLocked

`func (o *FComment) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *FComment) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *FComment) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *FComment) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *FComment) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *FComment) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetFlagCount

`func (o *FComment) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *FComment) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *FComment) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *FComment) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### SetFlagCountNil

`func (o *FComment) SetFlagCountNil(b bool)`

 SetFlagCountNil sets the value for FlagCount to be an explicit nil

### UnsetFlagCount
`func (o *FComment) UnsetFlagCount()`

UnsetFlagCount ensures that no value is present for FlagCount, not even an explicit nil
### GetRating

`func (o *FComment) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *FComment) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *FComment) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *FComment) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *FComment) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *FComment) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetDisplayLabel

`func (o *FComment) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *FComment) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *FComment) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *FComment) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *FComment) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *FComment) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetFromProductId

`func (o *FComment) GetFromProductId() int32`

GetFromProductId returns the FromProductId field if non-nil, zero value otherwise.

### GetFromProductIdOk

`func (o *FComment) GetFromProductIdOk() (*int32, bool)`

GetFromProductIdOk returns a tuple with the FromProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromProductId

`func (o *FComment) SetFromProductId(v int32)`

SetFromProductId sets FromProductId field to given value.

### HasFromProductId

`func (o *FComment) HasFromProductId() bool`

HasFromProductId returns a boolean if a field has been set.

### GetMeta

`func (o *FComment) GetMeta() FCommentMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FComment) GetMetaOk() (*FCommentMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FComment) SetMeta(v FCommentMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FComment) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *FComment) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *FComment) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetIpHash

`func (o *FComment) GetIpHash() string`

GetIpHash returns the IpHash field if non-nil, zero value otherwise.

### GetIpHashOk

`func (o *FComment) GetIpHashOk() (*string, bool)`

GetIpHashOk returns a tuple with the IpHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHash

`func (o *FComment) SetIpHash(v string)`

SetIpHash sets IpHash field to given value.

### HasIpHash

`func (o *FComment) HasIpHash() bool`

HasIpHash returns a boolean if a field has been set.

### GetMentions

`func (o *FComment) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *FComment) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *FComment) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *FComment) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetHashTags

`func (o *FComment) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *FComment) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *FComment) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *FComment) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetBadges

`func (o *FComment) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *FComment) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *FComment) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *FComment) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### SetBadgesNil

`func (o *FComment) SetBadgesNil(b bool)`

 SetBadgesNil sets the value for Badges to be an explicit nil

### UnsetBadges
`func (o *FComment) UnsetBadges()`

UnsetBadges ensures that no value is present for Badges, not even an explicit nil
### GetDomain

`func (o *FComment) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *FComment) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *FComment) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *FComment) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *FComment) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *FComment) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetVeteranBadgeProcessed

`func (o *FComment) GetVeteranBadgeProcessed() string`

GetVeteranBadgeProcessed returns the VeteranBadgeProcessed field if non-nil, zero value otherwise.

### GetVeteranBadgeProcessedOk

`func (o *FComment) GetVeteranBadgeProcessedOk() (*string, bool)`

GetVeteranBadgeProcessedOk returns a tuple with the VeteranBadgeProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVeteranBadgeProcessed

`func (o *FComment) SetVeteranBadgeProcessed(v string)`

SetVeteranBadgeProcessed sets VeteranBadgeProcessed field to given value.

### HasVeteranBadgeProcessed

`func (o *FComment) HasVeteranBadgeProcessed() bool`

HasVeteranBadgeProcessed returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *FComment) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *FComment) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *FComment) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *FComment) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### SetModerationGroupIdsNil

`func (o *FComment) SetModerationGroupIdsNil(b bool)`

 SetModerationGroupIdsNil sets the value for ModerationGroupIds to be an explicit nil

### UnsetModerationGroupIds
`func (o *FComment) UnsetModerationGroupIds()`

UnsetModerationGroupIds ensures that no value is present for ModerationGroupIds, not even an explicit nil
### GetDidProcessBadges

`func (o *FComment) GetDidProcessBadges() bool`

GetDidProcessBadges returns the DidProcessBadges field if non-nil, zero value otherwise.

### GetDidProcessBadgesOk

`func (o *FComment) GetDidProcessBadgesOk() (*bool, bool)`

GetDidProcessBadgesOk returns a tuple with the DidProcessBadges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidProcessBadges

`func (o *FComment) SetDidProcessBadges(v bool)`

SetDidProcessBadges sets DidProcessBadges field to given value.

### HasDidProcessBadges

`func (o *FComment) HasDidProcessBadges() bool`

HasDidProcessBadges returns a boolean if a field has been set.

### GetFromOfflineRestore

`func (o *FComment) GetFromOfflineRestore() bool`

GetFromOfflineRestore returns the FromOfflineRestore field if non-nil, zero value otherwise.

### GetFromOfflineRestoreOk

`func (o *FComment) GetFromOfflineRestoreOk() (*bool, bool)`

GetFromOfflineRestoreOk returns a tuple with the FromOfflineRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOfflineRestore

`func (o *FComment) SetFromOfflineRestore(v bool)`

SetFromOfflineRestore sets FromOfflineRestore field to given value.

### HasFromOfflineRestore

`func (o *FComment) HasFromOfflineRestore() bool`

HasFromOfflineRestore returns a boolean if a field has been set.

### GetAutoplayJobId

`func (o *FComment) GetAutoplayJobId() string`

GetAutoplayJobId returns the AutoplayJobId field if non-nil, zero value otherwise.

### GetAutoplayJobIdOk

`func (o *FComment) GetAutoplayJobIdOk() (*string, bool)`

GetAutoplayJobIdOk returns a tuple with the AutoplayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoplayJobId

`func (o *FComment) SetAutoplayJobId(v string)`

SetAutoplayJobId sets AutoplayJobId field to given value.

### HasAutoplayJobId

`func (o *FComment) HasAutoplayJobId() bool`

HasAutoplayJobId returns a boolean if a field has been set.

### GetAutoplayDelayMS

`func (o *FComment) GetAutoplayDelayMS() int64`

GetAutoplayDelayMS returns the AutoplayDelayMS field if non-nil, zero value otherwise.

### GetAutoplayDelayMSOk

`func (o *FComment) GetAutoplayDelayMSOk() (*int64, bool)`

GetAutoplayDelayMSOk returns a tuple with the AutoplayDelayMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoplayDelayMS

`func (o *FComment) SetAutoplayDelayMS(v int64)`

SetAutoplayDelayMS sets AutoplayDelayMS field to given value.

### HasAutoplayDelayMS

`func (o *FComment) HasAutoplayDelayMS() bool`

HasAutoplayDelayMS returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *FComment) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *FComment) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *FComment) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *FComment) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetLogs

`func (o *FComment) GetLogs() []CommentLogEntry`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *FComment) GetLogsOk() (*[]CommentLogEntry, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *FComment) SetLogs(v []CommentLogEntry)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *FComment) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *FComment) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *FComment) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetGroupIds

`func (o *FComment) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *FComment) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *FComment) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *FComment) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### SetGroupIdsNil

`func (o *FComment) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *FComment) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil
### GetViewCount

`func (o *FComment) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *FComment) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *FComment) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *FComment) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### SetViewCountNil

`func (o *FComment) SetViewCountNil(b bool)`

 SetViewCountNil sets the value for ViewCount to be an explicit nil

### UnsetViewCount
`func (o *FComment) UnsetViewCount()`

UnsetViewCount ensures that no value is present for ViewCount, not even an explicit nil
### GetRequiresVerification

`func (o *FComment) GetRequiresVerification() bool`

GetRequiresVerification returns the RequiresVerification field if non-nil, zero value otherwise.

### GetRequiresVerificationOk

`func (o *FComment) GetRequiresVerificationOk() (*bool, bool)`

GetRequiresVerificationOk returns a tuple with the RequiresVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVerification

`func (o *FComment) SetRequiresVerification(v bool)`

SetRequiresVerification sets RequiresVerification field to given value.

### HasRequiresVerification

`func (o *FComment) HasRequiresVerification() bool`

HasRequiresVerification returns a boolean if a field has been set.

### GetEditKey

`func (o *FComment) GetEditKey() string`

GetEditKey returns the EditKey field if non-nil, zero value otherwise.

### GetEditKeyOk

`func (o *FComment) GetEditKeyOk() (*string, bool)`

GetEditKeyOk returns a tuple with the EditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditKey

`func (o *FComment) SetEditKey(v string)`

SetEditKey sets EditKey field to given value.

### HasEditKey

`func (o *FComment) HasEditKey() bool`

HasEditKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


