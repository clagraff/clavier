package clavier

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// State represents the toggle-state of a control.
type State int

const (
	StillDeactivatedState State = iota
	JustDeactivatedState
	JustActivatedState
	StillActivatedState
)

var (
	// These functions will be used to check the state of keys and mouse buttons.
	// They are set to the default ebiten functions but can be overridden for testing.
	isKeyPressedFunc         = ebiten.IsKeyPressed
	isMouseButtonPressedFunc = ebiten.IsMouseButtonPressed
)

type Control interface {
	Update()
	State() State
}

type keyComboName string

func makeKeyComboName(keys ...ebiten.Key) keyComboName {
	var names []string
	for _, key := range keys {
		names = append(names, key.String())
	}
	return keyComboName(strings.Join(names, "+"))
}

var keyControls = make(map[ebiten.Key]*keyControl)
var keyComboControls = make(map[keyComboName]*keyComboControl)
var mouseButtonControls = make(map[ebiten.MouseButton]*mouseButtonControl)
var userControls = make(map[Control]Control)

func RegisterCustomControl(c Control) (unregister func()) {
	userControls[c] = c
	return func() {
		delete(userControls, c)
	}
}

type keyControl struct {
	key   ebiten.Key
	state State
}

func (kc *keyControl) Update() {
	kc.state = updateKeyState(kc.key, kc.state)
}

func (kc *keyControl) State() State {
	return kc.state
}

type keyComboControl struct {
	keys  []ebiten.Key
	state State
}

func (kcc *keyComboControl) Update() {
	kcc.state = updateKeyComboState(kcc.keys, kcc.state)
}

func (kcc *keyComboControl) State() State {
	return kcc.state
}

type mouseButtonControl struct {
	button ebiten.MouseButton
	state  State
}

func (mbc *mouseButtonControl) Update() {
	mbc.state = updateMouseButtonState(mbc.button, mbc.state)
}

func (mbc *mouseButtonControl) State() State {
	return mbc.state
}

func Update() {
	for _, kc := range keyControls {
		kc.Update()
	}

	for _, mbc := range mouseButtonControls {
		mbc.Update()
	}

	for _, combo := range keyComboControls {
		combo.Update()
	}

	for _, uc := range userControls {
		uc.Update()
	}
}

func updateKeyState(key ebiten.Key, prevState State) State {
	currentPressed := isKeyPressedFunc(key)

	switch {
	case currentPressed && (prevState == StillDeactivatedState || prevState == JustDeactivatedState):
		return JustActivatedState
	case currentPressed && (prevState == JustActivatedState || prevState == StillActivatedState):
		return StillActivatedState
	case !currentPressed && (prevState == StillActivatedState || prevState == JustActivatedState):
		return JustDeactivatedState
	case !currentPressed && (prevState == StillDeactivatedState || prevState == JustDeactivatedState):
		return StillDeactivatedState
	}

	return prevState
}

func updateMouseButtonState(btn ebiten.MouseButton, prevState State) State {
	currentPressed := isMouseButtonPressedFunc(btn)

	switch {
	case currentPressed && (prevState == StillDeactivatedState || prevState == JustDeactivatedState):
		return JustActivatedState
	case currentPressed && (prevState == JustActivatedState || prevState == StillActivatedState):
		return StillActivatedState
	case !currentPressed && (prevState == StillActivatedState || prevState == JustActivatedState):
		return JustDeactivatedState
	case !currentPressed && (prevState == StillDeactivatedState || prevState == JustDeactivatedState):
		return StillDeactivatedState
	}

	return prevState
}

func updateKeyComboState(keys []ebiten.Key, prevState State) State {
	allPressed := true
	for _, key := range keys {
		if keyControl, ok := keyControls[key]; !ok || (keyControl.state != StillActivatedState && keyControl.state != JustActivatedState) {
			allPressed = false
			break
		}
	}

	switch {
	case allPressed && (prevState == StillDeactivatedState || prevState == JustDeactivatedState):
		return JustActivatedState
	case allPressed && (prevState == JustActivatedState || prevState == StillActivatedState):
		return StillActivatedState
	case !allPressed && (prevState == StillActivatedState || prevState == JustActivatedState):
		return JustDeactivatedState
	case !allPressed && (prevState == StillDeactivatedState || prevState == JustDeactivatedState):
		return StillDeactivatedState
	}

	return prevState
}

func Key(key ebiten.Key) Control {
	if _, exists := keyControls[key]; !exists {
		keyControls[key] = &keyControl{key: key, state: StillDeactivatedState}
	}
	return keyControls[key]
}

func KeyCombo(keys ...ebiten.Key) Control {
	comboName := makeKeyComboName(keys...)
	if _, exists := keyComboControls[comboName]; !exists {
		keyComboControls[comboName] = &keyComboControl{keys: keys, state: StillDeactivatedState}
	}
	return keyComboControls[comboName]
}

func MouseButton(btn ebiten.MouseButton) Control {
	if _, exists := mouseButtonControls[btn]; !exists {
		mouseButtonControls[btn] = &mouseButtonControl{button: btn, state: StillDeactivatedState}
	}
	return mouseButtonControls[btn]
}

func Active(c Control) bool {
	return JustActivated(c) || StillActivated(c)
}

func JustActivated(c Control) bool {
	return c.State() == JustActivatedState
}

func StillActivated(c Control) bool {
	return c.State() == StillActivatedState
}

func JustDeactivated(c Control) bool {
	return c.State() == JustDeactivatedState
}

func StillDeactivated(c Control) bool {
	return c.State() == StillDeactivatedState
}

func Deactivated(c Control) bool {
	return JustDeactivated(c) || StillDeactivated(c)
}

func init() {
	for index := ebiten.KeyA; index < ebiten.KeyMax; index++ {
		keyControls[index] = &keyControl{
			key:   index,
			state: StillDeactivatedState,
		}
	}
	for index := ebiten.MouseButtonLeft; index < ebiten.MouseButtonMiddle; index++ {
		mouseButtonControls[index] = &mouseButtonControl{
			button: index,
			state:  StillDeactivatedState,
		}
	}
}
