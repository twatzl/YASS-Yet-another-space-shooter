package windowService

var windowService WindowService = &windowServiceImpl{}

func Get() WindowService {
	return windowService
}