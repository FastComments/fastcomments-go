# GetPublicPagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | **NullableString** |  | 
**Pages** | [**[]PublicPage**](PublicPage.md) |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 

## Methods

### NewGetPublicPagesResponse

`func NewGetPublicPagesResponse(nextCursor NullableString, pages []PublicPage, status APIStatus, ) *GetPublicPagesResponse`

NewGetPublicPagesResponse instantiates a new GetPublicPagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicPagesResponseWithDefaults

`func NewGetPublicPagesResponseWithDefaults() *GetPublicPagesResponse`

NewGetPublicPagesResponseWithDefaults instantiates a new GetPublicPagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *GetPublicPagesResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetPublicPagesResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetPublicPagesResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetPublicPagesResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetPublicPagesResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPages

`func (o *GetPublicPagesResponse) GetPages() []PublicPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetPublicPagesResponse) GetPagesOk() (*[]PublicPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetPublicPagesResponse) SetPages(v []PublicPage)`

SetPages sets Pages field to given value.


### GetStatus

`func (o *GetPublicPagesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPublicPagesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPublicPagesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


