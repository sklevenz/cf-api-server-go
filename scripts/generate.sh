#!/usr/bin/env bash


# Dateien definieren
CONFIG_FILE="./config.yaml"
API_SPEC="./spec/openapi.yaml"
OUTPUT_DIR="./internal/gen"
OUTPUT_FILE="$OUTPUT_DIR/api_generated.go"

# Stelle sicher, dass das Zielverzeichnis existiert
mkdir -p "$OUTPUT_DIR"

# Prüfe, ob oapi-codegen installiert ist
if ! command -v oapi-codegen &> /dev/null; then
    echo "Fehler: oapi-codegen ist nicht installiert. Installiere es mit:"
    echo "  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest"
    exit 1
fi

# Prüfe, ob die API-Spezifikation existiert
if [ ! -f "$API_SPEC" ]; then
    echo "Fehler: OpenAPI-Spezifikationsdatei nicht gefunden: $API_SPEC"
    exit 1
fi

# Prüfe, ob die Config existiert
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Fehler: Config-Datei nicht gefunden: $CONFIG_FILE"
    exit 1
fi

# Führe die Code-Generierung aus
echo "Generiere API-Code..."
oapi-codegen --config="$CONFIG_FILE" "$API_SPEC" > "$OUTPUT_FILE"

if [ $? -eq 0 ]; then
    echo "✅ API-Code erfolgreich generiert: $OUTPUT_FILE"
else
    echo "❌ Fehler bei der Code-Generierung!"
    exit 1
fi

echo $OUTPUT_FILE