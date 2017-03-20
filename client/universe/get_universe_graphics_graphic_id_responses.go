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

// GetUniverseGraphicsGraphicIDReader is a Reader for the GetUniverseGraphicsGraphicID structure.
type GetUniverseGraphicsGraphicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseGraphicsGraphicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseGraphicsGraphicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseGraphicsGraphicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseGraphicsGraphicIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseGraphicsGraphicIDOK creates a GetUniverseGraphicsGraphicIDOK with default headers values
func NewGetUniverseGraphicsGraphicIDOK() *GetUniverseGraphicsGraphicIDOK {
	return &GetUniverseGraphicsGraphicIDOK{}
}

/*GetUniverseGraphicsGraphicIDOK handles this case with default header values.

Information about a graphic
*/
type GetUniverseGraphicsGraphicIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetUniverseGraphicsGraphicIDOKBody
}

func (o *GetUniverseGraphicsGraphicIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseGraphicsGraphicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseGraphicsGraphicIDNotFound creates a GetUniverseGraphicsGraphicIDNotFound with default headers values
func NewGetUniverseGraphicsGraphicIDNotFound() *GetUniverseGraphicsGraphicIDNotFound {
	return &GetUniverseGraphicsGraphicIDNotFound{}
}

/*GetUniverseGraphicsGraphicIDNotFound handles this case with default header values.

Graphic not found
*/
type GetUniverseGraphicsGraphicIDNotFound struct {
	Payload GetUniverseGraphicsGraphicIDNotFoundBody
}

func (o *GetUniverseGraphicsGraphicIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseGraphicsGraphicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGraphicIDInternalServerError creates a GetUniverseGraphicsGraphicIDInternalServerError with default headers values
func NewGetUniverseGraphicsGraphicIDInternalServerError() *GetUniverseGraphicsGraphicIDInternalServerError {
	return &GetUniverseGraphicsGraphicIDInternalServerError{}
}

/*GetUniverseGraphicsGraphicIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseGraphicsGraphicIDInternalServerError struct {
	Payload GetUniverseGraphicsGraphicIDInternalServerErrorBody
}

func (o *GetUniverseGraphicsGraphicIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/graphics/{graphic_id}/][%d] getUniverseGraphicsGraphicIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseGraphicsGraphicIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseGraphicsGraphicIDInternalServerErrorBody get_universe_graphics_graphic_id_internal_server_error
//
// Internal server error
swagger:model GetUniverseGraphicsGraphicIDInternalServerErrorBody
*/
type GetUniverseGraphicsGraphicIDInternalServerErrorBody struct {

	// get_universe_graphics_graphic_id_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get universe graphics graphic ID internal server error body
func (o *GetUniverseGraphicsGraphicIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseGraphicsGraphicIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetUniverseGraphicsGraphicIDNotFoundBody get_universe_graphics_graphic_id_not_found
//
// Not found
swagger:model GetUniverseGraphicsGraphicIDNotFoundBody
*/
type GetUniverseGraphicsGraphicIDNotFoundBody struct {

	// get_universe_graphics_graphic_id_404_not_found
	//
	// Not found message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get universe graphics graphic ID not found body
func (o *GetUniverseGraphicsGraphicIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseGraphicsGraphicIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetUniverseGraphicsGraphicIDOKBody get_universe_graphics_graphic_id_ok
//
// 200 ok object
swagger:model GetUniverseGraphicsGraphicIDOKBody
*/
type GetUniverseGraphicsGraphicIDOKBody struct {

	// get_universe_graphics_graphic_id_collision_file
	//
	// collision_file string
	// Required: true
	CollisionFile *string `json:"collision_file"`

	// get_universe_graphics_graphic_id_graphic_file
	//
	// graphic_file string
	// Required: true
	GraphicFile *string `json:"graphic_file"`

	// get_universe_graphics_graphic_id_graphic_id
	//
	// graphic_id integer
	// Required: true
	GraphicID *int32 `json:"graphic_id"`

	// get_universe_graphics_graphic_id_icon_folder
	//
	// icon_folder string
	// Required: true
	IconFolder *string `json:"icon_folder"`

	// get_universe_graphics_graphic_id_sof_dna
	//
	// sof_dna string
	// Required: true
	SofDna *string `json:"sof_dna"`

	// get_universe_graphics_graphic_id_sof_fation_name
	//
	// sof_fation_name string
	// Required: true
	SofFationName *string `json:"sof_fation_name"`

	// get_universe_graphics_graphic_id_sof_hull_name
	//
	// sof_hull_name string
	// Required: true
	SofHullName *string `json:"sof_hull_name"`

	// get_universe_graphics_graphic_id_sof_race_name
	//
	// sof_race_name string
	// Required: true
	SofRaceName *string `json:"sof_race_name"`
}

// Validate validates this get universe graphics graphic ID o k body
func (o *GetUniverseGraphicsGraphicIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCollisionFile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGraphicFile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateGraphicID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIconFolder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSofDna(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSofFationName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSofHullName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSofRaceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateCollisionFile(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"collision_file", "body", o.CollisionFile); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateGraphicFile(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"graphic_file", "body", o.GraphicFile); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateGraphicID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"graphic_id", "body", o.GraphicID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateIconFolder(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"icon_folder", "body", o.IconFolder); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateSofDna(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"sof_dna", "body", o.SofDna); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateSofFationName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"sof_fation_name", "body", o.SofFationName); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateSofHullName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"sof_hull_name", "body", o.SofHullName); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGraphicsGraphicIDOKBody) validateSofRaceName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGraphicsGraphicIdOK"+"."+"sof_race_name", "body", o.SofRaceName); err != nil {
		return err
	}

	return nil
}