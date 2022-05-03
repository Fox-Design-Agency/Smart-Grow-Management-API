package routemodels

import "github.com/lib/pq"

type SlimNutrient struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Nutrient struct {
	SlimNutrient
	CreatedAt pq.NullTime `json:"createdAt" db:"created_at"`
	CreatedAtMember
	UpdatedAt pq.NullTime `json:"updatedAt" db:"updated_by"`
	UpdatedAtMember
}
