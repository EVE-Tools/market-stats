package sovereignty

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

// GetSovereigntyCampaignsReader is a Reader for the GetSovereigntyCampaigns structure.
type GetSovereigntyCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSovereigntyCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSovereigntyCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSovereigntyCampaignsOK creates a GetSovereigntyCampaignsOK with default headers values
func NewGetSovereigntyCampaignsOK() *GetSovereigntyCampaignsOK {
	return &GetSovereigntyCampaignsOK{}
}

/*GetSovereigntyCampaignsOK handles this case with default header values.

A list of sovereignty campaigns
*/
type GetSovereigntyCampaignsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetSovereigntyCampaignsOKBodyItems0
}

func (o *GetSovereigntyCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyCampaignsInternalServerError creates a GetSovereigntyCampaignsInternalServerError with default headers values
func NewGetSovereigntyCampaignsInternalServerError() *GetSovereigntyCampaignsInternalServerError {
	return &GetSovereigntyCampaignsInternalServerError{}
}

/*GetSovereigntyCampaignsInternalServerError handles this case with default header values.

Internal server error
*/
type GetSovereigntyCampaignsInternalServerError struct {
	Payload GetSovereigntyCampaignsInternalServerErrorBody
}

func (o *GetSovereigntyCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSovereigntyCampaignsInternalServerErrorBody get_sovereignty_campaigns_internal_server_error
//
// Internal server error
swagger:model GetSovereigntyCampaignsInternalServerErrorBody
*/
type GetSovereigntyCampaignsInternalServerErrorBody struct {

	// get_sovereignty_campaigns_500_internal_server_error
	//
	// Internal server error message
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get sovereignty campaigns internal server error body
func (o *GetSovereigntyCampaignsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetSovereigntyCampaignsInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getSovereigntyCampaignsInternalServerError"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

/*GetSovereigntyCampaignsOKBodyItems0 get_sovereignty_campaigns_200_ok
//
// 200 ok object
swagger:model GetSovereigntyCampaignsOKBodyItems0
*/
type GetSovereigntyCampaignsOKBodyItems0 struct {

	// get_sovereignty_campaigns_attackers_score
	//
	// Score for all attacking parties, only present in Defense Events.
	//
	AttackersScore float32 `json:"attackers_score,omitempty"`

	// get_sovereignty_campaigns_campaign_id
	//
	// Unique ID for this campaign.
	// Required: true
	CampaignID *int32 `json:"campaign_id"`

	// get_sovereignty_campaigns_constellation_id
	//
	// The constellation in which the campaign will take place.
	//
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_sovereignty_campaigns_defender_id
	//
	// Defending alliance, only present in Defense Events
	//
	DefenderID int32 `json:"defender_id,omitempty"`

	// get_sovereignty_campaigns_defender_score
	//
	// Score for the defending alliance, only present in Defense Events.
	//
	DefenderScore float32 `json:"defender_score,omitempty"`

	// get_sovereignty_campaigns_event_type
	//
	// Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as "Defense Events", station_freeport as "Freeport Events".
	//
	// Required: true
	EventType *string `json:"event_type"`

	// get_sovereignty_campaigns_participants
	//
	// Alliance participating and their respective scores, only present in Freeport Events.
	//
	Participants []*GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0 `json:"participants"`

	// get_sovereignty_campaigns_solar_system_id
	//
	// The solar system the structure is located in.
	//
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_sovereignty_campaigns_start_time
	//
	// Time the event is scheduled to start.
	//
	// Required: true
	StartTime *strfmt.DateTime `json:"start_time"`

	// get_sovereignty_campaigns_structure_id
	//
	// The structure item ID that is related to this campaign.
	//
	// Required: true
	StructureID *int64 `json:"structure_id"`
}

// Validate validates this get sovereignty campaigns o k body items0
func (o *GetSovereigntyCampaignsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCampaignID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateConstellationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateEventType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateParticipants(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStructureID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaign_id", "body", o.CampaignID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", o.ConstellationID); err != nil {
		return err
	}

	return nil
}

var getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcu_defense","ihub_defense","station_defense","station_freeport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum = append(getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum, v)
	}
}

const (
	// GetSovereigntyCampaignsOKBodyItems0EventTypeTcuDefense captures enum value "tcu_defense"
	GetSovereigntyCampaignsOKBodyItems0EventTypeTcuDefense string = "tcu_defense"
	// GetSovereigntyCampaignsOKBodyItems0EventTypeIhubDefense captures enum value "ihub_defense"
	GetSovereigntyCampaignsOKBodyItems0EventTypeIhubDefense string = "ihub_defense"
	// GetSovereigntyCampaignsOKBodyItems0EventTypeStationDefense captures enum value "station_defense"
	GetSovereigntyCampaignsOKBodyItems0EventTypeStationDefense string = "station_defense"
	// GetSovereigntyCampaignsOKBodyItems0EventTypeStationFreeport captures enum value "station_freeport"
	GetSovereigntyCampaignsOKBodyItems0EventTypeStationFreeport string = "station_freeport"
)

// prop value enum
func (o *GetSovereigntyCampaignsOKBodyItems0) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", o.EventType); err != nil {
		return err
	}

	// value enum
	if err := o.validateEventTypeEnum("event_type", "body", *o.EventType); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateParticipants(formats strfmt.Registry) error {

	if swag.IsZero(o.Participants) { // not required
		return nil
	}

	for i := 0; i < len(o.Participants); i++ {

		if swag.IsZero(o.Participants[i]) { // not required
			continue
		}

		if o.Participants[i] != nil {

			if err := o.Participants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", o.StartTime); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", o.StructureID); err != nil {
		return err
	}

	return nil
}

/*GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0 get_sovereignty_campaigns_participant
//
// participant object
swagger:model GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0
*/
type GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0 struct {

	// get_sovereignty_campaigns_alliance_id
	//
	// alliance_id integer
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_sovereignty_campaigns_score
	//
	// score number
	// Required: true
	Score *float32 `json:"score"`
}

// Validate validates this get sovereignty campaigns o k body items0 participants items0
func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateScore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance_id", "body", o.AllianceID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("score", "body", o.Score); err != nil {
		return err
	}

	return nil
}