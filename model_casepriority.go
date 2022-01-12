/*
DFIRTrack

OpenAPI 3 - Documentation of DFIRTrack API

API version: v2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Casepriority struct for Casepriority
type Casepriority struct {
	CasepriorityId *int32 `json:"casepriority_id,omitempty"`
	CasepriorityName string `json:"casepriority_name"`
}

// NewCasepriority instantiates a new Casepriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCasepriority(casepriorityName string) *Casepriority {
	this := Casepriority{}
	this.CasepriorityName = casepriorityName
	return &this
}

// NewCasepriorityWithDefaults instantiates a new Casepriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCasepriorityWithDefaults() *Casepriority {
	this := Casepriority{}
	return &this
}

// GetCasepriorityId returns the CasepriorityId field value if set, zero value otherwise.
func (o *Casepriority) GetCasepriorityId() int32 {
	if o == nil || o.CasepriorityId == nil {
		var ret int32
		return ret
	}
	return *o.CasepriorityId
}

// GetCasepriorityIdOk returns a tuple with the CasepriorityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Casepriority) GetCasepriorityIdOk() (*int32, bool) {
	if o == nil || o.CasepriorityId == nil {
		return nil, false
	}
	return o.CasepriorityId, true
}

// HasCasepriorityId returns a boolean if a field has been set.
func (o *Casepriority) HasCasepriorityId() bool {
	if o != nil && o.CasepriorityId != nil {
		return true
	}

	return false
}

// SetCasepriorityId gets a reference to the given int32 and assigns it to the CasepriorityId field.
func (o *Casepriority) SetCasepriorityId(v int32) {
	o.CasepriorityId = &v
}

// GetCasepriorityName returns the CasepriorityName field value
func (o *Casepriority) GetCasepriorityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CasepriorityName
}

// GetCasepriorityNameOk returns a tuple with the CasepriorityName field value
// and a boolean to check if the value has been set.
func (o *Casepriority) GetCasepriorityNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CasepriorityName, true
}

// SetCasepriorityName sets field value
func (o *Casepriority) SetCasepriorityName(v string) {
	o.CasepriorityName = v
}

func (o Casepriority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CasepriorityId != nil {
		toSerialize["casepriority_id"] = o.CasepriorityId
	}
	if true {
		toSerialize["casepriority_name"] = o.CasepriorityName
	}
	return json.Marshal(toSerialize)
}

type NullableCasepriority struct {
	value *Casepriority
	isSet bool
}

func (v NullableCasepriority) Get() *Casepriority {
	return v.value
}

func (v *NullableCasepriority) Set(val *Casepriority) {
	v.value = val
	v.isSet = true
}

func (v NullableCasepriority) IsSet() bool {
	return v.isSet
}

func (v *NullableCasepriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCasepriority(val *Casepriority) *NullableCasepriority {
	return &NullableCasepriority{value: val, isSet: true}
}

func (v NullableCasepriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCasepriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


