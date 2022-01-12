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

// Taskpriority struct for Taskpriority
type Taskpriority struct {
	TaskpriorityId *int32 `json:"taskpriority_id,omitempty"`
	TaskpriorityName string `json:"taskpriority_name"`
}

// NewTaskpriority instantiates a new Taskpriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskpriority(taskpriorityName string) *Taskpriority {
	this := Taskpriority{}
	this.TaskpriorityName = taskpriorityName
	return &this
}

// NewTaskpriorityWithDefaults instantiates a new Taskpriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskpriorityWithDefaults() *Taskpriority {
	this := Taskpriority{}
	return &this
}

// GetTaskpriorityId returns the TaskpriorityId field value if set, zero value otherwise.
func (o *Taskpriority) GetTaskpriorityId() int32 {
	if o == nil || o.TaskpriorityId == nil {
		var ret int32
		return ret
	}
	return *o.TaskpriorityId
}

// GetTaskpriorityIdOk returns a tuple with the TaskpriorityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taskpriority) GetTaskpriorityIdOk() (*int32, bool) {
	if o == nil || o.TaskpriorityId == nil {
		return nil, false
	}
	return o.TaskpriorityId, true
}

// HasTaskpriorityId returns a boolean if a field has been set.
func (o *Taskpriority) HasTaskpriorityId() bool {
	if o != nil && o.TaskpriorityId != nil {
		return true
	}

	return false
}

// SetTaskpriorityId gets a reference to the given int32 and assigns it to the TaskpriorityId field.
func (o *Taskpriority) SetTaskpriorityId(v int32) {
	o.TaskpriorityId = &v
}

// GetTaskpriorityName returns the TaskpriorityName field value
func (o *Taskpriority) GetTaskpriorityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskpriorityName
}

// GetTaskpriorityNameOk returns a tuple with the TaskpriorityName field value
// and a boolean to check if the value has been set.
func (o *Taskpriority) GetTaskpriorityNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TaskpriorityName, true
}

// SetTaskpriorityName sets field value
func (o *Taskpriority) SetTaskpriorityName(v string) {
	o.TaskpriorityName = v
}

func (o Taskpriority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaskpriorityId != nil {
		toSerialize["taskpriority_id"] = o.TaskpriorityId
	}
	if true {
		toSerialize["taskpriority_name"] = o.TaskpriorityName
	}
	return json.Marshal(toSerialize)
}

type NullableTaskpriority struct {
	value *Taskpriority
	isSet bool
}

func (v NullableTaskpriority) Get() *Taskpriority {
	return v.value
}

func (v *NullableTaskpriority) Set(val *Taskpriority) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskpriority) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskpriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskpriority(val *Taskpriority) *NullableTaskpriority {
	return &NullableTaskpriority{value: val, isSet: true}
}

func (v NullableTaskpriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskpriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


