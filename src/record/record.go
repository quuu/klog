package record

import (
	"errors"
)

type Record interface {
	Date() Date

	Summary() string
	SetSummary(string)

	Entries() []Entry
	AddDuration(Duration)
	AddRange(Range)
	OpenRange() OpenRangeStart
	StartOpenRange(OpenRangeStart)
	EndOpenRange(Time) error
}

func NewRecord(date Date) Record {
	return &record{
		date: date,
	}
}

type record struct {
	date    Date
	summary string
	entries []Entry
}

func (r *record) Date() Date {
	return r.date
}

func (r *record) Summary() string {
	return r.summary
}

func (r *record) SetSummary(summary string) {
	r.summary = summary
}

func (r *record) Entries() []Entry {
	return r.entries
}

func (r *record) Durations() []Duration {
	var durations []Duration
	for _, e := range r.entries {
		d, isDuration := e.Value().(Duration)
		if isDuration {
			durations = append(durations, d)
		}
	}
	return durations
}

func (r *record) AddDuration(d Duration) {
	r.entries = append(r.entries, entry{value: d, summary: ""})
}

func (r *record) Ranges() []Range {
	var ranges []Range
	for _, e := range r.entries {
		tr, isRange := e.Value().(Range)
		if isRange {
			ranges = append(ranges, tr)
		}
	}
	return ranges
}

func (r *record) AddRange(tr Range) {
	r.entries = append(r.entries, entry{value: tr, summary: ""})
}

func (r *record) OpenRange() OpenRangeStart {
	for _, e := range r.entries {
		t, isStartTime := e.Value().(Time)
		if isStartTime {
			return t
		}
	}
	return nil
}

func (r *record) StartOpenRange(t OpenRangeStart) {
	r.entries = append(r.entries, entry{value: t, summary: ""})
}

func (r *record) EndOpenRange(end Time) error {
	for i, e := range r.entries {
		t, isStartTime := e.Value().(Time)
		if isStartTime {
			tr, err := NewRange(t, end)
			if err != nil {
				return err
			}
			r.entries[i] = entry{value: tr, summary: ""}
			return nil
		}
	}
	return errors.New("NO_OPEN_RANGE")
}