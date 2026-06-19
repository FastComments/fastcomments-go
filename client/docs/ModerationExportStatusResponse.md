# ModerationExportStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**JobStatus** | **string** |  | 
**RecordCount** | **int32** |  | 
**DownloadUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewModerationExportStatusResponse

`func NewModerationExportStatusResponse(status string, jobStatus string, recordCount int32, ) *ModerationExportStatusResponse`

NewModerationExportStatusResponse instantiates a new ModerationExportStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationExportStatusResponseWithDefaults

`func NewModerationExportStatusResponseWithDefaults() *ModerationExportStatusResponse`

NewModerationExportStatusResponseWithDefaults instantiates a new ModerationExportStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModerationExportStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModerationExportStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModerationExportStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetJobStatus

`func (o *ModerationExportStatusResponse) GetJobStatus() string`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *ModerationExportStatusResponse) GetJobStatusOk() (*string, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *ModerationExportStatusResponse) SetJobStatus(v string)`

SetJobStatus sets JobStatus field to given value.


### GetRecordCount

`func (o *ModerationExportStatusResponse) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *ModerationExportStatusResponse) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *ModerationExportStatusResponse) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.


### GetDownloadUrl

`func (o *ModerationExportStatusResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ModerationExportStatusResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ModerationExportStatusResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *ModerationExportStatusResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


