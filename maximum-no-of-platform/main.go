package main

import (
	"fmt"
	"sort"
)

type Event struct {
	time      int
	isArrival bool
}

type EventHeap []Event

func (h EventHeap) Len() int           { return len(h) }
func (h EventHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h EventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EventHeap) Push(x interface{}) {
	*h = append(*h, x.(Event))
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Naive Approach
func minPlatformsRequiredNaive(arrivals, departures []int) int {
	n := len(arrivals)
	platformCount := 0
	minPlatforms := 0

	for i := 0; i < n; i++ {
		platformCount = 1 // At least one platform is required for each interval
		for j := i + 1; j < n; j++ {
			if (arrivals[i] >= arrivals[j] && arrivals[i] <= departures[j]) ||
				(arrivals[j] >= arrivals[i] && arrivals[j] <= departures[i]) {
				platformCount++
			}
		}
		if platformCount > minPlatforms {
			minPlatforms = platformCount
		}
	}

	return minPlatforms
}

// Heap Approach
func minPlatformsRequiredHeap(arrivals, departures []int) int {
	events := make([]Event, 2*len(arrivals))

	for i, time := range arrivals {
		events[i] = Event{time, true}
	}

	for i, time := range departures {
		events[len(arrivals)+i] = Event{time, false}
	}

	sort.Sort(EventHeap(events))
	platformCount := 0
	maxPlatforms := 0

	for _, event := range events {
		if event.isArrival {
			platformCount++
			if platformCount > maxPlatforms {
				maxPlatforms = platformCount
			}
		} else {
			platformCount--
		}
	}

	return maxPlatforms
}

// Sorting Approach (Naive)
func minPlatformsRequiredSort(arrivals, departures []int) int {
	events := make([]Event, 2*len(arrivals))

	for i, time := range arrivals {
		events[i] = Event{time, true}
	}

	for i, time := range departures {
		events[len(arrivals)+i] = Event{time, false}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].isArrival
		}
		return events[i].time < events[j].time
	})

	platformCount := 0
	activeTrains := 0

	for _, event := range events {
		if event.isArrival {
			activeTrains++
			if activeTrains > platformCount {
				platformCount = activeTrains
			}
		} else {
			activeTrains--
		}
	}

	return platformCount
}

// Sweep Line Algorithm
func minPlatformsRequiredSweepLine(arrivals, departures []int) int {
	events := make([]Event, 2*len(arrivals))

	for i, time := range arrivals {
		events[i] = Event{time, true}
	}

	for i, time := range departures {
		events[len(arrivals)+i] = Event{time, false}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].isArrival
		}
		return events[i].time < events[j].time
	})

	platformCount := 0
	activeTrains := 0

	for _, event := range events {
		if event.isArrival {
			activeTrains++
			if activeTrains > platformCount {
				platformCount = activeTrains
			}
		} else {
			activeTrains--
		}
	}

	return platformCount
}

// Set-based Approach
func minPlatformsRequiredSet(arrivals, departures []int) int {
	events := make([]Event, 2*len(arrivals))

	for i, time := range arrivals {
		events[i] = Event{time, true}
	}

	for i, time := range departures {
		events[len(arrivals)+i] = Event{time, false}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			// Handle arrivals before departures to avoid double-counting
			return events[i].isArrival && !events[j].isArrival
		}
		return events[i].time < events[j].time
	})

	platformCount := 0
	minPlatforms := 0

	for _, event := range events {
		if event.isArrival {
			platformCount++
		} else {
			platformCount--
		}

		if platformCount > minPlatforms {
			minPlatforms = platformCount
		}
	}

	return minPlatforms
}

func main() {
	arrivals := []int{900, 940, 950, 1100, 1500, 1800}
	departures := []int{910, 1200, 1120, 1130, 1900, 2000}

	// Naive Approach
	platformsNaive := minPlatformsRequiredNaive(arrivals, departures)
	fmt.Printf("Minimum platforms required (Naive Approach): %d\n", platformsNaive)

	// Heap Approach
	platformsHeap := minPlatformsRequiredHeap(arrivals, departures)
	fmt.Printf("Minimum platforms required (Heap Approach): %d\n", platformsHeap)

	// Sorting Approach (Naive)
	platformsSort := minPlatformsRequiredSort(arrivals, departures)
	fmt.Printf("Minimum platforms required (Sorting Approach): %d\n", platformsSort)

	// Sweep Line Algorithm
	platformsSweepLine := minPlatformsRequiredSweepLine(arrivals, departures)
	fmt.Printf("Minimum platforms required (Sweep Line Algorithm): %d\n", platformsSweepLine)

	// Set-based Approach
	platformsSet := minPlatformsRequiredSet(arrivals, departures)
	fmt.Printf("Minimum platforms required (Set-based Approach): %d\n", platformsSet)
}
