package main

type Event interface {
	display()
}

// Base struct
type Log struct {
	ID     int
	Source string
	Body   string
}

type SystemLog struct {
	Log
	Severity string
}

// implement display method for SystemLog
func (s SystemLog) display() {
	println(s.ID, s.Source, s.Body)
}

func main() {
	var events []Event

	slog := SystemLog{
		Log: Log{
			ID:     1,
			Source: "System",
			Body:   "System is running",
		},
		Severity: "Info",
	}

	events = append(events, slog)

	slog = SystemLog{
		Log{2, "System", "System is running"},
		"Error",
	}

	events = append(events, slog)

	for _, event := range events {
		event.display()
	}

}
