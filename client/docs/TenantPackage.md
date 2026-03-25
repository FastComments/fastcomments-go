# TenantPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TenantId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**MonthlyCostUSD** | **NullableFloat64** |  | 
**YearlyCostUSD** | **NullableFloat64** |  | 
**MonthlyStripePlanId** | **NullableString** |  | 
**YearlyStripePlanId** | **NullableString** |  | 
**MaxMonthlyPageLoads** | **float64** |  | 
**MaxMonthlyAPICredits** | **float64** |  | 
**MaxMonthlySmallWidgetsCredits** | **float64** |  | 
**MaxMonthlyComments** | **float64** |  | 
**MaxConcurrentUsers** | **float64** |  | 
**MaxTenantUsers** | **float64** |  | 
**MaxSSOUsers** | **float64** |  | 
**MaxModerators** | **float64** |  | 
**MaxDomains** | **float64** |  | 
**MaxWhiteLabeledTenants** | **float64** |  | 
**MaxMonthlyEventLogRequests** | **float64** |  | 
**MaxCustomCollectionSize** | **float64** |  | 
**HasWhiteLabeling** | **bool** |  | 
**HasDebranding** | **bool** |  | 
**HasLLMSpamDetection** | **bool** |  | 
**ForWhoText** | **string** |  | 
**FeatureTaglines** | **[]string** |  | 
**HasAuditing** | **bool** |  | 
**HasFlexPricing** | **bool** |  | 
**EnableSAML** | Pointer to **bool** |  | [optional] 
**EnableCanvasLTI** | Pointer to **bool** |  | [optional] 
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
**IsSSOBillingMonthlyActiveUsers** | Pointer to **bool** |  | [optional] 

## Methods

### NewTenantPackage

`func NewTenantPackage(id string, name string, tenantId string, createdAt time.Time, monthlyCostUSD NullableFloat64, yearlyCostUSD NullableFloat64, monthlyStripePlanId NullableString, yearlyStripePlanId NullableString, maxMonthlyPageLoads float64, maxMonthlyAPICredits float64, maxMonthlySmallWidgetsCredits float64, maxMonthlyComments float64, maxConcurrentUsers float64, maxTenantUsers float64, maxSSOUsers float64, maxModerators float64, maxDomains float64, maxWhiteLabeledTenants float64, maxMonthlyEventLogRequests float64, maxCustomCollectionSize float64, hasWhiteLabeling bool, hasDebranding bool, hasLLMSpamDetection bool, forWhoText string, featureTaglines []string, hasAuditing bool, hasFlexPricing bool, ) *TenantPackage`

NewTenantPackage instantiates a new TenantPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPackageWithDefaults

`func NewTenantPackageWithDefaults() *TenantPackage`

NewTenantPackageWithDefaults instantiates a new TenantPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantPackage) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TenantPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantPackage) SetName(v string)`

SetName sets Name field to given value.


### GetTenantId

`func (o *TenantPackage) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantPackage) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantPackage) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCreatedAt

`func (o *TenantPackage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantPackage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantPackage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMonthlyCostUSD

`func (o *TenantPackage) GetMonthlyCostUSD() float64`

GetMonthlyCostUSD returns the MonthlyCostUSD field if non-nil, zero value otherwise.

### GetMonthlyCostUSDOk

`func (o *TenantPackage) GetMonthlyCostUSDOk() (*float64, bool)`

GetMonthlyCostUSDOk returns a tuple with the MonthlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCostUSD

`func (o *TenantPackage) SetMonthlyCostUSD(v float64)`

SetMonthlyCostUSD sets MonthlyCostUSD field to given value.


### SetMonthlyCostUSDNil

`func (o *TenantPackage) SetMonthlyCostUSDNil(b bool)`

 SetMonthlyCostUSDNil sets the value for MonthlyCostUSD to be an explicit nil

### UnsetMonthlyCostUSD
`func (o *TenantPackage) UnsetMonthlyCostUSD()`

UnsetMonthlyCostUSD ensures that no value is present for MonthlyCostUSD, not even an explicit nil
### GetYearlyCostUSD

`func (o *TenantPackage) GetYearlyCostUSD() float64`

GetYearlyCostUSD returns the YearlyCostUSD field if non-nil, zero value otherwise.

### GetYearlyCostUSDOk

`func (o *TenantPackage) GetYearlyCostUSDOk() (*float64, bool)`

GetYearlyCostUSDOk returns a tuple with the YearlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCostUSD

`func (o *TenantPackage) SetYearlyCostUSD(v float64)`

SetYearlyCostUSD sets YearlyCostUSD field to given value.


### SetYearlyCostUSDNil

`func (o *TenantPackage) SetYearlyCostUSDNil(b bool)`

 SetYearlyCostUSDNil sets the value for YearlyCostUSD to be an explicit nil

### UnsetYearlyCostUSD
`func (o *TenantPackage) UnsetYearlyCostUSD()`

UnsetYearlyCostUSD ensures that no value is present for YearlyCostUSD, not even an explicit nil
### GetMonthlyStripePlanId

`func (o *TenantPackage) GetMonthlyStripePlanId() string`

GetMonthlyStripePlanId returns the MonthlyStripePlanId field if non-nil, zero value otherwise.

### GetMonthlyStripePlanIdOk

`func (o *TenantPackage) GetMonthlyStripePlanIdOk() (*string, bool)`

GetMonthlyStripePlanIdOk returns a tuple with the MonthlyStripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyStripePlanId

`func (o *TenantPackage) SetMonthlyStripePlanId(v string)`

SetMonthlyStripePlanId sets MonthlyStripePlanId field to given value.


### SetMonthlyStripePlanIdNil

`func (o *TenantPackage) SetMonthlyStripePlanIdNil(b bool)`

 SetMonthlyStripePlanIdNil sets the value for MonthlyStripePlanId to be an explicit nil

### UnsetMonthlyStripePlanId
`func (o *TenantPackage) UnsetMonthlyStripePlanId()`

UnsetMonthlyStripePlanId ensures that no value is present for MonthlyStripePlanId, not even an explicit nil
### GetYearlyStripePlanId

`func (o *TenantPackage) GetYearlyStripePlanId() string`

GetYearlyStripePlanId returns the YearlyStripePlanId field if non-nil, zero value otherwise.

### GetYearlyStripePlanIdOk

`func (o *TenantPackage) GetYearlyStripePlanIdOk() (*string, bool)`

GetYearlyStripePlanIdOk returns a tuple with the YearlyStripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyStripePlanId

`func (o *TenantPackage) SetYearlyStripePlanId(v string)`

SetYearlyStripePlanId sets YearlyStripePlanId field to given value.


### SetYearlyStripePlanIdNil

`func (o *TenantPackage) SetYearlyStripePlanIdNil(b bool)`

 SetYearlyStripePlanIdNil sets the value for YearlyStripePlanId to be an explicit nil

### UnsetYearlyStripePlanId
`func (o *TenantPackage) UnsetYearlyStripePlanId()`

UnsetYearlyStripePlanId ensures that no value is present for YearlyStripePlanId, not even an explicit nil
### GetMaxMonthlyPageLoads

`func (o *TenantPackage) GetMaxMonthlyPageLoads() float64`

GetMaxMonthlyPageLoads returns the MaxMonthlyPageLoads field if non-nil, zero value otherwise.

### GetMaxMonthlyPageLoadsOk

`func (o *TenantPackage) GetMaxMonthlyPageLoadsOk() (*float64, bool)`

GetMaxMonthlyPageLoadsOk returns a tuple with the MaxMonthlyPageLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyPageLoads

`func (o *TenantPackage) SetMaxMonthlyPageLoads(v float64)`

SetMaxMonthlyPageLoads sets MaxMonthlyPageLoads field to given value.


### GetMaxMonthlyAPICredits

`func (o *TenantPackage) GetMaxMonthlyAPICredits() float64`

GetMaxMonthlyAPICredits returns the MaxMonthlyAPICredits field if non-nil, zero value otherwise.

### GetMaxMonthlyAPICreditsOk

`func (o *TenantPackage) GetMaxMonthlyAPICreditsOk() (*float64, bool)`

GetMaxMonthlyAPICreditsOk returns a tuple with the MaxMonthlyAPICredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyAPICredits

`func (o *TenantPackage) SetMaxMonthlyAPICredits(v float64)`

SetMaxMonthlyAPICredits sets MaxMonthlyAPICredits field to given value.


### GetMaxMonthlySmallWidgetsCredits

`func (o *TenantPackage) GetMaxMonthlySmallWidgetsCredits() float64`

GetMaxMonthlySmallWidgetsCredits returns the MaxMonthlySmallWidgetsCredits field if non-nil, zero value otherwise.

### GetMaxMonthlySmallWidgetsCreditsOk

`func (o *TenantPackage) GetMaxMonthlySmallWidgetsCreditsOk() (*float64, bool)`

GetMaxMonthlySmallWidgetsCreditsOk returns a tuple with the MaxMonthlySmallWidgetsCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlySmallWidgetsCredits

`func (o *TenantPackage) SetMaxMonthlySmallWidgetsCredits(v float64)`

SetMaxMonthlySmallWidgetsCredits sets MaxMonthlySmallWidgetsCredits field to given value.


### GetMaxMonthlyComments

`func (o *TenantPackage) GetMaxMonthlyComments() float64`

GetMaxMonthlyComments returns the MaxMonthlyComments field if non-nil, zero value otherwise.

### GetMaxMonthlyCommentsOk

`func (o *TenantPackage) GetMaxMonthlyCommentsOk() (*float64, bool)`

GetMaxMonthlyCommentsOk returns a tuple with the MaxMonthlyComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyComments

`func (o *TenantPackage) SetMaxMonthlyComments(v float64)`

SetMaxMonthlyComments sets MaxMonthlyComments field to given value.


### GetMaxConcurrentUsers

`func (o *TenantPackage) GetMaxConcurrentUsers() float64`

GetMaxConcurrentUsers returns the MaxConcurrentUsers field if non-nil, zero value otherwise.

### GetMaxConcurrentUsersOk

`func (o *TenantPackage) GetMaxConcurrentUsersOk() (*float64, bool)`

GetMaxConcurrentUsersOk returns a tuple with the MaxConcurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUsers

`func (o *TenantPackage) SetMaxConcurrentUsers(v float64)`

SetMaxConcurrentUsers sets MaxConcurrentUsers field to given value.


### GetMaxTenantUsers

`func (o *TenantPackage) GetMaxTenantUsers() float64`

GetMaxTenantUsers returns the MaxTenantUsers field if non-nil, zero value otherwise.

### GetMaxTenantUsersOk

`func (o *TenantPackage) GetMaxTenantUsersOk() (*float64, bool)`

GetMaxTenantUsersOk returns a tuple with the MaxTenantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTenantUsers

`func (o *TenantPackage) SetMaxTenantUsers(v float64)`

SetMaxTenantUsers sets MaxTenantUsers field to given value.


### GetMaxSSOUsers

`func (o *TenantPackage) GetMaxSSOUsers() float64`

GetMaxSSOUsers returns the MaxSSOUsers field if non-nil, zero value otherwise.

### GetMaxSSOUsersOk

`func (o *TenantPackage) GetMaxSSOUsersOk() (*float64, bool)`

GetMaxSSOUsersOk returns a tuple with the MaxSSOUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSSOUsers

`func (o *TenantPackage) SetMaxSSOUsers(v float64)`

SetMaxSSOUsers sets MaxSSOUsers field to given value.


### GetMaxModerators

`func (o *TenantPackage) GetMaxModerators() float64`

GetMaxModerators returns the MaxModerators field if non-nil, zero value otherwise.

### GetMaxModeratorsOk

`func (o *TenantPackage) GetMaxModeratorsOk() (*float64, bool)`

GetMaxModeratorsOk returns a tuple with the MaxModerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxModerators

`func (o *TenantPackage) SetMaxModerators(v float64)`

SetMaxModerators sets MaxModerators field to given value.


### GetMaxDomains

`func (o *TenantPackage) GetMaxDomains() float64`

GetMaxDomains returns the MaxDomains field if non-nil, zero value otherwise.

### GetMaxDomainsOk

`func (o *TenantPackage) GetMaxDomainsOk() (*float64, bool)`

GetMaxDomainsOk returns a tuple with the MaxDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDomains

`func (o *TenantPackage) SetMaxDomains(v float64)`

SetMaxDomains sets MaxDomains field to given value.


### GetMaxWhiteLabeledTenants

`func (o *TenantPackage) GetMaxWhiteLabeledTenants() float64`

GetMaxWhiteLabeledTenants returns the MaxWhiteLabeledTenants field if non-nil, zero value otherwise.

### GetMaxWhiteLabeledTenantsOk

`func (o *TenantPackage) GetMaxWhiteLabeledTenantsOk() (*float64, bool)`

GetMaxWhiteLabeledTenantsOk returns a tuple with the MaxWhiteLabeledTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWhiteLabeledTenants

`func (o *TenantPackage) SetMaxWhiteLabeledTenants(v float64)`

SetMaxWhiteLabeledTenants sets MaxWhiteLabeledTenants field to given value.


### GetMaxMonthlyEventLogRequests

`func (o *TenantPackage) GetMaxMonthlyEventLogRequests() float64`

GetMaxMonthlyEventLogRequests returns the MaxMonthlyEventLogRequests field if non-nil, zero value otherwise.

### GetMaxMonthlyEventLogRequestsOk

`func (o *TenantPackage) GetMaxMonthlyEventLogRequestsOk() (*float64, bool)`

GetMaxMonthlyEventLogRequestsOk returns a tuple with the MaxMonthlyEventLogRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyEventLogRequests

`func (o *TenantPackage) SetMaxMonthlyEventLogRequests(v float64)`

SetMaxMonthlyEventLogRequests sets MaxMonthlyEventLogRequests field to given value.


### GetMaxCustomCollectionSize

`func (o *TenantPackage) GetMaxCustomCollectionSize() float64`

GetMaxCustomCollectionSize returns the MaxCustomCollectionSize field if non-nil, zero value otherwise.

### GetMaxCustomCollectionSizeOk

`func (o *TenantPackage) GetMaxCustomCollectionSizeOk() (*float64, bool)`

GetMaxCustomCollectionSizeOk returns a tuple with the MaxCustomCollectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCustomCollectionSize

`func (o *TenantPackage) SetMaxCustomCollectionSize(v float64)`

SetMaxCustomCollectionSize sets MaxCustomCollectionSize field to given value.


### GetHasWhiteLabeling

`func (o *TenantPackage) GetHasWhiteLabeling() bool`

GetHasWhiteLabeling returns the HasWhiteLabeling field if non-nil, zero value otherwise.

### GetHasWhiteLabelingOk

`func (o *TenantPackage) GetHasWhiteLabelingOk() (*bool, bool)`

GetHasWhiteLabelingOk returns a tuple with the HasWhiteLabeling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWhiteLabeling

`func (o *TenantPackage) SetHasWhiteLabeling(v bool)`

SetHasWhiteLabeling sets HasWhiteLabeling field to given value.


### GetHasDebranding

`func (o *TenantPackage) GetHasDebranding() bool`

GetHasDebranding returns the HasDebranding field if non-nil, zero value otherwise.

### GetHasDebrandingOk

`func (o *TenantPackage) GetHasDebrandingOk() (*bool, bool)`

GetHasDebrandingOk returns a tuple with the HasDebranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDebranding

`func (o *TenantPackage) SetHasDebranding(v bool)`

SetHasDebranding sets HasDebranding field to given value.


### GetHasLLMSpamDetection

`func (o *TenantPackage) GetHasLLMSpamDetection() bool`

GetHasLLMSpamDetection returns the HasLLMSpamDetection field if non-nil, zero value otherwise.

### GetHasLLMSpamDetectionOk

`func (o *TenantPackage) GetHasLLMSpamDetectionOk() (*bool, bool)`

GetHasLLMSpamDetectionOk returns a tuple with the HasLLMSpamDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLLMSpamDetection

`func (o *TenantPackage) SetHasLLMSpamDetection(v bool)`

SetHasLLMSpamDetection sets HasLLMSpamDetection field to given value.


### GetForWhoText

`func (o *TenantPackage) GetForWhoText() string`

GetForWhoText returns the ForWhoText field if non-nil, zero value otherwise.

### GetForWhoTextOk

`func (o *TenantPackage) GetForWhoTextOk() (*string, bool)`

GetForWhoTextOk returns a tuple with the ForWhoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForWhoText

`func (o *TenantPackage) SetForWhoText(v string)`

SetForWhoText sets ForWhoText field to given value.


### GetFeatureTaglines

`func (o *TenantPackage) GetFeatureTaglines() []string`

GetFeatureTaglines returns the FeatureTaglines field if non-nil, zero value otherwise.

### GetFeatureTaglinesOk

`func (o *TenantPackage) GetFeatureTaglinesOk() (*[]string, bool)`

GetFeatureTaglinesOk returns a tuple with the FeatureTaglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTaglines

`func (o *TenantPackage) SetFeatureTaglines(v []string)`

SetFeatureTaglines sets FeatureTaglines field to given value.


### GetHasAuditing

`func (o *TenantPackage) GetHasAuditing() bool`

GetHasAuditing returns the HasAuditing field if non-nil, zero value otherwise.

### GetHasAuditingOk

`func (o *TenantPackage) GetHasAuditingOk() (*bool, bool)`

GetHasAuditingOk returns a tuple with the HasAuditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAuditing

`func (o *TenantPackage) SetHasAuditing(v bool)`

SetHasAuditing sets HasAuditing field to given value.


### GetHasFlexPricing

`func (o *TenantPackage) GetHasFlexPricing() bool`

GetHasFlexPricing returns the HasFlexPricing field if non-nil, zero value otherwise.

### GetHasFlexPricingOk

`func (o *TenantPackage) GetHasFlexPricingOk() (*bool, bool)`

GetHasFlexPricingOk returns a tuple with the HasFlexPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFlexPricing

`func (o *TenantPackage) SetHasFlexPricing(v bool)`

SetHasFlexPricing sets HasFlexPricing field to given value.


### GetEnableSAML

`func (o *TenantPackage) GetEnableSAML() bool`

GetEnableSAML returns the EnableSAML field if non-nil, zero value otherwise.

### GetEnableSAMLOk

`func (o *TenantPackage) GetEnableSAMLOk() (*bool, bool)`

GetEnableSAMLOk returns a tuple with the EnableSAML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSAML

`func (o *TenantPackage) SetEnableSAML(v bool)`

SetEnableSAML sets EnableSAML field to given value.

### HasEnableSAML

`func (o *TenantPackage) HasEnableSAML() bool`

HasEnableSAML returns a boolean if a field has been set.

### GetEnableCanvasLTI

`func (o *TenantPackage) GetEnableCanvasLTI() bool`

GetEnableCanvasLTI returns the EnableCanvasLTI field if non-nil, zero value otherwise.

### GetEnableCanvasLTIOk

`func (o *TenantPackage) GetEnableCanvasLTIOk() (*bool, bool)`

GetEnableCanvasLTIOk returns a tuple with the EnableCanvasLTI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCanvasLTI

`func (o *TenantPackage) SetEnableCanvasLTI(v bool)`

SetEnableCanvasLTI sets EnableCanvasLTI field to given value.

### HasEnableCanvasLTI

`func (o *TenantPackage) HasEnableCanvasLTI() bool`

HasEnableCanvasLTI returns a boolean if a field has been set.

### GetFlexPageLoadCostCents

`func (o *TenantPackage) GetFlexPageLoadCostCents() float64`

GetFlexPageLoadCostCents returns the FlexPageLoadCostCents field if non-nil, zero value otherwise.

### GetFlexPageLoadCostCentsOk

`func (o *TenantPackage) GetFlexPageLoadCostCentsOk() (*float64, bool)`

GetFlexPageLoadCostCentsOk returns a tuple with the FlexPageLoadCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadCostCents

`func (o *TenantPackage) SetFlexPageLoadCostCents(v float64)`

SetFlexPageLoadCostCents sets FlexPageLoadCostCents field to given value.

### HasFlexPageLoadCostCents

`func (o *TenantPackage) HasFlexPageLoadCostCents() bool`

HasFlexPageLoadCostCents returns a boolean if a field has been set.

### GetFlexPageLoadUnit

`func (o *TenantPackage) GetFlexPageLoadUnit() float64`

GetFlexPageLoadUnit returns the FlexPageLoadUnit field if non-nil, zero value otherwise.

### GetFlexPageLoadUnitOk

`func (o *TenantPackage) GetFlexPageLoadUnitOk() (*float64, bool)`

GetFlexPageLoadUnitOk returns a tuple with the FlexPageLoadUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadUnit

`func (o *TenantPackage) SetFlexPageLoadUnit(v float64)`

SetFlexPageLoadUnit sets FlexPageLoadUnit field to given value.

### HasFlexPageLoadUnit

`func (o *TenantPackage) HasFlexPageLoadUnit() bool`

HasFlexPageLoadUnit returns a boolean if a field has been set.

### GetFlexCommentCostCents

`func (o *TenantPackage) GetFlexCommentCostCents() float64`

GetFlexCommentCostCents returns the FlexCommentCostCents field if non-nil, zero value otherwise.

### GetFlexCommentCostCentsOk

`func (o *TenantPackage) GetFlexCommentCostCentsOk() (*float64, bool)`

GetFlexCommentCostCentsOk returns a tuple with the FlexCommentCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentCostCents

`func (o *TenantPackage) SetFlexCommentCostCents(v float64)`

SetFlexCommentCostCents sets FlexCommentCostCents field to given value.

### HasFlexCommentCostCents

`func (o *TenantPackage) HasFlexCommentCostCents() bool`

HasFlexCommentCostCents returns a boolean if a field has been set.

### GetFlexCommentUnit

`func (o *TenantPackage) GetFlexCommentUnit() float64`

GetFlexCommentUnit returns the FlexCommentUnit field if non-nil, zero value otherwise.

### GetFlexCommentUnitOk

`func (o *TenantPackage) GetFlexCommentUnitOk() (*float64, bool)`

GetFlexCommentUnitOk returns a tuple with the FlexCommentUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentUnit

`func (o *TenantPackage) SetFlexCommentUnit(v float64)`

SetFlexCommentUnit sets FlexCommentUnit field to given value.

### HasFlexCommentUnit

`func (o *TenantPackage) HasFlexCommentUnit() bool`

HasFlexCommentUnit returns a boolean if a field has been set.

### GetFlexSSOUserCostCents

`func (o *TenantPackage) GetFlexSSOUserCostCents() float64`

GetFlexSSOUserCostCents returns the FlexSSOUserCostCents field if non-nil, zero value otherwise.

### GetFlexSSOUserCostCentsOk

`func (o *TenantPackage) GetFlexSSOUserCostCentsOk() (*float64, bool)`

GetFlexSSOUserCostCentsOk returns a tuple with the FlexSSOUserCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserCostCents

`func (o *TenantPackage) SetFlexSSOUserCostCents(v float64)`

SetFlexSSOUserCostCents sets FlexSSOUserCostCents field to given value.

### HasFlexSSOUserCostCents

`func (o *TenantPackage) HasFlexSSOUserCostCents() bool`

HasFlexSSOUserCostCents returns a boolean if a field has been set.

### GetFlexSSOUserUnit

`func (o *TenantPackage) GetFlexSSOUserUnit() float64`

GetFlexSSOUserUnit returns the FlexSSOUserUnit field if non-nil, zero value otherwise.

### GetFlexSSOUserUnitOk

`func (o *TenantPackage) GetFlexSSOUserUnitOk() (*float64, bool)`

GetFlexSSOUserUnitOk returns a tuple with the FlexSSOUserUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserUnit

`func (o *TenantPackage) SetFlexSSOUserUnit(v float64)`

SetFlexSSOUserUnit sets FlexSSOUserUnit field to given value.

### HasFlexSSOUserUnit

`func (o *TenantPackage) HasFlexSSOUserUnit() bool`

HasFlexSSOUserUnit returns a boolean if a field has been set.

### GetFlexAPICreditCostCents

`func (o *TenantPackage) GetFlexAPICreditCostCents() float64`

GetFlexAPICreditCostCents returns the FlexAPICreditCostCents field if non-nil, zero value otherwise.

### GetFlexAPICreditCostCentsOk

`func (o *TenantPackage) GetFlexAPICreditCostCentsOk() (*float64, bool)`

GetFlexAPICreditCostCentsOk returns a tuple with the FlexAPICreditCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditCostCents

`func (o *TenantPackage) SetFlexAPICreditCostCents(v float64)`

SetFlexAPICreditCostCents sets FlexAPICreditCostCents field to given value.

### HasFlexAPICreditCostCents

`func (o *TenantPackage) HasFlexAPICreditCostCents() bool`

HasFlexAPICreditCostCents returns a boolean if a field has been set.

### GetFlexAPICreditUnit

`func (o *TenantPackage) GetFlexAPICreditUnit() float64`

GetFlexAPICreditUnit returns the FlexAPICreditUnit field if non-nil, zero value otherwise.

### GetFlexAPICreditUnitOk

`func (o *TenantPackage) GetFlexAPICreditUnitOk() (*float64, bool)`

GetFlexAPICreditUnitOk returns a tuple with the FlexAPICreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditUnit

`func (o *TenantPackage) SetFlexAPICreditUnit(v float64)`

SetFlexAPICreditUnit sets FlexAPICreditUnit field to given value.

### HasFlexAPICreditUnit

`func (o *TenantPackage) HasFlexAPICreditUnit() bool`

HasFlexAPICreditUnit returns a boolean if a field has been set.

### GetFlexSmallWidgetsCreditCostCents

`func (o *TenantPackage) GetFlexSmallWidgetsCreditCostCents() float64`

GetFlexSmallWidgetsCreditCostCents returns the FlexSmallWidgetsCreditCostCents field if non-nil, zero value otherwise.

### GetFlexSmallWidgetsCreditCostCentsOk

`func (o *TenantPackage) GetFlexSmallWidgetsCreditCostCentsOk() (*float64, bool)`

GetFlexSmallWidgetsCreditCostCentsOk returns a tuple with the FlexSmallWidgetsCreditCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSmallWidgetsCreditCostCents

`func (o *TenantPackage) SetFlexSmallWidgetsCreditCostCents(v float64)`

SetFlexSmallWidgetsCreditCostCents sets FlexSmallWidgetsCreditCostCents field to given value.

### HasFlexSmallWidgetsCreditCostCents

`func (o *TenantPackage) HasFlexSmallWidgetsCreditCostCents() bool`

HasFlexSmallWidgetsCreditCostCents returns a boolean if a field has been set.

### GetFlexSmallWidgetsCreditUnit

`func (o *TenantPackage) GetFlexSmallWidgetsCreditUnit() float64`

GetFlexSmallWidgetsCreditUnit returns the FlexSmallWidgetsCreditUnit field if non-nil, zero value otherwise.

### GetFlexSmallWidgetsCreditUnitOk

`func (o *TenantPackage) GetFlexSmallWidgetsCreditUnitOk() (*float64, bool)`

GetFlexSmallWidgetsCreditUnitOk returns a tuple with the FlexSmallWidgetsCreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSmallWidgetsCreditUnit

`func (o *TenantPackage) SetFlexSmallWidgetsCreditUnit(v float64)`

SetFlexSmallWidgetsCreditUnit sets FlexSmallWidgetsCreditUnit field to given value.

### HasFlexSmallWidgetsCreditUnit

`func (o *TenantPackage) HasFlexSmallWidgetsCreditUnit() bool`

HasFlexSmallWidgetsCreditUnit returns a boolean if a field has been set.

### GetFlexModeratorCostCents

`func (o *TenantPackage) GetFlexModeratorCostCents() float64`

GetFlexModeratorCostCents returns the FlexModeratorCostCents field if non-nil, zero value otherwise.

### GetFlexModeratorCostCentsOk

`func (o *TenantPackage) GetFlexModeratorCostCentsOk() (*float64, bool)`

GetFlexModeratorCostCentsOk returns a tuple with the FlexModeratorCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorCostCents

`func (o *TenantPackage) SetFlexModeratorCostCents(v float64)`

SetFlexModeratorCostCents sets FlexModeratorCostCents field to given value.

### HasFlexModeratorCostCents

`func (o *TenantPackage) HasFlexModeratorCostCents() bool`

HasFlexModeratorCostCents returns a boolean if a field has been set.

### GetFlexModeratorUnit

`func (o *TenantPackage) GetFlexModeratorUnit() float64`

GetFlexModeratorUnit returns the FlexModeratorUnit field if non-nil, zero value otherwise.

### GetFlexModeratorUnitOk

`func (o *TenantPackage) GetFlexModeratorUnitOk() (*float64, bool)`

GetFlexModeratorUnitOk returns a tuple with the FlexModeratorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorUnit

`func (o *TenantPackage) SetFlexModeratorUnit(v float64)`

SetFlexModeratorUnit sets FlexModeratorUnit field to given value.

### HasFlexModeratorUnit

`func (o *TenantPackage) HasFlexModeratorUnit() bool`

HasFlexModeratorUnit returns a boolean if a field has been set.

### GetFlexAdminCostCents

`func (o *TenantPackage) GetFlexAdminCostCents() float64`

GetFlexAdminCostCents returns the FlexAdminCostCents field if non-nil, zero value otherwise.

### GetFlexAdminCostCentsOk

`func (o *TenantPackage) GetFlexAdminCostCentsOk() (*float64, bool)`

GetFlexAdminCostCentsOk returns a tuple with the FlexAdminCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminCostCents

`func (o *TenantPackage) SetFlexAdminCostCents(v float64)`

SetFlexAdminCostCents sets FlexAdminCostCents field to given value.

### HasFlexAdminCostCents

`func (o *TenantPackage) HasFlexAdminCostCents() bool`

HasFlexAdminCostCents returns a boolean if a field has been set.

### GetFlexAdminUnit

`func (o *TenantPackage) GetFlexAdminUnit() float64`

GetFlexAdminUnit returns the FlexAdminUnit field if non-nil, zero value otherwise.

### GetFlexAdminUnitOk

`func (o *TenantPackage) GetFlexAdminUnitOk() (*float64, bool)`

GetFlexAdminUnitOk returns a tuple with the FlexAdminUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminUnit

`func (o *TenantPackage) SetFlexAdminUnit(v float64)`

SetFlexAdminUnit sets FlexAdminUnit field to given value.

### HasFlexAdminUnit

`func (o *TenantPackage) HasFlexAdminUnit() bool`

HasFlexAdminUnit returns a boolean if a field has been set.

### GetFlexDomainCostCents

`func (o *TenantPackage) GetFlexDomainCostCents() float64`

GetFlexDomainCostCents returns the FlexDomainCostCents field if non-nil, zero value otherwise.

### GetFlexDomainCostCentsOk

`func (o *TenantPackage) GetFlexDomainCostCentsOk() (*float64, bool)`

GetFlexDomainCostCentsOk returns a tuple with the FlexDomainCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainCostCents

`func (o *TenantPackage) SetFlexDomainCostCents(v float64)`

SetFlexDomainCostCents sets FlexDomainCostCents field to given value.

### HasFlexDomainCostCents

`func (o *TenantPackage) HasFlexDomainCostCents() bool`

HasFlexDomainCostCents returns a boolean if a field has been set.

### GetFlexDomainUnit

`func (o *TenantPackage) GetFlexDomainUnit() float64`

GetFlexDomainUnit returns the FlexDomainUnit field if non-nil, zero value otherwise.

### GetFlexDomainUnitOk

`func (o *TenantPackage) GetFlexDomainUnitOk() (*float64, bool)`

GetFlexDomainUnitOk returns a tuple with the FlexDomainUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainUnit

`func (o *TenantPackage) SetFlexDomainUnit(v float64)`

SetFlexDomainUnit sets FlexDomainUnit field to given value.

### HasFlexDomainUnit

`func (o *TenantPackage) HasFlexDomainUnit() bool`

HasFlexDomainUnit returns a boolean if a field has been set.

### GetFlexChatGPTCostCents

`func (o *TenantPackage) GetFlexChatGPTCostCents() float64`

GetFlexChatGPTCostCents returns the FlexChatGPTCostCents field if non-nil, zero value otherwise.

### GetFlexChatGPTCostCentsOk

`func (o *TenantPackage) GetFlexChatGPTCostCentsOk() (*float64, bool)`

GetFlexChatGPTCostCentsOk returns a tuple with the FlexChatGPTCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexChatGPTCostCents

`func (o *TenantPackage) SetFlexChatGPTCostCents(v float64)`

SetFlexChatGPTCostCents sets FlexChatGPTCostCents field to given value.

### HasFlexChatGPTCostCents

`func (o *TenantPackage) HasFlexChatGPTCostCents() bool`

HasFlexChatGPTCostCents returns a boolean if a field has been set.

### GetFlexChatGPTUnit

`func (o *TenantPackage) GetFlexChatGPTUnit() float64`

GetFlexChatGPTUnit returns the FlexChatGPTUnit field if non-nil, zero value otherwise.

### GetFlexChatGPTUnitOk

`func (o *TenantPackage) GetFlexChatGPTUnitOk() (*float64, bool)`

GetFlexChatGPTUnitOk returns a tuple with the FlexChatGPTUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexChatGPTUnit

`func (o *TenantPackage) SetFlexChatGPTUnit(v float64)`

SetFlexChatGPTUnit sets FlexChatGPTUnit field to given value.

### HasFlexChatGPTUnit

`func (o *TenantPackage) HasFlexChatGPTUnit() bool`

HasFlexChatGPTUnit returns a boolean if a field has been set.

### GetFlexMinimumCostCents

`func (o *TenantPackage) GetFlexMinimumCostCents() float64`

GetFlexMinimumCostCents returns the FlexMinimumCostCents field if non-nil, zero value otherwise.

### GetFlexMinimumCostCentsOk

`func (o *TenantPackage) GetFlexMinimumCostCentsOk() (*float64, bool)`

GetFlexMinimumCostCentsOk returns a tuple with the FlexMinimumCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMinimumCostCents

`func (o *TenantPackage) SetFlexMinimumCostCents(v float64)`

SetFlexMinimumCostCents sets FlexMinimumCostCents field to given value.

### HasFlexMinimumCostCents

`func (o *TenantPackage) HasFlexMinimumCostCents() bool`

HasFlexMinimumCostCents returns a boolean if a field has been set.

### GetFlexManagedTenantCostCents

`func (o *TenantPackage) GetFlexManagedTenantCostCents() float64`

GetFlexManagedTenantCostCents returns the FlexManagedTenantCostCents field if non-nil, zero value otherwise.

### GetFlexManagedTenantCostCentsOk

`func (o *TenantPackage) GetFlexManagedTenantCostCentsOk() (*float64, bool)`

GetFlexManagedTenantCostCentsOk returns a tuple with the FlexManagedTenantCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexManagedTenantCostCents

`func (o *TenantPackage) SetFlexManagedTenantCostCents(v float64)`

SetFlexManagedTenantCostCents sets FlexManagedTenantCostCents field to given value.

### HasFlexManagedTenantCostCents

`func (o *TenantPackage) HasFlexManagedTenantCostCents() bool`

HasFlexManagedTenantCostCents returns a boolean if a field has been set.

### GetFlexSSOAdminCostCents

`func (o *TenantPackage) GetFlexSSOAdminCostCents() float64`

GetFlexSSOAdminCostCents returns the FlexSSOAdminCostCents field if non-nil, zero value otherwise.

### GetFlexSSOAdminCostCentsOk

`func (o *TenantPackage) GetFlexSSOAdminCostCentsOk() (*float64, bool)`

GetFlexSSOAdminCostCentsOk returns a tuple with the FlexSSOAdminCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOAdminCostCents

`func (o *TenantPackage) SetFlexSSOAdminCostCents(v float64)`

SetFlexSSOAdminCostCents sets FlexSSOAdminCostCents field to given value.

### HasFlexSSOAdminCostCents

`func (o *TenantPackage) HasFlexSSOAdminCostCents() bool`

HasFlexSSOAdminCostCents returns a boolean if a field has been set.

### GetFlexSSOAdminUnit

`func (o *TenantPackage) GetFlexSSOAdminUnit() float64`

GetFlexSSOAdminUnit returns the FlexSSOAdminUnit field if non-nil, zero value otherwise.

### GetFlexSSOAdminUnitOk

`func (o *TenantPackage) GetFlexSSOAdminUnitOk() (*float64, bool)`

GetFlexSSOAdminUnitOk returns a tuple with the FlexSSOAdminUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOAdminUnit

`func (o *TenantPackage) SetFlexSSOAdminUnit(v float64)`

SetFlexSSOAdminUnit sets FlexSSOAdminUnit field to given value.

### HasFlexSSOAdminUnit

`func (o *TenantPackage) HasFlexSSOAdminUnit() bool`

HasFlexSSOAdminUnit returns a boolean if a field has been set.

### GetFlexSSOModeratorCostCents

`func (o *TenantPackage) GetFlexSSOModeratorCostCents() float64`

GetFlexSSOModeratorCostCents returns the FlexSSOModeratorCostCents field if non-nil, zero value otherwise.

### GetFlexSSOModeratorCostCentsOk

`func (o *TenantPackage) GetFlexSSOModeratorCostCentsOk() (*float64, bool)`

GetFlexSSOModeratorCostCentsOk returns a tuple with the FlexSSOModeratorCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOModeratorCostCents

`func (o *TenantPackage) SetFlexSSOModeratorCostCents(v float64)`

SetFlexSSOModeratorCostCents sets FlexSSOModeratorCostCents field to given value.

### HasFlexSSOModeratorCostCents

`func (o *TenantPackage) HasFlexSSOModeratorCostCents() bool`

HasFlexSSOModeratorCostCents returns a boolean if a field has been set.

### GetFlexSSOModeratorUnit

`func (o *TenantPackage) GetFlexSSOModeratorUnit() float64`

GetFlexSSOModeratorUnit returns the FlexSSOModeratorUnit field if non-nil, zero value otherwise.

### GetFlexSSOModeratorUnitOk

`func (o *TenantPackage) GetFlexSSOModeratorUnitOk() (*float64, bool)`

GetFlexSSOModeratorUnitOk returns a tuple with the FlexSSOModeratorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOModeratorUnit

`func (o *TenantPackage) SetFlexSSOModeratorUnit(v float64)`

SetFlexSSOModeratorUnit sets FlexSSOModeratorUnit field to given value.

### HasFlexSSOModeratorUnit

`func (o *TenantPackage) HasFlexSSOModeratorUnit() bool`

HasFlexSSOModeratorUnit returns a boolean if a field has been set.

### GetIsSSOBillingMonthlyActiveUsers

`func (o *TenantPackage) GetIsSSOBillingMonthlyActiveUsers() bool`

GetIsSSOBillingMonthlyActiveUsers returns the IsSSOBillingMonthlyActiveUsers field if non-nil, zero value otherwise.

### GetIsSSOBillingMonthlyActiveUsersOk

`func (o *TenantPackage) GetIsSSOBillingMonthlyActiveUsersOk() (*bool, bool)`

GetIsSSOBillingMonthlyActiveUsersOk returns a tuple with the IsSSOBillingMonthlyActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSSOBillingMonthlyActiveUsers

`func (o *TenantPackage) SetIsSSOBillingMonthlyActiveUsers(v bool)`

SetIsSSOBillingMonthlyActiveUsers sets IsSSOBillingMonthlyActiveUsers field to given value.

### HasIsSSOBillingMonthlyActiveUsers

`func (o *TenantPackage) HasIsSSOBillingMonthlyActiveUsers() bool`

HasIsSSOBillingMonthlyActiveUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


