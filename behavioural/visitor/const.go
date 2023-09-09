package visitor

import "time"

type Severity string

const (
	Mild     Severity = "MILD"
	Moderate Severity = "MODERATE"
	Severe   Severity = "SEVERE"
)

type Sex string

const (
	Female Sex = "FEMALE"
	Male   Sex = "MALE"
)

type MedicalTreatment string

const (
	ToothScaling       MedicalTreatment = "TOOTH_SCALING"
	ToothFiling        MedicalTreatment = "TOOTH_FILING"
	ToothSurgery       MedicalTreatment = "TOOTH_SURGERY"
	CheckBloodPressure MedicalTreatment = "CHECK_BLOOD_PRESSURE"
)

type TreatmentHistory struct {
	visitTime        time.Time
	medicalTreatment MedicalTreatment
	doneBy           string
}

type MedicalRecord struct {
	isHighBloodSugar    bool
	isHighBloodPressure bool
	treatmentHistory    []TreatmentHistory
}
