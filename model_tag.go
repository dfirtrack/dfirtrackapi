/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v1.5.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Tag struct for Tag
type Tag struct {
	TagId *int32 `json:"tag_id,omitempty"`
	TagName string `json:"tag_name"`
	Tagcolor int32 `json:"tagcolor"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(tagName string, tagcolor int32, ) *Tag {
	this := Tag{}
	this.TagName = tagName
	this.Tagcolor = tagcolor
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *Tag) GetTagId() int32 {
	if o == nil || o.TagId == nil {
		var ret int32
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetTagIdOk() (*int32, bool) {
	if o == nil || o.TagId == nil {
		return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *Tag) HasTagId() bool {
	if o != nil && o.TagId != nil {
		return true
	}

	return false
}

// SetTagId gets a reference to the given int32 and assigns it to the TagId field.
func (o *Tag) SetTagId(v int32) {
	o.TagId = &v
}

// GetTagName returns the TagName field value
func (o *Tag) GetTagName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value
// and a boolean to check if the value has been set.
func (o *Tag) GetTagNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TagName, true
}

// SetTagName sets field value
func (o *Tag) SetTagName(v string) {
	o.TagName = v
}

// GetTagcolor returns the Tagcolor field value
func (o *Tag) GetTagcolor() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Tagcolor
}

// GetTagcolorOk returns a tuple with the Tagcolor field value
// and a boolean to check if the value has been set.
func (o *Tag) GetTagcolorOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tagcolor, true
}

// SetTagcolor sets field value
func (o *Tag) SetTagcolor(v int32) {
	o.Tagcolor = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TagId != nil {
		toSerialize["tag_id"] = o.TagId
	}
	if true {
		toSerialize["tag_name"] = o.TagName
	}
	if true {
		toSerialize["tagcolor"] = o.Tagcolor
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


