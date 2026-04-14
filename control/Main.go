package main
import "fmt"
type City struct {
	city string
}

type Student struct {
	City
	name string
	age  int
}

type Students struct {
	students []Student
}

func (s *Students) AddStudent(student Student) {
	s.students = append(s.students, student)
}

func (s Students) GetAllStudents() []Student {
	return s.students
}

func (s Students) FindStudent(name string) *Student {
	for i := 0; i < len(s.students); i++ {
		if s.students[i].name == name {
			return &s.students[i]
		}
	}
	return nil
}

func (s *Students) DeleteStudent(name string, age int) *Student {
	for i := range s.students {
		if s.students[i].name == name && s.students[i].age == age {
			removed := s.students[i]
			s.students = append(s.students[:i], s.students[i+1:]...)
			return &removed
		}
	}
	return nil
}

func main() {
	allStudents := Students{}

	for {
		var choice int
		fmt.Println("\nМеню:")
		fmt.Println("1 - Добавить студента")
		fmt.Println("2 - Показать всех студентов")
		fmt.Println("3 - Найти студента")
		fmt.Println("4 - Удалить студента")
		fmt.Println("0 - Выход")
		fmt.Print("Выберите действие: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name, city string
			var age int
			fmt.Print("Введите имя студента: ")
			fmt.Scanln(&name)
			fmt.Print("Введите возраст студента: ")
			fmt.Scanln(&age)
			fmt.Print("Введите город студента: ")
			fmt.Scanln(&city)

			allStudents.AddStudent(Student{
				name: name,
				age:  age,
				City: City{city: city},
			})
			fmt.Println("Студент добавлен!")

		case 2:
			fmt.Println("\nВсе студенты:")
			for _, student := range allStudents.GetAllStudents() {
				fmt.Printf("Имя: %s, Возраст: %d, Город: %s\n", student.name, student.age, student.City.city)
			}

		case 3:
			var name string
			fmt.Print("Введите имя студента для поиска: ")
			fmt.Scanln(&name)
			found := allStudents.FindStudent(name)
			if found != nil {
				fmt.Printf("Найден студент: %s, Возраст: %d, Город: %s\n", found.name, found.age, found.City.city)
			} else {
				fmt.Println("Студент не найден")
			}

		case 4:
			var name string
			var age int
			fmt.Print("Введите имя студента для удаления: ")
			fmt.Scanln(&name)
			fmt.Print("Введите возраст студента для удаления: ")
			fmt.Scanln(&age)
			deleted := allStudents.DeleteStudent(name, age)
			if deleted != nil {
				fmt.Printf("Студент %s удалён!\n", deleted.name)
			} else {
				fmt.Println("Студент не найден для удаления")
			}

		case 0:
			fmt.Println("Выход...")
			return

		default:
			fmt.Println("Неверная команда, попробуйте снова")
		}
	}
}