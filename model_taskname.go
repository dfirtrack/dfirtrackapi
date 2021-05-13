/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v1.2.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Taskname struct for Taskname
type Taskname struct {
	TasknameId *int32 `json:"taskname_id,omitempty"`
	TasknameName string `json:"taskname_name"`
}

// NewTaskname instantiates a new Taskname object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskname(tasknameName string, ) *Taskname {
	this := Taskname{}
	this.TasknameName = tasknameName
	return &this
}

// NewTasknameWithDefaults instantiates a new Taskname object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasknameWithDefaults() *Taskname {
	this := Taskname{}
	return &this
}

// GetTasknameId returns the TasknameId field value if set, zero value otherwise.
func (o *Taskname) GetTasknameId() int32 {
	if o == nil || o.TasknameId == nil {
		var ret int32
		return ret
	}
	return *o.TasknameId
}

// GetTasknameIdOk returns a tuple with the TasknameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taskname) GetTasknameIdOk() (*int32, bool) {
	if o == nil || o.TasknameId == nil {
		return nil, false
	}
	return o.TasknameId, true
}

// HasTasknameId returns a boolean if a field has been set.
func (o *Taskname) HasTasknameId() bool {
	if o != nil && o.TasknameId != nil {
		return true
	}

	return false
}

// SetTasknameId gets a reference to the given int32 and assigns it to the TasknameId field.
func (o *Taskname) SetTasknameId(v int32) {
	o.TasknameId = &v
}

// GetTasknameName returns the TasknameName field value
func (o *Taskname) GetTasknameName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TasknameName
}

// GetTasknameNameOk returns a tuple with the TasknameName field value
// and a boolean to check if the value has been set.
func (o *Taskname) GetTasknameNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TasknameName, true
}

// SetTasknameName sets field value
func (o *Taskname) SetTasknameName(v string) {
	o.TasknameName = v
}

func (o Taskname) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TasknameId != nil {
		toSerialize["taskname_id"] = o.TasknameId
	}
	if true {
		toSerialize["taskname_name"] = o.TasknameName
	}
	return json.Marshal(toSerialize)
}

type NullableTaskname struct {
	value *Taskname
	isSet bool
}

func (v NullableTaskname) Get() *Taskname {
	return v.value
}

func (v *NullableTaskname) Set(val *Taskname) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskname) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskname) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskname(val *Taskname) *NullableTaskname {
	return &NullableTaskname{value: val, isSet: true}
}

func (v NullableTaskname) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskname) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


