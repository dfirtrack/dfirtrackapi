# Go API client for dfirtrackapi

OpenAPI 3 - Documentation of DFIRTrack API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2.0.0
- Package version: 2.2.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./dfirtrackapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiApi* | [**CreateArtifact**](docs/ApiApi.md#createartifact) | **Post** /api/artifact/ | 
*ApiApi* | [**CreateArtifacttype**](docs/ApiApi.md#createartifacttype) | **Post** /api/artifacttype/ | 
*ApiApi* | [**CreateAuthToken**](docs/ApiApi.md#createauthtoken) | **Post** /api/token-auth/ | 
*ApiApi* | [**CreateCase**](docs/ApiApi.md#createcase) | **Post** /api/case/ | 
*ApiApi* | [**CreateCasetype**](docs/ApiApi.md#createcasetype) | **Post** /api/casetype/ | 
*ApiApi* | [**CreateCompany**](docs/ApiApi.md#createcompany) | **Post** /api/company/ | 
*ApiApi* | [**CreateContact**](docs/ApiApi.md#createcontact) | **Post** /api/contact/ | 
*ApiApi* | [**CreateDivision**](docs/ApiApi.md#createdivision) | **Post** /api/division/ | 
*ApiApi* | [**CreateDnsname**](docs/ApiApi.md#creatednsname) | **Post** /api/dnsname/ | 
*ApiApi* | [**CreateDomain**](docs/ApiApi.md#createdomain) | **Post** /api/domain/ | 
*ApiApi* | [**CreateDomainuser**](docs/ApiApi.md#createdomainuser) | **Post** /api/domainuser/ | 
*ApiApi* | [**CreateHeadline**](docs/ApiApi.md#createheadline) | **Post** /api/headline/ | 
*ApiApi* | [**CreateIp**](docs/ApiApi.md#createip) | **Post** /api/ip/ | 
*ApiApi* | [**CreateLocation**](docs/ApiApi.md#createlocation) | **Post** /api/location/ | 
*ApiApi* | [**CreateNote**](docs/ApiApi.md#createnote) | **Post** /api/note/ | 
*ApiApi* | [**CreateOs**](docs/ApiApi.md#createos) | **Post** /api/os/ | 
*ApiApi* | [**CreateOsarch**](docs/ApiApi.md#createosarch) | **Post** /api/osarch/ | 
*ApiApi* | [**CreateReason**](docs/ApiApi.md#createreason) | **Post** /api/reason/ | 
*ApiApi* | [**CreateRecommendation**](docs/ApiApi.md#createrecommendation) | **Post** /api/recommendation/ | 
*ApiApi* | [**CreateReportitem**](docs/ApiApi.md#createreportitem) | **Post** /api/reportitem/ | 
*ApiApi* | [**CreateServiceprovider**](docs/ApiApi.md#createserviceprovider) | **Post** /api/serviceprovider/ | 
*ApiApi* | [**CreateSystem**](docs/ApiApi.md#createsystem) | **Post** /api/system/ | 
*ApiApi* | [**CreateSystemtype**](docs/ApiApi.md#createsystemtype) | **Post** /api/systemtype/ | 
*ApiApi* | [**CreateSystemuser**](docs/ApiApi.md#createsystemuser) | **Post** /api/systemuser/ | 
*ApiApi* | [**CreateTag**](docs/ApiApi.md#createtag) | **Post** /api/tag/ | 
*ApiApi* | [**CreateTask**](docs/ApiApi.md#createtask) | **Post** /api/task/ | 
*ApiApi* | [**CreateTaskname**](docs/ApiApi.md#createtaskname) | **Post** /api/taskname/ | 
*ApiApi* | [**ListAnalysisstatus**](docs/ApiApi.md#listanalysisstatus) | **Get** /api/analysisstatus/ | 
*ApiApi* | [**ListArtifactprioritys**](docs/ApiApi.md#listartifactprioritys) | **Get** /api/artifactpriority/ | 
*ApiApi* | [**ListArtifacts**](docs/ApiApi.md#listartifacts) | **Get** /api/artifact/ | 
*ApiApi* | [**ListArtifactstatus**](docs/ApiApi.md#listartifactstatus) | **Get** /api/artifactstatus/ | 
*ApiApi* | [**ListArtifacttypes**](docs/ApiApi.md#listartifacttypes) | **Get** /api/artifacttype/ | 
*ApiApi* | [**ListCaseprioritys**](docs/ApiApi.md#listcaseprioritys) | **Get** /api/casepriority/ | 
*ApiApi* | [**ListCases**](docs/ApiApi.md#listcases) | **Get** /api/case/ | 
*ApiApi* | [**ListCasestatus**](docs/ApiApi.md#listcasestatus) | **Get** /api/casestatus/ | 
*ApiApi* | [**ListCasetypes**](docs/ApiApi.md#listcasetypes) | **Get** /api/casetype/ | 
*ApiApi* | [**ListCompanys**](docs/ApiApi.md#listcompanys) | **Get** /api/company/ | 
*ApiApi* | [**ListContacts**](docs/ApiApi.md#listcontacts) | **Get** /api/contact/ | 
*ApiApi* | [**ListDivisions**](docs/ApiApi.md#listdivisions) | **Get** /api/division/ | 
*ApiApi* | [**ListDnsnames**](docs/ApiApi.md#listdnsnames) | **Get** /api/dnsname/ | 
*ApiApi* | [**ListDomains**](docs/ApiApi.md#listdomains) | **Get** /api/domain/ | 
*ApiApi* | [**ListDomainusers**](docs/ApiApi.md#listdomainusers) | **Get** /api/domainuser/ | 
*ApiApi* | [**ListHeadlines**](docs/ApiApi.md#listheadlines) | **Get** /api/headline/ | 
*ApiApi* | [**ListIps**](docs/ApiApi.md#listips) | **Get** /api/ip/ | 
*ApiApi* | [**ListLocations**](docs/ApiApi.md#listlocations) | **Get** /api/location/ | 
*ApiApi* | [**ListNotes**](docs/ApiApi.md#listnotes) | **Get** /api/note/ | 
*ApiApi* | [**ListNotestatus**](docs/ApiApi.md#listnotestatus) | **Get** /api/notestatus/ | 
*ApiApi* | [**ListOs**](docs/ApiApi.md#listos) | **Get** /api/os/ | 
*ApiApi* | [**ListOsarchs**](docs/ApiApi.md#listosarchs) | **Get** /api/osarch/ | 
*ApiApi* | [**ListReasons**](docs/ApiApi.md#listreasons) | **Get** /api/reason/ | 
*ApiApi* | [**ListRecommendations**](docs/ApiApi.md#listrecommendations) | **Get** /api/recommendation/ | 
*ApiApi* | [**ListReportitems**](docs/ApiApi.md#listreportitems) | **Get** /api/reportitem/ | 
*ApiApi* | [**ListServiceproviders**](docs/ApiApi.md#listserviceproviders) | **Get** /api/serviceprovider/ | 
*ApiApi* | [**ListSystems**](docs/ApiApi.md#listsystems) | **Get** /api/system/ | 
*ApiApi* | [**ListSystemstatus**](docs/ApiApi.md#listsystemstatus) | **Get** /api/systemstatus/ | 
*ApiApi* | [**ListSystemtypes**](docs/ApiApi.md#listsystemtypes) | **Get** /api/systemtype/ | 
*ApiApi* | [**ListSystemusers**](docs/ApiApi.md#listsystemusers) | **Get** /api/systemuser/ | 
*ApiApi* | [**ListTagcolors**](docs/ApiApi.md#listtagcolors) | **Get** /api/tagcolor/ | 
*ApiApi* | [**ListTags**](docs/ApiApi.md#listtags) | **Get** /api/tag/ | 
*ApiApi* | [**ListTasknames**](docs/ApiApi.md#listtasknames) | **Get** /api/taskname/ | 
*ApiApi* | [**ListTaskprioritys**](docs/ApiApi.md#listtaskprioritys) | **Get** /api/taskpriority/ | 
*ApiApi* | [**ListTasks**](docs/ApiApi.md#listtasks) | **Get** /api/task/ | 
*ApiApi* | [**ListTaskstatus**](docs/ApiApi.md#listtaskstatus) | **Get** /api/taskstatus/ | 
*ApiApi* | [**PartialUpdateArtifact**](docs/ApiApi.md#partialupdateartifact) | **Patch** /api/artifact/{artifact_id}/ | 
*ApiApi* | [**PartialUpdateArtifacttype**](docs/ApiApi.md#partialupdateartifacttype) | **Patch** /api/artifacttype/{artifacttype_id}/ | 
*ApiApi* | [**PartialUpdateCase**](docs/ApiApi.md#partialupdatecase) | **Patch** /api/case/{case_id}/ | 
*ApiApi* | [**PartialUpdateCasetype**](docs/ApiApi.md#partialupdatecasetype) | **Patch** /api/casetype/{casetype_id}/ | 
*ApiApi* | [**PartialUpdateCompany**](docs/ApiApi.md#partialupdatecompany) | **Patch** /api/company/{company_id}/ | 
*ApiApi* | [**PartialUpdateContact**](docs/ApiApi.md#partialupdatecontact) | **Patch** /api/contact/{contact_id}/ | 
*ApiApi* | [**PartialUpdateDivision**](docs/ApiApi.md#partialupdatedivision) | **Patch** /api/division/{division_id}/ | 
*ApiApi* | [**PartialUpdateDnsname**](docs/ApiApi.md#partialupdatednsname) | **Patch** /api/dnsname/{dnsname_id}/ | 
*ApiApi* | [**PartialUpdateDomain**](docs/ApiApi.md#partialupdatedomain) | **Patch** /api/domain/{domain_id}/ | 
*ApiApi* | [**PartialUpdateDomainuser**](docs/ApiApi.md#partialupdatedomainuser) | **Patch** /api/domainuser/{domainuser_id}/ | 
*ApiApi* | [**PartialUpdateHeadline**](docs/ApiApi.md#partialupdateheadline) | **Patch** /api/headline/{headline_id}/ | 
*ApiApi* | [**PartialUpdateIp**](docs/ApiApi.md#partialupdateip) | **Patch** /api/ip/{ip_id}/ | 
*ApiApi* | [**PartialUpdateLocation**](docs/ApiApi.md#partialupdatelocation) | **Patch** /api/location/{location_id}/ | 
*ApiApi* | [**PartialUpdateNote**](docs/ApiApi.md#partialupdatenote) | **Patch** /api/note/{note_id}/ | 
*ApiApi* | [**PartialUpdateOs**](docs/ApiApi.md#partialupdateos) | **Patch** /api/os/{os_id}/ | 
*ApiApi* | [**PartialUpdateOsarch**](docs/ApiApi.md#partialupdateosarch) | **Patch** /api/osarch/{osarch_id}/ | 
*ApiApi* | [**PartialUpdateReason**](docs/ApiApi.md#partialupdatereason) | **Patch** /api/reason/{reason_id}/ | 
*ApiApi* | [**PartialUpdateRecommendation**](docs/ApiApi.md#partialupdaterecommendation) | **Patch** /api/recommendation/{recommendation_id}/ | 
*ApiApi* | [**PartialUpdateReportitem**](docs/ApiApi.md#partialupdatereportitem) | **Patch** /api/reportitem/{reportitem_id}/ | 
*ApiApi* | [**PartialUpdateServiceprovider**](docs/ApiApi.md#partialupdateserviceprovider) | **Patch** /api/serviceprovider/{serviceprovider_id}/ | 
*ApiApi* | [**PartialUpdateSystem**](docs/ApiApi.md#partialupdatesystem) | **Patch** /api/system/{system_id}/ | 
*ApiApi* | [**PartialUpdateSystemtype**](docs/ApiApi.md#partialupdatesystemtype) | **Patch** /api/systemtype/{systemtype_id}/ | 
*ApiApi* | [**PartialUpdateSystemuser**](docs/ApiApi.md#partialupdatesystemuser) | **Patch** /api/systemuser/{systemuser_id}/ | 
*ApiApi* | [**PartialUpdateTag**](docs/ApiApi.md#partialupdatetag) | **Patch** /api/tag/{tag_id}/ | 
*ApiApi* | [**PartialUpdateTask**](docs/ApiApi.md#partialupdatetask) | **Patch** /api/task/{task_id}/ | 
*ApiApi* | [**PartialUpdateTaskname**](docs/ApiApi.md#partialupdatetaskname) | **Patch** /api/taskname/{taskname_id}/ | 
*ApiApi* | [**RetrieveAnalysisstatus**](docs/ApiApi.md#retrieveanalysisstatus) | **Get** /api/analysisstatus/{analysisstatus_id}/ | 
*ApiApi* | [**RetrieveArtifact**](docs/ApiApi.md#retrieveartifact) | **Get** /api/artifact/{artifact_id}/ | 
*ApiApi* | [**RetrieveArtifactpriority**](docs/ApiApi.md#retrieveartifactpriority) | **Get** /api/artifactpriority/{artifactpriority_id}/ | 
*ApiApi* | [**RetrieveArtifactstatus**](docs/ApiApi.md#retrieveartifactstatus) | **Get** /api/artifactstatus/{artifactstatus_id}/ | 
*ApiApi* | [**RetrieveArtifacttype**](docs/ApiApi.md#retrieveartifacttype) | **Get** /api/artifacttype/{artifacttype_id}/ | 
*ApiApi* | [**RetrieveCase**](docs/ApiApi.md#retrievecase) | **Get** /api/case/{case_id}/ | 
*ApiApi* | [**RetrieveCasepriority**](docs/ApiApi.md#retrievecasepriority) | **Get** /api/casepriority/{casepriority_id}/ | 
*ApiApi* | [**RetrieveCasestatus**](docs/ApiApi.md#retrievecasestatus) | **Get** /api/casestatus/{casestatus_id}/ | 
*ApiApi* | [**RetrieveCasetype**](docs/ApiApi.md#retrievecasetype) | **Get** /api/casetype/{casetype_id}/ | 
*ApiApi* | [**RetrieveCompany**](docs/ApiApi.md#retrievecompany) | **Get** /api/company/{company_id}/ | 
*ApiApi* | [**RetrieveContact**](docs/ApiApi.md#retrievecontact) | **Get** /api/contact/{contact_id}/ | 
*ApiApi* | [**RetrieveDivision**](docs/ApiApi.md#retrievedivision) | **Get** /api/division/{division_id}/ | 
*ApiApi* | [**RetrieveDnsname**](docs/ApiApi.md#retrievednsname) | **Get** /api/dnsname/{dnsname_id}/ | 
*ApiApi* | [**RetrieveDomain**](docs/ApiApi.md#retrievedomain) | **Get** /api/domain/{domain_id}/ | 
*ApiApi* | [**RetrieveDomainuser**](docs/ApiApi.md#retrievedomainuser) | **Get** /api/domainuser/{domainuser_id}/ | 
*ApiApi* | [**RetrieveHeadline**](docs/ApiApi.md#retrieveheadline) | **Get** /api/headline/{headline_id}/ | 
*ApiApi* | [**RetrieveIp**](docs/ApiApi.md#retrieveip) | **Get** /api/ip/{ip_id}/ | 
*ApiApi* | [**RetrieveLocation**](docs/ApiApi.md#retrievelocation) | **Get** /api/location/{location_id}/ | 
*ApiApi* | [**RetrieveNote**](docs/ApiApi.md#retrievenote) | **Get** /api/note/{note_id}/ | 
*ApiApi* | [**RetrieveNotestatus**](docs/ApiApi.md#retrievenotestatus) | **Get** /api/notestatus/{notestatus_id}/ | 
*ApiApi* | [**RetrieveOs**](docs/ApiApi.md#retrieveos) | **Get** /api/os/{os_id}/ | 
*ApiApi* | [**RetrieveOsarch**](docs/ApiApi.md#retrieveosarch) | **Get** /api/osarch/{osarch_id}/ | 
*ApiApi* | [**RetrieveReason**](docs/ApiApi.md#retrievereason) | **Get** /api/reason/{reason_id}/ | 
*ApiApi* | [**RetrieveRecommendation**](docs/ApiApi.md#retrieverecommendation) | **Get** /api/recommendation/{recommendation_id}/ | 
*ApiApi* | [**RetrieveReportitem**](docs/ApiApi.md#retrievereportitem) | **Get** /api/reportitem/{reportitem_id}/ | 
*ApiApi* | [**RetrieveServiceprovider**](docs/ApiApi.md#retrieveserviceprovider) | **Get** /api/serviceprovider/{serviceprovider_id}/ | 
*ApiApi* | [**RetrieveSystem**](docs/ApiApi.md#retrievesystem) | **Get** /api/system/{system_id}/ | 
*ApiApi* | [**RetrieveSystemstatus**](docs/ApiApi.md#retrievesystemstatus) | **Get** /api/systemstatus/{systemstatus_id}/ | 
*ApiApi* | [**RetrieveSystemtype**](docs/ApiApi.md#retrievesystemtype) | **Get** /api/systemtype/{systemtype_id}/ | 
*ApiApi* | [**RetrieveSystemuser**](docs/ApiApi.md#retrievesystemuser) | **Get** /api/systemuser/{systemuser_id}/ | 
*ApiApi* | [**RetrieveTag**](docs/ApiApi.md#retrievetag) | **Get** /api/tag/{tag_id}/ | 
*ApiApi* | [**RetrieveTagcolor**](docs/ApiApi.md#retrievetagcolor) | **Get** /api/tagcolor/{tagcolor_id}/ | 
*ApiApi* | [**RetrieveTask**](docs/ApiApi.md#retrievetask) | **Get** /api/task/{task_id}/ | 
*ApiApi* | [**RetrieveTaskname**](docs/ApiApi.md#retrievetaskname) | **Get** /api/taskname/{taskname_id}/ | 
*ApiApi* | [**RetrieveTaskpriority**](docs/ApiApi.md#retrievetaskpriority) | **Get** /api/taskpriority/{taskpriority_id}/ | 
*ApiApi* | [**RetrieveTaskstatus**](docs/ApiApi.md#retrievetaskstatus) | **Get** /api/taskstatus/{taskstatus_id}/ | 
*ApiApi* | [**UpdateArtifact**](docs/ApiApi.md#updateartifact) | **Put** /api/artifact/{artifact_id}/ | 
*ApiApi* | [**UpdateArtifacttype**](docs/ApiApi.md#updateartifacttype) | **Put** /api/artifacttype/{artifacttype_id}/ | 
*ApiApi* | [**UpdateCase**](docs/ApiApi.md#updatecase) | **Put** /api/case/{case_id}/ | 
*ApiApi* | [**UpdateCasetype**](docs/ApiApi.md#updatecasetype) | **Put** /api/casetype/{casetype_id}/ | 
*ApiApi* | [**UpdateCompany**](docs/ApiApi.md#updatecompany) | **Put** /api/company/{company_id}/ | 
*ApiApi* | [**UpdateContact**](docs/ApiApi.md#updatecontact) | **Put** /api/contact/{contact_id}/ | 
*ApiApi* | [**UpdateDivision**](docs/ApiApi.md#updatedivision) | **Put** /api/division/{division_id}/ | 
*ApiApi* | [**UpdateDnsname**](docs/ApiApi.md#updatednsname) | **Put** /api/dnsname/{dnsname_id}/ | 
*ApiApi* | [**UpdateDomain**](docs/ApiApi.md#updatedomain) | **Put** /api/domain/{domain_id}/ | 
*ApiApi* | [**UpdateDomainuser**](docs/ApiApi.md#updatedomainuser) | **Put** /api/domainuser/{domainuser_id}/ | 
*ApiApi* | [**UpdateHeadline**](docs/ApiApi.md#updateheadline) | **Put** /api/headline/{headline_id}/ | 
*ApiApi* | [**UpdateIp**](docs/ApiApi.md#updateip) | **Put** /api/ip/{ip_id}/ | 
*ApiApi* | [**UpdateLocation**](docs/ApiApi.md#updatelocation) | **Put** /api/location/{location_id}/ | 
*ApiApi* | [**UpdateNote**](docs/ApiApi.md#updatenote) | **Put** /api/note/{note_id}/ | 
*ApiApi* | [**UpdateOs**](docs/ApiApi.md#updateos) | **Put** /api/os/{os_id}/ | 
*ApiApi* | [**UpdateOsarch**](docs/ApiApi.md#updateosarch) | **Put** /api/osarch/{osarch_id}/ | 
*ApiApi* | [**UpdateReason**](docs/ApiApi.md#updatereason) | **Put** /api/reason/{reason_id}/ | 
*ApiApi* | [**UpdateRecommendation**](docs/ApiApi.md#updaterecommendation) | **Put** /api/recommendation/{recommendation_id}/ | 
*ApiApi* | [**UpdateReportitem**](docs/ApiApi.md#updatereportitem) | **Put** /api/reportitem/{reportitem_id}/ | 
*ApiApi* | [**UpdateServiceprovider**](docs/ApiApi.md#updateserviceprovider) | **Put** /api/serviceprovider/{serviceprovider_id}/ | 
*ApiApi* | [**UpdateSystem**](docs/ApiApi.md#updatesystem) | **Put** /api/system/{system_id}/ | 
*ApiApi* | [**UpdateSystemtype**](docs/ApiApi.md#updatesystemtype) | **Put** /api/systemtype/{systemtype_id}/ | 
*ApiApi* | [**UpdateSystemuser**](docs/ApiApi.md#updatesystemuser) | **Put** /api/systemuser/{systemuser_id}/ | 
*ApiApi* | [**UpdateTag**](docs/ApiApi.md#updatetag) | **Put** /api/tag/{tag_id}/ | 
*ApiApi* | [**UpdateTask**](docs/ApiApi.md#updatetask) | **Put** /api/task/{task_id}/ | 
*ApiApi* | [**UpdateTaskname**](docs/ApiApi.md#updatetaskname) | **Put** /api/taskname/{taskname_id}/ | 


## Documentation For Models

 - [Analysisstatus](docs/Analysisstatus.md)
 - [Artifact](docs/Artifact.md)
 - [Artifactpriority](docs/Artifactpriority.md)
 - [Artifactstatus](docs/Artifactstatus.md)
 - [Artifacttype](docs/Artifacttype.md)
 - [AuthToken](docs/AuthToken.md)
 - [Case](docs/Case.md)
 - [Casepriority](docs/Casepriority.md)
 - [Casestatus](docs/Casestatus.md)
 - [Casetype](docs/Casetype.md)
 - [Company](docs/Company.md)
 - [Contact](docs/Contact.md)
 - [Division](docs/Division.md)
 - [Dnsname](docs/Dnsname.md)
 - [Domain](docs/Domain.md)
 - [Domainuser](docs/Domainuser.md)
 - [Headline](docs/Headline.md)
 - [Ip](docs/Ip.md)
 - [Location](docs/Location.md)
 - [Note](docs/Note.md)
 - [Notestatus](docs/Notestatus.md)
 - [Os](docs/Os.md)
 - [Osarch](docs/Osarch.md)
 - [Reason](docs/Reason.md)
 - [Recommendation](docs/Recommendation.md)
 - [Reportitem](docs/Reportitem.md)
 - [Serviceprovider](docs/Serviceprovider.md)
 - [System](docs/System.md)
 - [Systemstatus](docs/Systemstatus.md)
 - [Systemtype](docs/Systemtype.md)
 - [Systemuser](docs/Systemuser.md)
 - [Tag](docs/Tag.md)
 - [Tagcolor](docs/Tagcolor.md)
 - [Task](docs/Task.md)
 - [Taskname](docs/Taskname.md)
 - [Taskpriority](docs/Taskpriority.md)
 - [Taskstatus](docs/Taskstatus.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



