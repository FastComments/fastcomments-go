# CommentLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**D** | **time.Time** |  | 
**T** | [**CommentLogType**](CommentLogType.md) |  | 
**Da** | Pointer to [**CommentLogData**](CommentLogData.md) |  | [optional] 

## Methods

### NewCommentLogEntry

`func NewCommentLogEntry(d time.Time, t CommentLogType, ) *CommentLogEntry`

NewCommentLogEntry instantiates a new CommentLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentLogEntryWithDefaults

`func NewCommentLogEntryWithDefaults() *CommentLogEntry`

NewCommentLogEntryWithDefaults instantiates a new CommentLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetD

`func (o *CommentLogEntry) GetD() time.Time`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *CommentLogEntry) GetDOk() (*time.Time, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *CommentLogEntry) SetD(v time.Time)`

SetD sets D field to given value.


### GetT

`func (o *CommentLogEntry) GetT() CommentLogType`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CommentLogEntry) GetTOk() (*CommentLogType, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CommentLogEntry) SetT(v CommentLogType)`

SetT sets T field to given value.


### GetDa

`func (o *CommentLogEntry) GetDa() CommentLogData`

GetDa returns the Da field if non-nil, zero value otherwise.

### GetDaOk

`func (o *CommentLogEntry) GetDaOk() (*CommentLogData, bool)`

GetDaOk returns a tuple with the Da field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDa

`func (o *CommentLogEntry) SetDa(v CommentLogData)`

SetDa sets Da field to given value.

### HasDa

`func (o *CommentLogEntry) HasDa() bool`

HasDa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


