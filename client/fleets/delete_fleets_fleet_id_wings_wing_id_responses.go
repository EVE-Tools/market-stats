package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFleetsFleetIDWingsWingIDReader is a Reader for the DeleteFleetsFleetIDWingsWingID structure.
type DeleteFleetsFleetIDWingsWingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetsFleetIDWingsWingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFleetsFleetIDWingsWingIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteFleetsFleetIDWingsWingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFleetsFleetIDWingsWingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteFleetsFleetIDWingsWingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFleetsFleetIDWingsWingIDNoContent creates a DeleteFleetsFleetIDWingsWingIDNoContent with default headers values
func NewDeleteFleetsFleetIDWingsWingIDNoContent() *DeleteFleetsFleetIDWingsWingIDNoContent {
	return &DeleteFleetsFleetIDWingsWingIDNoContent{}
}

/*DeleteFleetsFleetIDWingsWingIDNoContent handles this case with default header values.

Wing deleted
*/
type DeleteFleetsFleetIDWingsWingIDNoContent struct {
}

func (o *DeleteFleetsFleetIDWingsWingIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDWingsWingIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDForbidden creates a DeleteFleetsFleetIDWingsWingIDForbidden with default headers values
func NewDeleteFleetsFleetIDWingsWingIDForbidden() *DeleteFleetsFleetIDWingsWingIDForbidden {
	return &DeleteFleetsFleetIDWingsWingIDForbidden{}
}

/*DeleteFleetsFleetIDWingsWingIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteFleetsFleetIDWingsWingIDForbidden struct {
	Payload DeleteFleetsFleetIDWingsWingIDForbiddenBody
}

func (o *DeleteFleetsFleetIDWingsWingIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDNotFound creates a DeleteFleetsFleetIDWingsWingIDNotFound with default headers values
func NewDeleteFleetsFleetIDWingsWingIDNotFound() *DeleteFleetsFleetIDWingsWingIDNotFound {
	return &DeleteFleetsFleetIDWingsWingIDNotFound{}
}

/*DeleteFleetsFleetIDWingsWingIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type DeleteFleetsFleetIDWingsWingIDNotFound struct {
	Payload DeleteFleetsFleetIDWingsWingIDNotFoundBody
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDInternalServerError creates a DeleteFleetsFleetIDWingsWingIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDWingsWingIDInternalServerError() *DeleteFleetsFleetIDWingsWingIDInternalServerError {
	return &DeleteFleetsFleetIDWingsWingIDInternalServerError{}
}

/*DeleteFleetsFleetIDWingsWingIDInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteFleetsFleetIDWingsWingIDInternalServerError struct {
	Payload DeleteFleetsFleetIDWingsWingIDInternalServerErrorBody
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDWingsWingIDForbiddenBody delete_fleets_fleet_id_wings_wing_id_forbidden
//
// Forbidden
swagger:model DeleteFleetsFleetIDWingsWingIDForbiddenBody
*/
type DeleteFleetsFleetIDWingsWingIDForbiddenBody struct {

	// delete_fleets_fleet_id_wings_wing_id_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this delete fleets fleet ID wings wing ID forbidden body
func (o *DeleteFleetsFleetIDWingsWingIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteFleetsFleetIDWingsWingIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteFleetsFleetIdWingsWingIdForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDWingsWingIDInternalServerErrorBody delete_fleets_fleet_id_wings_wing_id_internal_server_error
//
// Internal server error
swagger:model DeleteFleetsFleetIDWingsWingIDInternalServerErrorBody
*/
type DeleteFleetsFleetIDWingsWingIDInternalServerErrorBody struct {

	// delete_fleets_fleet_id_wings_wing_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this delete fleets fleet ID wings wing ID internal server error body
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteFleetsFleetIdWingsWingIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDWingsWingIDNotFoundBody delete_fleets_fleet_id_wings_wing_id_not_found
//
// Not found
swagger:model DeleteFleetsFleetIDWingsWingIDNotFoundBody
*/
type DeleteFleetsFleetIDWingsWingIDNotFoundBody struct {

	// delete_fleets_fleet_id_wings_wing_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this delete fleets fleet ID wings wing ID not found body
func (o *DeleteFleetsFleetIDWingsWingIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteFleetsFleetIdWingsWingIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}