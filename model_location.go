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

// Location struct for Location
type Location struct {
	LocationId *int32 `json:"location_id,omitempty"`
	LocationName string `json:"location_name"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation(locationName string) *Location {
	this := Location{}
	this.LocationName = locationName
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *Location) GetLocationId() int32 {
	if o == nil || o.LocationId == nil {
		var ret int32
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLocationIdOk() (*int32, bool) {
	if o == nil || o.LocationId == nil {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *Location) HasLocationId() bool {
	if o != nil && o.LocationId != nil {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given int32 and assigns it to the LocationId field.
func (o *Location) SetLocationId(v int32) {
	o.LocationId = &v
}

// GetLocationName returns the LocationName field value
func (o *Location) GetLocationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value
// and a boolean to check if the value has been set.
func (o *Location) GetLocationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LocationName, true
}

// SetLocationName sets field value
func (o *Location) SetLocationName(v string) {
	o.LocationName = v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocationId != nil {
		toSerialize["location_id"] = o.LocationId
	}
	if true {
		toSerialize["location_name"] = o.LocationName
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


