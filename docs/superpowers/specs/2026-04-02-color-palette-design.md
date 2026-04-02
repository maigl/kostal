# Design: Runtime Color Palette Configuration

**Date:** 2026-04-02
**Status:** Approved

## Overview

Add runtime color palette configuration to the Kostal solar dashboard, allowing users to change colors without redeploying. Users can paste coolors.co color codes and apply them via an API call or startup flag.

## Configuration

### Config File

- Controlled via `KOSTAL_CONFIG` env var (default: `config.json` in working directory)
- Format:
```json
{
  "palette": "c41b5c-08415c-6b818c-f1bf98"
}
```

### Startup Options

1. **Flag:** `--palette c41b5c-08415c-6b818c-f1bf98`
2. **Env var:** `KOSTAL_PALETTE=c41b5c-08415c-6b818c-f1bf98`
3. **Config file:** Written on first runtime update

Flag and env var take precedence over config file.

### Docker Compose Usage

```yaml
environment:
  - KOSTAL_PALETTE=c41b5c-08415c-6b818c-f1bf98
```

## API

### POST /colors

Update palette at runtime.

**Request:**
```
POST /colors
Content-Type: text/plain

c41b5c-08415c-6b818c-f1bf98-eee5e9
```

**Response:** Empty 200 OK

**Behavior:**
1. Parse hex codes from body
2. Update in-memory palette
3. Write to config file (if KOSTAL_CONFIG set)

### Color Mapping

- Parse hex codes from input (6-character hex, with or without `#`)
- Support 4-5 codes (dash-separated or comma-separated)
- Map first 4 codes to CSS variables: `--color1`, `--color2`, `--color3`, `--color4`
- Ignore 5th code if present

## HTML Generation

- `web/frame.html` uses Go templates with CSS variables
- On each request, parse palette and inject `--color1` through `--color4` into the `:root` CSS block
- Fallback to hardcoded colors if no palette configured

## Implementation Notes

- Use Go's `text/template` to inject CSS variables
- Config struct in `pkg/config/config.go`
- Handler in `pkg/handler/handler.go`
- Update HTML template to accept palette as template data
