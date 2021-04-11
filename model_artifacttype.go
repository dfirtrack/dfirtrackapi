/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Artifacttype struct for Artifacttype
type Artifacttype struct {
	ArtifacttypeName string `json:"artifacttype_name"`
}

// NewArtifacttype instantiates a new Artifacttype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifacttype(artifacttypeName string, ) *Artifacttype {
	this := Artifacttype{}
	this.ArtifacttypeName = artifacttypeName
	return &this
}

// NewArtifacttypeWithDefaults instantiates a new Artifacttype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifacttypeWithDefaults() *Artifacttype {
	this := Artifacttype{}
	return &this
}

// GetArtifacttypeName returns the ArtifacttypeName field value
func (o *Artifacttype) GetArtifacttypeName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifacttypeName
}

// GetArtifacttypeNameOk returns a tuple with the ArtifacttypeName field value
// and a boolean to check if the value has been set.
func (o *Artifacttype) GetArtifacttypeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifacttypeName, true
}

// SetArtifacttypeName sets field value
func (o *Artifacttype) SetArtifacttypeName(v string) {
	o.ArtifacttypeName = v
}

func (o Artifacttype) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifacttype_name"] = o.ArtifacttypeName
	}
	return json.Marshal(toSerialize)
}

type NullableArtifacttype struct {
	value *Artifacttype
	isSet bool
}

func (v NullableArtifacttype) Get() *Artifacttype {
	return v.value
}

func (v *NullableArtifacttype) Set(val *Artifacttype) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifacttype) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifacttype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifacttype(val *Artifacttype) *NullableArtifacttype {
	return &NullableArtifacttype{value: val, isSet: true}
}

func (v NullableArtifacttype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifacttype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


