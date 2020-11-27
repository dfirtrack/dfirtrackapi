/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: 0.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi-api-go-client

import (
	"encoding/json"
)

// Osarch struct for Osarch
type Osarch struct {
	OsarchId *int32 `json:"osarch_id,omitempty"`
	OsarchName string `json:"osarch_name"`
}

// NewOsarch instantiates a new Osarch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsarch(osarchName string, ) *Osarch {
	this := Osarch{}
	this.OsarchName = osarchName
	return &this
}

// NewOsarchWithDefaults instantiates a new Osarch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsarchWithDefaults() *Osarch {
	this := Osarch{}
	return &this
}

// GetOsarchId returns the OsarchId field value if set, zero value otherwise.
func (o *Osarch) GetOsarchId() int32 {
	if o == nil || o.OsarchId == nil {
		var ret int32
		return ret
	}
	return *o.OsarchId
}

// GetOsarchIdOk returns a tuple with the OsarchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Osarch) GetOsarchIdOk() (*int32, bool) {
	if o == nil || o.OsarchId == nil {
		return nil, false
	}
	return o.OsarchId, true
}

// HasOsarchId returns a boolean if a field has been set.
func (o *Osarch) HasOsarchId() bool {
	if o != nil && o.OsarchId != nil {
		return true
	}

	return false
}

// SetOsarchId gets a reference to the given int32 and assigns it to the OsarchId field.
func (o *Osarch) SetOsarchId(v int32) {
	o.OsarchId = &v
}

// GetOsarchName returns the OsarchName field value
func (o *Osarch) GetOsarchName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.OsarchName
}

// GetOsarchNameOk returns a tuple with the OsarchName field value
// and a boolean to check if the value has been set.
func (o *Osarch) GetOsarchNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsarchName, true
}

// SetOsarchName sets field value
func (o *Osarch) SetOsarchName(v string) {
	o.OsarchName = v
}

func (o Osarch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OsarchId != nil {
		toSerialize["osarch_id"] = o.OsarchId
	}
	if true {
		toSerialize["osarch_name"] = o.OsarchName
	}
	return json.Marshal(toSerialize)
}

type NullableOsarch struct {
	value *Osarch
	isSet bool
}

func (v NullableOsarch) Get() *Osarch {
	return v.value
}

func (v *NullableOsarch) Set(val *Osarch) {
	v.value = val
	v.isSet = true
}

func (v NullableOsarch) IsSet() bool {
	return v.isSet
}

func (v *NullableOsarch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsarch(val *Osarch) *NullableOsarch {
	return &NullableOsarch{value: val, isSet: true}
}

func (v NullableOsarch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsarch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


