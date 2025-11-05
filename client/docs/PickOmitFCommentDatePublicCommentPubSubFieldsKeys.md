# PickOmitFCommentDatePublicCommentPubSubFieldsKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **interface{}** |  | 
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UrlId** | **string** |  | 
**Url** | **string** |  | 
**PageTitle** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AnonUserId** | Pointer to **string** |  | [optional] 
**CommenterName** | **string** |  | 
**CommenterLink** | Pointer to **string** |  | [optional] 
**Comment** | **string** |  | 
**CommentHTML** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VotesUp** | Pointer to **int32** |  | [optional] 
**VotesDown** | Pointer to **int32** |  | [optional] 
**ExpireAt** | Pointer to **time.Time** |  | [optional] 
**Verified** | **bool** |  | 
**Reviewed** | Pointer to **bool** |  | [optional] 
**AvatarSrc** | Pointer to **string** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**HasImages** | Pointer to **bool** |  | [optional] 
**HasLinks** | Pointer to **bool** |  | [optional] 
**HasCode** | Pointer to **bool** |  | [optional] 
**Approved** | **bool** |  | 
**Locale** | **string** |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDeletedUser** | Pointer to **bool** |  | [optional] 
**IsBannedUser** | Pointer to **bool** |  | [optional] 
**IsByAdmin** | Pointer to **bool** |  | [optional] 
**IsByModerator** | Pointer to **bool** |  | [optional] 
**IsPinned** | Pointer to **bool** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**FlagCount** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **float64** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**Badges** | Pointer to [**[]CommentUserBadgeInfo**](CommentUserBadgeInfo.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**FeedbackIds** | Pointer to **[]string** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**ViewCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewPickOmitFCommentDatePublicCommentPubSubFieldsKeys

`func NewPickOmitFCommentDatePublicCommentPubSubFieldsKeys(date interface{}, id string, tenantId string, urlId string, url string, commenterName string, comment string, commentHTML string, verified bool, approved bool, locale string, ) *PickOmitFCommentDatePublicCommentPubSubFieldsKeys`

NewPickOmitFCommentDatePublicCommentPubSubFieldsKeys instantiates a new PickOmitFCommentDatePublicCommentPubSubFieldsKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickOmitFCommentDatePublicCommentPubSubFieldsKeysWithDefaults

`func NewPickOmitFCommentDatePublicCommentPubSubFieldsKeysWithDefaults() *PickOmitFCommentDatePublicCommentPubSubFieldsKeys`

NewPickOmitFCommentDatePublicCommentPubSubFieldsKeysWithDefaults instantiates a new PickOmitFCommentDatePublicCommentPubSubFieldsKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetDate() interface{}`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetDateOk() (*interface{}, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetDate(v interface{})`

SetDate sets Date field to given value.


### SetDateNil

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUrlId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetUrlId() string`

GetUrlId returns the UrlId field if non-nil, zero value otherwise.

### GetUrlIdOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetUrlIdOk() (*string, bool)`

GetUrlIdOk returns a tuple with the UrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetUrlId(v string)`

SetUrlId sets UrlId field to given value.


### GetUrl

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPageTitle

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetUserId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAnonUserId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetAnonUserId() string`

GetAnonUserId returns the AnonUserId field if non-nil, zero value otherwise.

### GetAnonUserIdOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetAnonUserIdOk() (*string, bool)`

GetAnonUserIdOk returns a tuple with the AnonUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonUserId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetAnonUserId(v string)`

SetAnonUserId sets AnonUserId field to given value.

### HasAnonUserId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasAnonUserId() bool`

HasAnonUserId returns a boolean if a field has been set.

### GetCommenterName

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetCommenterLink

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommenterLink() string`

GetCommenterLink returns the CommenterLink field if non-nil, zero value otherwise.

### GetCommenterLinkOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommenterLinkOk() (*string, bool)`

GetCommenterLinkOk returns a tuple with the CommenterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterLink

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetCommenterLink(v string)`

SetCommenterLink sets CommenterLink field to given value.

### HasCommenterLink

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasCommenterLink() bool`

HasCommenterLink returns a boolean if a field has been set.

### GetComment

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCommentHTML

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommentHTML() string`

GetCommentHTML returns the CommentHTML field if non-nil, zero value otherwise.

### GetCommentHTMLOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetCommentHTMLOk() (*string, bool)`

GetCommentHTMLOk returns a tuple with the CommentHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTML

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetCommentHTML(v string)`

SetCommentHTML sets CommentHTML field to given value.


### GetParentId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetVotes

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVotesUp

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVotesUp() int32`

GetVotesUp returns the VotesUp field if non-nil, zero value otherwise.

### GetVotesUpOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVotesUpOk() (*int32, bool)`

GetVotesUpOk returns a tuple with the VotesUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesUp

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetVotesUp(v int32)`

SetVotesUp sets VotesUp field to given value.

### HasVotesUp

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasVotesUp() bool`

HasVotesUp returns a boolean if a field has been set.

### GetVotesDown

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVotesDown() int32`

GetVotesDown returns the VotesDown field if non-nil, zero value otherwise.

### GetVotesDownOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVotesDownOk() (*int32, bool)`

GetVotesDownOk returns a tuple with the VotesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesDown

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetVotesDown(v int32)`

SetVotesDown sets VotesDown field to given value.

### HasVotesDown

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasVotesDown() bool`

HasVotesDown returns a boolean if a field has been set.

### GetExpireAt

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetVerified

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetReviewed

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetAvatarSrc

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetAvatarSrc() string`

GetAvatarSrc returns the AvatarSrc field if non-nil, zero value otherwise.

### GetAvatarSrcOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetAvatarSrcOk() (*string, bool)`

GetAvatarSrcOk returns a tuple with the AvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSrc

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetAvatarSrc(v string)`

SetAvatarSrc sets AvatarSrc field to given value.

### HasAvatarSrc

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasAvatarSrc() bool`

HasAvatarSrc returns a boolean if a field has been set.

### GetIsSpam

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetHasImages

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetHasImages() bool`

GetHasImages returns the HasImages field if non-nil, zero value otherwise.

### GetHasImagesOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetHasImagesOk() (*bool, bool)`

GetHasImagesOk returns a tuple with the HasImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasImages

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetHasImages(v bool)`

SetHasImages sets HasImages field to given value.

### HasHasImages

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasHasImages() bool`

HasHasImages returns a boolean if a field has been set.

### GetHasLinks

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetHasLinks() bool`

GetHasLinks returns the HasLinks field if non-nil, zero value otherwise.

### GetHasLinksOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetHasLinksOk() (*bool, bool)`

GetHasLinksOk returns a tuple with the HasLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinks

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetHasLinks(v bool)`

SetHasLinks sets HasLinks field to given value.

### HasHasLinks

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasHasLinks() bool`

HasHasLinks returns a boolean if a field has been set.

### GetHasCode

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetHasCode() bool`

GetHasCode returns the HasCode field if non-nil, zero value otherwise.

### GetHasCodeOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetHasCodeOk() (*bool, bool)`

GetHasCodeOk returns a tuple with the HasCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCode

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetHasCode(v bool)`

SetHasCode sets HasCode field to given value.

### HasHasCode

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasHasCode() bool`

HasHasCode returns a boolean if a field has been set.

### GetApproved

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetLocale

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetIsDeleted

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDeletedUser

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsDeletedUser() bool`

GetIsDeletedUser returns the IsDeletedUser field if non-nil, zero value otherwise.

### GetIsDeletedUserOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsDeletedUserOk() (*bool, bool)`

GetIsDeletedUserOk returns a tuple with the IsDeletedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletedUser

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsDeletedUser(v bool)`

SetIsDeletedUser sets IsDeletedUser field to given value.

### HasIsDeletedUser

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsDeletedUser() bool`

HasIsDeletedUser returns a boolean if a field has been set.

### GetIsBannedUser

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsBannedUser() bool`

GetIsBannedUser returns the IsBannedUser field if non-nil, zero value otherwise.

### GetIsBannedUserOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsBannedUserOk() (*bool, bool)`

GetIsBannedUserOk returns a tuple with the IsBannedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBannedUser

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsBannedUser(v bool)`

SetIsBannedUser sets IsBannedUser field to given value.

### HasIsBannedUser

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsBannedUser() bool`

HasIsBannedUser returns a boolean if a field has been set.

### GetIsByAdmin

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsByAdmin() bool`

GetIsByAdmin returns the IsByAdmin field if non-nil, zero value otherwise.

### GetIsByAdminOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsByAdminOk() (*bool, bool)`

GetIsByAdminOk returns a tuple with the IsByAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByAdmin

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsByAdmin(v bool)`

SetIsByAdmin sets IsByAdmin field to given value.

### HasIsByAdmin

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsByAdmin() bool`

HasIsByAdmin returns a boolean if a field has been set.

### GetIsByModerator

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsByModerator() bool`

GetIsByModerator returns the IsByModerator field if non-nil, zero value otherwise.

### GetIsByModeratorOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsByModeratorOk() (*bool, bool)`

GetIsByModeratorOk returns a tuple with the IsByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByModerator

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsByModerator(v bool)`

SetIsByModerator sets IsByModerator field to given value.

### HasIsByModerator

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsByModerator() bool`

HasIsByModerator returns a boolean if a field has been set.

### GetIsPinned

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsPinned() bool`

GetIsPinned returns the IsPinned field if non-nil, zero value otherwise.

### GetIsPinnedOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsPinnedOk() (*bool, bool)`

GetIsPinnedOk returns a tuple with the IsPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinned

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsPinned(v bool)`

SetIsPinned sets IsPinned field to given value.

### HasIsPinned

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsPinned() bool`

HasIsPinned returns a boolean if a field has been set.

### GetIsLocked

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetFlagCount

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### GetRating

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetBadges

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetBadges() []CommentUserBadgeInfo`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetBadgesOk() (*[]CommentUserBadgeInfo, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetBadges(v []CommentUserBadgeInfo)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetDomain

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFeedbackIds

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetFeedbackIds() []string`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetFeedbackIdsOk() (*[]string, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetFeedbackIds(v []string)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### GetGroupIds

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetViewCount

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *PickOmitFCommentDatePublicCommentPubSubFieldsKeys) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


