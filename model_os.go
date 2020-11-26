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

// Os struct for Os
type Os struct {
	OsId *int32 `json:"os_id,omitempty"`
	OsName string `json:"os_name"`
}

// NewOs instantiates a new Os object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOs(osName string, ) *Os {
	this := Os{}
	this.OsName = osName
	return &this
}

// NewOsWithDefaults instantiates a new Os object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsWithDefaults() *Os {
	this := Os{}
	return &this
}

// GetOsId returns the OsId field value if set, zero value otherwise.
func (o *Os) GetOsId() int32 {
	if o == nil || o.OsId == nil {
		var ret int32
		return ret
	}
	return *o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Os) GetOsIdOk() (*int32, bool) {
	if o == nil || o.OsId == nil {
		return nil, false
	}
	return o.OsId, true
}

// HasOsId returns a boolean if a field has been set.
func (o *Os) HasOsId() bool {
	if o != nil && o.OsId != nil {
		return true
	}

	return false
}

// SetOsId gets a reference to the given int32 and assigns it to the OsId field.
func (o *Os) SetOsId(v int32) {
	o.OsId = &v
}

// GetOsName returns the OsName field value
func (o *Os) GetOsName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value
// and a boolean to check if the value has been set.
func (o *Os) GetOsNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsName, true
}

// SetOsName sets field value
func (o *Os) SetOsName(v string) {
	o.OsName = v
}

func (o Os) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OsId != nil {
		toSerialize["os_id"] = o.OsId
	}
	if true {
		toSerialize["os_name"] = o.OsName
	}
	return json.Marshal(toSerialize)
}

type NullableOs struct {
	value *Os
	isSet bool
}

func (v NullableOs) Get() *Os {
	return v.value
}

func (v *NullableOs) Set(val *Os) {
	v.value = val
	v.isSet = true
}

func (v NullableOs) IsSet() bool {
	return v.isSet
}

func (v *NullableOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOs(val *Os) *NullableOs {
	return &NullableOs{value: val, isSet: true}
}

func (v NullableOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


