# ModerationAPIComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLocalDeleted** | Pointer to **bool** |  | [optional] 
**ReplyCount** | Pointer to **float64** |  | [optional] 
**FeedbackResults** | Pointer to **[]string** |  | [optional] 
**IsVotedUp** | Pointer to **bool** |  | [optional] 
**IsVotedDown** | Pointer to **bool** |  | [optional] 
**MyVoteId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**Url** | **string** |  | 
**PageTitle** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**AnonUserId** | Pointer to **NullableString** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **NullableString** |  | [optional] 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Date** | **NullableTime** |  | 
**LocalDateString** | Pointer to **NullableString** |  | [optional] 
**Votes** | Pointer to **NullableFloat64** |  | [optional] 
**VotesUp** | Pointer to **NullableFloat64** |  | [optional] 
**VotesDown** | Pointer to **NullableFloat64** |  | [optional] 
**ExpireAt** | Pointer to **NullableTime** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**PermNotSpam** | Pointer to **bool** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HasCode** | Pointer to **bool** |  | [optional] 
**Approved** | **bool** |  | 
**Locale** | **NullableString** |  | 
**IsBannedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **NullableBool** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**FlagCount** | Pointer to **NullableFloat64** |  | [optional] 
**DisplayLabel** | Pointer to **NullableString** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Verified** | **bool** |  | 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewModerationAPIComment

`func NewModerationAPIComment(id string, tenantId string, urlId string, url string, commenterName string, commentHTML string, date NullableTime, approved bool, locale NullableString, verified bool, ) *ModerationAPIComment`

NewModerationAPIComment instantiates a new ModerationAPIComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationAPICommentWithDefaults

`func NewModerationAPICommentWithDefaults() *ModerationAPIComment`

NewModerationAPICommentWithDefaults instantiates a new ModerationAPIComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocalDeleted

`func (o *ModerationAPIComment) GetIsLocalDeleted() bool`

GetIsLocalDeleted returns the IsLocalDeleted field if non-nil, zero value otherwise.

### GetIsLocalDeletedOk

`func (o *ModerationAPIComment) GetIsLocalDeletedOk() (*bool, bool)`

GetIsLocalDeletedOk returns a tuple with the IsLocalDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalDeleted

`func (o *ModerationAPIComment) SetIsLocalDeleted(v bool)`

SetIsLocalDeleted sets IsLocalDeleted field to given value.

### HasIsLocalDeleted

`func (o *ModerationAPIComment) HasIsLocalDeleted() bool`

HasIsLocalDeleted returns a boolean if a field has been set.

### GetReplyCount

`func (o *ModerationAPIComment) GetReplyCount() float64`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *ModerationAPIComment) GetReplyCountOk() (*float64, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *ModerationAPIComment) SetReplyCount(v float64)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *ModerationAPIComment) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetFeedbackResults

`func (o *ModerationAPIComment) GetFeedbackResults() []string`

GetFeedbackResults returns the FeedbackResults field if non-nil, zero value otherwise.

### GetFeedbackResultsOk

`func (o *ModerationAPIComment) GetFeedbackResultsOk() (*[]string, bool)`

GetFeedbackResultsOk returns a tuple with the FeedbackResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackResults

`func (o *ModerationAPIComment) SetFeedbackResults(v []string)`

SetFeedbackResults sets FeedbackResults field to given value.

### HasFeedbackResults

`func (o *ModerationAPIComment) HasFeedbackResults() bool`

HasFeedbackResults returns a boolean if a field has been set.

### GetIsVotedUp

`func (o *ModerationAPIComment) GetIsVotedUp() bool`

GetIsVotedUp returns the IsVotedUp field if non-nil, zero value otherwise.

### GetIsVotedUpOk

`func (o *ModerationAPIComment) GetIsVotedUpOk() (*bool, bool)`

GetIsVotedUpOk returns a tuple with the IsVotedUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVotedUp

`func (o *ModerationAPIComment) SetIsVotedUp(v bool)`

SetIsVotedUp sets IsVotedUp field to given value.

### HasIsVotedUp

`func (o *ModerationAPIComment) HasIsVotedUp() bool`

HasIsVotedUp returns a boolean if a field has been set.

### GetIsVotedDown

`func (o *ModerationAPIComment) GetIsVotedDown() bool`

GetIsVotedDown returns the IsVotedDown field if non-nil, zero value otherwise.

### GetIsVotedDownOk

`func (o *ModerationAPIComment) GetIsVotedDownOk() (*bool, bool)`

GetIsVotedDownOk returns a tuple with the IsVotedDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVotedDown

`func (o *ModerationAPIComment) SetIsVotedDown(v bool)`

SetIsVotedDown sets IsVotedDown field to given value.

### HasIsVotedDown

`func (o *ModerationAPIComment) HasIsVotedDown() bool`

HasIsVotedDown returns a boolean if a field has been set.

### GetMyVoteId

`func (o *ModerationAPIComment) GetMyVoteId() string`

GetMyVoteId returns the MyVoteId field if non-nil, zero value otherwise.

### GetMyVoteIdOk

`func (o *ModerationAPIComment) GetMyVoteIdOk() (*string, bool)`

GetMyVoteIdOk returns a tuple with the MyVoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyVoteId

`func (o *ModerationAPIComment) SetMyVoteId(v string)`

SetMyVoteId sets MyVoteId field to given value.

### HasMyVoteId

`func (o *ModerationAPIComment) HasMyVoteId() bool`

HasMyVoteId returns a boolean if a field has been set.

### GetId

`func (o *ModerationAPIComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModerationAPIComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModerationAPIComment) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ModerationAPIComment) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ModerationAPIComment) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ModerationAPIComment) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *ModerationAPIComment) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *ModerationAPIComment) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *ModerationAPIComment) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrl

`func (o *ModerationAPIComment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModerationAPIComment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModerationAPIComment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *ModerationAPIComment) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *ModerationAPIComment) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *ModerationAPIComment) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *ModerationAPIComment) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### SetPageTitleNil

`func (o *ModerationAPIComment) SetPageTitleNil(b bool)`

 SetPageTitleNil sets the value for PageTitle to be an explicit nil

### UnsetPageTitle
`func (o *ModerationAPIComment) UnsetPageTitle()`

UnsetPageTitle ensures that no value is present for PageTitle, not even an explicit nil
### GetUserId

`func (o *ModerationAPIComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModerationAPIComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModerationAPIComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModerationAPIComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ModerationAPIComment) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ModerationAPIComment) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAnonUserId

`func (o *ModerationAPIComment) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *ModerationAPIComment) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *ModerationAPIComment) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *ModerationAPIComment) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### SetAnonUserIdNil

`func (o *ModerationAPIComment) SetAnonUserIdNil(b bool)`

 SetAnonUserIdNil sets the value for AnonUserId to be an explicit nil

### UnsetAnonUserId
`func (o *ModerationAPIComment) UnsetAnonUserId()`

UnsetAnonUserId ensures that no value is present for AnonUserId, not even an explicit nil
### GetCommenterName

`func (o *ModerationAPIComment) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *ModerationAPIComment) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *ModerationAPIComment) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *ModerationAPIComment) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *ModerationAPIComment) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *ModerationAPIComment) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *ModerationAPIComment) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *ModerationAPIComment) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *ModerationAPIComment) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetCommentHTML

`func (o *ModerationAPIComment) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *ModerationAPIComment) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *ModerationAPIComment) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *ModerationAPIComment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModerationAPIComment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModerationAPIComment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModerationAPIComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ModerationAPIComment) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ModerationAPIComment) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetDate

`func (o *ModerationAPIComment) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ModerationAPIComment) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ModerationAPIComment) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *ModerationAPIComment) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ModerationAPIComment) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetLocalDateString

`func (o *ModerationAPIComment) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *ModerationAPIComment) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *ModerationAPIComment) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *ModerationAPIComment) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### SetLocalDateStringNil

`func (o *ModerationAPIComment) SetLocalDateStringNil(b bool)`

 SetLocalDateStringNil sets the value for LocalDateString to be an explicit nil

### UnsetLocalDateString
`func (o *ModerationAPIComment) UnsetLocalDateString()`

UnsetLocalDateString ensures that no value is present for LocalDateString, not even an explicit nil
### GetVotes

`func (o *ModerationAPIComment) GetVotes() float64`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *ModerationAPIComment) GetVotesOk() (*float64, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *ModerationAPIComment) SetVotes(v float64)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *ModerationAPIComment) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### SetVotesNil

`func (o *ModerationAPIComment) SetVotesNil(b bool)`

 SetVotesNil sets the value for Votes to be an explicit nil

### UnsetVotes
`func (o *ModerationAPIComment) UnsetVotes()`

UnsetVotes ensures that no value is present for Votes, not even an explicit nil
### GetVotesUp

`func (o *ModerationAPIComment) GetVotesUp() float64`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *ModerationAPIComment) GetVotesUpOk() (*float64, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *ModerationAPIComment) SetVotesUp(v float64)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *ModerationAPIComment) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### SetVotesUpNil

`func (o *ModerationAPIComment) SetVotesUpNil(b bool)`

 SetVotesUpNil sets the value for VotesUp to be an explicit nil

### UnsetVotesUp
`func (o *ModerationAPIComment) UnsetVotesUp()`

UnsetVotesUp ensures that no value is present for VotesUp, not even an explicit nil
### GetVotesDown

`func (o *ModerationAPIComment) GetVotesDown() float64`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *ModerationAPIComment) GetVotesDownOk() (*float64, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *ModerationAPIComment) SetVotesDown(v float64)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *ModerationAPIComment) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### SetVotesDownNil

`func (o *ModerationAPIComment) SetVotesDownNil(b bool)`

 SetVotesDownNil sets the value for VotesDown to be an explicit nil

### UnsetVotesDown
`func (o *ModerationAPIComment) UnsetVotesDown()`

UnsetVotesDown ensures that no value is present for VotesDown, not even an explicit nil
### GetExpireAt

`func (o *ModerationAPIComment) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *ModerationAPIComment) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *ModerationAPIComment) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *ModerationAPIComment) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### SetExpireAtNil

`func (o *ModerationAPIComment) SetExpireAtNil(b bool)`

 SetExpireAtNil sets the value for ExpireAt to be an explicit nil

### UnsetExpireAt
`func (o *ModerationAPIComment) UnsetExpireAt()`

UnsetExpireAt ensures that no value is present for ExpireAt, not even an explicit nil
### GetReviewed

`func (o *ModerationAPIComment) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *ModerationAPIComment) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *ModerationAPIComment) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *ModerationAPIComment) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *ModerationAPIComment) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *ModerationAPIComment) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *ModerationAPIComment) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *ModerationAPIComment) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *ModerationAPIComment) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *ModerationAPIComment) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetIsSpam

`func (o *ModerationAPIComment) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *ModerationAPIComment) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *ModerationAPIComment) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *ModerationAPIComment) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetPermNotSpam

`func (o *ModerationAPIComment) GetPermNotSpam() bool`

GetPermNotSpam returns the PermNotSpam field if non-nil, zero value otherwise.

### GetPermNotSpamOk

`func (o *ModerationAPIComment) GetPermNotSpamOk() (*bool, bool)`

GetPermNotSpamOk returns a tuple with the PermNotSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermNotSpam

`func (o *ModerationAPIComment) SetPermNotSpam(v bool)`

SetPermNotSpam sets PermNotSpam field to given value.

### HasPermNotSpam

`func (o *ModerationAPIComment) HasPermNotSpam() bool`

HasPermNotSpam returns a boolean if a field has been set.

### GetHasLinks

`func (o *ModerationAPIComment) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *ModerationAPIComment) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *ModerationAPIComment) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *ModerationAPIComment) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasCode

`func (o *ModerationAPIComment) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *ModerationAPIComment) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *ModerationAPIComment) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *ModerationAPIComment) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetApproved

`func (o *ModerationAPIComment) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ModerationAPIComment) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ModerationAPIComment) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetLocale

`func (o *ModerationAPIComment) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ModerationAPIComment) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ModerationAPIComment) SetLocale(v string)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *ModerationAPIComment) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *ModerationAPIComment) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetIsBannedUser

`func (o *ModerationAPIComment) GetIsBannedUser() bool`

GetIsBannedUser returns the IsBannedUser field if non-nil, zero value otherwise.

### GetIsBannedUserOk

`func (o *ModerationAPIComment) GetIsBannedUserOk() (*bool, bool)`

GetIsBannedUserOk returns a tuple with the IsBannedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBannedUser

`func (o *ModerationAPIComment) SetIsBannedUser(v bool)`

SetIsBannedUser sets IsBannedUser field to given value.

### HasIsBannedUser

`func (o *ModerationAPIComment) HasIsBannedUser() bool`

HasIsBannedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *ModerationAPIComment) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *ModerationAPIComment) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *ModerationAPIComment) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *ModerationAPIComment) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *ModerationAPIComment) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *ModerationAPIComment) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *ModerationAPIComment) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *ModerationAPIComment) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *ModerationAPIComment) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *ModerationAPIComment) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *ModerationAPIComment) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *ModerationAPIComment) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### SetIsPinnedNil

`func (o *ModerationAPIComment) SetIsPinnedNil(b bool)`

 SetIsPinnedNil sets the value for IsPinned to be an explicit nil

### UnsetIsPinned
`func (o *ModerationAPIComment) UnsetIsPinned()`

UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
### GetIsLocked

`func (o *ModerationAPIComment) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ModerationAPIComment) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ModerationAPIComment) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ModerationAPIComment) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *ModerationAPIComment) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *ModerationAPIComment) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetFlagCount

`func (o *ModerationAPIComment) GetFlagCount() float64`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *ModerationAPIComment) GetFlagCountOk() (*float64, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *ModerationAPIComment) SetFlagCount(v float64)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *ModerationAPIComment) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### SetFlagCountNil

`func (o *ModerationAPIComment) SetFlagCountNil(b bool)`

 SetFlagCountNil sets the value for FlagCount to be an explicit nil

### UnsetFlagCount
`func (o *ModerationAPIComment) UnsetFlagCount()`

UnsetFlagCount ensures that no value is present for FlagCount, not even an explicit nil
### GetDisplayLabel

`func (o *ModerationAPIComment) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *ModerationAPIComment) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *ModerationAPIComment) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *ModerationAPIComment) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *ModerationAPIComment) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *ModerationAPIComment) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetBadges

`func (o *ModerationAPIComment) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *ModerationAPIComment) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *ModerationAPIComment) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *ModerationAPIComment) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### SetBadgesNil

`func (o *ModerationAPIComment) SetBadgesNil(b bool)`

 SetBadgesNil sets the value for Badges to be an explicit nil

### UnsetBadges
`func (o *ModerationAPIComment) UnsetBadges()`

UnsetBadges ensures that no value is present for Badges, not even an explicit nil
### GetVerified

`func (o *ModerationAPIComment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ModerationAPIComment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ModerationAPIComment) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetFeedbackIds

`func (o *ModerationAPIComment) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *ModerationAPIComment) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *ModerationAPIComment) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *ModerationAPIComment) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ModerationAPIComment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ModerationAPIComment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ModerationAPIComment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ModerationAPIComment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


