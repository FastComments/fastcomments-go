# PostBulkPreBanSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**TotalRelatedCommentCount** | **int32** |  | 
**EmailDomains** | **[]string** |  | 
**Emails** | **[]string** |  | 
**UserIds** | **[]string** |  | 
**IpHashes** | **[]string** |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewPostBulkPreBanSummaryResponse

`func NewPostBulkPreBanSummaryResponse(status APIStatus, totalRelatedCommentCount int32, emailDomains []string, emails []string, userIds []string, ipHashes []string, reason string, code string, ) *PostBulkPreBanSummaryResponse`

NewPostBulkPreBanSummaryResponse instantiates a new PostBulkPreBanSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBulkPreBanSummaryResponseWithDefaults

`func NewPostBulkPreBanSummaryResponseWithDefaults() *PostBulkPreBanSummaryResponse`

NewPostBulkPreBanSummaryResponseWithDefaults instantiates a new PostBulkPreBanSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PostBulkPreBanSummaryResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostBulkPreBanSummaryResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostBulkPreBanSummaryResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetTotalRelatedCommentCount

`func (o *PostBulkPreBanSummaryResponse) GetTotalRelatedCommentCount() int32`

GetTotalRelatedCommentCount returns the TotalRelatedCommentCount field if non-nil, zero value otherwise.

### GetTotalRelatedCommentCountOk

`func (o *PostBulkPreBanSummaryResponse) GetTotalRelatedCommentCountOk() (*int32, bool)`

GetTotalRelatedCommentCountOk returns a tuple with the TotalRelatedCommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRelatedCommentCount

`func (o *PostBulkPreBanSummaryResponse) SetTotalRelatedCommentCount(v int32)`

SetTotalRelatedCommentCount sets TotalRelatedCommentCount field to given value.


### GetEmailDomains

`func (o *PostBulkPreBanSummaryResponse) GetEmailDomains() []string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *PostBulkPreBanSummaryResponse) GetEmailDomainsOk() (*[]string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *PostBulkPreBanSummaryResponse) SetEmailDomains(v []string)`

SetEmailDomains sets EmailDomains field to given value.


### GetEmails

`func (o *PostBulkPreBanSummaryResponse) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *PostBulkPreBanSummaryResponse) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *PostBulkPreBanSummaryResponse) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetUserIds

`func (o *PostBulkPreBanSummaryResponse) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *PostBulkPreBanSummaryResponse) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *PostBulkPreBanSummaryResponse) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetIpHashes

`func (o *PostBulkPreBanSummaryResponse) GetIpHashes() []string`

GetIpHashes returns the IpHashes field if non-nil, zero value otherwise.

### GetIpHashesOk

`func (o *PostBulkPreBanSummaryResponse) GetIpHashesOk() (*[]string, bool)`

GetIpHashesOk returns a tuple with the IpHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpHashes

`func (o *PostBulkPreBanSummaryResponse) SetIpHashes(v []string)`

SetIpHashes sets IpHashes field to given value.


### GetReason

`func (o *PostBulkPreBanSummaryResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostBulkPreBanSummaryResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostBulkPreBanSummaryResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *PostBulkPreBanSummaryResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostBulkPreBanSummaryResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostBulkPreBanSummaryResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *PostBulkPreBanSummaryResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *PostBulkPreBanSummaryResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *PostBulkPreBanSummaryResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *PostBulkPreBanSummaryResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *PostBulkPreBanSummaryResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *PostBulkPreBanSummaryResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *PostBulkPreBanSummaryResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *PostBulkPreBanSummaryResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *PostBulkPreBanSummaryResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *PostBulkPreBanSummaryResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *PostBulkPreBanSummaryResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *PostBulkPreBanSummaryResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *PostBulkPreBanSummaryResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *PostBulkPreBanSummaryResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *PostBulkPreBanSummaryResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *PostBulkPreBanSummaryResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *PostBulkPreBanSummaryResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *PostBulkPreBanSummaryResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *PostBulkPreBanSummaryResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *PostBulkPreBanSummaryResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


