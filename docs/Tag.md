# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | Pointer to **int32** |  | [optional] [readonly] 
**TagName** | **string** |  | 
**TagNote** | Pointer to **NullableString** |  | [optional] 
**Tagcolor** | **int32** |  | 

## Methods

### NewTag

`func NewTag(tagName string, tagcolor int32, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *Tag) GetTagId() int32`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *Tag) GetTagIdOk() (*int32, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *Tag) SetTagId(v int32)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *Tag) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetTagName

`func (o *Tag) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *Tag) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *Tag) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTagNote

`func (o *Tag) GetTagNote() string`

GetTagNote returns the TagNote field if non-nil, zero value otherwise.

### GetTagNoteOk

`func (o *Tag) GetTagNoteOk() (*string, bool)`

GetTagNoteOk returns a tuple with the TagNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNote

`func (o *Tag) SetTagNote(v string)`

SetTagNote sets TagNote field to given value.

### HasTagNote

`func (o *Tag) HasTagNote() bool`

HasTagNote returns a boolean if a field has been set.

### SetTagNoteNil

`func (o *Tag) SetTagNoteNil(b bool)`

 SetTagNoteNil sets the value for TagNote to be an explicit nil

### UnsetTagNote
`func (o *Tag) UnsetTagNote()`

UnsetTagNote ensures that no value is present for TagNote, not even an explicit nil
### GetTagcolor

`func (o *Tag) GetTagcolor() int32`

GetTagcolor returns the Tagcolor field if non-nil, zero value otherwise.

### GetTagcolorOk

`func (o *Tag) GetTagcolorOk() (*int32, bool)`

GetTagcolorOk returns a tuple with the Tagcolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagcolor

`func (o *Tag) SetTagcolor(v int32)`

SetTagcolor sets Tagcolor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


