package main

import (
	"html/template"
	"log"
	"net/http"
)

// Experience represents civilian job experience.
type Experience struct {
	Role        string
	Company     string
	Duration    string
	Description string
}

// MilitaryExperience represents a military role and associated honors.
type MilitaryExperience struct {
	Role        string
	Unit        string
	Duration    string
	Description string
	Awards      []string
}

// PageData holds all data passed to the template.
type PageData struct {
	Name               string
	Title              string
	Summary            string
	Experiences        []Experience
	MilitaryExperience []MilitaryExperience
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s\n", r.URL.Path)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		log.Printf("Template error: %v\n", err)
		return
	}

	data := PageData{
		Name:  "Mark De La Garza",
		Title: "Senior Systems Engineer",
		Summary: `Senior MES and Automation Engineer with over 15 years of experience leading complex automation, 
MES integration, and industrial control initiatives across aerospace, defense, and smart manufacturing sectors. 
Extensive background in robotic systems, SQL-backed MES architecture, PLC troubleshooting, and secure military-grade payload systems.`,
		Experiences: []Experience{
			{
				Role:        "Senior MES Engineer",
				Company:     "Navistar",
				Duration:    "2023 – Present",
				Description: `Led MES integrations across 30+ production stations. Implemented custom logic in Rockwell FTPC and Kepware 
for real-time data capture and production optimization. Supported EV transition projects and root cause investigations.`,
			},
			{
				Role:        "Systems Engineer",
				Company:     "Symbotic",
				Duration:    "2021 – 2023",
				Description: `Integrated ABB robotics, Studio 5000-based PLCs, and conveyor control logic in autonomous distribution facilities. 
Managed AGV diagnostics, robotic arm tuning, and root cause analyses to increase system uptime.`,
			},
			{
				Role:        "Payload Integration Lead",
				Company:     "General Atomics",
				Duration:    "2017 – 2022",
				Description: `Directed integration, test, and validation of advanced ISR payloads on MQ-9 and Predator-C UAVs. 
Coordinated OT&E, cryptographic compliance, and MIL-STD interface design across USAF and allied missions.`,
			},
		},
		MilitaryExperience: []MilitaryExperience{
			{
				Role:     "Weapons Standardization & Safety Manager",
				Unit:     "USAF – Indian Springs, NV & Houston, TX",
				Duration: "2016 – 2019",
				Description: `Managed safety protocols for 177 personnel and secured $180M in mission-critical armament systems. 
Directed 425+ live munition operations with zero incidents. Oversaw certifications and readiness across deployment cycles.`,
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
				Unit:     "USAF – 432 Maintenance Group, Creech AFB",
				Duration: "2010 – 2016",
				Description: `Performed weapons systems diagnostics and armament support on MQ-1 and MQ-9 Reaper platforms. 
Led T.O. compliance tracking, crew certification, and NATO deployment prep. Conducted site-wide load standardization training.`,
				Awards: []string{
					"AF Outstanding Unit Award",
					"Global War on Terrorism Expeditionary Medal",
					"AF Good Conduct Medal",
					"AF Training Ribbon",
				},
			},
		},
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Template execution error: %v\n", err)
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files (CSS, images, PDFs)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle root
	http.HandleFunc("/", handler)

	log.Println("Server running on http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
