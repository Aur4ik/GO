package main

import "fmt"

type Participant interface {
	GetInfo()
	Act()
}

//=============юзер================
type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}

//=============геттеры================
func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int {
	return u.age
}

//=============сеттеры================
func (u *User) SetName(name string) error {
	if name != "" {
		u.name = name
		return nil
	} else {
		return fmt.Errorf("Имя не может быть пустым !")
	}
}
func (u *User) SetAge(age int) error {
	if age >= 0{
		u.age = age
		return nil
	} else {
		return fmt.Errorf("Возраст не может быть отрицательным")
	}
}

//=============Курсы================

type Course struct {
	Title         string
	MaxScore      int
	StudentsCount int
	MaxStudents int
}

func (c Course) GetInfo() string {
	return fmt.Sprintf(
		"Course: %s | Max Score: %d | Students: %d",
		c.Title,
		c.MaxScore,
		c.StudentsCount,
	)
}

func (c *Course) AddStudent() error {
	if c.StudentsCount < c.MaxStudents {
		c.StudentsCount++
		return nil
	}
	return fmt.Errorf("достигнут лимит студентов")
}

//=============Студент================
type Student struct {
	User
	Score int
	Course *Course
}

func (s *Student) Study() {
	if s.Score < s.Course.MaxScore {
		s.Score += 10
		if s.Score > s.Course.MaxScore {
			s.Score = s.Course.MaxScore
		}
	}
}

func (s *Student) GetInfo() {
	fmt.Printf(
		"Student: %s | Score: %d | Course: %s\n",
		s.GetName(),
		s.Score,
		s.Course.Title,
	)
}

//=============Учитель================
type Teacher struct {
	User
	Subject string
}

func (t *Teacher) Teach() {
	fmt.Printf("Teacher %s is teaching %s\n", t.GetName(), t.Subject)
}

func (t Teacher) GetInfo() {
	fmt.Printf(
		"Teacher: %s | Subject: %s\n",
		t.GetName(),
		t.Subject,
	)
}
func (t *Teacher) GradeStudent(s *Student, points int) {
	s.Score += points
	if s.Score > s.Course.MaxScore {
		s.Score = s.Course.MaxScore
	}
}
func AverageScore(students []*Student) float64 {
	if len(students) == 0 {
		return 0
	}

	total := 0

	for _, s := range students {
		total += s.Score
	}

	return float64(total) / float64(len(students))
}



func BestStudent(students []*Student) *Student {
	if len(students) == 0 {
		return nil
	}

	best := students[0]

	for _, s := range students {
		if s.Score > best.Score {
			best = s
		}
	}

	return best
}


//=============Полиморфизм================
func (s *Student) Act() {
	s.Study()
}

func (t Teacher) Act() {
	t.Teach()
}

func main() {


	course := &Course{
		Title:       "Go Backend",
		MaxScore:    100,
		MaxStudents: 3,
	}


	s1 := &Student{
		User:   *NewUser("Aurik", 20),
		Course: course,
	}

	s2 := &Student{
		User:   *NewUser("Ali", 19),
		Course: course,
	}

	s3 := &Student{
		User:   *NewUser("John", 22),
		Course: course,
	}


	t1 := &Teacher{
		User:    *NewUser("Mr. Smith", 40),
		Subject: "Golang",
	}


	students := []*Student{s1, s2, s3}

	for range students {
		if err := course.AddStudent(); err != nil {
			fmt.Println("Error:", err)
		}
	}


	fmt.Println("=== COURSE INFO ===")
	fmt.Println(course.GetInfo())


	participants := []Participant{s1, s2, s3, t1}

	fmt.Println("\n=== PARTICIPANTS ===")
	for _, p := range participants {
		p.GetInfo()
	}

	fmt.Println("\n=== ACTIONS ===")
	for _, p := range participants {
		p.Act()
	}


	fmt.Println("\n=== GRADING ===")
	t1.GradeStudent(s1, 20)
	t1.GradeStudent(s2, 15)
	t1.GradeStudent(s3, 30)


	fmt.Println("\n=== AFTER ACTIONS ===")
	for _, p := range participants {
		p.GetInfo()
	}


	avg := AverageScore(students)
	fmt.Printf("\nAverage score: %.2f\n", avg)


	best := BestStudent(students)
	if best != nil {
		fmt.Printf("Best student: %s (%d points)\n", best.GetName(), best.Score)
	}
}
