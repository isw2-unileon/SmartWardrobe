package clip

import (
	"os"
	"runtime"
)

// sharedLibPath returns the ONNX Runtime shared library path.
// Priority:
// 1. ORT_LIB_PATH environment variable
// 2. Bundled runtime in ./lib
// 3. OS default path
func sharedLibPath() string {

	// Use explicit environment override if provided.
	if p := os.Getenv("ORT_LIB_PATH"); p != "" {
		return p
	}

	// Use bundled runtime library located in ./lib.
	switch runtime.GOOS {

	case "darwin":
		if runtime.GOARCH == "arm64" {
			return "./lib/libonnxruntime_arm64.dylib"
		}
		return "./lib/libonnxruntime.dylib"

	case "windows":
		return "./lib/onnxruntime.dll"

	default:
		return "./lib/libonnxruntime.so"
	}
}
