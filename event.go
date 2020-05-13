package ical

import (
	"log"
	"reflect"
	"strings"
	"time"
)

// Event struct type is a iCalendar component, defined
// as VEVENT in RFC 5545
type Event struct {
	// Required
	UID     string
	DTSTAMP string
	// Required only if Calendar object does not
	// specify the METHOD property
	DTSTART string
	DTEND   string
	// Optional
	SUMMARY     string
	CLASS       string
	CREATED     string
	DESCRIPTION string
	GEO         string
	LASTMOD     string // LAST-MOD
	LOCATION    string
	ORGANIZER   string // TODO: SPECIAL TYPE: CAL_ADDRESS
	PRIORITY    string
	SEQ         string
	STATUS      string
	TRANSP      string
	URL         string
	RECURID     string // TODO: SPECIAL TYPE!
	RRULE       string // TODO: SPECIAL TYPE!
	// Optional but should not be declared together
	DURATION string // TODO: SPECIAL TYPE PT0H0M0S
	// Optional and supports multiple declarations
	ATTACH     []string
	ATTENDEE   []string // TODO: SPECIAL TYPE: CAL_ADDRESS
	CATEGORIES []string
	COMMENT    []string
	CONTACT    []string
	EXDATE     []string
	RSTATUS    []string
	RELATED    []string
	RESOURCES  []string
	RDATE      []string
	XPROP      []string // X-PROP
	IANAPROP   []string // IANA-PROP
}

var eventComponents = []string{"UID", "DTSTAMP", "DTSTART", "DTEND", "SUMMARY", "CLASS", "CREATED", "DESCRIPTION", "GEO", "LASTMOD", "LOCATION",
	"ORGANIZER", "PRIORITY", "SEQ", "STATUS", "TRANSP", "URL", "RECURID", "RRULE", "DURATION", "ATTACH",
	"ATTENDEE", "CATEGORIES", "COMMENT", "CONTACT", "EXDATE", "RSTATUS", "RELATED", "RESOURCES", "RDATE", "XPROP", "IANAPROP"}

// NewEvent creates an instance of struct Event
func NewEvent() *Event {

	// Creates a new instance
	e := new(Event)

	// Get timestamp
	currentTimestamp := FormatDateTime(time.Now())

	// Assign struct values
	e.DTSTAMP = currentTimestamp

	return e
}

func (e *Event) isReady() bool {
	if e.UID == "" {
		return false
	}
	if e.DTSTAMP == "" {
		return false
	}
	if e.DTSTART == "" {
		return false
	}
	if e.DTEND != "" && e.DURATION != "" {
		return false
	}
	return true
}

// GenerateEventProp method creates .ics contents
func (e *Event) GenerateEventProp() string {
	// Validate first
	status := e.isReady()
	if !status {
		log.Fatal("Event is not ready!")
	}

	// Create object
	var str strings.Builder

	// Write headers
	str.WriteString("BEGIN:VEVENT\r\n")

	// Required params will be taken from event automatically.
	for _, item := range eventComponents {
		v := reflect.ValueOf(*e)
		itemValue := v.FieldByName(item)
		isNotValid := itemValue.IsZero() || itemValue.Kind() == reflect.Ptr && itemValue.IsNil()
		if !isNotValid {
			str.WriteString(item + ":" + itemValue.String() + "\r\n")
			log.Println(item + ":" + itemValue.String() + "\r\n")
		}

	}

	// Write footers
	str.WriteString("END:VEVENT\r\n")

	return str.String()
}
