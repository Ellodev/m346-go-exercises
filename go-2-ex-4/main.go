package main

import "fmt"

type Student struct {
	Firstname string
	Lastname  string
}

type Class struct {
	Students []Student
}

func main() {
	s1 := Student{Firstname: "John", Lastname: "Doe"}
	s2 := Student{Firstname: "Hans", Lastname: "Muster"}
	s3 := Student{Firstname: "Paul", Lastname: "Bucherer"}

	s4 := Student{Firstname: "Patrick", Lastname: "Müller"}
	s5 := Student{Firstname: "Aline", Lastname: "Lustenberger"}
	s6 := Student{Firstname: "Pascal", Lastname: "Meier"}

	classA := Class{
		Students: []Student{s1, s2, s3},
	}

	classB := Class{
		Students: []Student{s4, s5, s6},
	}

	moduleNames := map[int]string{
		104: "Datenmodell implementieren",
		117: "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren",
		346: "Cloud Lösungen konzipieren und realisieren",
	}

	moduleAllocations := map[int][]Class{
		104: {classA},
		117: {classA, classB},
		346: {classB},
	}

	fmt.Println(classA)
	fmt.Println(classB)
	fmt.Println(moduleNames)
	fmt.Println(moduleAllocations)
}
