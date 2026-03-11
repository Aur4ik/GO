package main

import "fmt"

// 1
// type Book struct {
// 	Title       string
// 	Author      string
// 	pages       int
// 	isAvailible bool
// }

// func newBook(Title string, Author string, pages int, isAvailible bool) *Book {
// 	return &Book{
// 		Title:       Title,
// 		Author:      Author,
// 		pages:       pages,
// 		isAvailible: isAvailible,
// 	}
// }
// func (b Book) Info() string {
// 	return fmt.Sprintf("Title: %s, Author: %s, Pages: %d, Available: %t",
// 		b.Title, b.Author, b.pages, b.isAvailible)
// }
// func (b *Book) Borrow() {
// 	if !b.isAvailible {
// 		fmt.Println("Книга уже выдана")
// 		return
// 	}

// 	b.isAvailible = false
// 	fmt.Println("Книга выдана")
// }
// func (b *Book) returnBook(){
// 	if b.isAvailible == false {
// 		fmt.Println("Вы не брали эту книгу или уже ее вернули")
// 		return
// 	}
// 	fmt.Println("Вы вернули книгу")
// 	b.isAvailible = false
// }
// func (b Book) returnPages() int {
// 	return b.pages
// }
// func (b *Book) setPage(p int) int{
// 	if(p < 0){
// 		fmt.Println("Страниц не может быть ментше одной")
// 		return b.pages
// 	}
// 	b.pages += p
// 	return b.pages
// }

// type Worker interface{
// 	Work() string
// 	GetName() string
// }

// type Programmer struct{
// 	Name string
// 	Language string
// }
// func (p Programmer) Work() string{
// 	return fmt.Sprintf("Programmer %s use %s", p.Name, p.Language)

// }
// func (p Programmer) GetName() string{
// 	return p.Name
// }
// type Designer struct{
// 	Name string
// 	Tool string
// }
// func (d Designer) Work() string{
// 	return fmt.Sprintf("Programmer %s use %s", d.Name, d.Tool)
// }
// func (d Designer) GetName() string{
// 	return d.Name
// }
// func ShowWork(w Worker){
//     fmt.Println(w.GetName())
//     fmt.Println(w.Work())
// }

type Product struct {
	Name     string
	price    float64
	Quantity int
}

func NewProduct(Name string, price float64, Quantity int) *Product {

	if(price < 0){
		price = 0
	}

	return &Product{
		Name:     Name,
		price:    price,
		Quantity: Quantity,
	}
}
func (p Product)GetProduct() string{
	return fmt.Sprintf("Product: %s,  price: %f, Quantity: %d", p.Name, p.price, p.Quantity)
}
func(p *Product) SetPrice(newPrice float64) float64{
	if newPrice < 0 {
		newPrice = 0
	}
	p.price = newPrice
	return p.price
}
func (p *Product)Buy(amount int){
	if amount < 0 {
		fmt.Println("Error, amount cant be a negative")
		return
	}
	if p.Quantity < amount {
		fmt.Println("Товара нету в таком количестве")
		return
	}
	p.Quantity -= amount
	fmt.Println("Вы купили товар, осталось :", p.Quantity)
}
func(p *Product)Restock(amount int){
	if amount < 0 {
		amount = 0
	}
	p.Quantity += amount
	fmt.Println("Товар пополнен на ", amount)
	fmt.Println("Сейчас на складе ", p.Quantity, " шт")
}
func(p Product)Info(){
	fmt.Println("Название:",p.Name)
	fmt.Println("цена:",p.price)
	fmt.Println("кол-во:",p.Quantity)
}
func main() {
	NewProduct("qweret", -111, 1)


	// p := Programmer{
	// 	Name: "Ali",
	// 	Language: "Java",
	// }
	// d := Designer{
	// 	Name: "Ivan",
	// 	Tool: "Figma",
	// }

	// // fmt.Println(p.GetName())
	// // fmt.Println(d.GetName())
	// // p.Work()
	// // d.Work()
	// ShowWork(p)
	// ShowWork(d)

}
