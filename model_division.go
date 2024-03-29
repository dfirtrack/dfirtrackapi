/*
DFIRTrack

OpenAPI 3 - Documentation of DFIRTrack API

API version: v2.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Division struct for Division
type Division struct {
	DivisionId *int32 `json:"division_id,omitempty"`
	DivisionName string `json:"division_name"`
}

// NewDivision instantiates a new Division object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDivision(divisionName string) *Division {
	this := Division{}
	this.DivisionName = divisionName
	return &this
}

// NewDivisionWithDefaults instantiates a new Division object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDivisionWithDefaults() *Division {
	this := Division{}
	return &this
}

// GetDivisionId returns the DivisionId field value if set, zero value otherwise.
func (o *Division) GetDivisionId() int32 {
	if o == nil || o.DivisionId == nil {
		var ret int32
		return ret
	}
	return *o.DivisionId
}

// GetDivisionIdOk returns a tuple with the DivisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Division) GetDivisionIdOk() (*int32, bool) {
	if o == nil || o.DivisionId == nil {
		return nil, false
	}
	return o.DivisionId, true
}

// HasDivisionId returns a boolean if a field has been set.
func (o *Division) HasDivisionId() bool {
	if o != nil && o.DivisionId != nil {
		return true
	}

	return false
}

// SetDivisionId gets a reference to the given int32 and assigns it to the DivisionId field.
func (o *Division) SetDivisionId(v int32) {
	o.DivisionId = &v
}

// GetDivisionName returns the DivisionName field value
func (o *Division) GetDivisionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DivisionName
}

// GetDivisionNameOk returns a tuple with the DivisionName field value
// and a boolean to check if the value has been set.
func (o *Division) GetDivisionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DivisionName, true
}

// SetDivisionName sets field value
func (o *Division) SetDivisionName(v string) {
	o.DivisionName = v
}

func (o Division) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DivisionId != nil {
		toSerialize["division_id"] = o.DivisionId
	}
	if true {
		toSerialize["division_name"] = o.DivisionName
	}
	return json.Marshal(toSerialize)
}

type NullableDivision struct {
	value *Division
	isSet bool
}

func (v NullableDivision) Get() *Division {
	return v.value
}

func (v *NullableDivision) Set(val *Division) {
	v.value = val
	v.isSet = true
}

func (v NullableDivision) IsSet() bool {
	return v.isSet
}

func (v *NullableDivision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDivision(val *Division) *NullableDivision {
	return &NullableDivision{value: val, isSet: true}
}

func (v NullableDivision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDivision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


