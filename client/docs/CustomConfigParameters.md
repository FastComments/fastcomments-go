# CustomConfigParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteAndRelativeDates** | Pointer to **bool** |  | [optional] 
**AbsoluteDates** | Pointer to **bool** |  | [optional] 
**AllowAnon** | Pointer to **bool** |  | [optional] 
**AllowAnonFlag** | Pointer to **bool** |  | [optional] 
**AllowAnonVotes** | Pointer to **bool** |  | [optional] 
**AllowedLanguages** | Pointer to **[]string** |  | [optional] 
**CollapseReplies** | Pointer to **bool** |  | [optional] 
**CommentCountFormat** | Pointer to **NullableString** |  | [optional] 
**CommentHTMLRenderingMode** | Pointer to [**CommentHTMLRenderingMode**](CommentHTMLRenderingMode.md) |  | [optional] 
**CommentThreadDeleteMode** | Pointer to [**NullableCommentThreadDeletionMode**](CommentThreadDeletionMode.md) |  | [optional] 
**CommenterNameFormat** | Pointer to [**NullableCommenterNameFormats**](CommenterNameFormats.md) |  | [optional] 
**CountAboveToggle** | Pointer to **int32** |  | [optional] 
**CustomCSS** | Pointer to **NullableString** |  | [optional] 
**DefaultAvatarSrc** | Pointer to **NullableString** |  | [optional] 
**DefaultSortDirection** | Pointer to [**NullableSortDirections**](SortDirections.md) |  | [optional] 
**DefaultUsername** | Pointer to **NullableString** |  | [optional] 
**DisableAutoAdminMigration** | Pointer to **bool** |  | [optional] 
**DisableAutoHashTagCreation** | Pointer to **bool** |  | [optional] 
**DisableBlocking** | Pointer to **bool** |  | [optional] 
**DisableCommenterCommentDelete** | Pointer to **bool** |  | [optional] 
**DisableCommenterCommentEdit** | Pointer to **bool** |  | [optional] 
**DisableEmailInputs** | Pointer to **bool** |  | [optional] 
**DisableLiveCommenting** | Pointer to **bool** |  | [optional] 
**DisableNotificationBell** | Pointer to **bool** |  | [optional] 
**DisableProfiles** | Pointer to **bool** |  | [optional] 
**DisableSuccessMessage** | Pointer to **bool** |  | [optional] 
**DisableToolbar** | Pointer to **bool** |  | [optional] 
**DisableUnverifiedLabel** | Pointer to **bool** |  | [optional] 
**DisableVoting** | Pointer to **bool** |  | [optional] 
**EnableCommenterLinks** | Pointer to **bool** |  | [optional] 
**EnableSearch** | Pointer to **bool** |  | [optional] 
**EnableSpoilers** | Pointer to **bool** |  | [optional] 
**EnableThirdPartyCookieBypass** | Pointer to **bool** |  | [optional] 
**EnableViewCounts** | Pointer to **bool** |  | [optional] 
**EnableVoteList** | Pointer to **bool** |  | [optional] 
**EnableWYSIWYG** | Pointer to **bool** |  | [optional] 
**GifRating** | Pointer to [**GifRating**](GifRating.md) |  | [optional] 
**HasDarkBackground** | Pointer to **bool** |  | [optional] 
**HeaderHTML** | Pointer to **NullableString** |  | [optional] 
**HideAvatars** | Pointer to **bool** |  | [optional] 
**HideCommentsUnderCountTextFormat** | Pointer to **NullableString** |  | [optional] 
**ImageContentProfanityLevel** | Pointer to [**ImageContentProfanityLevel**](ImageContentProfanityLevel.md) |  | [optional] 
**InputAfterComments** | Pointer to **bool** |  | [optional] 
**LimitCommentsByGroups** | Pointer to **bool** |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**MaxCommentCharacterLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxCommentCreatedCountPUPM** | Pointer to **NullableInt32** |  | [optional] 
**NoCustomConfig** | Pointer to **bool** |  | [optional] 
**NoImageUploads** | Pointer to **bool** |  | [optional] 
**NoStyles** | Pointer to **bool** |  | [optional] 
**PageSize** | Pointer to **NullableInt32** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**NoNewRootComments** | Pointer to **bool** |  | [optional] 
**RequireSSO** | Pointer to **bool** |  | [optional] 
**EnableResizeHandle** | Pointer to **bool** |  | [optional] 
**RestrictedLinkDomains** | Pointer to **[]string** |  | [optional] 
**ShowBadgesInTopBar** | Pointer to **bool** |  | [optional] 
**ShowCommentSaveSuccess** | Pointer to **bool** |  | [optional] 
**ShowLiveRightAway** | Pointer to **bool** |  | [optional] 
**ShowQuestion** | Pointer to **bool** |  | [optional] 
**SpamRules** | Pointer to [**[]SpamRule**](SpamRule.md) |  | [optional] 
**SsoSecLvl** | Pointer to [**SSOSecurityLevel**](SSOSecurityLevel.md) |  | [optional] 
**Translations** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 
**UseShowCommentsToggle** | Pointer to **bool** |  | [optional] 
**UseSingleLineCommentInput** | Pointer to **bool** |  | [optional] 
**VoteStyle** | Pointer to [**VoteStyle**](VoteStyle.md) |  | [optional] 
**WidgetQuestionId** | Pointer to **string** |  | [optional] 
**WidgetQuestionResultsStyle** | Pointer to [**CommentQuestionResultsRenderingType**](CommentQuestionResultsRenderingType.md) |  | [optional] 
**WidgetQuestionStyle** | Pointer to [**QuestionRenderingType**](QuestionRenderingType.md) |  | [optional] 
**WidgetQuestionWhenToSave** | Pointer to [**QuestionWhenSave**](QuestionWhenSave.md) |  | [optional] 
**WidgetQuestionsRequired** | Pointer to [**CommentQuestionsRequired**](CommentQuestionsRequired.md) |  | [optional] 
**WidgetSubQuestionVisibility** | Pointer to [**QuestionSubQuestionVisibility**](QuestionSubQuestionVisibility.md) |  | [optional] 
**Wrap** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomConfigParameters

`func NewCustomConfigParameters() *CustomConfigParameters`

NewCustomConfigParameters instantiates a new CustomConfigParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConfigParametersWithDefaults

`func NewCustomConfigParametersWithDefaults() *CustomConfigParameters`

NewCustomConfigParametersWithDefaults instantiates a new CustomConfigParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteAndRelativeDates

`func (o *CustomConfigParameters) GetAbsoluteAndRelativeDates() bool`

GetAbsoluteAndRelativeDates returns the AbsoluteAndRelativeDates field if non-nil, zero value otherwise.

### GetAbsoluteAndRelativeDatesOk

`func (o *CustomConfigParameters) GetAbsoluteAndRelativeDatesOk() (*bool, bool)`

GetAbsoluteAndRelativeDatesOk returns a tuple with the AbsoluteAndRelativeDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteAndRelativeDates

`func (o *CustomConfigParameters) SetAbsoluteAndRelativeDates(v bool)`

SetAbsoluteAndRelativeDates sets AbsoluteAndRelativeDates field to given value.

### HasAbsoluteAndRelativeDates

`func (o *CustomConfigParameters) HasAbsoluteAndRelativeDates() bool`

HasAbsoluteAndRelativeDates returns a boolean if a field has been set.

### GetAbsoluteDates

`func (o *CustomConfigParameters) GetAbsoluteDates() bool`

GetAbsoluteDates returns the AbsoluteDates field if non-nil, zero value otherwise.

### GetAbsoluteDatesOk

`func (o *CustomConfigParameters) GetAbsoluteDatesOk() (*bool, bool)`

GetAbsoluteDatesOk returns a tuple with the AbsoluteDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteDates

`func (o *CustomConfigParameters) SetAbsoluteDates(v bool)`

SetAbsoluteDates sets AbsoluteDates field to given value.

### HasAbsoluteDates

`func (o *CustomConfigParameters) HasAbsoluteDates() bool`

HasAbsoluteDates returns a boolean if a field has been set.

### GetAllowAnon

`func (o *CustomConfigParameters) GetAllowAnon() bool`

GetAllowAnon returns the AllowAnon field if non-nil, zero value otherwise.

### GetAllowAnonOk

`func (o *CustomConfigParameters) GetAllowAnonOk() (*bool, bool)`

GetAllowAnonOk returns a tuple with the AllowAnon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnon

`func (o *CustomConfigParameters) SetAllowAnon(v bool)`

SetAllowAnon sets AllowAnon field to given value.

### HasAllowAnon

`func (o *CustomConfigParameters) HasAllowAnon() bool`

HasAllowAnon returns a boolean if a field has been set.

### GetAllowAnonFlag

`func (o *CustomConfigParameters) GetAllowAnonFlag() bool`

GetAllowAnonFlag returns the AllowAnonFlag field if non-nil, zero value otherwise.

### GetAllowAnonFlagOk

`func (o *CustomConfigParameters) GetAllowAnonFlagOk() (*bool, bool)`

GetAllowAnonFlagOk returns a tuple with the AllowAnonFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonFlag

`func (o *CustomConfigParameters) SetAllowAnonFlag(v bool)`

SetAllowAnonFlag sets AllowAnonFlag field to given value.

### HasAllowAnonFlag

`func (o *CustomConfigParameters) HasAllowAnonFlag() bool`

HasAllowAnonFlag returns a boolean if a field has been set.

### GetAllowAnonVotes

`func (o *CustomConfigParameters) GetAllowAnonVotes() bool`

GetAllowAnonVotes returns the AllowAnonVotes field if non-nil, zero value otherwise.

### GetAllowAnonVotesOk

`func (o *CustomConfigParameters) GetAllowAnonVotesOk() (*bool, bool)`

GetAllowAnonVotesOk returns a tuple with the AllowAnonVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonVotes

`func (o *CustomConfigParameters) SetAllowAnonVotes(v bool)`

SetAllowAnonVotes sets AllowAnonVotes field to given value.

### HasAllowAnonVotes

`func (o *CustomConfigParameters) HasAllowAnonVotes() bool`

HasAllowAnonVotes returns a boolean if a field has been set.

### GetAllowedLanguages

`func (o *CustomConfigParameters) GetAllowedLanguages() []string`

GetAllowedLanguages returns the AllowedLanguages field if non-nil, zero value otherwise.

### GetAllowedLanguagesOk

`func (o *CustomConfigParameters) GetAllowedLanguagesOk() (*[]string, bool)`

GetAllowedLanguagesOk returns a tuple with the AllowedLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedLanguages

`func (o *CustomConfigParameters) SetAllowedLanguages(v []string)`

SetAllowedLanguages sets AllowedLanguages field to given value.

### HasAllowedLanguages

`func (o *CustomConfigParameters) HasAllowedLanguages() bool`

HasAllowedLanguages returns a boolean if a field has been set.

### SetAllowedLanguagesNil

`func (o *CustomConfigParameters) SetAllowedLanguagesNil(b bool)`

 SetAllowedLanguagesNil sets the value for AllowedLanguages to be an explicit nil

### UnsetAllowedLanguages
`func (o *CustomConfigParameters) UnsetAllowedLanguages()`

UnsetAllowedLanguages ensures that no value is present for AllowedLanguages, not even an explicit nil
### GetCollapseReplies

`func (o *CustomConfigParameters) GetCollapseReplies() bool`

GetCollapseReplies returns the CollapseReplies field if non-nil, zero value otherwise.

### GetCollapseRepliesOk

`func (o *CustomConfigParameters) GetCollapseRepliesOk() (*bool, bool)`

GetCollapseRepliesOk returns a tuple with the CollapseReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseReplies

`func (o *CustomConfigParameters) SetCollapseReplies(v bool)`

SetCollapseReplies sets CollapseReplies field to given value.

### HasCollapseReplies

`func (o *CustomConfigParameters) HasCollapseReplies() bool`

HasCollapseReplies returns a boolean if a field has been set.

### GetCommentCountFormat

`func (o *CustomConfigParameters) GetCommentCountFormat() string`

GetCommentCountFormat returns the CommentCountFormat field if non-nil, zero value otherwise.

### GetCommentCountFormatOk

`func (o *CustomConfigParameters) GetCommentCountFormatOk() (*string, bool)`

GetCommentCountFormatOk returns a tuple with the CommentCountFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCountFormat

`func (o *CustomConfigParameters) SetCommentCountFormat(v string)`

SetCommentCountFormat sets CommentCountFormat field to given value.

### HasCommentCountFormat

`func (o *CustomConfigParameters) HasCommentCountFormat() bool`

HasCommentCountFormat returns a boolean if a field has been set.

### SetCommentCountFormatNil

`func (o *CustomConfigParameters) SetCommentCountFormatNil(b bool)`

 SetCommentCountFormatNil sets the value for CommentCountFormat to be an explicit nil

### UnsetCommentCountFormat
`func (o *CustomConfigParameters) UnsetCommentCountFormat()`

UnsetCommentCountFormat ensures that no value is present for CommentCountFormat, not even an explicit nil
### GetCommentHTMLRenderingMode

`func (o *CustomConfigParameters) GetCommentHTMLRenderingMode() CommentHTMLRenderingMode`

GetCommentHTMLRenderingMode returns the CommentHTMLRenderingMode field if non-nil, zero value otherwise.

### GetCommentHTMLRenderingModeOk

`func (o *CustomConfigParameters) GetCommentHTMLRenderingModeOk() (*CommentHTMLRenderingMode, bool)`

GetCommentHTMLRenderingModeOk returns a tuple with the CommentHTMLRenderingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentHTMLRenderingMode

`func (o *CustomConfigParameters) SetCommentHTMLRenderingMode(v CommentHTMLRenderingMode)`

SetCommentHTMLRenderingMode sets CommentHTMLRenderingMode field to given value.

### HasCommentHTMLRenderingMode

`func (o *CustomConfigParameters) HasCommentHTMLRenderingMode() bool`

HasCommentHTMLRenderingMode returns a boolean if a field has been set.

### GetCommentThreadDeleteMode

`func (o *CustomConfigParameters) GetCommentThreadDeleteMode() CommentThreadDeletionMode`

GetCommentThreadDeleteMode returns the CommentThreadDeleteMode field if non-nil, zero value otherwise.

### GetCommentThreadDeleteModeOk

`func (o *CustomConfigParameters) GetCommentThreadDeleteModeOk() (*CommentThreadDeletionMode, bool)`

GetCommentThreadDeleteModeOk returns a tuple with the CommentThreadDeleteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentThreadDeleteMode

`func (o *CustomConfigParameters) SetCommentThreadDeleteMode(v CommentThreadDeletionMode)`

SetCommentThreadDeleteMode sets CommentThreadDeleteMode field to given value.

### HasCommentThreadDeleteMode

`func (o *CustomConfigParameters) HasCommentThreadDeleteMode() bool`

HasCommentThreadDeleteMode returns a boolean if a field has been set.

### SetCommentThreadDeleteModeNil

`func (o *CustomConfigParameters) SetCommentThreadDeleteModeNil(b bool)`

 SetCommentThreadDeleteModeNil sets the value for CommentThreadDeleteMode to be an explicit nil

### UnsetCommentThreadDeleteMode
`func (o *CustomConfigParameters) UnsetCommentThreadDeleteMode()`

UnsetCommentThreadDeleteMode ensures that no value is present for CommentThreadDeleteMode, not even an explicit nil
### GetCommenterNameFormat

`func (o *CustomConfigParameters) GetCommenterNameFormat() CommenterNameFormats`

GetCommenterNameFormat returns the CommenterNameFormat field if non-nil, zero value otherwise.

### GetCommenterNameFormatOk

`func (o *CustomConfigParameters) GetCommenterNameFormatOk() (*CommenterNameFormats, bool)`

GetCommenterNameFormatOk returns a tuple with the CommenterNameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterNameFormat

`func (o *CustomConfigParameters) SetCommenterNameFormat(v CommenterNameFormats)`

SetCommenterNameFormat sets CommenterNameFormat field to given value.

### HasCommenterNameFormat

`func (o *CustomConfigParameters) HasCommenterNameFormat() bool`

HasCommenterNameFormat returns a boolean if a field has been set.

### SetCommenterNameFormatNil

`func (o *CustomConfigParameters) SetCommenterNameFormatNil(b bool)`

 SetCommenterNameFormatNil sets the value for CommenterNameFormat to be an explicit nil

### UnsetCommenterNameFormat
`func (o *CustomConfigParameters) UnsetCommenterNameFormat()`

UnsetCommenterNameFormat ensures that no value is present for CommenterNameFormat, not even an explicit nil
### GetCountAboveToggle

`func (o *CustomConfigParameters) GetCountAboveToggle() int32`

GetCountAboveToggle returns the CountAboveToggle field if non-nil, zero value otherwise.

### GetCountAboveToggleOk

`func (o *CustomConfigParameters) GetCountAboveToggleOk() (*int32, bool)`

GetCountAboveToggleOk returns a tuple with the CountAboveToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAboveToggle

`func (o *CustomConfigParameters) SetCountAboveToggle(v int32)`

SetCountAboveToggle sets CountAboveToggle field to given value.

### HasCountAboveToggle

`func (o *CustomConfigParameters) HasCountAboveToggle() bool`

HasCountAboveToggle returns a boolean if a field has been set.

### GetCustomCSS

`func (o *CustomConfigParameters) GetCustomCSS() string`

GetCustomCSS returns the CustomCSS field if non-nil, zero value otherwise.

### GetCustomCSSOk

`func (o *CustomConfigParameters) GetCustomCSSOk() (*string, bool)`

GetCustomCSSOk returns a tuple with the CustomCSS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCSS

`func (o *CustomConfigParameters) SetCustomCSS(v string)`

SetCustomCSS sets CustomCSS field to given value.

### HasCustomCSS

`func (o *CustomConfigParameters) HasCustomCSS() bool`

HasCustomCSS returns a boolean if a field has been set.

### SetCustomCSSNil

`func (o *CustomConfigParameters) SetCustomCSSNil(b bool)`

 SetCustomCSSNil sets the value for CustomCSS to be an explicit nil

### UnsetCustomCSS
`func (o *CustomConfigParameters) UnsetCustomCSS()`

UnsetCustomCSS ensures that no value is present for CustomCSS, not even an explicit nil
### GetDefaultAvatarSrc

`func (o *CustomConfigParameters) GetDefaultAvatarSrc() string`

GetDefaultAvatarSrc returns the DefaultAvatarSrc field if non-nil, zero value otherwise.

### GetDefaultAvatarSrcOk

`func (o *CustomConfigParameters) GetDefaultAvatarSrcOk() (*string, bool)`

GetDefaultAvatarSrcOk returns a tuple with the DefaultAvatarSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatarSrc

`func (o *CustomConfigParameters) SetDefaultAvatarSrc(v string)`

SetDefaultAvatarSrc sets DefaultAvatarSrc field to given value.

### HasDefaultAvatarSrc

`func (o *CustomConfigParameters) HasDefaultAvatarSrc() bool`

HasDefaultAvatarSrc returns a boolean if a field has been set.

### SetDefaultAvatarSrcNil

`func (o *CustomConfigParameters) SetDefaultAvatarSrcNil(b bool)`

 SetDefaultAvatarSrcNil sets the value for DefaultAvatarSrc to be an explicit nil

### UnsetDefaultAvatarSrc
`func (o *CustomConfigParameters) UnsetDefaultAvatarSrc()`

UnsetDefaultAvatarSrc ensures that no value is present for DefaultAvatarSrc, not even an explicit nil
### GetDefaultSortDirection

`func (o *CustomConfigParameters) GetDefaultSortDirection() SortDirections`

GetDefaultSortDirection returns the DefaultSortDirection field if non-nil, zero value otherwise.

### GetDefaultSortDirectionOk

`func (o *CustomConfigParameters) GetDefaultSortDirectionOk() (*SortDirections, bool)`

GetDefaultSortDirectionOk returns a tuple with the DefaultSortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortDirection

`func (o *CustomConfigParameters) SetDefaultSortDirection(v SortDirections)`

SetDefaultSortDirection sets DefaultSortDirection field to given value.

### HasDefaultSortDirection

`func (o *CustomConfigParameters) HasDefaultSortDirection() bool`

HasDefaultSortDirection returns a boolean if a field has been set.

### SetDefaultSortDirectionNil

`func (o *CustomConfigParameters) SetDefaultSortDirectionNil(b bool)`

 SetDefaultSortDirectionNil sets the value for DefaultSortDirection to be an explicit nil

### UnsetDefaultSortDirection
`func (o *CustomConfigParameters) UnsetDefaultSortDirection()`

UnsetDefaultSortDirection ensures that no value is present for DefaultSortDirection, not even an explicit nil
### GetDefaultUsername

`func (o *CustomConfigParameters) GetDefaultUsername() string`

GetDefaultUsername returns the DefaultUsername field if non-nil, zero value otherwise.

### GetDefaultUsernameOk

`func (o *CustomConfigParameters) GetDefaultUsernameOk() (*string, bool)`

GetDefaultUsernameOk returns a tuple with the DefaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsername

`func (o *CustomConfigParameters) SetDefaultUsername(v string)`

SetDefaultUsername sets DefaultUsername field to given value.

### HasDefaultUsername

`func (o *CustomConfigParameters) HasDefaultUsername() bool`

HasDefaultUsername returns a boolean if a field has been set.

### SetDefaultUsernameNil

`func (o *CustomConfigParameters) SetDefaultUsernameNil(b bool)`

 SetDefaultUsernameNil sets the value for DefaultUsername to be an explicit nil

### UnsetDefaultUsername
`func (o *CustomConfigParameters) UnsetDefaultUsername()`

UnsetDefaultUsername ensures that no value is present for DefaultUsername, not even an explicit nil
### GetDisableAutoAdminMigration

`func (o *CustomConfigParameters) GetDisableAutoAdminMigration() bool`

GetDisableAutoAdminMigration returns the DisableAutoAdminMigration field if non-nil, zero value otherwise.

### GetDisableAutoAdminMigrationOk

`func (o *CustomConfigParameters) GetDisableAutoAdminMigrationOk() (*bool, bool)`

GetDisableAutoAdminMigrationOk returns a tuple with the DisableAutoAdminMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoAdminMigration

`func (o *CustomConfigParameters) SetDisableAutoAdminMigration(v bool)`

SetDisableAutoAdminMigration sets DisableAutoAdminMigration field to given value.

### HasDisableAutoAdminMigration

`func (o *CustomConfigParameters) HasDisableAutoAdminMigration() bool`

HasDisableAutoAdminMigration returns a boolean if a field has been set.

### GetDisableAutoHashTagCreation

`func (o *CustomConfigParameters) GetDisableAutoHashTagCreation() bool`

GetDisableAutoHashTagCreation returns the DisableAutoHashTagCreation field if non-nil, zero value otherwise.

### GetDisableAutoHashTagCreationOk

`func (o *CustomConfigParameters) GetDisableAutoHashTagCreationOk() (*bool, bool)`

GetDisableAutoHashTagCreationOk returns a tuple with the DisableAutoHashTagCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoHashTagCreation

`func (o *CustomConfigParameters) SetDisableAutoHashTagCreation(v bool)`

SetDisableAutoHashTagCreation sets DisableAutoHashTagCreation field to given value.

### HasDisableAutoHashTagCreation

`func (o *CustomConfigParameters) HasDisableAutoHashTagCreation() bool`

HasDisableAutoHashTagCreation returns a boolean if a field has been set.

### GetDisableBlocking

`func (o *CustomConfigParameters) GetDisableBlocking() bool`

GetDisableBlocking returns the DisableBlocking field if non-nil, zero value otherwise.

### GetDisableBlockingOk

`func (o *CustomConfigParameters) GetDisableBlockingOk() (*bool, bool)`

GetDisableBlockingOk returns a tuple with the DisableBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBlocking

`func (o *CustomConfigParameters) SetDisableBlocking(v bool)`

SetDisableBlocking sets DisableBlocking field to given value.

### HasDisableBlocking

`func (o *CustomConfigParameters) HasDisableBlocking() bool`

HasDisableBlocking returns a boolean if a field has been set.

### GetDisableCommenterCommentDelete

`func (o *CustomConfigParameters) GetDisableCommenterCommentDelete() bool`

GetDisableCommenterCommentDelete returns the DisableCommenterCommentDelete field if non-nil, zero value otherwise.

### GetDisableCommenterCommentDeleteOk

`func (o *CustomConfigParameters) GetDisableCommenterCommentDeleteOk() (*bool, bool)`

GetDisableCommenterCommentDeleteOk returns a tuple with the DisableCommenterCommentDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommenterCommentDelete

`func (o *CustomConfigParameters) SetDisableCommenterCommentDelete(v bool)`

SetDisableCommenterCommentDelete sets DisableCommenterCommentDelete field to given value.

### HasDisableCommenterCommentDelete

`func (o *CustomConfigParameters) HasDisableCommenterCommentDelete() bool`

HasDisableCommenterCommentDelete returns a boolean if a field has been set.

### GetDisableCommenterCommentEdit

`func (o *CustomConfigParameters) GetDisableCommenterCommentEdit() bool`

GetDisableCommenterCommentEdit returns the DisableCommenterCommentEdit field if non-nil, zero value otherwise.

### GetDisableCommenterCommentEditOk

`func (o *CustomConfigParameters) GetDisableCommenterCommentEditOk() (*bool, bool)`

GetDisableCommenterCommentEditOk returns a tuple with the DisableCommenterCommentEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommenterCommentEdit

`func (o *CustomConfigParameters) SetDisableCommenterCommentEdit(v bool)`

SetDisableCommenterCommentEdit sets DisableCommenterCommentEdit field to given value.

### HasDisableCommenterCommentEdit

`func (o *CustomConfigParameters) HasDisableCommenterCommentEdit() bool`

HasDisableCommenterCommentEdit returns a boolean if a field has been set.

### GetDisableEmailInputs

`func (o *CustomConfigParameters) GetDisableEmailInputs() bool`

GetDisableEmailInputs returns the DisableEmailInputs field if non-nil, zero value otherwise.

### GetDisableEmailInputsOk

`func (o *CustomConfigParameters) GetDisableEmailInputsOk() (*bool, bool)`

GetDisableEmailInputsOk returns a tuple with the DisableEmailInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailInputs

`func (o *CustomConfigParameters) SetDisableEmailInputs(v bool)`

SetDisableEmailInputs sets DisableEmailInputs field to given value.

### HasDisableEmailInputs

`func (o *CustomConfigParameters) HasDisableEmailInputs() bool`

HasDisableEmailInputs returns a boolean if a field has been set.

### GetDisableLiveCommenting

`func (o *CustomConfigParameters) GetDisableLiveCommenting() bool`

GetDisableLiveCommenting returns the DisableLiveCommenting field if non-nil, zero value otherwise.

### GetDisableLiveCommentingOk

`func (o *CustomConfigParameters) GetDisableLiveCommentingOk() (*bool, bool)`

GetDisableLiveCommentingOk returns a tuple with the DisableLiveCommenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLiveCommenting

`func (o *CustomConfigParameters) SetDisableLiveCommenting(v bool)`

SetDisableLiveCommenting sets DisableLiveCommenting field to given value.

### HasDisableLiveCommenting

`func (o *CustomConfigParameters) HasDisableLiveCommenting() bool`

HasDisableLiveCommenting returns a boolean if a field has been set.

### GetDisableNotificationBell

`func (o *CustomConfigParameters) GetDisableNotificationBell() bool`

GetDisableNotificationBell returns the DisableNotificationBell field if non-nil, zero value otherwise.

### GetDisableNotificationBellOk

`func (o *CustomConfigParameters) GetDisableNotificationBellOk() (*bool, bool)`

GetDisableNotificationBellOk returns a tuple with the DisableNotificationBell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotificationBell

`func (o *CustomConfigParameters) SetDisableNotificationBell(v bool)`

SetDisableNotificationBell sets DisableNotificationBell field to given value.

### HasDisableNotificationBell

`func (o *CustomConfigParameters) HasDisableNotificationBell() bool`

HasDisableNotificationBell returns a boolean if a field has been set.

### GetDisableProfiles

`func (o *CustomConfigParameters) GetDisableProfiles() bool`

GetDisableProfiles returns the DisableProfiles field if non-nil, zero value otherwise.

### GetDisableProfilesOk

`func (o *CustomConfigParameters) GetDisableProfilesOk() (*bool, bool)`

GetDisableProfilesOk returns a tuple with the DisableProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableProfiles

`func (o *CustomConfigParameters) SetDisableProfiles(v bool)`

SetDisableProfiles sets DisableProfiles field to given value.

### HasDisableProfiles

`func (o *CustomConfigParameters) HasDisableProfiles() bool`

HasDisableProfiles returns a boolean if a field has been set.

### GetDisableSuccessMessage

`func (o *CustomConfigParameters) GetDisableSuccessMessage() bool`

GetDisableSuccessMessage returns the DisableSuccessMessage field if non-nil, zero value otherwise.

### GetDisableSuccessMessageOk

`func (o *CustomConfigParameters) GetDisableSuccessMessageOk() (*bool, bool)`

GetDisableSuccessMessageOk returns a tuple with the DisableSuccessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSuccessMessage

`func (o *CustomConfigParameters) SetDisableSuccessMessage(v bool)`

SetDisableSuccessMessage sets DisableSuccessMessage field to given value.

### HasDisableSuccessMessage

`func (o *CustomConfigParameters) HasDisableSuccessMessage() bool`

HasDisableSuccessMessage returns a boolean if a field has been set.

### GetDisableToolbar

`func (o *CustomConfigParameters) GetDisableToolbar() bool`

GetDisableToolbar returns the DisableToolbar field if non-nil, zero value otherwise.

### GetDisableToolbarOk

`func (o *CustomConfigParameters) GetDisableToolbarOk() (*bool, bool)`

GetDisableToolbarOk returns a tuple with the DisableToolbar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableToolbar

`func (o *CustomConfigParameters) SetDisableToolbar(v bool)`

SetDisableToolbar sets DisableToolbar field to given value.

### HasDisableToolbar

`func (o *CustomConfigParameters) HasDisableToolbar() bool`

HasDisableToolbar returns a boolean if a field has been set.

### GetDisableUnverifiedLabel

`func (o *CustomConfigParameters) GetDisableUnverifiedLabel() bool`

GetDisableUnverifiedLabel returns the DisableUnverifiedLabel field if non-nil, zero value otherwise.

### GetDisableUnverifiedLabelOk

`func (o *CustomConfigParameters) GetDisableUnverifiedLabelOk() (*bool, bool)`

GetDisableUnverifiedLabelOk returns a tuple with the DisableUnverifiedLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUnverifiedLabel

`func (o *CustomConfigParameters) SetDisableUnverifiedLabel(v bool)`

SetDisableUnverifiedLabel sets DisableUnverifiedLabel field to given value.

### HasDisableUnverifiedLabel

`func (o *CustomConfigParameters) HasDisableUnverifiedLabel() bool`

HasDisableUnverifiedLabel returns a boolean if a field has been set.

### GetDisableVoting

`func (o *CustomConfigParameters) GetDisableVoting() bool`

GetDisableVoting returns the DisableVoting field if non-nil, zero value otherwise.

### GetDisableVotingOk

`func (o *CustomConfigParameters) GetDisableVotingOk() (*bool, bool)`

GetDisableVotingOk returns a tuple with the DisableVoting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableVoting

`func (o *CustomConfigParameters) SetDisableVoting(v bool)`

SetDisableVoting sets DisableVoting field to given value.

### HasDisableVoting

`func (o *CustomConfigParameters) HasDisableVoting() bool`

HasDisableVoting returns a boolean if a field has been set.

### GetEnableCommenterLinks

`func (o *CustomConfigParameters) GetEnableCommenterLinks() bool`

GetEnableCommenterLinks returns the EnableCommenterLinks field if non-nil, zero value otherwise.

### GetEnableCommenterLinksOk

`func (o *CustomConfigParameters) GetEnableCommenterLinksOk() (*bool, bool)`

GetEnableCommenterLinksOk returns a tuple with the EnableCommenterLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCommenterLinks

`func (o *CustomConfigParameters) SetEnableCommenterLinks(v bool)`

SetEnableCommenterLinks sets EnableCommenterLinks field to given value.

### HasEnableCommenterLinks

`func (o *CustomConfigParameters) HasEnableCommenterLinks() bool`

HasEnableCommenterLinks returns a boolean if a field has been set.

### GetEnableSearch

`func (o *CustomConfigParameters) GetEnableSearch() bool`

GetEnableSearch returns the EnableSearch field if non-nil, zero value otherwise.

### GetEnableSearchOk

`func (o *CustomConfigParameters) GetEnableSearchOk() (*bool, bool)`

GetEnableSearchOk returns a tuple with the EnableSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSearch

`func (o *CustomConfigParameters) SetEnableSearch(v bool)`

SetEnableSearch sets EnableSearch field to given value.

### HasEnableSearch

`func (o *CustomConfigParameters) HasEnableSearch() bool`

HasEnableSearch returns a boolean if a field has been set.

### GetEnableSpoilers

`func (o *CustomConfigParameters) GetEnableSpoilers() bool`

GetEnableSpoilers returns the EnableSpoilers field if non-nil, zero value otherwise.

### GetEnableSpoilersOk

`func (o *CustomConfigParameters) GetEnableSpoilersOk() (*bool, bool)`

GetEnableSpoilersOk returns a tuple with the EnableSpoilers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpoilers

`func (o *CustomConfigParameters) SetEnableSpoilers(v bool)`

SetEnableSpoilers sets EnableSpoilers field to given value.

### HasEnableSpoilers

`func (o *CustomConfigParameters) HasEnableSpoilers() bool`

HasEnableSpoilers returns a boolean if a field has been set.

### GetEnableThirdPartyCookieBypass

`func (o *CustomConfigParameters) GetEnableThirdPartyCookieBypass() bool`

GetEnableThirdPartyCookieBypass returns the EnableThirdPartyCookieBypass field if non-nil, zero value otherwise.

### GetEnableThirdPartyCookieBypassOk

`func (o *CustomConfigParameters) GetEnableThirdPartyCookieBypassOk() (*bool, bool)`

GetEnableThirdPartyCookieBypassOk returns a tuple with the EnableThirdPartyCookieBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThirdPartyCookieBypass

`func (o *CustomConfigParameters) SetEnableThirdPartyCookieBypass(v bool)`

SetEnableThirdPartyCookieBypass sets EnableThirdPartyCookieBypass field to given value.

### HasEnableThirdPartyCookieBypass

`func (o *CustomConfigParameters) HasEnableThirdPartyCookieBypass() bool`

HasEnableThirdPartyCookieBypass returns a boolean if a field has been set.

### GetEnableViewCounts

`func (o *CustomConfigParameters) GetEnableViewCounts() bool`

GetEnableViewCounts returns the EnableViewCounts field if non-nil, zero value otherwise.

### GetEnableViewCountsOk

`func (o *CustomConfigParameters) GetEnableViewCountsOk() (*bool, bool)`

GetEnableViewCountsOk returns a tuple with the EnableViewCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableViewCounts

`func (o *CustomConfigParameters) SetEnableViewCounts(v bool)`

SetEnableViewCounts sets EnableViewCounts field to given value.

### HasEnableViewCounts

`func (o *CustomConfigParameters) HasEnableViewCounts() bool`

HasEnableViewCounts returns a boolean if a field has been set.

### GetEnableVoteList

`func (o *CustomConfigParameters) GetEnableVoteList() bool`

GetEnableVoteList returns the EnableVoteList field if non-nil, zero value otherwise.

### GetEnableVoteListOk

`func (o *CustomConfigParameters) GetEnableVoteListOk() (*bool, bool)`

GetEnableVoteListOk returns a tuple with the EnableVoteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVoteList

`func (o *CustomConfigParameters) SetEnableVoteList(v bool)`

SetEnableVoteList sets EnableVoteList field to given value.

### HasEnableVoteList

`func (o *CustomConfigParameters) HasEnableVoteList() bool`

HasEnableVoteList returns a boolean if a field has been set.

### GetEnableWYSIWYG

`func (o *CustomConfigParameters) GetEnableWYSIWYG() bool`

GetEnableWYSIWYG returns the EnableWYSIWYG field if non-nil, zero value otherwise.

### GetEnableWYSIWYGOk

`func (o *CustomConfigParameters) GetEnableWYSIWYGOk() (*bool, bool)`

GetEnableWYSIWYGOk returns a tuple with the EnableWYSIWYG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWYSIWYG

`func (o *CustomConfigParameters) SetEnableWYSIWYG(v bool)`

SetEnableWYSIWYG sets EnableWYSIWYG field to given value.

### HasEnableWYSIWYG

`func (o *CustomConfigParameters) HasEnableWYSIWYG() bool`

HasEnableWYSIWYG returns a boolean if a field has been set.

### GetGifRating

`func (o *CustomConfigParameters) GetGifRating() GifRating`

GetGifRating returns the GifRating field if non-nil, zero value otherwise.

### GetGifRatingOk

`func (o *CustomConfigParameters) GetGifRatingOk() (*GifRating, bool)`

GetGifRatingOk returns a tuple with the GifRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifRating

`func (o *CustomConfigParameters) SetGifRating(v GifRating)`

SetGifRating sets GifRating field to given value.

### HasGifRating

`func (o *CustomConfigParameters) HasGifRating() bool`

HasGifRating returns a boolean if a field has been set.

### GetHasDarkBackground

`func (o *CustomConfigParameters) GetHasDarkBackground() bool`

GetHasDarkBackground returns the HasDarkBackground field if non-nil, zero value otherwise.

### GetHasDarkBackgroundOk

`func (o *CustomConfigParameters) GetHasDarkBackgroundOk() (*bool, bool)`

GetHasDarkBackgroundOk returns a tuple with the HasDarkBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDarkBackground

`func (o *CustomConfigParameters) SetHasDarkBackground(v bool)`

SetHasDarkBackground sets HasDarkBackground field to given value.

### HasHasDarkBackground

`func (o *CustomConfigParameters) HasHasDarkBackground() bool`

HasHasDarkBackground returns a boolean if a field has been set.

### GetHeaderHTML

`func (o *CustomConfigParameters) GetHeaderHTML() string`

GetHeaderHTML returns the HeaderHTML field if non-nil, zero value otherwise.

### GetHeaderHTMLOk

`func (o *CustomConfigParameters) GetHeaderHTMLOk() (*string, bool)`

GetHeaderHTMLOk returns a tuple with the HeaderHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHTML

`func (o *CustomConfigParameters) SetHeaderHTML(v string)`

SetHeaderHTML sets HeaderHTML field to given value.

### HasHeaderHTML

`func (o *CustomConfigParameters) HasHeaderHTML() bool`

HasHeaderHTML returns a boolean if a field has been set.

### SetHeaderHTMLNil

`func (o *CustomConfigParameters) SetHeaderHTMLNil(b bool)`

 SetHeaderHTMLNil sets the value for HeaderHTML to be an explicit nil

### UnsetHeaderHTML
`func (o *CustomConfigParameters) UnsetHeaderHTML()`

UnsetHeaderHTML ensures that no value is present for HeaderHTML, not even an explicit nil
### GetHideAvatars

`func (o *CustomConfigParameters) GetHideAvatars() bool`

GetHideAvatars returns the HideAvatars field if non-nil, zero value otherwise.

### GetHideAvatarsOk

`func (o *CustomConfigParameters) GetHideAvatarsOk() (*bool, bool)`

GetHideAvatarsOk returns a tuple with the HideAvatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAvatars

`func (o *CustomConfigParameters) SetHideAvatars(v bool)`

SetHideAvatars sets HideAvatars field to given value.

### HasHideAvatars

`func (o *CustomConfigParameters) HasHideAvatars() bool`

HasHideAvatars returns a boolean if a field has been set.

### GetHideCommentsUnderCountTextFormat

`func (o *CustomConfigParameters) GetHideCommentsUnderCountTextFormat() string`

GetHideCommentsUnderCountTextFormat returns the HideCommentsUnderCountTextFormat field if non-nil, zero value otherwise.

### GetHideCommentsUnderCountTextFormatOk

`func (o *CustomConfigParameters) GetHideCommentsUnderCountTextFormatOk() (*string, bool)`

GetHideCommentsUnderCountTextFormatOk returns a tuple with the HideCommentsUnderCountTextFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideCommentsUnderCountTextFormat

`func (o *CustomConfigParameters) SetHideCommentsUnderCountTextFormat(v string)`

SetHideCommentsUnderCountTextFormat sets HideCommentsUnderCountTextFormat field to given value.

### HasHideCommentsUnderCountTextFormat

`func (o *CustomConfigParameters) HasHideCommentsUnderCountTextFormat() bool`

HasHideCommentsUnderCountTextFormat returns a boolean if a field has been set.

### SetHideCommentsUnderCountTextFormatNil

`func (o *CustomConfigParameters) SetHideCommentsUnderCountTextFormatNil(b bool)`

 SetHideCommentsUnderCountTextFormatNil sets the value for HideCommentsUnderCountTextFormat to be an explicit nil

### UnsetHideCommentsUnderCountTextFormat
`func (o *CustomConfigParameters) UnsetHideCommentsUnderCountTextFormat()`

UnsetHideCommentsUnderCountTextFormat ensures that no value is present for HideCommentsUnderCountTextFormat, not even an explicit nil
### GetImageContentProfanityLevel

`func (o *CustomConfigParameters) GetImageContentProfanityLevel() ImageContentProfanityLevel`

GetImageContentProfanityLevel returns the ImageContentProfanityLevel field if non-nil, zero value otherwise.

### GetImageContentProfanityLevelOk

`func (o *CustomConfigParameters) GetImageContentProfanityLevelOk() (*ImageContentProfanityLevel, bool)`

GetImageContentProfanityLevelOk returns a tuple with the ImageContentProfanityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContentProfanityLevel

`func (o *CustomConfigParameters) SetImageContentProfanityLevel(v ImageContentProfanityLevel)`

SetImageContentProfanityLevel sets ImageContentProfanityLevel field to given value.

### HasImageContentProfanityLevel

`func (o *CustomConfigParameters) HasImageContentProfanityLevel() bool`

HasImageContentProfanityLevel returns a boolean if a field has been set.

### GetInputAfterComments

`func (o *CustomConfigParameters) GetInputAfterComments() bool`

GetInputAfterComments returns the InputAfterComments field if non-nil, zero value otherwise.

### GetInputAfterCommentsOk

`func (o *CustomConfigParameters) GetInputAfterCommentsOk() (*bool, bool)`

GetInputAfterCommentsOk returns a tuple with the InputAfterComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAfterComments

`func (o *CustomConfigParameters) SetInputAfterComments(v bool)`

SetInputAfterComments sets InputAfterComments field to given value.

### HasInputAfterComments

`func (o *CustomConfigParameters) HasInputAfterComments() bool`

HasInputAfterComments returns a boolean if a field has been set.

### GetLimitCommentsByGroups

`func (o *CustomConfigParameters) GetLimitCommentsByGroups() bool`

GetLimitCommentsByGroups returns the LimitCommentsByGroups field if non-nil, zero value otherwise.

### GetLimitCommentsByGroupsOk

`func (o *CustomConfigParameters) GetLimitCommentsByGroupsOk() (*bool, bool)`

GetLimitCommentsByGroupsOk returns a tuple with the LimitCommentsByGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCommentsByGroups

`func (o *CustomConfigParameters) SetLimitCommentsByGroups(v bool)`

SetLimitCommentsByGroups sets LimitCommentsByGroups field to given value.

### HasLimitCommentsByGroups

`func (o *CustomConfigParameters) HasLimitCommentsByGroups() bool`

HasLimitCommentsByGroups returns a boolean if a field has been set.

### GetLocale

`func (o *CustomConfigParameters) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CustomConfigParameters) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CustomConfigParameters) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CustomConfigParameters) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *CustomConfigParameters) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *CustomConfigParameters) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetMaxCommentCharacterLength

`func (o *CustomConfigParameters) GetMaxCommentCharacterLength() int32`

GetMaxCommentCharacterLength returns the MaxCommentCharacterLength field if non-nil, zero value otherwise.

### GetMaxCommentCharacterLengthOk

`func (o *CustomConfigParameters) GetMaxCommentCharacterLengthOk() (*int32, bool)`

GetMaxCommentCharacterLengthOk returns a tuple with the MaxCommentCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCommentCharacterLength

`func (o *CustomConfigParameters) SetMaxCommentCharacterLength(v int32)`

SetMaxCommentCharacterLength sets MaxCommentCharacterLength field to given value.

### HasMaxCommentCharacterLength

`func (o *CustomConfigParameters) HasMaxCommentCharacterLength() bool`

HasMaxCommentCharacterLength returns a boolean if a field has been set.

### SetMaxCommentCharacterLengthNil

`func (o *CustomConfigParameters) SetMaxCommentCharacterLengthNil(b bool)`

 SetMaxCommentCharacterLengthNil sets the value for MaxCommentCharacterLength to be an explicit nil

### UnsetMaxCommentCharacterLength
`func (o *CustomConfigParameters) UnsetMaxCommentCharacterLength()`

UnsetMaxCommentCharacterLength ensures that no value is present for MaxCommentCharacterLength, not even an explicit nil
### GetMaxCommentCreatedCountPUPM

`func (o *CustomConfigParameters) GetMaxCommentCreatedCountPUPM() int32`

GetMaxCommentCreatedCountPUPM returns the MaxCommentCreatedCountPUPM field if non-nil, zero value otherwise.

### GetMaxCommentCreatedCountPUPMOk

`func (o *CustomConfigParameters) GetMaxCommentCreatedCountPUPMOk() (*int32, bool)`

GetMaxCommentCreatedCountPUPMOk returns a tuple with the MaxCommentCreatedCountPUPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCommentCreatedCountPUPM

`func (o *CustomConfigParameters) SetMaxCommentCreatedCountPUPM(v int32)`

SetMaxCommentCreatedCountPUPM sets MaxCommentCreatedCountPUPM field to given value.

### HasMaxCommentCreatedCountPUPM

`func (o *CustomConfigParameters) HasMaxCommentCreatedCountPUPM() bool`

HasMaxCommentCreatedCountPUPM returns a boolean if a field has been set.

### SetMaxCommentCreatedCountPUPMNil

`func (o *CustomConfigParameters) SetMaxCommentCreatedCountPUPMNil(b bool)`

 SetMaxCommentCreatedCountPUPMNil sets the value for MaxCommentCreatedCountPUPM to be an explicit nil

### UnsetMaxCommentCreatedCountPUPM
`func (o *CustomConfigParameters) UnsetMaxCommentCreatedCountPUPM()`

UnsetMaxCommentCreatedCountPUPM ensures that no value is present for MaxCommentCreatedCountPUPM, not even an explicit nil
### GetNoCustomConfig

`func (o *CustomConfigParameters) GetNoCustomConfig() bool`

GetNoCustomConfig returns the NoCustomConfig field if non-nil, zero value otherwise.

### GetNoCustomConfigOk

`func (o *CustomConfigParameters) GetNoCustomConfigOk() (*bool, bool)`

GetNoCustomConfigOk returns a tuple with the NoCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCustomConfig

`func (o *CustomConfigParameters) SetNoCustomConfig(v bool)`

SetNoCustomConfig sets NoCustomConfig field to given value.

### HasNoCustomConfig

`func (o *CustomConfigParameters) HasNoCustomConfig() bool`

HasNoCustomConfig returns a boolean if a field has been set.

### GetNoImageUploads

`func (o *CustomConfigParameters) GetNoImageUploads() bool`

GetNoImageUploads returns the NoImageUploads field if non-nil, zero value otherwise.

### GetNoImageUploadsOk

`func (o *CustomConfigParameters) GetNoImageUploadsOk() (*bool, bool)`

GetNoImageUploadsOk returns a tuple with the NoImageUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoImageUploads

`func (o *CustomConfigParameters) SetNoImageUploads(v bool)`

SetNoImageUploads sets NoImageUploads field to given value.

### HasNoImageUploads

`func (o *CustomConfigParameters) HasNoImageUploads() bool`

HasNoImageUploads returns a boolean if a field has been set.

### GetNoStyles

`func (o *CustomConfigParameters) GetNoStyles() bool`

GetNoStyles returns the NoStyles field if non-nil, zero value otherwise.

### GetNoStylesOk

`func (o *CustomConfigParameters) GetNoStylesOk() (*bool, bool)`

GetNoStylesOk returns a tuple with the NoStyles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoStyles

`func (o *CustomConfigParameters) SetNoStyles(v bool)`

SetNoStyles sets NoStyles field to given value.

### HasNoStyles

`func (o *CustomConfigParameters) HasNoStyles() bool`

HasNoStyles returns a boolean if a field has been set.

### GetPageSize

`func (o *CustomConfigParameters) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CustomConfigParameters) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CustomConfigParameters) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CustomConfigParameters) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *CustomConfigParameters) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *CustomConfigParameters) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetReadonly

`func (o *CustomConfigParameters) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *CustomConfigParameters) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *CustomConfigParameters) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *CustomConfigParameters) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetNoNewRootComments

`func (o *CustomConfigParameters) GetNoNewRootComments() bool`

GetNoNewRootComments returns the NoNewRootComments field if non-nil, zero value otherwise.

### GetNoNewRootCommentsOk

`func (o *CustomConfigParameters) GetNoNewRootCommentsOk() (*bool, bool)`

GetNoNewRootCommentsOk returns a tuple with the NoNewRootComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoNewRootComments

`func (o *CustomConfigParameters) SetNoNewRootComments(v bool)`

SetNoNewRootComments sets NoNewRootComments field to given value.

### HasNoNewRootComments

`func (o *CustomConfigParameters) HasNoNewRootComments() bool`

HasNoNewRootComments returns a boolean if a field has been set.

### GetRequireSSO

`func (o *CustomConfigParameters) GetRequireSSO() bool`

GetRequireSSO returns the RequireSSO field if non-nil, zero value otherwise.

### GetRequireSSOOk

`func (o *CustomConfigParameters) GetRequireSSOOk() (*bool, bool)`

GetRequireSSOOk returns a tuple with the RequireSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSSO

`func (o *CustomConfigParameters) SetRequireSSO(v bool)`

SetRequireSSO sets RequireSSO field to given value.

### HasRequireSSO

`func (o *CustomConfigParameters) HasRequireSSO() bool`

HasRequireSSO returns a boolean if a field has been set.

### GetEnableResizeHandle

`func (o *CustomConfigParameters) GetEnableResizeHandle() bool`

GetEnableResizeHandle returns the EnableResizeHandle field if non-nil, zero value otherwise.

### GetEnableResizeHandleOk

`func (o *CustomConfigParameters) GetEnableResizeHandleOk() (*bool, bool)`

GetEnableResizeHandleOk returns a tuple with the EnableResizeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableResizeHandle

`func (o *CustomConfigParameters) SetEnableResizeHandle(v bool)`

SetEnableResizeHandle sets EnableResizeHandle field to given value.

### HasEnableResizeHandle

`func (o *CustomConfigParameters) HasEnableResizeHandle() bool`

HasEnableResizeHandle returns a boolean if a field has been set.

### GetRestrictedLinkDomains

`func (o *CustomConfigParameters) GetRestrictedLinkDomains() []string`

GetRestrictedLinkDomains returns the RestrictedLinkDomains field if non-nil, zero value otherwise.

### GetRestrictedLinkDomainsOk

`func (o *CustomConfigParameters) GetRestrictedLinkDomainsOk() (*[]string, bool)`

GetRestrictedLinkDomainsOk returns a tuple with the RestrictedLinkDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedLinkDomains

`func (o *CustomConfigParameters) SetRestrictedLinkDomains(v []string)`

SetRestrictedLinkDomains sets RestrictedLinkDomains field to given value.

### HasRestrictedLinkDomains

`func (o *CustomConfigParameters) HasRestrictedLinkDomains() bool`

HasRestrictedLinkDomains returns a boolean if a field has been set.

### SetRestrictedLinkDomainsNil

`func (o *CustomConfigParameters) SetRestrictedLinkDomainsNil(b bool)`

 SetRestrictedLinkDomainsNil sets the value for RestrictedLinkDomains to be an explicit nil

### UnsetRestrictedLinkDomains
`func (o *CustomConfigParameters) UnsetRestrictedLinkDomains()`

UnsetRestrictedLinkDomains ensures that no value is present for RestrictedLinkDomains, not even an explicit nil
### GetShowBadgesInTopBar

`func (o *CustomConfigParameters) GetShowBadgesInTopBar() bool`

GetShowBadgesInTopBar returns the ShowBadgesInTopBar field if non-nil, zero value otherwise.

### GetShowBadgesInTopBarOk

`func (o *CustomConfigParameters) GetShowBadgesInTopBarOk() (*bool, bool)`

GetShowBadgesInTopBarOk returns a tuple with the ShowBadgesInTopBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBadgesInTopBar

`func (o *CustomConfigParameters) SetShowBadgesInTopBar(v bool)`

SetShowBadgesInTopBar sets ShowBadgesInTopBar field to given value.

### HasShowBadgesInTopBar

`func (o *CustomConfigParameters) HasShowBadgesInTopBar() bool`

HasShowBadgesInTopBar returns a boolean if a field has been set.

### GetShowCommentSaveSuccess

`func (o *CustomConfigParameters) GetShowCommentSaveSuccess() bool`

GetShowCommentSaveSuccess returns the ShowCommentSaveSuccess field if non-nil, zero value otherwise.

### GetShowCommentSaveSuccessOk

`func (o *CustomConfigParameters) GetShowCommentSaveSuccessOk() (*bool, bool)`

GetShowCommentSaveSuccessOk returns a tuple with the ShowCommentSaveSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCommentSaveSuccess

`func (o *CustomConfigParameters) SetShowCommentSaveSuccess(v bool)`

SetShowCommentSaveSuccess sets ShowCommentSaveSuccess field to given value.

### HasShowCommentSaveSuccess

`func (o *CustomConfigParameters) HasShowCommentSaveSuccess() bool`

HasShowCommentSaveSuccess returns a boolean if a field has been set.

### GetShowLiveRightAway

`func (o *CustomConfigParameters) GetShowLiveRightAway() bool`

GetShowLiveRightAway returns the ShowLiveRightAway field if non-nil, zero value otherwise.

### GetShowLiveRightAwayOk

`func (o *CustomConfigParameters) GetShowLiveRightAwayOk() (*bool, bool)`

GetShowLiveRightAwayOk returns a tuple with the ShowLiveRightAway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLiveRightAway

`func (o *CustomConfigParameters) SetShowLiveRightAway(v bool)`

SetShowLiveRightAway sets ShowLiveRightAway field to given value.

### HasShowLiveRightAway

`func (o *CustomConfigParameters) HasShowLiveRightAway() bool`

HasShowLiveRightAway returns a boolean if a field has been set.

### GetShowQuestion

`func (o *CustomConfigParameters) GetShowQuestion() bool`

GetShowQuestion returns the ShowQuestion field if non-nil, zero value otherwise.

### GetShowQuestionOk

`func (o *CustomConfigParameters) GetShowQuestionOk() (*bool, bool)`

GetShowQuestionOk returns a tuple with the ShowQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowQuestion

`func (o *CustomConfigParameters) SetShowQuestion(v bool)`

SetShowQuestion sets ShowQuestion field to given value.

### HasShowQuestion

`func (o *CustomConfigParameters) HasShowQuestion() bool`

HasShowQuestion returns a boolean if a field has been set.

### GetSpamRules

`func (o *CustomConfigParameters) GetSpamRules() []SpamRule`

GetSpamRules returns the SpamRules field if non-nil, zero value otherwise.

### GetSpamRulesOk

`func (o *CustomConfigParameters) GetSpamRulesOk() (*[]SpamRule, bool)`

GetSpamRulesOk returns a tuple with the SpamRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamRules

`func (o *CustomConfigParameters) SetSpamRules(v []SpamRule)`

SetSpamRules sets SpamRules field to given value.

### HasSpamRules

`func (o *CustomConfigParameters) HasSpamRules() bool`

HasSpamRules returns a boolean if a field has been set.

### GetSsoSecLvl

`func (o *CustomConfigParameters) GetSsoSecLvl() SSOSecurityLevel`

GetSsoSecLvl returns the SsoSecLvl field if non-nil, zero value otherwise.

### GetSsoSecLvlOk

`func (o *CustomConfigParameters) GetSsoSecLvlOk() (*SSOSecurityLevel, bool)`

GetSsoSecLvlOk returns a tuple with the SsoSecLvl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSecLvl

`func (o *CustomConfigParameters) SetSsoSecLvl(v SSOSecurityLevel)`

SetSsoSecLvl sets SsoSecLvl field to given value.

### HasSsoSecLvl

`func (o *CustomConfigParameters) HasSsoSecLvl() bool`

HasSsoSecLvl returns a boolean if a field has been set.

### GetTranslations

`func (o *CustomConfigParameters) GetTranslations() map[string]string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *CustomConfigParameters) GetTranslationsOk() (*map[string]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *CustomConfigParameters) SetTranslations(v map[string]string)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *CustomConfigParameters) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### SetTranslationsNil

`func (o *CustomConfigParameters) SetTranslationsNil(b bool)`

 SetTranslationsNil sets the value for Translations to be an explicit nil

### UnsetTranslations
`func (o *CustomConfigParameters) UnsetTranslations()`

UnsetTranslations ensures that no value is present for Translations, not even an explicit nil
### GetUseShowCommentsToggle

`func (o *CustomConfigParameters) GetUseShowCommentsToggle() bool`

GetUseShowCommentsToggle returns the UseShowCommentsToggle field if non-nil, zero value otherwise.

### GetUseShowCommentsToggleOk

`func (o *CustomConfigParameters) GetUseShowCommentsToggleOk() (*bool, bool)`

GetUseShowCommentsToggleOk returns a tuple with the UseShowCommentsToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseShowCommentsToggle

`func (o *CustomConfigParameters) SetUseShowCommentsToggle(v bool)`

SetUseShowCommentsToggle sets UseShowCommentsToggle field to given value.

### HasUseShowCommentsToggle

`func (o *CustomConfigParameters) HasUseShowCommentsToggle() bool`

HasUseShowCommentsToggle returns a boolean if a field has been set.

### GetUseSingleLineCommentInput

`func (o *CustomConfigParameters) GetUseSingleLineCommentInput() bool`

GetUseSingleLineCommentInput returns the UseSingleLineCommentInput field if non-nil, zero value otherwise.

### GetUseSingleLineCommentInputOk

`func (o *CustomConfigParameters) GetUseSingleLineCommentInputOk() (*bool, bool)`

GetUseSingleLineCommentInputOk returns a tuple with the UseSingleLineCommentInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSingleLineCommentInput

`func (o *CustomConfigParameters) SetUseSingleLineCommentInput(v bool)`

SetUseSingleLineCommentInput sets UseSingleLineCommentInput field to given value.

### HasUseSingleLineCommentInput

`func (o *CustomConfigParameters) HasUseSingleLineCommentInput() bool`

HasUseSingleLineCommentInput returns a boolean if a field has been set.

### GetVoteStyle

`func (o *CustomConfigParameters) GetVoteStyle() VoteStyle`

GetVoteStyle returns the VoteStyle field if non-nil, zero value otherwise.

### GetVoteStyleOk

`func (o *CustomConfigParameters) GetVoteStyleOk() (*VoteStyle, bool)`

GetVoteStyleOk returns a tuple with the VoteStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteStyle

`func (o *CustomConfigParameters) SetVoteStyle(v VoteStyle)`

SetVoteStyle sets VoteStyle field to given value.

### HasVoteStyle

`func (o *CustomConfigParameters) HasVoteStyle() bool`

HasVoteStyle returns a boolean if a field has been set.

### GetWidgetQuestionId

`func (o *CustomConfigParameters) GetWidgetQuestionId() string`

GetWidgetQuestionId returns the WidgetQuestionId field if non-nil, zero value otherwise.

### GetWidgetQuestionIdOk

`func (o *CustomConfigParameters) GetWidgetQuestionIdOk() (*string, bool)`

GetWidgetQuestionIdOk returns a tuple with the WidgetQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetQuestionId

`func (o *CustomConfigParameters) SetWidgetQuestionId(v string)`

SetWidgetQuestionId sets WidgetQuestionId field to given value.

### HasWidgetQuestionId

`func (o *CustomConfigParameters) HasWidgetQuestionId() bool`

HasWidgetQuestionId returns a boolean if a field has been set.

### GetWidgetQuestionResultsStyle

`func (o *CustomConfigParameters) GetWidgetQuestionResultsStyle() CommentQuestionResultsRenderingType`

GetWidgetQuestionResultsStyle returns the WidgetQuestionResultsStyle field if non-nil, zero value otherwise.

### GetWidgetQuestionResultsStyleOk

`func (o *CustomConfigParameters) GetWidgetQuestionResultsStyleOk() (*CommentQuestionResultsRenderingType, bool)`

GetWidgetQuestionResultsStyleOk returns a tuple with the WidgetQuestionResultsStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetQuestionResultsStyle

`func (o *CustomConfigParameters) SetWidgetQuestionResultsStyle(v CommentQuestionResultsRenderingType)`

SetWidgetQuestionResultsStyle sets WidgetQuestionResultsStyle field to given value.

### HasWidgetQuestionResultsStyle

`func (o *CustomConfigParameters) HasWidgetQuestionResultsStyle() bool`

HasWidgetQuestionResultsStyle returns a boolean if a field has been set.

### GetWidgetQuestionStyle

`func (o *CustomConfigParameters) GetWidgetQuestionStyle() QuestionRenderingType`

GetWidgetQuestionStyle returns the WidgetQuestionStyle field if non-nil, zero value otherwise.

### GetWidgetQuestionStyleOk

`func (o *CustomConfigParameters) GetWidgetQuestionStyleOk() (*QuestionRenderingType, bool)`

GetWidgetQuestionStyleOk returns a tuple with the WidgetQuestionStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetQuestionStyle

`func (o *CustomConfigParameters) SetWidgetQuestionStyle(v QuestionRenderingType)`

SetWidgetQuestionStyle sets WidgetQuestionStyle field to given value.

### HasWidgetQuestionStyle

`func (o *CustomConfigParameters) HasWidgetQuestionStyle() bool`

HasWidgetQuestionStyle returns a boolean if a field has been set.

### GetWidgetQuestionWhenToSave

`func (o *CustomConfigParameters) GetWidgetQuestionWhenToSave() QuestionWhenSave`

GetWidgetQuestionWhenToSave returns the WidgetQuestionWhenToSave field if non-nil, zero value otherwise.

### GetWidgetQuestionWhenToSaveOk

`func (o *CustomConfigParameters) GetWidgetQuestionWhenToSaveOk() (*QuestionWhenSave, bool)`

GetWidgetQuestionWhenToSaveOk returns a tuple with the WidgetQuestionWhenToSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetQuestionWhenToSave

`func (o *CustomConfigParameters) SetWidgetQuestionWhenToSave(v QuestionWhenSave)`

SetWidgetQuestionWhenToSave sets WidgetQuestionWhenToSave field to given value.

### HasWidgetQuestionWhenToSave

`func (o *CustomConfigParameters) HasWidgetQuestionWhenToSave() bool`

HasWidgetQuestionWhenToSave returns a boolean if a field has been set.

### GetWidgetQuestionsRequired

`func (o *CustomConfigParameters) GetWidgetQuestionsRequired() CommentQuestionsRequired`

GetWidgetQuestionsRequired returns the WidgetQuestionsRequired field if non-nil, zero value otherwise.

### GetWidgetQuestionsRequiredOk

`func (o *CustomConfigParameters) GetWidgetQuestionsRequiredOk() (*CommentQuestionsRequired, bool)`

GetWidgetQuestionsRequiredOk returns a tuple with the WidgetQuestionsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetQuestionsRequired

`func (o *CustomConfigParameters) SetWidgetQuestionsRequired(v CommentQuestionsRequired)`

SetWidgetQuestionsRequired sets WidgetQuestionsRequired field to given value.

### HasWidgetQuestionsRequired

`func (o *CustomConfigParameters) HasWidgetQuestionsRequired() bool`

HasWidgetQuestionsRequired returns a boolean if a field has been set.

### GetWidgetSubQuestionVisibility

`func (o *CustomConfigParameters) GetWidgetSubQuestionVisibility() QuestionSubQuestionVisibility`

GetWidgetSubQuestionVisibility returns the WidgetSubQuestionVisibility field if non-nil, zero value otherwise.

### GetWidgetSubQuestionVisibilityOk

`func (o *CustomConfigParameters) GetWidgetSubQuestionVisibilityOk() (*QuestionSubQuestionVisibility, bool)`

GetWidgetSubQuestionVisibilityOk returns a tuple with the WidgetSubQuestionVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetSubQuestionVisibility

`func (o *CustomConfigParameters) SetWidgetSubQuestionVisibility(v QuestionSubQuestionVisibility)`

SetWidgetSubQuestionVisibility sets WidgetSubQuestionVisibility field to given value.

### HasWidgetSubQuestionVisibility

`func (o *CustomConfigParameters) HasWidgetSubQuestionVisibility() bool`

HasWidgetSubQuestionVisibility returns a boolean if a field has been set.

### GetWrap

`func (o *CustomConfigParameters) GetWrap() bool`

GetWrap returns the Wrap field if non-nil, zero value otherwise.

### GetWrapOk

`func (o *CustomConfigParameters) GetWrapOk() (*bool, bool)`

GetWrapOk returns a tuple with the Wrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrap

`func (o *CustomConfigParameters) SetWrap(v bool)`

SetWrap sets Wrap field to given value.

### HasWrap

`func (o *CustomConfigParameters) HasWrap() bool`

HasWrap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


