package storages

type Tasks struct {
	Id string `json:"id"`
	Content string `json:"content"`
	UserId string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type Users struct {
	Id string `json:"id"`
	Password string `json:"password"`
	MaxTodo int `json:"max_todo"`
}

type RestResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}