# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **int32** |  | [optional] [readonly] 
**CompanyName** | **string** |  | 
**Division** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCompany

`func NewCompany(companyName string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *Company) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Company) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Company) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Company) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *Company) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Company) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Company) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetDivision

`func (o *Company) GetDivision() int32`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *Company) GetDivisionOk() (*int32, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *Company) SetDivision(v int32)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *Company) HasDivision() bool`

HasDivision returns a boolean if a field has been set.

### SetDivisionNil

`func (o *Company) SetDivisionNil(b bool)`

 SetDivisionNil sets the value for Division to be an explicit nil

### UnsetDivision
`func (o *Company) UnsetDivision()`

UnsetDivision ensures that no value is present for Division, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


