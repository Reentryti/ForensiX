package model

import "time"

// List of different severity level
type Level string

const (
	LowLevel    Level = "low"
	MediumLevel Level = "medium"
	HighLevel   Level = "high"
)

// Type of forensic event
type EventType string

const (
	EventAuth        EventType = "authentification"
	EventExecution   EventType = "execution"
	EventPersistence EventType = "persistence"
	EventFileSystem  EventType = "filesystem"
	EventNetwork     EventType = "network"
	EventProcess     EventType = "process"
	Eventuser        EventType = "user"
	EventSystem      EventType = "system"
	EventPackage     EventType = "package"
	EventDevice      EventType = "device"
)

// List of involved actor
type Actor struct {
	User string
	UID  int
	IP   string
}

// Event representation structure
type Event struct {
	EventID   string
	Type      EventType
	Timestamp time.Time
	Artefact  string
	Severity  Level
	Actor     Actor
	Action    string
	Result    string
}
