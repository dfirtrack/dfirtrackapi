# Case

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | Pointer to **int32** |  | [optional] [readonly] 
**CaseName** | **string** |  | 
**CaseIsIncident** | **bool** |  | 
**CaseCreatedByUserId** | **int32** |  | 
**CaseCreateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


