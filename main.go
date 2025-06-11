package main

import (
	"html/template"
	"net/http"
	"log"
)

type Experience struct {
	Role        string
	Company     string
	Duration    string
	Description string
}

type MilitaryExperience struct {
	Role        string
	Unit        string
	Duration    string
	Description string
	Awards      []string
}

type PageData struct {
	Name               string
	Title              string
	Summary            string
	Experiences        []Experience
	MilitaryExperience []MilitaryExperience
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received at:", r.URL.Path)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := PageData{
		Name:  "Mark De La Garza",
		Title: "Senior System Engineer",
		Summary: "Senior MES and Automation Engineer with over 15 years of hands-on experience across aerospace, defense, and smart manufacturing. Proven leader in integrating robotics, MES, and control systems for mission-critical operations. U.S. Air Force Veteran with over a decade supporting DoD and State Department operations overseas, decorated with multiple commendations for excellence in safety, systems readiness, and technical execution.",
		Experiences: []Experience{
			{
				Role:        "Senior MES Engineer",
				Company:     "Navistar",
				Duration:    "2023 – Present",
				Description: "Integrated MES systems, led root cause analysis, and optimized production flow.",
			},
			{
				Role:        "Systems Engineer",
				Company:     "Symbotic",
				Duration:    "2021 – 2023",
				Description: "Implemented robotic automation in large-scale distribution centers.",
			},
			{
				Role:        "Payload Integration Lead",
				Company:     "General Atomics",
				Duration:    "2017 – 2022",
				Description: "Led systems integration and validation for advanced UAV platforms including MQ-9 and Predator-C.",
			},
		},
		MilitaryExperience: []MilitaryExperience{
			{
				Role:     "Weapons Standardization & Safety Manager",
				Unit:     "United States Air Force – Indian Springs, NV & Houston, TX",
				Duration: "2016 – 2019",
				Description: "Led safety programs for 177 personnel and safeguarded $180M in aerospace equipment. Directed 425+ live munitions operations during a 40-month overseas deployment with a flawless zero-mishap record. Managed standardization and readiness for global defense missions. Elevated to Lead Safety Coordinator overseeing load crew certifications and operational continuity.",
				Awards: []string{
					"Air Force Commendation Medal",
					"Air Force Achievement Medal",
					"Meritorious Unit Award (3 OLC)",
					"Afghanistan Campaign Medal (2 Service Stars)",
					"Air Force Expeditionary Service Ribbon (Gold Border, 1 OLC)",
				},
			},
			{
				Role:     "Aircraft Armament Systems Specialist",
				Unit:     "United States Air Force – 432 Maintenance Group, Creech AFB",
				Duration: "2010 – 2016",
				Description: "Performed maintenance and armament support for MQ-9 and MQ-1 platforms. Managed evaluations, technical orders, and readiness tracking. Supported mission deployments and NATO-aligned operations in Afghanistan and the Middle East.",
				Awards: []string{
					"AF Outstanding Unit Award",
					"Global War on Terrorism Expeditionary Medal",
					"AF Good Conduct Medal",
					"AF Training Ribbon",
				},
			},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)

	log.Println("Server started at http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
