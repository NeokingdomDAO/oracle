package odoo

import "time"

type Schedule uint

const (
	DailySchedule Schedule = iota
	WeeklySchedule
	MonthlySchedule
)

type Interval struct {
	Start      time.Time
	End        time.Time
	EndInclude time.Time
	Schedule   Schedule
}

func NewInterval(t time.Time, s Schedule) Interval {
	var i Interval
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	switch s {
	case DailySchedule:
		i = Interval{
			Start:      midnight,
			End:        midnight.AddDate(0, 0, 1),
			EndInclude: midnight.AddDate(0, 0, 1).Add(time.Second * -1),
			Schedule:   s,
		}
	case WeeklySchedule:
		weekday := (int(midnight.Weekday()) + 6) % 7
		i = Interval{
			Start:      midnight.AddDate(0, 0, -weekday),
			End:        midnight.AddDate(0, 0, 7-weekday),
			EndInclude: midnight.AddDate(0, 0, 7-weekday).Add(time.Second * -1),
			Schedule:   s,
		}
	case MonthlySchedule:
		firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		i = Interval{
			Start:      firstDay,
			End:        firstDay.AddDate(0, 1, 0),
			EndInclude: firstDay.AddDate(0, 1, -1).Add(time.Second * -1),
			Schedule:   s,
		}
	}
	return i
}

func (i Interval) Increment() Interval {
	var p Interval
	switch i.Schedule {
	case DailySchedule:
		p = Interval{
			Start:      i.Start.AddDate(0, 0, 1),
			End:        i.End.AddDate(0, 0, 1),
			EndInclude: i.EndInclude.AddDate(0, 0, 1),
			Schedule:   i.Schedule,
		}
	case WeeklySchedule:
		p = Interval{
			Start:      i.Start.AddDate(0, 0, 7),
			End:        i.End.AddDate(0, 0, 7),
			EndInclude: i.EndInclude.AddDate(0, 0, 7),
			Schedule:   i.Schedule,
		}
	case MonthlySchedule:
		p = Interval{
			Start:      i.Start.AddDate(0, 1, 0),
			End:        i.End.AddDate(0, 1, 0),
			EndInclude: i.EndInclude.AddDate(0, 1, 0),
			Schedule:   i.Schedule,
		}
	}
	return p
}

func (i Interval) Decrement() Interval {
	var p Interval
	switch i.Schedule {
	case DailySchedule:
		p = Interval{
			Start:      i.Start.AddDate(0, 0, -1),
			End:        i.End.AddDate(0, 0, -1),
			EndInclude: i.EndInclude.AddDate(0, 0, -1),
			Schedule:   i.Schedule,
		}
	case WeeklySchedule:
		p = Interval{
			Start:      i.Start.AddDate(0, 0, -7),
			End:        i.End.AddDate(0, 0, -7),
			EndInclude: i.EndInclude.AddDate(0, 0, -7),
			Schedule:   i.Schedule,
		}
	case MonthlySchedule:
		p = Interval{
			Start:      i.Start.AddDate(0, -1, 0),
			End:        i.End.AddDate(0, -1, 0),
			EndInclude: i.EndInclude.AddDate(0, -1, 0),
			Schedule:   i.Schedule,
		}
	}
	return p
}

func (i Interval) Format() (string, string) {
	return i.Start.Format("2006-01-02"), i.End.Format("2006-01-02")
}

func (i Interval) String() string {
	return ""
}
