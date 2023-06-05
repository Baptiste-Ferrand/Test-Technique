package pv

import "fmt"

type Equipment struct {
	Measures *PvMeasures
}

type PvMeasures struct {
	Ppv float32
}

func (pv *Equipment) Cycle() {
	var err error
	pv.Measures, err = pv.GetPvMeasure()
	if err != nil {
		fmt.Printf("Error while getting measures from the pv: %v", err)
	}
}

func (pv *Equipment) GetPvMeasure() (*PvMeasures, error) {

	var (
		err  error
		data *PvMeasures
	)

	if data, err = pv.Measures.getPvMeasure(); err != nil {
		fmt.Printf("Error while getting measures from the pv: %v", err)
		return nil, err
	}

	pv.Measures = data
	return pv.Measures, nil
}

// Fake method made for example
func (pv *PvMeasures) getPvMeasure() (*PvMeasures, error) {
	// Send modbus request to get the measures from the pv

	// If there is an communication error, return an error
	return &PvMeasures{
		Ppv: 200,
	}, nil
}
