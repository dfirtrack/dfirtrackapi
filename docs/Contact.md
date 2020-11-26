# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | Pointer to **int32** |  | [optional] [readonly] 
**ContactName** | **string** |  | 
**ContactEmail** | **string** |  | 
**ContactPhone** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContact

`func NewContact(contactName string, contactEmail string, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *Contact) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *Contact) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *Contact) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *Contact) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetContactName

`func (o *Contact) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Contact) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Contact) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetContactEmail

`func (o *Contact) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Contact) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Contact) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetContactPhone

`func (o *Contact) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *Contact) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *Contact) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.

### HasContactPhone

`func (o *Contact) HasContactPhone() bool`

HasContactPhone returns a boolean if a field has been set.

### SetContactPhoneNil

`func (o *Contact) SetContactPhoneNil(b bool)`

 SetContactPhoneNil sets the value for ContactPhone to be an explicit nil

### UnsetContactPhone
`func (o *Contact) UnsetContactPhone()`

UnsetContactPhone ensures that no value is present for ContactPhone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


