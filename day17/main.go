package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// creating models for courses

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `jsosn:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullaname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middleware .. helper

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Course management with golang")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "reactjs", CoursePrice: 344, Author: &Author{Fullname: "Bikram", Website: "hello.com.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")

	r.HandleFunc("/coures", getAllCourses).Methods("GET")

	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	r.HandleFunc("/course",createOneCourse).Methods("POST")

	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")

	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELET")




	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api development with golang </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from any request
	params := mux.Vars(r)

	// loop through courses , find mathing id and return the respsonse

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id ")
	return

}

func createOneCourse(w http.ResponseController, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEecoder(w).Encode("Please send some data")
	}

	// what about  {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data insid json")
	}

	// generate unique id, string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update  course")
	w.Header().Set("Content-Type", "application/json")

	// get id for req
	params := mux.Vars(r)
	// loop , id , remov ,add with id in params

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one  course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1]...)
			break
		}
	}

}

// routing
