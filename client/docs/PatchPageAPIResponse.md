# PatchPageAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Page** | Pointer to [**APIPage**](APIPage.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewPatchPageAPIResponse

`func NewPatchPageAPIResponse(status string, ) *PatchPageAPIResponse`

NewPatchPageAPIResponse instantiates a new PatchPageAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPageAPIResponseWithDefaults

`func NewPatchPageAPIResponseWithDefaults() *PatchPageAPIResponse`

NewPatchPageAPIResponseWithDefaults instantiates a new PatchPageAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *PatchPageAPIResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PatchPageAPIResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PatchPageAPIResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PatchPageAPIResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *PatchPageAPIResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PatchPageAPIResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PatchPageAPIResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PatchPageAPIResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPage

`func (o *PatchPageAPIResponse) GetPage() APIPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PatchPageAPIResponse) GetPageOk() (*APIPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PatchPageAPIResponse) SetPage(v APIPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PatchPageAPIResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetStatus

`func (o *PatchPageAPIResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchPageAPIResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchPageAPIResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


