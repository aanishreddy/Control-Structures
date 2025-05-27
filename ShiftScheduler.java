import java.util.*;

public class ShiftScheduler {

    static final String[] SHIFTS = {"morning", "afternoon", "evening"};
    static final String[] DAYS = {"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"};

    static class Employee {
        String name;
        Map<String, String> preferences = new HashMap<>(); // day -> preferred shift
        Map<String, String> assigned = new HashMap<>();    // day -> assigned shift
        int daysWorked = 0;

        Employee(String name) {
            this.name = name;
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Employee> employees = new ArrayList<>();

        System.out.print("Enter number of employees: ");
        int n = Integer.parseInt(scanner.nextLine());

        for (int i = 0; i < n; i++) {
            System.out.print("Enter name of employee " + (i + 1) + ": ");
            String name = scanner.nextLine();
            Employee emp = new Employee(name);

            for (String day : DAYS) {
                System.out.print("Preferred shift for " + name + " on " + day + " (morning/afternoon/evening): ");
                String shift = scanner.nextLine().toLowerCase();
                emp.preferences.put(day, shift);
            }

            employees.add(emp);
        }

        Map<String, Map<String, List<String>>> schedule = assignShifts(employees);
        printSchedule(schedule);
    }

    static Map<String, Map<String, List<String>>> assignShifts(List<Employee> employees) {
        Map<String, Map<String, List<String>>> schedule = new LinkedHashMap<>();
        Random random = new Random();

        for (String day : DAYS) {
            Map<String, List<String>> daySchedule = new HashMap<>();
            for (String shift : SHIFTS) {
                daySchedule.put(shift, new ArrayList<>());
            }

            for (Employee emp : employees) {
                if (emp.daysWorked >= 5 || emp.assigned.containsKey(day)) continue;

                String preferred = emp.preferences.getOrDefault(day, "");
                if (daySchedule.get(preferred).size() < 2) {
                    daySchedule.get(preferred).add(emp.name);
                    emp.assigned.put(day, preferred);
                    emp.daysWorked++;
                } else {
                    for (String alt : SHIFTS) {
                        if (daySchedule.get(alt).size() < 2) {
                            daySchedule.get(alt).add(emp.name);
                            emp.assigned.put(day, alt);
                            emp.daysWorked++;
                            break;
                        }
                    }
                }
            }

            // Fill unassigned spots randomly
            for (String shift : SHIFTS) {
                while (daySchedule.get(shift).size() < 2) {
                    List<Employee> candidates = new ArrayList<>();
                    for (Employee e : employees) {
                        if (e.daysWorked < 5 && !e.assigned.containsKey(day)) {
                            candidates.add(e);
                        }
                    }
                    if (candidates.isEmpty()) break;

                    Employee chosen = candidates.get(random.nextInt(candidates.size()));
                    daySchedule.get(shift).add(chosen.name);
                    chosen.assigned.put(day, shift);
                    chosen.daysWorked++;
                }
            }

            schedule.put(day, daySchedule);
        }

        return schedule;
    }

    static void printSchedule(Map<String, Map<String, List<String>>> schedule) {
        System.out.println("\nFinal Weekly Schedule:");
        for (String day : DAYS) {
            System.out.println(day + ":");
            Map<String, List<String>> shifts = schedule.get(day);
            for (String shift : SHIFTS) {
                System.out.println("  " + shift + ": " + shifts.get(shift));
            }
        }
    }
}
