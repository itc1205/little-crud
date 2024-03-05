package logger

type LogMessage struct {
	Id          int32
	ProjectId   int32
	Name        string
	Description string
	Priority    int32
	Removed     bool
	EventTime   string
}
