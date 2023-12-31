-- name: GetChord :one
SELECT * FROM chords
WHERE id = $1 LIMIT 1;

-- name: ListChords :many 
SELECT * FROM chords
ORDER BY name;

-- name: CreateChord :one
INSERT INTO chords (
  name, root
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateChord :exec
UPDATE chords
  set name = $2,
  root = $3
WHERE id = $1;

-- name: DeleteChord :exec
DELETE FROM chords
WHERE id = $1;