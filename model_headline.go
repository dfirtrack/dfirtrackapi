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

// Headline struct for Headline
type Headline struct {
	HeadlineId *int32 `json:"headline_id,omitempty"`
	HeadlineName string `json:"headline_name"`
}

// NewHeadline instantiates a new Headline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeadline(headlineName string) *Headline {
	this := Headline{}
	this.HeadlineName = headlineName
	return &this
}

// NewHeadlineWithDefaults instantiates a new Headline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadlineWithDefaults() *Headline {
	this := Headline{}
	return &this
}

// GetHeadlineId returns the HeadlineId field value if set, zero value otherwise.
func (o *Headline) GetHeadlineId() int32 {
	if o == nil || o.HeadlineId == nil {
		var ret int32
		return ret
	}
	return *o.HeadlineId
}

// GetHeadlineIdOk returns a tuple with the HeadlineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Headline) GetHeadlineIdOk() (*int32, bool) {
	if o == nil || o.HeadlineId == nil {
		return nil, false
	}
	return o.HeadlineId, true
}

// HasHeadlineId returns a boolean if a field has been set.
func (o *Headline) HasHeadlineId() bool {
	if o != nil && o.HeadlineId != nil {
		return true
	}

	return false
}

// SetHeadlineId gets a reference to the given int32 and assigns it to the HeadlineId field.
func (o *Headline) SetHeadlineId(v int32) {
	o.HeadlineId = &v
}

// GetHeadlineName returns the HeadlineName field value
func (o *Headline) GetHeadlineName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeadlineName
}

// GetHeadlineNameOk returns a tuple with the HeadlineName field value
// and a boolean to check if the value has been set.
func (o *Headline) GetHeadlineNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HeadlineName, true
}

// SetHeadlineName sets field value
func (o *Headline) SetHeadlineName(v string) {
	o.HeadlineName = v
}

func (o Headline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HeadlineId != nil {
		toSerialize["headline_id"] = o.HeadlineId
	}
	if true {
		toSerialize["headline_name"] = o.HeadlineName
	}
	return json.Marshal(toSerialize)
}

type NullableHeadline struct {
	value *Headline
	isSet bool
}

func (v NullableHeadline) Get() *Headline {
	return v.value
}

func (v *NullableHeadline) Set(val *Headline) {
	v.value = val
	v.isSet = true
}

func (v NullableHeadline) IsSet() bool {
	return v.isSet
}

func (v *NullableHeadline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeadline(val *Headline) *NullableHeadline {
	return &NullableHeadline{value: val, isSet: true}
}

func (v NullableHeadline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeadline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


