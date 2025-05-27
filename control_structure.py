import random

shifts = ["morning", "afternoon", "evening"]
days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"]

class Employee:
    def __init__(self, name, preferences):
        self.name = name
        self.preferences = preferences  # dict: day -> shift
        self.assigned = {}              # dict: day -> assigned shift
        self.days_worked = 0

def assign_shifts(employees):
    schedule = {day: {shift: [] for shift in shifts} for day in days}

    for day in days:
        for emp in employees:
            if emp.days_worked >= 5 or day in emp.assigned:
                continue
            pref_shift = emp.preferences.get(day)
            if pref_shift and len(schedule[day][pref_shift]) < 2:
                schedule[day][pref_shift].append(emp.name)
                emp.assigned[day] = pref_shift
                emp.days_worked += 1
            else:
                for alt in shifts:
                    if len(schedule[day][alt]) < 2:
                        schedule[day][alt].append(emp.name)
                        emp.assigned[day] = alt
                        emp.days_worked += 1
                        break

        # Fill remaining slots randomly
        for shift in shifts:
            while len(schedule[day][shift]) < 2:
                candidates = [e for e in employees if e.days_worked < 5 and day not in e.assigned]
                if not candidates:
                    break
                chosen = random.choice(candidates)
                schedule[day][shift].append(chosen.name)
                chosen.assigned[day] = shift
                chosen.days_worked += 1

    return schedule

def print_schedule(schedule):
    for day in days:
        print(f"{day}:")
        for shift in shifts:
            print(f"  {shift}: {schedule[day][shift]}")

# Static employee input
employees = [
    Employee("Alice", {
        "Monday": "morning", "Tuesday": "afternoon", "Wednesday": "evening",
        "Thursday": "morning", "Friday": "afternoon", "Saturday": "evening", "Sunday": "morning"
    }),
    Employee("Bob", {
        "Monday": "afternoon", "Tuesday": "morning", "Wednesday": "afternoon",
        "Thursday": "evening", "Friday": "morning", "Saturday": "afternoon", "Sunday": "evening"
    }),
    Employee("Charlie", {
        "Monday": "evening", "Tuesday": "evening", "Wednesday": "morning",
        "Thursday": "afternoon", "Friday": "evening", "Saturday": "morning", "Sunday": "afternoon"
    }),
]

schedule = assign_shifts(employees)
print_schedule(schedule)
