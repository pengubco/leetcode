package mycalendarii

type MyCalendarTwo struct {
	doubleBooked []Interval
	booked       []Interval
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(startTime int, endTime int) bool {
	for _, b := range c.doubleBooked {
		if startTime >= b.endTime || b.startTime >= endTime {
			continue
		}
		return false
	}
	for _, b := range c.booked {
		if startTime >= b.endTime || b.startTime >= endTime {
			continue
		}
		c.doubleBooked = append(c.doubleBooked, Interval{
			max(startTime, b.startTime), min(endTime, b.endTime),
		})
	}
	c.booked = append(c.booked, Interval{startTime, endTime})
	return true
}

type Interval struct {
	startTime int
	endTime   int
}
