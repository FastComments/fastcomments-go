# CreateCommentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **int64** |  | [optional] 
**LocalDateString** | Pointer to **string** |  | [optional] 
**LocalDateHours** | Pointer to **int32** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterEmail** | Pointer to **NullableString** |  | [optional] 
**CommenterLink** | Pointer to **NullableString** |  | [optional] 
**Comment** | **string** |  | 
**ProductId** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**AvatarSrc** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Mentions** | Pointer to [**[]CommentUserMentionInfo**](CommentUserMentionInfo.md) |  | [optional] 
**HashTags** | Pointer to [**[]CommentUserHashTagInfo**](CommentUserHashTagInfo.md) |  | [optional] 
**PageTitle** | Pointer to **string** |  | [optional] 
**IsFromMyAccountPage** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 
**UrlId** | **string** |  | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**ModerationGroupIds** | Pointer to **[]string** |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**FromOfflineRestore** | Pointer to **bool** |  | [optional] 
**AutoplayDelayMS** | Pointer to **int64** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**QuestionValues** | Pointer to [**map[string]RecordStringStringOrNumberValue**](RecordStringStringOrNumberValue.md) | Construct a type with a set of properties K of type T | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IsPinned** | Pointer to **bool** |  | [optional] 
**Locale** | **string** | Example: en_us | 
**Reviewed** | Pointer to **bool** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VotesDown** | Pointer to **int32** |  | [optional] 
**VotesUp** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateCommentParams

`func NewCreateCommentParams(commenterName string, comment string, url string, urlId string, locale string, ) *CreateCommentParams`

NewCreateCommentParams instantiates a new CreateCommentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentParamsWithDefaults

`func NewCreateCommentParamsWithDefaults() *CreateCommentParams`

NewCreateCommentParamsWithDefaults instantiates a new CreateCommentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CreateCommentParams) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateCommentParams) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateCommentParams) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *CreateCommentParams) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLocalDateString

`func (o *CreateCommentParams) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *CreateCommentParams) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *CreateCommentParams) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *CreateCommentParams) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### GetLocalDateHours

`func (o *CreateCommentParams) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *CreateCommentParams) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *CreateCommentParams) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *CreateCommentParams) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### GetCommenterName

`func (o *CreateCommentParams) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *CreateCommentParams) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *CreateCommentParams) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterEmail

`func (o *CreateCommentParams) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *CreateCommentParams) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *CreateCommentParams) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *CreateCommentParams) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *CreateCommentParams) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *CreateCommentParams) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetCommenterLink

`func (o *CreateCommentParams) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *CreateCommentParams) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *CreateCommentParams) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *CreateCommentParams) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *CreateCommentParams) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *CreateCommentParams) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetComment

`func (o *CreateCommentParams) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateCommentParams) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateCommentParams) SetComment(v string)`

SetComment sets Comment field to given value.


### GetProductId

`func (o *CreateCommentParams) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateCommentParams) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateCommentParams) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CreateCommentParams) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *CreateCommentParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateCommentParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateCommentParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateCommentParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateCommentParams) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateCommentParams) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAvatarSrc

`func (o *CreateCommentParams) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *CreateCommentParams) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *CreateCommentParams) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *CreateCommentParams) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *CreateCommentParams) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *CreateCommentParams) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetParentId

`func (o *CreateCommentParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateCommentParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateCommentParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateCommentParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateCommentParams) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateCommentParams) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetMentions

`func (o *CreateCommentParams) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *CreateCommentParams) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *CreateCommentParams) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *CreateCommentParams) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetHashTags

`func (o *CreateCommentParams) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *CreateCommentParams) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *CreateCommentParams) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *CreateCommentParams) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetPageTitle

`func (o *CreateCommentParams) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *CreateCommentParams) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *CreateCommentParams) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *CreateCommentParams) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetIsFromMyAccountPage

`func (o *CreateCommentParams) GetIsFromMyAccountPage() bool`

GetIsFromMyAccountPage returns the IsFromMyAccountPage field if non-nil, zero value otherwise.

### GetIsFromMyAccountPageOk

`func (o *CreateCommentParams) GetIsFromMyAccountPageOk() (*bool, bool)`

GetIsFromMyAccountPageOk returns a tuple with the IsFromMyAccountPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromMyAccountPage

`func (o *CreateCommentParams) SetIsFromMyAccountPage(v bool)`

SetIsFromMyAccountPage sets IsFromMyAccountPage field to given value.

### HasIsFromMyAccountPage

`func (o *CreateCommentParams) HasIsFromMyAccountPage() bool`

HasIsFromMyAccountPage returns a boolean if a field has been set.

### GetUrl

`func (o *CreateCommentParams) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateCommentParams) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateCommentParams) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlId

`func (o *CreateCommentParams) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *CreateCommentParams) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *CreateCommentParams) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetMeta

`func (o *CreateCommentParams) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateCommentParams) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateCommentParams) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateCommentParams) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *CreateCommentParams) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *CreateCommentParams) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *CreateCommentParams) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *CreateCommentParams) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### GetRating

`func (o *CreateCommentParams) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CreateCommentParams) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CreateCommentParams) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *CreateCommentParams) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFromOfflineRestore

`func (o *CreateCommentParams) GetFromOfflineRestore() bool`

GetFromOfflineRestore returns the FromOfflineRestore field if non-nil, zero value otherwise.

### GetFromOfflineRestoreOk

`func (o *CreateCommentParams) GetFromOfflineRestoreOk() (*bool, bool)`

GetFromOfflineRestoreOk returns a tuple with the FromOfflineRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOfflineRestore

`func (o *CreateCommentParams) SetFromOfflineRestore(v bool)`

SetFromOfflineRestore sets FromOfflineRestore field to given value.

### HasFromOfflineRestore

`func (o *CreateCommentParams) HasFromOfflineRestore() bool`

HasFromOfflineRestore returns a boolean if a field has been set.

### GetAutoplayDelayMS

`func (o *CreateCommentParams) GetAutoplayDelayMS() int64`

GetAutoplayDelayMS returns the AutoplayDelayMS field if non-nil, zero value otherwise.

### GetAutoplayDelayMSOk

`func (o *CreateCommentParams) GetAutoplayDelayMSOk() (*int64, bool)`

GetAutoplayDelayMSOk returns a tuple with the AutoplayDelayMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoplayDelayMS

`func (o *CreateCommentParams) SetAutoplayDelayMS(v int64)`

SetAutoplayDelayMS sets AutoplayDelayMS field to given value.

### HasAutoplayDelayMS

`func (o *CreateCommentParams) HasAutoplayDelayMS() bool`

HasAutoplayDelayMS returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *CreateCommentParams) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *CreateCommentParams) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *CreateCommentParams) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *CreateCommentParams) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetQuestionValues

`func (o *CreateCommentParams) GetQuestionValues() map[string]RecordStringStringOrNumberValue`

GetQuestionValues returns the QuestionValues field if non-nil, zero value otherwise.

### GetQuestionValuesOk

`func (o *CreateCommentParams) GetQuestionValuesOk() (*map[string]RecordStringStringOrNumberValue, bool)`

GetQuestionValuesOk returns a tuple with the QuestionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionValues

`func (o *CreateCommentParams) SetQuestionValues(v map[string]RecordStringStringOrNumberValue)`

SetQuestionValues sets QuestionValues field to given value.

### HasQuestionValues

`func (o *CreateCommentParams) HasQuestionValues() bool`

HasQuestionValues returns a boolean if a field has been set.

### GetApproved

`func (o *CreateCommentParams) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *CreateCommentParams) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *CreateCommentParams) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *CreateCommentParams) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetDomain

`func (o *CreateCommentParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateCommentParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateCommentParams) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateCommentParams) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIp

`func (o *CreateCommentParams) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateCommentParams) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateCommentParams) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateCommentParams) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsPinned

`func (o *CreateCommentParams) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *CreateCommentParams) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *CreateCommentParams) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *CreateCommentParams) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetLocale

`func (o *CreateCommentParams) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateCommentParams) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateCommentParams) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetReviewed

`func (o *CreateCommentParams) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *CreateCommentParams) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *CreateCommentParams) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *CreateCommentParams) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetVerified

`func (o *CreateCommentParams) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *CreateCommentParams) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *CreateCommentParams) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *CreateCommentParams) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVotes

`func (o *CreateCommentParams) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *CreateCommentParams) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *CreateCommentParams) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *CreateCommentParams) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesDown

`func (o *CreateCommentParams) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *CreateCommentParams) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *CreateCommentParams) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *CreateCommentParams) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetVotesUp

`func (o *CreateCommentParams) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *CreateCommentParams) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *CreateCommentParams) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *CreateCommentParams) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


