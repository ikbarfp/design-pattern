package visitor

import (
	"fmt"
	"time"
)

// Dentist abstraction of dentist behaviour
// and it acts as a visitor
type Dentist interface {
	GetSex() Sex
	DoScaling(p Patient)
	DoToothFilling(p Patient)
	DoSurgery(p Patient)
}

type JuniorDentist struct {
	sex  Sex
	cost float64
}

// GetSex get junior dentist sex
func (jd *JuniorDentist) GetSex() Sex {
	jd.setSex()
	return jd.sex
}

// setSex private function to
// set sex of junior dentist as default
func (jd *JuniorDentist) setSex() {
	jd.sex = Female
}

// DoScaling do scaling with subjective
// condition from junior dentist
func (jd *JuniorDentist) DoScaling(p Patient) {
	records := p.GetMedicalRecord()

	// In this example, junior dentist
	// want to give candy to young patient
	// i.e : only patient with age below 12 will get candy
	if p.GetAge() < 12 {
		fmt.Println("Here is candy, very brave young lads :)")
	}

	// In this example, junior dentist
	// had their own opinion about scaling
	// i.e : only patient with last time scaling > 30 days
	// can do a tooth scaling
	for _, record := range records.treatmentHistory {
		switch record.medicalTreatment {
		case ToothScaling:
			if time.Now().Sub(record.visitTime) > ((time.Hour * 24) * 30) {
				fmt.Println("Let's do a tooth scaling with me :D")
			} else {
				fmt.Println("Come back later after 30 days from your last scaling :)")
			}
		default:
			continue
		}
	}
	return
}

// DoToothFilling do tooth filling with subjective
// condition from junior dentist
func (jd *JuniorDentist) DoToothFilling(p Patient) {
	records := p.GetMedicalRecord()

	// In this example, junior dentist
	// want to give candy to young patient
	// i.e : only patient with age below 12 will get candy
	if p.GetAge() < 12 {
		fmt.Println("Here is candy, very brave young lads :)")
	}

	// In this example, junior dentist
	// had their own opinion about tooth filing
	// i.e : only patient with last time tooth surgery > 6 month
	// can do a tooth filing
	for _, record := range records.treatmentHistory {
		switch record.medicalTreatment {
		case ToothSurgery:
			if time.Now().Sub(record.visitTime) > (((time.Hour * 24) * 30) * 6) {
				fmt.Println("Let's do a tooth filing with me :D")
			} else {
				fmt.Println("Come back later after 6 months from your last tooth surgery :)")
				return
			}
		default:
			continue
		}
	}
	return
}

// DoSurgery do surgery with subjective
// condition from junior dentist
func (jd *JuniorDentist) DoSurgery(p Patient) {
	// In this example,
	// junior dentist can't do a surgery
	// if the severity above mild
	if p.GetSeverity() != Mild {
		fmt.Println("I'm so sorry, i can't do a major surgery :(")
	} else {
		fmt.Println("Aye! Come to surgery room, it won't be hurt :D")
	}
	return
}

type SeniorDentist struct {
	sex  Sex
	cost float64
}

// GetSex get senior dentist sex
func (sd *SeniorDentist) GetSex() Sex {
	sd.setSex()
	return sd.sex
}

// setSex private function to
// set sex of senior dentist as default
func (sd *SeniorDentist) setSex() {
	sd.sex = Male
}

// DoScaling do scaling with subjective
// condition from senior dentist
func (sd *SeniorDentist) DoScaling(p Patient) {
	records := p.GetMedicalRecord()

	// In this example, senior dentist
	// had their own opinion about scaling
	// i.e : only patient with last time scaling > 15 days
	// or last time tooth filing > 3 months
	// can do a scaling
	for _, record := range records.treatmentHistory {
		switch record.medicalTreatment {
		case ToothScaling:
			if time.Now().Sub(record.visitTime) > ((time.Hour * 24) * 15) {
				fmt.Println("Let's do a tooth scaling")
			} else {
				fmt.Println("Come back later after 15 days from your last scaling")
				return
			}
		case ToothFiling:
			if time.Now().Sub(record.visitTime) > (((time.Hour * 24) * 30) * 3) {
				fmt.Println("Let's do a tooth scaling")
			} else {
				fmt.Println("Come back later after 3 month from your last tooth filing")
				return
			}
		default:
			continue
		}
	}
	return
}

// DoToothFilling do tooth filling with subjective
// condition from senior dentist
func (sd *SeniorDentist) DoToothFilling(p Patient) {
	// In this example, senior dentist
	// accept only patient with severity under severe
	// to do a tooth filing with him
	switch p.GetSeverity() {
	case Severe:
		fmt.Println("Sorry, i can't do a tooth filing on you")
		return
	}

	// In this example,
	// senior dentist can do a tooth filing,
	// but it takes slower times than specialist can do
	fmt.Println("Let's do a tooth filing")
	time.Sleep(4 * time.Second)
}

// DoSurgery do surgery with subjective
// condition from senior dentist
func (sd *SeniorDentist) DoSurgery(p Patient) {
	// In this example, senior dentist
	// only accept patient with severity under severe
	// to do a surgery with him
	switch p.GetSeverity() {
	case Severe:
		fmt.Println("Sorry, i am hesitant to do a surgery on you")
		return
	}

	// In this example, senior dentist
	// only accept patient between 10 until 60 years old
	// to do a surgery with him
	age := p.GetAge()

	if age < 10 && age > 60 {
		fmt.Println("Sorry, i am hesitant to do a surgery on you")
		return
	}

	records := p.GetMedicalRecord()

	// In this example, senior dentist
	// had their own opinion about tooth surgery
	// i.e : only patient with last time tooth surgery > 6 month
	// can do a tooth surgery
	for _, record := range records.treatmentHistory {
		switch record.medicalTreatment {
		case ToothSurgery:
			if time.Now().Sub(record.visitTime) > (((time.Hour * 24) * 30) * 6) {

				// In this example,
				// senior dentist can do a surgery,
				// but it takes slower times than specialist can do
				fmt.Println("Let's do a tooth surgery")
				time.Sleep(4 * time.Second)
			} else {
				fmt.Println("Come back later after 6 months from your last tooth surgery")
			}
		default:
			continue
		}
	}
	return
}

type SpecialistDentist struct {
	sex  Sex
	cost float64
}

// GetSex get specialist dentist sex
func (spd *SpecialistDentist) GetSex() Sex {
	spd.setSex()
	return spd.sex
}

// setSex private function to
// set sex of specialist dentist as default
func (spd *SpecialistDentist) setSex() {
	spd.sex = Male
}

// DoScaling do scaling with subjective
// condition from specialist dentist
func (spd *SpecialistDentist) DoScaling(p Patient) {
	// In this example,
	// specialist dentist don't do scaling
	// because it's too easy for him
	fmt.Println("Go to my junior, i will not do that anymore!")
	return
}

// DoToothFilling do tooth filling with subjective
// condition from specialist dentist
func (spd *SpecialistDentist) DoToothFilling(p Patient) {
	// In this example, specialist dentist
	// only accept patient with severity above mild
	// to do a tooth filing with him
	switch p.GetSeverity() {
	case Mild:
		fmt.Println("Go to my junior, i will not do that anymore!")
		return
	}

	// In this example, specialist dentist
	// had their own opinion about tooth filing
	// i.e : only patient with below 60 years old
	// can do a tooth filing with him
	if p.GetAge() > 60 {
		fmt.Println("It's better to have dentures installed rather than fill cavities!")
		return
	}
}

// DoSurgery do surgery with subjective
// condition from specialist dentist
func (spd *SpecialistDentist) DoSurgery(p Patient) {
	records := p.GetMedicalRecord()

	// In this example, specialist dentist
	// had their own opinion about tooth surgery
	// i.e : only patient with below 70 years old
	// can do a tooth surgery with him
	if p.GetAge() > 70 {
		fmt.Println("Let me think to find another solution, i can't do this on you Sir!")
		return
	}

	// In this example, specialist dentist
	// had their own opinion about tooth surgery
	// i.e : only patient with controlled blood sugar
	// can do a tooth surgery with him
	if records.isHighBloodSugar {
		fmt.Println("Please control your blood sugar first!")
		return
	}

	// In this example, specialist dentist
	// had their own opinion about tooth surgery
	// i.e : only patient with controlled blood pressure
	// can do a tooth surgery with him
	if records.isHighBloodPressure {
		fmt.Println("Please control your blood sugar first!")
		return
	}

	// In this example,
	// senior dentist can do a surgery,
	// and it takes faster times than senior dentist can do
	fmt.Println("Let me handle this!")
	time.Sleep(2 * time.Second)
	return
}
