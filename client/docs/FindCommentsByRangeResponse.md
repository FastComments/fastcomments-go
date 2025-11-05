# FindCommentsByRangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]FindCommentsByRangeItem**](FindCommentsByRangeItem.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewFindCommentsByRangeResponse

`func NewFindCommentsByRangeResponse(results []FindCommentsByRangeItem, createdAt time.Time, ) *FindCommentsByRangeResponse`

NewFindCommentsByRangeResponse instantiates a new FindCommentsByRangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindCommentsByRangeResponseWithDefaults

`func NewFindCommentsByRangeResponseWithDefaults() *FindCommentsByRangeResponse`

NewFindCommentsByRangeResponseWithDefaults instantiates a new FindCommentsByRangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *FindCommentsByRangeResponse) GetResults() []FindCommentsByRangeItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FindCommentsByRangeResponse) GetResultsOk() (*[]FindCommentsByRangeItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FindCommentsByRangeResponse) SetResults(v []FindCommentsByRangeItem)`

SetResults sets Results field to given value.


### GetCreatedAt

`func (o *FindCommentsByRangeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FindCommentsByRangeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FindCommentsByRangeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


