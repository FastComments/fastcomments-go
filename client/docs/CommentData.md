# CommentData

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

## Methods

### NewCommentData

`func NewCommentData(commenterName string, comment string, url string, urlId string, ) *CommentData`

NewCommentData instantiates a new CommentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentDataWithDefaults

`func NewCommentDataWithDefaults() *CommentData`

NewCommentDataWithDefaults instantiates a new CommentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CommentData) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CommentData) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CommentData) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *CommentData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLocalDateString

`func (o *CommentData) GetLocalDateString() string`

GetLocalDateString returns the LocalDateString field if non-nil, zero value otherwise.

### GetLocalDateStringOk

`func (o *CommentData) GetLocalDateStringOk() (*string, bool)`

GetLocalDateStringOk returns a tuple with the LocalDateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateString

`func (o *CommentData) SetLocalDateString(v string)`

SetLocalDateString sets LocalDateString field to given value.

### HasLocalDateString

`func (o *CommentData) HasLocalDateString() bool`

HasLocalDateString returns a boolean if a field has been set.

### GetLocalDateHours

`func (o *CommentData) GetLocalDateHours() int32`

GetLocalDateHours returns the LocalDateHours field if non-nil, zero value otherwise.

### GetLocalDateHoursOk

`func (o *CommentData) GetLocalDateHoursOk() (*int32, bool)`

GetLocalDateHoursOk returns a tuple with the LocalDateHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDateHours

`func (o *CommentData) SetLocalDateHours(v int32)`

SetLocalDateHours sets LocalDateHours field to given value.

### HasLocalDateHours

`func (o *CommentData) HasLocalDateHours() bool`

HasLocalDateHours returns a boolean if a field has been set.

### GetCommenterName

`func (o *CommentData) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *CommentData) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *CommentData) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterEmail

`func (o *CommentData) GetCommenterEmail() string`

GetCommenterEmail returns the CommenterEmail field if non-nil, zero value otherwise.

### GetCommenterEmailOk

`func (o *CommentData) GetCommenterEmailOk() (*string, bool)`

GetCommenterEmailOk returns a tuple with the CommenterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterEmail

`func (o *CommentData) SetCommenterEmail(v string)`

SetCommenterEmail sets CommenterEmail field to given value.

### HasCommenterEmail

`func (o *CommentData) HasCommenterEmail() bool`

HasCommenterEmail returns a boolean if a field has been set.

### SetCommenterEmailNil

`func (o *CommentData) SetCommenterEmailNil(b bool)`

 SetCommenterEmailNil sets the value for CommenterEmail to be an explicit nil

### UnsetCommenterEmail
`func (o *CommentData) UnsetCommenterEmail()`

UnsetCommenterEmail ensures that no value is present for CommenterEmail, not even an explicit nil
### GetCommenterLink

`func (o *CommentData) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *CommentData) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *CommentData) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *CommentData) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### SetCommenterLinkNil

`func (o *CommentData) SetCommenterLinkNil(b bool)`

 SetCommenterLinkNil sets the value for CommenterLink to be an explicit nil

### UnsetCommenterLink
`func (o *CommentData) UnsetCommenterLink()`

UnsetCommenterLink ensures that no value is present for CommenterLink, not even an explicit nil
### GetComment

`func (o *CommentData) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CommentData) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CommentData) SetComment(v string)`

SetComment sets Comment field to given value.


### GetProductId

`func (o *CommentData) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CommentData) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CommentData) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CommentData) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *CommentData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CommentData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CommentData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CommentData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CommentData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CommentData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAvatarSrc

`func (o *CommentData) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *CommentData) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *CommentData) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *CommentData) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### SetAvatarSrcNil

`func (o *CommentData) SetAvatarSrcNil(b bool)`

 SetAvatarSrcNil sets the value for AvatarSrc to be an explicit nil

### UnsetAvatarSrc
`func (o *CommentData) UnsetAvatarSrc()`

UnsetAvatarSrc ensures that no value is present for AvatarSrc, not even an explicit nil
### GetParentId

`func (o *CommentData) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CommentData) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CommentData) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CommentData) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CommentData) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CommentData) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetMentions

`func (o *CommentData) GetMentions() []CommentUserMentionInfo`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *CommentData) GetMentionsOk() (*[]CommentUserMentionInfo, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *CommentData) SetMentions(v []CommentUserMentionInfo)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *CommentData) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetHashTags

`func (o *CommentData) GetHashTags() []CommentUserHashTagInfo`

GetHashTags returns the HashTags field if non-nil, zero value otherwise.

### GetHashTagsOk

`func (o *CommentData) GetHashTagsOk() (*[]CommentUserHashTagInfo, bool)`

GetHashTagsOk returns a tuple with the HashTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTags

`func (o *CommentData) SetHashTags(v []CommentUserHashTagInfo)`

SetHashTags sets HashTags field to given value.

### HasHashTags

`func (o *CommentData) HasHashTags() bool`

HasHashTags returns a boolean if a field has been set.

### GetPageTitle

`func (o *CommentData) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *CommentData) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *CommentData) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *CommentData) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetIsFromMyAccountPage

`func (o *CommentData) GetIsFromMyAccountPage() bool`

GetIsFromMyAccountPage returns the IsFromMyAccountPage field if non-nil, zero value otherwise.

### GetIsFromMyAccountPageOk

`func (o *CommentData) GetIsFromMyAccountPageOk() (*bool, bool)`

GetIsFromMyAccountPageOk returns a tuple with the IsFromMyAccountPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromMyAccountPage

`func (o *CommentData) SetIsFromMyAccountPage(v bool)`

SetIsFromMyAccountPage sets IsFromMyAccountPage field to given value.

### HasIsFromMyAccountPage

`func (o *CommentData) HasIsFromMyAccountPage() bool`

HasIsFromMyAccountPage returns a boolean if a field has been set.

### GetUrl

`func (o *CommentData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommentData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommentData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlId

`func (o *CommentData) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *CommentData) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *CommentData) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetMeta

`func (o *CommentData) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CommentData) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CommentData) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CommentData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModerationGroupIds

`func (o *CommentData) GetModerationGroupIds() []string`

GetModerationGroupIds returns the ModerationGroupIds field if non-nil, zero value otherwise.

### GetModerationGroupIdsOk

`func (o *CommentData) GetModerationGroupIdsOk() (*[]string, bool)`

GetModerationGroupIdsOk returns a tuple with the ModerationGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationGroupIds

`func (o *CommentData) SetModerationGroupIds(v []string)`

SetModerationGroupIds sets ModerationGroupIds field to given value.

### HasModerationGroupIds

`func (o *CommentData) HasModerationGroupIds() bool`

HasModerationGroupIds returns a boolean if a field has been set.

### GetRating

`func (o *CommentData) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CommentData) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CommentData) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *CommentData) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFromOfflineRestore

`func (o *CommentData) GetFromOfflineRestore() bool`

GetFromOfflineRestore returns the FromOfflineRestore field if non-nil, zero value otherwise.

### GetFromOfflineRestoreOk

`func (o *CommentData) GetFromOfflineRestoreOk() (*bool, bool)`

GetFromOfflineRestoreOk returns a tuple with the FromOfflineRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOfflineRestore

`func (o *CommentData) SetFromOfflineRestore(v bool)`

SetFromOfflineRestore sets FromOfflineRestore field to given value.

### HasFromOfflineRestore

`func (o *CommentData) HasFromOfflineRestore() bool`

HasFromOfflineRestore returns a boolean if a field has been set.

### GetAutoplayDelayMS

`func (o *CommentData) GetAutoplayDelayMS() int64`

GetAutoplayDelayMS returns the AutoplayDelayMS field if non-nil, zero value otherwise.

### GetAutoplayDelayMSOk

`func (o *CommentData) GetAutoplayDelayMSOk() (*int64, bool)`

GetAutoplayDelayMSOk returns a tuple with the AutoplayDelayMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoplayDelayMS

`func (o *CommentData) SetAutoplayDelayMS(v int64)`

SetAutoplayDelayMS sets AutoplayDelayMS field to given value.

### HasAutoplayDelayMS

`func (o *CommentData) HasAutoplayDelayMS() bool`

HasAutoplayDelayMS returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *CommentData) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *CommentData) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *CommentData) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *CommentData) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetQuestionValues

`func (o *CommentData) GetQuestionValues() map[string]RecordStringStringOrNumberValue`

GetQuestionValues returns the QuestionValues field if non-nil, zero value otherwise.

### GetQuestionValuesOk

`func (o *CommentData) GetQuestionValuesOk() (*map[string]RecordStringStringOrNumberValue, bool)`

GetQuestionValuesOk returns a tuple with the QuestionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionValues

`func (o *CommentData) SetQuestionValues(v map[string]RecordStringStringOrNumberValue)`

SetQuestionValues sets QuestionValues field to given value.

### HasQuestionValues

`func (o *CommentData) HasQuestionValues() bool`

HasQuestionValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


