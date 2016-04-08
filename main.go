package main

//TODO write a check to to see how many things the line includes.  Should be x number.  If not send email to update code.
//TODO set download all to execute at 2 am every night.
//TODO send me an email with the error output.
//TODO make monday - sunday a map
//TODO make a map for seasons courses

import (
	"log"
	"github.com/ZacharyJacobCollins/Scheduler/services"
	"github.com/ZacharyJacobCollins/Scheduler/models"
)

var fall   []models.Course
var summer []models.Course
var winter []models.Course

func main() {
	//services.DownloadFiles()
	fall, summer, winter = services.LoadSemesters()
	log.Print(fall[369].Title)
}


//TODO send me an email with the error output.
func check(e error) {
	if e != nil {
		log.Print(e)
	}
}
