package common

const simpleSuccessMessage = true

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
	Sort   interface{} `json:"sort,omitempty"`
}

func FullSuccessResponse(data, filter, paging, sort interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter, Sort: sort}
}

func SuccessResponse(data interface{}) *successRes {
	return FullSuccessResponse(data, nil, nil, nil)

}

func SimpleSuccessResponse() *successRes {
	return FullSuccessResponse(simpleSuccessMessage, nil, nil, nil)
}
