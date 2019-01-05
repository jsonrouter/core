package openapi

func pointerFloat64(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}
