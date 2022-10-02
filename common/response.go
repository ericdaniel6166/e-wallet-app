package common

const simpleSuccessMessage = "data: true"

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func FullSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SuccessResponse(data interface{}) *successRes {
	return FullSuccessResponse(data, nil, nil)

}

func SimpleSuccessResponse() *successRes {
	return FullSuccessResponse(simpleSuccessMessage, nil, nil)
}
