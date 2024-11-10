package mycalendari

type MyCalendar struct {
	events []event
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(startTime int, endTime int) bool {
	for _, e := range c.events {
		if endTime <= e.startTime || e.endTime <= startTime {
			continue
		}
		return false
	}
	c.events = append(c.events, event{startTime, endTime})
	return true
}

type event struct {
	startTime int
	endTime   int
}
