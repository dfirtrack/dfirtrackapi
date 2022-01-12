/*
DFIRTrack

OpenAPI 3 - Documentation of DFIRTrack API

API version: v2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Company struct for Company
type Company struct {
	CompanyId *int32 `json:"company_id,omitempty"`
	CompanyName string `json:"company_name"`
	Division NullableInt32 `json:"division,omitempty"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany(companyName string) *Company {
	this := Company{}
	this.CompanyName = companyName
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *Company) GetCompanyId() int32 {
	if o == nil || o.CompanyId == nil {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCompanyIdOk() (*int32, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *Company) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *Company) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetCompanyName returns the CompanyName field value
func (o *Company) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *Company) GetCompanyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *Company) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetDivision returns the Division field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetDivision() int32 {
	if o == nil || o.Division.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Division.Get()
}

// GetDivisionOk returns a tuple with the Division field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetDivisionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Division.Get(), o.Division.IsSet()
}

// HasDivision returns a boolean if a field has been set.
func (o *Company) HasDivision() bool {
	if o != nil && o.Division.IsSet() {
		return true
	}

	return false
}

// SetDivision gets a reference to the given NullableInt32 and assigns it to the Division field.
func (o *Company) SetDivision(v int32) {
	o.Division.Set(&v)
}
// SetDivisionNil sets the value for Division to be an explicit nil
func (o *Company) SetDivisionNil() {
	o.Division.Set(nil)
}

// UnsetDivision ensures that no value is present for Division, not even an explicit nil
func (o *Company) UnsetDivision() {
	o.Division.Unset()
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["company_name"] = o.CompanyName
	}
	if o.Division.IsSet() {
		toSerialize["division"] = o.Division.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


