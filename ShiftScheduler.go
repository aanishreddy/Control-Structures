package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

type Employee struct {
    Name        string
    Preferences map[string]string
    Assigned    map[string]string
    DaysWorked  int
}

var shifts = []string{"morning", "afternoon", "evening"}
var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func collectInput() []*Employee {
    reader := bufio.NewReader(os.Stdin)
    var employees []*Employee

    fmt.Print("Enter number of employees: ")
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Printf("Enter name of employee %d: ", i+1)
        nameInput, _ := reader.ReadString('\n')
        name := strings.TrimSpace(nameInput)

        prefs := make(map[string]string)
        for _, day := range days {
            fmt.Printf("Enter %s's preferred shift on %s (morning/afternoon/evening): ", name, day)
            shiftInput, _ := reader.ReadString('\n')
            shift := strings.ToLower(strings.TrimSpace(shiftInput))
            prefs[day] = shift
        }

        emp := &Employee{
            Name:        name,
            Preferences: prefs,
            Assigned:    make(map[string]string),
            DaysWorked:  0,
        }
        employees = append(employees, emp)
    }
    return employees
}

func assignShifts(employees []*Employee) map[string]map[string][]string {
    schedule := make(map[string]map[string][]string)
    for _, day := range days {
        schedule[day] = map[string][]string{"morning": {}, "afternoon": {}, "evening": {}}
    }

    rand.Seed(time.Now().UnixNano())
    for _, day := range days {
        for _, emp := range employees {
            if emp.DaysWorked >= 5 || emp.Assigned[day] != "" {
                continue
            }
            if shift, ok := emp.Preferences[day]; ok && len(schedule[day][shift]) < 2 {
                schedule[day][shift] = append(schedule[day][shift], emp.Name)
                emp.Assigned[day] = shift
                emp.DaysWorked++
            } else {
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

func printSchedule(schedule map[string]map[string][]string) {
    for _, day := range days {
        fmt.Println(day + ":")
        for _, shift := range shifts {
            fmt.Printf("  %s: %v\n", shift, schedule[day][shift])
        }
    }
}

func main() {
    employees := collectInput()
    schedule := assignShifts(employees)
    fmt.Println("\nFinal Weekly Schedule:")
    printSchedule(schedule)
}
