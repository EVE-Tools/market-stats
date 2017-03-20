package insurance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetInsurancePricesReader is a Reader for the GetInsurancePrices structure.
type GetInsurancePricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInsurancePricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInsurancePricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetInsurancePricesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInsurancePricesOK creates a GetInsurancePricesOK with default headers values
func NewGetInsurancePricesOK() *GetInsurancePricesOK {
	return &GetInsurancePricesOK{}
}

/*GetInsurancePricesOK handles this case with default header values.

A list of insurance levels for all ship types
*/
type GetInsurancePricesOK struct {
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

	Payload []*GetInsurancePricesOKBodyItems0
}

func (o *GetInsurancePricesOK) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesOK  %+v", 200, o.Payload)
}

func (o *GetInsurancePricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInsurancePricesInternalServerError creates a GetInsurancePricesInternalServerError with default headers values
func NewGetInsurancePricesInternalServerError() *GetInsurancePricesInternalServerError {
	return &GetInsurancePricesInternalServerError{}
}

/*GetInsurancePricesInternalServerError handles this case with default header values.

Internal server error
*/
type GetInsurancePricesInternalServerError struct {
	Payload GetInsurancePricesInternalServerErrorBody
}

func (o *GetInsurancePricesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInsurancePricesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInsurancePricesInternalServerErrorBody get_insurance_prices_internal_server_error
//
// Internal server error
swagger:model GetInsurancePricesInternalServerErrorBody
*/
type GetInsurancePricesInternalServerErrorBody struct {

	// get_insurance_prices_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get insurance prices internal server error body
func (o *GetInsurancePricesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetInsurancePricesInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getInsurancePricesInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetInsurancePricesOKBodyItems0 get_insurance_prices_200_ok
//
// 200 ok object
swagger:model GetInsurancePricesOKBodyItems0
*/
type GetInsurancePricesOKBodyItems0 struct {

	// get_insurance_prices_levels
	//
	// A list of a available insurance levels for this ship type
	// Required: true
	Levels []*GetInsurancePricesOKBodyItems0LevelsItems0 `json:"levels"`

	// get_insurance_prices_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get insurance prices o k body items0
func (o *GetInsurancePricesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLevels(formats); err != nil {
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

func (o *GetInsurancePricesOKBodyItems0) validateLevels(formats strfmt.Registry) error {

	if err := validate.Required("levels", "body", o.Levels); err != nil {
		return err
	}

	for i := 0; i < len(o.Levels); i++ {

		if swag.IsZero(o.Levels[i]) { // not required
			continue
		}

		if o.Levels[i] != nil {

			if err := o.Levels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("levels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetInsurancePricesOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

/*GetInsurancePricesOKBodyItems0LevelsItems0 get_insurance_prices_level
//
// level object
swagger:model GetInsurancePricesOKBodyItems0LevelsItems0
*/
type GetInsurancePricesOKBodyItems0LevelsItems0 struct {

	// get_insurance_prices_cost
	//
	// cost number
	// Required: true
	Cost *float32 `json:"cost"`

	// get_insurance_prices_name
	//
	// Localized insurance level
	// Required: true
	Name *string `json:"name"`

	// get_insurance_prices_payout
	//
	// payout number
	// Required: true
	Payout *float32 `json:"payout"`
}

// Validate validates this get insurance prices o k body items0 levels items0
func (o *GetInsurancePricesOKBodyItems0LevelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validatePayout(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInsurancePricesOKBodyItems0LevelsItems0) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", o.Cost); err != nil {
		return err
	}

	return nil
}

func (o *GetInsurancePricesOKBodyItems0LevelsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetInsurancePricesOKBodyItems0LevelsItems0) validatePayout(formats strfmt.Registry) error {

	if err := validate.Required("payout", "body", o.Payout); err != nil {
		return err
	}

	return nil
}