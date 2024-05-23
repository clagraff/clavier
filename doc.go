// Package clavier provides a simple way to manage and track the state of input controls
// such as keyboard keys, mouse buttons, and key combinations in an Ebiten-based game.
//
// The package offers a generic Control interface, which can represent different types of input
// controls. It includes utility functions to check the state of these controls, like whether
// they are activated, just activated, still activated, just deactivated, or still deactivated.
//
// Example usage:
//
//	package main
//
//	import (
//	    "github.com/yourusername/clavier"
//	    "github.com/hajimehoshi/ebiten/v2"
//	    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
//	    "log"
//	)
//
//	type Game struct{}
//
//	func (g *Game) Update() error {
//	    clavier.Update()
//	    spaceKey := clavier.Key(ebiten.KeySpace)
//	    if clavier.JustActivated(spaceKey) {
//	        println("Space key just activated!")
//	    }
//	    return nil
//	}
//
//	func (g *Game) Draw(screen *ebiten.Image) {
//	    ebitenutil.DebugPrint(screen, "Press Space")
//	}
//
//	func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	    return 320, 240
//	}
//
//	func main() {
//	    game := &Game{}
//	    ebiten.SetWindowSize(640, 480)
//	    ebiten.SetWindowTitle("Clavier Example")
//	    if err := ebiten.RunGame(game); err != nil {
//	        log.Fatal(err)
//	    }
//	}
package clavier
