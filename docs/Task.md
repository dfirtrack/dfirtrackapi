# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **int32** |  | [optional] [readonly] 
**ParentTask** | Pointer to **NullableInt32** |  | [optional] 
**Taskname** | **int32** |  | 
**Taskpriority** | **int32** |  | 
**Taskstatus** | **int32** |  | 
**System** | Pointer to **NullableInt32** |  | [optional] 
**TaskAssignedToUserId** | Pointer to **NullableInt32** |  | [optional] 
**Tag** | Pointer to **[]int32** |  | [optional] 
**TaskScheduledTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**TaskStartedTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**TaskFinishedTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**TaskDueTime** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] 
**TaskCreateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**TaskModifyTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**TaskCreatedByUserId** | **int32** |  | 
**TaskModifiedByUserId** | **int32** |  | 

## Methods

### NewTask

`func NewTask(taskname int32, taskpriority int32, taskstatus int32, taskCreatedByUserId int32, taskModifiedByUserId int32, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *Task) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Task) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Task) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Task) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetParentTask

`func (o *Task) GetParentTask() int32`

GetParentTask returns the ParentTask field if non-nil, zero value otherwise.

### GetParentTaskOk

`func (o *Task) GetParentTaskOk() (*int32, bool)`

GetParentTaskOk returns a tuple with the ParentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTask

`func (o *Task) SetParentTask(v int32)`

SetParentTask sets ParentTask field to given value.

### HasParentTask

`func (o *Task) HasParentTask() bool`

HasParentTask returns a boolean if a field has been set.

### SetParentTaskNil

`func (o *Task) SetParentTaskNil(b bool)`

 SetParentTaskNil sets the value for ParentTask to be an explicit nil

### UnsetParentTask
`func (o *Task) UnsetParentTask()`

UnsetParentTask ensures that no value is present for ParentTask, not even an explicit nil
### GetTaskname

`func (o *Task) GetTaskname() int32`

GetTaskname returns the Taskname field if non-nil, zero value otherwise.

### GetTasknameOk

`func (o *Task) GetTasknameOk() (*int32, bool)`

GetTasknameOk returns a tuple with the Taskname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskname

`func (o *Task) SetTaskname(v int32)`

SetTaskname sets Taskname field to given value.


### GetTaskpriority

`func (o *Task) GetTaskpriority() int32`

GetTaskpriority returns the Taskpriority field if non-nil, zero value otherwise.

### GetTaskpriorityOk

`func (o *Task) GetTaskpriorityOk() (*int32, bool)`

GetTaskpriorityOk returns a tuple with the Taskpriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskpriority

`func (o *Task) SetTaskpriority(v int32)`

SetTaskpriority sets Taskpriority field to given value.


### GetTaskstatus

`func (o *Task) GetTaskstatus() int32`

GetTaskstatus returns the Taskstatus field if non-nil, zero value otherwise.

### GetTaskstatusOk

`func (o *Task) GetTaskstatusOk() (*int32, bool)`

GetTaskstatusOk returns a tuple with the Taskstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskstatus

`func (o *Task) SetTaskstatus(v int32)`

SetTaskstatus sets Taskstatus field to given value.


### GetSystem

`func (o *Task) GetSystem() int32`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Task) GetSystemOk() (*int32, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Task) SetSystem(v int32)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Task) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *Task) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *Task) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetTaskAssignedToUserId

`func (o *Task) GetTaskAssignedToUserId() int32`

GetTaskAssignedToUserId returns the TaskAssignedToUserId field if non-nil, zero value otherwise.

### GetTaskAssignedToUserIdOk

`func (o *Task) GetTaskAssignedToUserIdOk() (*int32, bool)`

GetTaskAssignedToUserIdOk returns a tuple with the TaskAssignedToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAssignedToUserId

`func (o *Task) SetTaskAssignedToUserId(v int32)`

SetTaskAssignedToUserId sets TaskAssignedToUserId field to given value.

### HasTaskAssignedToUserId

`func (o *Task) HasTaskAssignedToUserId() bool`

HasTaskAssignedToUserId returns a boolean if a field has been set.

### SetTaskAssignedToUserIdNil

`func (o *Task) SetTaskAssignedToUserIdNil(b bool)`

 SetTaskAssignedToUserIdNil sets the value for TaskAssignedToUserId to be an explicit nil

### UnsetTaskAssignedToUserId
`func (o *Task) UnsetTaskAssignedToUserId()`

UnsetTaskAssignedToUserId ensures that no value is present for TaskAssignedToUserId, not even an explicit nil
### GetTag

`func (o *Task) GetTag() []int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Task) GetTagOk() (*[]int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Task) SetTag(v []int32)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Task) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTaskScheduledTime

`func (o *Task) GetTaskScheduledTime() time.Time`

GetTaskScheduledTime returns the TaskScheduledTime field if non-nil, zero value otherwise.

### GetTaskScheduledTimeOk

`func (o *Task) GetTaskScheduledTimeOk() (*time.Time, bool)`

GetTaskScheduledTimeOk returns a tuple with the TaskScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskScheduledTime

`func (o *Task) SetTaskScheduledTime(v time.Time)`

SetTaskScheduledTime sets TaskScheduledTime field to given value.

### HasTaskScheduledTime

`func (o *Task) HasTaskScheduledTime() bool`

HasTaskScheduledTime returns a boolean if a field has been set.

### SetTaskScheduledTimeNil

`func (o *Task) SetTaskScheduledTimeNil(b bool)`

 SetTaskScheduledTimeNil sets the value for TaskScheduledTime to be an explicit nil

### UnsetTaskScheduledTime
`func (o *Task) UnsetTaskScheduledTime()`

UnsetTaskScheduledTime ensures that no value is present for TaskScheduledTime, not even an explicit nil
### GetTaskStartedTime

`func (o *Task) GetTaskStartedTime() time.Time`

GetTaskStartedTime returns the TaskStartedTime field if non-nil, zero value otherwise.

### GetTaskStartedTimeOk

`func (o *Task) GetTaskStartedTimeOk() (*time.Time, bool)`

GetTaskStartedTimeOk returns a tuple with the TaskStartedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStartedTime

`func (o *Task) SetTaskStartedTime(v time.Time)`

SetTaskStartedTime sets TaskStartedTime field to given value.

### HasTaskStartedTime

`func (o *Task) HasTaskStartedTime() bool`

HasTaskStartedTime returns a boolean if a field has been set.

### SetTaskStartedTimeNil

`func (o *Task) SetTaskStartedTimeNil(b bool)`

 SetTaskStartedTimeNil sets the value for TaskStartedTime to be an explicit nil

### UnsetTaskStartedTime
`func (o *Task) UnsetTaskStartedTime()`

UnsetTaskStartedTime ensures that no value is present for TaskStartedTime, not even an explicit nil
### GetTaskFinishedTime

`func (o *Task) GetTaskFinishedTime() time.Time`

GetTaskFinishedTime returns the TaskFinishedTime field if non-nil, zero value otherwise.

### GetTaskFinishedTimeOk

`func (o *Task) GetTaskFinishedTimeOk() (*time.Time, bool)`

GetTaskFinishedTimeOk returns a tuple with the TaskFinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFinishedTime

`func (o *Task) SetTaskFinishedTime(v time.Time)`

SetTaskFinishedTime sets TaskFinishedTime field to given value.

### HasTaskFinishedTime

`func (o *Task) HasTaskFinishedTime() bool`

HasTaskFinishedTime returns a boolean if a field has been set.

### SetTaskFinishedTimeNil

`func (o *Task) SetTaskFinishedTimeNil(b bool)`

 SetTaskFinishedTimeNil sets the value for TaskFinishedTime to be an explicit nil

### UnsetTaskFinishedTime
`func (o *Task) UnsetTaskFinishedTime()`

UnsetTaskFinishedTime ensures that no value is present for TaskFinishedTime, not even an explicit nil
### GetTaskDueTime

`func (o *Task) GetTaskDueTime() time.Time`

GetTaskDueTime returns the TaskDueTime field if non-nil, zero value otherwise.

### GetTaskDueTimeOk

`func (o *Task) GetTaskDueTimeOk() (*time.Time, bool)`

GetTaskDueTimeOk returns a tuple with the TaskDueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDueTime

`func (o *Task) SetTaskDueTime(v time.Time)`

SetTaskDueTime sets TaskDueTime field to given value.

### HasTaskDueTime

`func (o *Task) HasTaskDueTime() bool`

HasTaskDueTime returns a boolean if a field has been set.

### SetTaskDueTimeNil

`func (o *Task) SetTaskDueTimeNil(b bool)`

 SetTaskDueTimeNil sets the value for TaskDueTime to be an explicit nil

### UnsetTaskDueTime
`func (o *Task) UnsetTaskDueTime()`

UnsetTaskDueTime ensures that no value is present for TaskDueTime, not even an explicit nil
### GetTaskCreateTime

`func (o *Task) GetTaskCreateTime() time.Time`

GetTaskCreateTime returns the TaskCreateTime field if non-nil, zero value otherwise.

### GetTaskCreateTimeOk

`func (o *Task) GetTaskCreateTimeOk() (*time.Time, bool)`

GetTaskCreateTimeOk returns a tuple with the TaskCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCreateTime

`func (o *Task) SetTaskCreateTime(v time.Time)`

SetTaskCreateTime sets TaskCreateTime field to given value.

### HasTaskCreateTime

`func (o *Task) HasTaskCreateTime() bool`

HasTaskCreateTime returns a boolean if a field has been set.

### GetTaskModifyTime

`func (o *Task) GetTaskModifyTime() time.Time`

GetTaskModifyTime returns the TaskModifyTime field if non-nil, zero value otherwise.

### GetTaskModifyTimeOk

`func (o *Task) GetTaskModifyTimeOk() (*time.Time, bool)`

GetTaskModifyTimeOk returns a tuple with the TaskModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskModifyTime

`func (o *Task) SetTaskModifyTime(v time.Time)`

SetTaskModifyTime sets TaskModifyTime field to given value.

### HasTaskModifyTime

`func (o *Task) HasTaskModifyTime() bool`

HasTaskModifyTime returns a boolean if a field has been set.

### GetTaskCreatedByUserId

`func (o *Task) GetTaskCreatedByUserId() int32`

GetTaskCreatedByUserId returns the TaskCreatedByUserId field if non-nil, zero value otherwise.

### GetTaskCreatedByUserIdOk

`func (o *Task) GetTaskCreatedByUserIdOk() (*int32, bool)`

GetTaskCreatedByUserIdOk returns a tuple with the TaskCreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCreatedByUserId

`func (o *Task) SetTaskCreatedByUserId(v int32)`

SetTaskCreatedByUserId sets TaskCreatedByUserId field to given value.


### GetTaskModifiedByUserId

`func (o *Task) GetTaskModifiedByUserId() int32`

GetTaskModifiedByUserId returns the TaskModifiedByUserId field if non-nil, zero value otherwise.

### GetTaskModifiedByUserIdOk

`func (o *Task) GetTaskModifiedByUserIdOk() (*int32, bool)`

GetTaskModifiedByUserIdOk returns a tuple with the TaskModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskModifiedByUserId

`func (o *Task) SetTaskModifiedByUserId(v int32)`

SetTaskModifiedByUserId sets TaskModifiedByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


