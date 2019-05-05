package simulationSystem

var simulationSystem SimulationSystemControl = nil

func Get() SimulationSystem {
	return simulationSystem
}

func GetControl() SimulationSystemControl {
	return simulationSystem
}

func Provide(service SimulationSystemControl) {
	simulationSystem = service
}
