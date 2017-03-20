package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAlliancesAllianceIDParams creates a new GetAlliancesAllianceIDParams object
// with the default values initialized.
func NewGetAlliancesAllianceIDParams() *GetAlliancesAllianceIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlliancesAllianceIDParamsWithTimeout creates a new GetAlliancesAllianceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlliancesAllianceIDParamsWithTimeout(timeout time.Duration) *GetAlliancesAllianceIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetAlliancesAllianceIDParamsWithContext creates a new GetAlliancesAllianceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlliancesAllianceIDParamsWithContext(ctx context.Context) *GetAlliancesAllianceIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetAlliancesAllianceIDParamsWithHTTPClient creates a new GetAlliancesAllianceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAlliancesAllianceIDParamsWithHTTPClient(client *http.Client) *GetAlliancesAllianceIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetAlliancesAllianceIDParams contains all the parameters to send to the API endpoint
for the get alliances alliance id operation typically these are written to a http.Request
*/
type GetAlliancesAllianceIDParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*AllianceID
	  An Eve alliance ID

	*/
	AllianceID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithTimeout(timeout time.Duration) *GetAlliancesAllianceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithContext(ctx context.Context) *GetAlliancesAllianceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithHTTPClient(client *http.Client) *GetAlliancesAllianceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithXUserAgent(xUserAgent *string) *GetAlliancesAllianceIDParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithAllianceID adds the allianceID to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithAllianceID(allianceID int32) *GetAlliancesAllianceIDParams {
	o.SetAllianceID(allianceID)
	return o
}

// SetAllianceID adds the allianceId to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetAllianceID(allianceID int32) {
	o.AllianceID = allianceID
}

// WithDatasource adds the datasource to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithDatasource(datasource *string) *GetAlliancesAllianceIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithUserAgent adds the userAgent to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) WithUserAgent(userAgent *string) *GetAlliancesAllianceIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get alliances alliance id params
func (o *GetAlliancesAllianceIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlliancesAllianceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	// path param alliance_id
	if err := r.SetPathParam("alliance_id", swag.FormatInt32(o.AllianceID)); err != nil {
		return err
	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string
		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {
			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}

	}

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string
		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {
			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}