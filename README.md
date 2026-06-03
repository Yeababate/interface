# 🎨 ART Interface

A web interface for encoding and decoding text-based art strings. Built with Go and styled with Tailwind CSS.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running the Server](#running-the-server)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Project Structure](#project-structure)
- [Notes](#notes)

---

## Overview

ART Interface is a web application that wraps the ART encoder/decoder in a clean, browser-based interface. Instead of using the command line, artists can paste their encoded strings into a text box, click a button, and instantly see the resulting artwork.

The server is written in Go and uses Go's `html/template` package to render the interface. Tailwind CSS is loaded via CDN for styling.

---

## Requirements

- [Go](https://go.dev/dl/) 1.21 or higher

---

## Installation

Clone the repository and navigate into it:

```bash
git clone https://gitea.kood.tech/yeabsiraabateguangul/interface.git
cd interface
```

---

## Running the Server

```bash
go run .
```

The server starts on port **8080**. Open your browser and go to:

```
http://localhost:8080
```

---

## Usage

1. Open `http://localhost:8080` in your browser
2. Paste your ART string into the text box
3. Click **Decode** to convert an encoded string into artwork
4. Click **Encode** to convert artwork back into an encoded string
5. The result is displayed below the input along with the HTTP status code

If your input is invalid, a dialog will appear on screen with an error message. Your input is preserved so you can correct it without retyping.



---

## Endpoints

| Method | Route | Description | Success | Error |
|--------|-------|-------------|---------|-------|
| `GET` | `/` | Returns the main web interface | `200 OK` | — |
| `POST` | `/art` | Encodes or decodes a submitted ART string | `202 Accepted` | `400 Bad Request` |

The `POST /art` endpoint reads two form fields:

| Field | Values | Description |
|-------|--------|-------------|
| `artToCode` | any string | The input string to encode or decode |
| `action` | `encode` or `decode` | Which operation to perform |

Any route not listed above returns `404 Not Found`.

---

## Project Structure

```
art-interface/
├── main.go              # Server entry point, route handlers, and data structs
├── decode.go            # decodes the given string
├── encode.go            # encodes the given string
├── templates/
│   └── index.html       # Main HTML template rendered by Go
└── README.md
```

### Key files

**`main.go`** — defines three things:
- `Data` struct — holds `Output`, `OutputStatus`, and `ArtToCode` passed to the template
- `homeHandler` — serves `GET /` and renders the page with empty data
- `codeHandler` — handles `POST /art`, runs encode or decode, and re-renders the page with the result

**`templates/index.html`** — the web interface. Uses Go template syntax (`{{ .Output }}`, `{{ .OutputStatus }}` etc.) to display results returned by the server.

---

## Notes
- The encoder only supports characters that repeat **1 or 2 times**