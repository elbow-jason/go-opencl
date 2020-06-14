// +build cl10

package cl

// ArgName is not supported by OpenCL 1.0
func (k *Kernel) ArgName(index int) (string, error) {
	return "", ErrUnsupported
}
