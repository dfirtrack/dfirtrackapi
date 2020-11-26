# Systemuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemuserId** | Pointer to **int32** |  | [optional] [readonly] 
**SystemuserName** | **string** |  | 
**System** | **int32** |  | 
**SystemuserLastlogonTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**SystemuserIsSystemadmin** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSystemuser

`func NewSystemuser(systemuserName string, system int32, ) *Systemuser`

NewSystemuser instantiates a new Systemuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemuserWithDefaults

`func NewSystemuserWithDefaults() *Systemuser`

NewSystemuserWithDefaults instantiates a new Systemuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemuserId

`func (o *Systemuser) GetSystemuserId() int32`

GetSystemuserId returns the SystemuserId field if non-nil, zero value otherwise.

### GetSystemuserIdOk

`func (o *Systemuser) GetSystemuserIdOk() (*int32, bool)`

GetSystemuserIdOk returns a tuple with the SystemuserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemuserId

`func (o *Systemuser) SetSystemuserId(v int32)`

SetSystemuserId sets SystemuserId field to given value.

### HasSystemuserId

`func (o *Systemuser) HasSystemuserId() bool`

HasSystemuserId returns a boolean if a field has been set.

### GetSystemuserName

`func (o *Systemuser) GetSystemuserName() string`

GetSystemuserName returns the SystemuserName field if non-nil, zero value otherwise.

### GetSystemuserNameOk

`func (o *Systemuser) GetSystemuserNameOk() (*string, bool)`

GetSystemuserNameOk returns a tuple with the SystemuserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemuserName

`func (o *Systemuser) SetSystemuserName(v string)`

SetSystemuserName sets SystemuserName field to given value.


### GetSystem

`func (o *Systemuser) GetSystem() int32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Systemuser) GetSystemOk() (*int32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Systemuser) SetSystem(v int32)`

SetSystem sets System field to given value.


### GetSystemuserLastlogonTime

`func (o *Systemuser) GetSystemuserLastlogonTime() time.Time`

GetSystemuserLastlogonTime returns the SystemuserLastlogonTime field if non-nil, zero value otherwise.

### GetSystemuserLastlogonTimeOk

`func (o *Systemuser) GetSystemuserLastlogonTimeOk() (*time.Time, bool)`

GetSystemuserLastlogonTimeOk returns a tuple with the SystemuserLastlogonTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemuserLastlogonTime

`func (o *Systemuser) SetSystemuserLastlogonTime(v time.Time)`

SetSystemuserLastlogonTime sets SystemuserLastlogonTime field to given value.

### HasSystemuserLastlogonTime

`func (o *Systemuser) HasSystemuserLastlogonTime() bool`

HasSystemuserLastlogonTime returns a boolean if a field has been set.

### SetSystemuserLastlogonTimeNil

`func (o *Systemuser) SetSystemuserLastlogonTimeNil(b bool)`

 SetSystemuserLastlogonTimeNil sets the value for SystemuserLastlogonTime to be an explicit nil

### UnsetSystemuserLastlogonTime
`func (o *Systemuser) UnsetSystemuserLastlogonTime()`

UnsetSystemuserLastlogonTime ensures that no value is present for SystemuserLastlogonTime, not even an explicit nil
### GetSystemuserIsSystemadmin

`func (o *Systemuser) GetSystemuserIsSystemadmin() bool`

GetSystemuserIsSystemadmin returns the SystemuserIsSystemadmin field if non-nil, zero value otherwise.

### GetSystemuserIsSystemadminOk

`func (o *Systemuser) GetSystemuserIsSystemadminOk() (*bool, bool)`

GetSystemuserIsSystemadminOk returns a tuple with the SystemuserIsSystemadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemuserIsSystemadmin

`func (o *Systemuser) SetSystemuserIsSystemadmin(v bool)`

SetSystemuserIsSystemadmin sets SystemuserIsSystemadmin field to given value.

### HasSystemuserIsSystemadmin

`func (o *Systemuser) HasSystemuserIsSystemadmin() bool`

HasSystemuserIsSystemadmin returns a boolean if a field has been set.

### SetSystemuserIsSystemadminNil

`func (o *Systemuser) SetSystemuserIsSystemadminNil(b bool)`

 SetSystemuserIsSystemadminNil sets the value for SystemuserIsSystemadmin to be an explicit nil

### UnsetSystemuserIsSystemadmin
`func (o *Systemuser) UnsetSystemuserIsSystemadmin()`

UnsetSystemuserIsSystemadmin ensures that no value is present for SystemuserIsSystemadmin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


