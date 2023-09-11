package visitor

func Execute() {

	// Initialize all type of dentist
	juniorDentist := &JuniorDentist{}
	seniorDentist := &JuniorDentist{}
	specialistDentist := &SpecialistDentist{}

	// Initialize new patient 1
	patient1 := &PlaquePatient{}

	patient1.Accept(juniorDentist)
	juniorDentist.DoScaling(patient1)
	juniorDentist.DoToothFilling(patient1)
	juniorDentist.DoSurgery(patient1)

	patient1.Accept(seniorDentist)
	seniorDentist.DoScaling(patient1)
	seniorDentist.DoToothFilling(patient1)
	seniorDentist.DoSurgery(patient1)

	patient1.Accept(specialistDentist)
	specialistDentist.DoScaling(patient1)
	specialistDentist.DoToothFilling(patient1)
	specialistDentist.DoSurgery(patient1)

	// Initialize new patient 2
	patient2 := &CavitiesPatient{}

	patient2.Accept(juniorDentist)
	juniorDentist.DoScaling(patient2)
	juniorDentist.DoToothFilling(patient2)
	juniorDentist.DoSurgery(patient2)

	patient2.Accept(seniorDentist)
	seniorDentist.DoScaling(patient2)
	seniorDentist.DoToothFilling(patient2)
	seniorDentist.DoSurgery(patient2)

	patient2.Accept(specialistDentist)
	specialistDentist.DoScaling(patient2)
	specialistDentist.DoToothFilling(patient2)
	specialistDentist.DoSurgery(patient2)

	// Initialize new patient 3
	patient3 := &BrokenToothPatient{}

	patient3.Accept(juniorDentist)
	juniorDentist.DoScaling(patient3)
	juniorDentist.DoToothFilling(patient3)
	juniorDentist.DoSurgery(patient3)

	patient3.Accept(seniorDentist)
	seniorDentist.DoScaling(patient3)
	seniorDentist.DoToothFilling(patient3)
	seniorDentist.DoSurgery(patient3)

	patient3.Accept(specialistDentist)
	specialistDentist.DoScaling(patient3)
	specialistDentist.DoToothFilling(patient3)
	specialistDentist.DoSurgery(patient3)

}
