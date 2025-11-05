# PublicAPIDeleteCommentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**PickFCommentIsDeletedOrCommentHTMLOrCommenterNameOrUserId**](PickFCommentIsDeletedOrCommentHTMLOrCommenterNameOrUserId.md) |  | [optional] 
**HardRemoved** | **bool** |  | 
**Status** | [**ImportedAPIStatusSUCCESS**](ImportedAPIStatusSUCCESS.md) |  | 

## Methods

### NewPublicAPIDeleteCommentResponse

`func NewPublicAPIDeleteCommentResponse(hardRemoved bool, status ImportedAPIStatusSUCCESS, ) *PublicAPIDeleteCommentResponse`

NewPublicAPIDeleteCommentResponse instantiates a new PublicAPIDeleteCommentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAPIDeleteCommentResponseWithDefaults

`func NewPublicAPIDeleteCommentResponseWithDefaults() *PublicAPIDeleteCommentResponse`

NewPublicAPIDeleteCommentResponseWithDefaults instantiates a new PublicAPIDeleteCommentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PublicAPIDeleteCommentResponse) GetComment() PickFCommentIsDeletedOrCommentHTMLOrCommenterNameOrUserId`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PublicAPIDeleteCommentResponse) GetCommentOk() (*PickFCommentIsDeletedOrCommentHTMLOrCommenterNameOrUserId, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PublicAPIDeleteCommentResponse) SetComment(v PickFCommentIsDeletedOrCommentHTMLOrCommenterNameOrUserId)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PublicAPIDeleteCommentResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetHardRemoved

`func (o *PublicAPIDeleteCommentResponse) GetHardRemoved() bool`

GetHardRemoved returns the HardRemoved field if non-nil, zero value otherwise.

### GetHardRemovedOk

`func (o *PublicAPIDeleteCommentResponse) GetHardRemovedOk() (*bool, bool)`

GetHardRemovedOk returns a tuple with the HardRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardRemoved

`func (o *PublicAPIDeleteCommentResponse) SetHardRemoved(v bool)`

SetHardRemoved sets HardRemoved field to given value.


### GetStatus

`func (o *PublicAPIDeleteCommentResponse) GetStatus() ImportedAPIStatusSUCCESS`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicAPIDeleteCommentResponse) GetStatusOk() (*ImportedAPIStatusSUCCESS, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicAPIDeleteCommentResponse) SetStatus(v ImportedAPIStatusSUCCESS)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


