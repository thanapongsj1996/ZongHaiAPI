package models

type JSONResponse struct {
	Status       bool         `json:"status"`
	Data         interface{}  `json:"data,omitempty"`
	Paging       PagingResult `json:"paging,omitempty"`
	ErrorMessage string       `json:"errorMessage,omitempty"`
}

func SuccessResponse(response JSONResponse) JSONResponse {
	response.Status = true
	return response
}

func ErrorResponse(response JSONResponse, message string) JSONResponse {
	response.Status = false
	response.ErrorMessage = message
	return response
}

type PagingResult struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	PrevPage  int   `json:"prevPage"`
	NextPage  int   `json:"nextPage"`
	Count     int64 `json:"count"`
	TotalPage int   `json:"totalPage"`
}

func (j *JSONResponse) AddPaging(p *PagingResult) {
	j.Paging = *p
}
