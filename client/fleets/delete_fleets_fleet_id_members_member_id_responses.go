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

// DeleteFleetsFleetIDMembersMemberIDReader is a Reader for the DeleteFleetsFleetIDMembersMemberID structure.
type DeleteFleetsFleetIDMembersMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetsFleetIDMembersMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFleetsFleetIDMembersMemberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteFleetsFleetIDMembersMemberIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFleetsFleetIDMembersMemberIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteFleetsFleetIDMembersMemberIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFleetsFleetIDMembersMemberIDNoContent creates a DeleteFleetsFleetIDMembersMemberIDNoContent with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDNoContent() *DeleteFleetsFleetIDMembersMemberIDNoContent {
	return &DeleteFleetsFleetIDMembersMemberIDNoContent{}
}

/*DeleteFleetsFleetIDMembersMemberIDNoContent handles this case with default header values.

Fleet member kicked
*/
type DeleteFleetsFleetIDMembersMemberIDNoContent struct {
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDForbidden creates a DeleteFleetsFleetIDMembersMemberIDForbidden with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDForbidden() *DeleteFleetsFleetIDMembersMemberIDForbidden {
	return &DeleteFleetsFleetIDMembersMemberIDForbidden{}
}

/*DeleteFleetsFleetIDMembersMemberIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteFleetsFleetIDMembersMemberIDForbidden struct {
	Payload DeleteFleetsFleetIDMembersMemberIDForbiddenBody
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDNotFound creates a DeleteFleetsFleetIDMembersMemberIDNotFound with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDNotFound() *DeleteFleetsFleetIDMembersMemberIDNotFound {
	return &DeleteFleetsFleetIDMembersMemberIDNotFound{}
}

/*DeleteFleetsFleetIDMembersMemberIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type DeleteFleetsFleetIDMembersMemberIDNotFound struct {
	Payload DeleteFleetsFleetIDMembersMemberIDNotFoundBody
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDInternalServerError creates a DeleteFleetsFleetIDMembersMemberIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDInternalServerError() *DeleteFleetsFleetIDMembersMemberIDInternalServerError {
	return &DeleteFleetsFleetIDMembersMemberIDInternalServerError{}
}

/*DeleteFleetsFleetIDMembersMemberIDInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteFleetsFleetIDMembersMemberIDInternalServerError struct {
	Payload DeleteFleetsFleetIDMembersMemberIDInternalServerErrorBody
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDMembersMemberIDForbiddenBody delete_fleets_fleet_id_members_member_id_forbidden
//
// Forbidden
swagger:model DeleteFleetsFleetIDMembersMemberIDForbiddenBody
*/
type DeleteFleetsFleetIDMembersMemberIDForbiddenBody struct {

	// delete_fleets_fleet_id_members_member_id_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this delete fleets fleet ID members member ID forbidden body
func (o *DeleteFleetsFleetIDMembersMemberIDForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDMembersMemberIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteFleetsFleetIdMembersMemberIdForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDMembersMemberIDInternalServerErrorBody delete_fleets_fleet_id_members_member_id_internal_server_error
//
// Internal server error
swagger:model DeleteFleetsFleetIDMembersMemberIDInternalServerErrorBody
*/
type DeleteFleetsFleetIDMembersMemberIDInternalServerErrorBody struct {

	// delete_fleets_fleet_id_members_member_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this delete fleets fleet ID members member ID internal server error body
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteFleetsFleetIdMembersMemberIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*DeleteFleetsFleetIDMembersMemberIDNotFoundBody delete_fleets_fleet_id_members_member_id_not_found
//
// Not found
swagger:model DeleteFleetsFleetIDMembersMemberIDNotFoundBody
*/
type DeleteFleetsFleetIDMembersMemberIDNotFoundBody struct {

	// delete_fleets_fleet_id_members_member_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this delete fleets fleet ID members member ID not found body
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("deleteFleetsFleetIdMembersMemberIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}