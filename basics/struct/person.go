package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street     string
	City       string
	State      string
	Country    string
	PostalCode string
}

type PhoneNumber struct {
	Type   string // e.g., "mobile", "home", "work"
	Number string
}

type Email struct {
	Type    string // e.g., "personal", "work"
	Address string
}

type SocialMedia struct {
	Platform string // e.g., "Twitter", "LinkedIn", "Facebook"
	Handle   string
	URL      string
}

type Employment struct {
	CompanyName      string
	JobTitle         string
	StartDate        time.Time
	EndDate          time.Time
	Responsibilities []string
}

type Education struct {
	Institution  string
	Degree       string
	FieldOfStudy string
	StartDate    time.Time
	EndDate      time.Time
	GPA          float64
}

type Language struct {
	Name        string
	Proficiency string // e.g., "Native", "Fluent", "Intermediate"
}

type PhysicalAttributes struct {
	Height       float64 // in centimeters
	Weight       float64 // in kilograms
	EyeColor     string
	HairColor    string
	BloodType    string
	Disabilities []string
}

type EmergencyContact struct {
	Name         string
	Relationship string // e.g., "Spouse", "Parent"
	Phone        PhoneNumber
	Email        string
	Address      Address
}

type PersonalIdentification struct {
	NationalID       string
	PassportNumber   string
	DriverLicense    string
	SocialSecurityNo string
	TaxID            string
}

type Person struct {
	FirstName          string
	LastName           string
	MiddleName         string
	Nickname           string
	DateOfBirth        time.Time
	PlaceOfBirth       string
	Gender             string
	MaritalStatus      string
	Nationality        string
	Ethnicity          string
	Religion           string
	SexualOrientation  string
	ResidentialAddress Address
	MailingAddress     Address
	PhoneNumbers       []PhoneNumber
	Emails             []Email
	SocialMedias       []SocialMedia
	EmergencyContacts  []EmergencyContact
	PersonalID         PersonalIdentification
	PhysicalDetails    PhysicalAttributes
	LanguagesSpoken    []Language
	EducationHistory   []Education
	EmploymentHistory  []Employment
	Hobbies            []string
	FavoriteBooks      []string
	FavoriteMovies     []string
	FavoriteMusic      []string
}

func main() {
	// Example instantiation of the Person struct
	person := Person{
		FirstName:         "John",
		LastName:          "Doe",
		Nickname:          "Johnny",
		DateOfBirth:       time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC),
		PlaceOfBirth:      "New York, USA",
		Gender:            "Male",
		MaritalStatus:     "Single",
		Nationality:       "American",
		Ethnicity:         "Caucasian",
		Religion:          "Christian",
		SexualOrientation: "Heterosexual",
		ResidentialAddress: Address{
			Street:     "123 Main St",
			City:       "New York",
			State:      "NY",
			Country:    "USA",
			PostalCode: "10001",
		},
		PhoneNumbers: []PhoneNumber{
			{Type: "mobile", Number: "123-456-7890"},
			{Type: "home", Number: "098-765-4321"},
		},
		Emails: []Email{
			{Type: "personal", Address: "john.doe@example.com"},
			{Type: "work", Address: "j.doe@company.com"},
		},
		SocialMedias: []SocialMedia{
			{Platform: "Twitter", Handle: "@johndoe", URL: "https://twitter.com/johndoe"},
		},
		EmergencyContacts: []EmergencyContact{
			{
				Name:         "Jane Doe",
				Relationship: "Sister",
				Phone:        PhoneNumber{Type: "mobile", Number: "987-654-3210"},
				Email:        "jane.doe@example.com",
				Address: Address{
					Street:     "456 Elm St",
					City:       "Los Angeles",
					State:      "CA",
					Country:    "USA",
					PostalCode: "90001",
				},
			},
		},
		PersonalID: PersonalIdentification{
			NationalID:       "123456789",
			PassportNumber:   "X1234567",
			DriverLicense:    "D9876543",
			SocialSecurityNo: "123-45-6789",
			TaxID:            "987-65-4321",
		},
		PhysicalDetails: PhysicalAttributes{
			Height:       180.34,
			Weight:       75.5,
			EyeColor:     "Brown",
			HairColor:    "Black",
			BloodType:    "O+",
			Disabilities: []string{"None"},
		},
		LanguagesSpoken: []Language{
			{Name: "English", Proficiency: "Native"},
			{Name: "Spanish", Proficiency: "Fluent"},
		},
		EducationHistory: []Education{
			{
				Institution:  "University of Example",
				Degree:       "Bachelor of Science",
				FieldOfStudy: "Computer Science",
				StartDate:    time.Date(2008, time.September, 1, 0, 0, 0, 0, time.UTC),
				EndDate:      time.Date(2012, time.May, 15, 0, 0, 0, 0, time.UTC),
				GPA:          3.8,
			},
		},
		EmploymentHistory: []Employment{
			{
				CompanyName:      "Example Corp",
				JobTitle:         "Software Engineer",
				StartDate:        time.Date(2013, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndDate:          time.Now(),
				Responsibilities: []string{"Developing software", "Leading projects"},
			},
		},
		Hobbies:        []string{"Reading", "Hiking", "Photography"},
		FavoriteBooks:  []string{"1984 by George Orwell", "To Kill a Mockingbird by Harper Lee"},
		FavoriteMovies: []string{"The Shawshank Redemption", "Inception"},
		FavoriteMusic:  []string{"Rock", "Classical"},
	}

	fmt.Print(person)
}
