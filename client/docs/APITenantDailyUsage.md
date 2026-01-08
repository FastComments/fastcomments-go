# APITenantDailyUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**YearNumber** | **float64** |  | 
**MonthNumber** | **float64** |  | 
**DayNumber** | **float64** |  | 
**CommentFetchCount** | **float64** |  | 
**CommentCreateCount** | **float64** |  | 
**ConversationCreateCount** | **float64** |  | 
**VoteCount** | **float64** |  | 
**AccountCreatedCount** | **float64** |  | 
**UserMentionSearch** | **float64** |  | 
**HashTagSearch** | **float64** |  | 
**GifSearchTrending** | **float64** |  | 
**GifSearch** | **float64** |  | 
**ApiCreditsUsed** | **float64** |  | 
**CreatedAt** | **time.Time** |  | 
**Billed** | **bool** |  | 
**Ignored** | **bool** |  | 
**ApiErrorCount** | **float64** |  | 

## Methods

### NewAPITenantDailyUsage

`func NewAPITenantDailyUsage(id string, tenantId string, yearNumber float64, monthNumber float64, dayNumber float64, commentFetchCount float64, commentCreateCount float64, conversationCreateCount float64, voteCount float64, accountCreatedCount float64, userMentionSearch float64, hashTagSearch float64, gifSearchTrending float64, gifSearch float64, apiCreditsUsed float64, createdAt time.Time, billed bool, ignored bool, apiErrorCount float64, ) *APITenantDailyUsage`

NewAPITenantDailyUsage instantiates a new APITenantDailyUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITenantDailyUsageWithDefaults

`func NewAPITenantDailyUsageWithDefaults() *APITenantDailyUsage`

NewAPITenantDailyUsageWithDefaults instantiates a new APITenantDailyUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITenantDailyUsage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITenantDailyUsage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITenantDailyUsage) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *APITenantDailyUsage) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *APITenantDailyUsage) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *APITenantDailyUsage) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetYearNumber

`func (o *APITenantDailyUsage) GetYearNumber() float64`

GetYearNumber returns the YearNumber field if non-nil, zero value otherwise.

### GetYearNumberOk

`func (o *APITenantDailyUsage) GetYearNumberOk() (*float64, bool)`

GetYearNumberOk returns a tuple with the YearNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearNumber

`func (o *APITenantDailyUsage) SetYearNumber(v float64)`

SetYearNumber sets YearNumber field to given value.


### GetMonthNumber

`func (o *APITenantDailyUsage) GetMonthNumber() float64`

GetMonthNumber returns the MonthNumber field if non-nil, zero value otherwise.

### GetMonthNumberOk

`func (o *APITenantDailyUsage) GetMonthNumberOk() (*float64, bool)`

GetMonthNumberOk returns a tuple with the MonthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthNumber

`func (o *APITenantDailyUsage) SetMonthNumber(v float64)`

SetMonthNumber sets MonthNumber field to given value.


### GetDayNumber

`func (o *APITenantDailyUsage) GetDayNumber() float64`

GetDayNumber returns the DayNumber field if non-nil, zero value otherwise.

### GetDayNumberOk

`func (o *APITenantDailyUsage) GetDayNumberOk() (*float64, bool)`

GetDayNumberOk returns a tuple with the DayNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayNumber

`func (o *APITenantDailyUsage) SetDayNumber(v float64)`

SetDayNumber sets DayNumber field to given value.


### GetCommentFetchCount

`func (o *APITenantDailyUsage) GetCommentFetchCount() float64`

GetCommentFetchCount returns the CommentFetchCount field if non-nil, zero value otherwise.

### GetCommentFetchCountOk

`func (o *APITenantDailyUsage) GetCommentFetchCountOk() (*float64, bool)`

GetCommentFetchCountOk returns a tuple with the CommentFetchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentFetchCount

`func (o *APITenantDailyUsage) SetCommentFetchCount(v float64)`

SetCommentFetchCount sets CommentFetchCount field to given value.


### GetCommentCreateCount

`func (o *APITenantDailyUsage) GetCommentCreateCount() float64`

GetCommentCreateCount returns the CommentCreateCount field if non-nil, zero value otherwise.

### GetCommentCreateCountOk

`func (o *APITenantDailyUsage) GetCommentCreateCountOk() (*float64, bool)`

GetCommentCreateCountOk returns a tuple with the CommentCreateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCreateCount

`func (o *APITenantDailyUsage) SetCommentCreateCount(v float64)`

SetCommentCreateCount sets CommentCreateCount field to given value.


### GetConversationCreateCount

`func (o *APITenantDailyUsage) GetConversationCreateCount() float64`

GetConversationCreateCount returns the ConversationCreateCount field if non-nil, zero value otherwise.

### GetConversationCreateCountOk

`func (o *APITenantDailyUsage) GetConversationCreateCountOk() (*float64, bool)`

GetConversationCreateCountOk returns a tuple with the ConversationCreateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationCreateCount

`func (o *APITenantDailyUsage) SetConversationCreateCount(v float64)`

SetConversationCreateCount sets ConversationCreateCount field to given value.


### GetVoteCount

`func (o *APITenantDailyUsage) GetVoteCount() float64`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *APITenantDailyUsage) GetVoteCountOk() (*float64, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *APITenantDailyUsage) SetVoteCount(v float64)`

SetVoteCount sets VoteCount field to given value.


### GetAccountCreatedCount

`func (o *APITenantDailyUsage) GetAccountCreatedCount() float64`

GetAccountCreatedCount returns the AccountCreatedCount field if non-nil, zero value otherwise.

### GetAccountCreatedCountOk

`func (o *APITenantDailyUsage) GetAccountCreatedCountOk() (*float64, bool)`

GetAccountCreatedCountOk returns a tuple with the AccountCreatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreatedCount

`func (o *APITenantDailyUsage) SetAccountCreatedCount(v float64)`

SetAccountCreatedCount sets AccountCreatedCount field to given value.


### GetUserMentionSearch

`func (o *APITenantDailyUsage) GetUserMentionSearch() float64`

GetUserMentionSearch returns the UserMentionSearch field if non-nil, zero value otherwise.

### GetUserMentionSearchOk

`func (o *APITenantDailyUsage) GetUserMentionSearchOk() (*float64, bool)`

GetUserMentionSearchOk returns a tuple with the UserMentionSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMentionSearch

`func (o *APITenantDailyUsage) SetUserMentionSearch(v float64)`

SetUserMentionSearch sets UserMentionSearch field to given value.


### GetHashTagSearch

`func (o *APITenantDailyUsage) GetHashTagSearch() float64`

GetHashTagSearch returns the HashTagSearch field if non-nil, zero value otherwise.

### GetHashTagSearchOk

`func (o *APITenantDailyUsage) GetHashTagSearchOk() (*float64, bool)`

GetHashTagSearchOk returns a tuple with the HashTagSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTagSearch

`func (o *APITenantDailyUsage) SetHashTagSearch(v float64)`

SetHashTagSearch sets HashTagSearch field to given value.


### GetGifSearchTrending

`func (o *APITenantDailyUsage) GetGifSearchTrending() float64`

GetGifSearchTrending returns the GifSearchTrending field if non-nil, zero value otherwise.

### GetGifSearchTrendingOk

`func (o *APITenantDailyUsage) GetGifSearchTrendingOk() (*float64, bool)`

GetGifSearchTrendingOk returns a tuple with the GifSearchTrending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifSearchTrending

`func (o *APITenantDailyUsage) SetGifSearchTrending(v float64)`

SetGifSearchTrending sets GifSearchTrending field to given value.


### GetGifSearch

`func (o *APITenantDailyUsage) GetGifSearch() float64`

GetGifSearch returns the GifSearch field if non-nil, zero value otherwise.

### GetGifSearchOk

`func (o *APITenantDailyUsage) GetGifSearchOk() (*float64, bool)`

GetGifSearchOk returns a tuple with the GifSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifSearch

`func (o *APITenantDailyUsage) SetGifSearch(v float64)`

SetGifSearch sets GifSearch field to given value.


### GetApiCreditsUsed

`func (o *APITenantDailyUsage) GetApiCreditsUsed() float64`

GetApiCreditsUsed returns the ApiCreditsUsed field if non-nil, zero value otherwise.

### GetApiCreditsUsedOk

`func (o *APITenantDailyUsage) GetApiCreditsUsedOk() (*float64, bool)`

GetApiCreditsUsedOk returns a tuple with the ApiCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCreditsUsed

`func (o *APITenantDailyUsage) SetApiCreditsUsed(v float64)`

SetApiCreditsUsed sets ApiCreditsUsed field to given value.


### GetCreatedAt

`func (o *APITenantDailyUsage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APITenantDailyUsage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APITenantDailyUsage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetBilled

`func (o *APITenantDailyUsage) GetBilled() bool`

GetBilled returns the Billed field if non-nil, zero value otherwise.

### GetBilledOk

`func (o *APITenantDailyUsage) GetBilledOk() (*bool, bool)`

GetBilledOk returns a tuple with the Billed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilled

`func (o *APITenantDailyUsage) SetBilled(v bool)`

SetBilled sets Billed field to given value.


### GetIgnored

`func (o *APITenantDailyUsage) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *APITenantDailyUsage) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *APITenantDailyUsage) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.


### GetApiErrorCount

`func (o *APITenantDailyUsage) GetApiErrorCount() float64`

GetApiErrorCount returns the ApiErrorCount field if non-nil, zero value otherwise.

### GetApiErrorCountOk

`func (o *APITenantDailyUsage) GetApiErrorCountOk() (*float64, bool)`

GetApiErrorCountOk returns a tuple with the ApiErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiErrorCount

`func (o *APITenantDailyUsage) SetApiErrorCount(v float64)`

SetApiErrorCount sets ApiErrorCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


