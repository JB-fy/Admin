package internal

type MyGenDataHandleMethod = uint

const (
	ReturnEmpty    MyGenDataHandleMethod = 0  //默认返回空
	ReturnType     MyGenDataHandleMethod = 1  //返回根据字段数据类型解析的数据
	ReturnTypeName MyGenDataHandleMethod = 2  //返回根据字段命名类型解析的数据
	ReturnUnion    MyGenDataHandleMethod = 10 //返回两种类型解析的数据
)

type MyGenDataSliceHandler struct {
	Method       MyGenDataHandleMethod //根据该字段返回解析的数据
	DataType     []string              //根据字段数据类型解析的数据
	DataTypeName []string              //根据字段命名类型解析的数据
}

func (myGenDataSliceHandlerThis *MyGenDataSliceHandler) GetData() []string {
	switch myGenDataSliceHandlerThis.Method {
	case ReturnType:
		return myGenDataSliceHandlerThis.DataType
	case ReturnTypeName:
		return myGenDataSliceHandlerThis.DataTypeName
	case ReturnUnion:
		return append(myGenDataSliceHandlerThis.DataType, myGenDataSliceHandlerThis.DataTypeName...)
	default:
		return nil
	}
}

type MyGenDataStrHandler struct {
	Method       MyGenDataHandleMethod //根据该字段返回解析的数据
	DataType     string                //根据字段数据类型解析的数据
	DataTypeName string                //根据字段命名类型解析的数据
}

func (myGenDataStrHandlerThis *MyGenDataStrHandler) GetData() string {
	switch myGenDataStrHandlerThis.Method {
	case ReturnType:
		return myGenDataStrHandlerThis.DataType
	case ReturnTypeName:
		return myGenDataStrHandlerThis.DataTypeName
	case ReturnUnion:
		return myGenDataStrHandlerThis.DataType + myGenDataStrHandlerThis.DataTypeName
	default:
		return ``
	}
}
