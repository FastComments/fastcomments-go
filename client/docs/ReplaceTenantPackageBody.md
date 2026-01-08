# ReplaceTenantPackageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MonthlyCostUSD** | **float64** |  | 
**YearlyCostUSD** | **float64** |  | 
**MaxMonthlyPageLoads** | **float64** |  | 
**MaxMonthlyAPICredits** | **float64** |  | 
**MaxMonthlyComments** | **float64** |  | 
**MaxConcurrentUsers** | **float64** |  | 
**MaxTenantUsers** | **float64** |  | 
**MaxSSOUsers** | **float64** |  | 
**MaxModerators** | **float64** |  | 
**MaxDomains** | **float64** |  | 
**HasDebranding** | **bool** |  | 
**ForWhoText** | **string** |  | 
**FeatureTaglines** | **[]string** |  | 
**HasFlexPricing** | **bool** |  | 
**FlexPageLoadCostCents** | Pointer to **float64** |  | [optional] 
**FlexPageLoadUnit** | Pointer to **float64** |  | [optional] 
**FlexCommentCostCents** | Pointer to **float64** |  | [optional] 
**FlexCommentUnit** | Pointer to **float64** |  | [optional] 
**FlexSSOUserCostCents** | Pointer to **float64** |  | [optional] 
**FlexSSOUserUnit** | Pointer to **float64** |  | [optional] 
**FlexAPICreditCostCents** | Pointer to **float64** |  | [optional] 
**FlexAPICreditUnit** | Pointer to **float64** |  | [optional] 
**FlexModeratorCostCents** | Pointer to **float64** |  | [optional] 
**FlexModeratorUnit** | Pointer to **float64** |  | [optional] 
**FlexAdminCostCents** | Pointer to **float64** |  | [optional] 
**FlexAdminUnit** | Pointer to **float64** |  | [optional] 
**FlexDomainCostCents** | Pointer to **float64** |  | [optional] 
**FlexDomainUnit** | Pointer to **float64** |  | [optional] 
**FlexMinimumCostCents** | Pointer to **float64** |  | [optional] 

## Methods

### NewReplaceTenantPackageBody

`func NewReplaceTenantPackageBody(name string, monthlyCostUSD float64, yearlyCostUSD float64, maxMonthlyPageLoads float64, maxMonthlyAPICredits float64, maxMonthlyComments float64, maxConcurrentUsers float64, maxTenantUsers float64, maxSSOUsers float64, maxModerators float64, maxDomains float64, hasDebranding bool, forWhoText string, featureTaglines []string, hasFlexPricing bool, ) *ReplaceTenantPackageBody`

NewReplaceTenantPackageBody instantiates a new ReplaceTenantPackageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceTenantPackageBodyWithDefaults

`func NewReplaceTenantPackageBodyWithDefaults() *ReplaceTenantPackageBody`

NewReplaceTenantPackageBodyWithDefaults instantiates a new ReplaceTenantPackageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReplaceTenantPackageBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplaceTenantPackageBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplaceTenantPackageBody) SetName(v string)`

SetName sets Name field to given value.


### GetMonthlyCostUSD

`func (o *ReplaceTenantPackageBody) GetMonthlyCostUSD() float64`

GetMonthlyCostUSD returns the MonthlyCostUSD field if non-nil, zero value otherwise.

### GetMonthlyCostUSDOk

`func (o *ReplaceTenantPackageBody) GetMonthlyCostUSDOk() (*float64, bool)`

GetMonthlyCostUSDOk returns a tuple with the MonthlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCostUSD

`func (o *ReplaceTenantPackageBody) SetMonthlyCostUSD(v float64)`

SetMonthlyCostUSD sets MonthlyCostUSD field to given value.


### GetYearlyCostUSD

`func (o *ReplaceTenantPackageBody) GetYearlyCostUSD() float64`

GetYearlyCostUSD returns the YearlyCostUSD field if non-nil, zero value otherwise.

### GetYearlyCostUSDOk

`func (o *ReplaceTenantPackageBody) GetYearlyCostUSDOk() (*float64, bool)`

GetYearlyCostUSDOk returns a tuple with the YearlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCostUSD

`func (o *ReplaceTenantPackageBody) SetYearlyCostUSD(v float64)`

SetYearlyCostUSD sets YearlyCostUSD field to given value.


### GetMaxMonthlyPageLoads

`func (o *ReplaceTenantPackageBody) GetMaxMonthlyPageLoads() float64`

GetMaxMonthlyPageLoads returns the MaxMonthlyPageLoads field if non-nil, zero value otherwise.

### GetMaxMonthlyPageLoadsOk

`func (o *ReplaceTenantPackageBody) GetMaxMonthlyPageLoadsOk() (*float64, bool)`

GetMaxMonthlyPageLoadsOk returns a tuple with the MaxMonthlyPageLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyPageLoads

`func (o *ReplaceTenantPackageBody) SetMaxMonthlyPageLoads(v float64)`

SetMaxMonthlyPageLoads sets MaxMonthlyPageLoads field to given value.


### GetMaxMonthlyAPICredits

`func (o *ReplaceTenantPackageBody) GetMaxMonthlyAPICredits() float64`

GetMaxMonthlyAPICredits returns the MaxMonthlyAPICredits field if non-nil, zero value otherwise.

### GetMaxMonthlyAPICreditsOk

`func (o *ReplaceTenantPackageBody) GetMaxMonthlyAPICreditsOk() (*float64, bool)`

GetMaxMonthlyAPICreditsOk returns a tuple with the MaxMonthlyAPICredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyAPICredits

`func (o *ReplaceTenantPackageBody) SetMaxMonthlyAPICredits(v float64)`

SetMaxMonthlyAPICredits sets MaxMonthlyAPICredits field to given value.


### GetMaxMonthlyComments

`func (o *ReplaceTenantPackageBody) GetMaxMonthlyComments() float64`

GetMaxMonthlyComments returns the MaxMonthlyComments field if non-nil, zero value otherwise.

### GetMaxMonthlyCommentsOk

`func (o *ReplaceTenantPackageBody) GetMaxMonthlyCommentsOk() (*float64, bool)`

GetMaxMonthlyCommentsOk returns a tuple with the MaxMonthlyComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyComments

`func (o *ReplaceTenantPackageBody) SetMaxMonthlyComments(v float64)`

SetMaxMonthlyComments sets MaxMonthlyComments field to given value.


### GetMaxConcurrentUsers

`func (o *ReplaceTenantPackageBody) GetMaxConcurrentUsers() float64`

GetMaxConcurrentUsers returns the MaxConcurrentUsers field if non-nil, zero value otherwise.

### GetMaxConcurrentUsersOk

`func (o *ReplaceTenantPackageBody) GetMaxConcurrentUsersOk() (*float64, bool)`

GetMaxConcurrentUsersOk returns a tuple with the MaxConcurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUsers

`func (o *ReplaceTenantPackageBody) SetMaxConcurrentUsers(v float64)`

SetMaxConcurrentUsers sets MaxConcurrentUsers field to given value.


### GetMaxTenantUsers

`func (o *ReplaceTenantPackageBody) GetMaxTenantUsers() float64`

GetMaxTenantUsers returns the MaxTenantUsers field if non-nil, zero value otherwise.

### GetMaxTenantUsersOk

`func (o *ReplaceTenantPackageBody) GetMaxTenantUsersOk() (*float64, bool)`

GetMaxTenantUsersOk returns a tuple with the MaxTenantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTenantUsers

`func (o *ReplaceTenantPackageBody) SetMaxTenantUsers(v float64)`

SetMaxTenantUsers sets MaxTenantUsers field to given value.


### GetMaxSSOUsers

`func (o *ReplaceTenantPackageBody) GetMaxSSOUsers() float64`

GetMaxSSOUsers returns the MaxSSOUsers field if non-nil, zero value otherwise.

### GetMaxSSOUsersOk

`func (o *ReplaceTenantPackageBody) GetMaxSSOUsersOk() (*float64, bool)`

GetMaxSSOUsersOk returns a tuple with the MaxSSOUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSSOUsers

`func (o *ReplaceTenantPackageBody) SetMaxSSOUsers(v float64)`

SetMaxSSOUsers sets MaxSSOUsers field to given value.


### GetMaxModerators

`func (o *ReplaceTenantPackageBody) GetMaxModerators() float64`

GetMaxModerators returns the MaxModerators field if non-nil, zero value otherwise.

### GetMaxModeratorsOk

`func (o *ReplaceTenantPackageBody) GetMaxModeratorsOk() (*float64, bool)`

GetMaxModeratorsOk returns a tuple with the MaxModerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxModerators

`func (o *ReplaceTenantPackageBody) SetMaxModerators(v float64)`

SetMaxModerators sets MaxModerators field to given value.


### GetMaxDomains

`func (o *ReplaceTenantPackageBody) GetMaxDomains() float64`

GetMaxDomains returns the MaxDomains field if non-nil, zero value otherwise.

### GetMaxDomainsOk

`func (o *ReplaceTenantPackageBody) GetMaxDomainsOk() (*float64, bool)`

GetMaxDomainsOk returns a tuple with the MaxDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDomains

`func (o *ReplaceTenantPackageBody) SetMaxDomains(v float64)`

SetMaxDomains sets MaxDomains field to given value.


### GetHasDebranding

`func (o *ReplaceTenantPackageBody) GetHasDebranding() bool`

GetHasDebranding returns the HasDebranding field if non-nil, zero value otherwise.

### GetHasDebrandingOk

`func (o *ReplaceTenantPackageBody) GetHasDebrandingOk() (*bool, bool)`

GetHasDebrandingOk returns a tuple with the HasDebranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDebranding

`func (o *ReplaceTenantPackageBody) SetHasDebranding(v bool)`

SetHasDebranding sets HasDebranding field to given value.


### GetForWhoText

`func (o *ReplaceTenantPackageBody) GetForWhoText() string`

GetForWhoText returns the ForWhoText field if non-nil, zero value otherwise.

### GetForWhoTextOk

`func (o *ReplaceTenantPackageBody) GetForWhoTextOk() (*string, bool)`

GetForWhoTextOk returns a tuple with the ForWhoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForWhoText

`func (o *ReplaceTenantPackageBody) SetForWhoText(v string)`

SetForWhoText sets ForWhoText field to given value.


### GetFeatureTaglines

`func (o *ReplaceTenantPackageBody) GetFeatureTaglines() []string`

GetFeatureTaglines returns the FeatureTaglines field if non-nil, zero value otherwise.

### GetFeatureTaglinesOk

`func (o *ReplaceTenantPackageBody) GetFeatureTaglinesOk() (*[]string, bool)`

GetFeatureTaglinesOk returns a tuple with the FeatureTaglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTaglines

`func (o *ReplaceTenantPackageBody) SetFeatureTaglines(v []string)`

SetFeatureTaglines sets FeatureTaglines field to given value.


### GetHasFlexPricing

`func (o *ReplaceTenantPackageBody) GetHasFlexPricing() bool`

GetHasFlexPricing returns the HasFlexPricing field if non-nil, zero value otherwise.

### GetHasFlexPricingOk

`func (o *ReplaceTenantPackageBody) GetHasFlexPricingOk() (*bool, bool)`

GetHasFlexPricingOk returns a tuple with the HasFlexPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFlexPricing

`func (o *ReplaceTenantPackageBody) SetHasFlexPricing(v bool)`

SetHasFlexPricing sets HasFlexPricing field to given value.


### GetFlexPageLoadCostCents

`func (o *ReplaceTenantPackageBody) GetFlexPageLoadCostCents() float64`

GetFlexPageLoadCostCents returns the FlexPageLoadCostCents field if non-nil, zero value otherwise.

### GetFlexPageLoadCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexPageLoadCostCentsOk() (*float64, bool)`

GetFlexPageLoadCostCentsOk returns a tuple with the FlexPageLoadCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadCostCents

`func (o *ReplaceTenantPackageBody) SetFlexPageLoadCostCents(v float64)`

SetFlexPageLoadCostCents sets FlexPageLoadCostCents field to given value.

### HasFlexPageLoadCostCents

`func (o *ReplaceTenantPackageBody) HasFlexPageLoadCostCents() bool`

HasFlexPageLoadCostCents returns a boolean if a field has been set.

### GetFlexPageLoadUnit

`func (o *ReplaceTenantPackageBody) GetFlexPageLoadUnit() float64`

GetFlexPageLoadUnit returns the FlexPageLoadUnit field if non-nil, zero value otherwise.

### GetFlexPageLoadUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexPageLoadUnitOk() (*float64, bool)`

GetFlexPageLoadUnitOk returns a tuple with the FlexPageLoadUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadUnit

`func (o *ReplaceTenantPackageBody) SetFlexPageLoadUnit(v float64)`

SetFlexPageLoadUnit sets FlexPageLoadUnit field to given value.

### HasFlexPageLoadUnit

`func (o *ReplaceTenantPackageBody) HasFlexPageLoadUnit() bool`

HasFlexPageLoadUnit returns a boolean if a field has been set.

### GetFlexCommentCostCents

`func (o *ReplaceTenantPackageBody) GetFlexCommentCostCents() float64`

GetFlexCommentCostCents returns the FlexCommentCostCents field if non-nil, zero value otherwise.

### GetFlexCommentCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexCommentCostCentsOk() (*float64, bool)`

GetFlexCommentCostCentsOk returns a tuple with the FlexCommentCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentCostCents

`func (o *ReplaceTenantPackageBody) SetFlexCommentCostCents(v float64)`

SetFlexCommentCostCents sets FlexCommentCostCents field to given value.

### HasFlexCommentCostCents

`func (o *ReplaceTenantPackageBody) HasFlexCommentCostCents() bool`

HasFlexCommentCostCents returns a boolean if a field has been set.

### GetFlexCommentUnit

`func (o *ReplaceTenantPackageBody) GetFlexCommentUnit() float64`

GetFlexCommentUnit returns the FlexCommentUnit field if non-nil, zero value otherwise.

### GetFlexCommentUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexCommentUnitOk() (*float64, bool)`

GetFlexCommentUnitOk returns a tuple with the FlexCommentUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentUnit

`func (o *ReplaceTenantPackageBody) SetFlexCommentUnit(v float64)`

SetFlexCommentUnit sets FlexCommentUnit field to given value.

### HasFlexCommentUnit

`func (o *ReplaceTenantPackageBody) HasFlexCommentUnit() bool`

HasFlexCommentUnit returns a boolean if a field has been set.

### GetFlexSSOUserCostCents

`func (o *ReplaceTenantPackageBody) GetFlexSSOUserCostCents() float64`

GetFlexSSOUserCostCents returns the FlexSSOUserCostCents field if non-nil, zero value otherwise.

### GetFlexSSOUserCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexSSOUserCostCentsOk() (*float64, bool)`

GetFlexSSOUserCostCentsOk returns a tuple with the FlexSSOUserCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserCostCents

`func (o *ReplaceTenantPackageBody) SetFlexSSOUserCostCents(v float64)`

SetFlexSSOUserCostCents sets FlexSSOUserCostCents field to given value.

### HasFlexSSOUserCostCents

`func (o *ReplaceTenantPackageBody) HasFlexSSOUserCostCents() bool`

HasFlexSSOUserCostCents returns a boolean if a field has been set.

### GetFlexSSOUserUnit

`func (o *ReplaceTenantPackageBody) GetFlexSSOUserUnit() float64`

GetFlexSSOUserUnit returns the FlexSSOUserUnit field if non-nil, zero value otherwise.

### GetFlexSSOUserUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexSSOUserUnitOk() (*float64, bool)`

GetFlexSSOUserUnitOk returns a tuple with the FlexSSOUserUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserUnit

`func (o *ReplaceTenantPackageBody) SetFlexSSOUserUnit(v float64)`

SetFlexSSOUserUnit sets FlexSSOUserUnit field to given value.

### HasFlexSSOUserUnit

`func (o *ReplaceTenantPackageBody) HasFlexSSOUserUnit() bool`

HasFlexSSOUserUnit returns a boolean if a field has been set.

### GetFlexAPICreditCostCents

`func (o *ReplaceTenantPackageBody) GetFlexAPICreditCostCents() float64`

GetFlexAPICreditCostCents returns the FlexAPICreditCostCents field if non-nil, zero value otherwise.

### GetFlexAPICreditCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexAPICreditCostCentsOk() (*float64, bool)`

GetFlexAPICreditCostCentsOk returns a tuple with the FlexAPICreditCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditCostCents

`func (o *ReplaceTenantPackageBody) SetFlexAPICreditCostCents(v float64)`

SetFlexAPICreditCostCents sets FlexAPICreditCostCents field to given value.

### HasFlexAPICreditCostCents

`func (o *ReplaceTenantPackageBody) HasFlexAPICreditCostCents() bool`

HasFlexAPICreditCostCents returns a boolean if a field has been set.

### GetFlexAPICreditUnit

`func (o *ReplaceTenantPackageBody) GetFlexAPICreditUnit() float64`

GetFlexAPICreditUnit returns the FlexAPICreditUnit field if non-nil, zero value otherwise.

### GetFlexAPICreditUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexAPICreditUnitOk() (*float64, bool)`

GetFlexAPICreditUnitOk returns a tuple with the FlexAPICreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditUnit

`func (o *ReplaceTenantPackageBody) SetFlexAPICreditUnit(v float64)`

SetFlexAPICreditUnit sets FlexAPICreditUnit field to given value.

### HasFlexAPICreditUnit

`func (o *ReplaceTenantPackageBody) HasFlexAPICreditUnit() bool`

HasFlexAPICreditUnit returns a boolean if a field has been set.

### GetFlexModeratorCostCents

`func (o *ReplaceTenantPackageBody) GetFlexModeratorCostCents() float64`

GetFlexModeratorCostCents returns the FlexModeratorCostCents field if non-nil, zero value otherwise.

### GetFlexModeratorCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexModeratorCostCentsOk() (*float64, bool)`

GetFlexModeratorCostCentsOk returns a tuple with the FlexModeratorCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorCostCents

`func (o *ReplaceTenantPackageBody) SetFlexModeratorCostCents(v float64)`

SetFlexModeratorCostCents sets FlexModeratorCostCents field to given value.

### HasFlexModeratorCostCents

`func (o *ReplaceTenantPackageBody) HasFlexModeratorCostCents() bool`

HasFlexModeratorCostCents returns a boolean if a field has been set.

### GetFlexModeratorUnit

`func (o *ReplaceTenantPackageBody) GetFlexModeratorUnit() float64`

GetFlexModeratorUnit returns the FlexModeratorUnit field if non-nil, zero value otherwise.

### GetFlexModeratorUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexModeratorUnitOk() (*float64, bool)`

GetFlexModeratorUnitOk returns a tuple with the FlexModeratorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorUnit

`func (o *ReplaceTenantPackageBody) SetFlexModeratorUnit(v float64)`

SetFlexModeratorUnit sets FlexModeratorUnit field to given value.

### HasFlexModeratorUnit

`func (o *ReplaceTenantPackageBody) HasFlexModeratorUnit() bool`

HasFlexModeratorUnit returns a boolean if a field has been set.

### GetFlexAdminCostCents

`func (o *ReplaceTenantPackageBody) GetFlexAdminCostCents() float64`

GetFlexAdminCostCents returns the FlexAdminCostCents field if non-nil, zero value otherwise.

### GetFlexAdminCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexAdminCostCentsOk() (*float64, bool)`

GetFlexAdminCostCentsOk returns a tuple with the FlexAdminCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminCostCents

`func (o *ReplaceTenantPackageBody) SetFlexAdminCostCents(v float64)`

SetFlexAdminCostCents sets FlexAdminCostCents field to given value.

### HasFlexAdminCostCents

`func (o *ReplaceTenantPackageBody) HasFlexAdminCostCents() bool`

HasFlexAdminCostCents returns a boolean if a field has been set.

### GetFlexAdminUnit

`func (o *ReplaceTenantPackageBody) GetFlexAdminUnit() float64`

GetFlexAdminUnit returns the FlexAdminUnit field if non-nil, zero value otherwise.

### GetFlexAdminUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexAdminUnitOk() (*float64, bool)`

GetFlexAdminUnitOk returns a tuple with the FlexAdminUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminUnit

`func (o *ReplaceTenantPackageBody) SetFlexAdminUnit(v float64)`

SetFlexAdminUnit sets FlexAdminUnit field to given value.

### HasFlexAdminUnit

`func (o *ReplaceTenantPackageBody) HasFlexAdminUnit() bool`

HasFlexAdminUnit returns a boolean if a field has been set.

### GetFlexDomainCostCents

`func (o *ReplaceTenantPackageBody) GetFlexDomainCostCents() float64`

GetFlexDomainCostCents returns the FlexDomainCostCents field if non-nil, zero value otherwise.

### GetFlexDomainCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexDomainCostCentsOk() (*float64, bool)`

GetFlexDomainCostCentsOk returns a tuple with the FlexDomainCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainCostCents

`func (o *ReplaceTenantPackageBody) SetFlexDomainCostCents(v float64)`

SetFlexDomainCostCents sets FlexDomainCostCents field to given value.

### HasFlexDomainCostCents

`func (o *ReplaceTenantPackageBody) HasFlexDomainCostCents() bool`

HasFlexDomainCostCents returns a boolean if a field has been set.

### GetFlexDomainUnit

`func (o *ReplaceTenantPackageBody) GetFlexDomainUnit() float64`

GetFlexDomainUnit returns the FlexDomainUnit field if non-nil, zero value otherwise.

### GetFlexDomainUnitOk

`func (o *ReplaceTenantPackageBody) GetFlexDomainUnitOk() (*float64, bool)`

GetFlexDomainUnitOk returns a tuple with the FlexDomainUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainUnit

`func (o *ReplaceTenantPackageBody) SetFlexDomainUnit(v float64)`

SetFlexDomainUnit sets FlexDomainUnit field to given value.

### HasFlexDomainUnit

`func (o *ReplaceTenantPackageBody) HasFlexDomainUnit() bool`

HasFlexDomainUnit returns a boolean if a field has been set.

### GetFlexMinimumCostCents

`func (o *ReplaceTenantPackageBody) GetFlexMinimumCostCents() float64`

GetFlexMinimumCostCents returns the FlexMinimumCostCents field if non-nil, zero value otherwise.

### GetFlexMinimumCostCentsOk

`func (o *ReplaceTenantPackageBody) GetFlexMinimumCostCentsOk() (*float64, bool)`

GetFlexMinimumCostCentsOk returns a tuple with the FlexMinimumCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMinimumCostCents

`func (o *ReplaceTenantPackageBody) SetFlexMinimumCostCents(v float64)`

SetFlexMinimumCostCents sets FlexMinimumCostCents field to given value.

### HasFlexMinimumCostCents

`func (o *ReplaceTenantPackageBody) HasFlexMinimumCostCents() bool`

HasFlexMinimumCostCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


