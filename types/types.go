package types

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
