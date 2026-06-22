# GetCommentsPublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Code** | **string** |  | 
**Reason** | **string** |  | 
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
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCommentsPublicResponse

`func NewGetCommentsPublicResponse(status APIStatus, code string, reason string, comments []PublicComment, user NullableUserSessionInfo, pageNumber int32, ) *GetCommentsPublicResponse`

NewGetCommentsPublicResponse instantiates a new GetCommentsPublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentsPublicResponseWithDefaults

`func NewGetCommentsPublicResponseWithDefaults() *GetCommentsPublicResponse`

NewGetCommentsPublicResponseWithDefaults instantiates a new GetCommentsPublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetCommentsPublicResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetCommentsPublicResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetCommentsPublicResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetCommentsPublicResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetCommentsPublicResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentsPublicResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentsPublicResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCode

`func (o *GetCommentsPublicResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCommentsPublicResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCommentsPublicResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetReason

`func (o *GetCommentsPublicResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetCommentsPublicResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetCommentsPublicResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTranslatedWarning

`func (o *GetCommentsPublicResponse) GetTranslatedWarning() string`

GetTranslatedWarning returns the TranslatedWarning field if non-nil, zero value otherwise.

### GetTranslatedWarningOk

`func (o *GetCommentsPublicResponse) GetTranslatedWarningOk() (*string, bool)`

GetTranslatedWarningOk returns a tuple with the TranslatedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedWarning

`func (o *GetCommentsPublicResponse) SetTranslatedWarning(v string)`

SetTranslatedWarning sets TranslatedWarning field to given value.

### HasTranslatedWarning

`func (o *GetCommentsPublicResponse) HasTranslatedWarning() bool`

HasTranslatedWarning returns a boolean if a field has been set.

### GetComments

`func (o *GetCommentsPublicResponse) GetComments() []PublicComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetCommentsPublicResponse) GetCommentsOk() (*[]PublicComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetCommentsPublicResponse) SetComments(v []PublicComment)`

SetComments sets Comments field to given value.


### GetUser

`func (o *GetCommentsPublicResponse) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetCommentsPublicResponse) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetCommentsPublicResponse) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *GetCommentsPublicResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetCommentsPublicResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdClean

`func (o *GetCommentsPublicResponse) GetUrlIdClean() string`

GetUrlIdClean returns the UrlIdClean field if non-nil, zero value otherwise.

### GetUrlIdCleanOk

`func (o *GetCommentsPublicResponse) GetUrlIdCleanOk() (*string, bool)`

GetUrlIdCleanOk returns a tuple with the UrlIdClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdClean

`func (o *GetCommentsPublicResponse) SetUrlIdClean(v string)`

SetUrlIdClean sets UrlIdClean field to given value.

### HasUrlIdClean

`func (o *GetCommentsPublicResponse) HasUrlIdClean() bool`

HasUrlIdClean returns a boolean if a field has been set.

### GetLastGenDate

`func (o *GetCommentsPublicResponse) GetLastGenDate() int64`

GetLastGenDate returns the LastGenDate field if non-nil, zero value otherwise.

### GetLastGenDateOk

`func (o *GetCommentsPublicResponse) GetLastGenDateOk() (*int64, bool)`

GetLastGenDateOk returns a tuple with the LastGenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGenDate

`func (o *GetCommentsPublicResponse) SetLastGenDate(v int64)`

SetLastGenDate sets LastGenDate field to given value.

### HasLastGenDate

`func (o *GetCommentsPublicResponse) HasLastGenDate() bool`

HasLastGenDate returns a boolean if a field has been set.

### SetLastGenDateNil

`func (o *GetCommentsPublicResponse) SetLastGenDateNil(b bool)`

 SetLastGenDateNil sets the value for LastGenDate to be an explicit nil

### UnsetLastGenDate
`func (o *GetCommentsPublicResponse) UnsetLastGenDate()`

UnsetLastGenDate ensures that no value is present for LastGenDate, not even an explicit nil
### GetIncludesPastPages

`func (o *GetCommentsPublicResponse) GetIncludesPastPages() bool`

GetIncludesPastPages returns the IncludesPastPages field if non-nil, zero value otherwise.

### GetIncludesPastPagesOk

`func (o *GetCommentsPublicResponse) GetIncludesPastPagesOk() (*bool, bool)`

GetIncludesPastPagesOk returns a tuple with the IncludesPastPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPastPages

`func (o *GetCommentsPublicResponse) SetIncludesPastPages(v bool)`

SetIncludesPastPages sets IncludesPastPages field to given value.

### HasIncludesPastPages

`func (o *GetCommentsPublicResponse) HasIncludesPastPages() bool`

HasIncludesPastPages returns a boolean if a field has been set.

### GetIsDemo

`func (o *GetCommentsPublicResponse) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *GetCommentsPublicResponse) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *GetCommentsPublicResponse) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *GetCommentsPublicResponse) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetCommentCount

`func (o *GetCommentsPublicResponse) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetCommentsPublicResponse) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetCommentsPublicResponse) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *GetCommentsPublicResponse) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetIsSiteAdmin

`func (o *GetCommentsPublicResponse) GetIsSiteAdmin() bool`

GetIsSiteAdmin returns the IsSiteAdmin field if non-nil, zero value otherwise.

### GetIsSiteAdminOk

`func (o *GetCommentsPublicResponse) GetIsSiteAdminOk() (*bool, bool)`

GetIsSiteAdminOk returns a tuple with the IsSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSiteAdmin

`func (o *GetCommentsPublicResponse) SetIsSiteAdmin(v bool)`

SetIsSiteAdmin sets IsSiteAdmin field to given value.

### HasIsSiteAdmin

`func (o *GetCommentsPublicResponse) HasIsSiteAdmin() bool`

HasIsSiteAdmin returns a boolean if a field has been set.

### GetHasBillingIssue

`func (o *GetCommentsPublicResponse) GetHasBillingIssue() bool`

GetHasBillingIssue returns the HasBillingIssue field if non-nil, zero value otherwise.

### GetHasBillingIssueOk

`func (o *GetCommentsPublicResponse) GetHasBillingIssueOk() (*bool, bool)`

GetHasBillingIssueOk returns a tuple with the HasBillingIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBillingIssue

`func (o *GetCommentsPublicResponse) SetHasBillingIssue(v bool)`

SetHasBillingIssue sets HasBillingIssue field to given value.

### HasHasBillingIssue

`func (o *GetCommentsPublicResponse) HasHasBillingIssue() bool`

HasHasBillingIssue returns a boolean if a field has been set.

### GetModuleData

`func (o *GetCommentsPublicResponse) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *GetCommentsPublicResponse) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *GetCommentsPublicResponse) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *GetCommentsPublicResponse) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetPageNumber

`func (o *GetCommentsPublicResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *GetCommentsPublicResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *GetCommentsPublicResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetIsWhiteLabeled

`func (o *GetCommentsPublicResponse) GetIsWhiteLabeled() bool`

GetIsWhiteLabeled returns the IsWhiteLabeled field if non-nil, zero value otherwise.

### GetIsWhiteLabeledOk

`func (o *GetCommentsPublicResponse) GetIsWhiteLabeledOk() (*bool, bool)`

GetIsWhiteLabeledOk returns a tuple with the IsWhiteLabeled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteLabeled

`func (o *GetCommentsPublicResponse) SetIsWhiteLabeled(v bool)`

SetIsWhiteLabeled sets IsWhiteLabeled field to given value.

### HasIsWhiteLabeled

`func (o *GetCommentsPublicResponse) HasIsWhiteLabeled() bool`

HasIsWhiteLabeled returns a boolean if a field has been set.

### GetIsProd

`func (o *GetCommentsPublicResponse) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *GetCommentsPublicResponse) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *GetCommentsPublicResponse) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *GetCommentsPublicResponse) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetIsCrawler

`func (o *GetCommentsPublicResponse) GetIsCrawler() bool`

GetIsCrawler returns the IsCrawler field if non-nil, zero value otherwise.

### GetIsCrawlerOk

`func (o *GetCommentsPublicResponse) GetIsCrawlerOk() (*bool, bool)`

GetIsCrawlerOk returns a tuple with the IsCrawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCrawler

`func (o *GetCommentsPublicResponse) SetIsCrawler(v bool)`

SetIsCrawler sets IsCrawler field to given value.

### HasIsCrawler

`func (o *GetCommentsPublicResponse) HasIsCrawler() bool`

HasIsCrawler returns a boolean if a field has been set.

### GetNotificationCount

`func (o *GetCommentsPublicResponse) GetNotificationCount() int32`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *GetCommentsPublicResponse) GetNotificationCountOk() (*int32, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *GetCommentsPublicResponse) SetNotificationCount(v int32)`

SetNotificationCount sets NotificationCount field to given value.

### HasNotificationCount

`func (o *GetCommentsPublicResponse) HasNotificationCount() bool`

HasNotificationCount returns a boolean if a field has been set.

### GetHasMore

`func (o *GetCommentsPublicResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCommentsPublicResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCommentsPublicResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetCommentsPublicResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetIsClosed

`func (o *GetCommentsPublicResponse) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *GetCommentsPublicResponse) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *GetCommentsPublicResponse) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *GetCommentsPublicResponse) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetPresencePollState

`func (o *GetCommentsPublicResponse) GetPresencePollState() int32`

GetPresencePollState returns the PresencePollState field if non-nil, zero value otherwise.

### GetPresencePollStateOk

`func (o *GetCommentsPublicResponse) GetPresencePollStateOk() (*int32, bool)`

GetPresencePollStateOk returns a tuple with the PresencePollState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePollState

`func (o *GetCommentsPublicResponse) SetPresencePollState(v int32)`

SetPresencePollState sets PresencePollState field to given value.

### HasPresencePollState

`func (o *GetCommentsPublicResponse) HasPresencePollState() bool`

HasPresencePollState returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetCommentsPublicResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetCommentsPublicResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetCommentsPublicResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetCommentsPublicResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetUrlIdWS

`func (o *GetCommentsPublicResponse) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetCommentsPublicResponse) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetCommentsPublicResponse) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.

### HasUrlIdWS

`func (o *GetCommentsPublicResponse) HasUrlIdWS() bool`

HasUrlIdWS returns a boolean if a field has been set.

### GetUserIdWS

`func (o *GetCommentsPublicResponse) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *GetCommentsPublicResponse) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *GetCommentsPublicResponse) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *GetCommentsPublicResponse) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetTenantIdWS

`func (o *GetCommentsPublicResponse) GetTenantIdWS() string`

GetTenantIdWS returns the TenantIdWS field if non-nil, zero value otherwise.

### GetTenantIdWSOk

`func (o *GetCommentsPublicResponse) GetTenantIdWSOk() (*string, bool)`

GetTenantIdWSOk returns a tuple with the TenantIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdWS

`func (o *GetCommentsPublicResponse) SetTenantIdWS(v string)`

SetTenantIdWS sets TenantIdWS field to given value.

### HasTenantIdWS

`func (o *GetCommentsPublicResponse) HasTenantIdWS() bool`

HasTenantIdWS returns a boolean if a field has been set.

### GetSecondaryCode

`func (o *GetCommentsPublicResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetCommentsPublicResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetCommentsPublicResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetCommentsPublicResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetCommentsPublicResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetCommentsPublicResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetCommentsPublicResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetCommentsPublicResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetCommentsPublicResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetCommentsPublicResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetCommentsPublicResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetCommentsPublicResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetCommentsPublicResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetCommentsPublicResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetCommentsPublicResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetCommentsPublicResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


