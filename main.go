/*											OBJECTIF
The EMS objective is to ensure that the industrial site power consumption remains under a maximal
value PmaxSite. This means, the EMS must keep PmaxSite < Ppoc at all times. Exporting
power to the Grid is tolerated
*/

/*
• the industrial facility load
• an Energy Storage System (ESS, e.g. a battery) of capacity ess_capacity in kWh
• a photovoltaic (PV) power plant of peak power pv_peak in kW
*/
//ess=kwh
//

//Pmaxdisch>0
//Pload !<Pmaxch
//Pload !>Pmaxdisch

//Ppoc = Pess + Ppv + Pload.
// Pload = -Pess - Ppv + Ppoc
/*
Read-only values
• Pess: current ESS active power output in kW (< 0 means charge / > 0 means discharge)
• Pmaxch: current ESS maximal charge power capability in kW (necessarily <= 0 by
convention)
• Pmaxdisch: current ESS maximal discharge power capability in kW (necessarily >= 0 by
convention)
• Eess: current ESS stored energy in kWh (necessarily >= 0)
Write values
• setpointPess: active power setpoint computed by the EMS in kW (< 0 for charge
setpoint, > 0 for discharge setpoint)
*/

package main

import (
	"fmt"

	Battery "github.com/Baptiste-Ferrand/testTechnique/pkg/battery"
	Industrial "github.com/Baptiste-Ferrand/testTechnique/pkg/industrial"
	Pv "github.com/Baptiste-Ferrand/testTechnique/pkg/pv"
)

// main function
func main() {

	// Equipment Initialization
	var battery Battery.Equipment
	var industrial Industrial.Equipment
	var pv Pv.Equipment

	for {
		battery.Cycle()
		pv.Cycle()
		industrial.Cycle(battery.Measures.Pess, pv.Measures.Ppv)
		fmt.Printf("%v %v", battery.Measures.Pess, pv.Measures.Ppv)

		// Charging Ess
		if -industrial.Measures.Pload < pv.Measures.Ppv {
			battery.Setpoint((-industrial.Measures.Pload - pv.Measures.Ppv))
		}

		// Discharging Ess
		if -industrial.Measures.Pload > pv.Measures.Ppv {
			battery.Setpoint((-industrial.Measures.Pload - pv.Measures.Ppv))
		}

		// Save Discharging Ess
		if battery.Measures.Pess > battery.Measures.Pmaxdisch*0.80 {
			battery.Setpoint((battery.Measures.Pmaxdisch * 0.80))
			// Take -(battery.Measures.Pess - battery.Measures.Pmaxdisch*0.80) in Poc
		}

		// Save Chargin Ess
		if battery.Measures.Pess < battery.Measures.Pmaxch*0.80 {
			battery.Setpoint((battery.Measures.Pmaxdisch * 0.80))
		}

		if -industrial.Measures.Pload > pv.Measures.Ppv+battery.Measures.Pess {
			// Take energy in Poc
		}

		if industrial.Measures.PmaxSite > industrial.Measures.Ppoc*1.2 {
			// Take less in Poc
		}

	}
}
