package models

type (
	TileGeneratorFunction func(params any) ([]GeneratedTile, error)

	GeneratedTile struct {
		Label  string
		Params any
	}
)
