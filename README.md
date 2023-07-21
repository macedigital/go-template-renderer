# go-template-renderer

Small utility for testing go template rendering with a super-fast feedback loop

## Usage

No additional dependencies are required, only standard library is used. Just build the binary like this `go build -o tpl-renderer main.go`.

Basic usage is `./tpl-renderer <PATH_TO_TEMPLATE> <PATH_TO_VALUES>.json` and output will be the rendered template, or the last error that occurred.

### Example

Render `examples/example.tpl.yaml` using values coming from `examples/example.json` and linting the output with `yamllint`.

```sh
./tpl-renderer examples/example.tpl.yaml examples/example.json | yamllint -
```

## License

MIT
