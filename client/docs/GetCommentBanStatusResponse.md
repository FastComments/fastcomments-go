# GetCommentBanStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**EmailDomain** | **NullableString** |  | 
**CanIPBan** | **NullableBool** |  | 

## Methods

### NewGetCommentBanStatusResponse

`func NewGetCommentBanStatusResponse(status string, emailDomain NullableString, canIPBan NullableBool, ) *GetCommentBanStatusResponse`

NewGetCommentBanStatusResponse instantiates a new GetCommentBanStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentBanStatusResponseWithDefaults

`func NewGetCommentBanStatusResponseWithDefaults() *GetCommentBanStatusResponse`

NewGetCommentBanStatusResponseWithDefaults instantiates a new GetCommentBanStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCommentBanStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentBanStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentBanStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEmailDomain

`func (o *GetCommentBanStatusResponse) GetEmailDomain() string`

GetEmailDomain returns the EmailDomain field if non-nil, zero value otherwise.

### GetEmailDomainOk

`func (o *GetCommentBanStatusResponse) GetEmailDomainOk() (*string, bool)`

GetEmailDomainOk returns a tuple with the EmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomain

`func (o *GetCommentBanStatusResponse) SetEmailDomain(v string)`

SetEmailDomain sets EmailDomain field to given value.


### SetEmailDomainNil

`func (o *GetCommentBanStatusResponse) SetEmailDomainNil(b bool)`

 SetEmailDomainNil sets the value for EmailDomain to be an explicit nil

### UnsetEmailDomain
`func (o *GetCommentBanStatusResponse) UnsetEmailDomain()`

UnsetEmailDomain ensures that no value is present for EmailDomain, not even an explicit nil
### GetCanIPBan

`func (o *GetCommentBanStatusResponse) GetCanIPBan() bool`

GetCanIPBan returns the CanIPBan field if non-nil, zero value otherwise.

### GetCanIPBanOk

`func (o *GetCommentBanStatusResponse) GetCanIPBanOk() (*bool, bool)`

GetCanIPBanOk returns a tuple with the CanIPBan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanIPBan

`func (o *GetCommentBanStatusResponse) SetCanIPBan(v bool)`

SetCanIPBan sets CanIPBan field to given value.


### SetCanIPBanNil

`func (o *GetCommentBanStatusResponse) SetCanIPBanNil(b bool)`

 SetCanIPBanNil sets the value for CanIPBan to be an explicit nil

### UnsetCanIPBan
`func (o *GetCommentBanStatusResponse) UnsetCanIPBan()`

UnsetCanIPBan ensures that no value is present for CanIPBan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


