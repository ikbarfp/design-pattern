package visitor

// Patient abstraction of patient behaviour
type Patient interface {
	Accept(d Dentist)
	GetSeverity() Severity
	GetAge() int
	GetMedicalRecord() *MedicalRecord
}

type PlaquePatient struct {
	sexPreference Sex
	severity      Severity
}

// Accept method for plaque patient
func (pp *PlaquePatient) Accept(d Dentist) {
	// Visitor can have specific condition
	// 	i.e : this plaque patient only want
	// 	female dentist to handle their cases
	if d.GetSex() == pp.sexPreference {
		d.DoScaling(pp)
	}
}

// GetSeverity method for plaque patient
func (pp *PlaquePatient) GetSeverity() Severity {
	pp.setSeverity()
	return pp.severity
}

// setSeverity private function to
// set severity for plaque patient
func (pp *PlaquePatient) setSeverity() {
	pp.severity = Mild
}

// setSexPreference private function to
// set sex preference for plaque patient
// because only this patient has preference sex
func (pp *PlaquePatient) setSexPreference() {
	pp.sexPreference = Female
}

func (pp *PlaquePatient) GetAge() int {
	return 18
}

func (pp *PlaquePatient) GetMedicalRecord() *MedicalRecord {
	record := &MedicalRecord{}

	treatmentHistory := make([]TreatmentHistory, 0)

	record.treatmentHistory = treatmentHistory

	return record
}

type CavitiesPatient struct {
	severity Severity
}

// Accept method for cavities patient
func (cp *CavitiesPatient) Accept(d Dentist) {
	d.DoToothFilling(cp)
}

// GetSeverity method for cavities patient
func (cp *CavitiesPatient) GetSeverity() Severity {
	cp.setSeverity()
	return cp.severity
}

// setSeverity private function to
// set severity for cavities patient
func (cp *CavitiesPatient) setSeverity() {
	cp.severity = Moderate
}

func (cp *CavitiesPatient) GetAge() int {
	return 45
}

func (cp *CavitiesPatient) GetMedicalRecord() *MedicalRecord {
	record := &MedicalRecord{}

	treatmentHistory := make([]TreatmentHistory, 0)

	record.treatmentHistory = treatmentHistory

	return record
}

type BrokenToothPatient struct {
	severity Severity
}

// Accept method for broken tooth patient
func (btp *BrokenToothPatient) Accept(d Dentist) {
	d.DoToothFilling(btp)
}

// GetSeverity method for broken tooth patient
func (btp *BrokenToothPatient) GetSeverity() Severity {
	btp.setSeverity()
	return btp.severity
}

// setSeverity private function to
// set severity for broken tooth patient
func (btp *BrokenToothPatient) setSeverity() {
	btp.severity = Severe
}

func (btp *BrokenToothPatient) GetAge() int {
	return 60
}

func (btp *BrokenToothPatient) GetMedicalRecord() *MedicalRecord {
	record := &MedicalRecord{}

	treatmentHistory := make([]TreatmentHistory, 0)

	record.treatmentHistory = treatmentHistory

	return record
}
