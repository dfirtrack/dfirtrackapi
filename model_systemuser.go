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
	"time"
)

// Systemuser struct for Systemuser
type Systemuser struct {
	SystemuserId *int32 `json:"systemuser_id,omitempty"`
	SystemuserName string `json:"systemuser_name"`
	System int32 `json:"system"`
	SystemuserLastlogonTime NullableTime `json:"systemuser_lastlogon_time,omitempty"`
	SystemuserIsSystemadmin NullableBool `json:"systemuser_is_systemadmin,omitempty"`
}

// NewSystemuser instantiates a new Systemuser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemuser(systemuserName string, system int32, ) *Systemuser {
	this := Systemuser{}
	this.SystemuserName = systemuserName
	this.System = system
	return &this
}

// NewSystemuserWithDefaults instantiates a new Systemuser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemuserWithDefaults() *Systemuser {
	this := Systemuser{}
	return &this
}

// GetSystemuserId returns the SystemuserId field value if set, zero value otherwise.
func (o *Systemuser) GetSystemuserId() int32 {
	if o == nil || o.SystemuserId == nil {
		var ret int32
		return ret
	}
	return *o.SystemuserId
}

// GetSystemuserIdOk returns a tuple with the SystemuserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuser) GetSystemuserIdOk() (*int32, bool) {
	if o == nil || o.SystemuserId == nil {
		return nil, false
	}
	return o.SystemuserId, true
}

// HasSystemuserId returns a boolean if a field has been set.
func (o *Systemuser) HasSystemuserId() bool {
	if o != nil && o.SystemuserId != nil {
		return true
	}

	return false
}

// SetSystemuserId gets a reference to the given int32 and assigns it to the SystemuserId field.
func (o *Systemuser) SetSystemuserId(v int32) {
	o.SystemuserId = &v
}

// GetSystemuserName returns the SystemuserName field value
func (o *Systemuser) GetSystemuserName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SystemuserName
}

// GetSystemuserNameOk returns a tuple with the SystemuserName field value
// and a boolean to check if the value has been set.
func (o *Systemuser) GetSystemuserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemuserName, true
}

// SetSystemuserName sets field value
func (o *Systemuser) SetSystemuserName(v string) {
	o.SystemuserName = v
}

// GetSystem returns the System field value
func (o *Systemuser) GetSystem() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *Systemuser) GetSystemOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *Systemuser) SetSystem(v int32) {
	o.System = v
}

// GetSystemuserLastlogonTime returns the SystemuserLastlogonTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Systemuser) GetSystemuserLastlogonTime() time.Time {
	if o == nil || o.SystemuserLastlogonTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SystemuserLastlogonTime.Get()
}

// GetSystemuserLastlogonTimeOk returns a tuple with the SystemuserLastlogonTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Systemuser) GetSystemuserLastlogonTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemuserLastlogonTime.Get(), o.SystemuserLastlogonTime.IsSet()
}

// HasSystemuserLastlogonTime returns a boolean if a field has been set.
func (o *Systemuser) HasSystemuserLastlogonTime() bool {
	if o != nil && o.SystemuserLastlogonTime.IsSet() {
		return true
	}

	return false
}

// SetSystemuserLastlogonTime gets a reference to the given NullableTime and assigns it to the SystemuserLastlogonTime field.
func (o *Systemuser) SetSystemuserLastlogonTime(v time.Time) {
	o.SystemuserLastlogonTime.Set(&v)
}
// SetSystemuserLastlogonTimeNil sets the value for SystemuserLastlogonTime to be an explicit nil
func (o *Systemuser) SetSystemuserLastlogonTimeNil() {
	o.SystemuserLastlogonTime.Set(nil)
}

// UnsetSystemuserLastlogonTime ensures that no value is present for SystemuserLastlogonTime, not even an explicit nil
func (o *Systemuser) UnsetSystemuserLastlogonTime() {
	o.SystemuserLastlogonTime.Unset()
}

// GetSystemuserIsSystemadmin returns the SystemuserIsSystemadmin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Systemuser) GetSystemuserIsSystemadmin() bool {
	if o == nil || o.SystemuserIsSystemadmin.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SystemuserIsSystemadmin.Get()
}

// GetSystemuserIsSystemadminOk returns a tuple with the SystemuserIsSystemadmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Systemuser) GetSystemuserIsSystemadminOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemuserIsSystemadmin.Get(), o.SystemuserIsSystemadmin.IsSet()
}

// HasSystemuserIsSystemadmin returns a boolean if a field has been set.
func (o *Systemuser) HasSystemuserIsSystemadmin() bool {
	if o != nil && o.SystemuserIsSystemadmin.IsSet() {
		return true
	}

	return false
}

// SetSystemuserIsSystemadmin gets a reference to the given NullableBool and assigns it to the SystemuserIsSystemadmin field.
func (o *Systemuser) SetSystemuserIsSystemadmin(v bool) {
	o.SystemuserIsSystemadmin.Set(&v)
}
// SetSystemuserIsSystemadminNil sets the value for SystemuserIsSystemadmin to be an explicit nil
func (o *Systemuser) SetSystemuserIsSystemadminNil() {
	o.SystemuserIsSystemadmin.Set(nil)
}

// UnsetSystemuserIsSystemadmin ensures that no value is present for SystemuserIsSystemadmin, not even an explicit nil
func (o *Systemuser) UnsetSystemuserIsSystemadmin() {
	o.SystemuserIsSystemadmin.Unset()
}

func (o Systemuser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemuserId != nil {
		toSerialize["systemuser_id"] = o.SystemuserId
	}
	if true {
		toSerialize["systemuser_name"] = o.SystemuserName
	}
	if true {
		toSerialize["system"] = o.System
	}
	if o.SystemuserLastlogonTime.IsSet() {
		toSerialize["systemuser_lastlogon_time"] = o.SystemuserLastlogonTime.Get()
	}
	if o.SystemuserIsSystemadmin.IsSet() {
		toSerialize["systemuser_is_systemadmin"] = o.SystemuserIsSystemadmin.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSystemuser struct {
	value *Systemuser
	isSet bool
}

func (v NullableSystemuser) Get() *Systemuser {
	return v.value
}

func (v *NullableSystemuser) Set(val *Systemuser) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemuser) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemuser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemuser(val *Systemuser) *NullableSystemuser {
	return &NullableSystemuser{value: val, isSet: true}
}

func (v NullableSystemuser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemuser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


