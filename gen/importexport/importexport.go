// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package importexport provides a client for AWS Import/Export.
package importexport

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// ImportExport is a client for AWS Import/Export.
type ImportExport struct {
	client *aws.QueryClient
}

// New returns a new ImportExport client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *ImportExport {
	if client == nil {
		client = http.DefaultClient
	}

	service := "importexport"
	endpoint, service, region := endpoints.Lookup("importexport", region)

	return &ImportExport{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2010-06-01",
		},
	}
}

// CancelJob this operation cancels a specified job. Only the job owner can
// cancel it. The operation fails if the job has already started or is
// complete.
func (c *ImportExport) CancelJob(req *CancelJobInput) (resp *CancelJobResult, err error) {
	resp = &CancelJobResult{}
	err = c.client.Do("CancelJob", "POST", "/?Operation=CancelJob", req, resp)
	return
}

// CreateJob this operation initiates the process of scheduling an upload
// or download of your data. You include in the request a manifest that
// describes the data transfer specifics. The response to the request
// includes a job ID, which you can use in other operations, a signature
// that you use to identify your storage device, and the address where you
// should ship your storage device.
func (c *ImportExport) CreateJob(req *CreateJobInput) (resp *CreateJobResult, err error) {
	resp = &CreateJobResult{}
	err = c.client.Do("CreateJob", "POST", "/?Operation=CreateJob", req, resp)
	return
}

// GetStatus this operation returns information about a job, including
// where the job is in the processing pipeline, the status of the results,
// and the signature value associated with the job. You can only return
// information about jobs you own.
func (c *ImportExport) GetStatus(req *GetStatusInput) (resp *GetStatusResult, err error) {
	resp = &GetStatusResult{}
	err = c.client.Do("GetStatus", "POST", "/?Operation=GetStatus", req, resp)
	return
}

// ListJobs this operation returns the jobs associated with the requester.
// AWS Import/Export lists the jobs in reverse chronological order based on
// the date of creation. For example if Job Test1 was created 2009Dec30 and
// Test2 was created 2010Feb05, the ListJobs operation would return Test2
// followed by Test1.
func (c *ImportExport) ListJobs(req *ListJobsInput) (resp *ListJobsResult, err error) {
	resp = &ListJobsResult{}
	err = c.client.Do("ListJobs", "POST", "/?Operation=ListJobs", req, resp)
	return
}

// UpdateJob you use this operation to change the parameters specified in
// the original manifest file by supplying a new manifest file. The
// manifest file attached to this request replaces the original manifest
// file. You can only use the operation after a CreateJob request but
// before the data transfer starts and you can only use it on jobs you own.
func (c *ImportExport) UpdateJob(req *UpdateJobInput) (resp *UpdateJobResult, err error) {
	resp = &UpdateJobResult{}
	err = c.client.Do("UpdateJob", "POST", "/?Operation=UpdateJob", req, resp)
	return
}

// CancelJobInput is undocumented.
type CancelJobInput struct {
	JobID aws.StringValue `xml:"JobId"`
}

// CancelJobOutput is undocumented.
type CancelJobOutput struct {
	Success aws.BooleanValue `xml:"CancelJobResult>Success"`
}

// CreateJobInput is undocumented.
type CreateJobInput struct {
	JobType          aws.StringValue  `xml:"JobType"`
	Manifest         aws.StringValue  `xml:"Manifest"`
	ManifestAddendum aws.StringValue  `xml:"ManifestAddendum"`
	ValidateOnly     aws.BooleanValue `xml:"ValidateOnly"`
}

// CreateJobOutput is undocumented.
type CreateJobOutput struct {
	AWSShippingAddress    aws.StringValue `xml:"CreateJobResult>AwsShippingAddress"`
	JobID                 aws.StringValue `xml:"CreateJobResult>JobId"`
	JobType               aws.StringValue `xml:"CreateJobResult>JobType"`
	Signature             aws.StringValue `xml:"CreateJobResult>Signature"`
	SignatureFileContents aws.StringValue `xml:"CreateJobResult>SignatureFileContents"`
	WarningMessage        aws.StringValue `xml:"CreateJobResult>WarningMessage"`
}

// GetStatusInput is undocumented.
type GetStatusInput struct {
	JobID aws.StringValue `xml:"JobId"`
}

// GetStatusOutput is undocumented.
type GetStatusOutput struct {
	AWSShippingAddress    aws.StringValue  `xml:"GetStatusResult>AwsShippingAddress"`
	Carrier               aws.StringValue  `xml:"GetStatusResult>Carrier"`
	CreationDate          time.Time        `xml:"GetStatusResult>CreationDate"`
	CurrentManifest       aws.StringValue  `xml:"GetStatusResult>CurrentManifest"`
	ErrorCount            aws.IntegerValue `xml:"GetStatusResult>ErrorCount"`
	JobID                 aws.StringValue  `xml:"GetStatusResult>JobId"`
	JobType               aws.StringValue  `xml:"GetStatusResult>JobType"`
	LocationCode          aws.StringValue  `xml:"GetStatusResult>LocationCode"`
	LocationMessage       aws.StringValue  `xml:"GetStatusResult>LocationMessage"`
	LogBucket             aws.StringValue  `xml:"GetStatusResult>LogBucket"`
	LogKey                aws.StringValue  `xml:"GetStatusResult>LogKey"`
	ProgressCode          aws.StringValue  `xml:"GetStatusResult>ProgressCode"`
	ProgressMessage       aws.StringValue  `xml:"GetStatusResult>ProgressMessage"`
	Signature             aws.StringValue  `xml:"GetStatusResult>Signature"`
	SignatureFileContents aws.StringValue  `xml:"GetStatusResult>SignatureFileContents"`
	TrackingNumber        aws.StringValue  `xml:"GetStatusResult>TrackingNumber"`
}

// Job is undocumented.
type Job struct {
	CreationDate time.Time        `xml:"CreationDate"`
	IsCanceled   aws.BooleanValue `xml:"IsCanceled"`
	JobID        aws.StringValue  `xml:"JobId"`
	JobType      aws.StringValue  `xml:"JobType"`
}

// Possible values for ImportExport.
const (
	JobTypeExport = "Export"
	JobTypeImport = "Import"
)

// ListJobsInput is undocumented.
type ListJobsInput struct {
	Marker  aws.StringValue  `xml:"Marker"`
	MaxJobs aws.IntegerValue `xml:"MaxJobs"`
}

// ListJobsOutput is undocumented.
type ListJobsOutput struct {
	IsTruncated aws.BooleanValue `xml:"ListJobsResult>IsTruncated"`
	Jobs        []Job            `xml:"ListJobsResult>Jobs>member"`
}

// UpdateJobInput is undocumented.
type UpdateJobInput struct {
	JobID        aws.StringValue  `xml:"JobId"`
	JobType      aws.StringValue  `xml:"JobType"`
	Manifest     aws.StringValue  `xml:"Manifest"`
	ValidateOnly aws.BooleanValue `xml:"ValidateOnly"`
}

// UpdateJobOutput is undocumented.
type UpdateJobOutput struct {
	Success        aws.BooleanValue `xml:"UpdateJobResult>Success"`
	WarningMessage aws.StringValue  `xml:"UpdateJobResult>WarningMessage"`
}

// CancelJobResult is a wrapper for CancelJobOutput.
type CancelJobResult struct {
	Success aws.BooleanValue `xml:"CancelJobResult>Success"`
}

// CreateJobResult is a wrapper for CreateJobOutput.
type CreateJobResult struct {
	AWSShippingAddress    aws.StringValue `xml:"CreateJobResult>AwsShippingAddress"`
	JobID                 aws.StringValue `xml:"CreateJobResult>JobId"`
	JobType               aws.StringValue `xml:"CreateJobResult>JobType"`
	Signature             aws.StringValue `xml:"CreateJobResult>Signature"`
	SignatureFileContents aws.StringValue `xml:"CreateJobResult>SignatureFileContents"`
	WarningMessage        aws.StringValue `xml:"CreateJobResult>WarningMessage"`
}

// GetStatusResult is a wrapper for GetStatusOutput.
type GetStatusResult struct {
	AWSShippingAddress    aws.StringValue  `xml:"GetStatusResult>AwsShippingAddress"`
	Carrier               aws.StringValue  `xml:"GetStatusResult>Carrier"`
	CreationDate          time.Time        `xml:"GetStatusResult>CreationDate"`
	CurrentManifest       aws.StringValue  `xml:"GetStatusResult>CurrentManifest"`
	ErrorCount            aws.IntegerValue `xml:"GetStatusResult>ErrorCount"`
	JobID                 aws.StringValue  `xml:"GetStatusResult>JobId"`
	JobType               aws.StringValue  `xml:"GetStatusResult>JobType"`
	LocationCode          aws.StringValue  `xml:"GetStatusResult>LocationCode"`
	LocationMessage       aws.StringValue  `xml:"GetStatusResult>LocationMessage"`
	LogBucket             aws.StringValue  `xml:"GetStatusResult>LogBucket"`
	LogKey                aws.StringValue  `xml:"GetStatusResult>LogKey"`
	ProgressCode          aws.StringValue  `xml:"GetStatusResult>ProgressCode"`
	ProgressMessage       aws.StringValue  `xml:"GetStatusResult>ProgressMessage"`
	Signature             aws.StringValue  `xml:"GetStatusResult>Signature"`
	SignatureFileContents aws.StringValue  `xml:"GetStatusResult>SignatureFileContents"`
	TrackingNumber        aws.StringValue  `xml:"GetStatusResult>TrackingNumber"`
}

// ListJobsResult is a wrapper for ListJobsOutput.
type ListJobsResult struct {
	IsTruncated aws.BooleanValue `xml:"ListJobsResult>IsTruncated"`
	Jobs        []Job            `xml:"ListJobsResult>Jobs>member"`
}

// UpdateJobResult is a wrapper for UpdateJobOutput.
type UpdateJobResult struct {
	Success        aws.BooleanValue `xml:"UpdateJobResult>Success"`
	WarningMessage aws.StringValue  `xml:"UpdateJobResult>WarningMessage"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
