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

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// Demo DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("My API")
	router := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "Java", CoursePrice: 69, Author: &Author{FullName: "Fullname", Website: "localhost"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Go", CoursePrice: 420, Author: &Author{FullName: "Naam", Website: "127.0.0.1"}})

	//routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", router))
}

//controllers -file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	fmt.Println("Params:", params)

	// loop through the courses, find the matching id sent and then return the reponse
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given courseId:" + params["id"])
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Insert Course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body not sent")
		return
	}

	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Request body is empty")
		return
	}

	//generate unique courseId, string conversiion, append the new course to the courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab the courseId from the request
	params := mux.Vars(r)

	// loop through the courses and remove the course from the courses, update and add new course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	//if course is not found
	json.NewEncoder(w).Encode("Course not found with course Id " + params["id"])
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "appliction/json")

	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode("Course removed from courses")
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	return
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete AllCourses")
	w.Header().Set("Content-Type", "application/json")
	courses = nil
	json.NewEncoder(w).Encode("All Courses deleted")
}
