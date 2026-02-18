package main
import "fmt"

type Course struct {
	id	int
	title 	string
	Price 	float64
}

type Student struct {
	name 	string
	Courses  []Course
}

var catalog = make(map[int]Course)
var students = make(map[string]Student)

func addCourse(id int, title string, price float64) {
	course := Course{id, title, price}
	catalog[id] = course
}
func AddStudent(name string) {
	student := Student{name, []Course{}}
	students[name] = student
}

func Enroll(studentName string, courseId int) {
	student, studentExists := students[studentName]
	course, courseExists := catalog[courseId]
	if studentExists && courseExists {
		student.Courses = append(student.Courses, course)
		students[studentName] = student
	}
}

func TotalCost(studentName string) float64 {
	student, studentExists := students[studentName]
	if !studentExists {
		return 0
	}
	total := 0.0
	for _, course := range student.Courses {
		total += course.Price
	}
	return total
}

func PrintStudent(studentName string) {
	student, studentExists := students[studentName]
	if !studentExists {
		fmt.Println("Студент не найден")
		return
	}
	fmt.Printf("Студент: %s\n", student.name)
	fmt.Println("Курсы:")
	for _, course := range student.Courses {
		fmt.Printf("- %s (%.2f)\n", course.title, course.Price)
	}
	fmt.Printf("Общая стоимость: %.2f\n", TotalCost(studentName))
}




func main() {
	addCourse(1, "Go Programming", 100.0)
	addCourse(2, "Python Programming", 80.0)
	addCourse(3, "Java Programming", 90.0)

	AddStudent("Alice")
	AddStudent("Bob")

	Enroll("Alice", 1)
	Enroll("Alice", 2)
	Enroll("Bob", 3)

	PrintStudent("Alice")
	PrintStudent("Bob")
}