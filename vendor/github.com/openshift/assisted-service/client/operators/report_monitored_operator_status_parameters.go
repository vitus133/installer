// Code generated by go-swagger; DO NOT EDIT.

package operators

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

	"github.com/openshift/assisted-service/models"
)

// NewReportMonitoredOperatorStatusParams creates a new ReportMonitoredOperatorStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportMonitoredOperatorStatusParams() *ReportMonitoredOperatorStatusParams {
	return &ReportMonitoredOperatorStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportMonitoredOperatorStatusParamsWithTimeout creates a new ReportMonitoredOperatorStatusParams object
// with the ability to set a timeout on a request.
func NewReportMonitoredOperatorStatusParamsWithTimeout(timeout time.Duration) *ReportMonitoredOperatorStatusParams {
	return &ReportMonitoredOperatorStatusParams{
		timeout: timeout,
	}
}

// NewReportMonitoredOperatorStatusParamsWithContext creates a new ReportMonitoredOperatorStatusParams object
// with the ability to set a context for a request.
func NewReportMonitoredOperatorStatusParamsWithContext(ctx context.Context) *ReportMonitoredOperatorStatusParams {
	return &ReportMonitoredOperatorStatusParams{
		Context: ctx,
	}
}

// NewReportMonitoredOperatorStatusParamsWithHTTPClient creates a new ReportMonitoredOperatorStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportMonitoredOperatorStatusParamsWithHTTPClient(client *http.Client) *ReportMonitoredOperatorStatusParams {
	return &ReportMonitoredOperatorStatusParams{
		HTTPClient: client,
	}
}

/* ReportMonitoredOperatorStatusParams contains all the parameters to send to the API endpoint
   for the report monitored operator status operation.

   Typically these are written to a http.Request.
*/
type ReportMonitoredOperatorStatusParams struct {

	/* ClusterID.

	   The cluster whose operators are being monitored.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* ReportParams.

	   The operators monitor report.
	*/
	ReportParams *models.OperatorMonitorReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report monitored operator status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportMonitoredOperatorStatusParams) WithDefaults() *ReportMonitoredOperatorStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report monitored operator status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportMonitoredOperatorStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) WithTimeout(timeout time.Duration) *ReportMonitoredOperatorStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) WithContext(ctx context.Context) *ReportMonitoredOperatorStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) WithHTTPClient(client *http.Client) *ReportMonitoredOperatorStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) WithClusterID(clusterID strfmt.UUID) *ReportMonitoredOperatorStatusParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithReportParams adds the reportParams to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) WithReportParams(reportParams *models.OperatorMonitorReport) *ReportMonitoredOperatorStatusParams {
	o.SetReportParams(reportParams)
	return o
}

// SetReportParams adds the reportParams to the report monitored operator status params
func (o *ReportMonitoredOperatorStatusParams) SetReportParams(reportParams *models.OperatorMonitorReport) {
	o.ReportParams = reportParams
}

// WriteToRequest writes these params to a swagger request
func (o *ReportMonitoredOperatorStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}
	if o.ReportParams != nil {
		if err := r.SetBodyParam(o.ReportParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
