package web

// UpdateUserTagForm
type UpdateUserTagForm struct {
	UserID   int64  `json:"userid"`
	TagNames string `json:"tagNames"`
}
