package mail

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

// PostCharactersCharacterIDMailLabelsReader is a Reader for the PostCharactersCharacterIDMailLabels structure.
type PostCharactersCharacterIDMailLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDMailLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCharactersCharacterIDMailLabelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostCharactersCharacterIDMailLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDMailLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDMailLabelsCreated creates a PostCharactersCharacterIDMailLabelsCreated with default headers values
func NewPostCharactersCharacterIDMailLabelsCreated() *PostCharactersCharacterIDMailLabelsCreated {
	return &PostCharactersCharacterIDMailLabelsCreated{}
}

/*PostCharactersCharacterIDMailLabelsCreated handles this case with default header values.

Label created
*/
type PostCharactersCharacterIDMailLabelsCreated struct {
	Payload int64
}

func (o *PostCharactersCharacterIDMailLabelsCreated) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/labels/][%d] postCharactersCharacterIdMailLabelsCreated  %+v", 201, o.Payload)
}

func (o *PostCharactersCharacterIDMailLabelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailLabelsForbidden creates a PostCharactersCharacterIDMailLabelsForbidden with default headers values
func NewPostCharactersCharacterIDMailLabelsForbidden() *PostCharactersCharacterIDMailLabelsForbidden {
	return &PostCharactersCharacterIDMailLabelsForbidden{}
}

/*PostCharactersCharacterIDMailLabelsForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDMailLabelsForbidden struct {
	Payload PostCharactersCharacterIDMailLabelsForbiddenBody
}

func (o *PostCharactersCharacterIDMailLabelsForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/labels/][%d] postCharactersCharacterIdMailLabelsForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDMailLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailLabelsInternalServerError creates a PostCharactersCharacterIDMailLabelsInternalServerError with default headers values
func NewPostCharactersCharacterIDMailLabelsInternalServerError() *PostCharactersCharacterIDMailLabelsInternalServerError {
	return &PostCharactersCharacterIDMailLabelsInternalServerError{}
}

/*PostCharactersCharacterIDMailLabelsInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDMailLabelsInternalServerError struct {
	Payload PostCharactersCharacterIDMailLabelsInternalServerErrorBody
}

func (o *PostCharactersCharacterIDMailLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/labels/][%d] postCharactersCharacterIdMailLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDMailLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCharactersCharacterIDMailLabelsBody post_characters_character_id_mail_labels_label
//
// label object
swagger:model PostCharactersCharacterIDMailLabelsBody
*/
type PostCharactersCharacterIDMailLabelsBody struct {

	// post_characters_character_id_mail_labels_color
	//
	// Hexadecimal string representing label color,
	// in RGB format
	//
	Color *string `json:"color,omitempty"`

	// post_characters_character_id_mail_labels_name
	//
	// name string
	// Required: true
	// Max Length: 40
	// Min Length: 1
	Name *string `json:"name"`
}

/*PostCharactersCharacterIDMailLabelsForbiddenBody post_characters_character_id_mail_labels_forbidden
//
// Forbidden
swagger:model PostCharactersCharacterIDMailLabelsForbiddenBody
*/
type PostCharactersCharacterIDMailLabelsForbiddenBody struct {

	// post_characters_character_id_mail_labels_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post characters character ID mail labels forbidden body
func (o *PostCharactersCharacterIDMailLabelsForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCharactersCharacterIDMailLabelsForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postCharactersCharacterIdMailLabelsForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*PostCharactersCharacterIDMailLabelsInternalServerErrorBody post_characters_character_id_mail_labels_internal_server_error
//
// Internal server error
swagger:model PostCharactersCharacterIDMailLabelsInternalServerErrorBody
*/
type PostCharactersCharacterIDMailLabelsInternalServerErrorBody struct {

	// post_characters_character_id_mail_labels_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this post characters character ID mail labels internal server error body
func (o *PostCharactersCharacterIDMailLabelsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *PostCharactersCharacterIDMailLabelsInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postCharactersCharacterIdMailLabelsInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}