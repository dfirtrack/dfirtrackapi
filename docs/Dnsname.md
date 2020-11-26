# Dnsname

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsnameId** | Pointer to **int32** |  | [optional] [readonly] 
**DnsnameName** | **string** |  | 
**Domain** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewDnsname

`func NewDnsname(dnsnameName string, ) *Dnsname`

NewDnsname instantiates a new Dnsname object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsnameWithDefaults

`func NewDnsnameWithDefaults() *Dnsname`

NewDnsnameWithDefaults instantiates a new Dnsname object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsnameId

`func (o *Dnsname) GetDnsnameId() int32`

GetDnsnameId returns the DnsnameId field if non-nil, zero value otherwise.

### GetDnsnameIdOk

`func (o *Dnsname) GetDnsnameIdOk() (*int32, bool)`

GetDnsnameIdOk returns a tuple with the DnsnameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsnameId

`func (o *Dnsname) SetDnsnameId(v int32)`

SetDnsnameId sets DnsnameId field to given value.

### HasDnsnameId

`func (o *Dnsname) HasDnsnameId() bool`

HasDnsnameId returns a boolean if a field has been set.

### GetDnsnameName

`func (o *Dnsname) GetDnsnameName() string`

GetDnsnameName returns the DnsnameName field if non-nil, zero value otherwise.

### GetDnsnameNameOk

`func (o *Dnsname) GetDnsnameNameOk() (*string, bool)`

GetDnsnameNameOk returns a tuple with the DnsnameName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsnameName

`func (o *Dnsname) SetDnsnameName(v string)`

SetDnsnameName sets DnsnameName field to given value.


### GetDomain

`func (o *Dnsname) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Dnsname) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Dnsname) SetDomain(v int32)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Dnsname) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *Dnsname) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *Dnsname) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


