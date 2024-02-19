package resolvers

import "github.com/m-butterfield/heatmap-creator/server/app/data"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DS data.Store
}
