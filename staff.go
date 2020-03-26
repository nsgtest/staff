package main

import(
	"os"
	"fmt"
	"github.com/nsgtest/packages/structs"
)

type Colleague struct{
	Abbreviation, Gender, Forename, Surname, Email string
	Subjects string
}

func main(){
	if len(os.Args) > 1{
		switch os.Args[1]{
			case "init":
				if len(os.Args) == 3{
					s := structs.Struct{os.Args[2], Colleague{}}
					s.Write()
				} else {
					help()
				}
			case "add":
				if len(os.Args) == 9{
					s := structs.Struct{os.Args[2], Colleague{os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7], os.Args[8]}}
					s.Add([]int{0,2,4,6,8,10,32,34,36,38,40,42})
				} else {
					help()
				}
			case "update":
				if len(os.Args) == 9{
					s := structs.Struct{os.Args[2], Colleague{os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7], os.Args[8]}}
					s.Update([]int{17,19,21,23,25,27,29,31,47,49,51,53,55,57,59,61,62,63})
				} else {
					help()
				}
			case "remove":
				if len(os.Args) == 4{
					s := structs.Struct{os.Args[2], Colleague{os.Args[3], "", "", "", "", ""}}
					s.Remove([]int{1})
				} else {
					help()
				}
			case "list":
				if len(os.Args) == 3{
					s := structs.Struct{os.Args[2], Colleague{}}
					s.List()
				} else {
					help()
				}
			default:
				help()
		}
	} else {
		help()
	}
}

func (c Colleague) Interface(){
	fmt.Println("Smells Like Teen Spirit!")
}

func help(){
	fmt.Println("Help!")
}
