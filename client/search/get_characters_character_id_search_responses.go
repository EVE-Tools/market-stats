package search

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

// GetCharactersCharacterIDSearchReader is a Reader for the GetCharactersCharacterIDSearch structure.
type GetCharactersCharacterIDSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDSearchOK creates a GetCharactersCharacterIDSearchOK with default headers values
func NewGetCharactersCharacterIDSearchOK() *GetCharactersCharacterIDSearchOK {
	return &GetCharactersCharacterIDSearchOK{}
}

/*GetCharactersCharacterIDSearchOK handles this case with default header values.

A list of search results
*/
type GetCharactersCharacterIDSearchOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload GetCharactersCharacterIDSearchOKBody
}

func (o *GetCharactersCharacterIDSearchOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/search/][%d] getCharactersCharacterIdSearchOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDSearchForbidden creates a GetCharactersCharacterIDSearchForbidden with default headers values
func NewGetCharactersCharacterIDSearchForbidden() *GetCharactersCharacterIDSearchForbidden {
	return &GetCharactersCharacterIDSearchForbidden{}
}

/*GetCharactersCharacterIDSearchForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDSearchForbidden struct {
	Payload GetCharactersCharacterIDSearchForbiddenBody
}

func (o *GetCharactersCharacterIDSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/search/][%d] getCharactersCharacterIdSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSearchInternalServerError creates a GetCharactersCharacterIDSearchInternalServerError with default headers values
func NewGetCharactersCharacterIDSearchInternalServerError() *GetCharactersCharacterIDSearchInternalServerError {
	return &GetCharactersCharacterIDSearchInternalServerError{}
}

/*GetCharactersCharacterIDSearchInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDSearchInternalServerError struct {
	Payload GetCharactersCharacterIDSearchInternalServerErrorBody
}

func (o *GetCharactersCharacterIDSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/search/][%d] getCharactersCharacterIdSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDSearchForbiddenBody get_characters_character_id_search_forbidden
//
// Forbidden
swagger:model GetCharactersCharacterIDSearchForbiddenBody
*/
type GetCharactersCharacterIDSearchForbiddenBody struct {

	// get_characters_character_id_search_403_forbidden
	//
	// Forbidden message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID search forbidden body
func (o *GetCharactersCharacterIDSearchForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDSearchForbiddenBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchForbidden"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDSearchInternalServerErrorBody get_characters_character_id_search_internal_server_error
//
// Internal server error
swagger:model GetCharactersCharacterIDSearchInternalServerErrorBody
*/
type GetCharactersCharacterIDSearchInternalServerErrorBody struct {

	// get_characters_character_id_search_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get characters character ID search internal server error body
func (o *GetCharactersCharacterIDSearchInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDSearchInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDSearchOKBody get_characters_character_id_search_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDSearchOKBody
*/
type GetCharactersCharacterIDSearchOKBody struct {

	// get_characters_character_id_search_agent
	//
	// agent array
	// Required: true
	Agent []int32 `json:"agent"`

	// get_characters_character_id_search_alliance
	//
	// alliance array
	// Required: true
	Alliance []int32 `json:"alliance"`

	// get_characters_character_id_search_character
	//
	// character array
	// Required: true
	Character []int32 `json:"character"`

	// get_characters_character_id_search_constellation
	//
	// constellation array
	// Required: true
	Constellation []int32 `json:"constellation"`

	// get_characters_character_id_search_corporation
	//
	// corporation array
	// Required: true
	Corporation []int32 `json:"corporation"`

	// get_characters_character_id_search_faction
	//
	// faction array
	// Required: true
	Faction []int32 `json:"faction"`

	// get_characters_character_id_search_inventorytype
	//
	// inventorytype array
	// Required: true
	Inventorytype []int32 `json:"inventorytype"`

	// get_characters_character_id_search_region
	//
	// region array
	// Required: true
	Region []int32 `json:"region"`

	// get_characters_character_id_search_solarsystem
	//
	// solarsystem array
	// Required: true
	Solarsystem []int32 `json:"solarsystem"`

	// get_characters_character_id_search_station
	//
	// station array
	// Required: true
	Station []int32 `json:"station"`

	// get_characters_character_id_search_structure
	//
	// structure array
	// Required: true
	Structure []int64 `json:"structure"`
}

// Validate validates this get characters character ID search o k body
func (o *GetCharactersCharacterIDSearchOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAgent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateAlliance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCharacter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateConstellation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateCorporation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateFaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateInventorytype(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSolarsystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStructure(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateAgent(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"agent", "body", o.Agent); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateAlliance(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"alliance", "body", o.Alliance); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateCharacter(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"character", "body", o.Character); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateConstellation(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"constellation", "body", o.Constellation); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateCorporation(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"corporation", "body", o.Corporation); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateFaction(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"faction", "body", o.Faction); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateInventorytype(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"inventorytype", "body", o.Inventorytype); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"region", "body", o.Region); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateSolarsystem(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"solarsystem", "body", o.Solarsystem); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateStation(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"station", "body", o.Station); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSearchOKBody) validateStructure(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdSearchOK"+"."+"structure", "body", o.Structure); err != nil {
		return err
	}

	return nil
}