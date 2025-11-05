# PickFCommentAPICommentFieldsKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** |  | 
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**UrlIdRaw** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**PageTitle** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AnonUserId** | Pointer to **string** |  | [optional] 
**CommenterEmail** | Pointer to **string** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **string** |  | [optional] 
**Comment** | **string** |  | 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**LocalDateString** | Pointer to **string** |  | [optional] 
**LocalDateHours** | Pointer to **int32** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VotesUp** | Pointer to **int32** |  | [optional] 
**VotesDown** | Pointer to **int32** |  | [optional] 
**ExpireAt** | Pointer to **time.Time** |  | [optional] 
**Verified** | **bool** |  | 
**VerifiedDate** | Pointer to **time.Time** |  | [optional] 
**NotificationSentForParent** | Pointer to **bool** |  | [optional] 
**NotificationSentForParentTenant** | Pointer to **bool** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalParentId** | Pointer to **string** |  | [optional] 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**AiDeterminedSpam** | Pointer to **bool** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HasCode** | Pointer to **bool** |  | [optional] 
**Approved** | **bool** |  | 
**Locale** | **string** |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**FlagCount** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**FromProductId** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**PickFCommentAPICommentFieldsKeysMeta**](PickFCommentAPICommentFieldsKeysMeta.md) |  | [optional] 
**Mentions** | Pointer to [**[]CommentUserMentionInfo**](CommentUserMentionInfo.md) |  | [optional] 
**HashTags** | Pointer to [**[]CommentUserHashTagInfo**](CommentUserHashTagInfo.md) |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPickFCommentAPICommentFieldsKeys

`func NewPickFCommentAPICommentFieldsKeys(date time.Time, id string, tenantId string, urlId string, url string, commenterName string, comment string, commentHTML string, verified bool, approved bool, locale string, ) *PickFCommentAPICommentFieldsKeys`

NewPickFCommentAPICommentFieldsKeys instantiates a new PickFCommentAPICommentFieldsKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickFCommentAPICommentFieldsKeysWithDefaults

`func NewPickFCommentAPICommentFieldsKeysWithDefaults() *PickFCommentAPICommentFieldsKeys`

NewPickFCommentAPICommentFieldsKeysWithDefaults instantiates a new PickFCommentAPICommentFieldsKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PickFCommentAPICommentFieldsKeys) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PickFCommentAPICommentFieldsKeys) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PickFCommentAPICommentFieldsKeys) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetId

`func (o *PickFCommentAPICommentFieldsKeys) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PickFCommentAPICommentFieldsKeys) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *PickFCommentAPICommentFieldsKeys) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PickFCommentAPICommentFieldsKeys) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *PickFCommentAPICommentFieldsKeys) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *PickFCommentAPICommentFieldsKeys) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrlIdRaw

`func (o *PickFCommentAPICommentFieldsKeys) GetUrlIdRaw() string`

GetUrlIdRaw returns the UrlIdRaw field if non-nil, zero value otherwise.

### GetUrlIdRawOk

`func (o *PickFCommentAPICommentFieldsKeys) GetUrlIdRawOk() (*string, bool)`

GetUrlIdRawOk returns a tuple with the UrlIdRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdRaw

`func (o *PickFCommentAPICommentFieldsKeys) SetUrlIdRaw(v string)`

SetUrlIdRaw sets UrlIdRaw field to given value.

### HasUrlIdRaw

`func (o *PickFCommentAPICommentFieldsKeys) HasUrlIdRaw() bool`

HasUrlIdRaw returns a boolean if a field has been set.

### GetUrl

`func (o *PickFCommentAPICommentFieldsKeys) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PickFCommentAPICommentFieldsKeys) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PickFCommentAPICommentFieldsKeys) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *PickFCommentAPICommentFieldsKeys) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *PickFCommentAPICommentFieldsKeys) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *PickFCommentAPICommentFieldsKeys) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *PickFCommentAPICommentFieldsKeys) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetUserId

`func (o *PickFCommentAPICommentFieldsKeys) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PickFCommentAPICommentFieldsKeys) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PickFCommentAPICommentFieldsKeys) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAnonUserId

`func (o *PickFCommentAPICommentFieldsKeys) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *PickFCommentAPICommentFieldsKeys) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *PickFCommentAPICommentFieldsKeys) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetCommenterEmail

`func (o *PickFCommentAPICommentFieldsKeys) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *PickFCommentAPICommentFieldsKeys) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *PickFCommentAPICommentFieldsKeys) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *PickFCommentAPICommentFieldsKeys) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### GetCommenterName

`func (o *PickFCommentAPICommentFieldsKeys) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PickFCommentAPICommentFieldsKeys) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PickFCommentAPICommentFieldsKeys) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PickFCommentAPICommentFieldsKeys) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PickFCommentAPICommentFieldsKeys) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PickFCommentAPICommentFieldsKeys) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PickFCommentAPICommentFieldsKeys) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### GetComment

`func (o *PickFCommentAPICommentFieldsKeys) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PickFCommentAPICommentFieldsKeys) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PickFCommentAPICommentFieldsKeys) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *PickFCommentAPICommentFieldsKeys) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PickFCommentAPICommentFieldsKeys) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PickFCommentAPICommentFieldsKeys) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PickFCommentAPICommentFieldsKeys) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PickFCommentAPICommentFieldsKeys) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PickFCommentAPICommentFieldsKeys) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetLocalDateString

`func (o *PickFCommentAPICommentFieldsKeys) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *PickFCommentAPICommentFieldsKeys) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *PickFCommentAPICommentFieldsKeys) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *PickFCommentAPICommentFieldsKeys) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### GetLocalDateHours

`func (o *PickFCommentAPICommentFieldsKeys) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *PickFCommentAPICommentFieldsKeys) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *PickFCommentAPICommentFieldsKeys) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *PickFCommentAPICommentFieldsKeys) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### GetVotes

`func (o *PickFCommentAPICommentFieldsKeys) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PickFCommentAPICommentFieldsKeys) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PickFCommentAPICommentFieldsKeys) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PickFCommentAPICommentFieldsKeys) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesUp

`func (o *PickFCommentAPICommentFieldsKeys) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PickFCommentAPICommentFieldsKeys) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PickFCommentAPICommentFieldsKeys) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PickFCommentAPICommentFieldsKeys) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### GetVotesDown

`func (o *PickFCommentAPICommentFieldsKeys) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PickFCommentAPICommentFieldsKeys) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PickFCommentAPICommentFieldsKeys) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PickFCommentAPICommentFieldsKeys) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetExpireAt

`func (o *PickFCommentAPICommentFieldsKeys) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *PickFCommentAPICommentFieldsKeys) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *PickFCommentAPICommentFieldsKeys) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *PickFCommentAPICommentFieldsKeys) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetVerified

`func (o *PickFCommentAPICommentFieldsKeys) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PickFCommentAPICommentFieldsKeys) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PickFCommentAPICommentFieldsKeys) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedDate

`func (o *PickFCommentAPICommentFieldsKeys) GetVerifiedDate() time.Time`

GetVerifiedDate returns the VerifiedDate field if non-nil, zero value otherwise.

### GetVerifiedDateOk

`func (o *PickFCommentAPICommentFieldsKeys) GetVerifiedDateOk() (*time.Time, bool)`

GetVerifiedDateOk returns a tuple with the VerifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedDate

`func (o *PickFCommentAPICommentFieldsKeys) SetVerifiedDate(v time.Time)`

SetVerifiedDate sets VerifiedDate field to given value.

### HasVerifiedDate

`func (o *PickFCommentAPICommentFieldsKeys) HasVerifiedDate() bool`

HasVerifiedDate returns a boolean if a field has been set.

### GetNotificationSentForParent

`func (o *PickFCommentAPICommentFieldsKeys) GetNotificationSentForParent() bool`

GetNotificationSentForParent returns the NotificationSentForParent field if non-nil, zero value otherwise.

### GetNotificationSentForParentOk

`func (o *PickFCommentAPICommentFieldsKeys) GetNotificationSentForParentOk() (*bool, bool)`

GetNotificationSentForParentOk returns a tuple with the NotificationSentForParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParent

`func (o *PickFCommentAPICommentFieldsKeys) SetNotificationSentForParent(v bool)`

SetNotificationSentForParent sets NotificationSentForParent field to given value.

### HasNotificationSentForParent

`func (o *PickFCommentAPICommentFieldsKeys) HasNotificationSentForParent() bool`

HasNotificationSentForParent returns a boolean if a field has been set.

### GetNotificationSentForParentTenant

`func (o *PickFCommentAPICommentFieldsKeys) GetNotificationSentForParentTenant() bool`

GetNotificationSentForParentTenant returns the NotificationSentForParentTenant field if non-nil, zero value otherwise.

### GetNotificationSentForParentTenantOk

`func (o *PickFCommentAPICommentFieldsKeys) GetNotificationSentForParentTenantOk() (*bool, bool)`

GetNotificationSentForParentTenantOk returns a tuple with the NotificationSentForParentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSentForParentTenant

`func (o *PickFCommentAPICommentFieldsKeys) SetNotificationSentForParentTenant(v bool)`

SetNotificationSentForParentTenant sets NotificationSentForParentTenant field to given value.

### HasNotificationSentForParentTenant

`func (o *PickFCommentAPICommentFieldsKeys) HasNotificationSentForParentTenant() bool`

HasNotificationSentForParentTenant returns a boolean if a field has been set.

### GetReviewed

`func (o *PickFCommentAPICommentFieldsKeys) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *PickFCommentAPICommentFieldsKeys) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *PickFCommentAPICommentFieldsKeys) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *PickFCommentAPICommentFieldsKeys) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetExternalId

`func (o *PickFCommentAPICommentFieldsKeys) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PickFCommentAPICommentFieldsKeys) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PickFCommentAPICommentFieldsKeys) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalParentId

`func (o *PickFCommentAPICommentFieldsKeys) GetExternalParentId() string`

GetExternalParentId returns the ExternalParentId field if non-nil, zero value otherwise.

### GetExternalParentIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetExternalParentIdOk() (*string, bool)`

GetExternalParentIdOk returns a tuple with the ExternalParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParentId

`func (o *PickFCommentAPICommentFieldsKeys) SetExternalParentId(v string)`

SetExternalParentId sets ExternalParentId field to given value.

### HasExternalParentId

`func (o *PickFCommentAPICommentFieldsKeys) HasExternalParentId() bool`

HasExternalParentId returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *PickFCommentAPICommentFieldsKeys) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PickFCommentAPICommentFieldsKeys) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PickFCommentAPICommentFieldsKeys) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PickFCommentAPICommentFieldsKeys) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsSpam

`func (o *PickFCommentAPICommentFieldsKeys) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PickFCommentAPICommentFieldsKeys) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PickFCommentAPICommentFieldsKeys) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetAiDeterminedSpam

`func (o *PickFCommentAPICommentFieldsKeys) GetAiDeterminedSpam() bool`

GetAiDeterminedSpam returns the AiDeterminedSpam field if non-nil, zero value otherwise.

### GetAiDeterminedSpamOk

`func (o *PickFCommentAPICommentFieldsKeys) GetAiDeterminedSpamOk() (*bool, bool)`

GetAiDeterminedSpamOk returns a tuple with the AiDeterminedSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDeterminedSpam

`func (o *PickFCommentAPICommentFieldsKeys) SetAiDeterminedSpam(v bool)`

SetAiDeterminedSpam sets AiDeterminedSpam field to given value.

### HasAiDeterminedSpam

`func (o *PickFCommentAPICommentFieldsKeys) HasAiDeterminedSpam() bool`

HasAiDeterminedSpam returns a boolean if a field has been set.

### GetHasImages

`func (o *PickFCommentAPICommentFieldsKeys) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *PickFCommentAPICommentFieldsKeys) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *PickFCommentAPICommentFieldsKeys) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *PickFCommentAPICommentFieldsKeys) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetHasLinks

`func (o *PickFCommentAPICommentFieldsKeys) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *PickFCommentAPICommentFieldsKeys) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *PickFCommentAPICommentFieldsKeys) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *PickFCommentAPICommentFieldsKeys) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasCode

`func (o *PickFCommentAPICommentFieldsKeys) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *PickFCommentAPICommentFieldsKeys) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *PickFCommentAPICommentFieldsKeys) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *PickFCommentAPICommentFieldsKeys) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetApproved

`func (o *PickFCommentAPICommentFieldsKeys) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *PickFCommentAPICommentFieldsKeys) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *PickFCommentAPICommentFieldsKeys) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetLocale

`func (o *PickFCommentAPICommentFieldsKeys) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PickFCommentAPICommentFieldsKeys) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PickFCommentAPICommentFieldsKeys) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetIsDeleted

`func (o *PickFCommentAPICommentFieldsKeys) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PickFCommentAPICommentFieldsKeys) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PickFCommentAPICommentFieldsKeys) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PickFCommentAPICommentFieldsKeys) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PickFCommentAPICommentFieldsKeys) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PickFCommentAPICommentFieldsKeys) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PickFCommentAPICommentFieldsKeys) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PickFCommentAPICommentFieldsKeys) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PickFCommentAPICommentFieldsKeys) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PickFCommentAPICommentFieldsKeys) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PickFCommentAPICommentFieldsKeys) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PickFCommentAPICommentFieldsKeys) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PickFCommentAPICommentFieldsKeys) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PickFCommentAPICommentFieldsKeys) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PickFCommentAPICommentFieldsKeys) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsLocked

`func (o *PickFCommentAPICommentFieldsKeys) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PickFCommentAPICommentFieldsKeys) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PickFCommentAPICommentFieldsKeys) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PickFCommentAPICommentFieldsKeys) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetFlagCount

`func (o *PickFCommentAPICommentFieldsKeys) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *PickFCommentAPICommentFieldsKeys) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *PickFCommentAPICommentFieldsKeys) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *PickFCommentAPICommentFieldsKeys) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### GetRating

`func (o *PickFCommentAPICommentFieldsKeys) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PickFCommentAPICommentFieldsKeys) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PickFCommentAPICommentFieldsKeys) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PickFCommentAPICommentFieldsKeys) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PickFCommentAPICommentFieldsKeys) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PickFCommentAPICommentFieldsKeys) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PickFCommentAPICommentFieldsKeys) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PickFCommentAPICommentFieldsKeys) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetFromProductId

`func (o *PickFCommentAPICommentFieldsKeys) GetFromProductId() int32`

GetFromProductId returns the FromProductId field if non-nil, zero value otherwise.

### GetFromProductIdOk

`func (o *PickFCommentAPICommentFieldsKeys) GetFromProductIdOk() (*int32, bool)`

GetFromProductIdOk returns a tuple with the FromProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromProductId

`func (o *PickFCommentAPICommentFieldsKeys) SetFromProductId(v int32)`

SetFromProductId sets FromProductId field to given value.

### HasFromProductId

`func (o *PickFCommentAPICommentFieldsKeys) HasFromProductId() bool`

HasFromProductId returns a boolean if a field has been set.

### GetMeta

`func (o *PickFCommentAPICommentFieldsKeys) GetMeta() PickFCommentAPICommentFieldsKeysMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PickFCommentAPICommentFieldsKeys) GetMetaOk() (*PickFCommentAPICommentFieldsKeysMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PickFCommentAPICommentFieldsKeys) SetMeta(v PickFCommentAPICommentFieldsKeysMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PickFCommentAPICommentFieldsKeys) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMentions

`func (o *PickFCommentAPICommentFieldsKeys) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *PickFCommentAPICommentFieldsKeys) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *PickFCommentAPICommentFieldsKeys) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *PickFCommentAPICommentFieldsKeys) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetHashTags

`func (o *PickFCommentAPICommentFieldsKeys) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *PickFCommentAPICommentFieldsKeys) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *PickFCommentAPICommentFieldsKeys) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *PickFCommentAPICommentFieldsKeys) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetBadges

`func (o *PickFCommentAPICommentFieldsKeys) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PickFCommentAPICommentFieldsKeys) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PickFCommentAPICommentFieldsKeys) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PickFCommentAPICommentFieldsKeys) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetDomain

`func (o *PickFCommentAPICommentFieldsKeys) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PickFCommentAPICommentFieldsKeys) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PickFCommentAPICommentFieldsKeys) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PickFCommentAPICommentFieldsKeys) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *PickFCommentAPICommentFieldsKeys) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *PickFCommentAPICommentFieldsKeys) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *PickFCommentAPICommentFieldsKeys) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *PickFCommentAPICommentFieldsKeys) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *PickFCommentAPICommentFieldsKeys) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PickFCommentAPICommentFieldsKeys) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PickFCommentAPICommentFieldsKeys) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PickFCommentAPICommentFieldsKeys) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


