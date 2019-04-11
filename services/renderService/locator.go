package renderService

var audioService RenderService = nil

func ProvideRenderService(service RenderService) {
	audioService = service
}

func Get() RenderService {
	return audioService
}