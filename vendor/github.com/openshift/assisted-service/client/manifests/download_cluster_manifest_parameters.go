// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDownloadClusterManifestParams creates a new DownloadClusterManifestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadClusterManifestParams() *DownloadClusterManifestParams {
	return &DownloadClusterManifestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadClusterManifestParamsWithTimeout creates a new DownloadClusterManifestParams object
// with the ability to set a timeout on a request.
func NewDownloadClusterManifestParamsWithTimeout(timeout time.Duration) *DownloadClusterManifestParams {
	return &DownloadClusterManifestParams{
		timeout: timeout,
	}
}

// NewDownloadClusterManifestParamsWithContext creates a new DownloadClusterManifestParams object
// with the ability to set a context for a request.
func NewDownloadClusterManifestParamsWithContext(ctx context.Context) *DownloadClusterManifestParams {
	return &DownloadClusterManifestParams{
		Context: ctx,
	}
}

// NewDownloadClusterManifestParamsWithHTTPClient creates a new DownloadClusterManifestParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadClusterManifestParamsWithHTTPClient(client *http.Client) *DownloadClusterManifestParams {
	return &DownloadClusterManifestParams{
		HTTPClient: client,
	}
}

/* DownloadClusterManifestParams contains all the parameters to send to the API endpoint
   for the download cluster manifest operation.

   Typically these are written to a http.Request.
*/
type DownloadClusterManifestParams struct {

	/* ClusterID.

	   The cluster whose manifest should be downloaded.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* FileName.

	   The manifest file name to download.
	*/
	FileName string

	/* Folder.

	   The folder that contains the files. Manifests can be placed in 'manifests' or 'openshift' directories.

	   Default: "manifests"
	*/
	Folder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download cluster manifest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadClusterManifestParams) WithDefaults() *DownloadClusterManifestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download cluster manifest params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadClusterManifestParams) SetDefaults() {
	var (
		folderDefault = string("manifests")
	)

	val := DownloadClusterManifestParams{
		Folder: &folderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the download cluster manifest params
func (o *DownloadClusterManifestParams) WithTimeout(timeout time.Duration) *DownloadClusterManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download cluster manifest params
func (o *DownloadClusterManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download cluster manifest params
func (o *DownloadClusterManifestParams) WithContext(ctx context.Context) *DownloadClusterManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download cluster manifest params
func (o *DownloadClusterManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download cluster manifest params
func (o *DownloadClusterManifestParams) WithHTTPClient(client *http.Client) *DownloadClusterManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download cluster manifest params
func (o *DownloadClusterManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the download cluster manifest params
func (o *DownloadClusterManifestParams) WithClusterID(clusterID strfmt.UUID) *DownloadClusterManifestParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the download cluster manifest params
func (o *DownloadClusterManifestParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithFileName adds the fileName to the download cluster manifest params
func (o *DownloadClusterManifestParams) WithFileName(fileName string) *DownloadClusterManifestParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the download cluster manifest params
func (o *DownloadClusterManifestParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithFolder adds the folder to the download cluster manifest params
func (o *DownloadClusterManifestParams) WithFolder(folder *string) *DownloadClusterManifestParams {
	o.SetFolder(folder)
	return o
}

// SetFolder adds the folder to the download cluster manifest params
func (o *DownloadClusterManifestParams) SetFolder(folder *string) {
	o.Folder = folder
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadClusterManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// query param file_name
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {

		if err := r.SetQueryParam("file_name", qFileName); err != nil {
			return err
		}
	}

	if o.Folder != nil {

		// query param folder
		var qrFolder string

		if o.Folder != nil {
			qrFolder = *o.Folder
		}
		qFolder := qrFolder
		if qFolder != "" {

			if err := r.SetQueryParam("folder", qFolder); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
