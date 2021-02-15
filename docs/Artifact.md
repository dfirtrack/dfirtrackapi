# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactId** | Pointer to **int32** |  | [optional] [readonly] 
**ArtifactUuid** | Pointer to **string** |  | [optional] [readonly] 
**ArtifactName** | **string** |  | 
**Artifactstatus** | Pointer to **int32** |  | [optional] 
**Artifacttype** | **int32** |  | 
**Case** | Pointer to **NullableInt32** |  | [optional] 
**System** | **int32** |  | 
**ArtifactMd5** | Pointer to **NullableString** |  | [optional] 
**ArtifactSha1** | Pointer to **NullableString** |  | [optional] 
**ArtifactSha256** | Pointer to **NullableString** |  | [optional] 
**ArtifactSourcePath** | Pointer to **NullableString** |  | [optional] 
**ArtifactStoragePath** | Pointer to **string** |  | [optional] [readonly] 
**ArtifactAcquisitionTime** | Pointer to **NullableTime** |  | [optional] 
**ArtifactRequestedTime** | Pointer to **NullableTime** |  | [optional] 
**ArtifactCreateTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**ArtifactCreatedByUserId** | **int32** |  | 
**ArtifactModifyTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**ArtifactModifiedByUserId** | **int32** |  | 

## Methods

### NewArtifact

`func NewArtifact(artifactName string, artifacttype int32, system int32, artifactCreatedByUserId int32, artifactModifiedByUserId int32, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactId

`func (o *Artifact) GetArtifactId() int32`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *Artifact) GetArtifactIdOk() (*int32, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *Artifact) SetArtifactId(v int32)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *Artifact) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetArtifactUuid

`func (o *Artifact) GetArtifactUuid() string`

GetArtifactUuid returns the ArtifactUuid field if non-nil, zero value otherwise.

### GetArtifactUuidOk

`func (o *Artifact) GetArtifactUuidOk() (*string, bool)`

GetArtifactUuidOk returns a tuple with the ArtifactUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactUuid

`func (o *Artifact) SetArtifactUuid(v string)`

SetArtifactUuid sets ArtifactUuid field to given value.

### HasArtifactUuid

`func (o *Artifact) HasArtifactUuid() bool`

HasArtifactUuid returns a boolean if a field has been set.

### GetArtifactName

`func (o *Artifact) GetArtifactName() string`

GetArtifactName returns the ArtifactName field if non-nil, zero value otherwise.

### GetArtifactNameOk

`func (o *Artifact) GetArtifactNameOk() (*string, bool)`

GetArtifactNameOk returns a tuple with the ArtifactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactName

`func (o *Artifact) SetArtifactName(v string)`

SetArtifactName sets ArtifactName field to given value.


### GetArtifactstatus

`func (o *Artifact) GetArtifactstatus() int32`

GetArtifactstatus returns the Artifactstatus field if non-nil, zero value otherwise.

### GetArtifactstatusOk

`func (o *Artifact) GetArtifactstatusOk() (*int32, bool)`

GetArtifactstatusOk returns a tuple with the Artifactstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactstatus

`func (o *Artifact) SetArtifactstatus(v int32)`

SetArtifactstatus sets Artifactstatus field to given value.

### HasArtifactstatus

`func (o *Artifact) HasArtifactstatus() bool`

HasArtifactstatus returns a boolean if a field has been set.

### GetArtifacttype

`func (o *Artifact) GetArtifacttype() int32`

GetArtifacttype returns the Artifacttype field if non-nil, zero value otherwise.

### GetArtifacttypeOk

`func (o *Artifact) GetArtifacttypeOk() (*int32, bool)`

GetArtifacttypeOk returns a tuple with the Artifacttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacttype

`func (o *Artifact) SetArtifacttype(v int32)`

SetArtifacttype sets Artifacttype field to given value.


### GetCase

`func (o *Artifact) GetCase() int32`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *Artifact) GetCaseOk() (*int32, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *Artifact) SetCase(v int32)`

SetCase sets Case field to given value.

### HasCase

`func (o *Artifact) HasCase() bool`

HasCase returns a boolean if a field has been set.

### SetCaseNil

`func (o *Artifact) SetCaseNil(b bool)`

 SetCaseNil sets the value for Case to be an explicit nil

### UnsetCase
`func (o *Artifact) UnsetCase()`

UnsetCase ensures that no value is present for Case, not even an explicit nil
### GetSystem

`func (o *Artifact) GetSystem() int32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Artifact) GetSystemOk() (*int32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Artifact) SetSystem(v int32)`

SetSystem sets System field to given value.


### GetArtifactMd5

`func (o *Artifact) GetArtifactMd5() string`

GetArtifactMd5 returns the ArtifactMd5 field if non-nil, zero value otherwise.

### GetArtifactMd5Ok

`func (o *Artifact) GetArtifactMd5Ok() (*string, bool)`

GetArtifactMd5Ok returns a tuple with the ArtifactMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactMd5

`func (o *Artifact) SetArtifactMd5(v string)`

SetArtifactMd5 sets ArtifactMd5 field to given value.

### HasArtifactMd5

`func (o *Artifact) HasArtifactMd5() bool`

HasArtifactMd5 returns a boolean if a field has been set.

### SetArtifactMd5Nil

`func (o *Artifact) SetArtifactMd5Nil(b bool)`

 SetArtifactMd5Nil sets the value for ArtifactMd5 to be an explicit nil

### UnsetArtifactMd5
`func (o *Artifact) UnsetArtifactMd5()`

UnsetArtifactMd5 ensures that no value is present for ArtifactMd5, not even an explicit nil
### GetArtifactSha1

`func (o *Artifact) GetArtifactSha1() string`

GetArtifactSha1 returns the ArtifactSha1 field if non-nil, zero value otherwise.

### GetArtifactSha1Ok

`func (o *Artifact) GetArtifactSha1Ok() (*string, bool)`

GetArtifactSha1Ok returns a tuple with the ArtifactSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSha1

`func (o *Artifact) SetArtifactSha1(v string)`

SetArtifactSha1 sets ArtifactSha1 field to given value.

### HasArtifactSha1

`func (o *Artifact) HasArtifactSha1() bool`

HasArtifactSha1 returns a boolean if a field has been set.

### SetArtifactSha1Nil

`func (o *Artifact) SetArtifactSha1Nil(b bool)`

 SetArtifactSha1Nil sets the value for ArtifactSha1 to be an explicit nil

### UnsetArtifactSha1
`func (o *Artifact) UnsetArtifactSha1()`

UnsetArtifactSha1 ensures that no value is present for ArtifactSha1, not even an explicit nil
### GetArtifactSha256

`func (o *Artifact) GetArtifactSha256() string`

GetArtifactSha256 returns the ArtifactSha256 field if non-nil, zero value otherwise.

### GetArtifactSha256Ok

`func (o *Artifact) GetArtifactSha256Ok() (*string, bool)`

GetArtifactSha256Ok returns a tuple with the ArtifactSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSha256

`func (o *Artifact) SetArtifactSha256(v string)`

SetArtifactSha256 sets ArtifactSha256 field to given value.

### HasArtifactSha256

`func (o *Artifact) HasArtifactSha256() bool`

HasArtifactSha256 returns a boolean if a field has been set.

### SetArtifactSha256Nil

`func (o *Artifact) SetArtifactSha256Nil(b bool)`

 SetArtifactSha256Nil sets the value for ArtifactSha256 to be an explicit nil

### UnsetArtifactSha256
`func (o *Artifact) UnsetArtifactSha256()`

UnsetArtifactSha256 ensures that no value is present for ArtifactSha256, not even an explicit nil
### GetArtifactSourcePath

`func (o *Artifact) GetArtifactSourcePath() string`

GetArtifactSourcePath returns the ArtifactSourcePath field if non-nil, zero value otherwise.

### GetArtifactSourcePathOk

`func (o *Artifact) GetArtifactSourcePathOk() (*string, bool)`

GetArtifactSourcePathOk returns a tuple with the ArtifactSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSourcePath

`func (o *Artifact) SetArtifactSourcePath(v string)`

SetArtifactSourcePath sets ArtifactSourcePath field to given value.

### HasArtifactSourcePath

`func (o *Artifact) HasArtifactSourcePath() bool`

HasArtifactSourcePath returns a boolean if a field has been set.

### SetArtifactSourcePathNil

`func (o *Artifact) SetArtifactSourcePathNil(b bool)`

 SetArtifactSourcePathNil sets the value for ArtifactSourcePath to be an explicit nil

### UnsetArtifactSourcePath
`func (o *Artifact) UnsetArtifactSourcePath()`

UnsetArtifactSourcePath ensures that no value is present for ArtifactSourcePath, not even an explicit nil
### GetArtifactStoragePath

`func (o *Artifact) GetArtifactStoragePath() string`

GetArtifactStoragePath returns the ArtifactStoragePath field if non-nil, zero value otherwise.

### GetArtifactStoragePathOk

`func (o *Artifact) GetArtifactStoragePathOk() (*string, bool)`

GetArtifactStoragePathOk returns a tuple with the ArtifactStoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactStoragePath

`func (o *Artifact) SetArtifactStoragePath(v string)`

SetArtifactStoragePath sets ArtifactStoragePath field to given value.

### HasArtifactStoragePath

`func (o *Artifact) HasArtifactStoragePath() bool`

HasArtifactStoragePath returns a boolean if a field has been set.

### GetArtifactAcquisitionTime

`func (o *Artifact) GetArtifactAcquisitionTime() time.Time`

GetArtifactAcquisitionTime returns the ArtifactAcquisitionTime field if non-nil, zero value otherwise.

### GetArtifactAcquisitionTimeOk

`func (o *Artifact) GetArtifactAcquisitionTimeOk() (*time.Time, bool)`

GetArtifactAcquisitionTimeOk returns a tuple with the ArtifactAcquisitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactAcquisitionTime

`func (o *Artifact) SetArtifactAcquisitionTime(v time.Time)`

SetArtifactAcquisitionTime sets ArtifactAcquisitionTime field to given value.

### HasArtifactAcquisitionTime

`func (o *Artifact) HasArtifactAcquisitionTime() bool`

HasArtifactAcquisitionTime returns a boolean if a field has been set.

### SetArtifactAcquisitionTimeNil

`func (o *Artifact) SetArtifactAcquisitionTimeNil(b bool)`

 SetArtifactAcquisitionTimeNil sets the value for ArtifactAcquisitionTime to be an explicit nil

### UnsetArtifactAcquisitionTime
`func (o *Artifact) UnsetArtifactAcquisitionTime()`

UnsetArtifactAcquisitionTime ensures that no value is present for ArtifactAcquisitionTime, not even an explicit nil
### GetArtifactRequestedTime

`func (o *Artifact) GetArtifactRequestedTime() time.Time`

GetArtifactRequestedTime returns the ArtifactRequestedTime field if non-nil, zero value otherwise.

### GetArtifactRequestedTimeOk

`func (o *Artifact) GetArtifactRequestedTimeOk() (*time.Time, bool)`

GetArtifactRequestedTimeOk returns a tuple with the ArtifactRequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRequestedTime

`func (o *Artifact) SetArtifactRequestedTime(v time.Time)`

SetArtifactRequestedTime sets ArtifactRequestedTime field to given value.

### HasArtifactRequestedTime

`func (o *Artifact) HasArtifactRequestedTime() bool`

HasArtifactRequestedTime returns a boolean if a field has been set.

### SetArtifactRequestedTimeNil

`func (o *Artifact) SetArtifactRequestedTimeNil(b bool)`

 SetArtifactRequestedTimeNil sets the value for ArtifactRequestedTime to be an explicit nil

### UnsetArtifactRequestedTime
`func (o *Artifact) UnsetArtifactRequestedTime()`

UnsetArtifactRequestedTime ensures that no value is present for ArtifactRequestedTime, not even an explicit nil
### GetArtifactCreateTime

`func (o *Artifact) GetArtifactCreateTime() time.Time`

GetArtifactCreateTime returns the ArtifactCreateTime field if non-nil, zero value otherwise.

### GetArtifactCreateTimeOk

`func (o *Artifact) GetArtifactCreateTimeOk() (*time.Time, bool)`

GetArtifactCreateTimeOk returns a tuple with the ArtifactCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCreateTime

`func (o *Artifact) SetArtifactCreateTime(v time.Time)`

SetArtifactCreateTime sets ArtifactCreateTime field to given value.

### HasArtifactCreateTime

`func (o *Artifact) HasArtifactCreateTime() bool`

HasArtifactCreateTime returns a boolean if a field has been set.

### GetArtifactCreatedByUserId

`func (o *Artifact) GetArtifactCreatedByUserId() int32`

GetArtifactCreatedByUserId returns the ArtifactCreatedByUserId field if non-nil, zero value otherwise.

### GetArtifactCreatedByUserIdOk

`func (o *Artifact) GetArtifactCreatedByUserIdOk() (*int32, bool)`

GetArtifactCreatedByUserIdOk returns a tuple with the ArtifactCreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCreatedByUserId

`func (o *Artifact) SetArtifactCreatedByUserId(v int32)`

SetArtifactCreatedByUserId sets ArtifactCreatedByUserId field to given value.


### GetArtifactModifyTime

`func (o *Artifact) GetArtifactModifyTime() time.Time`

GetArtifactModifyTime returns the ArtifactModifyTime field if non-nil, zero value otherwise.

### GetArtifactModifyTimeOk

`func (o *Artifact) GetArtifactModifyTimeOk() (*time.Time, bool)`

GetArtifactModifyTimeOk returns a tuple with the ArtifactModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactModifyTime

`func (o *Artifact) SetArtifactModifyTime(v time.Time)`

SetArtifactModifyTime sets ArtifactModifyTime field to given value.

### HasArtifactModifyTime

`func (o *Artifact) HasArtifactModifyTime() bool`

HasArtifactModifyTime returns a boolean if a field has been set.

### GetArtifactModifiedByUserId

`func (o *Artifact) GetArtifactModifiedByUserId() int32`

GetArtifactModifiedByUserId returns the ArtifactModifiedByUserId field if non-nil, zero value otherwise.

### GetArtifactModifiedByUserIdOk

`func (o *Artifact) GetArtifactModifiedByUserIdOk() (*int32, bool)`

GetArtifactModifiedByUserIdOk returns a tuple with the ArtifactModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactModifiedByUserId

`func (o *Artifact) SetArtifactModifiedByUserId(v int32)`

SetArtifactModifiedByUserId sets ArtifactModifiedByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


