# GetCommentsPublic200Response

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

### NewGetCommentsPublic200Response

`func NewGetCommentsPublic200Response(status APIStatus, code string, reason string, comments []PublicComment, user NullableUserSessionInfo, pageNumber int32, ) *GetCommentsPublic200Response`

NewGetCommentsPublic200Response instantiates a new GetCommentsPublic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentsPublic200ResponseWithDefaults

`func NewGetCommentsPublic200ResponseWithDefaults() *GetCommentsPublic200Response`

NewGetCommentsPublic200ResponseWithDefaults instantiates a new GetCommentsPublic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GetCommentsPublic200Response) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetCommentsPublic200Response) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetCommentsPublic200Response) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetCommentsPublic200Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetCommentsPublic200Response) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentsPublic200Response) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentsPublic200Response) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetCode

`func (o *GetCommentsPublic200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCommentsPublic200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCommentsPublic200Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetReason

`func (o *GetCommentsPublic200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetCommentsPublic200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetCommentsPublic200Response) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTranslatedWarning

`func (o *GetCommentsPublic200Response) GetTranslatedWarning() string`

GetTranslatedWarning returns the TranslatedWarning field if non-nil, zero value otherwise.

### GetTranslatedWarningOk

`func (o *GetCommentsPublic200Response) GetTranslatedWarningOk() (*string, bool)`

GetTranslatedWarningOk returns a tuple with the TranslatedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedWarning

`func (o *GetCommentsPublic200Response) SetTranslatedWarning(v string)`

SetTranslatedWarning sets TranslatedWarning field to given value.

### HasTranslatedWarning

`func (o *GetCommentsPublic200Response) HasTranslatedWarning() bool`

HasTranslatedWarning returns a boolean if a field has been set.

### GetComments

`func (o *GetCommentsPublic200Response) GetComments() []PublicComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetCommentsPublic200Response) GetCommentsOk() (*[]PublicComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetCommentsPublic200Response) SetComments(v []PublicComment)`

SetComments sets Comments field to given value.


### GetUser

`func (o *GetCommentsPublic200Response) GetUser() UserSessionInfo`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetCommentsPublic200Response) GetUserOk() (*UserSessionInfo, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetCommentsPublic200Response) SetUser(v UserSessionInfo)`

SetUser sets User field to given value.


### SetUserNil

`func (o *GetCommentsPublic200Response) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetCommentsPublic200Response) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUrlIdClean

`func (o *GetCommentsPublic200Response) GetUrlIdClean() string`

GetUrlIdClean returns the UrlIdClean field if non-nil, zero value otherwise.

### GetUrlIdCleanOk

`func (o *GetCommentsPublic200Response) GetUrlIdCleanOk() (*string, bool)`

GetUrlIdCleanOk returns a tuple with the UrlIdClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdClean

`func (o *GetCommentsPublic200Response) SetUrlIdClean(v string)`

SetUrlIdClean sets UrlIdClean field to given value.

### HasUrlIdClean

`func (o *GetCommentsPublic200Response) HasUrlIdClean() bool`

HasUrlIdClean returns a boolean if a field has been set.

### GetLastGenDate

`func (o *GetCommentsPublic200Response) GetLastGenDate() int64`

GetLastGenDate returns the LastGenDate field if non-nil, zero value otherwise.

### GetLastGenDateOk

`func (o *GetCommentsPublic200Response) GetLastGenDateOk() (*int64, bool)`

GetLastGenDateOk returns a tuple with the LastGenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGenDate

`func (o *GetCommentsPublic200Response) SetLastGenDate(v int64)`

SetLastGenDate sets LastGenDate field to given value.

### HasLastGenDate

`func (o *GetCommentsPublic200Response) HasLastGenDate() bool`

HasLastGenDate returns a boolean if a field has been set.

### SetLastGenDateNil

`func (o *GetCommentsPublic200Response) SetLastGenDateNil(b bool)`

 SetLastGenDateNil sets the value for LastGenDate to be an explicit nil

### UnsetLastGenDate
`func (o *GetCommentsPublic200Response) UnsetLastGenDate()`

UnsetLastGenDate ensures that no value is present for LastGenDate, not even an explicit nil
### GetIncludesPastPages

`func (o *GetCommentsPublic200Response) GetIncludesPastPages() bool`

GetIncludesPastPages returns the IncludesPastPages field if non-nil, zero value otherwise.

### GetIncludesPastPagesOk

`func (o *GetCommentsPublic200Response) GetIncludesPastPagesOk() (*bool, bool)`

GetIncludesPastPagesOk returns a tuple with the IncludesPastPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesPastPages

`func (o *GetCommentsPublic200Response) SetIncludesPastPages(v bool)`

SetIncludesPastPages sets IncludesPastPages field to given value.

### HasIncludesPastPages

`func (o *GetCommentsPublic200Response) HasIncludesPastPages() bool`

HasIncludesPastPages returns a boolean if a field has been set.

### GetIsDemo

`func (o *GetCommentsPublic200Response) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *GetCommentsPublic200Response) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *GetCommentsPublic200Response) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *GetCommentsPublic200Response) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetCommentCount

`func (o *GetCommentsPublic200Response) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetCommentsPublic200Response) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetCommentsPublic200Response) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *GetCommentsPublic200Response) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetIsSiteAdmin

`func (o *GetCommentsPublic200Response) GetIsSiteAdmin() bool`

GetIsSiteAdmin returns the IsSiteAdmin field if non-nil, zero value otherwise.

### GetIsSiteAdminOk

`func (o *GetCommentsPublic200Response) GetIsSiteAdminOk() (*bool, bool)`

GetIsSiteAdminOk returns a tuple with the IsSiteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSiteAdmin

`func (o *GetCommentsPublic200Response) SetIsSiteAdmin(v bool)`

SetIsSiteAdmin sets IsSiteAdmin field to given value.

### HasIsSiteAdmin

`func (o *GetCommentsPublic200Response) HasIsSiteAdmin() bool`

HasIsSiteAdmin returns a boolean if a field has been set.

### GetHasBillingIssue

`func (o *GetCommentsPublic200Response) GetHasBillingIssue() bool`

GetHasBillingIssue returns the HasBillingIssue field if non-nil, zero value otherwise.

### GetHasBillingIssueOk

`func (o *GetCommentsPublic200Response) GetHasBillingIssueOk() (*bool, bool)`

GetHasBillingIssueOk returns a tuple with the HasBillingIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBillingIssue

`func (o *GetCommentsPublic200Response) SetHasBillingIssue(v bool)`

SetHasBillingIssue sets HasBillingIssue field to given value.

### HasHasBillingIssue

`func (o *GetCommentsPublic200Response) HasHasBillingIssue() bool`

HasHasBillingIssue returns a boolean if a field has been set.

### GetModuleData

`func (o *GetCommentsPublic200Response) GetModuleData() map[string]map[string]interface{}`

GetModuleData returns the ModuleData field if non-nil, zero value otherwise.

### GetModuleDataOk

`func (o *GetCommentsPublic200Response) GetModuleDataOk() (*map[string]map[string]interface{}, bool)`

GetModuleDataOk returns a tuple with the ModuleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleData

`func (o *GetCommentsPublic200Response) SetModuleData(v map[string]map[string]interface{})`

SetModuleData sets ModuleData field to given value.

### HasModuleData

`func (o *GetCommentsPublic200Response) HasModuleData() bool`

HasModuleData returns a boolean if a field has been set.

### GetPageNumber

`func (o *GetCommentsPublic200Response) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *GetCommentsPublic200Response) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *GetCommentsPublic200Response) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetIsWhiteLabeled

`func (o *GetCommentsPublic200Response) GetIsWhiteLabeled() bool`

GetIsWhiteLabeled returns the IsWhiteLabeled field if non-nil, zero value otherwise.

### GetIsWhiteLabeledOk

`func (o *GetCommentsPublic200Response) GetIsWhiteLabeledOk() (*bool, bool)`

GetIsWhiteLabeledOk returns a tuple with the IsWhiteLabeled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteLabeled

`func (o *GetCommentsPublic200Response) SetIsWhiteLabeled(v bool)`

SetIsWhiteLabeled sets IsWhiteLabeled field to given value.

### HasIsWhiteLabeled

`func (o *GetCommentsPublic200Response) HasIsWhiteLabeled() bool`

HasIsWhiteLabeled returns a boolean if a field has been set.

### GetIsProd

`func (o *GetCommentsPublic200Response) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *GetCommentsPublic200Response) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *GetCommentsPublic200Response) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *GetCommentsPublic200Response) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetIsCrawler

`func (o *GetCommentsPublic200Response) GetIsCrawler() bool`

GetIsCrawler returns the IsCrawler field if non-nil, zero value otherwise.

### GetIsCrawlerOk

`func (o *GetCommentsPublic200Response) GetIsCrawlerOk() (*bool, bool)`

GetIsCrawlerOk returns a tuple with the IsCrawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCrawler

`func (o *GetCommentsPublic200Response) SetIsCrawler(v bool)`

SetIsCrawler sets IsCrawler field to given value.

### HasIsCrawler

`func (o *GetCommentsPublic200Response) HasIsCrawler() bool`

HasIsCrawler returns a boolean if a field has been set.

### GetNotificationCount

`func (o *GetCommentsPublic200Response) GetNotificationCount() int32`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *GetCommentsPublic200Response) GetNotificationCountOk() (*int32, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *GetCommentsPublic200Response) SetNotificationCount(v int32)`

SetNotificationCount sets NotificationCount field to given value.

### HasNotificationCount

`func (o *GetCommentsPublic200Response) HasNotificationCount() bool`

HasNotificationCount returns a boolean if a field has been set.

### GetHasMore

`func (o *GetCommentsPublic200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCommentsPublic200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCommentsPublic200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetCommentsPublic200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetIsClosed

`func (o *GetCommentsPublic200Response) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *GetCommentsPublic200Response) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *GetCommentsPublic200Response) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *GetCommentsPublic200Response) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetPresencePollState

`func (o *GetCommentsPublic200Response) GetPresencePollState() int32`

GetPresencePollState returns the PresencePollState field if non-nil, zero value otherwise.

### GetPresencePollStateOk

`func (o *GetCommentsPublic200Response) GetPresencePollStateOk() (*int32, bool)`

GetPresencePollStateOk returns a tuple with the PresencePollState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePollState

`func (o *GetCommentsPublic200Response) SetPresencePollState(v int32)`

SetPresencePollState sets PresencePollState field to given value.

### HasPresencePollState

`func (o *GetCommentsPublic200Response) HasPresencePollState() bool`

HasPresencePollState returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetCommentsPublic200Response) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetCommentsPublic200Response) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetCommentsPublic200Response) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetCommentsPublic200Response) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetUrlIdWS

`func (o *GetCommentsPublic200Response) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetCommentsPublic200Response) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetCommentsPublic200Response) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.

### HasUrlIdWS

`func (o *GetCommentsPublic200Response) HasUrlIdWS() bool`

HasUrlIdWS returns a boolean if a field has been set.

### GetUserIdWS

`func (o *GetCommentsPublic200Response) GetUserIdWS() string`

GetUserIdWS returns the UserIdWS field if non-nil, zero value otherwise.

### GetUserIdWSOk

`func (o *GetCommentsPublic200Response) GetUserIdWSOk() (*string, bool)`

GetUserIdWSOk returns a tuple with the UserIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdWS

`func (o *GetCommentsPublic200Response) SetUserIdWS(v string)`

SetUserIdWS sets UserIdWS field to given value.

### HasUserIdWS

`func (o *GetCommentsPublic200Response) HasUserIdWS() bool`

HasUserIdWS returns a boolean if a field has been set.

### GetTenantIdWS

`func (o *GetCommentsPublic200Response) GetTenantIdWS() string`

GetTenantIdWS returns the TenantIdWS field if non-nil, zero value otherwise.

### GetTenantIdWSOk

`func (o *GetCommentsPublic200Response) GetTenantIdWSOk() (*string, bool)`

GetTenantIdWSOk returns a tuple with the TenantIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdWS

`func (o *GetCommentsPublic200Response) SetTenantIdWS(v string)`

SetTenantIdWS sets TenantIdWS field to given value.

### HasTenantIdWS

`func (o *GetCommentsPublic200Response) HasTenantIdWS() bool`

HasTenantIdWS returns a boolean if a field has been set.

### GetSecondaryCode

`func (o *GetCommentsPublic200Response) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetCommentsPublic200Response) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetCommentsPublic200Response) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetCommentsPublic200Response) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetCommentsPublic200Response) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetCommentsPublic200Response) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetCommentsPublic200Response) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetCommentsPublic200Response) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetCommentsPublic200Response) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetCommentsPublic200Response) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetCommentsPublic200Response) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetCommentsPublic200Response) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetCommentsPublic200Response) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetCommentsPublic200Response) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetCommentsPublic200Response) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetCommentsPublic200Response) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


