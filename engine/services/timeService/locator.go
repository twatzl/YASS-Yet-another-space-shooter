package timeService

var timeService TimeServiceControl = nil

func Get() TimeService {
	return timeService
}

func GetControl() TimeServiceControl {
	return timeService
}

func Provide(service TimeServiceControl) {
	timeService = service
}
