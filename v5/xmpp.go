package v5

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	xmpp "github.com/mattn/go-xmpp"
)

// XMPPFeedEvent are feeds that can be registered for the xmpp server
type XMPPFeedEvent string

// Note that not all feeds are available based on your chargepoint API plan
const (
	StationProvisioned           XMPPFeedEvent = "station_provisioned"
	StationUnprovisioned         XMPPFeedEvent = "station_unprovisioned"
	StationHeadSwap              XMPPFeedEvent = "station_head_swap"
	StationHiddenVisibleToggle   XMPPFeedEvent = "station_hidden_visible_toggle"
	StationTokenChange           XMPPFeedEvent = "station_token_change"
	StationPlanWarrantyChange    XMPPFeedEvent = "station_plan_warranty_change"
	StationAddressChange         XMPPFeedEvent = "station_address_change"
	StationNameChange            XMPPFeedEvent = "station_name_change"
	StationTimezoneChange        XMPPFeedEvent = "station_timezone_change"
	StationDescriptionChange     XMPPFeedEvent = "station_description_change"
	StationUsageStatusChange     XMPPFeedEvent = "station_usage_status_change"
	StationAlarm                 XMPPFeedEvent = "station_alarm"
	StationACLApplied            XMPPFeedEvent = "station_acl_applied"
	StationACLRemoved            XMPPFeedEvent = "station_acl_removed"
	StationPricingApplied        XMPPFeedEvent = "station_pricing_applied"
	StationPricingremoved        XMPPFeedEvent = "station_pricing_removed"
	StationReservationEnabled    XMPPFeedEvent = "station_reservation_enabled"
	StationReservationDisabled   XMPPFeedEvent = "station_reservation_disabled"
	StationChargingSessionStart  XMPPFeedEvent = "station_charging_session_start"
	StationChargingSessionStop   XMPPFeedEvent = "station_charging_session_stop"
	StationChargingSessionUpdate XMPPFeedEvent = "station_charging_session_update"
)

// XMPPClient is an XMPP listener
type XMPPClient struct {
	Server   string
	User     string
	Password string
	Messages chan Event
	talk     *xmpp.Client
	ctx      context.Context
}

// Event represents an event message from the chargepoint API
type Event struct {
	XMLName       xml.Name      `xml:"event"`
	Text          string        `xml:",chardata"`
	StationID     string        `xml:"stationID"`
	FeedEventName XMPPFeedEvent `xml:"feedEventName"`
	SessionID     string        `xml:"sessionID"`
	StartTime     string        `xml:"startTime"`
	PortNumber    string        `xml:"portNumber"`
	RfID          string        `xml:"rfID"`
	UserID        string        `xml:"userID"`
}

// NewXMPPClient returns a new xmpp client listener
func NewXMPPClient(ctx context.Context, server string, user string, password string, debug bool) (XMPPClient, error) {
	if !strings.HasSuffix("@webservice.chargepointportal.net", user) {
		user = fmt.Sprintf("%s%s", user, "@webservice.chargepointportal.net")
	}

	c := XMPPClient{
		Server:   server,
		User:     user,
		Password: password,
		Messages: make(chan Event, 10),
		ctx:      ctx,
	}

	xmpp.DefaultConfig = tls.Config{
		ServerName:         strings.Split(c.Server, ":")[0],
		InsecureSkipVerify: false,
	}

	// Note, the chargepoint API doesn't pass TLS
	options := xmpp.Options{
		Host:     c.Server,
		User:     c.User,
		Password: c.Password,
		NoTLS:    true,
		Debug:    debug,
	}

	talk, err := options.NewClient()
	if err != nil {
		return XMPPClient{}, err
	}
	c.talk = talk

	return c, nil
}

// StartReader will start reading from the xmpp channel and write data to a channel
func (x *XMPPClient) StartReader() {
	defer x.talk.Close()

	for {
		select {
		case <-x.ctx.Done():
			return
		default:
			chat, err := x.talk.Recv()
			if err != nil {
				log.Fatal(err)
			}

			switch v := chat.(type) {
			case xmpp.Chat:
				resp := Event{}
				err := xml.Unmarshal([]byte(v.Text), &resp)
				if err != nil {
					log.Fatal(err)
				}
				x.Messages <- resp
			}
		}
	}
}
