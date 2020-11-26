# Os

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsId** | Pointer to **int32** |  | [optional] [readonly] 
**OsName** | **string** |  | 

## Methods

### NewOs

`func NewOs(osName string, ) *Os`

NewOs instantiates a new Os object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsWithDefaults

`func NewOsWithDefaults() *Os`

NewOsWithDefaults instantiates a new Os object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsId

`func (o *Os) GetOsId() int32`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *Os) GetOsIdOk() (*int32, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *Os) SetOsId(v int32)`

SetOsId sets OsId field to given value.

### HasOsId

`func (o *Os) HasOsId() bool`

HasOsId returns a boolean if a field has been set.

### GetOsName

`func (o *Os) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *Os) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *Os) SetOsName(v string)`

SetOsName sets OsName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


