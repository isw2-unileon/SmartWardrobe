#!/usr/bin/env bash
# =============================================================================
# download_models.sh
#
# Downloads:
#   - ONNX Runtime shared library
#   - CLIP ViT-B/32 ONNX models
#
# Compatible with:
#   github.com/yalue/onnxruntime_go
#
# Usage:
#   ./scripts/download_models.sh
# =============================================================================

set -Eeuo pipefail

# -----------------------------------------------------------------------------
# Configuration
# -----------------------------------------------------------------------------

MODELS_DIR="./models"
LIB_DIR="./lib"

ORT_VERSION="1.25.0"

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

# -----------------------------------------------------------------------------
# Logging
# -----------------------------------------------------------------------------

log() {
    echo "[INFO] $1"
}

success() {
    echo "[OK]   $1"
}

error() {
    echo "[ERROR] $1" >&2
    exit 1
}

# -----------------------------------------------------------------------------
# Dependency checks
# -----------------------------------------------------------------------------

require_command() {
    command -v "$1" >/dev/null 2>&1 || \
        error "Missing required command: $1"
}

require_command curl
require_command tar

mkdir -p "$MODELS_DIR" "$LIB_DIR"

# -----------------------------------------------------------------------------
# Cleanup
# -----------------------------------------------------------------------------

TMP_DIR="$(mktemp -d)"

cleanup() {
    rm -rf "$TMP_DIR"
}

trap cleanup EXIT

# -----------------------------------------------------------------------------
# Download helper
# -----------------------------------------------------------------------------

download_if_missing() {
    local output="$1"
    local url="$2"
    local description="$3"

    if [[ -f "$output" ]]; then
        log "$description already exists"
        return
    fi

    log "Downloading $description..."

    curl -L --fail --progress-bar \
        "$url" \
        -o "$output"

    success "$description downloaded"
}

# -----------------------------------------------------------------------------
# Banner
# -----------------------------------------------------------------------------

echo "=================================================="
echo " CLIP Fashion Classifier — Environment Setup"
echo "=================================================="
echo " OS:      $OS"
echo " ARCH:    $ARCH"
echo " ORT:     $ORT_VERSION"
echo "=================================================="
echo

# -----------------------------------------------------------------------------
# ONNX Runtime
# -----------------------------------------------------------------------------

log "Preparing ONNX Runtime..."

case "$OS" in
    linux)
        if [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
            ORT_URL="https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VERSION}/onnxruntime-linux-aarch64-${ORT_VERSION}.tgz"
        else
            ORT_URL="https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VERSION}/onnxruntime-linux-x64-${ORT_VERSION}.tgz"
        fi
        ORT_DEST="$LIB_DIR/libonnxruntime.so"
        ;;
    darwin)
        if [[ "$ARCH" == "arm64" ]]; then
            ORT_URL="https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VERSION}/onnxruntime-osx-arm64-${ORT_VERSION}.tgz"
            ORT_DEST="$LIB_DIR/libonnxruntime_arm64.dylib"
        else
            ORT_URL="https://github.com/microsoft/onnxruntime/releases/download/v${ORT_VERSION}/onnxruntime-osx-x86_64-${ORT_VERSION}.tgz"
            ORT_DEST="$LIB_DIR/libonnxruntime.dylib"
        fi
        ;;
    *)
        error "Unsupported operating system: $OS"
        ;;
esac

if [[ ! -f "$ORT_DEST" ]]; then
    log "Downloading ONNX Runtime..."

    curl -L --fail --progress-bar \
        "$ORT_URL" \
        -o "$TMP_DIR/ort.tgz"

    tar -xzf "$TMP_DIR/ort.tgz" -C "$TMP_DIR"

    REAL_LIB=$(find "$TMP_DIR" -type f -name "libonnxruntime.so*" \
        | sort -V \
        | tail -1)

    cp "$REAL_LIB" "$ORT_DEST"

    success "ONNX Runtime installed"
else
    log "ONNX Runtime already installed"
fi

# -----------------------------------------------------------------------------
# CLIP Models
# -----------------------------------------------------------------------------

log "Preparing CLIP models..."

download_if_missing \
    "$MODELS_DIR/clip_image_encoder.onnx" \
    "https://huggingface.co/Xenova/clip-vit-base-patch32/resolve/main/onnx/vision_model.onnx" \
    "CLIP vision encoder"

download_if_missing \
    "$MODELS_DIR/clip_text_encoder.onnx" \
    "https://huggingface.co/Xenova/clip-vit-base-patch32/resolve/main/onnx/text_model.onnx" \
    "CLIP text encoder"

# -----------------------------------------------------------------------------
# Validation
# -----------------------------------------------------------------------------

[[ -f "$MODELS_DIR/clip_image_encoder.onnx" ]] || \
    error "Missing clip_image_encoder.onnx"

[[ -f "$MODELS_DIR/clip_text_encoder.onnx" ]] || \
    error "Missing clip_text_encoder.onnx"

[[ -f "$ORT_DEST" ]] || \
    error "Missing ONNX Runtime library"

# -----------------------------------------------------------------------------
# Summary
# -----------------------------------------------------------------------------

echo
echo "=================================================="
success "Environment setup completed"
echo "=================================================="
echo
echo "Next steps:"
echo "  go mod tidy"
echo "  go run cmd/server/main.go"
echo