# GetPagesAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Pages** | Pointer to [**[]APIPage**](APIPage.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewGetPagesAPIResponse

`func NewGetPagesAPIResponse(status string, ) *GetPagesAPIResponse`

NewGetPagesAPIResponse instantiates a new GetPagesAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPagesAPIResponseWithDefaults

`func NewGetPagesAPIResponseWithDefaults() *GetPagesAPIResponse`

NewGetPagesAPIResponseWithDefaults instantiates a new GetPagesAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *GetPagesAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetPagesAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetPagesAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetPagesAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *GetPagesAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPagesAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPagesAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPagesAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPages

`func (o *GetPagesAPIResponse) GetPages() []APIPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetPagesAPIResponse) GetPagesOk() (*[]APIPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetPagesAPIResponse) SetPages(v []APIPage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetPagesAPIResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetStatus

`func (o *GetPagesAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPagesAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPagesAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


