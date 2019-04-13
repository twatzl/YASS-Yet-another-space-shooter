package inputSystem

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/twatzl/pixel-test/services/windowService"
)

type KeyEvent int

const (
	KeyPressed KeyEvent = iota
	KeyJustPressed
	KeyJustReleased
	KeyRepeated
)

type InputPredicate func(button pixelgl.Button) bool

type InputHandler func()

type inputHandlerStruct struct {
	event   KeyEvent
	button  pixelgl.Button
	handler InputHandler
}

type InputSystem interface {
	RegisterKeyEventHandler(event KeyEvent, button pixelgl.Button, handler InputHandler) int
	RemoveKeyEventHandler(id int)
}

type InputSystemControl interface {
	InputSystem
	HandleInputs()
}

type inputSystemImpl struct {
	idCounter int
	handlers  map[int]inputHandlerStruct
}

func New() InputSystemControl {
	return &inputSystemImpl{
		idCounter: 0,
		handlers:  map[int]inputHandlerStruct{},
	}
}

func (is *inputSystemImpl) RegisterKeyEventHandler(event KeyEvent, button pixelgl.Button, handler InputHandler) int {
	ihs := inputHandlerStruct{
		event:   event,
		button:  button,
		handler: handler,
	}
	is.idCounter = is.idCounter + 1

	is.handlers[is.idCounter] = ihs
	return is.idCounter
}

func (is *inputSystemImpl) RemoveKeyEventHandler(id int) {
	delete(is.handlers, id)
}

func (is *inputSystemImpl) HandleInputs() {
	window := windowService.Get().GetWindow()

	/* a bit clumsy handling but necessary if we want to have a nice interface */
	for _, handler := range is.handlers {
		eventTriggered := false
		switch handler.event {
		case KeyPressed:
			eventTriggered = window.Pressed(handler.button)
			break
		case KeyJustPressed:
			eventTriggered = window.JustPressed(handler.button)
			break
		case KeyJustReleased:
			eventTriggered = window.JustReleased(handler.button)
			break
		case KeyRepeated:
			eventTriggered = window.Repeated(handler.button)
			break
		}

		if eventTriggered {
			handler.handler()
		}
	}
}
