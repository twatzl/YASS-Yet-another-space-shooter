package renderService

var renderService RenderService = nil

func ProvideRenderService(service RenderService) {
	renderService = service
}

func Get() RenderService {
	return renderService
}