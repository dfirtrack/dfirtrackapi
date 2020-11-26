# System

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemId** | Pointer to **int32** |  | [optional] [readonly] 
**SystemUuid** | Pointer to **string** |  | [optional] [readonly] 
**SystemName** | **string** |  | 
**Domain** | Pointer to **NullableInt32** |  | [optional] 
**Dnsname** | Pointer to **NullableInt32** |  | [optional] 
**Systemstatus** | **int32** |  | 
**Analysisstatus** | Pointer to **NullableInt32** |  | [optional] 
**Reason** | Pointer to **NullableInt32** |  | [optional] 
**Recommendation** | Pointer to **NullableInt32** |  | [optional] 
**Systemtype** | Pointer to **NullableInt32** |  | [optional] 
**Ip** | Pointer to **[]int32** |  | [optional] 
**Os** | Pointer to **NullableInt32** |  | [optional] 
**Osarch** | Pointer to **NullableInt32** |  | [optional] 
**SystemLastbootedTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**SystemDeprecatedTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**SystemIsVm** | Pointer to **NullableBool** |  | [optional] 
**HostSystem** | Pointer to **NullableInt32** |  | [optional] 
**Company** | Pointer to **[]int32** |  | [optional] 
**Location** | Pointer to **NullableInt32** |  | [optional] 
**Serviceprovider** | Pointer to **NullableInt32** |  | [optional] 
**Contact** | Pointer to **NullableInt32** |  | [optional] 
**Tag** | Pointer to **[]int32** |  | [optional] 
**Case** | Pointer to **[]int32** |  | [optional] 
**SystemApiTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**SystemCreateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**SystemCreatedByUserId** | **int32** |  | 
**SystemModifyTime** | [**time.Time**](time.Time.md) |  | 
**SystemModifiedByUserId** | **int32** |  | 
**SystemExportMarkdown** | Pointer to **bool** |  | [optional] 
**SystemExportSpreadsheet** | Pointer to **bool** |  | [optional] 

## Methods

### NewSystem

`func NewSystem(systemName string, systemstatus int32, systemCreatedByUserId int32, systemModifyTime time.Time, systemModifiedByUserId int32, ) *System`

NewSystem instantiates a new System object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWithDefaults

`func NewSystemWithDefaults() *System`

NewSystemWithDefaults instantiates a new System object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemId

`func (o *System) GetSystemId() int32`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *System) GetSystemIdOk() (*int32, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *System) SetSystemId(v int32)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *System) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetSystemUuid

`func (o *System) GetSystemUuid() string`

GetSystemUuid returns the SystemUuid field if non-nil, zero value otherwise.

### GetSystemUuidOk

`func (o *System) GetSystemUuidOk() (*string, bool)`

GetSystemUuidOk returns a tuple with the SystemUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUuid

`func (o *System) SetSystemUuid(v string)`

SetSystemUuid sets SystemUuid field to given value.

### HasSystemUuid

`func (o *System) HasSystemUuid() bool`

HasSystemUuid returns a boolean if a field has been set.

### GetSystemName

`func (o *System) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *System) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *System) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.


### GetDomain

`func (o *System) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *System) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *System) SetDomain(v int32)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *System) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *System) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *System) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDnsname

`func (o *System) GetDnsname() int32`

GetDnsname returns the Dnsname field if non-nil, zero value otherwise.

### GetDnsnameOk

`func (o *System) GetDnsnameOk() (*int32, bool)`

GetDnsnameOk returns a tuple with the Dnsname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsname

`func (o *System) SetDnsname(v int32)`

SetDnsname sets Dnsname field to given value.

### HasDnsname

`func (o *System) HasDnsname() bool`

HasDnsname returns a boolean if a field has been set.

### SetDnsnameNil

`func (o *System) SetDnsnameNil(b bool)`

 SetDnsnameNil sets the value for Dnsname to be an explicit nil

### UnsetDnsname
`func (o *System) UnsetDnsname()`

UnsetDnsname ensures that no value is present for Dnsname, not even an explicit nil
### GetSystemstatus

`func (o *System) GetSystemstatus() int32`

GetSystemstatus returns the Systemstatus field if non-nil, zero value otherwise.

### GetSystemstatusOk

`func (o *System) GetSystemstatusOk() (*int32, bool)`

GetSystemstatusOk returns a tuple with the Systemstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemstatus

`func (o *System) SetSystemstatus(v int32)`

SetSystemstatus sets Systemstatus field to given value.


### GetAnalysisstatus

`func (o *System) GetAnalysisstatus() int32`

GetAnalysisstatus returns the Analysisstatus field if non-nil, zero value otherwise.

### GetAnalysisstatusOk

`func (o *System) GetAnalysisstatusOk() (*int32, bool)`

GetAnalysisstatusOk returns a tuple with the Analysisstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisstatus

`func (o *System) SetAnalysisstatus(v int32)`

SetAnalysisstatus sets Analysisstatus field to given value.

### HasAnalysisstatus

`func (o *System) HasAnalysisstatus() bool`

HasAnalysisstatus returns a boolean if a field has been set.

### SetAnalysisstatusNil

`func (o *System) SetAnalysisstatusNil(b bool)`

 SetAnalysisstatusNil sets the value for Analysisstatus to be an explicit nil

### UnsetAnalysisstatus
`func (o *System) UnsetAnalysisstatus()`

UnsetAnalysisstatus ensures that no value is present for Analysisstatus, not even an explicit nil
### GetReason

`func (o *System) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *System) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *System) SetReason(v int32)`

SetReason sets Reason field to given value.

### HasReason

`func (o *System) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *System) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *System) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRecommendation

`func (o *System) GetRecommendation() int32`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *System) GetRecommendationOk() (*int32, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *System) SetRecommendation(v int32)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *System) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *System) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *System) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetSystemtype

`func (o *System) GetSystemtype() int32`

GetSystemtype returns the Systemtype field if non-nil, zero value otherwise.

### GetSystemtypeOk

`func (o *System) GetSystemtypeOk() (*int32, bool)`

GetSystemtypeOk returns a tuple with the Systemtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemtype

`func (o *System) SetSystemtype(v int32)`

SetSystemtype sets Systemtype field to given value.

### HasSystemtype

`func (o *System) HasSystemtype() bool`

HasSystemtype returns a boolean if a field has been set.

### SetSystemtypeNil

`func (o *System) SetSystemtypeNil(b bool)`

 SetSystemtypeNil sets the value for Systemtype to be an explicit nil

### UnsetSystemtype
`func (o *System) UnsetSystemtype()`

UnsetSystemtype ensures that no value is present for Systemtype, not even an explicit nil
### GetIp

`func (o *System) GetIp() []int32`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *System) GetIpOk() (*[]int32, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *System) SetIp(v []int32)`

SetIp sets Ip field to given value.

### HasIp

`func (o *System) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetOs

`func (o *System) GetOs() int32`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *System) GetOsOk() (*int32, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *System) SetOs(v int32)`

SetOs sets Os field to given value.

### HasOs

`func (o *System) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *System) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *System) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetOsarch

`func (o *System) GetOsarch() int32`

GetOsarch returns the Osarch field if non-nil, zero value otherwise.

### GetOsarchOk

`func (o *System) GetOsarchOk() (*int32, bool)`

GetOsarchOk returns a tuple with the Osarch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsarch

`func (o *System) SetOsarch(v int32)`

SetOsarch sets Osarch field to given value.

### HasOsarch

`func (o *System) HasOsarch() bool`

HasOsarch returns a boolean if a field has been set.

### SetOsarchNil

`func (o *System) SetOsarchNil(b bool)`

 SetOsarchNil sets the value for Osarch to be an explicit nil

### UnsetOsarch
`func (o *System) UnsetOsarch()`

UnsetOsarch ensures that no value is present for Osarch, not even an explicit nil
### GetSystemLastbootedTime

`func (o *System) GetSystemLastbootedTime() time.Time`

GetSystemLastbootedTime returns the SystemLastbootedTime field if non-nil, zero value otherwise.

### GetSystemLastbootedTimeOk

`func (o *System) GetSystemLastbootedTimeOk() (*time.Time, bool)`

GetSystemLastbootedTimeOk returns a tuple with the SystemLastbootedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLastbootedTime

`func (o *System) SetSystemLastbootedTime(v time.Time)`

SetSystemLastbootedTime sets SystemLastbootedTime field to given value.

### HasSystemLastbootedTime

`func (o *System) HasSystemLastbootedTime() bool`

HasSystemLastbootedTime returns a boolean if a field has been set.

### SetSystemLastbootedTimeNil

`func (o *System) SetSystemLastbootedTimeNil(b bool)`

 SetSystemLastbootedTimeNil sets the value for SystemLastbootedTime to be an explicit nil

### UnsetSystemLastbootedTime
`func (o *System) UnsetSystemLastbootedTime()`

UnsetSystemLastbootedTime ensures that no value is present for SystemLastbootedTime, not even an explicit nil
### GetSystemDeprecatedTime

`func (o *System) GetSystemDeprecatedTime() time.Time`

GetSystemDeprecatedTime returns the SystemDeprecatedTime field if non-nil, zero value otherwise.

### GetSystemDeprecatedTimeOk

`func (o *System) GetSystemDeprecatedTimeOk() (*time.Time, bool)`

GetSystemDeprecatedTimeOk returns a tuple with the SystemDeprecatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDeprecatedTime

`func (o *System) SetSystemDeprecatedTime(v time.Time)`

SetSystemDeprecatedTime sets SystemDeprecatedTime field to given value.

### HasSystemDeprecatedTime

`func (o *System) HasSystemDeprecatedTime() bool`

HasSystemDeprecatedTime returns a boolean if a field has been set.

### SetSystemDeprecatedTimeNil

`func (o *System) SetSystemDeprecatedTimeNil(b bool)`

 SetSystemDeprecatedTimeNil sets the value for SystemDeprecatedTime to be an explicit nil

### UnsetSystemDeprecatedTime
`func (o *System) UnsetSystemDeprecatedTime()`

UnsetSystemDeprecatedTime ensures that no value is present for SystemDeprecatedTime, not even an explicit nil
### GetSystemIsVm

`func (o *System) GetSystemIsVm() bool`

GetSystemIsVm returns the SystemIsVm field if non-nil, zero value otherwise.

### GetSystemIsVmOk

`func (o *System) GetSystemIsVmOk() (*bool, bool)`

GetSystemIsVmOk returns a tuple with the SystemIsVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIsVm

`func (o *System) SetSystemIsVm(v bool)`

SetSystemIsVm sets SystemIsVm field to given value.

### HasSystemIsVm

`func (o *System) HasSystemIsVm() bool`

HasSystemIsVm returns a boolean if a field has been set.

### SetSystemIsVmNil

`func (o *System) SetSystemIsVmNil(b bool)`

 SetSystemIsVmNil sets the value for SystemIsVm to be an explicit nil

### UnsetSystemIsVm
`func (o *System) UnsetSystemIsVm()`

UnsetSystemIsVm ensures that no value is present for SystemIsVm, not even an explicit nil
### GetHostSystem

`func (o *System) GetHostSystem() int32`

GetHostSystem returns the HostSystem field if non-nil, zero value otherwise.

### GetHostSystemOk

`func (o *System) GetHostSystemOk() (*int32, bool)`

GetHostSystemOk returns a tuple with the HostSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSystem

`func (o *System) SetHostSystem(v int32)`

SetHostSystem sets HostSystem field to given value.

### HasHostSystem

`func (o *System) HasHostSystem() bool`

HasHostSystem returns a boolean if a field has been set.

### SetHostSystemNil

`func (o *System) SetHostSystemNil(b bool)`

 SetHostSystemNil sets the value for HostSystem to be an explicit nil

### UnsetHostSystem
`func (o *System) UnsetHostSystem()`

UnsetHostSystem ensures that no value is present for HostSystem, not even an explicit nil
### GetCompany

`func (o *System) GetCompany() []int32`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *System) GetCompanyOk() (*[]int32, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *System) SetCompany(v []int32)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *System) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *System) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *System) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *System) SetLocation(v int32)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *System) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *System) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *System) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetServiceprovider

`func (o *System) GetServiceprovider() int32`

GetServiceprovider returns the Serviceprovider field if non-nil, zero value otherwise.

### GetServiceproviderOk

`func (o *System) GetServiceproviderOk() (*int32, bool)`

GetServiceproviderOk returns a tuple with the Serviceprovider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceprovider

`func (o *System) SetServiceprovider(v int32)`

SetServiceprovider sets Serviceprovider field to given value.

### HasServiceprovider

`func (o *System) HasServiceprovider() bool`

HasServiceprovider returns a boolean if a field has been set.

### SetServiceproviderNil

`func (o *System) SetServiceproviderNil(b bool)`

 SetServiceproviderNil sets the value for Serviceprovider to be an explicit nil

### UnsetServiceprovider
`func (o *System) UnsetServiceprovider()`

UnsetServiceprovider ensures that no value is present for Serviceprovider, not even an explicit nil
### GetContact

`func (o *System) GetContact() int32`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *System) GetContactOk() (*int32, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *System) SetContact(v int32)`

SetContact sets Contact field to given value.

### HasContact

`func (o *System) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *System) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *System) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTag

`func (o *System) GetTag() []int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *System) GetTagOk() (*[]int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *System) SetTag(v []int32)`

SetTag sets Tag field to given value.

### HasTag

`func (o *System) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCase

`func (o *System) GetCase() []int32`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *System) GetCaseOk() (*[]int32, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *System) SetCase(v []int32)`

SetCase sets Case field to given value.

### HasCase

`func (o *System) HasCase() bool`

HasCase returns a boolean if a field has been set.

### GetSystemApiTime

`func (o *System) GetSystemApiTime() time.Time`

GetSystemApiTime returns the SystemApiTime field if non-nil, zero value otherwise.

### GetSystemApiTimeOk

`func (o *System) GetSystemApiTimeOk() (*time.Time, bool)`

GetSystemApiTimeOk returns a tuple with the SystemApiTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemApiTime

`func (o *System) SetSystemApiTime(v time.Time)`

SetSystemApiTime sets SystemApiTime field to given value.

### HasSystemApiTime

`func (o *System) HasSystemApiTime() bool`

HasSystemApiTime returns a boolean if a field has been set.

### SetSystemApiTimeNil

`func (o *System) SetSystemApiTimeNil(b bool)`

 SetSystemApiTimeNil sets the value for SystemApiTime to be an explicit nil

### UnsetSystemApiTime
`func (o *System) UnsetSystemApiTime()`

UnsetSystemApiTime ensures that no value is present for SystemApiTime, not even an explicit nil
### GetSystemCreateTime

`func (o *System) GetSystemCreateTime() time.Time`

GetSystemCreateTime returns the SystemCreateTime field if non-nil, zero value otherwise.

### GetSystemCreateTimeOk

`func (o *System) GetSystemCreateTimeOk() (*time.Time, bool)`

GetSystemCreateTimeOk returns a tuple with the SystemCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCreateTime

`func (o *System) SetSystemCreateTime(v time.Time)`

SetSystemCreateTime sets SystemCreateTime field to given value.

### HasSystemCreateTime

`func (o *System) HasSystemCreateTime() bool`

HasSystemCreateTime returns a boolean if a field has been set.

### GetSystemCreatedByUserId

`func (o *System) GetSystemCreatedByUserId() int32`

GetSystemCreatedByUserId returns the SystemCreatedByUserId field if non-nil, zero value otherwise.

### GetSystemCreatedByUserIdOk

`func (o *System) GetSystemCreatedByUserIdOk() (*int32, bool)`

GetSystemCreatedByUserIdOk returns a tuple with the SystemCreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCreatedByUserId

`func (o *System) SetSystemCreatedByUserId(v int32)`

SetSystemCreatedByUserId sets SystemCreatedByUserId field to given value.


### GetSystemModifyTime

`func (o *System) GetSystemModifyTime() time.Time`

GetSystemModifyTime returns the SystemModifyTime field if non-nil, zero value otherwise.

### GetSystemModifyTimeOk

`func (o *System) GetSystemModifyTimeOk() (*time.Time, bool)`

GetSystemModifyTimeOk returns a tuple with the SystemModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModifyTime

`func (o *System) SetSystemModifyTime(v time.Time)`

SetSystemModifyTime sets SystemModifyTime field to given value.


### GetSystemModifiedByUserId

`func (o *System) GetSystemModifiedByUserId() int32`

GetSystemModifiedByUserId returns the SystemModifiedByUserId field if non-nil, zero value otherwise.

### GetSystemModifiedByUserIdOk

`func (o *System) GetSystemModifiedByUserIdOk() (*int32, bool)`

GetSystemModifiedByUserIdOk returns a tuple with the SystemModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemModifiedByUserId

`func (o *System) SetSystemModifiedByUserId(v int32)`

SetSystemModifiedByUserId sets SystemModifiedByUserId field to given value.


### GetSystemExportMarkdown

`func (o *System) GetSystemExportMarkdown() bool`

GetSystemExportMarkdown returns the SystemExportMarkdown field if non-nil, zero value otherwise.

### GetSystemExportMarkdownOk

`func (o *System) GetSystemExportMarkdownOk() (*bool, bool)`

GetSystemExportMarkdownOk returns a tuple with the SystemExportMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemExportMarkdown

`func (o *System) SetSystemExportMarkdown(v bool)`

SetSystemExportMarkdown sets SystemExportMarkdown field to given value.

### HasSystemExportMarkdown

`func (o *System) HasSystemExportMarkdown() bool`

HasSystemExportMarkdown returns a boolean if a field has been set.

### GetSystemExportSpreadsheet

`func (o *System) GetSystemExportSpreadsheet() bool`

GetSystemExportSpreadsheet returns the SystemExportSpreadsheet field if non-nil, zero value otherwise.

### GetSystemExportSpreadsheetOk

`func (o *System) GetSystemExportSpreadsheetOk() (*bool, bool)`

GetSystemExportSpreadsheetOk returns a tuple with the SystemExportSpreadsheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemExportSpreadsheet

`func (o *System) SetSystemExportSpreadsheet(v bool)`

SetSystemExportSpreadsheet sets SystemExportSpreadsheet field to given value.

### HasSystemExportSpreadsheet

`func (o *System) HasSystemExportSpreadsheet() bool`

HasSystemExportSpreadsheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


