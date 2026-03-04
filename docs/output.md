## Output

The output is a single self-contained JSON file. All requested languages are embedded inside each chord and type under a `lang` map, so consumers can support multiple locales without re-generating the file.

### Top-level shape

| Field       | Type                        | Description                                      |
|-------------|-----------------------------|--------------------------------------------------|
| `languages` | `string[]`                  | Language codes that were requested (e.g. `["en", "de"]`) |
| `types`     | `{ [id]: ChordType }`       | Map of chord type objects keyed by suffix        |
| `chords`    | `{ [id]: Chord }`           | Map of chord objects keyed by chord ID           |

### ChordType

| Field  | Type                  | Description                                        |
|--------|-----------------------|----------------------------------------------------|
| `id`   | `string`              | Chord type suffix (e.g. `"major"`, `"min"`, `"7"`) |
| `lang` | `{ [lang]: string }`  | Localized type name keyed by language code         |

### Chord

| Field             | Type                      | Description                                       |
|-------------------|---------------------------|---------------------------------------------------|
| `id`              | `string`                  | Unique chord identifier (e.g. `"cmajor"`)         |
| `nameShort`       | `string`                  | Abbreviated chord name (e.g. `"Cma"`)             |
| `lang`            | `{ [lang]: ChordLang }`   | Localized chord data keyed by language code       |
| `notes`           | `Note[]`                  | Notes that make up the chord                      |
| `root`            | `Note`                    | Root note of the chord                            |
| `guitarPositions` | `GuitarPosition[]`        | Guitar fingering positions                        |

### ChordLang

| Field  | Type     | Description                            |
|--------|----------|----------------------------------------|
| `name` | `string` | Full localized chord name (e.g. `"C Major"`, `"C Dur"`) |
| `type` | `string` | Localized type name (e.g. `"Major"`, `"Dur"`)           |

### Note

| Field  | Type     | Description              |
|--------|----------|--------------------------|
| `name` | `string` | Note name (e.g. `"C"`)  |

### GuitarPosition

| Field      | Type      | Description                                  |
|------------|-----------|----------------------------------------------|
| `frets`    | `int[]`   | Fret numbers per string (`-1` = muted)       |
| `fingers`  | `int[]`   | Finger assignments per string                |
| `barres`   | `int[]`   | Barre fret numbers                           |
| `capo`     | `bool`    | Whether a capo is used                       |
| `baseFret` | `int`     | Starting fret of the position diagram        |
| `midi`     | `int[]`   | MIDI note numbers for the sounding strings   |

---

### Example — single language (`--lang=en`)

```json
{
  "languages": ["en"],
  "types": {
    "major": {
      "id": "major",
      "lang": { "en": "Major" }
    }
  },
  "chords": {
    "cmajor": {
      "id": "cmajor",
      "nameShort": "Cma",
      "lang": {
        "en": { "name": "C Major", "type": "Major" }
      },
      "notes": [{ "name": "C" }, { "name": "E" }, { "name": "G" }],
      "root": { "name": "C" },
      "guitarPositions": [
        {
          "frets": [-1, 3, 2, 0, 1, 0],
          "fingers": [0, 3, 2, 0, 1, 0],
          "barres": [],
          "capo": false,
          "baseFret": 1,
          "midi": [48, 52, 55, 60, 64]
        }
      ]
    }
  }
}
```

### Example — multiple languages (`--lang=en,de`)

```json
{
  "languages": ["en", "de"],
  "types": {
    "major": {
      "id": "major",
      "lang": { "en": "Major", "de": "Dur" }
    }
  },
  "chords": {
    "cmajor": {
      "id": "cmajor",
      "nameShort": "Cma",
      "lang": {
        "en": { "name": "C Major", "type": "Major" },
        "de": { "name": "C Dur",   "type": "Dur"   }
      },
      "notes": [{ "name": "C" }, { "name": "E" }, { "name": "G" }],
      "root": { "name": "C" },
      "guitarPositions": [
        {
          "frets": [-1, 3, 2, 0, 1, 0],
          "fingers": [0, 3, 2, 0, 1, 0],
          "barres": [],
          "capo": false,
          "baseFret": 1,
          "midi": [48, 52, 55, 60, 64]
        }
      ]
    }
  }
}
```
