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
	"time"
)

// System struct for System
type System struct {
	SystemId *int32 `json:"system_id,omitempty"`
	SystemUuid *string `json:"system_uuid,omitempty"`
	SystemName string `json:"system_name"`
	Domain NullableInt32 `json:"domain,omitempty"`
	Dnsname NullableInt32 `json:"dnsname,omitempty"`
	Systemstatus int32 `json:"systemstatus"`
	Analysisstatus NullableInt32 `json:"analysisstatus,omitempty"`
	Reason NullableInt32 `json:"reason,omitempty"`
	Recommendation NullableInt32 `json:"recommendation,omitempty"`
	Systemtype NullableInt32 `json:"systemtype,omitempty"`
	Ip *[]int32 `json:"ip,omitempty"`
	Os NullableInt32 `json:"os,omitempty"`
	Osarch NullableInt32 `json:"osarch,omitempty"`
	SystemLastbootedTime NullableTime `json:"system_lastbooted_time,omitempty"`
	SystemDeprecatedTime NullableTime `json:"system_deprecated_time,omitempty"`
	SystemIsVm NullableBool `json:"system_is_vm,omitempty"`
	HostSystem NullableInt32 `json:"host_system,omitempty"`
	Company *[]int32 `json:"company,omitempty"`
	Location NullableInt32 `json:"location,omitempty"`
	Serviceprovider NullableInt32 `json:"serviceprovider,omitempty"`
	Contact NullableInt32 `json:"contact,omitempty"`
	Tag *[]int32 `json:"tag,omitempty"`
	Case *[]int32 `json:"case,omitempty"`
	SystemApiTime NullableTime `json:"system_api_time,omitempty"`
	SystemCreateTime *time.Time `json:"system_create_time,omitempty"`
	SystemCreatedByUserId int32 `json:"system_created_by_user_id"`
	SystemModifyTime time.Time `json:"system_modify_time"`
	SystemModifiedByUserId int32 `json:"system_modified_by_user_id"`
	SystemExportMarkdown *bool `json:"system_export_markdown,omitempty"`
	SystemExportSpreadsheet *bool `json:"system_export_spreadsheet,omitempty"`
}

// NewSystem instantiates a new System object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystem(systemName string, systemstatus int32, systemCreatedByUserId int32, systemModifyTime time.Time, systemModifiedByUserId int32, ) *System {
	this := System{}
	this.SystemName = systemName
	this.Systemstatus = systemstatus
	this.SystemCreatedByUserId = systemCreatedByUserId
	this.SystemModifyTime = systemModifyTime
	this.SystemModifiedByUserId = systemModifiedByUserId
	return &this
}

// NewSystemWithDefaults instantiates a new System object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemWithDefaults() *System {
	this := System{}
	return &this
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *System) GetSystemId() int32 {
	if o == nil || o.SystemId == nil {
		var ret int32
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemIdOk() (*int32, bool) {
	if o == nil || o.SystemId == nil {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *System) HasSystemId() bool {
	if o != nil && o.SystemId != nil {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given int32 and assigns it to the SystemId field.
func (o *System) SetSystemId(v int32) {
	o.SystemId = &v
}

// GetSystemUuid returns the SystemUuid field value if set, zero value otherwise.
func (o *System) GetSystemUuid() string {
	if o == nil || o.SystemUuid == nil {
		var ret string
		return ret
	}
	return *o.SystemUuid
}

// GetSystemUuidOk returns a tuple with the SystemUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemUuidOk() (*string, bool) {
	if o == nil || o.SystemUuid == nil {
		return nil, false
	}
	return o.SystemUuid, true
}

// HasSystemUuid returns a boolean if a field has been set.
func (o *System) HasSystemUuid() bool {
	if o != nil && o.SystemUuid != nil {
		return true
	}

	return false
}

// SetSystemUuid gets a reference to the given string and assigns it to the SystemUuid field.
func (o *System) SetSystemUuid(v string) {
	o.SystemUuid = &v
}

// GetSystemName returns the SystemName field value
func (o *System) GetSystemName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value
// and a boolean to check if the value has been set.
func (o *System) GetSystemNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemName, true
}

// SetSystemName sets field value
func (o *System) SetSystemName(v string) {
	o.SystemName = v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetDomain() int32 {
	if o == nil || o.Domain.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetDomainOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *System) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableInt32 and assigns it to the Domain field.
func (o *System) SetDomain(v int32) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *System) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *System) UnsetDomain() {
	o.Domain.Unset()
}

// GetDnsname returns the Dnsname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetDnsname() int32 {
	if o == nil || o.Dnsname.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Dnsname.Get()
}

// GetDnsnameOk returns a tuple with the Dnsname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetDnsnameOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dnsname.Get(), o.Dnsname.IsSet()
}

// HasDnsname returns a boolean if a field has been set.
func (o *System) HasDnsname() bool {
	if o != nil && o.Dnsname.IsSet() {
		return true
	}

	return false
}

// SetDnsname gets a reference to the given NullableInt32 and assigns it to the Dnsname field.
func (o *System) SetDnsname(v int32) {
	o.Dnsname.Set(&v)
}
// SetDnsnameNil sets the value for Dnsname to be an explicit nil
func (o *System) SetDnsnameNil() {
	o.Dnsname.Set(nil)
}

// UnsetDnsname ensures that no value is present for Dnsname, not even an explicit nil
func (o *System) UnsetDnsname() {
	o.Dnsname.Unset()
}

// GetSystemstatus returns the Systemstatus field value
func (o *System) GetSystemstatus() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Systemstatus
}

// GetSystemstatusOk returns a tuple with the Systemstatus field value
// and a boolean to check if the value has been set.
func (o *System) GetSystemstatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Systemstatus, true
}

// SetSystemstatus sets field value
func (o *System) SetSystemstatus(v int32) {
	o.Systemstatus = v
}

// GetAnalysisstatus returns the Analysisstatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetAnalysisstatus() int32 {
	if o == nil || o.Analysisstatus.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Analysisstatus.Get()
}

// GetAnalysisstatusOk returns a tuple with the Analysisstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetAnalysisstatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Analysisstatus.Get(), o.Analysisstatus.IsSet()
}

// HasAnalysisstatus returns a boolean if a field has been set.
func (o *System) HasAnalysisstatus() bool {
	if o != nil && o.Analysisstatus.IsSet() {
		return true
	}

	return false
}

// SetAnalysisstatus gets a reference to the given NullableInt32 and assigns it to the Analysisstatus field.
func (o *System) SetAnalysisstatus(v int32) {
	o.Analysisstatus.Set(&v)
}
// SetAnalysisstatusNil sets the value for Analysisstatus to be an explicit nil
func (o *System) SetAnalysisstatusNil() {
	o.Analysisstatus.Set(nil)
}

// UnsetAnalysisstatus ensures that no value is present for Analysisstatus, not even an explicit nil
func (o *System) UnsetAnalysisstatus() {
	o.Analysisstatus.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetReason() int32 {
	if o == nil || o.Reason.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetReasonOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *System) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableInt32 and assigns it to the Reason field.
func (o *System) SetReason(v int32) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *System) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *System) UnsetReason() {
	o.Reason.Unset()
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetRecommendation() int32 {
	if o == nil || o.Recommendation.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Recommendation.Get()
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetRecommendationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Recommendation.Get(), o.Recommendation.IsSet()
}

// HasRecommendation returns a boolean if a field has been set.
func (o *System) HasRecommendation() bool {
	if o != nil && o.Recommendation.IsSet() {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given NullableInt32 and assigns it to the Recommendation field.
func (o *System) SetRecommendation(v int32) {
	o.Recommendation.Set(&v)
}
// SetRecommendationNil sets the value for Recommendation to be an explicit nil
func (o *System) SetRecommendationNil() {
	o.Recommendation.Set(nil)
}

// UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
func (o *System) UnsetRecommendation() {
	o.Recommendation.Unset()
}

// GetSystemtype returns the Systemtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetSystemtype() int32 {
	if o == nil || o.Systemtype.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Systemtype.Get()
}

// GetSystemtypeOk returns a tuple with the Systemtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetSystemtypeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Systemtype.Get(), o.Systemtype.IsSet()
}

// HasSystemtype returns a boolean if a field has been set.
func (o *System) HasSystemtype() bool {
	if o != nil && o.Systemtype.IsSet() {
		return true
	}

	return false
}

// SetSystemtype gets a reference to the given NullableInt32 and assigns it to the Systemtype field.
func (o *System) SetSystemtype(v int32) {
	o.Systemtype.Set(&v)
}
// SetSystemtypeNil sets the value for Systemtype to be an explicit nil
func (o *System) SetSystemtypeNil() {
	o.Systemtype.Set(nil)
}

// UnsetSystemtype ensures that no value is present for Systemtype, not even an explicit nil
func (o *System) UnsetSystemtype() {
	o.Systemtype.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *System) GetIp() []int32 {
	if o == nil || o.Ip == nil {
		var ret []int32
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetIpOk() (*[]int32, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *System) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given []int32 and assigns it to the Ip field.
func (o *System) SetIp(v []int32) {
	o.Ip = &v
}

// GetOs returns the Os field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetOs() int32 {
	if o == nil || o.Os.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Os.Get()
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetOsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Os.Get(), o.Os.IsSet()
}

// HasOs returns a boolean if a field has been set.
func (o *System) HasOs() bool {
	if o != nil && o.Os.IsSet() {
		return true
	}

	return false
}

// SetOs gets a reference to the given NullableInt32 and assigns it to the Os field.
func (o *System) SetOs(v int32) {
	o.Os.Set(&v)
}
// SetOsNil sets the value for Os to be an explicit nil
func (o *System) SetOsNil() {
	o.Os.Set(nil)
}

// UnsetOs ensures that no value is present for Os, not even an explicit nil
func (o *System) UnsetOs() {
	o.Os.Unset()
}

// GetOsarch returns the Osarch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetOsarch() int32 {
	if o == nil || o.Osarch.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Osarch.Get()
}

// GetOsarchOk returns a tuple with the Osarch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetOsarchOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Osarch.Get(), o.Osarch.IsSet()
}

// HasOsarch returns a boolean if a field has been set.
func (o *System) HasOsarch() bool {
	if o != nil && o.Osarch.IsSet() {
		return true
	}

	return false
}

// SetOsarch gets a reference to the given NullableInt32 and assigns it to the Osarch field.
func (o *System) SetOsarch(v int32) {
	o.Osarch.Set(&v)
}
// SetOsarchNil sets the value for Osarch to be an explicit nil
func (o *System) SetOsarchNil() {
	o.Osarch.Set(nil)
}

// UnsetOsarch ensures that no value is present for Osarch, not even an explicit nil
func (o *System) UnsetOsarch() {
	o.Osarch.Unset()
}

// GetSystemLastbootedTime returns the SystemLastbootedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetSystemLastbootedTime() time.Time {
	if o == nil || o.SystemLastbootedTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SystemLastbootedTime.Get()
}

// GetSystemLastbootedTimeOk returns a tuple with the SystemLastbootedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetSystemLastbootedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemLastbootedTime.Get(), o.SystemLastbootedTime.IsSet()
}

// HasSystemLastbootedTime returns a boolean if a field has been set.
func (o *System) HasSystemLastbootedTime() bool {
	if o != nil && o.SystemLastbootedTime.IsSet() {
		return true
	}

	return false
}

// SetSystemLastbootedTime gets a reference to the given NullableTime and assigns it to the SystemLastbootedTime field.
func (o *System) SetSystemLastbootedTime(v time.Time) {
	o.SystemLastbootedTime.Set(&v)
}
// SetSystemLastbootedTimeNil sets the value for SystemLastbootedTime to be an explicit nil
func (o *System) SetSystemLastbootedTimeNil() {
	o.SystemLastbootedTime.Set(nil)
}

// UnsetSystemLastbootedTime ensures that no value is present for SystemLastbootedTime, not even an explicit nil
func (o *System) UnsetSystemLastbootedTime() {
	o.SystemLastbootedTime.Unset()
}

// GetSystemDeprecatedTime returns the SystemDeprecatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetSystemDeprecatedTime() time.Time {
	if o == nil || o.SystemDeprecatedTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SystemDeprecatedTime.Get()
}

// GetSystemDeprecatedTimeOk returns a tuple with the SystemDeprecatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetSystemDeprecatedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemDeprecatedTime.Get(), o.SystemDeprecatedTime.IsSet()
}

// HasSystemDeprecatedTime returns a boolean if a field has been set.
func (o *System) HasSystemDeprecatedTime() bool {
	if o != nil && o.SystemDeprecatedTime.IsSet() {
		return true
	}

	return false
}

// SetSystemDeprecatedTime gets a reference to the given NullableTime and assigns it to the SystemDeprecatedTime field.
func (o *System) SetSystemDeprecatedTime(v time.Time) {
	o.SystemDeprecatedTime.Set(&v)
}
// SetSystemDeprecatedTimeNil sets the value for SystemDeprecatedTime to be an explicit nil
func (o *System) SetSystemDeprecatedTimeNil() {
	o.SystemDeprecatedTime.Set(nil)
}

// UnsetSystemDeprecatedTime ensures that no value is present for SystemDeprecatedTime, not even an explicit nil
func (o *System) UnsetSystemDeprecatedTime() {
	o.SystemDeprecatedTime.Unset()
}

// GetSystemIsVm returns the SystemIsVm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetSystemIsVm() bool {
	if o == nil || o.SystemIsVm.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SystemIsVm.Get()
}

// GetSystemIsVmOk returns a tuple with the SystemIsVm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetSystemIsVmOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemIsVm.Get(), o.SystemIsVm.IsSet()
}

// HasSystemIsVm returns a boolean if a field has been set.
func (o *System) HasSystemIsVm() bool {
	if o != nil && o.SystemIsVm.IsSet() {
		return true
	}

	return false
}

// SetSystemIsVm gets a reference to the given NullableBool and assigns it to the SystemIsVm field.
func (o *System) SetSystemIsVm(v bool) {
	o.SystemIsVm.Set(&v)
}
// SetSystemIsVmNil sets the value for SystemIsVm to be an explicit nil
func (o *System) SetSystemIsVmNil() {
	o.SystemIsVm.Set(nil)
}

// UnsetSystemIsVm ensures that no value is present for SystemIsVm, not even an explicit nil
func (o *System) UnsetSystemIsVm() {
	o.SystemIsVm.Unset()
}

// GetHostSystem returns the HostSystem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetHostSystem() int32 {
	if o == nil || o.HostSystem.Get() == nil {
		var ret int32
		return ret
	}
	return *o.HostSystem.Get()
}

// GetHostSystemOk returns a tuple with the HostSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetHostSystemOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HostSystem.Get(), o.HostSystem.IsSet()
}

// HasHostSystem returns a boolean if a field has been set.
func (o *System) HasHostSystem() bool {
	if o != nil && o.HostSystem.IsSet() {
		return true
	}

	return false
}

// SetHostSystem gets a reference to the given NullableInt32 and assigns it to the HostSystem field.
func (o *System) SetHostSystem(v int32) {
	o.HostSystem.Set(&v)
}
// SetHostSystemNil sets the value for HostSystem to be an explicit nil
func (o *System) SetHostSystemNil() {
	o.HostSystem.Set(nil)
}

// UnsetHostSystem ensures that no value is present for HostSystem, not even an explicit nil
func (o *System) UnsetHostSystem() {
	o.HostSystem.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *System) GetCompany() []int32 {
	if o == nil || o.Company == nil {
		var ret []int32
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetCompanyOk() (*[]int32, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *System) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given []int32 and assigns it to the Company field.
func (o *System) SetCompany(v []int32) {
	o.Company = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetLocation() int32 {
	if o == nil || o.Location.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetLocationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *System) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableInt32 and assigns it to the Location field.
func (o *System) SetLocation(v int32) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *System) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *System) UnsetLocation() {
	o.Location.Unset()
}

// GetServiceprovider returns the Serviceprovider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetServiceprovider() int32 {
	if o == nil || o.Serviceprovider.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Serviceprovider.Get()
}

// GetServiceproviderOk returns a tuple with the Serviceprovider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetServiceproviderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Serviceprovider.Get(), o.Serviceprovider.IsSet()
}

// HasServiceprovider returns a boolean if a field has been set.
func (o *System) HasServiceprovider() bool {
	if o != nil && o.Serviceprovider.IsSet() {
		return true
	}

	return false
}

// SetServiceprovider gets a reference to the given NullableInt32 and assigns it to the Serviceprovider field.
func (o *System) SetServiceprovider(v int32) {
	o.Serviceprovider.Set(&v)
}
// SetServiceproviderNil sets the value for Serviceprovider to be an explicit nil
func (o *System) SetServiceproviderNil() {
	o.Serviceprovider.Set(nil)
}

// UnsetServiceprovider ensures that no value is present for Serviceprovider, not even an explicit nil
func (o *System) UnsetServiceprovider() {
	o.Serviceprovider.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetContact() int32 {
	if o == nil || o.Contact.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetContactOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *System) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableInt32 and assigns it to the Contact field.
func (o *System) SetContact(v int32) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *System) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *System) UnsetContact() {
	o.Contact.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *System) GetTag() []int32 {
	if o == nil || o.Tag == nil {
		var ret []int32
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetTagOk() (*[]int32, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *System) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []int32 and assigns it to the Tag field.
func (o *System) SetTag(v []int32) {
	o.Tag = &v
}

// GetCase returns the Case field value if set, zero value otherwise.
func (o *System) GetCase() []int32 {
	if o == nil || o.Case == nil {
		var ret []int32
		return ret
	}
	return *o.Case
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetCaseOk() (*[]int32, bool) {
	if o == nil || o.Case == nil {
		return nil, false
	}
	return o.Case, true
}

// HasCase returns a boolean if a field has been set.
func (o *System) HasCase() bool {
	if o != nil && o.Case != nil {
		return true
	}

	return false
}

// SetCase gets a reference to the given []int32 and assigns it to the Case field.
func (o *System) SetCase(v []int32) {
	o.Case = &v
}

// GetSystemApiTime returns the SystemApiTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetSystemApiTime() time.Time {
	if o == nil || o.SystemApiTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SystemApiTime.Get()
}

// GetSystemApiTimeOk returns a tuple with the SystemApiTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetSystemApiTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SystemApiTime.Get(), o.SystemApiTime.IsSet()
}

// HasSystemApiTime returns a boolean if a field has been set.
func (o *System) HasSystemApiTime() bool {
	if o != nil && o.SystemApiTime.IsSet() {
		return true
	}

	return false
}

// SetSystemApiTime gets a reference to the given NullableTime and assigns it to the SystemApiTime field.
func (o *System) SetSystemApiTime(v time.Time) {
	o.SystemApiTime.Set(&v)
}
// SetSystemApiTimeNil sets the value for SystemApiTime to be an explicit nil
func (o *System) SetSystemApiTimeNil() {
	o.SystemApiTime.Set(nil)
}

// UnsetSystemApiTime ensures that no value is present for SystemApiTime, not even an explicit nil
func (o *System) UnsetSystemApiTime() {
	o.SystemApiTime.Unset()
}

// GetSystemCreateTime returns the SystemCreateTime field value if set, zero value otherwise.
func (o *System) GetSystemCreateTime() time.Time {
	if o == nil || o.SystemCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.SystemCreateTime
}

// GetSystemCreateTimeOk returns a tuple with the SystemCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.SystemCreateTime == nil {
		return nil, false
	}
	return o.SystemCreateTime, true
}

// HasSystemCreateTime returns a boolean if a field has been set.
func (o *System) HasSystemCreateTime() bool {
	if o != nil && o.SystemCreateTime != nil {
		return true
	}

	return false
}

// SetSystemCreateTime gets a reference to the given time.Time and assigns it to the SystemCreateTime field.
func (o *System) SetSystemCreateTime(v time.Time) {
	o.SystemCreateTime = &v
}

// GetSystemCreatedByUserId returns the SystemCreatedByUserId field value
func (o *System) GetSystemCreatedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.SystemCreatedByUserId
}

// GetSystemCreatedByUserIdOk returns a tuple with the SystemCreatedByUserId field value
// and a boolean to check if the value has been set.
func (o *System) GetSystemCreatedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemCreatedByUserId, true
}

// SetSystemCreatedByUserId sets field value
func (o *System) SetSystemCreatedByUserId(v int32) {
	o.SystemCreatedByUserId = v
}

// GetSystemModifyTime returns the SystemModifyTime field value
func (o *System) GetSystemModifyTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.SystemModifyTime
}

// GetSystemModifyTimeOk returns a tuple with the SystemModifyTime field value
// and a boolean to check if the value has been set.
func (o *System) GetSystemModifyTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemModifyTime, true
}

// SetSystemModifyTime sets field value
func (o *System) SetSystemModifyTime(v time.Time) {
	o.SystemModifyTime = v
}

// GetSystemModifiedByUserId returns the SystemModifiedByUserId field value
func (o *System) GetSystemModifiedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.SystemModifiedByUserId
}

// GetSystemModifiedByUserIdOk returns a tuple with the SystemModifiedByUserId field value
// and a boolean to check if the value has been set.
func (o *System) GetSystemModifiedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemModifiedByUserId, true
}

// SetSystemModifiedByUserId sets field value
func (o *System) SetSystemModifiedByUserId(v int32) {
	o.SystemModifiedByUserId = v
}

// GetSystemExportMarkdown returns the SystemExportMarkdown field value if set, zero value otherwise.
func (o *System) GetSystemExportMarkdown() bool {
	if o == nil || o.SystemExportMarkdown == nil {
		var ret bool
		return ret
	}
	return *o.SystemExportMarkdown
}

// GetSystemExportMarkdownOk returns a tuple with the SystemExportMarkdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemExportMarkdownOk() (*bool, bool) {
	if o == nil || o.SystemExportMarkdown == nil {
		return nil, false
	}
	return o.SystemExportMarkdown, true
}

// HasSystemExportMarkdown returns a boolean if a field has been set.
func (o *System) HasSystemExportMarkdown() bool {
	if o != nil && o.SystemExportMarkdown != nil {
		return true
	}

	return false
}

// SetSystemExportMarkdown gets a reference to the given bool and assigns it to the SystemExportMarkdown field.
func (o *System) SetSystemExportMarkdown(v bool) {
	o.SystemExportMarkdown = &v
}

// GetSystemExportSpreadsheet returns the SystemExportSpreadsheet field value if set, zero value otherwise.
func (o *System) GetSystemExportSpreadsheet() bool {
	if o == nil || o.SystemExportSpreadsheet == nil {
		var ret bool
		return ret
	}
	return *o.SystemExportSpreadsheet
}

// GetSystemExportSpreadsheetOk returns a tuple with the SystemExportSpreadsheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemExportSpreadsheetOk() (*bool, bool) {
	if o == nil || o.SystemExportSpreadsheet == nil {
		return nil, false
	}
	return o.SystemExportSpreadsheet, true
}

// HasSystemExportSpreadsheet returns a boolean if a field has been set.
func (o *System) HasSystemExportSpreadsheet() bool {
	if o != nil && o.SystemExportSpreadsheet != nil {
		return true
	}

	return false
}

// SetSystemExportSpreadsheet gets a reference to the given bool and assigns it to the SystemExportSpreadsheet field.
func (o *System) SetSystemExportSpreadsheet(v bool) {
	o.SystemExportSpreadsheet = &v
}

func (o System) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemId != nil {
		toSerialize["system_id"] = o.SystemId
	}
	if o.SystemUuid != nil {
		toSerialize["system_uuid"] = o.SystemUuid
	}
	if true {
		toSerialize["system_name"] = o.SystemName
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Dnsname.IsSet() {
		toSerialize["dnsname"] = o.Dnsname.Get()
	}
	if true {
		toSerialize["systemstatus"] = o.Systemstatus
	}
	if o.Analysisstatus.IsSet() {
		toSerialize["analysisstatus"] = o.Analysisstatus.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Recommendation.IsSet() {
		toSerialize["recommendation"] = o.Recommendation.Get()
	}
	if o.Systemtype.IsSet() {
		toSerialize["systemtype"] = o.Systemtype.Get()
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Os.IsSet() {
		toSerialize["os"] = o.Os.Get()
	}
	if o.Osarch.IsSet() {
		toSerialize["osarch"] = o.Osarch.Get()
	}
	if o.SystemLastbootedTime.IsSet() {
		toSerialize["system_lastbooted_time"] = o.SystemLastbootedTime.Get()
	}
	if o.SystemDeprecatedTime.IsSet() {
		toSerialize["system_deprecated_time"] = o.SystemDeprecatedTime.Get()
	}
	if o.SystemIsVm.IsSet() {
		toSerialize["system_is_vm"] = o.SystemIsVm.Get()
	}
	if o.HostSystem.IsSet() {
		toSerialize["host_system"] = o.HostSystem.Get()
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Serviceprovider.IsSet() {
		toSerialize["serviceprovider"] = o.Serviceprovider.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Case != nil {
		toSerialize["case"] = o.Case
	}
	if o.SystemApiTime.IsSet() {
		toSerialize["system_api_time"] = o.SystemApiTime.Get()
	}
	if o.SystemCreateTime != nil {
		toSerialize["system_create_time"] = o.SystemCreateTime
	}
	if true {
		toSerialize["system_created_by_user_id"] = o.SystemCreatedByUserId
	}
	if true {
		toSerialize["system_modify_time"] = o.SystemModifyTime
	}
	if true {
		toSerialize["system_modified_by_user_id"] = o.SystemModifiedByUserId
	}
	if o.SystemExportMarkdown != nil {
		toSerialize["system_export_markdown"] = o.SystemExportMarkdown
	}
	if o.SystemExportSpreadsheet != nil {
		toSerialize["system_export_spreadsheet"] = o.SystemExportSpreadsheet
	}
	return json.Marshal(toSerialize)
}

type NullableSystem struct {
	value *System
	isSet bool
}

func (v NullableSystem) Get() *System {
	return v.value
}

func (v *NullableSystem) Set(val *System) {
	v.value = val
	v.isSet = true
}

func (v NullableSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystem(val *System) *NullableSystem {
	return &NullableSystem{value: val, isSet: true}
}

func (v NullableSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


