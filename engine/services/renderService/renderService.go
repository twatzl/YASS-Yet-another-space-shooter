package renderService

type RenderService interface {
	GetContext() RenderContext
	SetContext(ctx RenderContext)
}