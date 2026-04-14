package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//////////////////// 1 ////////////////////

type Student struct {
	Name  string
	Age   int
	Group string
}

func (s Student) Display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Group:", s.Group)
}

//////////////////// 2 + 6 ////////////////////

func (s Student) SaveToFile(filename string) error {
	// Проверки
	if s.Age < 0 {
		return errors.New("age cannot be negative")
	}
	if s.Name == "" {
		return errors.New("name cannot be empty")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := fmt.Sprintf("Name: %s\nAge: %d\nGroup: %s\n", s.Name, s.Age, s.Group)

	_, err = file.WriteString(data)
	return err
}

//////////////////// 3 ////////////////////

func ReadFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File content:\n", string(data))
}

//////////////////// 4 ////////////////////

type Book struct {
	Title  string
	Author string
}

func (b Book) Display() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
}

func (b Book) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := fmt.Sprintf("Title: %s\nAuthor: %s\n", b.Title, b.Author)

	_, err = file.WriteString(data)
	return err
}

//////////////////// 5 ////////////////////

type FileManager struct {
	Filename string
}

func (fm FileManager) Write(data string) error {
	file, err := os.Create(fm.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func (fm FileManager) Read() {
	data, err := os.ReadFile(fm.Filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File content:\n", string(data))
}

//////////////////// 7 ////////////////////

type Note struct {
	Text string
}

func main() {

	// 🟢 Задание 1
	student := Student{
		Name:  "Aurik",
		Age:   20,
		Group: "IT-101",
	}
	student.Display()

	// 🟡 2
	err := student.SaveToFile("student.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Student saved successfully")
	}

	//3
	ReadFromFile("student.txt")

	//4
	book := Book{
		Title:  "Go Basics",
		Author: "John Doe",
	}
	book.Display()

	err = book.SaveToFile("book.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	ReadFromFile("book.txt")

	// 5
	fm := FileManager{Filename: "data.txt"}

	data := fmt.Sprintf("Student: %s, Age: %d", student.Name, student.Age)

	err = fm.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fm.Read()

	// 7
	reader := bufio.NewScanner(os.Stdin)

	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("Enter notes (type 'exit' to stop):")

	for {
		reader.Scan()
		text := reader.Text()

		if text == "exit" {
			break
		}

		file.WriteString(text + "\n")
	}

	fmt.Println("\nAll notes:")

	dataNotes, err := os.ReadFile("notes.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(dataNotes))
}