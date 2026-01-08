# BulkCreateHashTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Results** | [**[]AddHashTag200Response**](AddHashTag200Response.md) |  | 

## Methods

### NewBulkCreateHashTagsResponse

`func NewBulkCreateHashTagsResponse(status APIStatus, results []AddHashTag200Response, ) *BulkCreateHashTagsResponse`

NewBulkCreateHashTagsResponse instantiates a new BulkCreateHashTagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateHashTagsResponseWithDefaults

`func NewBulkCreateHashTagsResponseWithDefaults() *BulkCreateHashTagsResponse`

NewBulkCreateHashTagsResponseWithDefaults instantiates a new BulkCreateHashTagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkCreateHashTagsResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkCreateHashTagsResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkCreateHashTagsResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetResults

`func (o *BulkCreateHashTagsResponse) GetResults() []AddHashTag200Response`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkCreateHashTagsResponse) GetResultsOk() (*[]AddHashTag200Response, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkCreateHashTagsResponse) SetResults(v []AddHashTag200Response)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


