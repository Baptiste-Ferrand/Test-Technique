package industrial

import "fmt"

type Equipment struct {
	Measures *IndustrialMeasures
}

type IndustrialMeasures struct {
	Pload    float32
	Ppoc     float32
	PmaxSite float32
}

func (industrial *Equipment) Cycle(Pess float32, Ppv float32) {
	var err error

	industrial.Measures, err = industrial.GetIndustrialMeasure(Pess, Ppv)
	if err != nil {
		fmt.Printf("Error while getting measures from the industrial: %v", err)
	}
}

func (industrial *Equipment) GetIndustrialMeasure(Pess float32, Ppv float32) (*IndustrialMeasures, error) {

	var (
		err  error
		data *IndustrialMeasures
	)

	if data, err = industrial.Measures.getIndustrialMeasure(Pess, Ppv); err != nil {
		fmt.Printf("Error while getting measures from the industrial: %v", err)
		return nil, err
	}

	industrial.Measures = data
	return industrial.Measures, nil
}

// Fake method made for example
func (pv *IndustrialMeasures) getIndustrialMeasure(Pess float32, Ppv float32) (*IndustrialMeasures, error) {
	// Send modbus request to get the measures from the industrial
	var ppoc float32 = -300
	var pload = pv.GetPloadCalcul(ppoc, Pess, Ppv)
	// If there is an communication error, return an error
	return &IndustrialMeasures{
		Pload:    pload,
		Ppoc:     ppoc,
		PmaxSite: -400,
	}, nil
}

func (pv *IndustrialMeasures) GetPloadCalcul(Ppoc float32, Pess float32, Ppv float32) float32 {
	var pload = Ppoc - Pess - Ppv
	return pload
}
