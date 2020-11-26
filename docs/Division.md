# Division

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivisionId** | Pointer to **int32** |  | [optional] [readonly] 
**DivisionName** | **string** |  | 

## Methods

### NewDivision

`func NewDivision(divisionName string, ) *Division`

NewDivision instantiates a new Division object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDivisionWithDefaults

`func NewDivisionWithDefaults() *Division`

NewDivisionWithDefaults instantiates a new Division object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisionId

`func (o *Division) GetDivisionId() int32`

GetDivisionId returns the DivisionId field if non-nil, zero value otherwise.

### GetDivisionIdOk

`func (o *Division) GetDivisionIdOk() (*int32, bool)`

GetDivisionIdOk returns a tuple with the DivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionId

`func (o *Division) SetDivisionId(v int32)`

SetDivisionId sets DivisionId field to given value.

### HasDivisionId

`func (o *Division) HasDivisionId() bool`

HasDivisionId returns a boolean if a field has been set.

### GetDivisionName

`func (o *Division) GetDivisionName() string`

GetDivisionName returns the DivisionName field if non-nil, zero value otherwise.

### GetDivisionNameOk

`func (o *Division) GetDivisionNameOk() (*string, bool)`

GetDivisionNameOk returns a tuple with the DivisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionName

`func (o *Division) SetDivisionName(v string)`

SetDivisionName sets DivisionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


