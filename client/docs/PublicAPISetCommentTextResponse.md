# PublicAPISetCommentTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | [**PickFCommentApprovedOrCommentHTML**](PickFCommentApprovedOrCommentHTML.md) |  | 
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 

## Methods

### NewPublicAPISetCommentTextResponse

`func NewPublicAPISetCommentTextResponse(comment PickFCommentApprovedOrCommentHTML, status ImportedAPIStatusSUCCESS, ) *PublicAPISetCommentTextResponse`

NewPublicAPISetCommentTextResponse instantiates a new PublicAPISetCommentTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAPISetCommentTextResponseWithDefaults

`func NewPublicAPISetCommentTextResponseWithDefaults() *PublicAPISetCommentTextResponse`

NewPublicAPISetCommentTextResponseWithDefaults instantiates a new PublicAPISetCommentTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PublicAPISetCommentTextResponse) GetComment() PickFCommentApprovedOrCommentHTML`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PublicAPISetCommentTextResponse) GetCommentOk() (*PickFCommentApprovedOrCommentHTML, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PublicAPISetCommentTextResponse) SetComment(v PickFCommentApprovedOrCommentHTML)`

SetComment sets Comment field to given value.


### GetStatus

`func (o *PublicAPISetCommentTextResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicAPISetCommentTextResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicAPISetCommentTextResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


