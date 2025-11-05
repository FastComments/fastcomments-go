# GetCommentsResponseWithPresencePublicComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Status** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**TranslatedWarning** | Pointer to **string** |  | [optional] 
**Comments** | [**[]PublicComment**](PublicComment.md) |  | 
**User** | [**NullableUserSessionInfo**](UserSessionInfo.md) |  | 
**UrlIdClean** | Pointer to **string** |  | [optional] 
**LastGenDate** | Pointer to **NullableInt64** |  | [optional] 
**IncludesPastPages** | Pointer to **bool** |  | [optional] 
**IsDemo** | Pointer to **bool** |  | [optional] 
**CommentCount** | Pointer to **int32** |  | [optional] 
**IsSiteAdmin** | Pointer to **bool** |  | [optional] 
**HasBillingIssue** | Pointer to **bool** |  | [optional] 
**ModuleData** | Pointer to **map[string]map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**PageNumber** | **int32** |  | 
**IsWhiteLabeled** | Pointer to **bool** |  | [optional] 
**IsProd** | Pointer to **bool** |  | [optional] 
**IsCrawler** | Pointer to **bool** |  | [optional] 
**NotificationCount** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**PresencePollState** | Pointer to **int32** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 
**UrlIdWS** | Pointer to **string** |  | [optional] 
**UserIdWS** | Pointer to **string** |  | [optional] 
**TenantIdWS** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCommentsResponseWithPresencePublicComment

`func NewGetCommentsResponseWithPresencePublicComment(status string, comments []PublicComment, user NullableUserSessionInfo, pageNumber int32, ) *GetCommentsResponseWithPresencePublicComment`

NewGetCommentsResponseWithPresencePublicComment instantiates a new GetCommentsResponseWithPresencePublicComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentsResponseWithPresencePublicCommentWithDefaults

`func NewGetCommentsResponseWithPresencePublicCommentWithDefaults() *GetCommentsResponseWithPresencePublicComment`

NewGetCommentsResponseWithPresencePublicCommentWithDefaults instantiates a new GetCommentsResponseWithPresencePublicComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetCommentsResponseWithPresencePublicComment) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetCommentsResponseWithPresencePublicComment) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetCommentsResponseWithPresencePublicComment) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetCommentsResponseWithPresencePublicComment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentsResponseWithPresencePublicComment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *GetCommentsResponseWithPresencePublicComment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCommentsResponseWithPresencePublicComment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCommentsResponseWithPresencePublicComment) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *GetCommentsResponseWithPresencePublicComment) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetCommentsResponseWithPresencePublicComment) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetCommentsResponseWithPresencePublicComment) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTranslatedWarning

`func (o *GetCommentsResponseWithPresencePublicComment) GetTranslatedWarning() string`

GetTranslatedWarning returns the TranslatedWarning field if non-nil, zero value otherwise.

### GetTranslatedWarningOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetTranslatedWarningOk() (*string, bool)`

GetTranslatedWarningOk returns a tuple with the TranslatedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedWarning

`func (o *GetCommentsResponseWithPresencePublicComment) SetTranslatedWarning(v string)`

SetTranslatedWarning sets TranslatedWarning field to given value.

### HasTranslatedWarning

`func (o *GetCommentsResponseWithPresencePublicComment) HasTranslatedWarning() bool`

HasTranslatedWarning returns a boolean if a field has been set.

### GetComments

`func (o *GetCommentsResponseWithPresencePublicComment) GetComments() []PublicComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetCommentsOk() (*[]PublicComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetCommentsResponseWithPresencePublicComment) SetComments(v []PublicComment)`

SetComments sets Comments field to given value.


### GetUser

`func (o *GetCommentsResponseWithPresencePublicComment) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetCommentsResponseWithPresencePublicComment) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *GetCommentsResponseWithPresencePublicComment) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetCommentsResponseWithPresencePublicComment) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdClean

`func (o *GetCommentsResponseWithPresencePublicComment) GetUrlIdClean() string`

GetUrlIdClean returns the UrlIdClean field if non-nil, zero value otherwise.

### GetUrlIdCleanOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetUrlIdCleanOk() (*string, bool)`

GetUrlIdCleanOk returns a tuple with the UrlIdClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdClean

`func (o *GetCommentsResponseWithPresencePublicComment) SetUrlIdClean(v string)`

SetUrlIdClean sets UrlIdClean field to given value.

### HasUrlIdClean

`func (o *GetCommentsResponseWithPresencePublicComment) HasUrlIdClean() bool`

HasUrlIdClean returns a boolean if a field has been set.

### GetLastGenDate

`func (o *GetCommentsResponseWithPresencePublicComment) GetLastGenDate() int64`

GetLastGenDate returns the LastGenDate field if non-nil, zero value otherwise.

### GetLastGenDateOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetLastGenDateOk() (*int64, bool)`

GetLastGenDateOk returns a tuple with the LastGenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGenDate

`func (o *GetCommentsResponseWithPresencePublicComment) SetLastGenDate(v int64)`

SetLastGenDate sets LastGenDate field to given value.

### HasLastGenDate

`func (o *GetCommentsResponseWithPresencePublicComment) HasLastGenDate() bool`

HasLastGenDate returns a boolean if a field has been set.

### SetLastGenDateNil

`func (o *GetCommentsResponseWithPresencePublicComment) SetLastGenDateNil(b bool)`

 SetLastGenDateNil sets the value for LastGenDate to be an explicit nil

### UnsetLastGenDate
`func (o *GetCommentsResponseWithPresencePublicComment) UnsetLastGenDate()`

UnsetLastGenDate ensures that no value is present for LastGenDate, not even an explicit nil
### GetIncludesPastPages

`func (o *GetCommentsResponseWithPresencePublicComment) GetIncludesPastPages() bool`

GetIncludesPastPages returns the IncludesPastPages field if non-nil, zero value otherwise.

### GetIncludesPastPagesOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIncludesPastPagesOk() (*bool, bool)`

GetIncludesPastPagesOk returns a tuple with the IncludesPastPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPastPages

`func (o *GetCommentsResponseWithPresencePublicComment) SetIncludesPastPages(v bool)`

SetIncludesPastPages sets IncludesPastPages field to given value.

### HasIncludesPastPages

`func (o *GetCommentsResponseWithPresencePublicComment) HasIncludesPastPages() bool`

HasIncludesPastPages returns a boolean if a field has been set.

### GetIsDemo

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *GetCommentsResponseWithPresencePublicComment) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *GetCommentsResponseWithPresencePublicComment) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetCommentCount

`func (o *GetCommentsResponseWithPresencePublicComment) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetCommentsResponseWithPresencePublicComment) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *GetCommentsResponseWithPresencePublicComment) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetIsSiteAdmin

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsSiteAdmin() bool`

GetIsSiteAdmin returns the IsSiteAdmin field if non-nil, zero value otherwise.

### GetIsSiteAdminOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsSiteAdminOk() (*bool, bool)`

GetIsSiteAdminOk returns a tuple with the IsSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSiteAdmin

`func (o *GetCommentsResponseWithPresencePublicComment) SetIsSiteAdmin(v bool)`

SetIsSiteAdmin sets IsSiteAdmin field to given value.

### HasIsSiteAdmin

`func (o *GetCommentsResponseWithPresencePublicComment) HasIsSiteAdmin() bool`

HasIsSiteAdmin returns a boolean if a field has been set.

### GetHasBillingIssue

`func (o *GetCommentsResponseWithPresencePublicComment) GetHasBillingIssue() bool`

GetHasBillingIssue returns the HasBillingIssue field if non-nil, zero value otherwise.

### GetHasBillingIssueOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetHasBillingIssueOk() (*bool, bool)`

GetHasBillingIssueOk returns a tuple with the HasBillingIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBillingIssue

`func (o *GetCommentsResponseWithPresencePublicComment) SetHasBillingIssue(v bool)`

SetHasBillingIssue sets HasBillingIssue field to given value.

### HasHasBillingIssue

`func (o *GetCommentsResponseWithPresencePublicComment) HasHasBillingIssue() bool`

HasHasBillingIssue returns a boolean if a field has been set.

### GetModuleData

`func (o *GetCommentsResponseWithPresencePublicComment) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *GetCommentsResponseWithPresencePublicComment) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *GetCommentsResponseWithPresencePublicComment) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetPageNumber

`func (o *GetCommentsResponseWithPresencePublicComment) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *GetCommentsResponseWithPresencePublicComment) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetIsWhiteLabeled

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsWhiteLabeled() bool`

GetIsWhiteLabeled returns the IsWhiteLabeled field if non-nil, zero value otherwise.

### GetIsWhiteLabeledOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsWhiteLabeledOk() (*bool, bool)`

GetIsWhiteLabeledOk returns a tuple with the IsWhiteLabeled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteLabeled

`func (o *GetCommentsResponseWithPresencePublicComment) SetIsWhiteLabeled(v bool)`

SetIsWhiteLabeled sets IsWhiteLabeled field to given value.

### HasIsWhiteLabeled

`func (o *GetCommentsResponseWithPresencePublicComment) HasIsWhiteLabeled() bool`

HasIsWhiteLabeled returns a boolean if a field has been set.

### GetIsProd

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *GetCommentsResponseWithPresencePublicComment) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *GetCommentsResponseWithPresencePublicComment) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetIsCrawler

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsCrawler() bool`

GetIsCrawler returns the IsCrawler field if non-nil, zero value otherwise.

### GetIsCrawlerOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsCrawlerOk() (*bool, bool)`

GetIsCrawlerOk returns a tuple with the IsCrawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCrawler

`func (o *GetCommentsResponseWithPresencePublicComment) SetIsCrawler(v bool)`

SetIsCrawler sets IsCrawler field to given value.

### HasIsCrawler

`func (o *GetCommentsResponseWithPresencePublicComment) HasIsCrawler() bool`

HasIsCrawler returns a boolean if a field has been set.

### GetNotificationCount

`func (o *GetCommentsResponseWithPresencePublicComment) GetNotificationCount() int32`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetNotificationCountOk() (*int32, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *GetCommentsResponseWithPresencePublicComment) SetNotificationCount(v int32)`

SetNotificationCount sets NotificationCount field to given value.

### HasNotificationCount

`func (o *GetCommentsResponseWithPresencePublicComment) HasNotificationCount() bool`

HasNotificationCount returns a boolean if a field has been set.

### GetHasMore

`func (o *GetCommentsResponseWithPresencePublicComment) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCommentsResponseWithPresencePublicComment) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetCommentsResponseWithPresencePublicComment) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetIsClosed

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *GetCommentsResponseWithPresencePublicComment) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *GetCommentsResponseWithPresencePublicComment) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetPresencePollState

`func (o *GetCommentsResponseWithPresencePublicComment) GetPresencePollState() int32`

GetPresencePollState returns the PresencePollState field if non-nil, zero value otherwise.

### GetPresencePollStateOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetPresencePollStateOk() (*int32, bool)`

GetPresencePollStateOk returns a tuple with the PresencePollState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePollState

`func (o *GetCommentsResponseWithPresencePublicComment) SetPresencePollState(v int32)`

SetPresencePollState sets PresencePollState field to given value.

### HasPresencePollState

`func (o *GetCommentsResponseWithPresencePublicComment) HasPresencePollState() bool`

HasPresencePollState returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetCommentsResponseWithPresencePublicComment) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetCommentsResponseWithPresencePublicComment) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetCommentsResponseWithPresencePublicComment) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetUrlIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.

### HasUrlIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) HasUrlIdWS() bool`

HasUrlIdWS returns a boolean if a field has been set.

### GetUserIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetTenantIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) GetTenantIdWS() string`

GetTenantIdWS returns the TenantIdWS field if non-nil, zero value otherwise.

### GetTenantIdWSOk

`func (o *GetCommentsResponseWithPresencePublicComment) GetTenantIdWSOk() (*string, bool)`

GetTenantIdWSOk returns a tuple with the TenantIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) SetTenantIdWS(v string)`

SetTenantIdWS sets TenantIdWS field to given value.

### HasTenantIdWS

`func (o *GetCommentsResponseWithPresencePublicComment) HasTenantIdWS() bool`

HasTenantIdWS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


