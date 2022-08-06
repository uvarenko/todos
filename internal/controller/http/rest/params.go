package rest

// get todos
type getTodosParams struct {
	Offset int `form:"offset" json:"offset"`
	Size   int `form:"size" json:"size"`
}

// patch todo
type updateAction string

const (
	done   updateAction = "/done"
	delete updateAction = "/delete"
)
