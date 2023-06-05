package battery

import (
	"fmt"
)

type Equipment struct {
	Measures  *ESSMeasures
	setpointP float32
}

// Function made for refresh the measures of the equipment and send the setpoint
func (ess *Equipment) Cycle() {
	var err error

	ess.Measures, err = ess.GetEssMeasure()
	if err != nil {
		fmt.Printf("Error while getting measures from the battery: %v", err)
	}

	// Send the setpoint to the equipment
	if err := ess.sendSetpoint(); err != nil {
		fmt.Printf("Error while sending setpoint to the battery: %v", err)
	}
}

// function for set the EssSetpoint
func (ess *Equipment) Setpoint(Pess float32) {
	ess.setpointP = Pess
	fmt.Printf("\n Pess =%v \n", Pess)
}

func (ess *Equipment) GetEssMeasure() (*ESSMeasures, error) {
	var (
		err  error
		data *ESSMeasures
	)

	if data, err = ess.Measures.getEssMeasure(); err != nil {
		fmt.Printf("Error while getting measures from the battery: %v", err)
		return nil, err
	}

	ess.Measures = data
	return ess.Measures, nil
}

// Fake method made for example
func (ess *Equipment) sendSetpoint() error {
	// Send modbus request to send the setpoint to the equipment

	// If there is an communication error, return an error
	return nil
}
