# GetCommentVoteUserNamesSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 
**VoteUserNames** | **[]string** |  | 
**HasMore** | **bool** |  | 

## Methods

### NewGetCommentVoteUserNamesSuccessResponse

`func NewGetCommentVoteUserNamesSuccessResponse(status ImportedAPIStatusSUCCESS, voteUserNames []string, hasMore bool, ) *GetCommentVoteUserNamesSuccessResponse`

NewGetCommentVoteUserNamesSuccessResponse instantiates a new GetCommentVoteUserNamesSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentVoteUserNamesSuccessResponseWithDefaults

`func NewGetCommentVoteUserNamesSuccessResponseWithDefaults() *GetCommentVoteUserNamesSuccessResponse`

NewGetCommentVoteUserNamesSuccessResponseWithDefaults instantiates a new GetCommentVoteUserNamesSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCommentVoteUserNamesSuccessResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentVoteUserNamesSuccessResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentVoteUserNamesSuccessResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.


### GetVoteUserNames

`func (o *GetCommentVoteUserNamesSuccessResponse) GetVoteUserNames() []string`

GetVoteUserNames returns the VoteUserNames field if non-nil, zero value otherwise.

### GetVoteUserNamesOk

`func (o *GetCommentVoteUserNamesSuccessResponse) GetVoteUserNamesOk() (*[]string, bool)`

GetVoteUserNamesOk returns a tuple with the VoteUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteUserNames

`func (o *GetCommentVoteUserNamesSuccessResponse) SetVoteUserNames(v []string)`

SetVoteUserNames sets VoteUserNames field to given value.


### GetHasMore

`func (o *GetCommentVoteUserNamesSuccessResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCommentVoteUserNamesSuccessResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCommentVoteUserNamesSuccessResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


