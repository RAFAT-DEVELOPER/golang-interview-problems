# Minimum Platforms Required

This repository contains implementations of different algorithms to find the minimum number of platforms required for a railway/bus station. This problem is commonly encountered in scheduling and logistics.

## Problem Statement

Given the arrival and departure times of trains/buses at a station, the goal is to determine the minimum number of platforms needed to ensure that no train/bus has to wait before departing.

## Example

For the example provided:

``` go
arrivals := []int{900, 940, 950, 1100, 1500, 1800}
departures := []int{910, 1200, 1120, 1130, 1900, 2000}
```
The expected output is 3, indicating that a minimum of 3 platforms are required to accommodate all trains/buses without delays.

Implemented Approaches
This repository includes implementations of the following approaches to solve the problem:

1. Naive Approach:

This approach takes every interval one by one and counts the number of overlapping intervals.
Time Complexity: O(n^2)
Space Complexity: O(1)

2. Heap Approach:

Uses a min heap to keep track of departure times.
Sorts the events based on time and processes them, maintaining the maximum number of overlapping intervals.
Time Complexity: O(n log n)
Space Complexity: O(n)

3. Sorting Approach :

Sorts the events based on time and processes them, maintaining the maximum number of overlapping intervals.
Time Complexity: O(n log n)
Space Complexity: O(n)

4. Sweep Line Algorithm:

Sorts the events based on time and processes them, maintaining the maximum number of overlapping intervals.
Time Complexity: O(n log n)
Space Complexity: O(n)

5. Set-Based Approach:

Sorts the events based on time, adding arrivals to a set and removing departures.
Maintains the maximum number of overlapping intervals by counting elements in the set.
Time Complexity: O(n log n)
Space Complexity: O(n)
How to Use
You can use the provided code implementations to find the minimum number of platforms required for your own dataset by calling the respective functions and passing your arrival and departure times.


## Example usage for Heap Approach
``` 
platformsHeap := minPlatformsRequiredHeap(arrivals, departures)
```
Feel free to explore and choose the approach that best suits your specific use case based on factors like dataset size and efficiency requirements.

## Summary
In summary, all approaches have a time complexity of O(n log n) due to sorting the events, except for the Naive Approach, which has a time complexity of O(n^2). The space complexity for all approaches is O(n) due to the storage of events or other data structures. The Heap and Sorting approaches are efficient for large datasets, while the Set-Based Approach is also efficient and maintains the benefits of a set data structure. The Naive Approach is the least efficient and should be avoided for large datasets. The choice of approach depends on your specific requirements and dataset size.
