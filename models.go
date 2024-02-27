package main

type Event interface {
	display()
}

type Log struct {
	ID     int
	Source string
	Body   string
}

type SystemLog struct {
	Log
	Severity string
}

func (l Log) display() {
	println(l.ID, l.Source, l.Body)
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

	for _, event := range events {
		event.display()
	}

}
