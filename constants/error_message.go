package constants

import "errors"

var (
	ErrCreateCanvas = errors.New("Error creating canvas")
	ErrGetCanvases  = errors.New("Error getting canvases")
	ErrGetCanvas    = errors.New("Error getting canvas")
	ErrUpdateCanvas = errors.New("Error updating canvas")
	ErrDeleteCanvas = errors.New("Error deleting canvas")
	ErrCreateShape  = errors.New("Error creating shape")
	ErrGetShapes    = errors.New("Error getting shapes")
	ErrGetShape     = errors.New("Error getting shape")
	ErrUpdateShape  = errors.New("Error updating shape")
	ErrDeleteShape  = errors.New("Error deleting shape")
	ErrInvalidShape = errors.New("Error invalid shape")
)
