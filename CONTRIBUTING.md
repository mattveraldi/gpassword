# Contributing to gpassword

Thanks for your interest in contributing! This document describes the preferred process for reporting issues, proposing changes, and submitting code to the project.

## Good first steps

- Read the project `README.md` to understand goals and how to run the app locally.
- Search existing issues to avoid duplicates before opening a new one.
- If you're unsure about a larger change, open an issue to discuss the approach first.

## Reporting issues

- Use clear, descriptive titles and include steps to reproduce the problem.
- Provide environment details where relevant (OS, Go version, browser, etc.).
- Attach logs or screenshots where helpful.

## Pull request process

1. Fork the repository and create a descriptive branch (e.g. `feature/add-length-setting`).
2. Keep changes focused and small; one feature or fix per PR.
3. Rebase or merge the latest `main` before opening the PR to keep history clean.
4. In the PR description, explain the problem, the proposed change, and any migration or manual steps required to test.
5. Link related issues in the PR description (e.g., `Fixes #12`).

## Code style and guidelines

- Use idiomatic Go for backend code. Follow `gofmt`/`gofumpt` formatting.
- Keep functions small and focused; prefer readability over cleverness.
- When editing templates in `views/`, keep markup simple and avoid inline JS unless necessary.
- When adding frontend CSS, prefer Tailwind utility classes already used in the repo.

## Tests

- There are currently no automated tests in the repository. Adding unit tests for `internal/generator` is a welcome improvement.
- If you add tests, ensure they run locally with `go test ./...` and keep them fast and deterministic.

## Local development

Run the development environment (Go server + Tailwind rebuild) with pnpm:

```bash
pnpm install
pnpm dev
```

Or run components separately:

```bash
# run Go server
go run gpassword.go

# build/watch Tailwind
pnpm run dev:css
```

Open `http://localhost:8080` to view the app.

## Security

- Avoid committing secrets or credentials. If you need to provide configuration, prefer environment variables and document them in `README.md`.
- The password generator uses `crypto/rand`. If you change randomness logic, please justify the approach in the PR.

## Accessibility and internationalization

- Keep accessibility in mind when modifying templates. Use semantic markup and provide labels for form controls.
- If you add text visible to users, consider how it will be translated in the future.

## License

By contributing, you agree that your contributions will be licensed under the project's MIT license (see `LICENSE.md`).

## Contact

For questions or design discussions, open an issue or contact the maintainer: Matteo Veraldi <mattveraldi@gmail.com>
