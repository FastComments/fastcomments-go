# GetV1PageLikesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlIdWS** | **string** |  | 
**DidLike** | **bool** |  | 
**CommentCount** | **int32** |  | 
**LikeCount** | **int32** |  | 
**Status** | [**APIStatus**](APIStatus.md) |  | 
**Reason** | **string** |  | 
**Code** | **string** |  | 
**SecondaryCode** | Pointer to **string** |  | [optional] 
**BannedUntil** | Pointer to **int64** |  | [optional] 
**MaxCharacterLength** | Pointer to **int32** |  | [optional] 
**TranslatedError** | Pointer to **string** |  | [optional] 
**CustomConfig** | Pointer to [**CustomConfigParameters**](CustomConfigParameters.md) |  | [optional] 

## Methods

### NewGetV1PageLikesResponse

`func NewGetV1PageLikesResponse(urlIdWS string, didLike bool, commentCount int32, likeCount int32, status APIStatus, reason string, code string, ) *GetV1PageLikesResponse`

NewGetV1PageLikesResponse instantiates a new GetV1PageLikesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetV1PageLikesResponseWithDefaults

`func NewGetV1PageLikesResponseWithDefaults() *GetV1PageLikesResponse`

NewGetV1PageLikesResponseWithDefaults instantiates a new GetV1PageLikesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlIdWS

`func (o *GetV1PageLikesResponse) GetUrlIdWS() string`

GetUrlIdWS returns the UrlIdWS field if non-nil, zero value otherwise.

### GetUrlIdWSOk

`func (o *GetV1PageLikesResponse) GetUrlIdWSOk() (*string, bool)`

GetUrlIdWSOk returns a tuple with the UrlIdWS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlIdWS

`func (o *GetV1PageLikesResponse) SetUrlIdWS(v string)`

SetUrlIdWS sets UrlIdWS field to given value.


### GetDidLike

`func (o *GetV1PageLikesResponse) GetDidLike() bool`

GetDidLike returns the DidLike field if non-nil, zero value otherwise.

### GetDidLikeOk

`func (o *GetV1PageLikesResponse) GetDidLikeOk() (*bool, bool)`

GetDidLikeOk returns a tuple with the DidLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidLike

`func (o *GetV1PageLikesResponse) SetDidLike(v bool)`

SetDidLike sets DidLike field to given value.


### GetCommentCount

`func (o *GetV1PageLikesResponse) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetV1PageLikesResponse) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetV1PageLikesResponse) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetLikeCount

`func (o *GetV1PageLikesResponse) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *GetV1PageLikesResponse) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *GetV1PageLikesResponse) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetStatus

`func (o *GetV1PageLikesResponse) GetStatus() APIStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetV1PageLikesResponse) GetStatusOk() (*APIStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetV1PageLikesResponse) SetStatus(v APIStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetV1PageLikesResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetV1PageLikesResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetV1PageLikesResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *GetV1PageLikesResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetV1PageLikesResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetV1PageLikesResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetSecondaryCode

`func (o *GetV1PageLikesResponse) GetSecondaryCode() string`

GetSecondaryCode returns the SecondaryCode field if non-nil, zero value otherwise.

### GetSecondaryCodeOk

`func (o *GetV1PageLikesResponse) GetSecondaryCodeOk() (*string, bool)`

GetSecondaryCodeOk returns a tuple with the SecondaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCode

`func (o *GetV1PageLikesResponse) SetSecondaryCode(v string)`

SetSecondaryCode sets SecondaryCode field to given value.

### HasSecondaryCode

`func (o *GetV1PageLikesResponse) HasSecondaryCode() bool`

HasSecondaryCode returns a boolean if a field has been set.

### GetBannedUntil

`func (o *GetV1PageLikesResponse) GetBannedUntil() int64`

GetBannedUntil returns the BannedUntil field if non-nil, zero value otherwise.

### GetBannedUntilOk

`func (o *GetV1PageLikesResponse) GetBannedUntilOk() (*int64, bool)`

GetBannedUntilOk returns a tuple with the BannedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedUntil

`func (o *GetV1PageLikesResponse) SetBannedUntil(v int64)`

SetBannedUntil sets BannedUntil field to given value.

### HasBannedUntil

`func (o *GetV1PageLikesResponse) HasBannedUntil() bool`

HasBannedUntil returns a boolean if a field has been set.

### GetMaxCharacterLength

`func (o *GetV1PageLikesResponse) GetMaxCharacterLength() int32`

GetMaxCharacterLength returns the MaxCharacterLength field if non-nil, zero value otherwise.

### GetMaxCharacterLengthOk

`func (o *GetV1PageLikesResponse) GetMaxCharacterLengthOk() (*int32, bool)`

GetMaxCharacterLengthOk returns a tuple with the MaxCharacterLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacterLength

`func (o *GetV1PageLikesResponse) SetMaxCharacterLength(v int32)`

SetMaxCharacterLength sets MaxCharacterLength field to given value.

### HasMaxCharacterLength

`func (o *GetV1PageLikesResponse) HasMaxCharacterLength() bool`

HasMaxCharacterLength returns a boolean if a field has been set.

### GetTranslatedError

`func (o *GetV1PageLikesResponse) GetTranslatedError() string`

GetTranslatedError returns the TranslatedError field if non-nil, zero value otherwise.

### GetTranslatedErrorOk

`func (o *GetV1PageLikesResponse) GetTranslatedErrorOk() (*string, bool)`

GetTranslatedErrorOk returns a tuple with the TranslatedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedError

`func (o *GetV1PageLikesResponse) SetTranslatedError(v string)`

SetTranslatedError sets TranslatedError field to given value.

### HasTranslatedError

`func (o *GetV1PageLikesResponse) HasTranslatedError() bool`

HasTranslatedError returns a boolean if a field has been set.

### GetCustomConfig

`func (o *GetV1PageLikesResponse) GetCustomConfig() CustomConfigParameters`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetV1PageLikesResponse) GetCustomConfigOk() (*CustomConfigParameters, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetV1PageLikesResponse) SetCustomConfig(v CustomConfigParameters)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetV1PageLikesResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


