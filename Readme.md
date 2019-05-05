# YASS - Yet Another Space Shooter

YASS is a simple local multiplayer space shooter game

## Structure

### assets
Contains sprites, sounds, maps and other stuff needed for running the game.

### engine

Contains the part of the code which is potentially reusable across other games and deals with creating a general game structure.

#### engine/game

Contains the core code for the game. Code which defines game objects, scenes, etc.

#### engine/components

Contains predefined components such as the transform component.

#### engine/services

Contains services. Services are global objects which provide some kind of service to game objects.

#### engine/systems

Systems are similar to services, but systems require the game objects to register and will actively notify game objects.

### scenes

Contains the files defining the scenes of the game as well as the code for the game objects.

### terrain

Contains the terrain implementation

## Compiling

On you own platform run

```go run```

Cross compoiling for Windows:

```GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -o game.exe main.go```