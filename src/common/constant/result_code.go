package constant

type ResponseCode int
type ResponseMsg string

const (
	SelectSuccessCode ResponseCode = 2005
	SelectSuccessMsg  ResponseMsg  = "查询成功"

	DataIsNilCode ResponseCode = 3007
	DataIsNilMsg  ResponseMsg  = "未找到符合条件的数据"
)
