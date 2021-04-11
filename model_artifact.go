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
	"time"
)

// Artifact struct for Artifact
type Artifact struct {
	ArtifactId *int32 `json:"artifact_id,omitempty"`
	ArtifactUuid *string `json:"artifact_uuid,omitempty"`
	ArtifactName string `json:"artifact_name"`
	Artifactstatus *int32 `json:"artifactstatus,omitempty"`
	Artifacttype int32 `json:"artifacttype"`
	Case NullableInt32 `json:"case,omitempty"`
	System int32 `json:"system"`
	ArtifactMd5 NullableString `json:"artifact_md5,omitempty"`
	ArtifactSha1 NullableString `json:"artifact_sha1,omitempty"`
	ArtifactSha256 NullableString `json:"artifact_sha256,omitempty"`
	ArtifactSourcePath NullableString `json:"artifact_source_path,omitempty"`
	ArtifactStoragePath *string `json:"artifact_storage_path,omitempty"`
	ArtifactAcquisitionTime NullableTime `json:"artifact_acquisition_time,omitempty"`
	ArtifactRequestedTime NullableTime `json:"artifact_requested_time,omitempty"`
	ArtifactCreateTime *time.Time `json:"artifact_create_time,omitempty"`
	ArtifactCreatedByUserId int32 `json:"artifact_created_by_user_id"`
	ArtifactModifyTime *time.Time `json:"artifact_modify_time,omitempty"`
	ArtifactModifiedByUserId int32 `json:"artifact_modified_by_user_id"`
}

// NewArtifact instantiates a new Artifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifact(artifactName string, artifacttype int32, system int32, artifactCreatedByUserId int32, artifactModifiedByUserId int32, ) *Artifact {
	this := Artifact{}
	this.ArtifactName = artifactName
	this.Artifacttype = artifacttype
	this.System = system
	this.ArtifactCreatedByUserId = artifactCreatedByUserId
	this.ArtifactModifiedByUserId = artifactModifiedByUserId
	return &this
}

// NewArtifactWithDefaults instantiates a new Artifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactWithDefaults() *Artifact {
	this := Artifact{}
	return &this
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *Artifact) GetArtifactId() int32 {
	if o == nil || o.ArtifactId == nil {
		var ret int32
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactIdOk() (*int32, bool) {
	if o == nil || o.ArtifactId == nil {
		return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *Artifact) HasArtifactId() bool {
	if o != nil && o.ArtifactId != nil {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given int32 and assigns it to the ArtifactId field.
func (o *Artifact) SetArtifactId(v int32) {
	o.ArtifactId = &v
}

// GetArtifactUuid returns the ArtifactUuid field value if set, zero value otherwise.
func (o *Artifact) GetArtifactUuid() string {
	if o == nil || o.ArtifactUuid == nil {
		var ret string
		return ret
	}
	return *o.ArtifactUuid
}

// GetArtifactUuidOk returns a tuple with the ArtifactUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactUuidOk() (*string, bool) {
	if o == nil || o.ArtifactUuid == nil {
		return nil, false
	}
	return o.ArtifactUuid, true
}

// HasArtifactUuid returns a boolean if a field has been set.
func (o *Artifact) HasArtifactUuid() bool {
	if o != nil && o.ArtifactUuid != nil {
		return true
	}

	return false
}

// SetArtifactUuid gets a reference to the given string and assigns it to the ArtifactUuid field.
func (o *Artifact) SetArtifactUuid(v string) {
	o.ArtifactUuid = &v
}

// GetArtifactName returns the ArtifactName field value
func (o *Artifact) GetArtifactName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ArtifactName
}

// GetArtifactNameOk returns a tuple with the ArtifactName field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactName, true
}

// SetArtifactName sets field value
func (o *Artifact) SetArtifactName(v string) {
	o.ArtifactName = v
}

// GetArtifactstatus returns the Artifactstatus field value if set, zero value otherwise.
func (o *Artifact) GetArtifactstatus() int32 {
	if o == nil || o.Artifactstatus == nil {
		var ret int32
		return ret
	}
	return *o.Artifactstatus
}

// GetArtifactstatusOk returns a tuple with the Artifactstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactstatusOk() (*int32, bool) {
	if o == nil || o.Artifactstatus == nil {
		return nil, false
	}
	return o.Artifactstatus, true
}

// HasArtifactstatus returns a boolean if a field has been set.
func (o *Artifact) HasArtifactstatus() bool {
	if o != nil && o.Artifactstatus != nil {
		return true
	}

	return false
}

// SetArtifactstatus gets a reference to the given int32 and assigns it to the Artifactstatus field.
func (o *Artifact) SetArtifactstatus(v int32) {
	o.Artifactstatus = &v
}

// GetArtifacttype returns the Artifacttype field value
func (o *Artifact) GetArtifacttype() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Artifacttype
}

// GetArtifacttypeOk returns a tuple with the Artifacttype field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifacttypeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Artifacttype, true
}

// SetArtifacttype sets field value
func (o *Artifact) SetArtifacttype(v int32) {
	o.Artifacttype = v
}

// GetCase returns the Case field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetCase() int32 {
	if o == nil || o.Case.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Case.Get()
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetCaseOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Case.Get(), o.Case.IsSet()
}

// HasCase returns a boolean if a field has been set.
func (o *Artifact) HasCase() bool {
	if o != nil && o.Case.IsSet() {
		return true
	}

	return false
}

// SetCase gets a reference to the given NullableInt32 and assigns it to the Case field.
func (o *Artifact) SetCase(v int32) {
	o.Case.Set(&v)
}
// SetCaseNil sets the value for Case to be an explicit nil
func (o *Artifact) SetCaseNil() {
	o.Case.Set(nil)
}

// UnsetCase ensures that no value is present for Case, not even an explicit nil
func (o *Artifact) UnsetCase() {
	o.Case.Unset()
}

// GetSystem returns the System field value
func (o *Artifact) GetSystem() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetSystemOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *Artifact) SetSystem(v int32) {
	o.System = v
}

// GetArtifactMd5 returns the ArtifactMd5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetArtifactMd5() string {
	if o == nil || o.ArtifactMd5.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArtifactMd5.Get()
}

// GetArtifactMd5Ok returns a tuple with the ArtifactMd5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetArtifactMd5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArtifactMd5.Get(), o.ArtifactMd5.IsSet()
}

// HasArtifactMd5 returns a boolean if a field has been set.
func (o *Artifact) HasArtifactMd5() bool {
	if o != nil && o.ArtifactMd5.IsSet() {
		return true
	}

	return false
}

// SetArtifactMd5 gets a reference to the given NullableString and assigns it to the ArtifactMd5 field.
func (o *Artifact) SetArtifactMd5(v string) {
	o.ArtifactMd5.Set(&v)
}
// SetArtifactMd5Nil sets the value for ArtifactMd5 to be an explicit nil
func (o *Artifact) SetArtifactMd5Nil() {
	o.ArtifactMd5.Set(nil)
}

// UnsetArtifactMd5 ensures that no value is present for ArtifactMd5, not even an explicit nil
func (o *Artifact) UnsetArtifactMd5() {
	o.ArtifactMd5.Unset()
}

// GetArtifactSha1 returns the ArtifactSha1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetArtifactSha1() string {
	if o == nil || o.ArtifactSha1.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArtifactSha1.Get()
}

// GetArtifactSha1Ok returns a tuple with the ArtifactSha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetArtifactSha1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArtifactSha1.Get(), o.ArtifactSha1.IsSet()
}

// HasArtifactSha1 returns a boolean if a field has been set.
func (o *Artifact) HasArtifactSha1() bool {
	if o != nil && o.ArtifactSha1.IsSet() {
		return true
	}

	return false
}

// SetArtifactSha1 gets a reference to the given NullableString and assigns it to the ArtifactSha1 field.
func (o *Artifact) SetArtifactSha1(v string) {
	o.ArtifactSha1.Set(&v)
}
// SetArtifactSha1Nil sets the value for ArtifactSha1 to be an explicit nil
func (o *Artifact) SetArtifactSha1Nil() {
	o.ArtifactSha1.Set(nil)
}

// UnsetArtifactSha1 ensures that no value is present for ArtifactSha1, not even an explicit nil
func (o *Artifact) UnsetArtifactSha1() {
	o.ArtifactSha1.Unset()
}

// GetArtifactSha256 returns the ArtifactSha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetArtifactSha256() string {
	if o == nil || o.ArtifactSha256.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArtifactSha256.Get()
}

// GetArtifactSha256Ok returns a tuple with the ArtifactSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetArtifactSha256Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArtifactSha256.Get(), o.ArtifactSha256.IsSet()
}

// HasArtifactSha256 returns a boolean if a field has been set.
func (o *Artifact) HasArtifactSha256() bool {
	if o != nil && o.ArtifactSha256.IsSet() {
		return true
	}

	return false
}

// SetArtifactSha256 gets a reference to the given NullableString and assigns it to the ArtifactSha256 field.
func (o *Artifact) SetArtifactSha256(v string) {
	o.ArtifactSha256.Set(&v)
}
// SetArtifactSha256Nil sets the value for ArtifactSha256 to be an explicit nil
func (o *Artifact) SetArtifactSha256Nil() {
	o.ArtifactSha256.Set(nil)
}

// UnsetArtifactSha256 ensures that no value is present for ArtifactSha256, not even an explicit nil
func (o *Artifact) UnsetArtifactSha256() {
	o.ArtifactSha256.Unset()
}

// GetArtifactSourcePath returns the ArtifactSourcePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetArtifactSourcePath() string {
	if o == nil || o.ArtifactSourcePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArtifactSourcePath.Get()
}

// GetArtifactSourcePathOk returns a tuple with the ArtifactSourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetArtifactSourcePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArtifactSourcePath.Get(), o.ArtifactSourcePath.IsSet()
}

// HasArtifactSourcePath returns a boolean if a field has been set.
func (o *Artifact) HasArtifactSourcePath() bool {
	if o != nil && o.ArtifactSourcePath.IsSet() {
		return true
	}

	return false
}

// SetArtifactSourcePath gets a reference to the given NullableString and assigns it to the ArtifactSourcePath field.
func (o *Artifact) SetArtifactSourcePath(v string) {
	o.ArtifactSourcePath.Set(&v)
}
// SetArtifactSourcePathNil sets the value for ArtifactSourcePath to be an explicit nil
func (o *Artifact) SetArtifactSourcePathNil() {
	o.ArtifactSourcePath.Set(nil)
}

// UnsetArtifactSourcePath ensures that no value is present for ArtifactSourcePath, not even an explicit nil
func (o *Artifact) UnsetArtifactSourcePath() {
	o.ArtifactSourcePath.Unset()
}

// GetArtifactStoragePath returns the ArtifactStoragePath field value if set, zero value otherwise.
func (o *Artifact) GetArtifactStoragePath() string {
	if o == nil || o.ArtifactStoragePath == nil {
		var ret string
		return ret
	}
	return *o.ArtifactStoragePath
}

// GetArtifactStoragePathOk returns a tuple with the ArtifactStoragePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactStoragePathOk() (*string, bool) {
	if o == nil || o.ArtifactStoragePath == nil {
		return nil, false
	}
	return o.ArtifactStoragePath, true
}

// HasArtifactStoragePath returns a boolean if a field has been set.
func (o *Artifact) HasArtifactStoragePath() bool {
	if o != nil && o.ArtifactStoragePath != nil {
		return true
	}

	return false
}

// SetArtifactStoragePath gets a reference to the given string and assigns it to the ArtifactStoragePath field.
func (o *Artifact) SetArtifactStoragePath(v string) {
	o.ArtifactStoragePath = &v
}

// GetArtifactAcquisitionTime returns the ArtifactAcquisitionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetArtifactAcquisitionTime() time.Time {
	if o == nil || o.ArtifactAcquisitionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArtifactAcquisitionTime.Get()
}

// GetArtifactAcquisitionTimeOk returns a tuple with the ArtifactAcquisitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetArtifactAcquisitionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArtifactAcquisitionTime.Get(), o.ArtifactAcquisitionTime.IsSet()
}

// HasArtifactAcquisitionTime returns a boolean if a field has been set.
func (o *Artifact) HasArtifactAcquisitionTime() bool {
	if o != nil && o.ArtifactAcquisitionTime.IsSet() {
		return true
	}

	return false
}

// SetArtifactAcquisitionTime gets a reference to the given NullableTime and assigns it to the ArtifactAcquisitionTime field.
func (o *Artifact) SetArtifactAcquisitionTime(v time.Time) {
	o.ArtifactAcquisitionTime.Set(&v)
}
// SetArtifactAcquisitionTimeNil sets the value for ArtifactAcquisitionTime to be an explicit nil
func (o *Artifact) SetArtifactAcquisitionTimeNil() {
	o.ArtifactAcquisitionTime.Set(nil)
}

// UnsetArtifactAcquisitionTime ensures that no value is present for ArtifactAcquisitionTime, not even an explicit nil
func (o *Artifact) UnsetArtifactAcquisitionTime() {
	o.ArtifactAcquisitionTime.Unset()
}

// GetArtifactRequestedTime returns the ArtifactRequestedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Artifact) GetArtifactRequestedTime() time.Time {
	if o == nil || o.ArtifactRequestedTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ArtifactRequestedTime.Get()
}

// GetArtifactRequestedTimeOk returns a tuple with the ArtifactRequestedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Artifact) GetArtifactRequestedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArtifactRequestedTime.Get(), o.ArtifactRequestedTime.IsSet()
}

// HasArtifactRequestedTime returns a boolean if a field has been set.
func (o *Artifact) HasArtifactRequestedTime() bool {
	if o != nil && o.ArtifactRequestedTime.IsSet() {
		return true
	}

	return false
}

// SetArtifactRequestedTime gets a reference to the given NullableTime and assigns it to the ArtifactRequestedTime field.
func (o *Artifact) SetArtifactRequestedTime(v time.Time) {
	o.ArtifactRequestedTime.Set(&v)
}
// SetArtifactRequestedTimeNil sets the value for ArtifactRequestedTime to be an explicit nil
func (o *Artifact) SetArtifactRequestedTimeNil() {
	o.ArtifactRequestedTime.Set(nil)
}

// UnsetArtifactRequestedTime ensures that no value is present for ArtifactRequestedTime, not even an explicit nil
func (o *Artifact) UnsetArtifactRequestedTime() {
	o.ArtifactRequestedTime.Unset()
}

// GetArtifactCreateTime returns the ArtifactCreateTime field value if set, zero value otherwise.
func (o *Artifact) GetArtifactCreateTime() time.Time {
	if o == nil || o.ArtifactCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ArtifactCreateTime
}

// GetArtifactCreateTimeOk returns a tuple with the ArtifactCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.ArtifactCreateTime == nil {
		return nil, false
	}
	return o.ArtifactCreateTime, true
}

// HasArtifactCreateTime returns a boolean if a field has been set.
func (o *Artifact) HasArtifactCreateTime() bool {
	if o != nil && o.ArtifactCreateTime != nil {
		return true
	}

	return false
}

// SetArtifactCreateTime gets a reference to the given time.Time and assigns it to the ArtifactCreateTime field.
func (o *Artifact) SetArtifactCreateTime(v time.Time) {
	o.ArtifactCreateTime = &v
}

// GetArtifactCreatedByUserId returns the ArtifactCreatedByUserId field value
func (o *Artifact) GetArtifactCreatedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ArtifactCreatedByUserId
}

// GetArtifactCreatedByUserIdOk returns a tuple with the ArtifactCreatedByUserId field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactCreatedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactCreatedByUserId, true
}

// SetArtifactCreatedByUserId sets field value
func (o *Artifact) SetArtifactCreatedByUserId(v int32) {
	o.ArtifactCreatedByUserId = v
}

// GetArtifactModifyTime returns the ArtifactModifyTime field value if set, zero value otherwise.
func (o *Artifact) GetArtifactModifyTime() time.Time {
	if o == nil || o.ArtifactModifyTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ArtifactModifyTime
}

// GetArtifactModifyTimeOk returns a tuple with the ArtifactModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactModifyTimeOk() (*time.Time, bool) {
	if o == nil || o.ArtifactModifyTime == nil {
		return nil, false
	}
	return o.ArtifactModifyTime, true
}

// HasArtifactModifyTime returns a boolean if a field has been set.
func (o *Artifact) HasArtifactModifyTime() bool {
	if o != nil && o.ArtifactModifyTime != nil {
		return true
	}

	return false
}

// SetArtifactModifyTime gets a reference to the given time.Time and assigns it to the ArtifactModifyTime field.
func (o *Artifact) SetArtifactModifyTime(v time.Time) {
	o.ArtifactModifyTime = &v
}

// GetArtifactModifiedByUserId returns the ArtifactModifiedByUserId field value
func (o *Artifact) GetArtifactModifiedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ArtifactModifiedByUserId
}

// GetArtifactModifiedByUserIdOk returns a tuple with the ArtifactModifiedByUserId field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetArtifactModifiedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ArtifactModifiedByUserId, true
}

// SetArtifactModifiedByUserId sets field value
func (o *Artifact) SetArtifactModifiedByUserId(v int32) {
	o.ArtifactModifiedByUserId = v
}

func (o Artifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactId != nil {
		toSerialize["artifact_id"] = o.ArtifactId
	}
	if o.ArtifactUuid != nil {
		toSerialize["artifact_uuid"] = o.ArtifactUuid
	}
	if true {
		toSerialize["artifact_name"] = o.ArtifactName
	}
	if o.Artifactstatus != nil {
		toSerialize["artifactstatus"] = o.Artifactstatus
	}
	if true {
		toSerialize["artifacttype"] = o.Artifacttype
	}
	if o.Case.IsSet() {
		toSerialize["case"] = o.Case.Get()
	}
	if true {
		toSerialize["system"] = o.System
	}
	if o.ArtifactMd5.IsSet() {
		toSerialize["artifact_md5"] = o.ArtifactMd5.Get()
	}
	if o.ArtifactSha1.IsSet() {
		toSerialize["artifact_sha1"] = o.ArtifactSha1.Get()
	}
	if o.ArtifactSha256.IsSet() {
		toSerialize["artifact_sha256"] = o.ArtifactSha256.Get()
	}
	if o.ArtifactSourcePath.IsSet() {
		toSerialize["artifact_source_path"] = o.ArtifactSourcePath.Get()
	}
	if o.ArtifactStoragePath != nil {
		toSerialize["artifact_storage_path"] = o.ArtifactStoragePath
	}
	if o.ArtifactAcquisitionTime.IsSet() {
		toSerialize["artifact_acquisition_time"] = o.ArtifactAcquisitionTime.Get()
	}
	if o.ArtifactRequestedTime.IsSet() {
		toSerialize["artifact_requested_time"] = o.ArtifactRequestedTime.Get()
	}
	if o.ArtifactCreateTime != nil {
		toSerialize["artifact_create_time"] = o.ArtifactCreateTime
	}
	if true {
		toSerialize["artifact_created_by_user_id"] = o.ArtifactCreatedByUserId
	}
	if o.ArtifactModifyTime != nil {
		toSerialize["artifact_modify_time"] = o.ArtifactModifyTime
	}
	if true {
		toSerialize["artifact_modified_by_user_id"] = o.ArtifactModifiedByUserId
	}
	return json.Marshal(toSerialize)
}

type NullableArtifact struct {
	value *Artifact
	isSet bool
}

func (v NullableArtifact) Get() *Artifact {
	return v.value
}

func (v *NullableArtifact) Set(val *Artifact) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifact(val *Artifact) *NullableArtifact {
	return &NullableArtifact{value: val, isSet: true}
}

func (v NullableArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


