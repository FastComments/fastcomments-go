# UpdateTenantBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**SignUpDate** | Pointer to **float64** |  | [optional] 
**PackageId** | Pointer to **string** |  | [optional] 
**PaymentFrequency** | Pointer to **float64** |  | [optional] 
**BillingInfoValid** | Pointer to **bool** |  | [optional] 
**BillingHandledExternally** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**IsSetup** | Pointer to **bool** |  | [optional] 
**DomainConfiguration** | Pointer to [**[]APIDomainConfiguration**](APIDomainConfiguration.md) |  | [optional] 
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
**ManagedByTenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateTenantBody

`func NewUpdateTenantBody() *UpdateTenantBody`

NewUpdateTenantBody instantiates a new UpdateTenantBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantBodyWithDefaults

`func NewUpdateTenantBodyWithDefaults() *UpdateTenantBody`

NewUpdateTenantBodyWithDefaults instantiates a new UpdateTenantBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTenantBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenantBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenantBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTenantBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateTenantBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateTenantBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateTenantBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateTenantBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSignUpDate

`func (o *UpdateTenantBody) GetSignUpDate() float64`

GetSignUpDate returns the SignUpDate field if non-nil, zero value otherwise.

### GetSignUpDateOk

`func (o *UpdateTenantBody) GetSignUpDateOk() (*float64, bool)`

GetSignUpDateOk returns a tuple with the SignUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpDate

`func (o *UpdateTenantBody) SetSignUpDate(v float64)`

SetSignUpDate sets SignUpDate field to given value.

### HasSignUpDate

`func (o *UpdateTenantBody) HasSignUpDate() bool`

HasSignUpDate returns a boolean if a field has been set.

### GetPackageId

`func (o *UpdateTenantBody) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *UpdateTenantBody) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *UpdateTenantBody) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *UpdateTenantBody) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPaymentFrequency

`func (o *UpdateTenantBody) GetPaymentFrequency() float64`

GetPaymentFrequency returns the PaymentFrequency field if non-nil, zero value otherwise.

### GetPaymentFrequencyOk

`func (o *UpdateTenantBody) GetPaymentFrequencyOk() (*float64, bool)`

GetPaymentFrequencyOk returns a tuple with the PaymentFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrequency

`func (o *UpdateTenantBody) SetPaymentFrequency(v float64)`

SetPaymentFrequency sets PaymentFrequency field to given value.

### HasPaymentFrequency

`func (o *UpdateTenantBody) HasPaymentFrequency() bool`

HasPaymentFrequency returns a boolean if a field has been set.

### GetBillingInfoValid

`func (o *UpdateTenantBody) GetBillingInfoValid() bool`

GetBillingInfoValid returns the BillingInfoValid field if non-nil, zero value otherwise.

### GetBillingInfoValidOk

`func (o *UpdateTenantBody) GetBillingInfoValidOk() (*bool, bool)`

GetBillingInfoValidOk returns a tuple with the BillingInfoValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfoValid

`func (o *UpdateTenantBody) SetBillingInfoValid(v bool)`

SetBillingInfoValid sets BillingInfoValid field to given value.

### HasBillingInfoValid

`func (o *UpdateTenantBody) HasBillingInfoValid() bool`

HasBillingInfoValid returns a boolean if a field has been set.

### GetBillingHandledExternally

`func (o *UpdateTenantBody) GetBillingHandledExternally() bool`

GetBillingHandledExternally returns the BillingHandledExternally field if non-nil, zero value otherwise.

### GetBillingHandledExternallyOk

`func (o *UpdateTenantBody) GetBillingHandledExternallyOk() (*bool, bool)`

GetBillingHandledExternallyOk returns a tuple with the BillingHandledExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingHandledExternally

`func (o *UpdateTenantBody) SetBillingHandledExternally(v bool)`

SetBillingHandledExternally sets BillingHandledExternally field to given value.

### HasBillingHandledExternally

`func (o *UpdateTenantBody) HasBillingHandledExternally() bool`

HasBillingHandledExternally returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateTenantBody) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateTenantBody) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateTenantBody) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateTenantBody) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsSetup

`func (o *UpdateTenantBody) GetIsSetup() bool`

GetIsSetup returns the IsSetup field if non-nil, zero value otherwise.

### GetIsSetupOk

`func (o *UpdateTenantBody) GetIsSetupOk() (*bool, bool)`

GetIsSetupOk returns a tuple with the IsSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetup

`func (o *UpdateTenantBody) SetIsSetup(v bool)`

SetIsSetup sets IsSetup field to given value.

### HasIsSetup

`func (o *UpdateTenantBody) HasIsSetup() bool`

HasIsSetup returns a boolean if a field has been set.

### GetDomainConfiguration

`func (o *UpdateTenantBody) GetDomainConfiguration() []APIDomainConfiguration`

GetDomainConfiguration returns the DomainConfiguration field if non-nil, zero value otherwise.

### GetDomainConfigurationOk

`func (o *UpdateTenantBody) GetDomainConfigurationOk() (*[]APIDomainConfiguration, bool)`

GetDomainConfigurationOk returns a tuple with the DomainConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainConfiguration

`func (o *UpdateTenantBody) SetDomainConfiguration(v []APIDomainConfiguration)`

SetDomainConfiguration sets DomainConfiguration field to given value.

### HasDomainConfiguration

`func (o *UpdateTenantBody) HasDomainConfiguration() bool`

HasDomainConfiguration returns a boolean if a field has been set.

### GetBillingInfo

`func (o *UpdateTenantBody) GetBillingInfo() BillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *UpdateTenantBody) GetBillingInfoOk() (*BillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *UpdateTenantBody) SetBillingInfo(v BillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *UpdateTenantBody) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *UpdateTenantBody) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *UpdateTenantBody) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *UpdateTenantBody) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *UpdateTenantBody) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *UpdateTenantBody) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *UpdateTenantBody) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *UpdateTenantBody) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *UpdateTenantBody) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetStripePlanId

`func (o *UpdateTenantBody) GetStripePlanId() string`

GetStripePlanId returns the StripePlanId field if non-nil, zero value otherwise.

### GetStripePlanIdOk

`func (o *UpdateTenantBody) GetStripePlanIdOk() (*string, bool)`

GetStripePlanIdOk returns a tuple with the StripePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePlanId

`func (o *UpdateTenantBody) SetStripePlanId(v string)`

SetStripePlanId sets StripePlanId field to given value.

### HasStripePlanId

`func (o *UpdateTenantBody) HasStripePlanId() bool`

HasStripePlanId returns a boolean if a field has been set.

### GetEnableProfanityFilter

`func (o *UpdateTenantBody) GetEnableProfanityFilter() bool`

GetEnableProfanityFilter returns the EnableProfanityFilter field if non-nil, zero value otherwise.

### GetEnableProfanityFilterOk

`func (o *UpdateTenantBody) GetEnableProfanityFilterOk() (*bool, bool)`

GetEnableProfanityFilterOk returns a tuple with the EnableProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfanityFilter

`func (o *UpdateTenantBody) SetEnableProfanityFilter(v bool)`

SetEnableProfanityFilter sets EnableProfanityFilter field to given value.

### HasEnableProfanityFilter

`func (o *UpdateTenantBody) HasEnableProfanityFilter() bool`

HasEnableProfanityFilter returns a boolean if a field has been set.

### GetEnableSpamFilter

`func (o *UpdateTenantBody) GetEnableSpamFilter() bool`

GetEnableSpamFilter returns the EnableSpamFilter field if non-nil, zero value otherwise.

### GetEnableSpamFilterOk

`func (o *UpdateTenantBody) GetEnableSpamFilterOk() (*bool, bool)`

GetEnableSpamFilterOk returns a tuple with the EnableSpamFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpamFilter

`func (o *UpdateTenantBody) SetEnableSpamFilter(v bool)`

SetEnableSpamFilter sets EnableSpamFilter field to given value.

### HasEnableSpamFilter

`func (o *UpdateTenantBody) HasEnableSpamFilter() bool`

HasEnableSpamFilter returns a boolean if a field has been set.

### GetRemoveUnverifiedComments

`func (o *UpdateTenantBody) GetRemoveUnverifiedComments() bool`

GetRemoveUnverifiedComments returns the RemoveUnverifiedComments field if non-nil, zero value otherwise.

### GetRemoveUnverifiedCommentsOk

`func (o *UpdateTenantBody) GetRemoveUnverifiedCommentsOk() (*bool, bool)`

GetRemoveUnverifiedCommentsOk returns a tuple with the RemoveUnverifiedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUnverifiedComments

`func (o *UpdateTenantBody) SetRemoveUnverifiedComments(v bool)`

SetRemoveUnverifiedComments sets RemoveUnverifiedComments field to given value.

### HasRemoveUnverifiedComments

`func (o *UpdateTenantBody) HasRemoveUnverifiedComments() bool`

HasRemoveUnverifiedComments returns a boolean if a field has been set.

### GetUnverifiedCommentsTTLms

`func (o *UpdateTenantBody) GetUnverifiedCommentsTTLms() float64`

GetUnverifiedCommentsTTLms returns the UnverifiedCommentsTTLms field if non-nil, zero value otherwise.

### GetUnverifiedCommentsTTLmsOk

`func (o *UpdateTenantBody) GetUnverifiedCommentsTTLmsOk() (*float64, bool)`

GetUnverifiedCommentsTTLmsOk returns a tuple with the UnverifiedCommentsTTLms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedCommentsTTLms

`func (o *UpdateTenantBody) SetUnverifiedCommentsTTLms(v float64)`

SetUnverifiedCommentsTTLms sets UnverifiedCommentsTTLms field to given value.

### HasUnverifiedCommentsTTLms

`func (o *UpdateTenantBody) HasUnverifiedCommentsTTLms() bool`

HasUnverifiedCommentsTTLms returns a boolean if a field has been set.

### GetCommentsRequireApproval

`func (o *UpdateTenantBody) GetCommentsRequireApproval() bool`

GetCommentsRequireApproval returns the CommentsRequireApproval field if non-nil, zero value otherwise.

### GetCommentsRequireApprovalOk

`func (o *UpdateTenantBody) GetCommentsRequireApprovalOk() (*bool, bool)`

GetCommentsRequireApprovalOk returns a tuple with the CommentsRequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsRequireApproval

`func (o *UpdateTenantBody) SetCommentsRequireApproval(v bool)`

SetCommentsRequireApproval sets CommentsRequireApproval field to given value.

### HasCommentsRequireApproval

`func (o *UpdateTenantBody) HasCommentsRequireApproval() bool`

HasCommentsRequireApproval returns a boolean if a field has been set.

### GetAutoApproveCommentOnVerification

`func (o *UpdateTenantBody) GetAutoApproveCommentOnVerification() bool`

GetAutoApproveCommentOnVerification returns the AutoApproveCommentOnVerification field if non-nil, zero value otherwise.

### GetAutoApproveCommentOnVerificationOk

`func (o *UpdateTenantBody) GetAutoApproveCommentOnVerificationOk() (*bool, bool)`

GetAutoApproveCommentOnVerificationOk returns a tuple with the AutoApproveCommentOnVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproveCommentOnVerification

`func (o *UpdateTenantBody) SetAutoApproveCommentOnVerification(v bool)`

SetAutoApproveCommentOnVerification sets AutoApproveCommentOnVerification field to given value.

### HasAutoApproveCommentOnVerification

`func (o *UpdateTenantBody) HasAutoApproveCommentOnVerification() bool`

HasAutoApproveCommentOnVerification returns a boolean if a field has been set.

### GetSendProfaneToSpam

`func (o *UpdateTenantBody) GetSendProfaneToSpam() bool`

GetSendProfaneToSpam returns the SendProfaneToSpam field if non-nil, zero value otherwise.

### GetSendProfaneToSpamOk

`func (o *UpdateTenantBody) GetSendProfaneToSpamOk() (*bool, bool)`

GetSendProfaneToSpamOk returns a tuple with the SendProfaneToSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendProfaneToSpam

`func (o *UpdateTenantBody) SetSendProfaneToSpam(v bool)`

SetSendProfaneToSpam sets SendProfaneToSpam field to given value.

### HasSendProfaneToSpam

`func (o *UpdateTenantBody) HasSendProfaneToSpam() bool`

HasSendProfaneToSpam returns a boolean if a field has been set.

### GetDeAnonIpAddr

`func (o *UpdateTenantBody) GetDeAnonIpAddr() float64`

GetDeAnonIpAddr returns the DeAnonIpAddr field if non-nil, zero value otherwise.

### GetDeAnonIpAddrOk

`func (o *UpdateTenantBody) GetDeAnonIpAddrOk() (*float64, bool)`

GetDeAnonIpAddrOk returns a tuple with the DeAnonIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeAnonIpAddr

`func (o *UpdateTenantBody) SetDeAnonIpAddr(v float64)`

SetDeAnonIpAddr sets DeAnonIpAddr field to given value.

### HasDeAnonIpAddr

`func (o *UpdateTenantBody) HasDeAnonIpAddr() bool`

HasDeAnonIpAddr returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateTenantBody) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateTenantBody) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateTenantBody) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateTenantBody) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetManagedByTenantId

`func (o *UpdateTenantBody) GetManagedByTenantId() string`

GetManagedByTenantId returns the ManagedByTenantId field if non-nil, zero value otherwise.

### GetManagedByTenantIdOk

`func (o *UpdateTenantBody) GetManagedByTenantIdOk() (*string, bool)`

GetManagedByTenantIdOk returns a tuple with the ManagedByTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedByTenantId

`func (o *UpdateTenantBody) SetManagedByTenantId(v string)`

SetManagedByTenantId sets ManagedByTenantId field to given value.

### HasManagedByTenantId

`func (o *UpdateTenantBody) HasManagedByTenantId() bool`

HasManagedByTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


