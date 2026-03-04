# 🔥 Infernote 🎵
[![Go Reference](https://pkg.go.dev/badge/github.com/demonyze/infernote.svg)](https://pkg.go.dev/github.com/demonyze/infernote)

Generate chords data from different resources in a custom JSON file.

## Getting Started

Install the `infernote` binary to your `$GOPATH/bin` directory

```bash
go install github.com/demonyze/infernote/cmd/infernote@latest
```

After that the `infernote` command should be available in your terminal.

```bash
infernote
```

## Parameters

```bash
--lang=en,de  # select output languages (comma-separated) - defaults to "en"
--name=chords.json  # output json filename - defaults to "infernote.json"
--path=my_dir  # output path - defaults to "output"
```

### Language Support

By default all data is generated in English. The output embeds all requested languages in a single file under a `lang` map keyed by language code. Pass multiple languages as a comma-separated list:

```bash
infernote --lang=en
infernote --lang=en,de
```

The output `languages` array reflects which languages were requested. Each chord and type carries a `lang` map so consumers can pick their preferred language at runtime without re-generating the file.

All supported languages can be found in the `/lang` directory.

> [!NOTE]
> Feel free to add new language files or update existing ones via pull requests.

## Output

See [docs/output.md](docs/output.md) for the full output schema and examples.

### Credit

This project uses data from the following resources:

- [chords-db](https://github.com/tombatossals/chords-db)
- [chord.rocks](https://chord.rocks)
