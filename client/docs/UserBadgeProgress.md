# UserBadgeProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**UserId** | **string** |  | 
**FirstCommentId** | **string** |  | 
**FirstCommentDate** | **time.Time** |  | 
**AutoTrustFactor** | Pointer to **float64** |  | [optional] 
**ManualTrustFactor** | Pointer to **float64** |  | [optional] 
**Progress** | **map[string]float64** | Construct a type with a set of properties K of type T | 

## Methods

### NewUserBadgeProgress

`func NewUserBadgeProgress(id string, tenantId string, userId string, firstCommentId string, firstCommentDate time.Time, progress map[string]float64, ) *UserBadgeProgress`

NewUserBadgeProgress instantiates a new UserBadgeProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBadgeProgressWithDefaults

`func NewUserBadgeProgressWithDefaults() *UserBadgeProgress`

NewUserBadgeProgressWithDefaults instantiates a new UserBadgeProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBadgeProgress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBadgeProgress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBadgeProgress) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *UserBadgeProgress) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserBadgeProgress) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserBadgeProgress) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUserId

`func (o *UserBadgeProgress) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserBadgeProgress) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserBadgeProgress) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFirstCommentId

`func (o *UserBadgeProgress) GetFirstCommentId() string`

GetFirstCommentId returns the FirstCommentId field if non-nil, zero value otherwise.

### GetFirstCommentIdOk

`func (o *UserBadgeProgress) GetFirstCommentIdOk() (*string, bool)`

GetFirstCommentIdOk returns a tuple with the FirstCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommentId

`func (o *UserBadgeProgress) SetFirstCommentId(v string)`

SetFirstCommentId sets FirstCommentId field to given value.


### GetFirstCommentDate

`func (o *UserBadgeProgress) GetFirstCommentDate() time.Time`

GetFirstCommentDate returns the FirstCommentDate field if non-nil, zero value otherwise.

### GetFirstCommentDateOk

`func (o *UserBadgeProgress) GetFirstCommentDateOk() (*time.Time, bool)`

GetFirstCommentDateOk returns a tuple with the FirstCommentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommentDate

`func (o *UserBadgeProgress) SetFirstCommentDate(v time.Time)`

SetFirstCommentDate sets FirstCommentDate field to given value.


### GetAutoTrustFactor

`func (o *UserBadgeProgress) GetAutoTrustFactor() float64`

GetAutoTrustFactor returns the AutoTrustFactor field if non-nil, zero value otherwise.

### GetAutoTrustFactorOk

`func (o *UserBadgeProgress) GetAutoTrustFactorOk() (*float64, bool)`

GetAutoTrustFactorOk returns a tuple with the AutoTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTrustFactor

`func (o *UserBadgeProgress) SetAutoTrustFactor(v float64)`

SetAutoTrustFactor sets AutoTrustFactor field to given value.

### HasAutoTrustFactor

`func (o *UserBadgeProgress) HasAutoTrustFactor() bool`

HasAutoTrustFactor returns a boolean if a field has been set.

### GetManualTrustFactor

`func (o *UserBadgeProgress) GetManualTrustFactor() float64`

GetManualTrustFactor returns the ManualTrustFactor field if non-nil, zero value otherwise.

### GetManualTrustFactorOk

`func (o *UserBadgeProgress) GetManualTrustFactorOk() (*float64, bool)`

GetManualTrustFactorOk returns a tuple with the ManualTrustFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualTrustFactor

`func (o *UserBadgeProgress) SetManualTrustFactor(v float64)`

SetManualTrustFactor sets ManualTrustFactor field to given value.

### HasManualTrustFactor

`func (o *UserBadgeProgress) HasManualTrustFactor() bool`

HasManualTrustFactor returns a boolean if a field has been set.

### GetProgress

`func (o *UserBadgeProgress) GetProgress() map[string]float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *UserBadgeProgress) GetProgressOk() (*map[string]float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *UserBadgeProgress) SetProgress(v map[string]float64)`

SetProgress sets Progress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


