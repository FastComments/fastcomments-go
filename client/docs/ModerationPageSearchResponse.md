# ModerationPageSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pages** | [**[]ModerationPageSearchProjected**](ModerationPageSearchProjected.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewModerationPageSearchResponse

`func NewModerationPageSearchResponse(pages []ModerationPageSearchProjected, status APIStatus, ) *ModerationPageSearchResponse`

NewModerationPageSearchResponse instantiates a new ModerationPageSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationPageSearchResponseWithDefaults

`func NewModerationPageSearchResponseWithDefaults() *ModerationPageSearchResponse`

NewModerationPageSearchResponseWithDefaults instantiates a new ModerationPageSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPages

`func (o *ModerationPageSearchResponse) GetPages() []ModerationPageSearchProjected`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ModerationPageSearchResponse) GetPagesOk() (*[]ModerationPageSearchProjected, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ModerationPageSearchResponse) SetPages(v []ModerationPageSearchProjected)`

SetPages sets Pages field to given value.


### GetStatus

`func (o *ModerationPageSearchResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationPageSearchResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationPageSearchResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


