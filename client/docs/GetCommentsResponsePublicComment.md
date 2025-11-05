# GetCommentsResponsePublicComment

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
**ModuleData** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**PageNumber** | **int32** |  | 
**IsWhiteLabeled** | Pointer to **bool** |  | [optional] 
**IsProd** | Pointer to **bool** |  | [optional] 
**IsCrawler** | Pointer to **bool** |  | [optional] 
**NotificationCount** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**IsClosed** | Pointer to **bool** |  | [optional] 
**PresencePollState** | Pointer to **int32** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetCommentsResponsePublicComment

`func NewGetCommentsResponsePublicComment(status string, comments []PublicComment, user NullableUserSessionInfo, pageNumber int32, ) *GetCommentsResponsePublicComment`

NewGetCommentsResponsePublicComment instantiates a new GetCommentsResponsePublicComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentsResponsePublicCommentWithDefaults

`func NewGetCommentsResponsePublicCommentWithDefaults() *GetCommentsResponsePublicComment`

NewGetCommentsResponsePublicCommentWithDefaults instantiates a new GetCommentsResponsePublicComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetCommentsResponsePublicComment) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetCommentsResponsePublicComment) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetCommentsResponsePublicComment) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetCommentsResponsePublicComment) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetCommentsResponsePublicComment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentsResponsePublicComment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentsResponsePublicComment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *GetCommentsResponsePublicComment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCommentsResponsePublicComment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCommentsResponsePublicComment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCommentsResponsePublicComment) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *GetCommentsResponsePublicComment) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetCommentsResponsePublicComment) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetCommentsResponsePublicComment) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetCommentsResponsePublicComment) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTranslatedWarning

`func (o *GetCommentsResponsePublicComment) GetTranslatedWarning() string`

GetTranslatedWarning returns the TranslatedWarning field if non-nil, zero value otherwise.

### GetTranslatedWarningOk

`func (o *GetCommentsResponsePublicComment) GetTranslatedWarningOk() (*string, bool)`

GetTranslatedWarningOk returns a tuple with the TranslatedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedWarning

`func (o *GetCommentsResponsePublicComment) SetTranslatedWarning(v string)`

SetTranslatedWarning sets TranslatedWarning field to given value.

### HasTranslatedWarning

`func (o *GetCommentsResponsePublicComment) HasTranslatedWarning() bool`

HasTranslatedWarning returns a boolean if a field has been set.

### GetComments

`func (o *GetCommentsResponsePublicComment) GetComments() []PublicComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetCommentsResponsePublicComment) GetCommentsOk() (*[]PublicComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetCommentsResponsePublicComment) SetComments(v []PublicComment)`

SetComments sets Comments field to given value.


### GetUser

`func (o *GetCommentsResponsePublicComment) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetCommentsResponsePublicComment) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetCommentsResponsePublicComment) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *GetCommentsResponsePublicComment) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetCommentsResponsePublicComment) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdClean

`func (o *GetCommentsResponsePublicComment) GetUrlIdClean() string`

GetUrlIdClean returns the UrlIdClean field if non-nil, zero value otherwise.

### GetUrlIdCleanOk

`func (o *GetCommentsResponsePublicComment) GetUrlIdCleanOk() (*string, bool)`

GetUrlIdCleanOk returns a tuple with the UrlIdClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdClean

`func (o *GetCommentsResponsePublicComment) SetUrlIdClean(v string)`

SetUrlIdClean sets UrlIdClean field to given value.

### HasUrlIdClean

`func (o *GetCommentsResponsePublicComment) HasUrlIdClean() bool`

HasUrlIdClean returns a boolean if a field has been set.

### GetLastGenDate

`func (o *GetCommentsResponsePublicComment) GetLastGenDate() int64`

GetLastGenDate returns the LastGenDate field if non-nil, zero value otherwise.

### GetLastGenDateOk

`func (o *GetCommentsResponsePublicComment) GetLastGenDateOk() (*int64, bool)`

GetLastGenDateOk returns a tuple with the LastGenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGenDate

`func (o *GetCommentsResponsePublicComment) SetLastGenDate(v int64)`

SetLastGenDate sets LastGenDate field to given value.

### HasLastGenDate

`func (o *GetCommentsResponsePublicComment) HasLastGenDate() bool`

HasLastGenDate returns a boolean if a field has been set.

### SetLastGenDateNil

`func (o *GetCommentsResponsePublicComment) SetLastGenDateNil(b bool)`

 SetLastGenDateNil sets the value for LastGenDate to be an explicit nil

### UnsetLastGenDate
`func (o *GetCommentsResponsePublicComment) UnsetLastGenDate()`

UnsetLastGenDate ensures that no value is present for LastGenDate, not even an explicit nil
### GetIncludesPastPages

`func (o *GetCommentsResponsePublicComment) GetIncludesPastPages() bool`

GetIncludesPastPages returns the IncludesPastPages field if non-nil, zero value otherwise.

### GetIncludesPastPagesOk

`func (o *GetCommentsResponsePublicComment) GetIncludesPastPagesOk() (*bool, bool)`

GetIncludesPastPagesOk returns a tuple with the IncludesPastPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPastPages

`func (o *GetCommentsResponsePublicComment) SetIncludesPastPages(v bool)`

SetIncludesPastPages sets IncludesPastPages field to given value.

### HasIncludesPastPages

`func (o *GetCommentsResponsePublicComment) HasIncludesPastPages() bool`

HasIncludesPastPages returns a boolean if a field has been set.

### GetIsDemo

`func (o *GetCommentsResponsePublicComment) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *GetCommentsResponsePublicComment) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *GetCommentsResponsePublicComment) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *GetCommentsResponsePublicComment) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetCommentCount

`func (o *GetCommentsResponsePublicComment) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetCommentsResponsePublicComment) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetCommentsResponsePublicComment) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *GetCommentsResponsePublicComment) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetIsSiteAdmin

`func (o *GetCommentsResponsePublicComment) GetIsSiteAdmin() bool`

GetIsSiteAdmin returns the IsSiteAdmin field if non-nil, zero value otherwise.

### GetIsSiteAdminOk

`func (o *GetCommentsResponsePublicComment) GetIsSiteAdminOk() (*bool, bool)`

GetIsSiteAdminOk returns a tuple with the IsSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSiteAdmin

`func (o *GetCommentsResponsePublicComment) SetIsSiteAdmin(v bool)`

SetIsSiteAdmin sets IsSiteAdmin field to given value.

### HasIsSiteAdmin

`func (o *GetCommentsResponsePublicComment) HasIsSiteAdmin() bool`

HasIsSiteAdmin returns a boolean if a field has been set.

### GetHasBillingIssue

`func (o *GetCommentsResponsePublicComment) GetHasBillingIssue() bool`

GetHasBillingIssue returns the HasBillingIssue field if non-nil, zero value otherwise.

### GetHasBillingIssueOk

`func (o *GetCommentsResponsePublicComment) GetHasBillingIssueOk() (*bool, bool)`

GetHasBillingIssueOk returns a tuple with the HasBillingIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBillingIssue

`func (o *GetCommentsResponsePublicComment) SetHasBillingIssue(v bool)`

SetHasBillingIssue sets HasBillingIssue field to given value.

### HasHasBillingIssue

`func (o *GetCommentsResponsePublicComment) HasHasBillingIssue() bool`

HasHasBillingIssue returns a boolean if a field has been set.

### GetModuleData

`func (o *GetCommentsResponsePublicComment) GetModuleData() map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *GetCommentsResponsePublicComment) GetModuleDataOk() (*map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *GetCommentsResponsePublicComment) SetModuleData(v map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *GetCommentsResponsePublicComment) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetPageNumber

`func (o *GetCommentsResponsePublicComment) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *GetCommentsResponsePublicComment) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *GetCommentsResponsePublicComment) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetIsWhiteLabeled

`func (o *GetCommentsResponsePublicComment) GetIsWhiteLabeled() bool`

GetIsWhiteLabeled returns the IsWhiteLabeled field if non-nil, zero value otherwise.

### GetIsWhiteLabeledOk

`func (o *GetCommentsResponsePublicComment) GetIsWhiteLabeledOk() (*bool, bool)`

GetIsWhiteLabeledOk returns a tuple with the IsWhiteLabeled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteLabeled

`func (o *GetCommentsResponsePublicComment) SetIsWhiteLabeled(v bool)`

SetIsWhiteLabeled sets IsWhiteLabeled field to given value.

### HasIsWhiteLabeled

`func (o *GetCommentsResponsePublicComment) HasIsWhiteLabeled() bool`

HasIsWhiteLabeled returns a boolean if a field has been set.

### GetIsProd

`func (o *GetCommentsResponsePublicComment) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *GetCommentsResponsePublicComment) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *GetCommentsResponsePublicComment) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *GetCommentsResponsePublicComment) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetIsCrawler

`func (o *GetCommentsResponsePublicComment) GetIsCrawler() bool`

GetIsCrawler returns the IsCrawler field if non-nil, zero value otherwise.

### GetIsCrawlerOk

`func (o *GetCommentsResponsePublicComment) GetIsCrawlerOk() (*bool, bool)`

GetIsCrawlerOk returns a tuple with the IsCrawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCrawler

`func (o *GetCommentsResponsePublicComment) SetIsCrawler(v bool)`

SetIsCrawler sets IsCrawler field to given value.

### HasIsCrawler

`func (o *GetCommentsResponsePublicComment) HasIsCrawler() bool`

HasIsCrawler returns a boolean if a field has been set.

### GetNotificationCount

`func (o *GetCommentsResponsePublicComment) GetNotificationCount() int32`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *GetCommentsResponsePublicComment) GetNotificationCountOk() (*int32, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *GetCommentsResponsePublicComment) SetNotificationCount(v int32)`

SetNotificationCount sets NotificationCount field to given value.

### HasNotificationCount

`func (o *GetCommentsResponsePublicComment) HasNotificationCount() bool`

HasNotificationCount returns a boolean if a field has been set.

### GetHasMore

`func (o *GetCommentsResponsePublicComment) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCommentsResponsePublicComment) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCommentsResponsePublicComment) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetCommentsResponsePublicComment) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetIsClosed

`func (o *GetCommentsResponsePublicComment) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *GetCommentsResponsePublicComment) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *GetCommentsResponsePublicComment) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *GetCommentsResponsePublicComment) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetPresencePollState

`func (o *GetCommentsResponsePublicComment) GetPresencePollState() int32`

GetPresencePollState returns the PresencePollState field if non-nil, zero value otherwise.

### GetPresencePollStateOk

`func (o *GetCommentsResponsePublicComment) GetPresencePollStateOk() (*int32, bool)`

GetPresencePollStateOk returns a tuple with the PresencePollState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePollState

`func (o *GetCommentsResponsePublicComment) SetPresencePollState(v int32)`

SetPresencePollState sets PresencePollState field to given value.

### HasPresencePollState

`func (o *GetCommentsResponsePublicComment) HasPresencePollState() bool`

HasPresencePollState returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetCommentsResponsePublicComment) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetCommentsResponsePublicComment) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetCommentsResponsePublicComment) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetCommentsResponsePublicComment) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


