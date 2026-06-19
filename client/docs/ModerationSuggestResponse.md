# ModerationSuggestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Pages** | Pointer to [**[]ModerationPageSearchProjected**](ModerationPageSearchProjected.md) |  | [optional] 
**Users** | Pointer to [**[]ModerationUserSearchProjected**](ModerationUserSearchProjected.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewModerationSuggestResponse

`func NewModerationSuggestResponse(status string, ) *ModerationSuggestResponse`

NewModerationSuggestResponse instantiates a new ModerationSuggestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationSuggestResponseWithDefaults

`func NewModerationSuggestResponseWithDefaults() *ModerationSuggestResponse`

NewModerationSuggestResponseWithDefaults instantiates a new ModerationSuggestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModerationSuggestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationSuggestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationSuggestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPages

`func (o *ModerationSuggestResponse) GetPages() []ModerationPageSearchProjected`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ModerationSuggestResponse) GetPagesOk() (*[]ModerationPageSearchProjected, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ModerationSuggestResponse) SetPages(v []ModerationPageSearchProjected)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ModerationSuggestResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetUsers

`func (o *ModerationSuggestResponse) GetUsers() []ModerationUserSearchProjected`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModerationSuggestResponse) GetUsersOk() (*[]ModerationUserSearchProjected, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModerationSuggestResponse) SetUsers(v []ModerationUserSearchProjected)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ModerationSuggestResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCode

`func (o *ModerationSuggestResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModerationSuggestResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModerationSuggestResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ModerationSuggestResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


