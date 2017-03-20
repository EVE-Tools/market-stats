package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCharactersCharacterIDMailReader is a Reader for the GetCharactersCharacterIDMail structure.
type GetCharactersCharacterIDMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDMailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailOK creates a GetCharactersCharacterIDMailOK with default headers values
func NewGetCharactersCharacterIDMailOK() *GetCharactersCharacterIDMailOK {
	return &GetCharactersCharacterIDMailOK{}
}

/*GetCharactersCharacterIDMailOK handles this case with default header values.

The requested mail
*/
type GetCharactersCharacterIDMailOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDMailOKBodyItems0
}

func (o *GetCharactersCharacterIDMailOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailForbidden creates a GetCharactersCharacterIDMailForbidden with default headers values
func NewGetCharactersCharacterIDMailForbidden() *GetCharactersCharacterIDMailForbidden {
	return &GetCharactersCharacterIDMailForbidden{}
}

/*GetCharactersCharacterIDMailForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailForbidden struct {
	Payload GetCharactersCharacterIDMailForbiddenBody
}

func (o *GetCharactersCharacterIDMailForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailInternalServerError creates a GetCharactersCharacterIDMailInternalServerError with default headers values
func NewGetCharactersCharacterIDMailInternalServerError() *GetCharactersCharacterIDMailInternalServerError {
	return &GetCharactersCharacterIDMailInternalServerError{}
}

/*GetCharactersCharacterIDMailInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailInternalServerError struct {
	Payload GetCharactersCharacterIDMailInternalServerErrorBody
}

func (o *GetCharactersCharacterIDMailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailForbiddenBody get_characters_character_id_mail_forbidden
//
// Forbidden
swagger:model GetCharactersCharacterIDMailForbiddenBody
*/
type GetCharactersCharacterIDMailForbiddenBody struct {

	// get_characters_character_id_mail_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID mail forbidden body
func (o *GetCharactersCharacterIDMailForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDMailForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailInternalServerErrorBody get_characters_character_id_mail_internal_server_error
//
// Internal server error
swagger:model GetCharactersCharacterIDMailInternalServerErrorBody
*/
type GetCharactersCharacterIDMailInternalServerErrorBody struct {

	// get_characters_character_id_mail_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID mail internal server error body
func (o *GetCharactersCharacterIDMailInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDMailInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdMailInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDMailOKBodyItems0 get_characters_character_id_mail_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDMailOKBodyItems0
*/
type GetCharactersCharacterIDMailOKBodyItems0 struct {

	// get_characters_character_id_mail_from
	//
	// From whom the mail was sent
	From int32 `json:"from,omitempty"`

	// get_characters_character_id_mail_is_read
	//
	// is_read boolean
	IsRead bool `json:"is_read,omitempty"`

	// get_characters_character_id_mail_labels
	//
	// labels array
	// Minimum: 0
	// Max Items: 25
	// Unique: true
	Labels []int64 `json:"labels"`

	// get_characters_character_id_mail_mail_id
	//
	// mail_id integer
	MailID int64 `json:"mail_id,omitempty"`

	// get_characters_character_id_mail_recipients
	//
	// Recipients of the mail
	// Max Items: 50
	// Min Items: 1
	// Unique: true
	Recipients []*GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0 `json:"recipients"`

	// get_characters_character_id_mail_subject
	//
	// Mail subject
	Subject string `json:"subject,omitempty"`

	// get_characters_character_id_mail_timestamp
	//
	// When the mail was sent
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this get characters character ID mail o k body items0
func (o *GetCharactersCharacterIDMailOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecipients(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDMailOKBodyItems0) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	iLabelsSize := int64(len(o.Labels))

	if err := validate.MaxItems("labels", "body", iLabelsSize, 25); err != nil {
		return err
	}

	if err := validate.UniqueItems("labels", "body", o.Labels); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDMailOKBodyItems0) validateRecipients(formats strfmt.Registry) error {

	if swag.IsZero(o.Recipients) { // not required
		return nil
	}

	iRecipientsSize := int64(len(o.Recipients))

	if err := validate.MinItems("recipients", "body", iRecipientsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("recipients", "body", iRecipientsSize, 50); err != nil {
		return err
	}

	if err := validate.UniqueItems("recipients", "body", o.Recipients); err != nil {
		return err
	}

	for i := 0; i < len(o.Recipients); i++ {

		if swag.IsZero(o.Recipients[i]) { // not required
			continue
		}

		if o.Recipients[i] != nil {

			if err := o.Recipients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

/*GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0 get_characters_character_id_mail_recipient
//
// recipient object
swagger:model GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0
*/
type GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0 struct {

	// get_characters_character_id_mail_recipient_id
	//
	// recipient_id integer
	// Required: true
	RecipientID *int32 `json:"recipient_id"`

	// get_characters_character_id_mail_recipient_type
	//
	// recipient_type string
	// Required: true
	RecipientType *string `json:"recipient_type"`
}

// Validate validates this get characters character ID mail o k body items0 recipients items0
func (o *GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecipientID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecipientType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0) validateRecipientID(formats strfmt.Registry) error {

	if err := validate.Required("recipient_id", "body", o.RecipientID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdMailOKBodyItems0RecipientsItems0TypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance","character","corporation","mailing_list"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdMailOKBodyItems0RecipientsItems0TypeRecipientTypePropEnum = append(getCharactersCharacterIdMailOKBodyItems0RecipientsItems0TypeRecipientTypePropEnum, v)
	}
}

const (
	// GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeAlliance captures enum value "alliance"
	GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeAlliance string = "alliance"
	// GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeCharacter captures enum value "character"
	GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeCharacter string = "character"
	// GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeCorporation captures enum value "corporation"
	GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeCorporation string = "corporation"
	// GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeMailingList captures enum value "mailing_list"
	GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0RecipientTypeMailingList string = "mailing_list"
)

// prop value enum
func (o *GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdMailOKBodyItems0RecipientsItems0TypeRecipientTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDMailOKBodyItems0RecipientsItems0) validateRecipientType(formats strfmt.Registry) error {

	if err := validate.Required("recipient_type", "body", o.RecipientType); err != nil {
		return err
	}

	// value enum
	if err := o.validateRecipientTypeEnum("recipient_type", "body", *o.RecipientType); err != nil {
		return err
	}

	return nil
}