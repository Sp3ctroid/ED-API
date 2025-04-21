package types

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Is_admin bool   `json:"is_admin,omitempty"`
}
type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type JSON_Status struct {
	Status string `json:"status"`
	Action string `json:"action"`
	Obj    any    `json:"object"`
}

func (j *JSON_Status) Response(st string, ac string, ob any) {
	j.Status = st
	j.Action = ac
	j.Obj = ob
}
