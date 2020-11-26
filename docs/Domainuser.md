# Domainuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainuserId** | Pointer to **int32** |  | [optional] [readonly] 
**DomainuserName** | **string** |  | 
**Domain** | **int32** |  | 
**DomainuserIsDomainadmin** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDomainuser

`func NewDomainuser(domainuserName string, domain int32, ) *Domainuser`

NewDomainuser instantiates a new Domainuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainuserWithDefaults

`func NewDomainuserWithDefaults() *Domainuser`

NewDomainuserWithDefaults instantiates a new Domainuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainuserId

`func (o *Domainuser) GetDomainuserId() int32`

GetDomainuserId returns the DomainuserId field if non-nil, zero value otherwise.

### GetDomainuserIdOk

`func (o *Domainuser) GetDomainuserIdOk() (*int32, bool)`

GetDomainuserIdOk returns a tuple with the DomainuserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainuserId

`func (o *Domainuser) SetDomainuserId(v int32)`

SetDomainuserId sets DomainuserId field to given value.

### HasDomainuserId

`func (o *Domainuser) HasDomainuserId() bool`

HasDomainuserId returns a boolean if a field has been set.

### GetDomainuserName

`func (o *Domainuser) GetDomainuserName() string`

GetDomainuserName returns the DomainuserName field if non-nil, zero value otherwise.

### GetDomainuserNameOk

`func (o *Domainuser) GetDomainuserNameOk() (*string, bool)`

GetDomainuserNameOk returns a tuple with the DomainuserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainuserName

`func (o *Domainuser) SetDomainuserName(v string)`

SetDomainuserName sets DomainuserName field to given value.


### GetDomain

`func (o *Domainuser) GetDomain() int32`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domainuser) GetDomainOk() (*int32, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domainuser) SetDomain(v int32)`

SetDomain sets Domain field to given value.


### GetDomainuserIsDomainadmin

`func (o *Domainuser) GetDomainuserIsDomainadmin() bool`

GetDomainuserIsDomainadmin returns the DomainuserIsDomainadmin field if non-nil, zero value otherwise.

### GetDomainuserIsDomainadminOk

`func (o *Domainuser) GetDomainuserIsDomainadminOk() (*bool, bool)`

GetDomainuserIsDomainadminOk returns a tuple with the DomainuserIsDomainadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainuserIsDomainadmin

`func (o *Domainuser) SetDomainuserIsDomainadmin(v bool)`

SetDomainuserIsDomainadmin sets DomainuserIsDomainadmin field to given value.

### HasDomainuserIsDomainadmin

`func (o *Domainuser) HasDomainuserIsDomainadmin() bool`

HasDomainuserIsDomainadmin returns a boolean if a field has been set.

### SetDomainuserIsDomainadminNil

`func (o *Domainuser) SetDomainuserIsDomainadminNil(b bool)`

 SetDomainuserIsDomainadminNil sets the value for DomainuserIsDomainadmin to be an explicit nil

### UnsetDomainuserIsDomainadmin
`func (o *Domainuser) UnsetDomainuserIsDomainadmin()`

UnsetDomainuserIsDomainadmin ensures that no value is present for DomainuserIsDomainadmin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


