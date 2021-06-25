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

// Artifactstatus struct for Artifactstatus
type Artifactstatus struct {
	ArtifactstatusId *int32 `json:"artifactstatus_id,omitempty"`
	ArtifactstatusName string `json:"artifactstatus_name"`
}

// NewArtifactstatus instantiates a new Artifactstatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactstatus(artifactstatusName string, ) *Artifactstatus {
	this := Artifactstatus{}
	this.ArtifactstatusName = artifactstatusName
	return &this
}

// NewArtifactstatusWithDefaults instantiates a new Artifactstatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactstatusWithDefaults() *Artifactstatus {
	this := Artifactstatus{}
	return &this
}

// GetArtifactstatusId returns the ArtifactstatusId field value if set, zero value otherwise.
func (o *Artifactstatus) GetArtifactstatusId() int32 {
	if o == nil || o.ArtifactstatusId == nil {
		var ret int32
		return ret
	}
	return *o.ArtifactstatusId
}

// GetArtifactstatusIdOk returns a tuple with the ArtifactstatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifactstatus) GetArtifactstatusIdOk() (*int32, bool) {
	if o == nil || o.ArtifactstatusId == nil {
		return nil, false
	}
	return o.ArtifactstatusId, true
}

// HasArtifactstatusId returns a boolean if a field has been set.
func (o *Artifactstatus) HasArtifactstatusId() bool {
	if o != nil && o.ArtifactstatusId != nil {
		return true
	}

	return false
}

// SetArtifactstatusId gets a reference to the given int32 and assigns it to the ArtifactstatusId field.
func (o *Artifactstatus) SetArtifactstatusId(v int32) {
	o.ArtifactstatusId = &v
}

// GetArtifactstatusName returns the ArtifactstatusName field value
func (o *Artifactstatus) GetArtifactstatusName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactstatusName
}

// GetArtifactstatusNameOk returns a tuple with the ArtifactstatusName field value
// and a boolean to check if the value has been set.
func (o *Artifactstatus) GetArtifactstatusNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactstatusName, true
}

// SetArtifactstatusName sets field value
func (o *Artifactstatus) SetArtifactstatusName(v string) {
	o.ArtifactstatusName = v
}

func (o Artifactstatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactstatusId != nil {
		toSerialize["artifactstatus_id"] = o.ArtifactstatusId
	}
	if true {
		toSerialize["artifactstatus_name"] = o.ArtifactstatusName
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactstatus struct {
	value *Artifactstatus
	isSet bool
}

func (v NullableArtifactstatus) Get() *Artifactstatus {
	return v.value
}

func (v *NullableArtifactstatus) Set(val *Artifactstatus) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactstatus) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactstatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactstatus(val *Artifactstatus) *NullableArtifactstatus {
	return &NullableArtifactstatus{value: val, isSet: true}
}

func (v NullableArtifactstatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactstatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


