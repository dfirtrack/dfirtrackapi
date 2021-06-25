# Reportitem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportitemId** | Pointer to **int32** |  | [optional] [readonly] 
**Case** | Pointer to **NullableInt32** |  | [optional] 
**Headline** | **int32** |  | 
**Notestatus** | Pointer to **int32** |  | [optional] 
**System** | **int32** |  | 
**Tag** | Pointer to **[]int32** |  | [optional] 
**ReportitemSubheadline** | **NullableString** |  | 
**ReportitemNote** | **string** |  | 
**ReportitemCreateTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**ReportitemCreatedByUserId** | **int32** |  | 
**ReportitemModifyTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**ReportitemModifiedByUserId** | **int32** |  | 

## Methods

### NewReportitem

`func NewReportitem(headline int32, system int32, reportitemSubheadline NullableString, reportitemNote string, reportitemCreatedByUserId int32, reportitemModifiedByUserId int32, ) *Reportitem`

NewReportitem instantiates a new Reportitem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportitemWithDefaults

`func NewReportitemWithDefaults() *Reportitem`

NewReportitemWithDefaults instantiates a new Reportitem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportitemId

`func (o *Reportitem) GetReportitemId() int32`

GetReportitemId returns the ReportitemId field if non-nil, zero value otherwise.

### GetReportitemIdOk

`func (o *Reportitem) GetReportitemIdOk() (*int32, bool)`

GetReportitemIdOk returns a tuple with the ReportitemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemId

`func (o *Reportitem) SetReportitemId(v int32)`

SetReportitemId sets ReportitemId field to given value.

### HasReportitemId

`func (o *Reportitem) HasReportitemId() bool`

HasReportitemId returns a boolean if a field has been set.

### GetCase

`func (o *Reportitem) GetCase() int32`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *Reportitem) GetCaseOk() (*int32, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *Reportitem) SetCase(v int32)`

SetCase sets Case field to given value.

### HasCase

`func (o *Reportitem) HasCase() bool`

HasCase returns a boolean if a field has been set.

### SetCaseNil

`func (o *Reportitem) SetCaseNil(b bool)`

 SetCaseNil sets the value for Case to be an explicit nil

### UnsetCase
`func (o *Reportitem) UnsetCase()`

UnsetCase ensures that no value is present for Case, not even an explicit nil
### GetHeadline

`func (o *Reportitem) GetHeadline() int32`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *Reportitem) GetHeadlineOk() (*int32, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *Reportitem) SetHeadline(v int32)`

SetHeadline sets Headline field to given value.


### GetNotestatus

`func (o *Reportitem) GetNotestatus() int32`

GetNotestatus returns the Notestatus field if non-nil, zero value otherwise.

### GetNotestatusOk

`func (o *Reportitem) GetNotestatusOk() (*int32, bool)`

GetNotestatusOk returns a tuple with the Notestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotestatus

`func (o *Reportitem) SetNotestatus(v int32)`

SetNotestatus sets Notestatus field to given value.

### HasNotestatus

`func (o *Reportitem) HasNotestatus() bool`

HasNotestatus returns a boolean if a field has been set.

### GetSystem

`func (o *Reportitem) GetSystem() int32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Reportitem) GetSystemOk() (*int32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Reportitem) SetSystem(v int32)`

SetSystem sets System field to given value.


### GetTag

`func (o *Reportitem) GetTag() []int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Reportitem) GetTagOk() (*[]int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Reportitem) SetTag(v []int32)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Reportitem) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetReportitemSubheadline

`func (o *Reportitem) GetReportitemSubheadline() string`

GetReportitemSubheadline returns the ReportitemSubheadline field if non-nil, zero value otherwise.

### GetReportitemSubheadlineOk

`func (o *Reportitem) GetReportitemSubheadlineOk() (*string, bool)`

GetReportitemSubheadlineOk returns a tuple with the ReportitemSubheadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemSubheadline

`func (o *Reportitem) SetReportitemSubheadline(v string)`

SetReportitemSubheadline sets ReportitemSubheadline field to given value.


### SetReportitemSubheadlineNil

`func (o *Reportitem) SetReportitemSubheadlineNil(b bool)`

 SetReportitemSubheadlineNil sets the value for ReportitemSubheadline to be an explicit nil

### UnsetReportitemSubheadline
`func (o *Reportitem) UnsetReportitemSubheadline()`

UnsetReportitemSubheadline ensures that no value is present for ReportitemSubheadline, not even an explicit nil
### GetReportitemNote

`func (o *Reportitem) GetReportitemNote() string`

GetReportitemNote returns the ReportitemNote field if non-nil, zero value otherwise.

### GetReportitemNoteOk

`func (o *Reportitem) GetReportitemNoteOk() (*string, bool)`

GetReportitemNoteOk returns a tuple with the ReportitemNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemNote

`func (o *Reportitem) SetReportitemNote(v string)`

SetReportitemNote sets ReportitemNote field to given value.


### GetReportitemCreateTime

`func (o *Reportitem) GetReportitemCreateTime() time.Time`

GetReportitemCreateTime returns the ReportitemCreateTime field if non-nil, zero value otherwise.

### GetReportitemCreateTimeOk

`func (o *Reportitem) GetReportitemCreateTimeOk() (*time.Time, bool)`

GetReportitemCreateTimeOk returns a tuple with the ReportitemCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemCreateTime

`func (o *Reportitem) SetReportitemCreateTime(v time.Time)`

SetReportitemCreateTime sets ReportitemCreateTime field to given value.

### HasReportitemCreateTime

`func (o *Reportitem) HasReportitemCreateTime() bool`

HasReportitemCreateTime returns a boolean if a field has been set.

### GetReportitemCreatedByUserId

`func (o *Reportitem) GetReportitemCreatedByUserId() int32`

GetReportitemCreatedByUserId returns the ReportitemCreatedByUserId field if non-nil, zero value otherwise.

### GetReportitemCreatedByUserIdOk

`func (o *Reportitem) GetReportitemCreatedByUserIdOk() (*int32, bool)`

GetReportitemCreatedByUserIdOk returns a tuple with the ReportitemCreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemCreatedByUserId

`func (o *Reportitem) SetReportitemCreatedByUserId(v int32)`

SetReportitemCreatedByUserId sets ReportitemCreatedByUserId field to given value.


### GetReportitemModifyTime

`func (o *Reportitem) GetReportitemModifyTime() time.Time`

GetReportitemModifyTime returns the ReportitemModifyTime field if non-nil, zero value otherwise.

### GetReportitemModifyTimeOk

`func (o *Reportitem) GetReportitemModifyTimeOk() (*time.Time, bool)`

GetReportitemModifyTimeOk returns a tuple with the ReportitemModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemModifyTime

`func (o *Reportitem) SetReportitemModifyTime(v time.Time)`

SetReportitemModifyTime sets ReportitemModifyTime field to given value.

### HasReportitemModifyTime

`func (o *Reportitem) HasReportitemModifyTime() bool`

HasReportitemModifyTime returns a boolean if a field has been set.

### GetReportitemModifiedByUserId

`func (o *Reportitem) GetReportitemModifiedByUserId() int32`

GetReportitemModifiedByUserId returns the ReportitemModifiedByUserId field if non-nil, zero value otherwise.

### GetReportitemModifiedByUserIdOk

`func (o *Reportitem) GetReportitemModifiedByUserIdOk() (*int32, bool)`

GetReportitemModifiedByUserIdOk returns a tuple with the ReportitemModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportitemModifiedByUserId

`func (o *Reportitem) SetReportitemModifiedByUserId(v int32)`

SetReportitemModifiedByUserId sets ReportitemModifiedByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


