# APITenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**SignUpDate** | **float64** |  | 
**PackageId** | **string** |  | 
**PaymentFrequency** | **float64** |  | 
**BillingInfoValid** | **bool** |  | 
**BillingHandledExternally** | Pointer to **bool** |  | [optional] 
**CreatedBy** | **string** |  | 
**IsSetup** | **bool** |  | 
**DomainConfiguration** | [**[]APIDomainConfiguration**](APIDomainConfiguration.md) |  | 
**BillingInfo** | Pointer to [**BillingInfo**](BillingInfo.md) |  | [optional] 
**StripeCustomerId** | Pointer to **string** |  | [optional] 
**StripeSubscriptionId** | Pointer to **string** |  | [optional] 
**StripePlanId** | Pointer to **string** |  | [optional] 
**EnableProfanityFilter** | **bool** |  | 
**EnableSpamFilter** | **bool** |  | 
**LastBillingIssueReminderDate** | Pointer to **time.Time** |  | [optional] 
**RemoveUnverifiedComments** | Pointer to **bool** |  | [optional] 
**UnverifiedCommentsTTLms** | Pointer to **NullableFloat64** |  | [optional] 
**CommentsRequireApproval** | Pointer to **bool** |  | [optional] 
**AutoApproveCommentOnVerification** | Pointer to **bool** |  | [optional] 
**SendProfaneToSpam** | Pointer to **bool** |  | [optional] 
**HasFlexPricing** | Pointer to **bool** |  | [optional] 
**HasAuditing** | Pointer to **bool** |  | [optional] 
**FlexLastBilledAmount** | Pointer to **float64** |  | [optional] 
**DeAnonIpAddr** | Pointer to **float64** |  | [optional] 
**Meta** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewAPITenant

`func NewAPITenant(id string, name string, signUpDate float64, packageId string, paymentFrequency float64, billingInfoValid bool, createdBy string, isSetup bool, domainConfiguration []APIDomainConfiguration, enableProfanityFilter bool, enableSpamFilter bool, ) *APITenant`

NewAPITenant instantiates a new APITenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITenantWithDefaults

`func NewAPITenantWithDefaults() *APITenant`

NewAPITenantWithDefaults instantiates a new APITenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITenant) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *APITenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APITenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APITenant) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *APITenant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *APITenant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *APITenant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *APITenant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSignUpDate

`func (o *APITenant) GetSignUpDate() float64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *APITenant) GetSignUpDateOk() (*float64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *APITenant) SetSignUpDate(v float64)`

SetSignUpDate sets SignUpDate field to given value.


### GetPackageId

`func (o *APITenant) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *APITenant) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *APITenant) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPaymentFrequency

`func (o *APITenant) GetPaymentFrequency() float64`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *APITenant) GetPaymentFrequencyOk() (*float64, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *APITenant) SetPaymentFrequency(v float64)`

SetPaymentFrequency sets PaymentFrequency field to given value.


### GetBillingInfoValid

`func (o *APITenant) GetBillingInfoValid() bool`

GetBillingInfoValid returns the BillingInfoValid field if non-nil, zero value otherwise.

### GetBillingInfoValidOk

`func (o *APITenant) GetBillingInfoValidOk() (*bool, bool)`

GetBillingInfoValidOk returns a tuple with the BillingInfoValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfoValid

`func (o *APITenant) SetBillingInfoValid(v bool)`

SetBillingInfoValid sets BillingInfoValid field to given value.


### GetBillingHandledExternally

`func (o *APITenant) GetBillingHandledExternally() bool`

GetBillingHandledExternally returns the BillingHandledExternally field if non-nil, zero value otherwise.

### GetBillingHandledExternallyOk

`func (o *APITenant) GetBillingHandledExternallyOk() (*bool, bool)`

GetBillingHandledExternallyOk returns a tuple with the BillingHandledExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHandledExternally

`func (o *APITenant) SetBillingHandledExternally(v bool)`

SetBillingHandledExternally sets BillingHandledExternally field to given value.

### HasBillingHandledExternally

`func (o *APITenant) HasBillingHandledExternally() bool`

HasBillingHandledExternally returns a boolean if a field has been set.

### GetCreatedBy

`func (o *APITenant) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APITenant) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APITenant) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetIsSetup

`func (o *APITenant) GetIsSetup() bool`

GetIsSetup returns the IsSetup field if non-nil, zero value otherwise.

### GetIsSetupOk

`func (o *APITenant) GetIsSetupOk() (*bool, bool)`

GetIsSetupOk returns a tuple with the IsSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetup

`func (o *APITenant) SetIsSetup(v bool)`

SetIsSetup sets IsSetup field to given value.


### GetDomainConfiguration

`func (o *APITenant) GetDomainConfiguration() []APIDomainConfiguration`

GetDomainConfiguration returns the DomainConfiguration field if non-nil, zero value otherwise.

### GetDomainConfigurationOk

`func (o *APITenant) GetDomainConfigurationOk() (*[]APIDomainConfiguration, bool)`

GetDomainConfigurationOk returns a tuple with the DomainConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainConfiguration

`func (o *APITenant) SetDomainConfiguration(v []APIDomainConfiguration)`

SetDomainConfiguration sets DomainConfiguration field to given value.


### GetBillingInfo

`func (o *APITenant) GetBillingInfo() BillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *APITenant) GetBillingInfoOk() (*BillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *APITenant) SetBillingInfo(v BillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *APITenant) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *APITenant) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *APITenant) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *APITenant) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *APITenant) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *APITenant) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *APITenant) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *APITenant) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *APITenant) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetStripePlanId

`func (o *APITenant) GetStripePlanId() string`

GetStripePlanId returns the StripePlanId field if non-nil, zero value otherwise.

### GetStripePlanIdOk

`func (o *APITenant) GetStripePlanIdOk() (*string, bool)`

GetStripePlanIdOk returns a tuple with the StripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePlanId

`func (o *APITenant) SetStripePlanId(v string)`

SetStripePlanId sets StripePlanId field to given value.

### HasStripePlanId

`func (o *APITenant) HasStripePlanId() bool`

HasStripePlanId returns a boolean if a field has been set.

### GetEnableProfanityFilter

`func (o *APITenant) GetEnableProfanityFilter() bool`

GetEnableProfanityFilter returns the EnableProfanityFilter field if non-nil, zero value otherwise.

### GetEnableProfanityFilterOk

`func (o *APITenant) GetEnableProfanityFilterOk() (*bool, bool)`

GetEnableProfanityFilterOk returns a tuple with the EnableProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfanityFilter

`func (o *APITenant) SetEnableProfanityFilter(v bool)`

SetEnableProfanityFilter sets EnableProfanityFilter field to given value.


### GetEnableSpamFilter

`func (o *APITenant) GetEnableSpamFilter() bool`

GetEnableSpamFilter returns the EnableSpamFilter field if non-nil, zero value otherwise.

### GetEnableSpamFilterOk

`func (o *APITenant) GetEnableSpamFilterOk() (*bool, bool)`

GetEnableSpamFilterOk returns a tuple with the EnableSpamFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpamFilter

`func (o *APITenant) SetEnableSpamFilter(v bool)`

SetEnableSpamFilter sets EnableSpamFilter field to given value.


### GetLastBillingIssueReminderDate

`func (o *APITenant) GetLastBillingIssueReminderDate() time.Time`

GetLastBillingIssueReminderDate returns the LastBillingIssueReminderDate field if non-nil, zero value otherwise.

### GetLastBillingIssueReminderDateOk

`func (o *APITenant) GetLastBillingIssueReminderDateOk() (*time.Time, bool)`

GetLastBillingIssueReminderDateOk returns a tuple with the LastBillingIssueReminderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBillingIssueReminderDate

`func (o *APITenant) SetLastBillingIssueReminderDate(v time.Time)`

SetLastBillingIssueReminderDate sets LastBillingIssueReminderDate field to given value.

### HasLastBillingIssueReminderDate

`func (o *APITenant) HasLastBillingIssueReminderDate() bool`

HasLastBillingIssueReminderDate returns a boolean if a field has been set.

### GetRemoveUnverifiedComments

`func (o *APITenant) GetRemoveUnverifiedComments() bool`

GetRemoveUnverifiedComments returns the RemoveUnverifiedComments field if non-nil, zero value otherwise.

### GetRemoveUnverifiedCommentsOk

`func (o *APITenant) GetRemoveUnverifiedCommentsOk() (*bool, bool)`

GetRemoveUnverifiedCommentsOk returns a tuple with the RemoveUnverifiedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUnverifiedComments

`func (o *APITenant) SetRemoveUnverifiedComments(v bool)`

SetRemoveUnverifiedComments sets RemoveUnverifiedComments field to given value.

### HasRemoveUnverifiedComments

`func (o *APITenant) HasRemoveUnverifiedComments() bool`

HasRemoveUnverifiedComments returns a boolean if a field has been set.

### GetUnverifiedCommentsTTLms

`func (o *APITenant) GetUnverifiedCommentsTTLms() float64`

GetUnverifiedCommentsTTLms returns the UnverifiedCommentsTTLms field if non-nil, zero value otherwise.

### GetUnverifiedCommentsTTLmsOk

`func (o *APITenant) GetUnverifiedCommentsTTLmsOk() (*float64, bool)`

GetUnverifiedCommentsTTLmsOk returns a tuple with the UnverifiedCommentsTTLms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedCommentsTTLms

`func (o *APITenant) SetUnverifiedCommentsTTLms(v float64)`

SetUnverifiedCommentsTTLms sets UnverifiedCommentsTTLms field to given value.

### HasUnverifiedCommentsTTLms

`func (o *APITenant) HasUnverifiedCommentsTTLms() bool`

HasUnverifiedCommentsTTLms returns a boolean if a field has been set.

### SetUnverifiedCommentsTTLmsNil

`func (o *APITenant) SetUnverifiedCommentsTTLmsNil(b bool)`

 SetUnverifiedCommentsTTLmsNil sets the value for UnverifiedCommentsTTLms to be an explicit nil

### UnsetUnverifiedCommentsTTLms
`func (o *APITenant) UnsetUnverifiedCommentsTTLms()`

UnsetUnverifiedCommentsTTLms ensures that no value is present for UnverifiedCommentsTTLms, not even an explicit nil
### GetCommentsRequireApproval

`func (o *APITenant) GetCommentsRequireApproval() bool`

GetCommentsRequireApproval returns the CommentsRequireApproval field if non-nil, zero value otherwise.

### GetCommentsRequireApprovalOk

`func (o *APITenant) GetCommentsRequireApprovalOk() (*bool, bool)`

GetCommentsRequireApprovalOk returns a tuple with the CommentsRequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsRequireApproval

`func (o *APITenant) SetCommentsRequireApproval(v bool)`

SetCommentsRequireApproval sets CommentsRequireApproval field to given value.

### HasCommentsRequireApproval

`func (o *APITenant) HasCommentsRequireApproval() bool`

HasCommentsRequireApproval returns a boolean if a field has been set.

### GetAutoApproveCommentOnVerification

`func (o *APITenant) GetAutoApproveCommentOnVerification() bool`

GetAutoApproveCommentOnVerification returns the AutoApproveCommentOnVerification field if non-nil, zero value otherwise.

### GetAutoApproveCommentOnVerificationOk

`func (o *APITenant) GetAutoApproveCommentOnVerificationOk() (*bool, bool)`

GetAutoApproveCommentOnVerificationOk returns a tuple with the AutoApproveCommentOnVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveCommentOnVerification

`func (o *APITenant) SetAutoApproveCommentOnVerification(v bool)`

SetAutoApproveCommentOnVerification sets AutoApproveCommentOnVerification field to given value.

### HasAutoApproveCommentOnVerification

`func (o *APITenant) HasAutoApproveCommentOnVerification() bool`

HasAutoApproveCommentOnVerification returns a boolean if a field has been set.

### GetSendProfaneToSpam

`func (o *APITenant) GetSendProfaneToSpam() bool`

GetSendProfaneToSpam returns the SendProfaneToSpam field if non-nil, zero value otherwise.

### GetSendProfaneToSpamOk

`func (o *APITenant) GetSendProfaneToSpamOk() (*bool, bool)`

GetSendProfaneToSpamOk returns a tuple with the SendProfaneToSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendProfaneToSpam

`func (o *APITenant) SetSendProfaneToSpam(v bool)`

SetSendProfaneToSpam sets SendProfaneToSpam field to given value.

### HasSendProfaneToSpam

`func (o *APITenant) HasSendProfaneToSpam() bool`

HasSendProfaneToSpam returns a boolean if a field has been set.

### GetHasFlexPricing

`func (o *APITenant) GetHasFlexPricing() bool`

GetHasFlexPricing returns the HasFlexPricing field if non-nil, zero value otherwise.

### GetHasFlexPricingOk

`func (o *APITenant) GetHasFlexPricingOk() (*bool, bool)`

GetHasFlexPricingOk returns a tuple with the HasFlexPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFlexPricing

`func (o *APITenant) SetHasFlexPricing(v bool)`

SetHasFlexPricing sets HasFlexPricing field to given value.

### HasHasFlexPricing

`func (o *APITenant) HasHasFlexPricing() bool`

HasHasFlexPricing returns a boolean if a field has been set.

### GetHasAuditing

`func (o *APITenant) GetHasAuditing() bool`

GetHasAuditing returns the HasAuditing field if non-nil, zero value otherwise.

### GetHasAuditingOk

`func (o *APITenant) GetHasAuditingOk() (*bool, bool)`

GetHasAuditingOk returns a tuple with the HasAuditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAuditing

`func (o *APITenant) SetHasAuditing(v bool)`

SetHasAuditing sets HasAuditing field to given value.

### HasHasAuditing

`func (o *APITenant) HasHasAuditing() bool`

HasHasAuditing returns a boolean if a field has been set.

### GetFlexLastBilledAmount

`func (o *APITenant) GetFlexLastBilledAmount() float64`

GetFlexLastBilledAmount returns the FlexLastBilledAmount field if non-nil, zero value otherwise.

### GetFlexLastBilledAmountOk

`func (o *APITenant) GetFlexLastBilledAmountOk() (*float64, bool)`

GetFlexLastBilledAmountOk returns a tuple with the FlexLastBilledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexLastBilledAmount

`func (o *APITenant) SetFlexLastBilledAmount(v float64)`

SetFlexLastBilledAmount sets FlexLastBilledAmount field to given value.

### HasFlexLastBilledAmount

`func (o *APITenant) HasFlexLastBilledAmount() bool`

HasFlexLastBilledAmount returns a boolean if a field has been set.

### GetDeAnonIpAddr

`func (o *APITenant) GetDeAnonIpAddr() float64`

GetDeAnonIpAddr returns the DeAnonIpAddr field if non-nil, zero value otherwise.

### GetDeAnonIpAddrOk

`func (o *APITenant) GetDeAnonIpAddrOk() (*float64, bool)`

GetDeAnonIpAddrOk returns a tuple with the DeAnonIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeAnonIpAddr

`func (o *APITenant) SetDeAnonIpAddr(v float64)`

SetDeAnonIpAddr sets DeAnonIpAddr field to given value.

### HasDeAnonIpAddr

`func (o *APITenant) HasDeAnonIpAddr() bool`

HasDeAnonIpAddr returns a boolean if a field has been set.

### GetMeta

`func (o *APITenant) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *APITenant) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *APITenant) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *APITenant) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


