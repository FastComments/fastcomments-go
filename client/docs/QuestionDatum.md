# QuestionDatum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V** | **map[string]float64** | Construct a type with a set of properties K of type T | 
**Total** | **int64** |  | 

## Methods

### NewQuestionDatum

`func NewQuestionDatum(v map[string]float64, total int64, ) *QuestionDatum`

NewQuestionDatum instantiates a new QuestionDatum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionDatumWithDefaults

`func NewQuestionDatumWithDefaults() *QuestionDatum`

NewQuestionDatumWithDefaults instantiates a new QuestionDatum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV

`func (o *QuestionDatum) GetV() map[string]float64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *QuestionDatum) GetVOk() (*map[string]float64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *QuestionDatum) SetV(v map[string]float64)`

SetV sets V field to given value.


### GetTotal

`func (o *QuestionDatum) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuestionDatum) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuestionDatum) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


