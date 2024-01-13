# Infernote

⚠️ Still heavily work in progress.

```bash
make chords
```

Generates a JSON file containing chords data.

## Output

### As Map (default)

```json
{
  "cmajor": {
    "id": "cmajor",
    "name": "C Major",
    "type": "major",
    "alternative_names": "C",
    "notes": [{ "name": "C" }, { "name": "E" }, { "name": "G" }],
    "root": { "name": "C" },
    "guitar_positions": [
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
```

### As Array

```json
[
  {
    "id": "cmajor",
    "name": "C Major",
    "type": "major",
    "alternative_names": "C",
    "notes": [{ "name": "C" }, { "name": "E" }, { "name": "G" }],
    "root": { "name": "C" },
    "guitar_positions": [
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
]
```

### Credit

This project uses data from the following resources:

- [chords-db](https://github.com/tombatossals/chords-db)
