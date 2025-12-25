# gpassword

A small, lightweight password generator web app written in Go using Echo and HTMX, styled with TailwindCSS. It serves a tiny UI to generate secure random passwords and allows toggling which symbols are included.

## Features

- Generate secure, random passwords (configurable symbol set)
- Small server written in Go with HTML templates and HTMX partial updates
- TailwindCSS for styling; simple frontend with copy-to-clipboard

## Quick Start

Prerequisites:

- Go >= 1.22
- pnpm >= 10.0

Install dependencies:

```bash
pnpm install
```

Run in development (starts Go server and rebuilds Tailwind CSS):

```bash
pnpm dev
```

Open http://localhost:8080 in your browser.

Build for production

```bash
pnpm run build:css   # build Tailwind CSS once
go build -o gpassword gpassword.go
./gpassword
```

## Endpoints

- `GET /` - Main page (index) with the password generator UI
- `GET /password` - Returns the password input partial (used by HTMX)
- `POST /password` - Generates a password using selected symbol choices

## Configuration & Usage

The generator exposes a set of symbols in the settings panel. Uncheck any symbol you don't want included and click the regenerate icon to get a new password without those symbols.

Password generation configuration is handled server-side in `internal/generator`.

## Project structure (important files)

- `gpassword.go` — application entrypoint and HTTP routes
- `internal/generator/generator.go` — password generation logic
- `views/` — HTML templates used by the server (HTMX partials)
- `static/` — generated CSS and other static assets
- `package.json` — frontend build scripts (Tailwind + concurrently)

See [gpassword.go](gpassword.go) for the server setup and route handlers.

## Development notes

- The dev script (`pnpm dev`) runs both the Go server and the Tailwind watcher concurrently. If you prefer, run the Go server separately with:

```bash
go run gpassword.go
```

and build/watch CSS separately:

```bash
pnpm run dev:css
```

## Security & randomness

The generator uses `crypto/rand` to pick characters from a charset. The default password length is defined in `internal/generator/generator.go`.

## Contributing

Feel free to open issues or submit pull requests. Small improvements that would be useful:

- Add tests for the generator logic
- Allow configurable password length via UI
- Improve accessibility and internationalization

## License

MIT — see LICENSE.md

---

Author: Matteo Veraldi <mattveraldi@gmail.com>
