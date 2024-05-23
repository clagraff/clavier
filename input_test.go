package clavier

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestInitialKeyState(t *testing.T) {
	spaceKey := Key(ebiten.KeySpace)

	// Initial state should be StillDeactivatedState
	if spaceKey.State() != StillDeactivatedState {
		t.Errorf("Expected initial state to be StillDeactivatedState, got %v", spaceKey.State())
	}
}

func TestKeyJustActivatedState(t *testing.T) {
	spaceKey := Key(ebiten.KeySpace)

	// Simulate key press
	isKeyPressedFunc = func(key ebiten.Key) bool {
		return key == ebiten.KeySpace
	}
	Update()

	// After key press, state should be JustActivatedState
	if spaceKey.State() != JustActivatedState {
		t.Errorf("Expected state to be JustActivatedState after key press, got %v", spaceKey.State())
	}
}

func TestKeyStillActivatedState(t *testing.T) {
	spaceKey := Key(ebiten.KeySpace)

	// Simulate key press
	isKeyPressedFunc = func(key ebiten.Key) bool {
		return key == ebiten.KeySpace
	}
	Update()
	Update()

	// After two updates, state should be StillActivatedState
	if spaceKey.State() != StillActivatedState {
		t.Errorf("Expected state to be StillActivatedState after two updates, got %v", spaceKey.State())
	}
}

func TestKeyJustDeactivatedState(t *testing.T) {
	spaceKey := Key(ebiten.KeySpace)

	// Simulate key press and release
	isKeyPressedFunc = func(key ebiten.Key) bool {
		return key == ebiten.KeySpace
	}
	Update()
	isKeyPressedFunc = func(key ebiten.Key) bool {
		return false
	}
	Update()

	// After key release, state should be JustDeactivatedState
	if spaceKey.State() != JustDeactivatedState {
		t.Errorf("Expected state to be JustDeactivatedState after key release, got %v", spaceKey.State())
	}
}

func TestKeyStillDeactivatedState(t *testing.T) {
	spaceKey := Key(ebiten.KeySpace)

	// Simulate key press and release
	isKeyPressedFunc = func(key ebiten.Key) bool {
		return key == ebiten.KeySpace
	}
	Update()
	isKeyPressedFunc = func(key ebiten.Key) bool {
		return false
	}
	Update()
	Update()

	// After another update, state should be StillDeactivatedState
	if spaceKey.State() != StillDeactivatedState {
		t.Errorf("Expected state to be StillDeactivatedState after another update, got %v", spaceKey.State())
	}
}

func TestInitialMouseButtonState(t *testing.T) {
	leftButton := MouseButton(ebiten.MouseButtonLeft)

	// Initial state should be StillDeactivatedState
	if leftButton.State() != StillDeactivatedState {
		t.Errorf("Expected initial state to be StillDeactivatedState, got %v", leftButton.State())
	}
}

func TestMouseButtonJustActivatedState(t *testing.T) {
	leftButton := MouseButton(ebiten.MouseButtonLeft)

	// Simulate mouse button press
	isMouseButtonPressedFunc = func(button ebiten.MouseButton) bool {
		return button == ebiten.MouseButtonLeft
	}
	Update()

	// After mouse button press, state should be JustActivatedState
	if leftButton.State() != JustActivatedState {
		t.Errorf("Expected state to be JustActivatedState after mouse button press, got %v", leftButton.State())
	}
}

func TestMouseButtonStillActivatedState(t *testing.T) {
	leftButton := MouseButton(ebiten.MouseButtonLeft)

	// Simulate mouse button press
	isMouseButtonPressedFunc = func(button ebiten.MouseButton) bool {
		return button == ebiten.MouseButtonLeft
	}
	Update()
	Update()

	// After two updates, state should be StillActivatedState
	if leftButton.State() != StillActivatedState {
		t.Errorf("Expected state to be StillActivatedState after two updates, got %v", leftButton.State())
	}
}

func TestMouseButtonJustDeactivatedState(t *testing.T) {
	leftButton := MouseButton(ebiten.MouseButtonLeft)

	// Simulate mouse button press and release
	isMouseButtonPressedFunc = func(button ebiten.MouseButton) bool {
		return button == ebiten.MouseButtonLeft
	}
	Update()
	isMouseButtonPressedFunc = func(button ebiten.MouseButton) bool {
		return false
	}
	Update()

	// After mouse button release, state should be JustDeactivatedState
	if leftButton.State() != JustDeactivatedState {
		t.Errorf("Expected state to be JustDeactivatedState after mouse button release, got %v", leftButton.State())
	}
}

func TestMouseButtonStillDeactivatedState(t *testing.T) {
	leftButton := MouseButton(ebiten.MouseButtonLeft)

	// Simulate mouse button press and release
	isMouseButtonPressedFunc = func(button ebiten.MouseButton) bool {
		return button == ebiten.MouseButtonLeft
	}
	Update()
	isMouseButtonPressedFunc = func(button ebiten.MouseButton) bool {
		return false
	}
	Update()
	Update()

	// After another update, state should be StillDeactivatedState
	if leftButton.State() != StillDeactivatedState {
		t.Errorf("Expected state to be StillDeactivatedState after another update, got %v", leftButton.State())
	}
}
