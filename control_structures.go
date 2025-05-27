package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Employee struct {
	Name        string
	Preferences map[string]string // preferred shift per day
	Assigned    map[string]string // actual assigned shift per day
	DaysWorked  int
}

var shifts = []string{"morning", "afternoon", "evening"}
var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func assignShifts(employees []*Employee) map[string]map[string][]string {
	schedule := make(map[string]map[string][]string)

	// Initialize empty schedule
	for _, day := range days {
		schedule[day] = map[string][]string{
			"morning":   {},
			"afternoon": {},
			"evening":   {},
		}
	}

	rand.Seed(time.Now().UnixNano())

	for _, day := range days {
		for _, emp := range employees {
			if emp.DaysWorked >= 5 || emp.Assigned[day] != "" {
				continue
			}
			// Assign preferred shift if available
			if shift, ok := emp.Preferences[day]; ok && len(schedule[day][shift]) < 2 {
				schedule[day][shift] = append(schedule[day][shift], emp.Name)
				emp.Assigned[day] = shift
				emp.DaysWorked++
			} else {
				// Try other available shifts
				for _, alt := range shifts {
					if len(schedule[day][alt]) < 2 {
						schedule[day][alt] = append(schedule[day][alt], emp.Name)
						emp.Assigned[day] = alt
						emp.DaysWorked++
						break
					}
				}
			}
		}

		// Fill remaining slots with random eligible employees
		for _, shift := range shifts {
			for len(schedule[day][shift]) < 2 {
				candidates := []*Employee{}
				for _, emp := range employees {
					if emp.DaysWorked < 5 && emp.Assigned[day] == "" {
						candidates = append(candidates, emp)
					}
				}
				if len(candidates) == 0 {
					break
				}
				chosen := candidates[rand.Intn(len(candidates))]
				schedule[day][shift] = append(schedule[day][shift], chosen.Name)
				chosen.Assigned[day] = shift
				chosen.DaysWorked++
			}
		}
	}

	return schedule
}

func main() {
	// Static employee input
	employees := []*Employee{
		{
			Name: "Alice",
			Preferences: map[string]string{
				"Monday": "morning", "Tuesday": "afternoon", "Wednesday": "evening",
				"Thursday": "morning", "Friday": "afternoon", "Saturday": "evening", "Sunday": "morning",
			},
			Assigned:   make(map[string]string),
			DaysWorked: 0,
		},
		{
			Name: "Bob",
			Preferences: map[string]string{
				"Monday": "afternoon", "Tuesday": "morning", "Wednesday": "afternoon",
				"Thursday": "evening", "Friday": "morning", "Saturday": "afternoon", "Sunday": "evening",
			},
			Assigned:   make(map[string]string),
			DaysWorked: 0,
		},
		{
			Name: "Charlie",
			Preferences: map[string]string{
				"Monday": "evening", "Tuesday": "evening", "Wednesday": "morning",
				"Thursday": "afternoon", "Friday": "evening", "Saturday": "morning", "Sunday": "afternoon",
			},
			Assigned:   make(map[string]string),
			DaysWorked: 0,
		},
	}

	schedule := assignShifts(employees)

	// Output final schedule
	for _, day := range days {
		fmt.Printf("%s:\n", day)
		for _, shift := range shifts {
			fmt.Printf("  %s: %v\n", shift, schedule[day][shift])
		}
	}
}
