package main

import (
	"html/template"
	"log"
	"net/http"
)

type Experience struct {
	Title       string
	Company     string
	Description string
	Period      string
}

type Education struct {
	Degree      string
	Institution string
	Period      string
	GPA         string
}

type Skill struct {
	Category string
	Items    []string
}

type PageData struct {
	Name         string
	Title        string
	Description  string
	Email        string
	Phone        string
	Location     string
	LinkedIn     string
	Website      string
	Experiences  []Experience
	Education    Education
	Skills       []Skill
	Certifications []string
}

func main() {
	http.HandleFunc("/", handleHome)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Name:        "Praveen Potnuri",
		Title:       "Full Stack Engineer",
		Description: "Innovative Full Stack Engineer with a strong background in development, currently enhancing cloud-based solutions with a proven track record of developing and implementing robust, user-friendly applications and automating key processes in both Python and Node.",
		Email:       "potnuripraveen284@gmail.com",
		Phone:       "+91 9959301876",
		Location:    "Hyderabad, India",
		LinkedIn:    "linkedin.com/in/praveenpotnuri/",
		Website:     "praveenpotnuri.netlify.app",
		Experiences: []Experience{
			{
				Title:       "Full Stack Engineer",
				Company:     "Fournine Cloud Solutions",
				Description: "Developed DbSAM, a gateway for developers to access internal databases in the cloud. Created client desktop application using ElectronJS and user management site using MERN stack. Worked on cloud functions, data migration, and automation for various clients.",
				Period:      "May 2022 - Present",
			},
			{
				Title:       "Python Developer",
				Company:     "Kaar Technologies",
				Description: "Led a project developing a sophisticated chatbot backend using Python, integrating with SAP systems and cloud platforms. Implemented tracking mechanisms for asynchronous data transfer jobs and engineered a Salesforce application using MERN stack.",
				Period:      "Dec 2020 - Apr 2022",
			},
		},
		Education: Education{
			Degree:      "Bachelor of Technology (B. Tech.)",
			Institution: "Sathyabama Institute of Science and Technology, Chennai (India)",
			Period:      "July 2017 - May 2021",
			GPA:         "9.2 CGPA",
		},
		Skills: []Skill{
			{
				Category: "Programming Languages",
				Items:    []string{"NodeJS", "Python", "JavaScript"},
			},
			{
				Category: "Frameworks",
				Items:    []string{"Django", "Flask", "React"},
			},
			{
				Category: "Cloud & DevOps",
				Items:    []string{"GCP", "Docker", "Kubernetes", "CI/CD", "GitHub Actions"},
			},
			{
				Category: "Other",
				Items:    []string{"MongoDB", "Linux", "Bash Scripting", "Data Structures", "Problem Solving"},
			},
		},
		Certifications: []string{"Professional Cloud Architect Certification (Expires on May 02, 2026)"},
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}