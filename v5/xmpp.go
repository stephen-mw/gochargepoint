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

// XMPPFeedEventName are feeds that can be registered for the xmpp server
type XMPPFeedEventName string

// Note that not all feeds are available based on your chargepoint API plan
const (
	StationProvisioned           XMPPFeedEventName = "station_provisioned"
	StationUnprovisioned         XMPPFeedEventName = "station_unprovisioned"
	StationHeadSwap              XMPPFeedEventName = "station_head_swap"
	StationHiddenVisibleToggle   XMPPFeedEventName = "station_hidden_visible_toggle"
	StationTokenChange           XMPPFeedEventName = "station_token_change"
	StationPlanWarrantyChange    XMPPFeedEventName = "station_plan_warranty_change"
	StationAddressChange         XMPPFeedEventName = "station_address_change"
	StationNameChange            XMPPFeedEventName = "station_name_change"
	StationTimezoneChange        XMPPFeedEventName = "station_timezone_change"
	StationDescriptionChange     XMPPFeedEventName = "station_description_change"
	StationUsageStatusChange     XMPPFeedEventName = "station_usage_status_change"
	StationAlarm                 XMPPFeedEventName = "station_alarm"
	StationACLApplied            XMPPFeedEventName = "station_acl_applied"
	StationACLRemoved            XMPPFeedEventName = "station_acl_removed"
	StationPricingApplied        XMPPFeedEventName = "station_pricing_applied"
	StationPricingremoved        XMPPFeedEventName = "station_pricing_removed"
	StationReservationEnabled    XMPPFeedEventName = "station_reservation_enabled"
	StationReservationDisabled   XMPPFeedEventName = "station_reservation_disabled"
	StationChargingSessionStart  XMPPFeedEventName = "station_charging_session_start"
	StationChargingSessionStop   XMPPFeedEventName = "station_charging_session_stop"
	StationChargingSessionUpdate XMPPFeedEventName = "station_charging_session_update"
)

// XMPPClient is an XMPP listener
type XMPPClient struct {
	Server   string
	User     string
	Password string
	Messages chan XMPPEvent
	talk     *xmpp.Client
	ctx      context.Context
}

// XMPPEvent represents an event message from the chargepoint API
type XMPPEvent struct {
	XMLName        xml.Name          `xml:"event"`
	EndTime        string            `xml:"endTime"`
	FeedEventName  XMPPFeedEventName `xml:"feedEventName"`
	FragmentNumber float32           `xml:"fragmentNumber"`
	NetEnergy      float32           `xml:"netEnergy"`
	PortNumber     string            `xml:"portNumber"`
	RfID           string            `xml:"rfID"`
	SessionID      string            `xml:"sessionID"`
	StartTime      string            `xml:"startTime"`
	StationID      string            `xml:"stationID"`
	StationTime    string            `xml:"stationTime"`
	Status         string            `xml:"status"`
	Text           string            `xml:",chardata"`
	UserID         string            `xml:"userID"`
	Raw            string            `xml:"-"`
}

// NewXMPPClient returns a new xmpp client listener
func NewXMPPClient(ctx context.Context, server string, user string, password string, ch chan XMPPEvent, debug bool) (XMPPClient, error) {
	if !strings.HasSuffix("@webservice.chargepointportal.net", user) {
		user = fmt.Sprintf("%s%s", user, "@webservice.chargepointportal.net")
	}

	c := XMPPClient{
		Server:   server,
		User:     user,
		Password: password,
		Messages: ch,
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
				resp := XMPPEvent{}
				resp.Raw = v.Text
				err := xml.Unmarshal([]byte(v.Text), &resp)
				if err != nil {
					log.Fatal(err)
				}
				x.Messages <- resp
			}
		}
	}
}
