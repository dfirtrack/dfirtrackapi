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

// Serviceprovider struct for Serviceprovider
type Serviceprovider struct {
	ServiceproviderId *int32 `json:"serviceprovider_id,omitempty"`
	ServiceproviderName string `json:"serviceprovider_name"`
}

// NewServiceprovider instantiates a new Serviceprovider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceprovider(serviceproviderName string, ) *Serviceprovider {
	this := Serviceprovider{}
	this.ServiceproviderName = serviceproviderName
	return &this
}

// NewServiceproviderWithDefaults instantiates a new Serviceprovider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceproviderWithDefaults() *Serviceprovider {
	this := Serviceprovider{}
	return &this
}

// GetServiceproviderId returns the ServiceproviderId field value if set, zero value otherwise.
func (o *Serviceprovider) GetServiceproviderId() int32 {
	if o == nil || o.ServiceproviderId == nil {
		var ret int32
		return ret
	}
	return *o.ServiceproviderId
}

// GetServiceproviderIdOk returns a tuple with the ServiceproviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Serviceprovider) GetServiceproviderIdOk() (*int32, bool) {
	if o == nil || o.ServiceproviderId == nil {
		return nil, false
	}
	return o.ServiceproviderId, true
}

// HasServiceproviderId returns a boolean if a field has been set.
func (o *Serviceprovider) HasServiceproviderId() bool {
	if o != nil && o.ServiceproviderId != nil {
		return true
	}

	return false
}

// SetServiceproviderId gets a reference to the given int32 and assigns it to the ServiceproviderId field.
func (o *Serviceprovider) SetServiceproviderId(v int32) {
	o.ServiceproviderId = &v
}

// GetServiceproviderName returns the ServiceproviderName field value
func (o *Serviceprovider) GetServiceproviderName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ServiceproviderName
}

// GetServiceproviderNameOk returns a tuple with the ServiceproviderName field value
// and a boolean to check if the value has been set.
func (o *Serviceprovider) GetServiceproviderNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceproviderName, true
}

// SetServiceproviderName sets field value
func (o *Serviceprovider) SetServiceproviderName(v string) {
	o.ServiceproviderName = v
}

func (o Serviceprovider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceproviderId != nil {
		toSerialize["serviceprovider_id"] = o.ServiceproviderId
	}
	if true {
		toSerialize["serviceprovider_name"] = o.ServiceproviderName
	}
	return json.Marshal(toSerialize)
}

type NullableServiceprovider struct {
	value *Serviceprovider
	isSet bool
}

func (v NullableServiceprovider) Get() *Serviceprovider {
	return v.value
}

func (v *NullableServiceprovider) Set(val *Serviceprovider) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceprovider) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceprovider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceprovider(val *Serviceprovider) *NullableServiceprovider {
	return &NullableServiceprovider{value: val, isSet: true}
}

func (v NullableServiceprovider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceprovider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


