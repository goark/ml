# Copilot Instructions for `goark/ml`

## Project purpose

`ml` is a CLI and library that fetches metadata from web pages and formats links
as markdown/wiki/html/csv/json output.

## Design principles

- Keep CLI behavior simple and deterministic.
- Keep output formatting stable for all supported styles.
- Reuse `github.com/goark/webinfo` for metadata extraction.
- Preserve compatibility of exported symbols when possible.

## Error handling

- Use `github.com/goark/errs` for internal error handling.
- Prefer `errs.Wrap`, `errs.Join`, and `errs.WithContext`.
- Keep `errors.Is` compatibility for callers.
- Keep sentinel errors stable (`ErrNullPointer`, `ErrNoImplement`, `ErrInvalidRequest`).
- Include useful context keys such as `url`, `style`, and `path`.

## Fetch behavior

- Fetch and metadata parsing are delegated to `github.com/goark/webinfo`.
- Keep `makelink.New` mapping behavior stable:
  - `Link.URL <- Webinfo.URL`
  - `Link.Location <- Webinfo.Location`
  - `Link.Canonical <- Webinfo.Canonical`
  - `Link.Title <- Webinfo.Title`
  - `Link.Description <- Webinfo.Description`
- Preserve URL fallback behavior for output rendering:
  - title fallback: `Title` -> `URL`
  - URL fallback: `Canonical` -> `Location` -> `URL`

## Output style behavior

- Keep style names and output formats stable:
  - `markdown`: `[title](url)`
  - `wiki`: `[url title]`
  - `html`: `<a href="url">title</a>`
  - `csv`: escaped CSV fields in fixed order
  - `json`: JSON encoded `Link`
- Keep CSV quote escaping behavior unchanged.

## Coding style

- Write idiomatic Go with straightforward control flow.
- Avoid unnecessary dependencies.
- Keep comments concise and in English.

## Testing and validation

- Add or update tests for behavior changes.
- Prefer local validation with Taskfile targets:
  - `task test`
  - `task govulncheck`

## Documentation

- Keep `README.md` aligned with actual CLI and exported API behavior.
- Keep examples concise and runnable.

## Release process

- Create release tags from `master`.
- Use semantic versioning tags in `vMAJOR.MINOR.PATCH` format.
- Ensure repository is clean and synced before tagging.

Release steps:

1. Ensure `master` is up to date.
2. Create annotated tag:
  - `git tag -a vX.Y.Z -m "Release vX.Y.Z"`
3. Push tag:
  - `git push origin vX.Y.Z`
4. Wait for the `build` workflow triggered by the tag push to complete.

Verification steps:

- Check tag exists:
  - `git tag -l "vX.Y.Z"`
- Check `build` workflow result for the tag run (success/failure).
- Check release exists:
  - `gh release view vX.Y.Z`
- Check generated release notes are present in the release page.
