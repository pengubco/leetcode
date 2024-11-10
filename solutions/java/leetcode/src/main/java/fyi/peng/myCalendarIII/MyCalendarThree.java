package fyi.peng.myCalendarIII;

import java.util.TreeSet;

public class MyCalendarThree {

    TreeSet<EventTime> eventTimes; 
    public MyCalendarThree() {
        this.eventTimes = new TreeSet<>((a, b) -> {
            if (a.time() != b.time()) {
                return a.time() - b.time();
            }
            if (!a.isStart()) {
                return -1;
            }
            return 1;
        });
    }
    
    public int book(int startTime, int endTime) {
        eventTimes.add(new EventTime(startTime, true));
        eventTimes.add(new EventTime(endTime, false));

        // scan-line
        int maxOverlap = 0;
        int overlap = 0;
        for (var e : this.eventTimes) {
            if (e.isStart()) {
                overlap++;
                maxOverlap = Math.max(maxOverlap, overlap);
            } else {
                overlap--;
            }
        }
        return maxOverlap;
    }
}

record EventTime(int time, boolean isStart) {
}
