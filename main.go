package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Day 4
// [1518-05-19 00:01] Guard #1801 begins shift

type Entry struct {
	Time time.Time
	Text string
}

type Guard struct {
	ID int
	Histogram [60]int // Minutes in the midnight hour
}

func (g Guard) timeAsleep() int {
	return sum(g.Histogram[:])
}

func (g Guard) mostCommonSleepTime() (int, int) {
	return max(g.Histogram[:])
}

type OrderedEntries []Entry
func (a OrderedEntries) Len() int { return len(a) }
func (a OrderedEntries) Swap(i, j int)  { a[i], a[j] = a[j], a[i] }
func (a OrderedEntries) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }


type OrderedBySleep []Guard
func (a OrderedBySleep) Len() int { return len(a) }
func (a OrderedBySleep) Swap(i, j int)  { a[i], a[j] = a[j], a[i] }
func (a OrderedBySleep) Less(i, j int) bool { return a[i].timeAsleep() < a[j].timeAsleep() }

type OrderedByMinuteAsleep []Guard
func (a OrderedByMinuteAsleep) Len() int { return len(a) }
func (a OrderedByMinuteAsleep) Swap(i, j int)  { a[i], a[j] = a[j], a[i] }
func (a OrderedByMinuteAsleep) Less(i, j int) bool {
	w, _ := a[i].mostCommonSleepTime()
	y, _ := a[j].mostCommonSleepTime()

	return w < y
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	entries := parse(f)

	guards := make(map[string]*Guard)
	currentGuardId := ""
	var currentGuard *Guard
	var ok bool
	begin, end := -1, -1

	for _, v := range entries {
		// Guard Identifier
		if v.Text[0:5] == "Guard" {
			// Change of guard, reset begin and end times
			begin, end = -1, -1
			currentGuardId = strings.Split(v.Text[7:], " ")[0]
			currentGuard, ok = guards[currentGuardId]
			if !ok {
				id, err := strconv.Atoi(currentGuardId)
				if err != nil {
					fmt.Println(err)
				}
				currentGuard = &Guard{ID: id}
				guards[currentGuardId] = currentGuard
			}
		}

		// Falls asleep
		if v.Text[0:5] == "falls" {
			begin = v.Time.Minute()
		}

		if v.Text[0:5] == "wakes" {
			end = v.Time.Minute()
		}

		if begin != -1 && end != -1 {
			for i := begin; i<end; i++ {
				currentGuard.Histogram[i] += 1
			}
			begin, end = -1, -1
		}
	}

	guardSlice := make([]Guard, 0)

	for _, v := range guards {
		guardSlice = append(guardSlice, *v)
	}

	// Part 1
	sort.Sort(OrderedBySleep(guardSlice))
	g := guardSlice[len(guardSlice) - 1]
	times, minute := max(g.Histogram[:])
	fmt.Println("Guard", g.ID, minute, times, g.ID * minute)

	// Part 2
	sort.Sort(OrderedByMinuteAsleep(guardSlice))
	g = guardSlice[len(guardSlice) - 1]
	times, minute = max(g.Histogram[:])
	fmt.Println("Guard", g.ID, minute, times, g.ID * minute)
}

func max(i []int) (int, int) {
	m, index := 0, 0
	for k, v := range i {
		if v > m {m = v; index = k}
	}
	return m, index
}

func sum(i []int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

func parse(f io.Reader) []Entry {
	entries := make(OrderedEntries, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		layout := "2006-01-02 15:04"
		t, err := time.Parse(layout, s[1:17])
		if err != nil {
			fmt.Println(err)
		}
		entries = append(entries, Entry{Time: t, Text: s[19:]})
	}
	sort.Sort(OrderedEntries(entries))
	return entries
}
