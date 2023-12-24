package main

import "fmt"

type book struct {
	title string
}

type studyItem interface {
	printName()
}

func (b *book) printName() {
	fmt.Printf(b.title)
}

/* func (b *book) printName() {
	fmt.Printf(b.title)
}
 */
func callStudyItem(si studyItem) {
	si.printName()
	fmt.Println(" -> Study Item called")
}

func main() {
	b1 := book{
		title: "freemind",
	}

	var b2 *book 
	b2 = &book{
		title: "mangoose",
	}

	fmt.Println(b1)
	fmt.Printf("%T \n",b1)

	fmt.Println(b2)
	fmt.Printf("%T \n",b2)

	//callStudyItem(b1)
	callStudyItem(b2)
}