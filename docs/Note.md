# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoteId** | Pointer to **int32** |  | [optional] [readonly] 
**NoteTitle** | **string** |  | 
**NoteContent** | **string** |  | 
**NoteVersion** | **int32** |  | 
**Case** | Pointer to **NullableInt32** |  | [optional] 
**Notestatus** | Pointer to **int32** |  | [optional] 
**Tag** | Pointer to **[]int32** |  | [optional] 
**NoteCreateTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**NoteCreatedByUserId** | **int32** |  | 
**NoteModifyTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**NoteModifiedByUserId** | **int32** |  | 

## Methods

### NewNote

`func NewNote(noteTitle string, noteContent string, noteVersion int32, noteCreatedByUserId int32, noteModifiedByUserId int32, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoteId

`func (o *Note) GetNoteId() int32`

GetNoteId returns the NoteId field if non-nil, zero value otherwise.

### GetNoteIdOk

`func (o *Note) GetNoteIdOk() (*int32, bool)`

GetNoteIdOk returns a tuple with the NoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteId

`func (o *Note) SetNoteId(v int32)`

SetNoteId sets NoteId field to given value.

### HasNoteId

`func (o *Note) HasNoteId() bool`

HasNoteId returns a boolean if a field has been set.

### GetNoteTitle

`func (o *Note) GetNoteTitle() string`

GetNoteTitle returns the NoteTitle field if non-nil, zero value otherwise.

### GetNoteTitleOk

`func (o *Note) GetNoteTitleOk() (*string, bool)`

GetNoteTitleOk returns a tuple with the NoteTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteTitle

`func (o *Note) SetNoteTitle(v string)`

SetNoteTitle sets NoteTitle field to given value.


### GetNoteContent

`func (o *Note) GetNoteContent() string`

GetNoteContent returns the NoteContent field if non-nil, zero value otherwise.

### GetNoteContentOk

`func (o *Note) GetNoteContentOk() (*string, bool)`

GetNoteContentOk returns a tuple with the NoteContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteContent

`func (o *Note) SetNoteContent(v string)`

SetNoteContent sets NoteContent field to given value.


### GetNoteVersion

`func (o *Note) GetNoteVersion() int32`

GetNoteVersion returns the NoteVersion field if non-nil, zero value otherwise.

### GetNoteVersionOk

`func (o *Note) GetNoteVersionOk() (*int32, bool)`

GetNoteVersionOk returns a tuple with the NoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteVersion

`func (o *Note) SetNoteVersion(v int32)`

SetNoteVersion sets NoteVersion field to given value.


### GetCase

`func (o *Note) GetCase() int32`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *Note) GetCaseOk() (*int32, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *Note) SetCase(v int32)`

SetCase sets Case field to given value.

### HasCase

`func (o *Note) HasCase() bool`

HasCase returns a boolean if a field has been set.

### SetCaseNil

`func (o *Note) SetCaseNil(b bool)`

 SetCaseNil sets the value for Case to be an explicit nil

### UnsetCase
`func (o *Note) UnsetCase()`

UnsetCase ensures that no value is present for Case, not even an explicit nil
### GetNotestatus

`func (o *Note) GetNotestatus() int32`

GetNotestatus returns the Notestatus field if non-nil, zero value otherwise.

### GetNotestatusOk

`func (o *Note) GetNotestatusOk() (*int32, bool)`

GetNotestatusOk returns a tuple with the Notestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotestatus

`func (o *Note) SetNotestatus(v int32)`

SetNotestatus sets Notestatus field to given value.

### HasNotestatus

`func (o *Note) HasNotestatus() bool`

HasNotestatus returns a boolean if a field has been set.

### GetTag

`func (o *Note) GetTag() []int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Note) GetTagOk() (*[]int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Note) SetTag(v []int32)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Note) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetNoteCreateTime

`func (o *Note) GetNoteCreateTime() time.Time`

GetNoteCreateTime returns the NoteCreateTime field if non-nil, zero value otherwise.

### GetNoteCreateTimeOk

`func (o *Note) GetNoteCreateTimeOk() (*time.Time, bool)`

GetNoteCreateTimeOk returns a tuple with the NoteCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCreateTime

`func (o *Note) SetNoteCreateTime(v time.Time)`

SetNoteCreateTime sets NoteCreateTime field to given value.

### HasNoteCreateTime

`func (o *Note) HasNoteCreateTime() bool`

HasNoteCreateTime returns a boolean if a field has been set.

### GetNoteCreatedByUserId

`func (o *Note) GetNoteCreatedByUserId() int32`

GetNoteCreatedByUserId returns the NoteCreatedByUserId field if non-nil, zero value otherwise.

### GetNoteCreatedByUserIdOk

`func (o *Note) GetNoteCreatedByUserIdOk() (*int32, bool)`

GetNoteCreatedByUserIdOk returns a tuple with the NoteCreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteCreatedByUserId

`func (o *Note) SetNoteCreatedByUserId(v int32)`

SetNoteCreatedByUserId sets NoteCreatedByUserId field to given value.


### GetNoteModifyTime

`func (o *Note) GetNoteModifyTime() time.Time`

GetNoteModifyTime returns the NoteModifyTime field if non-nil, zero value otherwise.

### GetNoteModifyTimeOk

`func (o *Note) GetNoteModifyTimeOk() (*time.Time, bool)`

GetNoteModifyTimeOk returns a tuple with the NoteModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteModifyTime

`func (o *Note) SetNoteModifyTime(v time.Time)`

SetNoteModifyTime sets NoteModifyTime field to given value.

### HasNoteModifyTime

`func (o *Note) HasNoteModifyTime() bool`

HasNoteModifyTime returns a boolean if a field has been set.

### GetNoteModifiedByUserId

`func (o *Note) GetNoteModifiedByUserId() int32`

GetNoteModifiedByUserId returns the NoteModifiedByUserId field if non-nil, zero value otherwise.

### GetNoteModifiedByUserIdOk

`func (o *Note) GetNoteModifiedByUserIdOk() (*int32, bool)`

GetNoteModifiedByUserIdOk returns a tuple with the NoteModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteModifiedByUserId

`func (o *Note) SetNoteModifiedByUserId(v int32)`

SetNoteModifiedByUserId sets NoteModifiedByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


