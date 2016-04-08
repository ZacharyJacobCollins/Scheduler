package main

//TODO write a check to to see how many things the line includes.  Should be x number.  If not send email to update code.
//TODO set download all to execute at 2 am every night.
//TODO send me an email with the error output.
//TODO make monday - sunday a map
//TODO make a map for seasons courses
//TODO more efficient search
//TODO send me an email with the error output.


import (
	"log"
	"github.com/ZacharyJacobCollins/Scheduler/services"
	"github.com/ZacharyJacobCollins/Scheduler/models"
	"fmt"
	"strings"
)

var fall   []models.Course
var summer []models.Course
var winter []models.Course

func main() {
	//services.DownloadFiles()
	fall, summer, winter = services.LoadSemesters()
	log.Print(fall[100].Category)
	prompt()
}

func prompt() {
	fmt.Print("Enter a professor you're looking for to see their location: ")
	var professor string
	fmt.Scan(&professor)
	//fmt.Print("Enter the class number of the course you're looking for: ")
	//var classNumber string
	//fmt.Scan(&classNumber)
	find(professor, fall)
}

func find(professor string, set []models.Course) {
	for _, course := range set {
		if (strings.TrimSpace(course.Professor) == professor) {
			fmt.Println("Here you are! ", course)
		}
	}
}

func check(e error) {
	if e != nil {
		log.Print(e)
	}
}
