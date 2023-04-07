package testLib

import "runtime"

func Version() string {
	return runtime.Version()
}
