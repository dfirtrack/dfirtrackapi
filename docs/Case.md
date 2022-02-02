# Case

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | Pointer to **int32** |  | [optional] [readonly] 
**CaseIdExternal** | Pointer to **NullableString** |  | [optional] 
**CaseName** | **string** |  | 
**Casepriority** | Pointer to **int32** |  | [optional] 
**Casestatus** | Pointer to **int32** |  | [optional] 
**Casetype** | Pointer to **NullableInt32** |  | [optional] 
**Tag** | Pointer to **[]int32** |  | [optional] 
**CaseIsIncident** | **bool** |  | 
**CaseStartTime** | Pointer to **NullableTime** |  | [optional] 
**CaseEndTime** | Pointer to **NullableTime** |  | [optional] 
**CaseCreatedByUserId** | **int32** |  | 
**CaseCreateTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**CaseModifiedByUserId** | Pointer to **NullableInt32** |  | [optional] 
**CaseModifyTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**CaseAssignedToUserId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCase

`func NewCase(caseName string, caseIsIncident bool, caseCreatedByUserId int32, ) *Case`

NewCase instantiates a new Case object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseWithDefaults

`func NewCaseWithDefaults() *Case`

NewCaseWithDefaults instantiates a new Case object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *Case) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *Case) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *Case) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *Case) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetCaseIdExternal

`func (o *Case) GetCaseIdExternal() string`

GetCaseIdExternal returns the CaseIdExternal field if non-nil, zero value otherwise.

### GetCaseIdExternalOk

`func (o *Case) GetCaseIdExternalOk() (*string, bool)`

GetCaseIdExternalOk returns a tuple with the CaseIdExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIdExternal

`func (o *Case) SetCaseIdExternal(v string)`

SetCaseIdExternal sets CaseIdExternal field to given value.

### HasCaseIdExternal

`func (o *Case) HasCaseIdExternal() bool`

HasCaseIdExternal returns a boolean if a field has been set.

### SetCaseIdExternalNil

`func (o *Case) SetCaseIdExternalNil(b bool)`

 SetCaseIdExternalNil sets the value for CaseIdExternal to be an explicit nil

### UnsetCaseIdExternal
`func (o *Case) UnsetCaseIdExternal()`

UnsetCaseIdExternal ensures that no value is present for CaseIdExternal, not even an explicit nil
### GetCaseName

`func (o *Case) GetCaseName() string`

GetCaseName returns the CaseName field if non-nil, zero value otherwise.

### GetCaseNameOk

`func (o *Case) GetCaseNameOk() (*string, bool)`

GetCaseNameOk returns a tuple with the CaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseName

`func (o *Case) SetCaseName(v string)`

SetCaseName sets CaseName field to given value.


### GetCasepriority

`func (o *Case) GetCasepriority() int32`

GetCasepriority returns the Casepriority field if non-nil, zero value otherwise.

### GetCasepriorityOk

`func (o *Case) GetCasepriorityOk() (*int32, bool)`

GetCasepriorityOk returns a tuple with the Casepriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasepriority

`func (o *Case) SetCasepriority(v int32)`

SetCasepriority sets Casepriority field to given value.

### HasCasepriority

`func (o *Case) HasCasepriority() bool`

HasCasepriority returns a boolean if a field has been set.

### GetCasestatus

`func (o *Case) GetCasestatus() int32`

GetCasestatus returns the Casestatus field if non-nil, zero value otherwise.

### GetCasestatusOk

`func (o *Case) GetCasestatusOk() (*int32, bool)`

GetCasestatusOk returns a tuple with the Casestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasestatus

`func (o *Case) SetCasestatus(v int32)`

SetCasestatus sets Casestatus field to given value.

### HasCasestatus

`func (o *Case) HasCasestatus() bool`

HasCasestatus returns a boolean if a field has been set.

### GetCasetype

`func (o *Case) GetCasetype() int32`

GetCasetype returns the Casetype field if non-nil, zero value otherwise.

### GetCasetypeOk

`func (o *Case) GetCasetypeOk() (*int32, bool)`

GetCasetypeOk returns a tuple with the Casetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasetype

`func (o *Case) SetCasetype(v int32)`

SetCasetype sets Casetype field to given value.

### HasCasetype

`func (o *Case) HasCasetype() bool`

HasCasetype returns a boolean if a field has been set.

### SetCasetypeNil

`func (o *Case) SetCasetypeNil(b bool)`

 SetCasetypeNil sets the value for Casetype to be an explicit nil

### UnsetCasetype
`func (o *Case) UnsetCasetype()`

UnsetCasetype ensures that no value is present for Casetype, not even an explicit nil
### GetTag

`func (o *Case) GetTag() []int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Case) GetTagOk() (*[]int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Case) SetTag(v []int32)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Case) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCaseIsIncident

`func (o *Case) GetCaseIsIncident() bool`

GetCaseIsIncident returns the CaseIsIncident field if non-nil, zero value otherwise.

### GetCaseIsIncidentOk

`func (o *Case) GetCaseIsIncidentOk() (*bool, bool)`

GetCaseIsIncidentOk returns a tuple with the CaseIsIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIsIncident

`func (o *Case) SetCaseIsIncident(v bool)`

SetCaseIsIncident sets CaseIsIncident field to given value.


### GetCaseStartTime

`func (o *Case) GetCaseStartTime() time.Time`

GetCaseStartTime returns the CaseStartTime field if non-nil, zero value otherwise.

### GetCaseStartTimeOk

`func (o *Case) GetCaseStartTimeOk() (*time.Time, bool)`

GetCaseStartTimeOk returns a tuple with the CaseStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseStartTime

`func (o *Case) SetCaseStartTime(v time.Time)`

SetCaseStartTime sets CaseStartTime field to given value.

### HasCaseStartTime

`func (o *Case) HasCaseStartTime() bool`

HasCaseStartTime returns a boolean if a field has been set.

### SetCaseStartTimeNil

`func (o *Case) SetCaseStartTimeNil(b bool)`

 SetCaseStartTimeNil sets the value for CaseStartTime to be an explicit nil

### UnsetCaseStartTime
`func (o *Case) UnsetCaseStartTime()`

UnsetCaseStartTime ensures that no value is present for CaseStartTime, not even an explicit nil
### GetCaseEndTime

`func (o *Case) GetCaseEndTime() time.Time`

GetCaseEndTime returns the CaseEndTime field if non-nil, zero value otherwise.

### GetCaseEndTimeOk

`func (o *Case) GetCaseEndTimeOk() (*time.Time, bool)`

GetCaseEndTimeOk returns a tuple with the CaseEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseEndTime

`func (o *Case) SetCaseEndTime(v time.Time)`

SetCaseEndTime sets CaseEndTime field to given value.

### HasCaseEndTime

`func (o *Case) HasCaseEndTime() bool`

HasCaseEndTime returns a boolean if a field has been set.

### SetCaseEndTimeNil

`func (o *Case) SetCaseEndTimeNil(b bool)`

 SetCaseEndTimeNil sets the value for CaseEndTime to be an explicit nil

### UnsetCaseEndTime
`func (o *Case) UnsetCaseEndTime()`

UnsetCaseEndTime ensures that no value is present for CaseEndTime, not even an explicit nil
### GetCaseCreatedByUserId

`func (o *Case) GetCaseCreatedByUserId() int32`

GetCaseCreatedByUserId returns the CaseCreatedByUserId field if non-nil, zero value otherwise.

### GetCaseCreatedByUserIdOk

`func (o *Case) GetCaseCreatedByUserIdOk() (*int32, bool)`

GetCaseCreatedByUserIdOk returns a tuple with the CaseCreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseCreatedByUserId

`func (o *Case) SetCaseCreatedByUserId(v int32)`

SetCaseCreatedByUserId sets CaseCreatedByUserId field to given value.


### GetCaseCreateTime

`func (o *Case) GetCaseCreateTime() time.Time`

GetCaseCreateTime returns the CaseCreateTime field if non-nil, zero value otherwise.

### GetCaseCreateTimeOk

`func (o *Case) GetCaseCreateTimeOk() (*time.Time, bool)`

GetCaseCreateTimeOk returns a tuple with the CaseCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseCreateTime

`func (o *Case) SetCaseCreateTime(v time.Time)`

SetCaseCreateTime sets CaseCreateTime field to given value.

### HasCaseCreateTime

`func (o *Case) HasCaseCreateTime() bool`

HasCaseCreateTime returns a boolean if a field has been set.

### GetCaseModifiedByUserId

`func (o *Case) GetCaseModifiedByUserId() int32`

GetCaseModifiedByUserId returns the CaseModifiedByUserId field if non-nil, zero value otherwise.

### GetCaseModifiedByUserIdOk

`func (o *Case) GetCaseModifiedByUserIdOk() (*int32, bool)`

GetCaseModifiedByUserIdOk returns a tuple with the CaseModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseModifiedByUserId

`func (o *Case) SetCaseModifiedByUserId(v int32)`

SetCaseModifiedByUserId sets CaseModifiedByUserId field to given value.

### HasCaseModifiedByUserId

`func (o *Case) HasCaseModifiedByUserId() bool`

HasCaseModifiedByUserId returns a boolean if a field has been set.

### SetCaseModifiedByUserIdNil

`func (o *Case) SetCaseModifiedByUserIdNil(b bool)`

 SetCaseModifiedByUserIdNil sets the value for CaseModifiedByUserId to be an explicit nil

### UnsetCaseModifiedByUserId
`func (o *Case) UnsetCaseModifiedByUserId()`

UnsetCaseModifiedByUserId ensures that no value is present for CaseModifiedByUserId, not even an explicit nil
### GetCaseModifyTime

`func (o *Case) GetCaseModifyTime() time.Time`

GetCaseModifyTime returns the CaseModifyTime field if non-nil, zero value otherwise.

### GetCaseModifyTimeOk

`func (o *Case) GetCaseModifyTimeOk() (*time.Time, bool)`

GetCaseModifyTimeOk returns a tuple with the CaseModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseModifyTime

`func (o *Case) SetCaseModifyTime(v time.Time)`

SetCaseModifyTime sets CaseModifyTime field to given value.

### HasCaseModifyTime

`func (o *Case) HasCaseModifyTime() bool`

HasCaseModifyTime returns a boolean if a field has been set.

### GetCaseAssignedToUserId

`func (o *Case) GetCaseAssignedToUserId() int32`

GetCaseAssignedToUserId returns the CaseAssignedToUserId field if non-nil, zero value otherwise.

### GetCaseAssignedToUserIdOk

`func (o *Case) GetCaseAssignedToUserIdOk() (*int32, bool)`

GetCaseAssignedToUserIdOk returns a tuple with the CaseAssignedToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseAssignedToUserId

`func (o *Case) SetCaseAssignedToUserId(v int32)`

SetCaseAssignedToUserId sets CaseAssignedToUserId field to given value.

### HasCaseAssignedToUserId

`func (o *Case) HasCaseAssignedToUserId() bool`

HasCaseAssignedToUserId returns a boolean if a field has been set.

### SetCaseAssignedToUserIdNil

`func (o *Case) SetCaseAssignedToUserIdNil(b bool)`

 SetCaseAssignedToUserIdNil sets the value for CaseAssignedToUserId to be an explicit nil

### UnsetCaseAssignedToUserId
`func (o *Case) UnsetCaseAssignedToUserId()`

UnsetCaseAssignedToUserId ensures that no value is present for CaseAssignedToUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


