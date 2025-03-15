package main 

import "fmt"

type Student struct {
	Name string;
	Age int;
	Marks []int;
}


func main() {

	var student1 Student
	var numSubj int

	fmt.Print("Enter Student Name: ")
	fmt.Scan(&student1.Name)

	fmt.Print("Enter Student Age: ")
	fmt.Scan(&student1.Age)

	fmt.Print("Enter Number of Subjects: ")
	fmt.Scan(&numSubj)

	student1.Marks = make([]int, numSubj)

	for i:=0; i < numSubj; i++ {
		fmt.Print("Enter Subject Marks: ")
		fmt.Scanln(&student1.Marks[i])
	}

	fmt.Printf("This is average marks of %s: %.2f\n",student1.Name, averageMarks(student1))
}

func averageMarks(student1 Student) float64{

	var Total = 0

	for i,_:= range student1.Marks {
		Total+=student1.Marks[i]
	}
	return float64(Total)/float64((len(student1.Marks)))
}