package battery

type ESSMeasures struct {
	Pess      float32
	Pmaxch    float32
	Pmaxdisch float32
	Eess      float32
	Soc       float32
}

// Fake method made for example
func (measures *ESSMeasures) getEssMeasure() (*ESSMeasures, error) {

	// Send modbus request to get the measures from the equipment
	// If there is an communication error, return an error
	return &ESSMeasures{
		Pess:      120,
		Pmaxch:    -400,
		Pmaxdisch: 400,
		Eess:      50,
		Soc:       60,
	}, nil
}
