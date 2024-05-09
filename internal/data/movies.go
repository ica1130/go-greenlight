package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"` // If we want to convert this field to a string in JSON, this is how it's done.
	Genres    []string  `json:"genres,omitempty"`  // converstion to string will work only on int*, unit*, float* or bool* types. on others it will have no effect.
	Version   int32     `json:"version"`
}
