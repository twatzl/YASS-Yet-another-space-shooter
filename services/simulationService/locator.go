package simulationService

var simulationService SimulationServiceControl = nil

func Get() SimulationService {
	return simulationService
}

func GetControl() SimulationServiceControl {
	return simulationService
}

func Provide(service SimulationServiceControl) {
	simulationService = service
}
