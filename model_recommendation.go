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

// Recommendation struct for Recommendation
type Recommendation struct {
	RecommendationId *int32 `json:"recommendation_id,omitempty"`
	RecommendationName string `json:"recommendation_name"`
}

// NewRecommendation instantiates a new Recommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendation(recommendationName string, ) *Recommendation {
	this := Recommendation{}
	this.RecommendationName = recommendationName
	return &this
}

// NewRecommendationWithDefaults instantiates a new Recommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationWithDefaults() *Recommendation {
	this := Recommendation{}
	return &this
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *Recommendation) GetRecommendationId() int32 {
	if o == nil || o.RecommendationId == nil {
		var ret int32
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recommendation) GetRecommendationIdOk() (*int32, bool) {
	if o == nil || o.RecommendationId == nil {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *Recommendation) HasRecommendationId() bool {
	if o != nil && o.RecommendationId != nil {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given int32 and assigns it to the RecommendationId field.
func (o *Recommendation) SetRecommendationId(v int32) {
	o.RecommendationId = &v
}

// GetRecommendationName returns the RecommendationName field value
func (o *Recommendation) GetRecommendationName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RecommendationName
}

// GetRecommendationNameOk returns a tuple with the RecommendationName field value
// and a boolean to check if the value has been set.
func (o *Recommendation) GetRecommendationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecommendationName, true
}

// SetRecommendationName sets field value
func (o *Recommendation) SetRecommendationName(v string) {
	o.RecommendationName = v
}

func (o Recommendation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecommendationId != nil {
		toSerialize["recommendation_id"] = o.RecommendationId
	}
	if true {
		toSerialize["recommendation_name"] = o.RecommendationName
	}
	return json.Marshal(toSerialize)
}

type NullableRecommendation struct {
	value *Recommendation
	isSet bool
}

func (v NullableRecommendation) Get() *Recommendation {
	return v.value
}

func (v *NullableRecommendation) Set(val *Recommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendation(val *Recommendation) *NullableRecommendation {
	return &NullableRecommendation{value: val, isSet: true}
}

func (v NullableRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


