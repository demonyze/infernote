# 🔥 Infernote 🎵

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
--lang=de # select output language - defaults to "en"
--name=chords.json # output json filename - defaults to "infernote.json"
--path=my_dir # output path - defaults to "output"
```

### Language Support

By default all data will be generated in english.
This can be changed by providing the `lang` parameter in the generate command:

```bash
infernote --lang=de
```

All supported languages can be found in the `/lang` directory.

> [!NOTE]
> Feel free to add new language files or update existing ones via pull requests.

### Credit

This project uses data from the following resources:

- [chords-db](https://github.com/tombatossals/chords-db)
- [chord.rocks](https://chord.rocks)
