package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Roll    int
	Subject []subject
}
type subject struct {
	sub   string
	marks int
}

func grade() {
	person := []Person{}
	var count int
	var count1 int
	var t int
	//taking the value
	fmt.Printf("if you want to input new entry press 1 if not press any key : ")
	fmt.Scanf("%d", &t)
	if t == 1 {
		fmt.Printf("Enter the number of student : ")
		fmt.Scanf("%d", &count)
		for p := 0; p < count; p++ {
			per := Person{}
			fmt.Println("#####################################")
			fmt.Printf("Enter student Name : ")
			fmt.Scanf("%s", &per.Name)
			fmt.Printf("Enter the roll of : ")
			fmt.Scanf("%d", &per.Roll)
			fmt.Printf("Enter the number of subject enrolled : ")
			fmt.Scanf("%d", &count1)
			for k := 0; k < count1; k++ {
				subject := subject{}
				fmt.Printf("Enter the name of %d subject : ", k+1)
				fmt.Scanf("%s", &subject.sub)
				fmt.Printf("Enter marks of %s : ", subject.sub)
				fmt.Scanf("%d", &subject.marks)
				per.Subject = append(per.Subject, subject)
			}
			person = append(person, per)
			fmt.Println("#######################################")

		}

	} else {
		fmt.Println("Thank you")
	}

	//grading
	var roll int
	fmt.Println("Enter the roll to see result")
	fmt.Scanf("%d", &roll)
	for _, value := range person {
		if roll == value.Roll {
			fmt.Println("#######################################")
			fmt.Println("Name : ", value.Name)
			fmt.Println("Roll : ", value.Roll)
			for _, value := range value.Subject {
				fmt.Printf("Subject : %s  Marks : %d", value.sub, value.marks)

				if value.marks <= 100 && value.marks >= 80 {
					fmt.Println(" Grade : A+")
				} else if value.marks <= 79 && value.marks >= 70 {
					fmt.Println(" Grade : A")
				} else if value.marks <= 69 && value.marks >= 60 {
					fmt.Println(" Grade : A-")
				} else if value.marks <= 59 && value.marks >= 50 {
					fmt.Println(" Grade : B")
				} else if value.marks <= 49 && value.marks >= 40 {
					fmt.Println(" Grade : C")
				} else {
					fmt.Println(" Grade : F")
				}

			}

			fmt.Println("###########################")

		}else{
			fmt.Println("The roll is not there")
		}
	}

}
