# GetPageByURLIdAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Page** | Pointer to [**APIPage**](APIPage.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewGetPageByURLIdAPIResponse

`func NewGetPageByURLIdAPIResponse(status string, ) *GetPageByURLIdAPIResponse`

NewGetPageByURLIdAPIResponse instantiates a new GetPageByURLIdAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPageByURLIdAPIResponseWithDefaults

`func NewGetPageByURLIdAPIResponseWithDefaults() *GetPageByURLIdAPIResponse`

NewGetPageByURLIdAPIResponseWithDefaults instantiates a new GetPageByURLIdAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *GetPageByURLIdAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetPageByURLIdAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetPageByURLIdAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetPageByURLIdAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *GetPageByURLIdAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPageByURLIdAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPageByURLIdAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPageByURLIdAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPage

`func (o *GetPageByURLIdAPIResponse) GetPage() APIPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPageByURLIdAPIResponse) GetPageOk() (*APIPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPageByURLIdAPIResponse) SetPage(v APIPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPageByURLIdAPIResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStatus

`func (o *GetPageByURLIdAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPageByURLIdAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPageByURLIdAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


