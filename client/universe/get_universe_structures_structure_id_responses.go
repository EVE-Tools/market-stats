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

// GetUniverseStructuresStructureIDReader is a Reader for the GetUniverseStructuresStructureID structure.
type GetUniverseStructuresStructureIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStructuresStructureIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseStructuresStructureIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetUniverseStructuresStructureIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUniverseStructuresStructureIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseStructuresStructureIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseStructuresStructureIDOK creates a GetUniverseStructuresStructureIDOK with default headers values
func NewGetUniverseStructuresStructureIDOK() *GetUniverseStructuresStructureIDOK {
	return &GetUniverseStructuresStructureIDOK{}
}

/*GetUniverseStructuresStructureIDOK handles this case with default header values.

Data about a structure
*/
type GetUniverseStructuresStructureIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetUniverseStructuresStructureIDOKBody
}

func (o *GetUniverseStructuresStructureIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresStructureIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresStructureIDForbidden creates a GetUniverseStructuresStructureIDForbidden with default headers values
func NewGetUniverseStructuresStructureIDForbidden() *GetUniverseStructuresStructureIDForbidden {
	return &GetUniverseStructuresStructureIDForbidden{}
}

/*GetUniverseStructuresStructureIDForbidden handles this case with default header values.

Forbidden
*/
type GetUniverseStructuresStructureIDForbidden struct {
	Payload GetUniverseStructuresStructureIDForbiddenBody
}

func (o *GetUniverseStructuresStructureIDForbidden) Error() string {
	return fmt.Sprintf("[GET /universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdForbidden  %+v", 403, o.Payload)
}

func (o *GetUniverseStructuresStructureIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDNotFound creates a GetUniverseStructuresStructureIDNotFound with default headers values
func NewGetUniverseStructuresStructureIDNotFound() *GetUniverseStructuresStructureIDNotFound {
	return &GetUniverseStructuresStructureIDNotFound{}
}

/*GetUniverseStructuresStructureIDNotFound handles this case with default header values.

Structure not found
*/
type GetUniverseStructuresStructureIDNotFound struct {
	Payload GetUniverseStructuresStructureIDNotFoundBody
}

func (o *GetUniverseStructuresStructureIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseStructuresStructureIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDInternalServerError creates a GetUniverseStructuresStructureIDInternalServerError with default headers values
func NewGetUniverseStructuresStructureIDInternalServerError() *GetUniverseStructuresStructureIDInternalServerError {
	return &GetUniverseStructuresStructureIDInternalServerError{}
}

/*GetUniverseStructuresStructureIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseStructuresStructureIDInternalServerError struct {
	Payload GetUniverseStructuresStructureIDInternalServerErrorBody
}

func (o *GetUniverseStructuresStructureIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresStructureIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseStructuresStructureIDForbiddenBody get_universe_structures_structure_id_forbidden
//
// Forbidden
swagger:model GetUniverseStructuresStructureIDForbiddenBody
*/
type GetUniverseStructuresStructureIDForbiddenBody struct {

	// get_universe_structures_structure_id_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get universe structures structure ID forbidden body
func (o *GetUniverseStructuresStructureIDForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseStructuresStructureIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetUniverseStructuresStructureIDInternalServerErrorBody get_universe_structures_structure_id_internal_server_error
//
// Internal server error
swagger:model GetUniverseStructuresStructureIDInternalServerErrorBody
*/
type GetUniverseStructuresStructureIDInternalServerErrorBody struct {

	// get_universe_structures_structure_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get universe structures structure ID internal server error body
func (o *GetUniverseStructuresStructureIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseStructuresStructureIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetUniverseStructuresStructureIDNotFoundBody get_universe_structures_structure_id_not_found
//
// Not found
swagger:model GetUniverseStructuresStructureIDNotFoundBody
*/
type GetUniverseStructuresStructureIDNotFoundBody struct {

	// get_universe_structures_structure_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get universe structures structure ID not found body
func (o *GetUniverseStructuresStructureIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseStructuresStructureIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetUniverseStructuresStructureIDOKBody get_universe_structures_structure_id_ok
//
// 200 ok object
swagger:model GetUniverseStructuresStructureIDOKBody
*/
type GetUniverseStructuresStructureIDOKBody struct {

	// get_universe_structures_structure_id_name
	//
	// The full name of the structure
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *GetUniverseStructuresStructureIDOKBodyPosition `json:"position"`

	// get_universe_structures_structure_id_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_universe_structures_structure_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get universe structures structure ID o k body
func (o *GetUniverseStructuresStructureIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {

		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStructuresStructureIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

/*GetUniverseStructuresStructureIDOKBodyPosition get_universe_structures_structure_id_position
//
// Coordinates of the structure in Cartesian space relative to the Sun, in metres.
//
swagger:model GetUniverseStructuresStructureIDOKBodyPosition
*/
type GetUniverseStructuresStructureIDOKBodyPosition struct {

	// get_universe_structures_structure_id_x
	//
	// x number
	// Required: true
	X *float32 `json:"x"`

	// get_universe_structures_structure_id_y
	//
	// y number
	// Required: true
	Y *float32 `json:"y"`

	// get_universe_structures_structure_id_z
	//
	// z number
	// Required: true
	Z *float32 `json:"z"`
}

// Validate validates this get universe structures structure ID o k body position
func (o *GetUniverseStructuresStructureIDOKBodyPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateY(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseStructuresStructureIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}