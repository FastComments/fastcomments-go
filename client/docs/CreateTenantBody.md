# CreateTenantBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DomainConfiguration** | [**[]APIDomainConfiguration**](APIDomainConfiguration.md) |  | 
**Email** | Pointer to **string** |  | [optional] 
**SignUpDate** | Pointer to **float64** |  | [optional] 
**PackageId** | Pointer to **string** |  | [optional] 
**PaymentFrequency** | Pointer to **float64** |  | [optional] 
**BillingInfoValid** | Pointer to **bool** |  | [optional] 
**BillingHandledExternally** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**IsSetup** | Pointer to **bool** |  | [optional] 
**BillingInfo** | Pointer to [**BillingInfo**](BillingInfo.md) |  | [optional] 
**StripeCustomerId** | Pointer to **string** |  | [optional] 
**StripeSubscriptionId** | Pointer to **string** |  | [optional] 
**StripePlanId** | Pointer to **string** |  | [optional] 
**EnableProfanityFilter** | Pointer to **bool** |  | [optional] 
**EnableSpamFilter** | Pointer to **bool** |  | [optional] 
**RemoveUnverifiedComments** | Pointer to **bool** |  | [optional] 
**UnverifiedCommentsTTLms** | Pointer to **float64** |  | [optional] 
**CommentsRequireApproval** | Pointer to **bool** |  | [optional] 
**AutoApproveCommentOnVerification** | Pointer to **bool** |  | [optional] 
**SendProfaneToSpam** | Pointer to **bool** |  | [optional] 
**DeAnonIpAddr** | Pointer to **float64** |  | [optional] 
**Meta** | Pointer to **map[string]string** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewCreateTenantBody

`func NewCreateTenantBody(name string, domainConfiguration []APIDomainConfiguration, ) *CreateTenantBody`

NewCreateTenantBody instantiates a new CreateTenantBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantBodyWithDefaults

`func NewCreateTenantBodyWithDefaults() *CreateTenantBody`

NewCreateTenantBodyWithDefaults instantiates a new CreateTenantBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTenantBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantBody) SetName(v string)`

SetName sets Name field to given value.


### GetDomainConfiguration

`func (o *CreateTenantBody) GetDomainConfiguration() []APIDomainConfiguration`

GetDomainConfiguration returns the DomainConfiguration field if non-nil, zero value otherwise.

### GetDomainConfigurationOk

`func (o *CreateTenantBody) GetDomainConfigurationOk() (*[]APIDomainConfiguration, bool)`

GetDomainConfigurationOk returns a tuple with the DomainConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainConfiguration

`func (o *CreateTenantBody) SetDomainConfiguration(v []APIDomainConfiguration)`

SetDomainConfiguration sets DomainConfiguration field to given value.


### GetEmail

`func (o *CreateTenantBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateTenantBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateTenantBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateTenantBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSignUpDate

`func (o *CreateTenantBody) GetSignUpDate() float64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *CreateTenantBody) GetSignUpDateOk() (*float64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *CreateTenantBody) SetSignUpDate(v float64)`

SetSignUpDate sets SignUpDate field to given value.

### HasSignUpDate

`func (o *CreateTenantBody) HasSignUpDate() bool`

HasSignUpDate returns a boolean if a field has been set.

### GetPackageId

`func (o *CreateTenantBody) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CreateTenantBody) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CreateTenantBody) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CreateTenantBody) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPaymentFrequency

`func (o *CreateTenantBody) GetPaymentFrequency() float64`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *CreateTenantBody) GetPaymentFrequencyOk() (*float64, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *CreateTenantBody) SetPaymentFrequency(v float64)`

SetPaymentFrequency sets PaymentFrequency field to given value.

### HasPaymentFrequency

`func (o *CreateTenantBody) HasPaymentFrequency() bool`

HasPaymentFrequency returns a boolean if a field has been set.

### GetBillingInfoValid

`func (o *CreateTenantBody) GetBillingInfoValid() bool`

GetBillingInfoValid returns the BillingInfoValid field if non-nil, zero value otherwise.

### GetBillingInfoValidOk

`func (o *CreateTenantBody) GetBillingInfoValidOk() (*bool, bool)`

GetBillingInfoValidOk returns a tuple with the BillingInfoValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfoValid

`func (o *CreateTenantBody) SetBillingInfoValid(v bool)`

SetBillingInfoValid sets BillingInfoValid field to given value.

### HasBillingInfoValid

`func (o *CreateTenantBody) HasBillingInfoValid() bool`

HasBillingInfoValid returns a boolean if a field has been set.

### GetBillingHandledExternally

`func (o *CreateTenantBody) GetBillingHandledExternally() bool`

GetBillingHandledExternally returns the BillingHandledExternally field if non-nil, zero value otherwise.

### GetBillingHandledExternallyOk

`func (o *CreateTenantBody) GetBillingHandledExternallyOk() (*bool, bool)`

GetBillingHandledExternallyOk returns a tuple with the BillingHandledExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHandledExternally

`func (o *CreateTenantBody) SetBillingHandledExternally(v bool)`

SetBillingHandledExternally sets BillingHandledExternally field to given value.

### HasBillingHandledExternally

`func (o *CreateTenantBody) HasBillingHandledExternally() bool`

HasBillingHandledExternally returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreateTenantBody) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateTenantBody) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateTenantBody) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateTenantBody) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsSetup

`func (o *CreateTenantBody) GetIsSetup() bool`

GetIsSetup returns the IsSetup field if non-nil, zero value otherwise.

### GetIsSetupOk

`func (o *CreateTenantBody) GetIsSetupOk() (*bool, bool)`

GetIsSetupOk returns a tuple with the IsSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetup

`func (o *CreateTenantBody) SetIsSetup(v bool)`

SetIsSetup sets IsSetup field to given value.

### HasIsSetup

`func (o *CreateTenantBody) HasIsSetup() bool`

HasIsSetup returns a boolean if a field has been set.

### GetBillingInfo

`func (o *CreateTenantBody) GetBillingInfo() BillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *CreateTenantBody) GetBillingInfoOk() (*BillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *CreateTenantBody) SetBillingInfo(v BillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *CreateTenantBody) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *CreateTenantBody) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CreateTenantBody) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CreateTenantBody) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *CreateTenantBody) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *CreateTenantBody) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *CreateTenantBody) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *CreateTenantBody) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *CreateTenantBody) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetStripePlanId

`func (o *CreateTenantBody) GetStripePlanId() string`

GetStripePlanId returns the StripePlanId field if non-nil, zero value otherwise.

### GetStripePlanIdOk

`func (o *CreateTenantBody) GetStripePlanIdOk() (*string, bool)`

GetStripePlanIdOk returns a tuple with the StripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePlanId

`func (o *CreateTenantBody) SetStripePlanId(v string)`

SetStripePlanId sets StripePlanId field to given value.

### HasStripePlanId

`func (o *CreateTenantBody) HasStripePlanId() bool`

HasStripePlanId returns a boolean if a field has been set.

### GetEnableProfanityFilter

`func (o *CreateTenantBody) GetEnableProfanityFilter() bool`

GetEnableProfanityFilter returns the EnableProfanityFilter field if non-nil, zero value otherwise.

### GetEnableProfanityFilterOk

`func (o *CreateTenantBody) GetEnableProfanityFilterOk() (*bool, bool)`

GetEnableProfanityFilterOk returns a tuple with the EnableProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfanityFilter

`func (o *CreateTenantBody) SetEnableProfanityFilter(v bool)`

SetEnableProfanityFilter sets EnableProfanityFilter field to given value.

### HasEnableProfanityFilter

`func (o *CreateTenantBody) HasEnableProfanityFilter() bool`

HasEnableProfanityFilter returns a boolean if a field has been set.

### GetEnableSpamFilter

`func (o *CreateTenantBody) GetEnableSpamFilter() bool`

GetEnableSpamFilter returns the EnableSpamFilter field if non-nil, zero value otherwise.

### GetEnableSpamFilterOk

`func (o *CreateTenantBody) GetEnableSpamFilterOk() (*bool, bool)`

GetEnableSpamFilterOk returns a tuple with the EnableSpamFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpamFilter

`func (o *CreateTenantBody) SetEnableSpamFilter(v bool)`

SetEnableSpamFilter sets EnableSpamFilter field to given value.

### HasEnableSpamFilter

`func (o *CreateTenantBody) HasEnableSpamFilter() bool`

HasEnableSpamFilter returns a boolean if a field has been set.

### GetRemoveUnverifiedComments

`func (o *CreateTenantBody) GetRemoveUnverifiedComments() bool`

GetRemoveUnverifiedComments returns the RemoveUnverifiedComments field if non-nil, zero value otherwise.

### GetRemoveUnverifiedCommentsOk

`func (o *CreateTenantBody) GetRemoveUnverifiedCommentsOk() (*bool, bool)`

GetRemoveUnverifiedCommentsOk returns a tuple with the RemoveUnverifiedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUnverifiedComments

`func (o *CreateTenantBody) SetRemoveUnverifiedComments(v bool)`

SetRemoveUnverifiedComments sets RemoveUnverifiedComments field to given value.

### HasRemoveUnverifiedComments

`func (o *CreateTenantBody) HasRemoveUnverifiedComments() bool`

HasRemoveUnverifiedComments returns a boolean if a field has been set.

### GetUnverifiedCommentsTTLms

`func (o *CreateTenantBody) GetUnverifiedCommentsTTLms() float64`

GetUnverifiedCommentsTTLms returns the UnverifiedCommentsTTLms field if non-nil, zero value otherwise.

### GetUnverifiedCommentsTTLmsOk

`func (o *CreateTenantBody) GetUnverifiedCommentsTTLmsOk() (*float64, bool)`

GetUnverifiedCommentsTTLmsOk returns a tuple with the UnverifiedCommentsTTLms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedCommentsTTLms

`func (o *CreateTenantBody) SetUnverifiedCommentsTTLms(v float64)`

SetUnverifiedCommentsTTLms sets UnverifiedCommentsTTLms field to given value.

### HasUnverifiedCommentsTTLms

`func (o *CreateTenantBody) HasUnverifiedCommentsTTLms() bool`

HasUnverifiedCommentsTTLms returns a boolean if a field has been set.

### GetCommentsRequireApproval

`func (o *CreateTenantBody) GetCommentsRequireApproval() bool`

GetCommentsRequireApproval returns the CommentsRequireApproval field if non-nil, zero value otherwise.

### GetCommentsRequireApprovalOk

`func (o *CreateTenantBody) GetCommentsRequireApprovalOk() (*bool, bool)`

GetCommentsRequireApprovalOk returns a tuple with the CommentsRequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsRequireApproval

`func (o *CreateTenantBody) SetCommentsRequireApproval(v bool)`

SetCommentsRequireApproval sets CommentsRequireApproval field to given value.

### HasCommentsRequireApproval

`func (o *CreateTenantBody) HasCommentsRequireApproval() bool`

HasCommentsRequireApproval returns a boolean if a field has been set.

### GetAutoApproveCommentOnVerification

`func (o *CreateTenantBody) GetAutoApproveCommentOnVerification() bool`

GetAutoApproveCommentOnVerification returns the AutoApproveCommentOnVerification field if non-nil, zero value otherwise.

### GetAutoApproveCommentOnVerificationOk

`func (o *CreateTenantBody) GetAutoApproveCommentOnVerificationOk() (*bool, bool)`

GetAutoApproveCommentOnVerificationOk returns a tuple with the AutoApproveCommentOnVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveCommentOnVerification

`func (o *CreateTenantBody) SetAutoApproveCommentOnVerification(v bool)`

SetAutoApproveCommentOnVerification sets AutoApproveCommentOnVerification field to given value.

### HasAutoApproveCommentOnVerification

`func (o *CreateTenantBody) HasAutoApproveCommentOnVerification() bool`

HasAutoApproveCommentOnVerification returns a boolean if a field has been set.

### GetSendProfaneToSpam

`func (o *CreateTenantBody) GetSendProfaneToSpam() bool`

GetSendProfaneToSpam returns the SendProfaneToSpam field if non-nil, zero value otherwise.

### GetSendProfaneToSpamOk

`func (o *CreateTenantBody) GetSendProfaneToSpamOk() (*bool, bool)`

GetSendProfaneToSpamOk returns a tuple with the SendProfaneToSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendProfaneToSpam

`func (o *CreateTenantBody) SetSendProfaneToSpam(v bool)`

SetSendProfaneToSpam sets SendProfaneToSpam field to given value.

### HasSendProfaneToSpam

`func (o *CreateTenantBody) HasSendProfaneToSpam() bool`

HasSendProfaneToSpam returns a boolean if a field has been set.

### GetDeAnonIpAddr

`func (o *CreateTenantBody) GetDeAnonIpAddr() float64`

GetDeAnonIpAddr returns the DeAnonIpAddr field if non-nil, zero value otherwise.

### GetDeAnonIpAddrOk

`func (o *CreateTenantBody) GetDeAnonIpAddrOk() (*float64, bool)`

GetDeAnonIpAddrOk returns a tuple with the DeAnonIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeAnonIpAddr

`func (o *CreateTenantBody) SetDeAnonIpAddr(v float64)`

SetDeAnonIpAddr sets DeAnonIpAddr field to given value.

### HasDeAnonIpAddr

`func (o *CreateTenantBody) HasDeAnonIpAddr() bool`

HasDeAnonIpAddr returns a boolean if a field has been set.

### GetMeta

`func (o *CreateTenantBody) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateTenantBody) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateTenantBody) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateTenantBody) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


