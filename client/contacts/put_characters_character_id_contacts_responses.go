package contacts

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

// PutCharactersCharacterIDContactsReader is a Reader for the PutCharactersCharacterIDContacts structure.
type PutCharactersCharacterIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCharactersCharacterIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutCharactersCharacterIDContactsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPutCharactersCharacterIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutCharactersCharacterIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCharactersCharacterIDContactsNoContent creates a PutCharactersCharacterIDContactsNoContent with default headers values
func NewPutCharactersCharacterIDContactsNoContent() *PutCharactersCharacterIDContactsNoContent {
	return &PutCharactersCharacterIDContactsNoContent{}
}

/*PutCharactersCharacterIDContactsNoContent handles this case with default header values.

Contacts updated
*/
type PutCharactersCharacterIDContactsNoContent struct {
}

func (o *PutCharactersCharacterIDContactsNoContent) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/contacts/][%d] putCharactersCharacterIdContactsNoContent ", 204)
}

func (o *PutCharactersCharacterIDContactsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCharactersCharacterIDContactsForbidden creates a PutCharactersCharacterIDContactsForbidden with default headers values
func NewPutCharactersCharacterIDContactsForbidden() *PutCharactersCharacterIDContactsForbidden {
	return &PutCharactersCharacterIDContactsForbidden{}
}

/*PutCharactersCharacterIDContactsForbidden handles this case with default header values.

Forbidden
*/
type PutCharactersCharacterIDContactsForbidden struct {
	Payload PutCharactersCharacterIDContactsForbiddenBody
}

func (o *PutCharactersCharacterIDContactsForbidden) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/contacts/][%d] putCharactersCharacterIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *PutCharactersCharacterIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDContactsInternalServerError creates a PutCharactersCharacterIDContactsInternalServerError with default headers values
func NewPutCharactersCharacterIDContactsInternalServerError() *PutCharactersCharacterIDContactsInternalServerError {
	return &PutCharactersCharacterIDContactsInternalServerError{}
}

/*PutCharactersCharacterIDContactsInternalServerError handles this case with default header values.

Internal server error
*/
type PutCharactersCharacterIDContactsInternalServerError struct {
	Payload PutCharactersCharacterIDContactsInternalServerErrorBody
}

func (o *PutCharactersCharacterIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /characters/{character_id}/contacts/][%d] putCharactersCharacterIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCharactersCharacterIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutCharactersCharacterIDContactsForbiddenBody put_characters_character_id_contacts_forbidden
//
// Forbidden
swagger:model PutCharactersCharacterIDContactsForbiddenBody
*/
type PutCharactersCharacterIDContactsForbiddenBody struct {

	// put_characters_character_id_contacts_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this put characters character ID contacts forbidden body
func (o *PutCharactersCharacterIDContactsForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCharactersCharacterIDContactsForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putCharactersCharacterIdContactsForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PutCharactersCharacterIDContactsInternalServerErrorBody put_characters_character_id_contacts_internal_server_error
//
// Internal server error
swagger:model PutCharactersCharacterIDContactsInternalServerErrorBody
*/
type PutCharactersCharacterIDContactsInternalServerErrorBody struct {

	// put_characters_character_id_contacts_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this put characters character ID contacts internal server error body
func (o *PutCharactersCharacterIDContactsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *PutCharactersCharacterIDContactsInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("putCharactersCharacterIdContactsInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}