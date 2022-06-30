package constants

import "errors"

var (
	ErrCreateCanvas = errors.New("error creating canvas")
	ErrGetCanvases  = errors.New("error getting canvases")
	ErrGetCanvas    = errors.New("error getting canvas")
	ErrUpdateCanvas = errors.New("error updating canvas")
	ErrDeleteCanvas = errors.New("error deleting canvas")
	ErrCreateShape  = errors.New("error creating shape")
	ErrGetShapes    = errors.New("error getting shapes")
	ErrGetShape     = errors.New("error getting shape")
	ErrUpdateShape  = errors.New("error updating shape")
	ErrDeleteShape  = errors.New("error deleting shape")
	ErrInvalidShape = errors.New("error invalid shape")
)
