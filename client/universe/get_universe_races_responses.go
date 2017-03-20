package universe

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

// GetUniverseRacesReader is a Reader for the GetUniverseRaces structure.
type GetUniverseRacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseRacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseRacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetUniverseRacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseRacesOK creates a GetUniverseRacesOK with default headers values
func NewGetUniverseRacesOK() *GetUniverseRacesOK {
	return &GetUniverseRacesOK{}
}

/*GetUniverseRacesOK handles this case with default header values.

A list of character races
*/
type GetUniverseRacesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetUniverseRacesOKBodyItems0
}

func (o *GetUniverseRacesOK) Error() string {
	return fmt.Sprintf("[GET /universe/races/][%d] getUniverseRacesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseRacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

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

// NewGetUniverseRacesInternalServerError creates a GetUniverseRacesInternalServerError with default headers values
func NewGetUniverseRacesInternalServerError() *GetUniverseRacesInternalServerError {
	return &GetUniverseRacesInternalServerError{}
}

/*GetUniverseRacesInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseRacesInternalServerError struct {
	Payload GetUniverseRacesInternalServerErrorBody
}

func (o *GetUniverseRacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/races/][%d] getUniverseRacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseRacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseRacesInternalServerErrorBody get_universe_races_internal_server_error
//
// Internal server error
swagger:model GetUniverseRacesInternalServerErrorBody
*/
type GetUniverseRacesInternalServerErrorBody struct {

	// get_universe_races_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get universe races internal server error body
func (o *GetUniverseRacesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseRacesInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseRacesInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetUniverseRacesOKBodyItems0 get_universe_races_200_ok
//
// 200 ok object
swagger:model GetUniverseRacesOKBodyItems0
*/
type GetUniverseRacesOKBodyItems0 struct {

	// get_universe_races_alliance_id
	//
	// The alliance generally associated with this race
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_universe_races_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_races_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_races_race_id
	//
	// race_id integer
	// Required: true
	RaceID *int32 `json:"race_id"`
}

// Validate validates this get universe races o k body items0
func (o *GetUniverseRacesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRaceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance_id", "body", o.AllianceID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateRaceID(formats strfmt.Registry) error {

	if err := validate.Required("race_id", "body", o.RaceID); err != nil {
		return err
	}

	return nil
}