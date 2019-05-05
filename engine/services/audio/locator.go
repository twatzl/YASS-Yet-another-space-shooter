package audio

var audioService AudioService = nil

func ProvideAudioService(service AudioService) {
	audioService = service
}

func GetAudioService() AudioService {
	return audioService
}