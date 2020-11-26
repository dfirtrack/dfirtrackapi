# Reason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonId** | Pointer to **int32** |  | [optional] [readonly] 
**ReasonName** | **string** |  | 

## Methods

### NewReason

`func NewReason(reasonName string, ) *Reason`

NewReason instantiates a new Reason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonWithDefaults

`func NewReasonWithDefaults() *Reason`

NewReasonWithDefaults instantiates a new Reason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonId

`func (o *Reason) GetReasonId() int32`

GetReasonId returns the ReasonId field if non-nil, zero value otherwise.

### GetReasonIdOk

`func (o *Reason) GetReasonIdOk() (*int32, bool)`

GetReasonIdOk returns a tuple with the ReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonId

`func (o *Reason) SetReasonId(v int32)`

SetReasonId sets ReasonId field to given value.

### HasReasonId

`func (o *Reason) HasReasonId() bool`

HasReasonId returns a boolean if a field has been set.

### GetReasonName

`func (o *Reason) GetReasonName() string`

GetReasonName returns the ReasonName field if non-nil, zero value otherwise.

### GetReasonNameOk

`func (o *Reason) GetReasonNameOk() (*string, bool)`

GetReasonNameOk returns a tuple with the ReasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonName

`func (o *Reason) SetReasonName(v string)`

SetReasonName sets ReasonName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


