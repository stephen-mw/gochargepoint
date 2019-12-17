package v5

import (
	"encoding/xml"
	"testing"
	"time"
)

// TestStationChargingSessionUpdate will test the StationChargingSessionUpdate event
func TestStationChargingSessionUpdate(t *testing.T) {
	msg := []byte(`<?xml version="1.0" encoding="UTF-8"?><event>  <stationID><![CDATA[1:124923]]></stationID>  <feedEventName><![CDATA[station_charging_session_update]]></feedEventName>  <sessionID><![CDATA[505543971]]></sessionID>  <fragmentNumber><![CDATA[1]]></fragmentNumber>  <netEnergy><![CDATA[0.825367]]></netEnergy>  <energy><![CDATA[0.825367]]></energy>  <stationTime><![CDATA[2019-08-07 01:51:44]]></stationTime>  <portNumber><![CDATA[2]]></portNumber></event>`)

	r := Event{}
	err := xml.Unmarshal(msg, &r)
	if err != nil {
		t.Fatal(err)
	}
	if r.StationID != "1:124923" {
		t.Fatal("StationID incorrectly")
	}
	if r.FeedEventName != StationChargingSessionUpdate {
		t.Fatal("Feed event name set incorrectly")
	}
	if r.SessionID != "505543971" {
		t.Fatal("SessionID incorrectly")
	}
	if r.FragmentNumber != 1 {
		t.Fatal("FragmentNumber incorrect")
	}
	if r.NetEnergy != 0.825367 {
		t.Fatal("NetEnergy incorrect")
	}
	if r.StationTime != "2019-08-07 01:51:44" {
		t.Fatal("StationTIme incorrect")
	}
	if r.PortNumber != "2" {
		t.Fatal("Port number incorrect")
	}

	ts, err := time.Parse("2006-01-02 15:04:05", r.StationTime)
	if err != nil {
		t.Fatal(err)
	}
	date, err := r.GetEventTime()
	if err != nil {
		t.Fatal(err)
	}
	if date != ts {
		t.Fatalf("bad timestamps. Expected %s but got %s", ts, date)
	}
}

// TestStationUsageStatusChange will test the StationUsageStatusChange event
func TestStationUsageStatusChange(t *testing.T) {
	msg := []byte(`<?xml version="1.0" encoding="UTF-8"?><event>  <stationID><![CDATA[1:116027]]></stationID>  <feedEventName><![CDATA[station_usage_status_change]]></feedEventName>  <portNumber><![CDATA[2]]></portNumber>  <status><![CDATA[AVAILABLE]]></status></event>`)

	r := Event{}
	err := xml.Unmarshal(msg, &r)
	if err != nil {
		t.Fatal(err)
	}
	if r.StationID != "1:116027" {
		t.Fatal("StationID incorrectly")
	}
	if r.FeedEventName != StationUsageStatusChange {
		t.Fatal("Feed event name set incorrectly")
	}

	date, err := r.GetEventTime()
	if err != nil {
		t.Fatal(err)
	}
	if !date.IsZero() {
		t.Fatalf("Timestamp should be zero")
	}
}

// TestStationChargingSessionStop will test the StationChargingSessionStart event
func TestStationChargingSessionStop(t *testing.T) {
	msg := []byte(`<?xml version="1.0" encoding="UTF-8"?><event>  <stationID><![CDATA[1:122945]]></stationID>  <feedEventName><![CDATA[station_charging_session_stop]]></feedEventName>  <sessionID><![CDATA[4007]]></sessionID>  <startTime><![CDATA[2019-08-06 19:57:47]]></startTime>  <endTime><![CDATA[2019-08-07 02:00:00]]></endTime>  <portNumber><![CDATA[1]]></portNumber></event>`)

	r := Event{}
	err := xml.Unmarshal(msg, &r)
	if err != nil {
		t.Fatal(err)
	}

	if r.StationID != "1:122945" {
		t.Fatal("StationID not set correctly")
	}
	if r.FeedEventName != StationChargingSessionStop {
		t.Fatal("Event name set incorrectly")
	}
	if r.SessionID != "4007" {
		t.Fatal("SessionID not set correct")
	}
	if r.StartTime != "2019-08-06 19:57:47" {
		t.Fatal("Startime not set")
	}
	if r.EndTime != "2019-08-07 02:00:00" {
		t.Fatal("Endtime not set correctly")
	}
	if r.PortNumber != "1" {
		t.Fatal("Port number not set correctly")
	}

	date, err := r.GetEventTime()
	if err != nil {
		t.Fatal(err)
	}
	if date.IsZero() {
		t.Fatalf("Bad timestamp: %s", date)
	}
}

// TestStationChargingSessionStart will test the StationChargingSessionStart event
func TestStationChargingSessionStart(t *testing.T) {
	msg := []byte(`<?xml version="1.0" encoding="UTF-8"?><event>  <stationID><![CDATA[1:116027]]></stationID>  <feedEventName><![CDATA[station_charging_session_start]]></feedEventName>  <sessionID><![CDATA[411]]></sessionID>  <startTime><![CDATA[2019-08-07T01:40:37Z]]></startTime>  <portNumber><![CDATA[1]]></portNumber>  <rfID><![CDATA[e007808be26f0b51]]></rfID>  <userID><![CDATA[171131]]></userID></event>`)

	r := Event{}
	err := xml.Unmarshal(msg, &r)
	if err != nil {
		t.Fatal(err)
	}

	if r.StationID != "1:116027" {
		t.Fatal("StationID not set")
	}
	if r.FeedEventName != StationChargingSessionStart {
		t.Fatal("Feed event name not set correctly")
	}
	if r.SessionID != "411" {
		t.Fatal("SessionID not set correctly")
	}
	if r.StartTime != "2019-08-07T01:40:37Z" {
		t.Fatal("Startime not set")
	}
	if r.EndTime != "" {
		t.Fatal("Endtime should not be set")
	}
	if r.PortNumber != "1" {
		t.Fatal("Port number not set")
	}

	date, err := r.GetEventTime()
	if err != nil {
		t.Fatal(err)
	}
	if date.IsZero() {
		t.Fatalf("Bad timestamp: %s", date)
	}
}
