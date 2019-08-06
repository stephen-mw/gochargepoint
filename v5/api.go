// Package v5 was auto-generated using wsdl2go.
// Chargepoint v5 wsdl can be found here: https://webservices.chargepoint.com/cp_api_5.0.wsdl
// API programming guide: https://na.chargepoint.com/programmers_guide/5.0 (requires login)
package v5

import (
	"context"
	"encoding/xml"
	"time"
)

// URI represents the chargepoint services v5.0 API
var URI = "https://webservices.chargepoint.com/webservices/chargepoint/services/5.0"

// against "unused imports"
var _ time.Time
var _ xml.Name

type PercentShedRange int32
type CurrencyCodeType string
type ConnectedUserStatusTypes string

const (
	ConnectedUserStatusTypesAPPROVED     ConnectedUserStatusTypes = "APPROVED"
	ConnectedUserStatusTypesNOTCONNECTED ConnectedUserStatusTypes = "NOT CONNECTED"
	ConnectedUserStatusTypesREJECTED     ConnectedUserStatusTypes = "REJECTED"
	ConnectedUserStatusTypesPENDING      ConnectedUserStatusTypes = "PENDING"
)

type ManagedUserStatusTypes string

const (
	ManagedUserStatusTypesAPPROVED   ManagedUserStatusTypes = "APPROVED"
	ManagedUserStatusTypesNOTMANAGED ManagedUserStatusTypes = "NOT MANAGED"
)

type TransTypes string

const (
	TransTypesSession                  TransTypes = "Session"
	TransTypesReservation              TransTypes = "Reservation"
	TransTypesReservationCancel        TransTypes = "Reservation Cancel"
	TransTypesReservationUsageAddendum TransTypes = "Reservation Usage Addendum"
)

type GetPublicStations struct {
	XMLName     xml.Name              `xml:"urn:dictionary:com.chargepoint.webservices getPublicStations"`
	SearchQuery *StationSearchRequest `xml:"searchQuery,omitempty"`
}
type GetPublicStationsResponse struct {
	XMLName      xml.Name       `xml:"urn:dictionary:com.chargepoint.webservices getPublicStationsResponse"`
	ResponseCode string         `xml:"responseCode,omitempty"`
	ResponseText string         `xml:"responseText,omitempty"`
	StationData  []*StationData `xml:"stationData,omitempty"`
}
type GetPublicStationStatus struct {
	XMLName     xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices getPublicStationStatus"`
	SearchQuery *StatusSearchdata `xml:"searchQuery,omitempty"`
}
type GetPublicStationStatusResponse struct {
	XMLName           xml.Name       `xml:"urn:dictionary:com.chargepoint.webservices getPublicStationStatusResponse"`
	ResponseCode      string         `xml:"responseCode,omitempty"`
	ResponseText      string         `xml:"responseText,omitempty"`
	StationStatusData []*OStatusdata `xml:"stationStatusData,omitempty"`
}
type GetStationStatus struct {
	XMLName     xml.Name                        `xml:"urn:dictionary:com.chargepoint.webservices getStationStatus"`
	SearchQuery *StatusSearchdataForGetStations `xml:"searchQuery,omitempty"`
}
type GetStationStatusResponse struct {
	XMLName      xml.Name       `xml:"urn:dictionary:com.chargepoint.webservices getStationStatusResponse"`
	ResponseCode string         `xml:"responseCode,omitempty"`
	ResponseText string         `xml:"responseText,omitempty"`
	StationData  []*OStatusdata `xml:"stationData,omitempty"`
	MoreFlag     int32          `xml:"moreFlag,omitempty"`
}
type GetAlarms struct {
	XMLName     xml.Name              `xml:"urn:dictionary:com.chargepoint.webservices getAlarms"`
	SearchQuery *GetAlarmsSearchQuery `xml:"searchQuery,omitempty"`
}
type GetAlarmsResponse struct {
	XMLName      xml.Name   `xml:"urn:dictionary:com.chargepoint.webservices getAlarmsResponse"`
	ResponseCode string     `xml:"responseCode,omitempty"`
	ResponseText string     `xml:"responseText,omitempty"`
	Alarms       []*Oalarms `xml:"Alarms,omitempty"`
	MoreFlag     int32      `xml:"moreFlag,omitempty"`
}
type ClearAlarms struct {
	XMLName     xml.Name                `xml:"urn:dictionary:com.chargepoint.webservices clearAlarms"`
	SearchQuery *ClearAlarmsSearchQuery `xml:"searchQuery,omitempty"`
}
type ClearAlarmsResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices clearAlarmsResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
}
type ShedLoad struct {
	XMLName   xml.Name                `xml:"urn:dictionary:com.chargepoint.webservices shedLoad"`
	ShedQuery *ShedLoadQueryInputData `xml:"shedQuery,omitempty"`
}
type ShedLoadResponse struct {
	XMLName               xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices shedLoadResponse"`
	ResponseCode          string            `xml:"responseCode,omitempty"`
	ResponseText          string            `xml:"responseText,omitempty"`
	Success               int32             `xml:"Success,omitempty"`
	SgID                  int32             `xml:"sgID,omitempty"`
	StationID             string            `xml:"stationID,omitempty"`
	AllowedLoadPerStation float64           `xml:"allowedLoadPerStation,omitempty"`
	PercentShedPerStation *PercentShedRange `xml:"percentShedPerStation,omitempty"`
	Ports                 struct {
		Port []struct {
			PortNumber         string            `xml:"portNumber,omitempty"`
			AllowedLoadPerPort float64           `xml:"allowedLoadPerPort,omitempty"`
			PercentShedPerPort *PercentShedRange `xml:"percentShedPerPort,omitempty"`
		} `xml:"Port,omitempty"`
	} `xml:"Ports,omitempty"`
}
type ClearShedState struct {
	XMLName   xml.Name            `xml:"urn:dictionary:com.chargepoint.webservices clearShedState"`
	ShedQuery *ShedQueryInputData `xml:"shedQuery,omitempty"`
}
type ClearShedStateResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices clearShedStateResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
	Success      int32    `xml:"Success,omitempty"`
	SgID         int32    `xml:"sgID,omitempty"`
	StationID    string   `xml:"stationID,omitempty"`
}
type GetLoad struct {
	XMLName     xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getLoad"`
	SearchQuery struct {
		SgID      int32  `xml:"sgID,omitempty"`
		StationID string `xml:"stationID,omitempty"`
	} `xml:"searchQuery,omitempty"`
}
type GetLoadResponse struct {
	XMLName      xml.Name           `xml:"urn:dictionary:com.chargepoint.webservices getLoadResponse"`
	ResponseCode string             `xml:"responseCode,omitempty"`
	ResponseText string             `xml:"responseText,omitempty"`
	SgID         int32              `xml:"sgID,omitempty"`
	NumStations  int32              `xml:"numStations,omitempty"`
	GroupName    string             `xml:"groupName,omitempty"`
	SgLoad       string             `xml:"sgLoad,omitempty"`
	StationData  []*Stationloaddata `xml:"stationData,omitempty"`
}
type GetStationGroups struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getStationGroups"`
	OrgID   string   `xml:"orgID,omitempty"`
}
type GetStationGroupsResponse struct {
	XMLName      xml.Name      `xml:"urn:dictionary:com.chargepoint.webservices getStationGroupsResponse"`
	ResponseCode string        `xml:"responseCode,omitempty"`
	ResponseText string        `xml:"responseText,omitempty"`
	GroupData    []*Groupsdata `xml:"groupData,omitempty"`
}
type GetUsers struct {
	XMLName     xml.Name               `xml:"urn:dictionary:com.chargepoint.webservices getUsers"`
	SearchQuery *GetUsersSearchRequest `xml:"searchQuery,omitempty"`
}
type GetUsersResponse struct {
	XMLName      xml.Name    `xml:"urn:dictionary:com.chargepoint.webservices getUsersResponse"`
	ResponseCode string      `xml:"responseCode,omitempty"`
	ResponseText string      `xml:"responseText,omitempty"`
	Users        *UserParams `xml:"users,omitempty"`
}

type GetTransactionData struct {
	XMLName     xml.Name                   `xml:"urn:dictionary:com.chargepoint.webservices getTransactionData"`
	SearchQuery *GetTransDataSearchRequest `xml:"searchQuery,omitempty"`
}
type GetTransactionDataResponse struct {
	XMLName      xml.Name         `xml:"urn:dictionary:com.chargepoint.webservices getTransactionDataResponse"`
	ResponseCode string           `xml:"responseCode,omitempty"`
	ResponseText string           `xml:"responseText,omitempty"`
	Transactions *TransDataParams `xml:"transactions,omitempty"`
	MoreFlag     int32            `xml:"MoreFlag,omitempty"`
}
type GetChargingSessionData struct {
	XMLName     xml.Name           `xml:"urn:dictionary:com.chargepoint.webservices getChargingSessionData"`
	SearchQuery *SessionSearchdata `xml:"searchQuery,omitempty"`
}
type GetChargingSessionDataResponse struct {
	XMLName             xml.Name                   `xml:"urn:dictionary:com.chargepoint.webservices getChargingSessionDataResponse"`
	ResponseCode        string                     `xml:"responseCode,omitempty"`
	ResponseText        string                     `xml:"responseText,omitempty"`
	ChargingSessionData []*SessionSearchResultdata `xml:"ChargingSessionData,omitempty"`
	MoreFlag            int32                      `xml:"MoreFlag,omitempty"`
}
type Get15minChargingSessionData struct {
	XMLName                xml.Name `xml:"urn:dictionary:com.chargepoint.webservices get15minChargingSessionData"`
	SessionID              int64    `xml:"sessionID,omitempty"`
	EnergyConsumedInterval bool     `xml:"energyConsumedInterval,omitempty"`
}
type Get15minChargingSessionDataResponse struct {
	XMLName        xml.Name                 `xml:"urn:dictionary:com.chargepoint.webservices get15minChargingSessionDataResponse"`
	ResponseCode   string                   `xml:"responseCode,omitempty"`
	ResponseText   string                   `xml:"responseText,omitempty"`
	SessionID      int64                    `xml:"sessionID,omitempty"`
	StationID      string                   `xml:"stationID,omitempty"`
	PortNumber     string                   `xml:"portNumber,omitempty"`
	FifteenminData []*OChargingSessionsData `xml:"fifteenminData,omitempty"`
}
type GetOrgsAndStationGroups struct {
	XMLName     xml.Name                            `xml:"urn:dictionary:com.chargepoint.webservices getOrgsAndStationGroups"`
	SearchQuery *GetOrgsAndStationGroupsSearchQuery `xml:"searchQuery,omitempty"`
}
type GetOrgsAndStationGroupsResponse struct {
	XMLName      xml.Name     `xml:"urn:dictionary:com.chargepoint.webservices getOrgsAndStationGroupsResponse"`
	ResponseCode string       `xml:"responseCode,omitempty"`
	ResponseText string       `xml:"responseText,omitempty"`
	OrgData      []*Ohostdata `xml:"orgData,omitempty"`
}
type GetStationGroupDetails struct {
	XMLName   xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getStationGroupDetails"`
	SgID      int32    `xml:"sgID,omitempty"`
	StationID string   `xml:"stationID,omitempty"`
}
type GetStationGroupDetailsResponse struct {
	XMLName      xml.Name            `xml:"urn:dictionary:com.chargepoint.webservices getStationGroupDetailsResponse"`
	ResponseCode string              `xml:"responseCode,omitempty"`
	ResponseText string              `xml:"responseText,omitempty"`
	GroupName    string              `xml:"groupName,omitempty"`
	NumStations  int32               `xml:"numStations,omitempty"`
	GroupData    []*ChildGroupData   `xml:"groupData,omitempty"`
	StationData  []*StationGroupData `xml:"stationData,omitempty"`
}
type GetStations struct {
	XMLName     xml.Name                      `xml:"urn:dictionary:com.chargepoint.webservices getStations"`
	SearchQuery *StationSearchRequestExtended `xml:"searchQuery,omitempty"`
}
type GetStationsResponse struct {
	XMLName      xml.Name               `xml:"urn:dictionary:com.chargepoint.webservices getStationsResponse"`
	ResponseCode string                 `xml:"responseCode,omitempty"`
	ResponseText string                 `xml:"responseText,omitempty"`
	StationData  []*StationDataExtended `xml:"stationData,omitempty"`
	MoreFlag     int32                  `xml:"moreFlag,omitempty"`
}
type GetStationRights struct {
	XMLName     xml.Name                    `xml:"urn:dictionary:com.chargepoint.webservices getStationRights"`
	SearchQuery *StationRightsSearchRequest `xml:"searchQuery,omitempty"`
}
type GetStationRightsResponse struct {
	XMLName      xml.Name      `xml:"urn:dictionary:com.chargepoint.webservices getStationRightsResponse"`
	ResponseCode string        `xml:"responseCode,omitempty"`
	ResponseText string        `xml:"responseText,omitempty"`
	RightsData   []*RightsData `xml:"rightsData,omitempty"`
	MoreFlag     int32         `xml:"moreFlag,omitempty"`
}
type GetStationRightsProfile struct {
	XMLName              xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getStationRightsProfile"`
	SgID                 int32    `xml:"sgID,omitempty"`
	StationRightsProfile string   `xml:"stationRightsProfile,omitempty"`
}
type GetStationRightsProfileResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getStationRightsProfileResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
	TaskList     []struct {
		Task        string `xml:"Task,omitempty"`
		Description string `xml:"Description,omitempty"`
	} `xml:"taskList,omitempty"`
}
type GetCPNInstances struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getCPNInstances"`
}
type GetCPNInstancesResponse struct {
	XMLName xml.Name         `xml:"urn:dictionary:com.chargepoint.webservices getCPNInstancesResponse"`
	CPN     []*Ocpninstances `xml:"CPN,omitempty"`
}
type RegisterFeeds struct {
	XMLName     xml.Name           `xml:"urn:dictionary:com.chargepoint.webservices registerFeeds"`
	Events      *DataEventName     `xml:"Events,omitempty"`
	SearchQuery *FeedStationSearch `xml:"searchQuery,omitempty"`
	FeedType    string             `xml:"feedType,omitempty"`
}
type RegisterFeedsResponse struct {
	XMLName        xml.Name `xml:"urn:dictionary:com.chargepoint.webservices registerFeedsResponse"`
	ResponseCode   string   `xml:"responseCode,omitempty"`
	ResponseText   string   `xml:"responseText,omitempty"`
	SubscriptionID int32    `xml:"subscriptionID,omitempty"`
}
type UpdateFeed struct {
	XMLName        xml.Name `xml:"urn:dictionary:com.chargepoint.webservices updateFeed"`
	SubscriptionID string   `xml:"subscriptionID,omitempty"`
	Refresh        bool     `xml:"Refresh,omitempty"`
}
type UpdateFeedResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices updateFeedResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
}
type UpdateUserStatus struct {
	XMLName         xml.Name `xml:"urn:dictionary:com.chargepoint.webservices updateUserStatus"`
	UserID          string   `xml:"userID,omitempty"`
	AssociationType string   `xml:"associationType,omitempty"`
	Status          string   `xml:"Status,omitempty"`
	CustomText      string   `xml:"customText,omitempty"`
}
type UpdateUserStatusResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices updateUserStatusResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
}
type SetWaitlistDone struct {
	XMLName           xml.Name `xml:"urn:dictionary:com.chargepoint.webservices setWaitlistDone"`
	StationID         string   `xml:"stationID,omitempty"`
	PortID            string   `xml:"portID,omitempty"`
	UserID            string   `xml:"userID,omitempty"`
	CustomMessageText string   `xml:"customMessageText,omitempty"`
}
type SetWaitlistDoneResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices setWaitlistDoneResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
}
type CreateLoadShedEvent struct {
	XMLName   xml.Name   `xml:"urn:dictionary:com.chargepoint.webservices createLoadShedEvent"`
	ShedEvent *ShedEvent `xml:"shedEvent,omitempty"`
}
type CreateLoadShedEventResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices createLoadShedEventResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
	EventId      int32    `xml:"eventId,omitempty"`
}
type RemoveLoadShedEvent struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices removeLoadShedEvent"`
	EventId int32    `xml:"eventId,omitempty"`
}
type RemoveLoadShedEventResponse struct {
	XMLName      xml.Name `xml:"urn:dictionary:com.chargepoint.webservices removeLoadShedEventResponse"`
	ResponseCode string   `xml:"responseCode,omitempty"`
	ResponseText string   `xml:"responseText,omitempty"`
	Response     string   `xml:"response,omitempty"`
}
type StationSearchRequest struct {
	XMLName             xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices stationSearchRequest"`
	StationID           string            `xml:"stationID,omitempty"`
	StationManufacturer string            `xml:"stationManufacturer,omitempty"`
	StationModel        string            `xml:"stationModel,omitempty"`
	StationName         string            `xml:"stationName,omitempty"`
	SerialNumber        string            `xml:"serialNumber,omitempty"`
	Address             string            `xml:"Address,omitempty"`
	City                string            `xml:"City,omitempty"`
	State               string            `xml:"State,omitempty"`
	Country             string            `xml:"Country,omitempty"`
	PostalCode          string            `xml:"postalCode,omitempty"`
	Proximity           float64           `xml:"Proximity,omitempty"`
	ProximityUnit       string            `xml:"proximityUnit,omitempty"`
	Connector           string            `xml:"Connector,omitempty"`
	Voltage             string            `xml:"Voltage,omitempty"`
	Current             string            `xml:"Current,omitempty"`
	Power               string            `xml:"Power,omitempty"`
	DemoSerialNumber    *SerialNumberData `xml:"demoSerialNumber,omitempty"`
	Reservable          bool              `xml:"Reservable,omitempty"`
	Geo                 *GeoData
	Level               string `xml:"Level,omitempty"`
	Mode                string `xml:"Mode,omitempty"`
	Pricing             *PricingOptions
}
type GeoData struct {
	XMLName xml.Name `xml:"Geo,omitempty"`
	Lat     string   `xml:"Lat,omitempty"`
	Long    string   `xml:"Long,omitempty"`
}
type PricingOptions struct {
	XMLName        xml.Name  `xml:"Pricing,omitempty"`
	StartTime      time.Time `xml:"startTime,omitempty"`
	Duration       int32     `xml:"Duration,omitempty"`
	EnergyRequired float64   `xml:"energyRequired,omitempty"`
	VehiclePower   float64   `xml:"vehiclePower,omitempty"`
}
type SerialNumberData struct {
	SerialNumber []string `xml:"serialNumber,omitempty"`
}

type StationData struct {
	XMLName             xml.Name              `xml:"urn:dictionary:com.chargepoint.webservices stationData"`
	StationID           string                `xml:"stationID,omitempty"`
	StationManufacturer string                `xml:"stationManufacturer,omitempty"`
	StationModel        string                `xml:"stationModel,omitempty"`
	StationMacAddr      string                `xml:"stationMacAddr,omitempty"`
	StationSerialNum    string                `xml:"stationSerialNum,omitempty"`
	Address             string                `xml:"Address,omitempty"`
	City                string                `xml:"City,omitempty"`
	State               string                `xml:"State,omitempty"`
	Country             string                `xml:"Country,omitempty"`
	PostalCode          string                `xml:"postalCode,omitempty"`
	Port                []*PortData           `xml:"Port,omitempty"`
	Pricing             *PricingSpecification `xml:"Pricing,omitempty"`
	NumPorts            int32                 `xml:"numPorts,omitempty"`
	MainPhone           string                `xml:"mainPhone,omitempty"`
	ModTimeStamp        time.Time             `xml:"modTimeStamp,omitempty"`
	TimezoneOffset      string                `xml:"timezoneOffset,omitempty"`
	CurrencyCode        CurrencyCodeType      `xml:"currencyCode,omitempty"`
}
type GetStationData struct {
	XMLName               xml.Name              `xml:"urn:dictionary:com.chargepoint.webservices getStationData"`
	StationID             string                `xml:"stationID,omitempty"`
	StationManufacturer   string                `xml:"stationManufacturer,omitempty"`
	StationModel          string                `xml:"stationModel,omitempty"`
	StationMacAddr        string                `xml:"stationMacAddr,omitempty"`
	StationSerialNum      string                `xml:"stationSerialNum,omitempty"`
	StationActivationDate time.Time             `xml:"stationActivationDate,omitempty"`
	Address               string                `xml:"Address,omitempty"`
	City                  string                `xml:"City,omitempty"`
	State                 string                `xml:"State,omitempty"`
	Country               string                `xml:"Country,omitempty"`
	PostalCode            string                `xml:"postalCode,omitempty"`
	Port                  []*PortData           `xml:"Port,omitempty"`
	Pricing               *PricingSpecification `xml:"Pricing,omitempty"`
	NumPorts              int32                 `xml:"numPorts,omitempty"`
	DriverName            string                `xml:"driverName,omitempty"`
	DriverAddress         string                `xml:"driverAddress,omitempty"`
	DriverEmail           string                `xml:"driverEmail,omitempty"`
	DriverPhoneNumber     string                `xml:"driverPhoneNumber,omitempty"`
	LastModifiedDate      time.Time             `xml:"lastModifiedDate,omitempty"`
	MainPhone             string                `xml:"mainPhone,omitempty"`
	ModTimeStamp          time.Time             `xml:"modTimeStamp,omitempty"`
	TimezoneOffset        string                `xml:"timezoneOffset,omitempty"`
	CurrencyCode          CurrencyCodeType      `xml:"currencyCode,omitempty"`
}
type PricingSpecification struct {
	Type                       string    `xml:"Type,omitempty"`
	StartTime                  time.Time `xml:"startTime,omitempty"`
	EndTime                    time.Time `xml:"endTime,omitempty"`
	MinPrice                   float64   `xml:"minPrice,omitempty"`
	MaxPrice                   float64   `xml:"maxPrice,omitempty"`
	InitialUnitPriceDuration   int32     `xml:"initialUnitPriceDuration,omitempty"`
	UnitPricePerHour           float64   `xml:"unitPricePerHour,omitempty"`
	UnitPricePerHourThereafter float64   `xml:"unitPricePerHourThereafter,omitempty"`
	UnitPricePerSession        float64   `xml:"unitPricePerSession,omitempty"`
	UnitPricePerKWh            float64   `xml:"unitPricePerKWh,omitempty"`
	SessionTime                string    `xml:"sessionTime,omitempty"`
}
type PortData struct {
	PortNumber    string    `xml:"portNumber,omitempty"`
	StationName   string    `xml:"stationName,omitempty"`
	Geo           *GeoData  `xml:"Geo,omitempty"`
	Description   string    `xml:"Description,omitempty"`
	Reservable    int32     `xml:"Reservable,omitempty"`
	Status        string    `xml:"Status,omitempty"`
	Level         string    `xml:"Level,omitempty"`
	TimeStamp     time.Time `xml:"timeStamp,omitempty"`
	Mode          string    `xml:"Mode,omitempty"`
	Connector     string    `xml:"Connector,omitempty"`
	Voltage       string    `xml:"Voltage,omitempty"`
	Current       string    `xml:"Current,omitempty"`
	Power         string    `xml:"Power,omitempty"`
	EstimatedCost float64   `xml:"estimatedCost,omitempty"`
}

// StatusSearchdata is used to search status
type StatusSearchdata struct {
	XMLName             xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices statusSearchdata"`
	StationID           string            `xml:"stationID,omitempty"`
	StationIDs          *StationIdList    `xml:"stationIDs,omitempty"`
	StationManufacturer string            `xml:"stationManufacturer,omitempty"`
	StationModel        string            `xml:"stationModel,omitempty"`
	StationName         string            `xml:"stationName,omitempty"`
	Address             string            `xml:"Address,omitempty"`
	City                string            `xml:"City,omitempty"`
	State               string            `xml:"State,omitempty"`
	Country             string            `xml:"Country,omitempty"`
	PostalCode          string            `xml:"postalCode,omitempty"`
	Proximity           float64           `xml:"Proximity,omitempty"`
	ProximityUnit       string            `xml:"proximityUnit,omitempty"`
	Connector           string            `xml:"Connector,omitempty"`
	Voltage             string            `xml:"Voltage,omitempty"`
	Current             string            `xml:"Current,omitempty"`
	Power               string            `xml:"Power,omitempty"`
	PortDetails         bool              `xml:"portDetails,omitempty"`
	DemoSerialNumber    *SerialNumberData `xml:"demoSerialNumber,omitempty"`
	Reservable          bool              `xml:"Reservable,omitempty"`
	Geo                 *GeoData          `xml:"Geo,omitempty"`
	Status              int32             `xml:"Status,omitempty"`
	Level               string            `xml:"Level,omitempty"`
	Mode                string            `xml:"Mode,omitempty"`
	Pricing             *PricingOptions   `xml:"Pricing,omitempty"`
}
type StatusSearchdataForGetStations struct {
	XMLName             xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices statusSearchdataForGetStations"`
	StationID           string            `xml:"stationID,omitempty"`
	StationIDs          *StationIdList    `xml:"stationIDs,omitempty"`
	StationManufacturer string            `xml:"stationManufacturer,omitempty"`
	StationModel        string            `xml:"stationModel,omitempty"`
	StationName         string            `xml:"stationName,omitempty"`
	Address             string            `xml:"Address,omitempty"`
	City                string            `xml:"City,omitempty"`
	State               string            `xml:"State,omitempty"`
	Country             string            `xml:"Country,omitempty"`
	PostalCode          string            `xml:"postalCode,omitempty"`
	Proximity           float64           `xml:"Proximity,omitempty"`
	ProximityUnit       string            `xml:"proximityUnit,omitempty"`
	Connector           string            `xml:"Connector,omitempty"`
	Voltage             string            `xml:"Voltage,omitempty"`
	Current             string            `xml:"Current,omitempty"`
	Power               string            `xml:"Power,omitempty"`
	PortDetails         bool              `xml:"portDetails,omitempty"`
	DemoSerialNumber    *SerialNumberData `xml:"demoSerialNumber,omitempty"`
	Reservable          bool              `xml:"Reservable,omitempty"`
	Geo                 *GeoData          `xml:"Geo,omitempty"`
	Status              int32             `xml:"Status,omitempty"`
	Level               string            `xml:"Level,omitempty"`
	Mode                string            `xml:"Mode,omitempty"`
	Pricing             *PricingOptions   `xml:"Pricing,omitempty"`
	StartRecord         int32             `xml:"startRecord,omitempty"`
	NumStations         int32             `xml:"numStations,omitempty"`
}
type StationIdList struct {
	XMLName   xml.Name `xml:"urn:dictionary:com.chargepoint.webservices stationIdList"`
	StationID []string `xml:"stationID,omitempty"`
}
type OStatusdata struct {
	XMLName   xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices oStatusdata"`
	StationID string            `xml:"stationID,omitempty"`
	Port      []*PortDataStatus `xml:"Port,omitempty"`
}
type PortDataStatus struct {
	XMLName    xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices portDataStatus"`
	PortNumber string    `xml:"portNumber,omitempty"`
	Status     string    `xml:"Status,omitempty"`
	TimeStamp  time.Time `xml:"TimeStamp,omitempty"`
	Connector  string    `xml:"Connector,omitempty"`
	Power      string    `xml:"Power,omitempty"`
}
type Oalarms struct {
	XMLName             xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices oalarms"`
	StationID           string    `xml:"stationID,omitempty"`
	StationName         string    `xml:"stationName,omitempty"`
	StationModel        string    `xml:"stationModel,omitempty"`
	OrgID               string    `xml:"orgID,omitempty"`
	OrganizationName    string    `xml:"organizationName,omitempty"`
	StationManufacturer string    `xml:"stationManufacturer,omitempty"`
	StationSerialNum    string    `xml:"stationSerialNum,omitempty"`
	PortNumber          string    `xml:"portNumber,omitempty"`
	AlarmType           string    `xml:"alarmType,omitempty"`
	AlarmTime           time.Time `xml:"alarmTime,omitempty"`
	RecordNumber        int32     `xml:"recordNumber,omitempty"`
}
type GetAlarmsSearchQuery struct {
	XMLName          xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices getAlarmsSearchQuery"`
	OrgID            string    `xml:"orgID,omitempty"`
	OrganizationName string    `xml:"organizationName,omitempty"`
	StationID        string    `xml:"stationID,omitempty"`
	StationName      string    `xml:"stationName,omitempty"`
	SgID             int32     `xml:"sgID,omitempty"`
	SgName           string    `xml:"sgName,omitempty"`
	StartTime        time.Time `xml:"startTime,omitempty"`
	EndTime          time.Time `xml:"endTime,omitempty"`
	PortNumber       string    `xml:"portNumber,omitempty"`
	StartRecord      int32     `xml:"startRecord,omitempty"`
	NumTransactions  int32     `xml:"numTransactions,omitempty"`
}
type ClearAlarmsSearchQuery struct {
	XMLName          xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices clearAlarmsSearchQuery"`
	OrgID            string    `xml:"orgID,omitempty"`
	OrganizationName string    `xml:"organizationName,omitempty"`
	StationID        string    `xml:"stationID,omitempty"`
	StationName      string    `xml:"stationName,omitempty"`
	SgID             int32     `xml:"sgID,omitempty"`
	SgName           string    `xml:"sgName,omitempty"`
	StartTime        time.Time `xml:"startTime,omitempty"`
	EndTime          time.Time `xml:"endTime,omitempty"`
	PortNumber       string    `xml:"portNumber,omitempty"`
	AlarmType        string    `xml:"alarmType,omitempty"`
	ClearReason      string    `xml:"clearReason,omitempty"`
}
type Stationloaddata struct {
	XMLName     xml.Name           `xml:"urn:dictionary:com.chargepoint.webservices stationloaddata"`
	StationID   string             `xml:"stationID,omitempty"`
	StationName string             `xml:"stationName,omitempty"`
	Address     string             `xml:"Address,omitempty"`
	StationLoad float64            `xml:"stationLoad,omitempty"`
	Port        []*StationPortData `xml:"Port,omitempty"`
}
type StationPortData struct {
	XMLName      xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices stationPortData"`
	PortNumber   string            `xml:"portNumber,omitempty"`
	UserID       string            `xml:"userID,omitempty"`
	CredentialID string            `xml:"credentialID,omitempty"`
	ShedState    int32             `xml:"shedState,omitempty"`
	PortLoad     float64           `xml:"portLoad,omitempty"`
	AllowedLoad  float64           `xml:"allowedLoad,omitempty"`
	PercentShed  *PercentShedRange `xml:"percentShed,omitempty"`
}
type ShedQueryInputData struct {
	XMLName     xml.Name              `xml:"urn:dictionary:com.chargepoint.webservices shedQueryInputData"`
	ShedGroup   *ShedGroupInputData   `xml:"shedGroup,omitempty"`
	ShedStation *ShedStationInputData `xml:"shedStation,omitempty"`
}
type ShedLoadQueryInputData struct {
	XMLName      xml.Name                  `xml:"urn:dictionary:com.chargepoint.webservices shedLoadQueryInputData"`
	ShedGroup    *ShedLoadGroupInputData   `xml:"shedGroup,omitempty"`
	ShedStation  *ShedLoadStationInputData `xml:"shedStation,omitempty"`
	TimeInterval int32                     `xml:"timeInterval,omitempty"`
}
type ShedGroupInputData struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices shedGroupInputData"`
	SgID    int32    `xml:"sgID,omitempty"`
}
type ShedStationInputData struct {
	XMLName   xml.Name        `xml:"urn:dictionary:com.chargepoint.webservices shedStationInputData"`
	StationID string          `xml:"stationID,omitempty"`
	Ports     *PortsDataBlock `xml:"Ports,omitempty"`
}
type PortsDataBlock struct {
	XMLName xml.Name        `xml:"urn:dictionary:com.chargepoint.webservices portsDataBlock"`
	Port    []*PortDataShed `xml:"Port,omitempty"`
}

type PortDataShed struct {
	XMLName    xml.Name `xml:"urn:dictionary:com.chargepoint.webservices portDataShed"`
	PortNumber string   `xml:"portNumber,omitempty"`
}

type ShedLoadGroupInputData struct {
	XMLName               xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices shedLoadGroupInputData"`
	SgID                  int32             `xml:"sgID,omitempty"`
	AllowedLoadPerStation float64           `xml:"allowedLoadPerStation,omitempty"`
	PercentShedPerStation *PercentShedRange `xml:"percentShedPerStation,omitempty"`
}

type ShedLoadStationInputData struct {
	XMLName               xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices shedLoadStationInputData"`
	StationID             string            `xml:"stationID,omitempty"`
	AllowedLoadPerStation float64           `xml:"allowedLoadPerStation,omitempty"`
	PercentShedPerStation *PercentShedRange `xml:"percentShedPerStation,omitempty"`
	Ports                 struct {
		Port []struct {
			PortNumber         string            `xml:"portNumber,omitempty"`
			AllowedLoadPerPort float64           `xml:"allowedLoadPerPort,omitempty"`
			PercentShedPerPort *PercentShedRange `xml:"percentShedPerPort,omitempty"`
		} `xml:"Port,omitempty"`
	} `xml:"Ports,omitempty"`
}

type Groupsdata struct {
	XMLName          xml.Name `xml:"urn:dictionary:com.chargepoint.webservices groupsdata"`
	SgID             int32    `xml:"sgID,omitempty"`
	OrgID            string   `xml:"orgID,omitempty"`
	SgName           string   `xml:"sgName,omitempty"`
	OrganizationName string   `xml:"organizationName,omitempty"`
	StationData      []struct {
		StationID string   `xml:"stationID,omitempty"`
		Geo       *GeoData `xml:"Geo,omitempty"`
	} `xml:"stationData,omitempty"`
}

type SessionSearchdata struct {
	XMLName       xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices sessionSearchdata"`
	StationID     string    `xml:"stationID,omitempty"`
	SessionID     int64     `xml:"sessionID,omitempty"`
	StationName   string    `xml:"stationName,omitempty"`
	Address       string    `xml:"Address,omitempty"`
	City          string    `xml:"City,omitempty"`
	State         string    `xml:"State,omitempty"`
	Country       string    `xml:"Country,omitempty"`
	PostalCode    string    `xml:"postalCode,omitempty"`
	Proximity     float64   `xml:"Proximity,omitempty"`
	ProximityUnit string    `xml:"proximityUnit,omitempty"`
	FromTimeStamp time.Time `xml:"fromTimeStamp,omitempty"`
	ToTimeStamp   time.Time `xml:"toTimeStamp,omitempty"`
	StartRecord   int64     `xml:"startRecord,omitempty"`
	Geo           *GeoData  `xml:"Geo,omitempty"`
}

type SessionSearchResultdata struct {
	XMLName      xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices sessionSearchResultdata"`
	StationID    string    `xml:"stationID,omitempty"`
	StationName  string    `xml:"stationName,omitempty"`
	PortNumber   string    `xml:"portNumber,omitempty"`
	Address      string    `xml:"Address,omitempty"`
	City         string    `xml:"City,omitempty"`
	State        string    `xml:"State,omitempty"`
	Country      string    `xml:"Country,omitempty"`
	PostalCode   string    `xml:"postalCode,omitempty"`
	SessionID    int64     `xml:"sessionID,omitempty"`
	Energy       float64   `xml:"Energy,omitempty"`
	StartTime    time.Time `xml:"startTime,omitempty"`
	EndTime      time.Time `xml:"endTime,omitempty"`
	UserID       string    `xml:"userID,omitempty"`
	RecordNumber int64     `xml:"recordNumber,omitempty"`
	CredentialID string    `xml:"credentialID,omitempty"`
}

type OChargingSessionsData struct {
	XMLName         xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices oChargingSessionsData"`
	StationTime     time.Time `xml:"stationTime,omitempty"`
	EnergyConsumed  float64   `xml:"energyConsumed,omitempty"`
	PeakPower       float64   `xml:"peakPower,omitempty"`
	RollingPowerAvg float64   `xml:"rollingPowerAvg,omitempty"`
}

type GetOrgsAndStationGroupsSearchQuery struct {
	XMLName          xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getOrgsAndStationGroupsSearchQuery"`
	OrgID            string   `xml:"orgID,omitempty"`
	OrganizationName string   `xml:"organizationName,omitempty"`
	SgID             int32    `xml:"sgID,omitempty"`
	SgName           string   `xml:"sgName,omitempty"`
}

type Ohostdata struct {
	XMLName          xml.Name `xml:"urn:dictionary:com.chargepoint.webservices ohostdata"`
	OrgID            string   `xml:"orgID,omitempty"`
	OrganizationName string   `xml:"organizationName,omitempty"`
	SgData           []struct {
		SgID          int32  `xml:"sgID,omitempty"`
		SgName        string `xml:"sgName,omitempty"`
		ParentGroupID string `xml:"parentGroupID,omitempty"`
	} `xml:"sgData,omitempty"`
}

type StationGroupData struct {
	StationID   string `xml:"stationID,omitempty"`
	StationName string `xml:"stationName,omitempty"`
	Address     string `xml:"Address,omitempty"`
}

type ChildGroupData struct {
	SgID   int32  `xml:"sgID,omitempty"`
	SgName string `xml:"sgName,omitempty"`
}

type StationSearchRequestExtended struct {
	StationID             string            `xml:"stationID,omitempty"`
	StationManufacturer   string            `xml:"stationManufacturer,omitempty"`
	StationModel          string            `xml:"stationModel,omitempty"`
	StationName           string            `xml:"stationName,omitempty"`
	SerialNumber          string            `xml:"serialNumber,omitempty"`
	Address               string            `xml:"Address,omitempty"`
	City                  string            `xml:"City,omitempty"`
	State                 string            `xml:"State,omitempty"`
	Country               string            `xml:"Country,omitempty"`
	PostalCode            string            `xml:"postalCode,omitempty"`
	Proximity             float64           `xml:"Proximity,omitempty"`
	ProximityUnit         string            `xml:"proximityUnit,omitempty"`
	Connector             string            `xml:"Connector,omitempty"`
	Voltage               string            `xml:"Voltage,omitempty"`
	Current               string            `xml:"Current,omitempty"`
	Power                 string            `xml:"Power,omitempty"`
	DemoSerialNumber      *SerialNumberData `xml:"demoSerialNumber,omitempty"`
	Reservable            bool              `xml:"Reservable,omitempty"`
	Geo                   *GeoData          `xml:"Geo,omitempty"`
	Level                 string            `xml:"Level,omitempty"`
	Mode                  string            `xml:"Mode,omitempty"`
	Pricing               *PricingOptions   `xml:"Pricing,omitempty"`
	OrgID                 string            `xml:"orgID,omitempty"`
	OrganizationName      string            `xml:"organizationName,omitempty"`
	SgID                  string            `xml:"sgID,omitempty"`
	SgName                string            `xml:"sgName,omitempty"`
	StationActivationDate time.Time         `xml:"stationActivationDate,omitempty"`
	StartRecord           int32             `xml:"startRecord,omitempty"`
	NumStations           int32             `xml:"numStations,omitempty"`
}

type StationDataExtended struct {
	*GetStationData
	XMLName          xml.Name `xml:"stationData"`
	OrgID            string   `xml:"orgID,omitempty"`
	OrganizationName string   `xml:"organizationName,omitempty"`
	SgID             string   `xml:"sgID,omitempty"`
	SgName           string   `xml:"sgName,omitempty"`
}

type StationRightsSearchRequest struct {
	XMLName             xml.Name          `xml:"urn:dictionary:com.chargepoint.webservices stationRightsSearchRequest"`
	StationID           string            `xml:"stationID,omitempty"`
	StationManufacturer string            `xml:"stationManufacturer,omitempty"`
	StationModel        string            `xml:"stationModel,omitempty"`
	StationName         string            `xml:"stationName,omitempty"`
	SerialNumber        string            `xml:"serialNumber,omitempty"`
	Address             string            `xml:"Address,omitempty"`
	City                string            `xml:"City,omitempty"`
	State               string            `xml:"State,omitempty"`
	Country             string            `xml:"Country,omitempty"`
	PostalCode          string            `xml:"postalCode,omitempty"`
	Proximity           float64           `xml:"Proximity,omitempty"`
	ProximityUnit       string            `xml:"proximityUnit,omitempty"`
	Connector           string            `xml:"Connector,omitempty"`
	Voltage             string            `xml:"Voltage,omitempty"`
	Current             string            `xml:"Current,omitempty"`
	Power               string            `xml:"Power,omitempty"`
	DemoSerialNumber    *SerialNumberData `xml:"demoSerialNumber,omitempty"`
	Reservable          bool              `xml:"Reservable,omitempty"`
	Geo                 *GeoData          `xml:"Geo,omitempty"`
	Level               string            `xml:"Level,omitempty"`
	Mode                string            `xml:"Mode,omitempty"`
	Pricing             *PricingOptions   `xml:"Pricing,omitempty"`
	OrgID               string            `xml:"orgID,omitempty"`
	OrganizationName    string            `xml:"organizationName,omitempty"`
	SgID                string            `xml:"sgID,omitempty"`
	SgName              string            `xml:"sgName,omitempty"`
	ProvisionDateRange  struct {
		StartDate time.Time `xml:"startDate,omitempty"`
		EndDate   time.Time `xml:"endDate,omitempty"`
	} `xml:"provisionDateRange,omitempty"`
	CurrentFault    string `xml:"currentFault,omitempty"`
	PortStatus      string `xml:"portStatus,omitempty"`
	AdminStatus     string `xml:"adminStatus,omitempty"`
	NetworkStatus   string `xml:"networkStatus,omitempty"`
	ProvisionStatus string `xml:"provisionStatus,omitempty"`
	StartRecord     int64  `xml:"startRecord,omitempty"`
}

type RightsData struct {
	SgID                 string               `xml:"sgID,omitempty"`
	SgName               string               `xml:"sgName,omitempty"`
	StationRightsProfile string               `xml:"stationRightsProfile,omitempty"`
	StationData          []*StationDataRights `xml:"stationData,omitempty"`
}

type StationDataRights struct {
	XMLName          xml.Name `xml:"urn:dictionary:com.chargepoint.webservices stationDataRights"`
	StationID        string   `xml:"stationID,omitempty"`
	StationName      string   `xml:"stationName,omitempty"`
	StationSerialNum string   `xml:"stationSerialNum,omitempty"`
	StationMacAddr   string   `xml:"stationMacAddr,omitempty"`
}

type UserParams struct {
	//XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices userParams"`
	User     []*UserData `xml:"user,omitempty"`
	MoreFlag int32       `xml:"moreFlag,omitempty"`
}

type GetUsersSearchRequest struct {
	Connection            *ConnectionDataRequest   `xml:"Connection,omitempty"`
	CredentialID          string                   `xml:"credentialID,omitempty"`
	FirstName             string                   `xml:"firstName,omitempty"`
	LastModifiedTimeStamp string                   `xml:"lastModifiedTimeStamp,omitempty"`
	LastName              string                   `xml:"lastName,omitempty"`
	ManagementRealm       *ManagementRealmRequest  `xml:"managementRealm,omitempty"`
	NumUsers              int32                    `xml:"numUsers,omitempty"`
	StartRecord           int32                    `xml:"startRecord,omitempty"`
	Status                ConnectedUserStatusTypes `xml:"status,omitempty"`
	UserID                string                   `xml:"userID,omitempty"`
}

type UserData struct {
	LastModifiedTimestamp time.Time            `xml:"lastModifiedTimestamp,omitempty"`
	UserID                string               `xml:"userID,omitempty"`
	FirstName             string               `xml:"firstName,omitempty"`
	LastName              string               `xml:"lastName,omitempty"`
	Connection            *ConnectionData      `xml:"Connection,omitempty"`
	ManagementRealm       *ManagementRealmData `xml:"managementRealm,omitempty"`
	CredentialIDs         *CredentialIDsData   `xml:"credentialIDs,omitempty"`
	RecordNumber          int32                `xml:"recordNumber,omitempty"`
}

type ConnectionDataRequest struct {
	Status     ConnectedUserStatusTypes `xml:"Status,omitempty"`
	CustomInfo *CustomInfoData          `xml:"customInfo,omitempty"`
}

type ConnectionData struct {
	Status           *ConnectedUserStatusTypes `xml:"Status,omitempty"`
	RequestTimeStamp time.Time                 `xml:"requestTimeStamp,omitempty"`
	CustomInfos      *CustomInfosData          `xml:"customInfos,omitempty"`
}
type ManagementRealmRequest struct {
	Status     *ManagedUserStatusTypes `xml:"Status,omitempty"`
	CustomInfo *CustomInfoData         `xml:"customInfo,omitempty"`
}
type ManagementRealmData struct {
	Status          *ManagedUserStatusTypes `xml:"Status,omitempty"`
	SignupTimeStamp time.Time               `xml:"signupTimeStamp,omitempty"`
	CustomInfos     *CustomInfosData        `xml:"customInfos,omitempty"`
}
type CustomInfosData struct {
	CustomInfo []*CustomInfoData `xml:"customInfo,omitempty"`
}
type CustomInfoData struct {
	Key   string `xml:"Key,omitempty"`
	Value string `xml:"Value,omitempty"`
}
type CredentialIDsData struct {
	CredentialID []string `xml:"credentialID,omitempty"`
}
type Ocpninstances struct {
	XMLName        xml.Name `xml:"urn:dictionary:com.chargepoint.webservices ocpninstances"`
	CpnID          string   `xml:"cpnID,omitempty"`
	CpnName        string   `xml:"cpnName,omitempty"`
	CpnDescription string   `xml:"cpnDescription,omitempty"`
}
type DataEventName struct {
	FeedEventName []XMPPFeedEvent `xml:"feedEventName,omitempty"`
}
type FeedStationSearch struct {
	StationID     string   `xml:"stationID,omitempty"`
	StationModel  string   `xml:"stationModel,omitempty"`
	StationName   string   `xml:"stationName,omitempty"`
	Address       string   `xml:"Address,omitempty"`
	City          string   `xml:"City,omitempty"`
	State         string   `xml:"State,omitempty"`
	Country       string   `xml:"Country,omitempty"`
	PostalCode    string   `xml:"postalCode,omitempty"`
	Level         string   `xml:"Level,omitempty"`
	Reservable    bool     `xml:"Reservable,omitempty"`
	Geo           *GeoData `xml:"Geo,omitempty"`
	Proximity     int32    `xml:"Proximity,omitempty"`
	ProximityUnit int32    `xml:"proximityUnit,omitempty"`
	Connector     string   `xml:"Connector,omitempty"`
	Voltage       string   `xml:"Voltage,omitempty"`
	Current       string   `xml:"Current,omitempty"`
	Power         string   `xml:"Power,omitempty"`
	SgID          int32    `xml:"sgID,omitempty"`
}
type GetTransDataSearchRequest struct {
	StationID                string      `xml:"stationID,omitempty"`
	StationName              string      `xml:"stationName,omitempty"`
	StationMacAddr           string      `xml:"stationMacAddr,omitempty"`
	OrgID                    string      `xml:"orgID,omitempty"`
	OrganizationName         string      `xml:"organizationName,omitempty"`
	PricingRuleName          string      `xml:"pricingRuleName,omitempty"`
	TransactionType          *TransTypes `xml:"transactionType,omitempty"`
	TransactionID            int64       `xml:"transactionID,omitempty"`
	FromTransactionTimeStamp time.Time   `xml:"fromTransactionTimeStamp,omitempty"`
	ToTransactionTimeStamp   time.Time   `xml:"toTransactionTimeStamp,omitempty"`
	StartRecord              int32       `xml:"startRecord,omitempty"`
	NumTransactions          int32       `xml:"numTransactions,omitempty"`
}
type TransDataParams struct {
	XMLName         xml.Name     `xml:"urn:dictionary:com.chargepoint.webservices transDataParams"`
	TransactionData []*TransData `xml:"transactionData,omitempty"`
}
type TransData struct {
	StationID             string            `xml:"stationID,omitempty"`
	StationName           string            `xml:"stationName,omitempty"`
	StationMacAddr        string            `xml:"stationMacAddr,omitempty"`
	PortNumber            int32             `xml:"portNumber,omitempty"`
	OrganizationName      string            `xml:"organizationName,omitempty"`
	PricingRuleName       string            `xml:"pricingRuleName,omitempty"`
	TransactionType       *TransTypes       `xml:"transactionType,omitempty"`
	TransactionID         int64             `xml:"transactionID,omitempty"`
	Energy                float64           `xml:"Energy,omitempty"`
	TransactionTime       time.Time         `xml:"transactionTime,omitempty"`
	StartTime             time.Time         `xml:"startTime,omitempty"`
	EndTime               time.Time         `xml:"endTime,omitempty"`
	Currency              *CurrencyCodeType `xml:"Currency,omitempty"`
	GrossAmount           float64           `xml:"grossAmount,omitempty"`
	FlexBillingServiceFee float64           `xml:"flexBillingServiceFee,omitempty"`
	NetRevenue            float64           `xml:"netRevenue,omitempty"`
	ExchangeRateUSD       float64           `xml:"exchangeRateUSD,omitempty"`
	RecordNumber          int32             `xml:"recordNumber,omitempty"`
}
type ShedEvent struct {
	XMLName               xml.Name  `xml:"urn:dictionary:com.chargepoint.webservices shedEvent"`
	StationID             string    `xml:"stationID,omitempty"`
	SgID                  string    `xml:"sgID,omitempty"`
	EventStartTime        time.Time `xml:"eventStartTime,omitempty"`
	EventDuration         int32     `xml:"eventDuration,omitempty"`
	AllowOptOut           bool      `xml:"allowOptOut,omitempty"`
	AllowedLoadPerPort    float64   `xml:"allowedLoadPerPort,omitempty"`
	PercentShedPerStation float64   `xml:"percentShedPerStation,omitempty"`
}

// Chargepointservices all services available
type Chargepointservices interface {
	ClearAlarms(request *ClearAlarms) (*ClearAlarmsResponse, error)
	ClearAlarmsContext(ctx context.Context, request *ClearAlarms) (*ClearAlarmsResponse, error)
	ClearShedState(request *ClearShedState) (*ClearShedStateResponse, error)
	ClearShedStateContext(ctx context.Context, request *ClearShedState) (*ClearShedStateResponse, error)
	CreateLoadShedEvent(request *CreateLoadShedEvent) (*CreateLoadShedEventResponse, error)
	CreateLoadShedEventContext(ctx context.Context, request *CreateLoadShedEvent) (*CreateLoadShedEventResponse, error)
	Get15minChargingSessionData(request *Get15minChargingSessionData) (*Get15minChargingSessionDataResponse, error)
	Get15minChargingSessionDataContext(ctx context.Context, request *Get15minChargingSessionData) (*Get15minChargingSessionDataResponse, error)
	GetAlarms(request *GetAlarms) (*GetAlarmsResponse, error)
	GetAlarmsContext(ctx context.Context, request *GetAlarms) (*GetAlarmsResponse, error)
	GetCPNInstances(request *GetCPNInstances) (*GetCPNInstancesResponse, error)
	GetCPNInstancesContext(ctx context.Context, request *GetCPNInstances) (*GetCPNInstancesResponse, error)
	GetChargingSessionData(request *GetChargingSessionData) (*GetChargingSessionDataResponse, error)
	GetChargingSessionDataContext(ctx context.Context, request *GetChargingSessionData) (*GetChargingSessionDataResponse, error)
	GetLoad(request *GetLoad) (*GetLoadResponse, error)
	GetLoadContext(ctx context.Context, request *GetLoad) (*GetLoadResponse, error)
	GetOrgsAndStationGroups(request *GetOrgsAndStationGroups) (*GetOrgsAndStationGroupsResponse, error)
	GetOrgsAndStationGroupsContext(ctx context.Context, request *GetOrgsAndStationGroups) (*GetOrgsAndStationGroupsResponse, error)
	GetPublicStationStatus(request *GetPublicStationStatus) (*GetPublicStationStatusResponse, error)
	GetPublicStationStatusContext(ctx context.Context, request *GetPublicStationStatus) (*GetPublicStationStatusResponse, error)
	GetPublicStations(request *GetPublicStations) (*GetPublicStationsResponse, error)
	GetPublicStationsContext(ctx context.Context, request *GetPublicStations) (*GetPublicStationsResponse, error)
	GetStationGroupDetails(request *GetStationGroupDetails) (*GetStationGroupDetailsResponse, error)
	GetStationGroupDetailsContext(ctx context.Context, request *GetStationGroupDetails) (*GetStationGroupDetailsResponse, error)
	GetStationGroups(request *GetStationGroups) (*GetStationGroupsResponse, error)
	GetStationGroupsContext(ctx context.Context, request *GetStationGroups) (*GetStationGroupsResponse, error)
	GetStationRights(request *GetStationRights) (*GetStationRightsResponse, error)
	GetStationRightsContext(ctx context.Context, request *GetStationRights) (*GetStationRightsResponse, error)
	GetStationRightsProfile(request *GetStationRightsProfile) (*GetStationRightsProfileResponse, error)
	GetStationRightsProfileContext(ctx context.Context, request *GetStationRightsProfile) (*GetStationRightsProfileResponse, error)
	GetStationStatus(request *GetStationStatus) (*GetStationStatusResponse, error)
	GetStationStatusContext(ctx context.Context, request *GetStationStatus) (*GetStationStatusResponse, error)
	GetStations(request *GetStations) (*GetStationsResponse, error)
	GetStationsContext(ctx context.Context, request *GetStations) (*GetStationsResponse, error)
	GetTransactionData(request *GetTransactionData) (*GetTransactionDataResponse, error)
	GetTransactionDataContext(ctx context.Context, request *GetTransactionData) (*GetTransactionDataResponse, error)
	GetUsers(request *GetUsers) (*GetUsersResponse, error)
	GetUsersContext(ctx context.Context, request *GetUsers) (*GetUsersResponse, error)
	RegisterFeeds(request *RegisterFeeds) (*RegisterFeedsResponse, error)
	RegisterFeedsContext(ctx context.Context, request *RegisterFeeds) (*RegisterFeedsResponse, error)
	RemoveLoadShedEvent(request *RemoveLoadShedEvent) (*RemoveLoadShedEventResponse, error)
	RemoveLoadShedEventContext(ctx context.Context, request *RemoveLoadShedEvent) (*RemoveLoadShedEventResponse, error)
	SetWaitlistDone(request *SetWaitlistDone) (*SetWaitlistDoneResponse, error)
	SetWaitlistDoneContext(ctx context.Context, request *SetWaitlistDone) (*SetWaitlistDoneResponse, error)
	ShedLoad(request *ShedLoad) (*ShedLoadResponse, error)
	ShedLoadContext(ctx context.Context, request *ShedLoad) (*ShedLoadResponse, error)
	UpdateFeed(request *UpdateFeed) (*UpdateFeedResponse, error)
	UpdateFeedContext(ctx context.Context, request *UpdateFeed) (*UpdateFeedResponse, error)
	UpdateUserStatus(request *UpdateUserStatus) (*UpdateUserStatusResponse, error)
	UpdateUserStatusContext(ctx context.Context, request *UpdateUserStatus) (*UpdateUserStatusResponse, error)
}

type chargepointservices struct {
	client *Client
}

// NewChargepointV5Client is a new chargepoint service
func NewChargepointV5Client(user string, password string) Chargepointservices {
	c := NewSoapClient(user, password, URI)
	return &chargepointservices{
		client: c,
	}
}

func (service *chargepointservices) GetPublicStationsContext(ctx context.Context, request *GetPublicStations) (*GetPublicStationsResponse, error) {
	response := new(GetPublicStationsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getPublicStations", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *chargepointservices) GetPublicStations(request *GetPublicStations) (*GetPublicStationsResponse, error) {
	return service.GetPublicStationsContext(
		context.Background(),
		request,
	)
}

func (service *chargepointservices) GetPublicStationStatusContext(ctx context.Context, request *GetPublicStationStatus) (*GetPublicStationStatusResponse, error) {
	response := new(GetPublicStationStatusResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getPublicStationStatus", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *chargepointservices) GetPublicStationStatus(request *GetPublicStationStatus) (*GetPublicStationStatusResponse, error) {
	return service.GetPublicStationStatusContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetStationStatusContext(ctx context.Context, request *GetStationStatus) (*GetStationStatusResponse, error) {
	response := new(GetStationStatusResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getStationStatus", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetStationStatus(request *GetStationStatus) (*GetStationStatusResponse, error) {
	return service.GetStationStatusContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetAlarmsContext(ctx context.Context, request *GetAlarms) (*GetAlarmsResponse, error) {
	response := new(GetAlarmsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getAlarms", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetAlarms(request *GetAlarms) (*GetAlarmsResponse, error) {
	return service.GetAlarmsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) ClearAlarmsContext(ctx context.Context, request *ClearAlarms) (*ClearAlarmsResponse, error) {
	response := new(ClearAlarmsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/clearAlarms", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) ClearAlarms(request *ClearAlarms) (*ClearAlarmsResponse, error) {
	return service.ClearAlarmsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetOrgsAndStationGroupsContext(ctx context.Context, request *GetOrgsAndStationGroups) (*GetOrgsAndStationGroupsResponse, error) {
	response := new(GetOrgsAndStationGroupsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getOrgsAndStationGroups", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetOrgsAndStationGroups(request *GetOrgsAndStationGroups) (*GetOrgsAndStationGroupsResponse, error) {
	return service.GetOrgsAndStationGroupsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) ShedLoadContext(ctx context.Context, request *ShedLoad) (*ShedLoadResponse, error) {
	response := new(ShedLoadResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/shedLoad", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) ShedLoad(request *ShedLoad) (*ShedLoadResponse, error) {
	return service.ShedLoadContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) ClearShedStateContext(ctx context.Context, request *ClearShedState) (*ClearShedStateResponse, error) {
	response := new(ClearShedStateResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/clearShedState", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) ClearShedState(request *ClearShedState) (*ClearShedStateResponse, error) {
	return service.ClearShedStateContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetLoadContext(ctx context.Context, request *GetLoad) (*GetLoadResponse, error) {
	response := new(GetLoadResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getLoad", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetLoad(request *GetLoad) (*GetLoadResponse, error) {
	return service.GetLoadContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetStationGroupsContext(ctx context.Context, request *GetStationGroups) (*GetStationGroupsResponse, error) {
	response := new(GetStationGroupsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getStationGroups", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetStationGroups(request *GetStationGroups) (*GetStationGroupsResponse, error) {
	return service.GetStationGroupsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetStationGroupDetailsContext(ctx context.Context, request *GetStationGroupDetails) (*GetStationGroupDetailsResponse, error) {
	response := new(GetStationGroupDetailsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getStationGroupDetails", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetStationGroupDetails(request *GetStationGroupDetails) (*GetStationGroupDetailsResponse, error) {
	return service.GetStationGroupDetailsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetCPNInstancesContext(ctx context.Context, request *GetCPNInstances) (*GetCPNInstancesResponse, error) {
	response := new(GetCPNInstancesResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getCPNInstances", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetCPNInstances(request *GetCPNInstances) (*GetCPNInstancesResponse, error) {
	return service.GetCPNInstancesContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetChargingSessionDataContext(ctx context.Context, request *GetChargingSessionData) (*GetChargingSessionDataResponse, error) {
	response := new(GetChargingSessionDataResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getChargingSessionData", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetChargingSessionData(request *GetChargingSessionData) (*GetChargingSessionDataResponse, error) {
	return service.GetChargingSessionDataContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) Get15minChargingSessionDataContext(ctx context.Context, request *Get15minChargingSessionData) (*Get15minChargingSessionDataResponse, error) {
	response := new(Get15minChargingSessionDataResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/get15minChargingSessionData", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) Get15minChargingSessionData(request *Get15minChargingSessionData) (*Get15minChargingSessionDataResponse, error) {
	return service.Get15minChargingSessionDataContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetStationsContext(ctx context.Context, request *GetStations) (*GetStationsResponse, error) {
	response := new(GetStationsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getStations", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetStations(request *GetStations) (*GetStationsResponse, error) {
	return service.GetStationsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetStationRightsContext(ctx context.Context, request *GetStationRights) (*GetStationRightsResponse, error) {
	response := new(GetStationRightsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getStationRights", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetStationRights(request *GetStationRights) (*GetStationRightsResponse, error) {
	return service.GetStationRightsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetStationRightsProfileContext(ctx context.Context, request *GetStationRightsProfile) (*GetStationRightsProfileResponse, error) {
	response := new(GetStationRightsProfileResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getStationRightsProfile", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetStationRightsProfile(request *GetStationRightsProfile) (*GetStationRightsProfileResponse, error) {
	return service.GetStationRightsProfileContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) RegisterFeedsContext(ctx context.Context, request *RegisterFeeds) (*RegisterFeedsResponse, error) {
	response := new(RegisterFeedsResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/registerFeeds", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) RegisterFeeds(request *RegisterFeeds) (*RegisterFeedsResponse, error) {
	return service.RegisterFeedsContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) UpdateFeedContext(ctx context.Context, request *UpdateFeed) (*UpdateFeedResponse, error) {
	response := new(UpdateFeedResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/updateFeed", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) UpdateFeed(request *UpdateFeed) (*UpdateFeedResponse, error) {
	return service.UpdateFeedContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetUsersContext(ctx context.Context, request *GetUsers) (*GetUsersResponse, error) {
	response := new(GetUsersResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getUsers", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetUsers(request *GetUsers) (*GetUsersResponse, error) {
	return service.GetUsersContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) UpdateUserStatusContext(ctx context.Context, request *UpdateUserStatus) (*UpdateUserStatusResponse, error) {
	response := new(UpdateUserStatusResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/updateUserStatus", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) UpdateUserStatus(request *UpdateUserStatus) (*UpdateUserStatusResponse, error) {
	return service.UpdateUserStatusContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) SetWaitlistDoneContext(ctx context.Context, request *SetWaitlistDone) (*SetWaitlistDoneResponse, error) {
	response := new(SetWaitlistDoneResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/setWaitlistDone", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) SetWaitlistDone(request *SetWaitlistDone) (*SetWaitlistDoneResponse, error) {
	return service.SetWaitlistDoneContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) GetTransactionDataContext(ctx context.Context, request *GetTransactionData) (*GetTransactionDataResponse, error) {
	response := new(GetTransactionDataResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/getTransactionData", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) GetTransactionData(request *GetTransactionData) (*GetTransactionDataResponse, error) {
	return service.GetTransactionDataContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) CreateLoadShedEventContext(ctx context.Context, request *CreateLoadShedEvent) (*CreateLoadShedEventResponse, error) {
	response := new(CreateLoadShedEventResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/createLoadShedEvent", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) CreateLoadShedEvent(request *CreateLoadShedEvent) (*CreateLoadShedEventResponse, error) {
	return service.CreateLoadShedEventContext(
		context.Background(),
		request,
	)
}
func (service *chargepointservices) RemoveLoadShedEventContext(ctx context.Context, request *RemoveLoadShedEvent) (*RemoveLoadShedEventResponse, error) {
	response := new(RemoveLoadShedEventResponse)
	err := service.client.CallContext(ctx, "urn:provider/interface/chargepointservices/removeLoadShedEvent", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *chargepointservices) RemoveLoadShedEvent(request *RemoveLoadShedEvent) (*RemoveLoadShedEventResponse, error) {
	return service.RemoveLoadShedEventContext(
		context.Background(),
		request,
	)
}
