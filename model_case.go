/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: 0.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrack-api-client

import (
	"encoding/json"
	"time"
)

// Case struct for Case
type Case struct {
	CaseId *int32 `json:"case_id,omitempty"`
	CaseName string `json:"case_name"`
	CaseIsIncident bool `json:"case_is_incident"`
	CaseCreatedByUserId int32 `json:"case_created_by_user_id"`
	CaseCreateTime *time.Time `json:"case_create_time,omitempty"`
}

// NewCase instantiates a new Case object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCase(caseName string, caseIsIncident bool, caseCreatedByUserId int32, ) *Case {
	this := Case{}
	this.CaseName = caseName
	this.CaseIsIncident = caseIsIncident
	this.CaseCreatedByUserId = caseCreatedByUserId
	return &this
}

// NewCaseWithDefaults instantiates a new Case object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseWithDefaults() *Case {
	this := Case{}
	return &this
}

// GetCaseId returns the CaseId field value if set, zero value otherwise.
func (o *Case) GetCaseId() int32 {
	if o == nil || o.CaseId == nil {
		var ret int32
		return ret
	}
	return *o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetCaseIdOk() (*int32, bool) {
	if o == nil || o.CaseId == nil {
		return nil, false
	}
	return o.CaseId, true
}

// HasCaseId returns a boolean if a field has been set.
func (o *Case) HasCaseId() bool {
	if o != nil && o.CaseId != nil {
		return true
	}

	return false
}

// SetCaseId gets a reference to the given int32 and assigns it to the CaseId field.
func (o *Case) SetCaseId(v int32) {
	o.CaseId = &v
}

// GetCaseName returns the CaseName field value
func (o *Case) GetCaseName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CaseName
}

// GetCaseNameOk returns a tuple with the CaseName field value
// and a boolean to check if the value has been set.
func (o *Case) GetCaseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CaseName, true
}

// SetCaseName sets field value
func (o *Case) SetCaseName(v string) {
	o.CaseName = v
}

// GetCaseIsIncident returns the CaseIsIncident field value
func (o *Case) GetCaseIsIncident() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.CaseIsIncident
}

// GetCaseIsIncidentOk returns a tuple with the CaseIsIncident field value
// and a boolean to check if the value has been set.
func (o *Case) GetCaseIsIncidentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CaseIsIncident, true
}

// SetCaseIsIncident sets field value
func (o *Case) SetCaseIsIncident(v bool) {
	o.CaseIsIncident = v
}

// GetCaseCreatedByUserId returns the CaseCreatedByUserId field value
func (o *Case) GetCaseCreatedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CaseCreatedByUserId
}

// GetCaseCreatedByUserIdOk returns a tuple with the CaseCreatedByUserId field value
// and a boolean to check if the value has been set.
func (o *Case) GetCaseCreatedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CaseCreatedByUserId, true
}

// SetCaseCreatedByUserId sets field value
func (o *Case) SetCaseCreatedByUserId(v int32) {
	o.CaseCreatedByUserId = v
}

// GetCaseCreateTime returns the CaseCreateTime field value if set, zero value otherwise.
func (o *Case) GetCaseCreateTime() time.Time {
	if o == nil || o.CaseCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CaseCreateTime
}

// GetCaseCreateTimeOk returns a tuple with the CaseCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetCaseCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CaseCreateTime == nil {
		return nil, false
	}
	return o.CaseCreateTime, true
}

// HasCaseCreateTime returns a boolean if a field has been set.
func (o *Case) HasCaseCreateTime() bool {
	if o != nil && o.CaseCreateTime != nil {
		return true
	}

	return false
}

// SetCaseCreateTime gets a reference to the given time.Time and assigns it to the CaseCreateTime field.
func (o *Case) SetCaseCreateTime(v time.Time) {
	o.CaseCreateTime = &v
}

func (o Case) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaseId != nil {
		toSerialize["case_id"] = o.CaseId
	}
	if true {
		toSerialize["case_name"] = o.CaseName
	}
	if true {
		toSerialize["case_is_incident"] = o.CaseIsIncident
	}
	if true {
		toSerialize["case_created_by_user_id"] = o.CaseCreatedByUserId
	}
	if o.CaseCreateTime != nil {
		toSerialize["case_create_time"] = o.CaseCreateTime
	}
	return json.Marshal(toSerialize)
}

type NullableCase struct {
	value *Case
	isSet bool
}

func (v NullableCase) Get() *Case {
	return v.value
}

func (v *NullableCase) Set(val *Case) {
	v.value = val
	v.isSet = true
}

func (v NullableCase) IsSet() bool {
	return v.isSet
}

func (v *NullableCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCase(val *Case) *NullableCase {
	return &NullableCase{value: val, isSet: true}
}

func (v NullableCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

