package services

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"io/ioutil"

	"github.com/ZacharyJacobCollins/Scheduler/models"
)


func DownloadFiles() {
	urls := []string{
		"http://it.emich.edu/bannerfiles/sctnenrl_fa.txt",
		"http://it.emich.edu/bannerfiles/sctnenrl_wi.txt",
		"http://it.emich.edu/bannerfiles/sctnenrl_su.txt",
	}

	for _, u := range urls {
		go download(u)
	}
}

func download(url string) {
	tokens := strings.Split(url, "/")
	filename := tokens[len(tokens)-1]
	r, err := http.Get(url)
	check(err)
	defer r.Body.Close()
	file, err := os.Create("files/" + filename)
	check(err)
	defer file.Close()
	n, err := io.Copy(file, r.Body)
	check(err)

	log.Println(n, " bytes downloaded at ", time.Now(), " into file ", filename)
}

func LoadSemesters() ([]models.Course, []models.Course, []models.Course) {
	fall   := loadCourses("files/sctnenrl_fa.txt")
	summer := loadCourses("files/sctnenrl_su.txt")
	winter := loadCourses("files/sctnenrl_wi.txt")
	return fall, summer, winter
}

func loadCourses(file string) (courses []models.Course) {
	text, err := ioutil.ReadFile(file)
	check(err)
	stringText := string(text)
	lines := strings.Split(stringText, "\n")
	for i, l := range lines {
		if (i < len(lines)-1) {
			//splits slice of lines, into slice of strings(data) seperated by ;
			data := strings.Split(l, ";")
			//export data slice which is a slice of strings from a single line
			courses = parseLine(data, courses)
		}
	}
	return courses
}

func parseLine(data []string, courses []models.Course) ([]models.Course) {
	courses = append(courses, models.Course{
		data[0],  data[3],  data[4],
		data[5],  data[9],  data[15],
		data[16], data[17], data[18],
		data[19], data[20], data[21],
		data[22], data[23], data[24],
		data[25], data[26], data[28],
	})
	return courses
}

func check(e error) {
	if e != nil {
		log.Print(e)
	}
}