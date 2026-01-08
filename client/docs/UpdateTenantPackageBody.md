# UpdateTenantPackageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**MonthlyCostUSD** | Pointer to **float64** |  | [optional] 
**YearlyCostUSD** | Pointer to **float64** |  | [optional] 
**MaxMonthlyPageLoads** | Pointer to **float64** |  | [optional] 
**MaxMonthlyAPICredits** | Pointer to **float64** |  | [optional] 
**MaxMonthlyComments** | Pointer to **float64** |  | [optional] 
**MaxConcurrentUsers** | Pointer to **float64** |  | [optional] 
**MaxTenantUsers** | Pointer to **float64** |  | [optional] 
**MaxSSOUsers** | Pointer to **float64** |  | [optional] 
**MaxModerators** | Pointer to **float64** |  | [optional] 
**MaxDomains** | Pointer to **float64** |  | [optional] 
**HasDebranding** | Pointer to **bool** |  | [optional] 
**HasWhiteLabeling** | Pointer to **bool** |  | [optional] 
**ForWhoText** | Pointer to **string** |  | [optional] 
**FeatureTaglines** | Pointer to **[]string** |  | [optional] 
**HasFlexPricing** | Pointer to **bool** |  | [optional] 
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

### NewUpdateTenantPackageBody

`func NewUpdateTenantPackageBody() *UpdateTenantPackageBody`

NewUpdateTenantPackageBody instantiates a new UpdateTenantPackageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantPackageBodyWithDefaults

`func NewUpdateTenantPackageBodyWithDefaults() *UpdateTenantPackageBody`

NewUpdateTenantPackageBodyWithDefaults instantiates a new UpdateTenantPackageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTenantPackageBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenantPackageBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenantPackageBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTenantPackageBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMonthlyCostUSD

`func (o *UpdateTenantPackageBody) GetMonthlyCostUSD() float64`

GetMonthlyCostUSD returns the MonthlyCostUSD field if non-nil, zero value otherwise.

### GetMonthlyCostUSDOk

`func (o *UpdateTenantPackageBody) GetMonthlyCostUSDOk() (*float64, bool)`

GetMonthlyCostUSDOk returns a tuple with the MonthlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCostUSD

`func (o *UpdateTenantPackageBody) SetMonthlyCostUSD(v float64)`

SetMonthlyCostUSD sets MonthlyCostUSD field to given value.

### HasMonthlyCostUSD

`func (o *UpdateTenantPackageBody) HasMonthlyCostUSD() bool`

HasMonthlyCostUSD returns a boolean if a field has been set.

### GetYearlyCostUSD

`func (o *UpdateTenantPackageBody) GetYearlyCostUSD() float64`

GetYearlyCostUSD returns the YearlyCostUSD field if non-nil, zero value otherwise.

### GetYearlyCostUSDOk

`func (o *UpdateTenantPackageBody) GetYearlyCostUSDOk() (*float64, bool)`

GetYearlyCostUSDOk returns a tuple with the YearlyCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCostUSD

`func (o *UpdateTenantPackageBody) SetYearlyCostUSD(v float64)`

SetYearlyCostUSD sets YearlyCostUSD field to given value.

### HasYearlyCostUSD

`func (o *UpdateTenantPackageBody) HasYearlyCostUSD() bool`

HasYearlyCostUSD returns a boolean if a field has been set.

### GetMaxMonthlyPageLoads

`func (o *UpdateTenantPackageBody) GetMaxMonthlyPageLoads() float64`

GetMaxMonthlyPageLoads returns the MaxMonthlyPageLoads field if non-nil, zero value otherwise.

### GetMaxMonthlyPageLoadsOk

`func (o *UpdateTenantPackageBody) GetMaxMonthlyPageLoadsOk() (*float64, bool)`

GetMaxMonthlyPageLoadsOk returns a tuple with the MaxMonthlyPageLoads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyPageLoads

`func (o *UpdateTenantPackageBody) SetMaxMonthlyPageLoads(v float64)`

SetMaxMonthlyPageLoads sets MaxMonthlyPageLoads field to given value.

### HasMaxMonthlyPageLoads

`func (o *UpdateTenantPackageBody) HasMaxMonthlyPageLoads() bool`

HasMaxMonthlyPageLoads returns a boolean if a field has been set.

### GetMaxMonthlyAPICredits

`func (o *UpdateTenantPackageBody) GetMaxMonthlyAPICredits() float64`

GetMaxMonthlyAPICredits returns the MaxMonthlyAPICredits field if non-nil, zero value otherwise.

### GetMaxMonthlyAPICreditsOk

`func (o *UpdateTenantPackageBody) GetMaxMonthlyAPICreditsOk() (*float64, bool)`

GetMaxMonthlyAPICreditsOk returns a tuple with the MaxMonthlyAPICredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyAPICredits

`func (o *UpdateTenantPackageBody) SetMaxMonthlyAPICredits(v float64)`

SetMaxMonthlyAPICredits sets MaxMonthlyAPICredits field to given value.

### HasMaxMonthlyAPICredits

`func (o *UpdateTenantPackageBody) HasMaxMonthlyAPICredits() bool`

HasMaxMonthlyAPICredits returns a boolean if a field has been set.

### GetMaxMonthlyComments

`func (o *UpdateTenantPackageBody) GetMaxMonthlyComments() float64`

GetMaxMonthlyComments returns the MaxMonthlyComments field if non-nil, zero value otherwise.

### GetMaxMonthlyCommentsOk

`func (o *UpdateTenantPackageBody) GetMaxMonthlyCommentsOk() (*float64, bool)`

GetMaxMonthlyCommentsOk returns a tuple with the MaxMonthlyComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMonthlyComments

`func (o *UpdateTenantPackageBody) SetMaxMonthlyComments(v float64)`

SetMaxMonthlyComments sets MaxMonthlyComments field to given value.

### HasMaxMonthlyComments

`func (o *UpdateTenantPackageBody) HasMaxMonthlyComments() bool`

HasMaxMonthlyComments returns a boolean if a field has been set.

### GetMaxConcurrentUsers

`func (o *UpdateTenantPackageBody) GetMaxConcurrentUsers() float64`

GetMaxConcurrentUsers returns the MaxConcurrentUsers field if non-nil, zero value otherwise.

### GetMaxConcurrentUsersOk

`func (o *UpdateTenantPackageBody) GetMaxConcurrentUsersOk() (*float64, bool)`

GetMaxConcurrentUsersOk returns a tuple with the MaxConcurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUsers

`func (o *UpdateTenantPackageBody) SetMaxConcurrentUsers(v float64)`

SetMaxConcurrentUsers sets MaxConcurrentUsers field to given value.

### HasMaxConcurrentUsers

`func (o *UpdateTenantPackageBody) HasMaxConcurrentUsers() bool`

HasMaxConcurrentUsers returns a boolean if a field has been set.

### GetMaxTenantUsers

`func (o *UpdateTenantPackageBody) GetMaxTenantUsers() float64`

GetMaxTenantUsers returns the MaxTenantUsers field if non-nil, zero value otherwise.

### GetMaxTenantUsersOk

`func (o *UpdateTenantPackageBody) GetMaxTenantUsersOk() (*float64, bool)`

GetMaxTenantUsersOk returns a tuple with the MaxTenantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTenantUsers

`func (o *UpdateTenantPackageBody) SetMaxTenantUsers(v float64)`

SetMaxTenantUsers sets MaxTenantUsers field to given value.

### HasMaxTenantUsers

`func (o *UpdateTenantPackageBody) HasMaxTenantUsers() bool`

HasMaxTenantUsers returns a boolean if a field has been set.

### GetMaxSSOUsers

`func (o *UpdateTenantPackageBody) GetMaxSSOUsers() float64`

GetMaxSSOUsers returns the MaxSSOUsers field if non-nil, zero value otherwise.

### GetMaxSSOUsersOk

`func (o *UpdateTenantPackageBody) GetMaxSSOUsersOk() (*float64, bool)`

GetMaxSSOUsersOk returns a tuple with the MaxSSOUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSSOUsers

`func (o *UpdateTenantPackageBody) SetMaxSSOUsers(v float64)`

SetMaxSSOUsers sets MaxSSOUsers field to given value.

### HasMaxSSOUsers

`func (o *UpdateTenantPackageBody) HasMaxSSOUsers() bool`

HasMaxSSOUsers returns a boolean if a field has been set.

### GetMaxModerators

`func (o *UpdateTenantPackageBody) GetMaxModerators() float64`

GetMaxModerators returns the MaxModerators field if non-nil, zero value otherwise.

### GetMaxModeratorsOk

`func (o *UpdateTenantPackageBody) GetMaxModeratorsOk() (*float64, bool)`

GetMaxModeratorsOk returns a tuple with the MaxModerators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxModerators

`func (o *UpdateTenantPackageBody) SetMaxModerators(v float64)`

SetMaxModerators sets MaxModerators field to given value.

### HasMaxModerators

`func (o *UpdateTenantPackageBody) HasMaxModerators() bool`

HasMaxModerators returns a boolean if a field has been set.

### GetMaxDomains

`func (o *UpdateTenantPackageBody) GetMaxDomains() float64`

GetMaxDomains returns the MaxDomains field if non-nil, zero value otherwise.

### GetMaxDomainsOk

`func (o *UpdateTenantPackageBody) GetMaxDomainsOk() (*float64, bool)`

GetMaxDomainsOk returns a tuple with the MaxDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDomains

`func (o *UpdateTenantPackageBody) SetMaxDomains(v float64)`

SetMaxDomains sets MaxDomains field to given value.

### HasMaxDomains

`func (o *UpdateTenantPackageBody) HasMaxDomains() bool`

HasMaxDomains returns a boolean if a field has been set.

### GetHasDebranding

`func (o *UpdateTenantPackageBody) GetHasDebranding() bool`

GetHasDebranding returns the HasDebranding field if non-nil, zero value otherwise.

### GetHasDebrandingOk

`func (o *UpdateTenantPackageBody) GetHasDebrandingOk() (*bool, bool)`

GetHasDebrandingOk returns a tuple with the HasDebranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDebranding

`func (o *UpdateTenantPackageBody) SetHasDebranding(v bool)`

SetHasDebranding sets HasDebranding field to given value.

### HasHasDebranding

`func (o *UpdateTenantPackageBody) HasHasDebranding() bool`

HasHasDebranding returns a boolean if a field has been set.

### GetHasWhiteLabeling

`func (o *UpdateTenantPackageBody) GetHasWhiteLabeling() bool`

GetHasWhiteLabeling returns the HasWhiteLabeling field if non-nil, zero value otherwise.

### GetHasWhiteLabelingOk

`func (o *UpdateTenantPackageBody) GetHasWhiteLabelingOk() (*bool, bool)`

GetHasWhiteLabelingOk returns a tuple with the HasWhiteLabeling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWhiteLabeling

`func (o *UpdateTenantPackageBody) SetHasWhiteLabeling(v bool)`

SetHasWhiteLabeling sets HasWhiteLabeling field to given value.

### HasHasWhiteLabeling

`func (o *UpdateTenantPackageBody) HasHasWhiteLabeling() bool`

HasHasWhiteLabeling returns a boolean if a field has been set.

### GetForWhoText

`func (o *UpdateTenantPackageBody) GetForWhoText() string`

GetForWhoText returns the ForWhoText field if non-nil, zero value otherwise.

### GetForWhoTextOk

`func (o *UpdateTenantPackageBody) GetForWhoTextOk() (*string, bool)`

GetForWhoTextOk returns a tuple with the ForWhoText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForWhoText

`func (o *UpdateTenantPackageBody) SetForWhoText(v string)`

SetForWhoText sets ForWhoText field to given value.

### HasForWhoText

`func (o *UpdateTenantPackageBody) HasForWhoText() bool`

HasForWhoText returns a boolean if a field has been set.

### GetFeatureTaglines

`func (o *UpdateTenantPackageBody) GetFeatureTaglines() []string`

GetFeatureTaglines returns the FeatureTaglines field if non-nil, zero value otherwise.

### GetFeatureTaglinesOk

`func (o *UpdateTenantPackageBody) GetFeatureTaglinesOk() (*[]string, bool)`

GetFeatureTaglinesOk returns a tuple with the FeatureTaglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTaglines

`func (o *UpdateTenantPackageBody) SetFeatureTaglines(v []string)`

SetFeatureTaglines sets FeatureTaglines field to given value.

### HasFeatureTaglines

`func (o *UpdateTenantPackageBody) HasFeatureTaglines() bool`

HasFeatureTaglines returns a boolean if a field has been set.

### GetHasFlexPricing

`func (o *UpdateTenantPackageBody) GetHasFlexPricing() bool`

GetHasFlexPricing returns the HasFlexPricing field if non-nil, zero value otherwise.

### GetHasFlexPricingOk

`func (o *UpdateTenantPackageBody) GetHasFlexPricingOk() (*bool, bool)`

GetHasFlexPricingOk returns a tuple with the HasFlexPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFlexPricing

`func (o *UpdateTenantPackageBody) SetHasFlexPricing(v bool)`

SetHasFlexPricing sets HasFlexPricing field to given value.

### HasHasFlexPricing

`func (o *UpdateTenantPackageBody) HasHasFlexPricing() bool`

HasHasFlexPricing returns a boolean if a field has been set.

### GetFlexPageLoadCostCents

`func (o *UpdateTenantPackageBody) GetFlexPageLoadCostCents() float64`

GetFlexPageLoadCostCents returns the FlexPageLoadCostCents field if non-nil, zero value otherwise.

### GetFlexPageLoadCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexPageLoadCostCentsOk() (*float64, bool)`

GetFlexPageLoadCostCentsOk returns a tuple with the FlexPageLoadCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadCostCents

`func (o *UpdateTenantPackageBody) SetFlexPageLoadCostCents(v float64)`

SetFlexPageLoadCostCents sets FlexPageLoadCostCents field to given value.

### HasFlexPageLoadCostCents

`func (o *UpdateTenantPackageBody) HasFlexPageLoadCostCents() bool`

HasFlexPageLoadCostCents returns a boolean if a field has been set.

### GetFlexPageLoadUnit

`func (o *UpdateTenantPackageBody) GetFlexPageLoadUnit() float64`

GetFlexPageLoadUnit returns the FlexPageLoadUnit field if non-nil, zero value otherwise.

### GetFlexPageLoadUnitOk

`func (o *UpdateTenantPackageBody) GetFlexPageLoadUnitOk() (*float64, bool)`

GetFlexPageLoadUnitOk returns a tuple with the FlexPageLoadUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexPageLoadUnit

`func (o *UpdateTenantPackageBody) SetFlexPageLoadUnit(v float64)`

SetFlexPageLoadUnit sets FlexPageLoadUnit field to given value.

### HasFlexPageLoadUnit

`func (o *UpdateTenantPackageBody) HasFlexPageLoadUnit() bool`

HasFlexPageLoadUnit returns a boolean if a field has been set.

### GetFlexCommentCostCents

`func (o *UpdateTenantPackageBody) GetFlexCommentCostCents() float64`

GetFlexCommentCostCents returns the FlexCommentCostCents field if non-nil, zero value otherwise.

### GetFlexCommentCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexCommentCostCentsOk() (*float64, bool)`

GetFlexCommentCostCentsOk returns a tuple with the FlexCommentCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentCostCents

`func (o *UpdateTenantPackageBody) SetFlexCommentCostCents(v float64)`

SetFlexCommentCostCents sets FlexCommentCostCents field to given value.

### HasFlexCommentCostCents

`func (o *UpdateTenantPackageBody) HasFlexCommentCostCents() bool`

HasFlexCommentCostCents returns a boolean if a field has been set.

### GetFlexCommentUnit

`func (o *UpdateTenantPackageBody) GetFlexCommentUnit() float64`

GetFlexCommentUnit returns the FlexCommentUnit field if non-nil, zero value otherwise.

### GetFlexCommentUnitOk

`func (o *UpdateTenantPackageBody) GetFlexCommentUnitOk() (*float64, bool)`

GetFlexCommentUnitOk returns a tuple with the FlexCommentUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexCommentUnit

`func (o *UpdateTenantPackageBody) SetFlexCommentUnit(v float64)`

SetFlexCommentUnit sets FlexCommentUnit field to given value.

### HasFlexCommentUnit

`func (o *UpdateTenantPackageBody) HasFlexCommentUnit() bool`

HasFlexCommentUnit returns a boolean if a field has been set.

### GetFlexSSOUserCostCents

`func (o *UpdateTenantPackageBody) GetFlexSSOUserCostCents() float64`

GetFlexSSOUserCostCents returns the FlexSSOUserCostCents field if non-nil, zero value otherwise.

### GetFlexSSOUserCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexSSOUserCostCentsOk() (*float64, bool)`

GetFlexSSOUserCostCentsOk returns a tuple with the FlexSSOUserCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserCostCents

`func (o *UpdateTenantPackageBody) SetFlexSSOUserCostCents(v float64)`

SetFlexSSOUserCostCents sets FlexSSOUserCostCents field to given value.

### HasFlexSSOUserCostCents

`func (o *UpdateTenantPackageBody) HasFlexSSOUserCostCents() bool`

HasFlexSSOUserCostCents returns a boolean if a field has been set.

### GetFlexSSOUserUnit

`func (o *UpdateTenantPackageBody) GetFlexSSOUserUnit() float64`

GetFlexSSOUserUnit returns the FlexSSOUserUnit field if non-nil, zero value otherwise.

### GetFlexSSOUserUnitOk

`func (o *UpdateTenantPackageBody) GetFlexSSOUserUnitOk() (*float64, bool)`

GetFlexSSOUserUnitOk returns a tuple with the FlexSSOUserUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexSSOUserUnit

`func (o *UpdateTenantPackageBody) SetFlexSSOUserUnit(v float64)`

SetFlexSSOUserUnit sets FlexSSOUserUnit field to given value.

### HasFlexSSOUserUnit

`func (o *UpdateTenantPackageBody) HasFlexSSOUserUnit() bool`

HasFlexSSOUserUnit returns a boolean if a field has been set.

### GetFlexAPICreditCostCents

`func (o *UpdateTenantPackageBody) GetFlexAPICreditCostCents() float64`

GetFlexAPICreditCostCents returns the FlexAPICreditCostCents field if non-nil, zero value otherwise.

### GetFlexAPICreditCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexAPICreditCostCentsOk() (*float64, bool)`

GetFlexAPICreditCostCentsOk returns a tuple with the FlexAPICreditCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditCostCents

`func (o *UpdateTenantPackageBody) SetFlexAPICreditCostCents(v float64)`

SetFlexAPICreditCostCents sets FlexAPICreditCostCents field to given value.

### HasFlexAPICreditCostCents

`func (o *UpdateTenantPackageBody) HasFlexAPICreditCostCents() bool`

HasFlexAPICreditCostCents returns a boolean if a field has been set.

### GetFlexAPICreditUnit

`func (o *UpdateTenantPackageBody) GetFlexAPICreditUnit() float64`

GetFlexAPICreditUnit returns the FlexAPICreditUnit field if non-nil, zero value otherwise.

### GetFlexAPICreditUnitOk

`func (o *UpdateTenantPackageBody) GetFlexAPICreditUnitOk() (*float64, bool)`

GetFlexAPICreditUnitOk returns a tuple with the FlexAPICreditUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAPICreditUnit

`func (o *UpdateTenantPackageBody) SetFlexAPICreditUnit(v float64)`

SetFlexAPICreditUnit sets FlexAPICreditUnit field to given value.

### HasFlexAPICreditUnit

`func (o *UpdateTenantPackageBody) HasFlexAPICreditUnit() bool`

HasFlexAPICreditUnit returns a boolean if a field has been set.

### GetFlexModeratorCostCents

`func (o *UpdateTenantPackageBody) GetFlexModeratorCostCents() float64`

GetFlexModeratorCostCents returns the FlexModeratorCostCents field if non-nil, zero value otherwise.

### GetFlexModeratorCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexModeratorCostCentsOk() (*float64, bool)`

GetFlexModeratorCostCentsOk returns a tuple with the FlexModeratorCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorCostCents

`func (o *UpdateTenantPackageBody) SetFlexModeratorCostCents(v float64)`

SetFlexModeratorCostCents sets FlexModeratorCostCents field to given value.

### HasFlexModeratorCostCents

`func (o *UpdateTenantPackageBody) HasFlexModeratorCostCents() bool`

HasFlexModeratorCostCents returns a boolean if a field has been set.

### GetFlexModeratorUnit

`func (o *UpdateTenantPackageBody) GetFlexModeratorUnit() float64`

GetFlexModeratorUnit returns the FlexModeratorUnit field if non-nil, zero value otherwise.

### GetFlexModeratorUnitOk

`func (o *UpdateTenantPackageBody) GetFlexModeratorUnitOk() (*float64, bool)`

GetFlexModeratorUnitOk returns a tuple with the FlexModeratorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexModeratorUnit

`func (o *UpdateTenantPackageBody) SetFlexModeratorUnit(v float64)`

SetFlexModeratorUnit sets FlexModeratorUnit field to given value.

### HasFlexModeratorUnit

`func (o *UpdateTenantPackageBody) HasFlexModeratorUnit() bool`

HasFlexModeratorUnit returns a boolean if a field has been set.

### GetFlexAdminCostCents

`func (o *UpdateTenantPackageBody) GetFlexAdminCostCents() float64`

GetFlexAdminCostCents returns the FlexAdminCostCents field if non-nil, zero value otherwise.

### GetFlexAdminCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexAdminCostCentsOk() (*float64, bool)`

GetFlexAdminCostCentsOk returns a tuple with the FlexAdminCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminCostCents

`func (o *UpdateTenantPackageBody) SetFlexAdminCostCents(v float64)`

SetFlexAdminCostCents sets FlexAdminCostCents field to given value.

### HasFlexAdminCostCents

`func (o *UpdateTenantPackageBody) HasFlexAdminCostCents() bool`

HasFlexAdminCostCents returns a boolean if a field has been set.

### GetFlexAdminUnit

`func (o *UpdateTenantPackageBody) GetFlexAdminUnit() float64`

GetFlexAdminUnit returns the FlexAdminUnit field if non-nil, zero value otherwise.

### GetFlexAdminUnitOk

`func (o *UpdateTenantPackageBody) GetFlexAdminUnitOk() (*float64, bool)`

GetFlexAdminUnitOk returns a tuple with the FlexAdminUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexAdminUnit

`func (o *UpdateTenantPackageBody) SetFlexAdminUnit(v float64)`

SetFlexAdminUnit sets FlexAdminUnit field to given value.

### HasFlexAdminUnit

`func (o *UpdateTenantPackageBody) HasFlexAdminUnit() bool`

HasFlexAdminUnit returns a boolean if a field has been set.

### GetFlexDomainCostCents

`func (o *UpdateTenantPackageBody) GetFlexDomainCostCents() float64`

GetFlexDomainCostCents returns the FlexDomainCostCents field if non-nil, zero value otherwise.

### GetFlexDomainCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexDomainCostCentsOk() (*float64, bool)`

GetFlexDomainCostCentsOk returns a tuple with the FlexDomainCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainCostCents

`func (o *UpdateTenantPackageBody) SetFlexDomainCostCents(v float64)`

SetFlexDomainCostCents sets FlexDomainCostCents field to given value.

### HasFlexDomainCostCents

`func (o *UpdateTenantPackageBody) HasFlexDomainCostCents() bool`

HasFlexDomainCostCents returns a boolean if a field has been set.

### GetFlexDomainUnit

`func (o *UpdateTenantPackageBody) GetFlexDomainUnit() float64`

GetFlexDomainUnit returns the FlexDomainUnit field if non-nil, zero value otherwise.

### GetFlexDomainUnitOk

`func (o *UpdateTenantPackageBody) GetFlexDomainUnitOk() (*float64, bool)`

GetFlexDomainUnitOk returns a tuple with the FlexDomainUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexDomainUnit

`func (o *UpdateTenantPackageBody) SetFlexDomainUnit(v float64)`

SetFlexDomainUnit sets FlexDomainUnit field to given value.

### HasFlexDomainUnit

`func (o *UpdateTenantPackageBody) HasFlexDomainUnit() bool`

HasFlexDomainUnit returns a boolean if a field has been set.

### GetFlexMinimumCostCents

`func (o *UpdateTenantPackageBody) GetFlexMinimumCostCents() float64`

GetFlexMinimumCostCents returns the FlexMinimumCostCents field if non-nil, zero value otherwise.

### GetFlexMinimumCostCentsOk

`func (o *UpdateTenantPackageBody) GetFlexMinimumCostCentsOk() (*float64, bool)`

GetFlexMinimumCostCentsOk returns a tuple with the FlexMinimumCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexMinimumCostCents

`func (o *UpdateTenantPackageBody) SetFlexMinimumCostCents(v float64)`

SetFlexMinimumCostCents sets FlexMinimumCostCents field to given value.

### HasFlexMinimumCostCents

`func (o *UpdateTenantPackageBody) HasFlexMinimumCostCents() bool`

HasFlexMinimumCostCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


