# CreateTenantPackageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MonthlyCostUSD** | Pointer to **NullableFloat64** |  | [optional] 
**YearlyCostUSD** | Pointer to **NullableFloat64** |  | [optional] 
**MonthlyStripePlanId** | Pointer to **NullableString** |  | [optional] 
**YearlyStripePlanId** | Pointer to **NullableString** |  | [optional] 
**MaxMonthlyPageLoads** | **float64** |  | 
**MaxMonthlyAPICredits** | **float64** |  | 
**MaxMonthlySmallWidgetsCredits** | Pointer to **float64** |  | [optional] 
**MaxMonthlyComments** | **float64** |  | 
**MaxConcurrentUsers** | **float64** |  | 
**MaxTenantUsers** | **float64** |  | 
**MaxSSOUsers** | **float64** |  | 
**MaxModerators** | **float64** |  | 
**MaxDomains** | **float64** |  | 
**MaxWhiteLabeledTenants** | Pointer to **float64** |  | [optional] 
**MaxMonthlyEventLogRequests** | Pointer to **float64** |  | [optional] 
**MaxCustomCollectionSize** | Pointer to **float64** |  | [optional] 
**HasWhiteLabeling** | Pointer to **bool** |  | [optional] 
**HasDebranding** | **bool** |  | 
**HasLLMSpamDetection** | Pointer to **bool** |  | [optional] 
**ForWhoText** | **string** |  | 
**FeatureTaglines** | **[]string** |  | 
**HasAuditing** | Pointer to **bool** |  | [optional] 
**HasFlexPricing** | **bool** |  | 
**EnableSAML** | Pointer to **bool** |  | [optional] 
**FlexPageLoadCostCents** | Pointer to **float64** |  | [optional] 
**FlexPageLoadUnit** | Pointer to **float64** |  | [optional] 
**FlexCommentCostCents** | Pointer to **float64** |  | [optional] 
**FlexCommentUnit** | Pointer to **float64** |  | [optional] 
**FlexSSOUserCostCents** | Pointer to **float64** |  | [optional] 
**FlexSSOUserUnit** | Pointer to **float64** |  | [optional] 
**FlexAPICreditCostCents** | Pointer to **float64** |  | [optional] 
**FlexAPICreditUnit** | Pointer to **float64** |  | [optional] 
**FlexSmallWidgetsCreditCostCents** | Pointer to **float64** |  | [optional] 
**FlexSmallWidgetsCreditUnit** | Pointer to **float64** |  | [optional] 
**FlexModeratorCostCents** | Pointer to **float64** |  | [optional] 
**FlexModeratorUnit** | Pointer to **float64** |  | [optional] 
**FlexAdminCostCents** | Pointer to **float64** |  | [optional] 
**FlexAdminUnit** | Pointer to **float64** |  | [optional] 
**FlexDomainCostCents** | Pointer to **float64** |  | [optional] 
**FlexDomainUnit** | Pointer to **float64** |  | [optional] 
**FlexChatGPTCostCents** | Pointer to **float64** |  | [optional] 
**FlexChatGPTUnit** | Pointer to **float64** |  | [optional] 
**FlexMinimumCostCents** | Pointer to **float64** |  | [optional] 
**FlexManagedTenantCostCents** | Pointer to **float64** |  | [optional] 
**FlexSSOAdminCostCents** | Pointer to **float64** |  | [optional] 
**FlexSSOAdminUnit** | Pointer to **float64** |  | [optional] 
**FlexSSOModeratorCostCents** | Pointer to **float64** |  | [optional] 
**FlexSSOModeratorUnit** | Pointer to **float64** |  | [optional] 

## Methods

### NewCreateTenantPackageBody

`func NewCreateTenantPackageBody(name string, maxMonthlyPageLoads float64, maxMonthlyAPICredits float64, maxMonthlyComments float64, maxConcurrentUsers float64, maxTenantUsers float64, maxSSOUsers float64, maxModerators float64, maxDomains float64, hasDebranding bool, forWhoText string, featureTaglines []string, hasFlexPricing bool, ) *CreateTenantPackageBody`

NewCreateTenantPackageBody instantiates a new CreateTenantPackageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantPackageBodyWithDefaults

`func NewCreateTenantPackageBodyWithDefaults() *CreateTenantPackageBody`

NewCreateTenantPackageBodyWithDefaults instantiates a new CreateTenantPackageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTenantPackageBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantPackageBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantPackageBody) SetName(v string)`

SetName sets Name field to given value.


### GetMonthlyCostUSD

`func (o *CreateTenantPackageBody) GetMonthlyCostUSD() float64`

GetMonthlyCostUSD returns the MonthlyCostUSD field if non-nil, zero value otherwise.

### GetMonthlyCostUSDOk

`func (o *CreateTenantPackageBody) GetMonthlyCostUSDOk() (*float64, bool)`

GetMonthlyCostUSDOk returns a tuple with the MonthlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCostUSD

`func (o *CreateTenantPackageBody) SetMonthlyCostUSD(v float64)`

SetMonthlyCostUSD sets MonthlyCostUSD field to given value.

### HasMonthlyCostUSD

`func (o *CreateTenantPackageBody) HasMonthlyCostUSD() bool`

HasMonthlyCostUSD returns a boolean if a field has been set.

### SetMonthlyCostUSDNil

`func (o *CreateTenantPackageBody) SetMonthlyCostUSDNil(b bool)`

 SetMonthlyCostUSDNil sets the value for MonthlyCostUSD to be an explicit nil

### UnsetMonthlyCostUSD
`func (o *CreateTenantPackageBody) UnsetMonthlyCostUSD()`

UnsetMonthlyCostUSD ensures that no value is present for MonthlyCostUSD, not even an explicit nil
### GetYearlyCostUSD

`func (o *CreateTenantPackageBody) GetYearlyCostUSD() float64`

GetYearlyCostUSD returns the YearlyCostUSD field if non-nil, zero value otherwise.

### GetYearlyCostUSDOk

`func (o *CreateTenantPackageBody) GetYearlyCostUSDOk() (*float64, bool)`

GetYearlyCostUSDOk returns a tuple with the YearlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCostUSD

`func (o *CreateTenantPackageBody) SetYearlyCostUSD(v float64)`

SetYearlyCostUSD sets YearlyCostUSD field to given value.

### HasYearlyCostUSD

`func (o *CreateTenantPackageBody) HasYearlyCostUSD() bool`

HasYearlyCostUSD returns a boolean if a field has been set.

### SetYearlyCostUSDNil

`func (o *CreateTenantPackageBody) SetYearlyCostUSDNil(b bool)`

 SetYearlyCostUSDNil sets the value for YearlyCostUSD to be an explicit nil

### UnsetYearlyCostUSD
`func (o *CreateTenantPackageBody) UnsetYearlyCostUSD()`

UnsetYearlyCostUSD ensures that no value is present for YearlyCostUSD, not even an explicit nil
### GetMonthlyStripePlanId

`func (o *CreateTenantPackageBody) GetMonthlyStripePlanId() string`

GetMonthlyStripePlanId returns the MonthlyStripePlanId field if non-nil, zero value otherwise.

### GetMonthlyStripePlanIdOk

`func (o *CreateTenantPackageBody) GetMonthlyStripePlanIdOk() (*string, bool)`

GetMonthlyStripePlanIdOk returns a tuple with the MonthlyStripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyStripePlanId

`func (o *CreateTenantPackageBody) SetMonthlyStripePlanId(v string)`

SetMonthlyStripePlanId sets MonthlyStripePlanId field to given value.

### HasMonthlyStripePlanId

`func (o *CreateTenantPackageBody) HasMonthlyStripePlanId() bool`

HasMonthlyStripePlanId returns a boolean if a field has been set.

### SetMonthlyStripePlanIdNil

`func (o *CreateTenantPackageBody) SetMonthlyStripePlanIdNil(b bool)`

 SetMonthlyStripePlanIdNil sets the value for MonthlyStripePlanId to be an explicit nil

### UnsetMonthlyStripePlanId
`func (o *CreateTenantPackageBody) UnsetMonthlyStripePlanId()`

UnsetMonthlyStripePlanId ensures that no value is present for MonthlyStripePlanId, not even an explicit nil
### GetYearlyStripePlanId

`func (o *CreateTenantPackageBody) GetYearlyStripePlanId() string`

GetYearlyStripePlanId returns the YearlyStripePlanId field if non-nil, zero value otherwise.

### GetYearlyStripePlanIdOk

`func (o *CreateTenantPackageBody) GetYearlyStripePlanIdOk() (*string, bool)`

GetYearlyStripePlanIdOk returns a tuple with the YearlyStripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyStripePlanId

`func (o *CreateTenantPackageBody) SetYearlyStripePlanId(v string)`

SetYearlyStripePlanId sets YearlyStripePlanId field to given value.

### HasYearlyStripePlanId

`func (o *CreateTenantPackageBody) HasYearlyStripePlanId() bool`

HasYearlyStripePlanId returns a boolean if a field has been set.

### SetYearlyStripePlanIdNil

`func (o *CreateTenantPackageBody) SetYearlyStripePlanIdNil(b bool)`

 SetYearlyStripePlanIdNil sets the value for YearlyStripePlanId to be an explicit nil

### UnsetYearlyStripePlanId
`func (o *CreateTenantPackageBody) UnsetYearlyStripePlanId()`

UnsetYearlyStripePlanId ensures that no value is present for YearlyStripePlanId, not even an explicit nil
### GetMaxMonthlyPageLoads

`func (o *CreateTenantPackageBody) GetMaxMonthlyPageLoads() float64`

GetMaxMonthlyPageLoads returns the MaxMonthlyPageLoads field if non-nil, zero value otherwise.

### GetMaxMonthlyPageLoadsOk

`func (o *CreateTenantPackageBody) GetMaxMonthlyPageLoadsOk() (*float64, bool)`

GetMaxMonthlyPageLoadsOk returns a tuple with the MaxMonthlyPageLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyPageLoads

`func (o *CreateTenantPackageBody) SetMaxMonthlyPageLoads(v float64)`

SetMaxMonthlyPageLoads sets MaxMonthlyPageLoads field to given value.


### GetMaxMonthlyAPICredits

`func (o *CreateTenantPackageBody) GetMaxMonthlyAPICredits() float64`

GetMaxMonthlyAPICredits returns the MaxMonthlyAPICredits field if non-nil, zero value otherwise.

### GetMaxMonthlyAPICreditsOk

`func (o *CreateTenantPackageBody) GetMaxMonthlyAPICreditsOk() (*float64, bool)`

GetMaxMonthlyAPICreditsOk returns a tuple with the MaxMonthlyAPICredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyAPICredits

`func (o *CreateTenantPackageBody) SetMaxMonthlyAPICredits(v float64)`

SetMaxMonthlyAPICredits sets MaxMonthlyAPICredits field to given value.


### GetMaxMonthlySmallWidgetsCredits

`func (o *CreateTenantPackageBody) GetMaxMonthlySmallWidgetsCredits() float64`

GetMaxMonthlySmallWidgetsCredits returns the MaxMonthlySmallWidgetsCredits field if non-nil, zero value otherwise.

### GetMaxMonthlySmallWidgetsCreditsOk

`func (o *CreateTenantPackageBody) GetMaxMonthlySmallWidgetsCreditsOk() (*float64, bool)`

GetMaxMonthlySmallWidgetsCreditsOk returns a tuple with the MaxMonthlySmallWidgetsCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlySmallWidgetsCredits

`func (o *CreateTenantPackageBody) SetMaxMonthlySmallWidgetsCredits(v float64)`

SetMaxMonthlySmallWidgetsCredits sets MaxMonthlySmallWidgetsCredits field to given value.

### HasMaxMonthlySmallWidgetsCredits

`func (o *CreateTenantPackageBody) HasMaxMonthlySmallWidgetsCredits() bool`

HasMaxMonthlySmallWidgetsCredits returns a boolean if a field has been set.

### GetMaxMonthlyComments

`func (o *CreateTenantPackageBody) GetMaxMonthlyComments() float64`

GetMaxMonthlyComments returns the MaxMonthlyComments field if non-nil, zero value otherwise.

### GetMaxMonthlyCommentsOk

`func (o *CreateTenantPackageBody) GetMaxMonthlyCommentsOk() (*float64, bool)`

GetMaxMonthlyCommentsOk returns a tuple with the MaxMonthlyComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyComments

`func (o *CreateTenantPackageBody) SetMaxMonthlyComments(v float64)`

SetMaxMonthlyComments sets MaxMonthlyComments field to given value.


### GetMaxConcurrentUsers

`func (o *CreateTenantPackageBody) GetMaxConcurrentUsers() float64`

GetMaxConcurrentUsers returns the MaxConcurrentUsers field if non-nil, zero value otherwise.

### GetMaxConcurrentUsersOk

`func (o *CreateTenantPackageBody) GetMaxConcurrentUsersOk() (*float64, bool)`

GetMaxConcurrentUsersOk returns a tuple with the MaxConcurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUsers

`func (o *CreateTenantPackageBody) SetMaxConcurrentUsers(v float64)`

SetMaxConcurrentUsers sets MaxConcurrentUsers field to given value.


### GetMaxTenantUsers

`func (o *CreateTenantPackageBody) GetMaxTenantUsers() float64`

GetMaxTenantUsers returns the MaxTenantUsers field if non-nil, zero value otherwise.

### GetMaxTenantUsersOk

`func (o *CreateTenantPackageBody) GetMaxTenantUsersOk() (*float64, bool)`

GetMaxTenantUsersOk returns a tuple with the MaxTenantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTenantUsers

`func (o *CreateTenantPackageBody) SetMaxTenantUsers(v float64)`

SetMaxTenantUsers sets MaxTenantUsers field to given value.


### GetMaxSSOUsers

`func (o *CreateTenantPackageBody) GetMaxSSOUsers() float64`

GetMaxSSOUsers returns the MaxSSOUsers field if non-nil, zero value otherwise.

### GetMaxSSOUsersOk

`func (o *CreateTenantPackageBody) GetMaxSSOUsersOk() (*float64, bool)`

GetMaxSSOUsersOk returns a tuple with the MaxSSOUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSSOUsers

`func (o *CreateTenantPackageBody) SetMaxSSOUsers(v float64)`

SetMaxSSOUsers sets MaxSSOUsers field to given value.


### GetMaxModerators

`func (o *CreateTenantPackageBody) GetMaxModerators() float64`

GetMaxModerators returns the MaxModerators field if non-nil, zero value otherwise.

### GetMaxModeratorsOk

`func (o *CreateTenantPackageBody) GetMaxModeratorsOk() (*float64, bool)`

GetMaxModeratorsOk returns a tuple with the MaxModerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxModerators

`func (o *CreateTenantPackageBody) SetMaxModerators(v float64)`

SetMaxModerators sets MaxModerators field to given value.


### GetMaxDomains

`func (o *CreateTenantPackageBody) GetMaxDomains() float64`

GetMaxDomains returns the MaxDomains field if non-nil, zero value otherwise.

### GetMaxDomainsOk

`func (o *CreateTenantPackageBody) GetMaxDomainsOk() (*float64, bool)`

GetMaxDomainsOk returns a tuple with the MaxDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDomains

`func (o *CreateTenantPackageBody) SetMaxDomains(v float64)`

SetMaxDomains sets MaxDomains field to given value.


### GetMaxWhiteLabeledTenants

`func (o *CreateTenantPackageBody) GetMaxWhiteLabeledTenants() float64`

GetMaxWhiteLabeledTenants returns the MaxWhiteLabeledTenants field if non-nil, zero value otherwise.

### GetMaxWhiteLabeledTenantsOk

`func (o *CreateTenantPackageBody) GetMaxWhiteLabeledTenantsOk() (*float64, bool)`

GetMaxWhiteLabeledTenantsOk returns a tuple with the MaxWhiteLabeledTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWhiteLabeledTenants

`func (o *CreateTenantPackageBody) SetMaxWhiteLabeledTenants(v float64)`

SetMaxWhiteLabeledTenants sets MaxWhiteLabeledTenants field to given value.

### HasMaxWhiteLabeledTenants

`func (o *CreateTenantPackageBody) HasMaxWhiteLabeledTenants() bool`

HasMaxWhiteLabeledTenants returns a boolean if a field has been set.

### GetMaxMonthlyEventLogRequests

`func (o *CreateTenantPackageBody) GetMaxMonthlyEventLogRequests() float64`

GetMaxMonthlyEventLogRequests returns the MaxMonthlyEventLogRequests field if non-nil, zero value otherwise.

### GetMaxMonthlyEventLogRequestsOk

`func (o *CreateTenantPackageBody) GetMaxMonthlyEventLogRequestsOk() (*float64, bool)`

GetMaxMonthlyEventLogRequestsOk returns a tuple with the MaxMonthlyEventLogRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyEventLogRequests

`func (o *CreateTenantPackageBody) SetMaxMonthlyEventLogRequests(v float64)`

SetMaxMonthlyEventLogRequests sets MaxMonthlyEventLogRequests field to given value.

### HasMaxMonthlyEventLogRequests

`func (o *CreateTenantPackageBody) HasMaxMonthlyEventLogRequests() bool`

HasMaxMonthlyEventLogRequests returns a boolean if a field has been set.

### GetMaxCustomCollectionSize

`func (o *CreateTenantPackageBody) GetMaxCustomCollectionSize() float64`

GetMaxCustomCollectionSize returns the MaxCustomCollectionSize field if non-nil, zero value otherwise.

### GetMaxCustomCollectionSizeOk

`func (o *CreateTenantPackageBody) GetMaxCustomCollectionSizeOk() (*float64, bool)`

GetMaxCustomCollectionSizeOk returns a tuple with the MaxCustomCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCustomCollectionSize

`func (o *CreateTenantPackageBody) SetMaxCustomCollectionSize(v float64)`

SetMaxCustomCollectionSize sets MaxCustomCollectionSize field to given value.

### HasMaxCustomCollectionSize

`func (o *CreateTenantPackageBody) HasMaxCustomCollectionSize() bool`

HasMaxCustomCollectionSize returns a boolean if a field has been set.

### GetHasWhiteLabeling

`func (o *CreateTenantPackageBody) GetHasWhiteLabeling() bool`

GetHasWhiteLabeling returns the HasWhiteLabeling field if non-nil, zero value otherwise.

### GetHasWhiteLabelingOk

`func (o *CreateTenantPackageBody) GetHasWhiteLabelingOk() (*bool, bool)`

GetHasWhiteLabelingOk returns a tuple with the HasWhiteLabeling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWhiteLabeling

`func (o *CreateTenantPackageBody) SetHasWhiteLabeling(v bool)`

SetHasWhiteLabeling sets HasWhiteLabeling field to given value.

### HasHasWhiteLabeling

`func (o *CreateTenantPackageBody) HasHasWhiteLabeling() bool`

HasHasWhiteLabeling returns a boolean if a field has been set.

### GetHasDebranding

`func (o *CreateTenantPackageBody) GetHasDebranding() bool`

GetHasDebranding returns the HasDebranding field if non-nil, zero value otherwise.

### GetHasDebrandingOk

`func (o *CreateTenantPackageBody) GetHasDebrandingOk() (*bool, bool)`

GetHasDebrandingOk returns a tuple with the HasDebranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDebranding

`func (o *CreateTenantPackageBody) SetHasDebranding(v bool)`

SetHasDebranding sets HasDebranding field to given value.


### GetHasLLMSpamDetection

`func (o *CreateTenantPackageBody) GetHasLLMSpamDetection() bool`

GetHasLLMSpamDetection returns the HasLLMSpamDetection field if non-nil, zero value otherwise.

### GetHasLLMSpamDetectionOk

`func (o *CreateTenantPackageBody) GetHasLLMSpamDetectionOk() (*bool, bool)`

GetHasLLMSpamDetectionOk returns a tuple with the HasLLMSpamDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLLMSpamDetection

`func (o *CreateTenantPackageBody) SetHasLLMSpamDetection(v bool)`

SetHasLLMSpamDetection sets HasLLMSpamDetection field to given value.

### HasHasLLMSpamDetection

`func (o *CreateTenantPackageBody) HasHasLLMSpamDetection() bool`

HasHasLLMSpamDetection returns a boolean if a field has been set.

### GetForWhoText

`func (o *CreateTenantPackageBody) GetForWhoText() string`

GetForWhoText returns the ForWhoText field if non-nil, zero value otherwise.

### GetForWhoTextOk

`func (o *CreateTenantPackageBody) GetForWhoTextOk() (*string, bool)`

GetForWhoTextOk returns a tuple with the ForWhoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForWhoText

`func (o *CreateTenantPackageBody) SetForWhoText(v string)`

SetForWhoText sets ForWhoText field to given value.


### GetFeatureTaglines

`func (o *CreateTenantPackageBody) GetFeatureTaglines() []string`

GetFeatureTaglines returns the FeatureTaglines field if non-nil, zero value otherwise.

### GetFeatureTaglinesOk

`func (o *CreateTenantPackageBody) GetFeatureTaglinesOk() (*[]string, bool)`

GetFeatureTaglinesOk returns a tuple with the FeatureTaglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTaglines

`func (o *CreateTenantPackageBody) SetFeatureTaglines(v []string)`

SetFeatureTaglines sets FeatureTaglines field to given value.


### GetHasAuditing

`func (o *CreateTenantPackageBody) GetHasAuditing() bool`

GetHasAuditing returns the HasAuditing field if non-nil, zero value otherwise.

### GetHasAuditingOk

`func (o *CreateTenantPackageBody) GetHasAuditingOk() (*bool, bool)`

GetHasAuditingOk returns a tuple with the HasAuditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAuditing

`func (o *CreateTenantPackageBody) SetHasAuditing(v bool)`

SetHasAuditing sets HasAuditing field to given value.

### HasHasAuditing

`func (o *CreateTenantPackageBody) HasHasAuditing() bool`

HasHasAuditing returns a boolean if a field has been set.

### GetHasFlexPricing

`func (o *CreateTenantPackageBody) GetHasFlexPricing() bool`

GetHasFlexPricing returns the HasFlexPricing field if non-nil, zero value otherwise.

### GetHasFlexPricingOk

`func (o *CreateTenantPackageBody) GetHasFlexPricingOk() (*bool, bool)`

GetHasFlexPricingOk returns a tuple with the HasFlexPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFlexPricing

`func (o *CreateTenantPackageBody) SetHasFlexPricing(v bool)`

SetHasFlexPricing sets HasFlexPricing field to given value.


### GetEnableSAML

`func (o *CreateTenantPackageBody) GetEnableSAML() bool`

GetEnableSAML returns the EnableSAML field if non-nil, zero value otherwise.

### GetEnableSAMLOk

`func (o *CreateTenantPackageBody) GetEnableSAMLOk() (*bool, bool)`

GetEnableSAMLOk returns a tuple with the EnableSAML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSAML

`func (o *CreateTenantPackageBody) SetEnableSAML(v bool)`

SetEnableSAML sets EnableSAML field to given value.

### HasEnableSAML

`func (o *CreateTenantPackageBody) HasEnableSAML() bool`

HasEnableSAML returns a boolean if a field has been set.

### GetFlexPageLoadCostCents

`func (o *CreateTenantPackageBody) GetFlexPageLoadCostCents() float64`

GetFlexPageLoadCostCents returns the FlexPageLoadCostCents field if non-nil, zero value otherwise.

### GetFlexPageLoadCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexPageLoadCostCentsOk() (*float64, bool)`

GetFlexPageLoadCostCentsOk returns a tuple with the FlexPageLoadCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadCostCents

`func (o *CreateTenantPackageBody) SetFlexPageLoadCostCents(v float64)`

SetFlexPageLoadCostCents sets FlexPageLoadCostCents field to given value.

### HasFlexPageLoadCostCents

`func (o *CreateTenantPackageBody) HasFlexPageLoadCostCents() bool`

HasFlexPageLoadCostCents returns a boolean if a field has been set.

### GetFlexPageLoadUnit

`func (o *CreateTenantPackageBody) GetFlexPageLoadUnit() float64`

GetFlexPageLoadUnit returns the FlexPageLoadUnit field if non-nil, zero value otherwise.

### GetFlexPageLoadUnitOk

`func (o *CreateTenantPackageBody) GetFlexPageLoadUnitOk() (*float64, bool)`

GetFlexPageLoadUnitOk returns a tuple with the FlexPageLoadUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadUnit

`func (o *CreateTenantPackageBody) SetFlexPageLoadUnit(v float64)`

SetFlexPageLoadUnit sets FlexPageLoadUnit field to given value.

### HasFlexPageLoadUnit

`func (o *CreateTenantPackageBody) HasFlexPageLoadUnit() bool`

HasFlexPageLoadUnit returns a boolean if a field has been set.

### GetFlexCommentCostCents

`func (o *CreateTenantPackageBody) GetFlexCommentCostCents() float64`

GetFlexCommentCostCents returns the FlexCommentCostCents field if non-nil, zero value otherwise.

### GetFlexCommentCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexCommentCostCentsOk() (*float64, bool)`

GetFlexCommentCostCentsOk returns a tuple with the FlexCommentCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentCostCents

`func (o *CreateTenantPackageBody) SetFlexCommentCostCents(v float64)`

SetFlexCommentCostCents sets FlexCommentCostCents field to given value.

### HasFlexCommentCostCents

`func (o *CreateTenantPackageBody) HasFlexCommentCostCents() bool`

HasFlexCommentCostCents returns a boolean if a field has been set.

### GetFlexCommentUnit

`func (o *CreateTenantPackageBody) GetFlexCommentUnit() float64`

GetFlexCommentUnit returns the FlexCommentUnit field if non-nil, zero value otherwise.

### GetFlexCommentUnitOk

`func (o *CreateTenantPackageBody) GetFlexCommentUnitOk() (*float64, bool)`

GetFlexCommentUnitOk returns a tuple with the FlexCommentUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentUnit

`func (o *CreateTenantPackageBody) SetFlexCommentUnit(v float64)`

SetFlexCommentUnit sets FlexCommentUnit field to given value.

### HasFlexCommentUnit

`func (o *CreateTenantPackageBody) HasFlexCommentUnit() bool`

HasFlexCommentUnit returns a boolean if a field has been set.

### GetFlexSSOUserCostCents

`func (o *CreateTenantPackageBody) GetFlexSSOUserCostCents() float64`

GetFlexSSOUserCostCents returns the FlexSSOUserCostCents field if non-nil, zero value otherwise.

### GetFlexSSOUserCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexSSOUserCostCentsOk() (*float64, bool)`

GetFlexSSOUserCostCentsOk returns a tuple with the FlexSSOUserCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserCostCents

`func (o *CreateTenantPackageBody) SetFlexSSOUserCostCents(v float64)`

SetFlexSSOUserCostCents sets FlexSSOUserCostCents field to given value.

### HasFlexSSOUserCostCents

`func (o *CreateTenantPackageBody) HasFlexSSOUserCostCents() bool`

HasFlexSSOUserCostCents returns a boolean if a field has been set.

### GetFlexSSOUserUnit

`func (o *CreateTenantPackageBody) GetFlexSSOUserUnit() float64`

GetFlexSSOUserUnit returns the FlexSSOUserUnit field if non-nil, zero value otherwise.

### GetFlexSSOUserUnitOk

`func (o *CreateTenantPackageBody) GetFlexSSOUserUnitOk() (*float64, bool)`

GetFlexSSOUserUnitOk returns a tuple with the FlexSSOUserUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserUnit

`func (o *CreateTenantPackageBody) SetFlexSSOUserUnit(v float64)`

SetFlexSSOUserUnit sets FlexSSOUserUnit field to given value.

### HasFlexSSOUserUnit

`func (o *CreateTenantPackageBody) HasFlexSSOUserUnit() bool`

HasFlexSSOUserUnit returns a boolean if a field has been set.

### GetFlexAPICreditCostCents

`func (o *CreateTenantPackageBody) GetFlexAPICreditCostCents() float64`

GetFlexAPICreditCostCents returns the FlexAPICreditCostCents field if non-nil, zero value otherwise.

### GetFlexAPICreditCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexAPICreditCostCentsOk() (*float64, bool)`

GetFlexAPICreditCostCentsOk returns a tuple with the FlexAPICreditCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditCostCents

`func (o *CreateTenantPackageBody) SetFlexAPICreditCostCents(v float64)`

SetFlexAPICreditCostCents sets FlexAPICreditCostCents field to given value.

### HasFlexAPICreditCostCents

`func (o *CreateTenantPackageBody) HasFlexAPICreditCostCents() bool`

HasFlexAPICreditCostCents returns a boolean if a field has been set.

### GetFlexAPICreditUnit

`func (o *CreateTenantPackageBody) GetFlexAPICreditUnit() float64`

GetFlexAPICreditUnit returns the FlexAPICreditUnit field if non-nil, zero value otherwise.

### GetFlexAPICreditUnitOk

`func (o *CreateTenantPackageBody) GetFlexAPICreditUnitOk() (*float64, bool)`

GetFlexAPICreditUnitOk returns a tuple with the FlexAPICreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditUnit

`func (o *CreateTenantPackageBody) SetFlexAPICreditUnit(v float64)`

SetFlexAPICreditUnit sets FlexAPICreditUnit field to given value.

### HasFlexAPICreditUnit

`func (o *CreateTenantPackageBody) HasFlexAPICreditUnit() bool`

HasFlexAPICreditUnit returns a boolean if a field has been set.

### GetFlexSmallWidgetsCreditCostCents

`func (o *CreateTenantPackageBody) GetFlexSmallWidgetsCreditCostCents() float64`

GetFlexSmallWidgetsCreditCostCents returns the FlexSmallWidgetsCreditCostCents field if non-nil, zero value otherwise.

### GetFlexSmallWidgetsCreditCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexSmallWidgetsCreditCostCentsOk() (*float64, bool)`

GetFlexSmallWidgetsCreditCostCentsOk returns a tuple with the FlexSmallWidgetsCreditCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSmallWidgetsCreditCostCents

`func (o *CreateTenantPackageBody) SetFlexSmallWidgetsCreditCostCents(v float64)`

SetFlexSmallWidgetsCreditCostCents sets FlexSmallWidgetsCreditCostCents field to given value.

### HasFlexSmallWidgetsCreditCostCents

`func (o *CreateTenantPackageBody) HasFlexSmallWidgetsCreditCostCents() bool`

HasFlexSmallWidgetsCreditCostCents returns a boolean if a field has been set.

### GetFlexSmallWidgetsCreditUnit

`func (o *CreateTenantPackageBody) GetFlexSmallWidgetsCreditUnit() float64`

GetFlexSmallWidgetsCreditUnit returns the FlexSmallWidgetsCreditUnit field if non-nil, zero value otherwise.

### GetFlexSmallWidgetsCreditUnitOk

`func (o *CreateTenantPackageBody) GetFlexSmallWidgetsCreditUnitOk() (*float64, bool)`

GetFlexSmallWidgetsCreditUnitOk returns a tuple with the FlexSmallWidgetsCreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSmallWidgetsCreditUnit

`func (o *CreateTenantPackageBody) SetFlexSmallWidgetsCreditUnit(v float64)`

SetFlexSmallWidgetsCreditUnit sets FlexSmallWidgetsCreditUnit field to given value.

### HasFlexSmallWidgetsCreditUnit

`func (o *CreateTenantPackageBody) HasFlexSmallWidgetsCreditUnit() bool`

HasFlexSmallWidgetsCreditUnit returns a boolean if a field has been set.

### GetFlexModeratorCostCents

`func (o *CreateTenantPackageBody) GetFlexModeratorCostCents() float64`

GetFlexModeratorCostCents returns the FlexModeratorCostCents field if non-nil, zero value otherwise.

### GetFlexModeratorCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexModeratorCostCentsOk() (*float64, bool)`

GetFlexModeratorCostCentsOk returns a tuple with the FlexModeratorCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorCostCents

`func (o *CreateTenantPackageBody) SetFlexModeratorCostCents(v float64)`

SetFlexModeratorCostCents sets FlexModeratorCostCents field to given value.

### HasFlexModeratorCostCents

`func (o *CreateTenantPackageBody) HasFlexModeratorCostCents() bool`

HasFlexModeratorCostCents returns a boolean if a field has been set.

### GetFlexModeratorUnit

`func (o *CreateTenantPackageBody) GetFlexModeratorUnit() float64`

GetFlexModeratorUnit returns the FlexModeratorUnit field if non-nil, zero value otherwise.

### GetFlexModeratorUnitOk

`func (o *CreateTenantPackageBody) GetFlexModeratorUnitOk() (*float64, bool)`

GetFlexModeratorUnitOk returns a tuple with the FlexModeratorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorUnit

`func (o *CreateTenantPackageBody) SetFlexModeratorUnit(v float64)`

SetFlexModeratorUnit sets FlexModeratorUnit field to given value.

### HasFlexModeratorUnit

`func (o *CreateTenantPackageBody) HasFlexModeratorUnit() bool`

HasFlexModeratorUnit returns a boolean if a field has been set.

### GetFlexAdminCostCents

`func (o *CreateTenantPackageBody) GetFlexAdminCostCents() float64`

GetFlexAdminCostCents returns the FlexAdminCostCents field if non-nil, zero value otherwise.

### GetFlexAdminCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexAdminCostCentsOk() (*float64, bool)`

GetFlexAdminCostCentsOk returns a tuple with the FlexAdminCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminCostCents

`func (o *CreateTenantPackageBody) SetFlexAdminCostCents(v float64)`

SetFlexAdminCostCents sets FlexAdminCostCents field to given value.

### HasFlexAdminCostCents

`func (o *CreateTenantPackageBody) HasFlexAdminCostCents() bool`

HasFlexAdminCostCents returns a boolean if a field has been set.

### GetFlexAdminUnit

`func (o *CreateTenantPackageBody) GetFlexAdminUnit() float64`

GetFlexAdminUnit returns the FlexAdminUnit field if non-nil, zero value otherwise.

### GetFlexAdminUnitOk

`func (o *CreateTenantPackageBody) GetFlexAdminUnitOk() (*float64, bool)`

GetFlexAdminUnitOk returns a tuple with the FlexAdminUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminUnit

`func (o *CreateTenantPackageBody) SetFlexAdminUnit(v float64)`

SetFlexAdminUnit sets FlexAdminUnit field to given value.

### HasFlexAdminUnit

`func (o *CreateTenantPackageBody) HasFlexAdminUnit() bool`

HasFlexAdminUnit returns a boolean if a field has been set.

### GetFlexDomainCostCents

`func (o *CreateTenantPackageBody) GetFlexDomainCostCents() float64`

GetFlexDomainCostCents returns the FlexDomainCostCents field if non-nil, zero value otherwise.

### GetFlexDomainCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexDomainCostCentsOk() (*float64, bool)`

GetFlexDomainCostCentsOk returns a tuple with the FlexDomainCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainCostCents

`func (o *CreateTenantPackageBody) SetFlexDomainCostCents(v float64)`

SetFlexDomainCostCents sets FlexDomainCostCents field to given value.

### HasFlexDomainCostCents

`func (o *CreateTenantPackageBody) HasFlexDomainCostCents() bool`

HasFlexDomainCostCents returns a boolean if a field has been set.

### GetFlexDomainUnit

`func (o *CreateTenantPackageBody) GetFlexDomainUnit() float64`

GetFlexDomainUnit returns the FlexDomainUnit field if non-nil, zero value otherwise.

### GetFlexDomainUnitOk

`func (o *CreateTenantPackageBody) GetFlexDomainUnitOk() (*float64, bool)`

GetFlexDomainUnitOk returns a tuple with the FlexDomainUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainUnit

`func (o *CreateTenantPackageBody) SetFlexDomainUnit(v float64)`

SetFlexDomainUnit sets FlexDomainUnit field to given value.

### HasFlexDomainUnit

`func (o *CreateTenantPackageBody) HasFlexDomainUnit() bool`

HasFlexDomainUnit returns a boolean if a field has been set.

### GetFlexChatGPTCostCents

`func (o *CreateTenantPackageBody) GetFlexChatGPTCostCents() float64`

GetFlexChatGPTCostCents returns the FlexChatGPTCostCents field if non-nil, zero value otherwise.

### GetFlexChatGPTCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexChatGPTCostCentsOk() (*float64, bool)`

GetFlexChatGPTCostCentsOk returns a tuple with the FlexChatGPTCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexChatGPTCostCents

`func (o *CreateTenantPackageBody) SetFlexChatGPTCostCents(v float64)`

SetFlexChatGPTCostCents sets FlexChatGPTCostCents field to given value.

### HasFlexChatGPTCostCents

`func (o *CreateTenantPackageBody) HasFlexChatGPTCostCents() bool`

HasFlexChatGPTCostCents returns a boolean if a field has been set.

### GetFlexChatGPTUnit

`func (o *CreateTenantPackageBody) GetFlexChatGPTUnit() float64`

GetFlexChatGPTUnit returns the FlexChatGPTUnit field if non-nil, zero value otherwise.

### GetFlexChatGPTUnitOk

`func (o *CreateTenantPackageBody) GetFlexChatGPTUnitOk() (*float64, bool)`

GetFlexChatGPTUnitOk returns a tuple with the FlexChatGPTUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexChatGPTUnit

`func (o *CreateTenantPackageBody) SetFlexChatGPTUnit(v float64)`

SetFlexChatGPTUnit sets FlexChatGPTUnit field to given value.

### HasFlexChatGPTUnit

`func (o *CreateTenantPackageBody) HasFlexChatGPTUnit() bool`

HasFlexChatGPTUnit returns a boolean if a field has been set.

### GetFlexMinimumCostCents

`func (o *CreateTenantPackageBody) GetFlexMinimumCostCents() float64`

GetFlexMinimumCostCents returns the FlexMinimumCostCents field if non-nil, zero value otherwise.

### GetFlexMinimumCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexMinimumCostCentsOk() (*float64, bool)`

GetFlexMinimumCostCentsOk returns a tuple with the FlexMinimumCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMinimumCostCents

`func (o *CreateTenantPackageBody) SetFlexMinimumCostCents(v float64)`

SetFlexMinimumCostCents sets FlexMinimumCostCents field to given value.

### HasFlexMinimumCostCents

`func (o *CreateTenantPackageBody) HasFlexMinimumCostCents() bool`

HasFlexMinimumCostCents returns a boolean if a field has been set.

### GetFlexManagedTenantCostCents

`func (o *CreateTenantPackageBody) GetFlexManagedTenantCostCents() float64`

GetFlexManagedTenantCostCents returns the FlexManagedTenantCostCents field if non-nil, zero value otherwise.

### GetFlexManagedTenantCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexManagedTenantCostCentsOk() (*float64, bool)`

GetFlexManagedTenantCostCentsOk returns a tuple with the FlexManagedTenantCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexManagedTenantCostCents

`func (o *CreateTenantPackageBody) SetFlexManagedTenantCostCents(v float64)`

SetFlexManagedTenantCostCents sets FlexManagedTenantCostCents field to given value.

### HasFlexManagedTenantCostCents

`func (o *CreateTenantPackageBody) HasFlexManagedTenantCostCents() bool`

HasFlexManagedTenantCostCents returns a boolean if a field has been set.

### GetFlexSSOAdminCostCents

`func (o *CreateTenantPackageBody) GetFlexSSOAdminCostCents() float64`

GetFlexSSOAdminCostCents returns the FlexSSOAdminCostCents field if non-nil, zero value otherwise.

### GetFlexSSOAdminCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexSSOAdminCostCentsOk() (*float64, bool)`

GetFlexSSOAdminCostCentsOk returns a tuple with the FlexSSOAdminCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOAdminCostCents

`func (o *CreateTenantPackageBody) SetFlexSSOAdminCostCents(v float64)`

SetFlexSSOAdminCostCents sets FlexSSOAdminCostCents field to given value.

### HasFlexSSOAdminCostCents

`func (o *CreateTenantPackageBody) HasFlexSSOAdminCostCents() bool`

HasFlexSSOAdminCostCents returns a boolean if a field has been set.

### GetFlexSSOAdminUnit

`func (o *CreateTenantPackageBody) GetFlexSSOAdminUnit() float64`

GetFlexSSOAdminUnit returns the FlexSSOAdminUnit field if non-nil, zero value otherwise.

### GetFlexSSOAdminUnitOk

`func (o *CreateTenantPackageBody) GetFlexSSOAdminUnitOk() (*float64, bool)`

GetFlexSSOAdminUnitOk returns a tuple with the FlexSSOAdminUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOAdminUnit

`func (o *CreateTenantPackageBody) SetFlexSSOAdminUnit(v float64)`

SetFlexSSOAdminUnit sets FlexSSOAdminUnit field to given value.

### HasFlexSSOAdminUnit

`func (o *CreateTenantPackageBody) HasFlexSSOAdminUnit() bool`

HasFlexSSOAdminUnit returns a boolean if a field has been set.

### GetFlexSSOModeratorCostCents

`func (o *CreateTenantPackageBody) GetFlexSSOModeratorCostCents() float64`

GetFlexSSOModeratorCostCents returns the FlexSSOModeratorCostCents field if non-nil, zero value otherwise.

### GetFlexSSOModeratorCostCentsOk

`func (o *CreateTenantPackageBody) GetFlexSSOModeratorCostCentsOk() (*float64, bool)`

GetFlexSSOModeratorCostCentsOk returns a tuple with the FlexSSOModeratorCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOModeratorCostCents

`func (o *CreateTenantPackageBody) SetFlexSSOModeratorCostCents(v float64)`

SetFlexSSOModeratorCostCents sets FlexSSOModeratorCostCents field to given value.

### HasFlexSSOModeratorCostCents

`func (o *CreateTenantPackageBody) HasFlexSSOModeratorCostCents() bool`

HasFlexSSOModeratorCostCents returns a boolean if a field has been set.

### GetFlexSSOModeratorUnit

`func (o *CreateTenantPackageBody) GetFlexSSOModeratorUnit() float64`

GetFlexSSOModeratorUnit returns the FlexSSOModeratorUnit field if non-nil, zero value otherwise.

### GetFlexSSOModeratorUnitOk

`func (o *CreateTenantPackageBody) GetFlexSSOModeratorUnitOk() (*float64, bool)`

GetFlexSSOModeratorUnitOk returns a tuple with the FlexSSOModeratorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOModeratorUnit

`func (o *CreateTenantPackageBody) SetFlexSSOModeratorUnit(v float64)`

SetFlexSSOModeratorUnit sets FlexSSOModeratorUnit field to given value.

### HasFlexSSOModeratorUnit

`func (o *CreateTenantPackageBody) HasFlexSSOModeratorUnit() bool`

HasFlexSSOModeratorUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


