# Recommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecommendationId** | Pointer to **int32** |  | [optional] [readonly] 
**RecommendationName** | **string** |  | 

## Methods

### NewRecommendation

`func NewRecommendation(recommendationName string, ) *Recommendation`

NewRecommendation instantiates a new Recommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationWithDefaults

`func NewRecommendationWithDefaults() *Recommendation`

NewRecommendationWithDefaults instantiates a new Recommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendationId

`func (o *Recommendation) GetRecommendationId() int32`

GetRecommendationId returns the RecommendationId field if non-nil, zero value otherwise.

### GetRecommendationIdOk

`func (o *Recommendation) GetRecommendationIdOk() (*int32, bool)`

GetRecommendationIdOk returns a tuple with the RecommendationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationId

`func (o *Recommendation) SetRecommendationId(v int32)`

SetRecommendationId sets RecommendationId field to given value.

### HasRecommendationId

`func (o *Recommendation) HasRecommendationId() bool`

HasRecommendationId returns a boolean if a field has been set.

### GetRecommendationName

`func (o *Recommendation) GetRecommendationName() string`

GetRecommendationName returns the RecommendationName field if non-nil, zero value otherwise.

### GetRecommendationNameOk

`func (o *Recommendation) GetRecommendationNameOk() (*string, bool)`

GetRecommendationNameOk returns a tuple with the RecommendationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationName

`func (o *Recommendation) SetRecommendationName(v string)`

SetRecommendationName sets RecommendationName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


