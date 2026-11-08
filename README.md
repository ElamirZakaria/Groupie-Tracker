# Groupie Tracker

A web application built with Go that consumes and visualizes data from the Groupie Trackers API. The app showcases famous music artists, their career milestones, band members, concert locations, and tour schedules, enriched with sample previews from the iTunes Search API.

---

## Features

- **Artist Showcase**: Displays a responsive grid of artists featuring high-resolution imagery, founding year, and first album release.
- **Detailed Artist Profiles**: Dedicated profile views showing active years, band member lists, and organized tour dates mapped to their respective locations.
- **Top Songs Preview**: Integrated track sampler leveraging the iTunes Search API to listen to audio snippets directly in the browser.
- **Concurrent API Ingestion**: Uses Go's concurrency primitives (`sync.WaitGroup` and goroutines) to aggregate data simultaneously across multiple endpoints (`artists`, `locations`, `dates`, and `relation`).
- **Custom Error Handling**: Styled, contextual HTTP error pages for `400 Bad Request`, `404 Not Found`, `405 Method Not Allowed`, and `500 Internal Server Error`.

---

## Tech Stack

- **Backend**: Go (Standard Library: `net/http`, `html/template`, `encoding/json`, `sync`)
- **Frontend**: HTML5, CSS3 (Modern CSS Grid, Flexbox, Custom Properties)
- **External APIs**:
  - [Groupie Trackers API](https://groupietrackers.01edu.ai/api)
  - [Apple iTunes Search API](https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/index.html)

---

## Project Structure

```text
groupie-tracker/
├── api/
│   └── api.go              # Data fetching, goroutine orchestration, and iTunes integration
├── handler/
│   └── handler.go          # HTTP handlers, routing logic, and template rendering
├── models/
│   └── models.go           # Struct definitions for API responses and application state
├── static/
│   └── css/
│       └── style.css       # Global styles, layout, and error card themes
├── templates/
│   ├── artist.html         # Artist detail and tour schedule view
│   ├── error.html          # Standardized error page template
│   └── index.html          # Homepage artist grid template
├── go.mod                  # Go module definition
├── main.go                 # Application entrypoint and router initialization
└── README.md
```

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) version **1.20** or higher installed on your machine.

### Installation & Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/groupie-tracker.git
   cd groupie-tracker
   ```

2. **Verify dependencies:**
   ```bash
   go mod tidy
   ```

3. **Start the server:**
   ```bash
   go run main.go
   ```

4. **Access the application:**
   Open your browser and navigate to:
   ```text
   http://localhost:8080
   ```

---

## Application Endpoints

| Route | Method | Description |
| :--- | :--- | :--- |
| `/` | `GET` | Homepage showing all artists. |
| `/artist?id={id}` | `GET` | Profile page for a specific artist by ID. |
| `/track?id={id}` | `GET` | Artist page with expanded iTunes audio track previews. |
| `/static/*` | `GET` | Static asset server (CSS, images). |

---

## License

This project is created for educational purposes as part of the software engineering curriculum.