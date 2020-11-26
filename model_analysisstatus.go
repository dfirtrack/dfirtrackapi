/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: 0.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrack-api-go-client

import (
	"encoding/json"
)

// Analysisstatus struct for Analysisstatus
type Analysisstatus struct {
	AnalysisstatusId *int32 `json:"analysisstatus_id,omitempty"`
	AnalysisstatusName string `json:"analysisstatus_name"`
}

// NewAnalysisstatus instantiates a new Analysisstatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisstatus(analysisstatusName string, ) *Analysisstatus {
	this := Analysisstatus{}
	this.AnalysisstatusName = analysisstatusName
	return &this
}

// NewAnalysisstatusWithDefaults instantiates a new Analysisstatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisstatusWithDefaults() *Analysisstatus {
	this := Analysisstatus{}
	return &this
}

// GetAnalysisstatusId returns the AnalysisstatusId field value if set, zero value otherwise.
func (o *Analysisstatus) GetAnalysisstatusId() int32 {
	if o == nil || o.AnalysisstatusId == nil {
		var ret int32
		return ret
	}
	return *o.AnalysisstatusId
}

// GetAnalysisstatusIdOk returns a tuple with the AnalysisstatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Analysisstatus) GetAnalysisstatusIdOk() (*int32, bool) {
	if o == nil || o.AnalysisstatusId == nil {
		return nil, false
	}
	return o.AnalysisstatusId, true
}

// HasAnalysisstatusId returns a boolean if a field has been set.
func (o *Analysisstatus) HasAnalysisstatusId() bool {
	if o != nil && o.AnalysisstatusId != nil {
		return true
	}

	return false
}

// SetAnalysisstatusId gets a reference to the given int32 and assigns it to the AnalysisstatusId field.
func (o *Analysisstatus) SetAnalysisstatusId(v int32) {
	o.AnalysisstatusId = &v
}

// GetAnalysisstatusName returns the AnalysisstatusName field value
func (o *Analysisstatus) GetAnalysisstatusName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AnalysisstatusName
}

// GetAnalysisstatusNameOk returns a tuple with the AnalysisstatusName field value
// and a boolean to check if the value has been set.
func (o *Analysisstatus) GetAnalysisstatusNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AnalysisstatusName, true
}

// SetAnalysisstatusName sets field value
func (o *Analysisstatus) SetAnalysisstatusName(v string) {
	o.AnalysisstatusName = v
}

func (o Analysisstatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalysisstatusId != nil {
		toSerialize["analysisstatus_id"] = o.AnalysisstatusId
	}
	if true {
		toSerialize["analysisstatus_name"] = o.AnalysisstatusName
	}
	return json.Marshal(toSerialize)
}

type NullableAnalysisstatus struct {
	value *Analysisstatus
	isSet bool
}

func (v NullableAnalysisstatus) Get() *Analysisstatus {
	return v.value
}

func (v *NullableAnalysisstatus) Set(val *Analysisstatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisstatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisstatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisstatus(val *Analysisstatus) *NullableAnalysisstatus {
	return &NullableAnalysisstatus{value: val, isSet: true}
}

func (v NullableAnalysisstatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisstatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


