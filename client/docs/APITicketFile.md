# APITicketFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**S3Key** | **string** |  | 
**OriginalFileName** | **string** |  | 
**SizeBytes** | **int32** |  | 
**ContentType** | **string** |  | 
**UploadedByUserId** | **string** |  | 
**UploadedAt** | **string** |  | 
**Url** | **string** |  | 
**ExpiresAt** | **string** |  | 
**Expired** | Pointer to **bool** |  | [optional] 

## Methods

### NewAPITicketFile

`func NewAPITicketFile(id string, s3Key string, originalFileName string, sizeBytes int32, contentType string, uploadedByUserId string, uploadedAt string, url string, expiresAt string, ) *APITicketFile`

NewAPITicketFile instantiates a new APITicketFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITicketFileWithDefaults

`func NewAPITicketFileWithDefaults() *APITicketFile`

NewAPITicketFileWithDefaults instantiates a new APITicketFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APITicketFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APITicketFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APITicketFile) SetId(v string)`

SetId sets Id field to given value.


### GetS3Key

`func (o *APITicketFile) GetS3Key() string`

GetS3Key returns the S3Key field if non-nil, zero value otherwise.

### GetS3KeyOk

`func (o *APITicketFile) GetS3KeyOk() (*string, bool)`

GetS3KeyOk returns a tuple with the S3Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Key

`func (o *APITicketFile) SetS3Key(v string)`

SetS3Key sets S3Key field to given value.


### GetOriginalFileName

`func (o *APITicketFile) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *APITicketFile) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *APITicketFile) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.


### GetSizeBytes

`func (o *APITicketFile) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *APITicketFile) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *APITicketFile) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetContentType

`func (o *APITicketFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *APITicketFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *APITicketFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetUploadedByUserId

`func (o *APITicketFile) GetUploadedByUserId() string`

GetUploadedByUserId returns the UploadedByUserId field if non-nil, zero value otherwise.

### GetUploadedByUserIdOk

`func (o *APITicketFile) GetUploadedByUserIdOk() (*string, bool)`

GetUploadedByUserIdOk returns a tuple with the UploadedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedByUserId

`func (o *APITicketFile) SetUploadedByUserId(v string)`

SetUploadedByUserId sets UploadedByUserId field to given value.


### GetUploadedAt

`func (o *APITicketFile) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *APITicketFile) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *APITicketFile) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.


### GetUrl

`func (o *APITicketFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *APITicketFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *APITicketFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *APITicketFile) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *APITicketFile) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *APITicketFile) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExpired

`func (o *APITicketFile) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *APITicketFile) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *APITicketFile) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *APITicketFile) HasExpired() bool`

HasExpired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


