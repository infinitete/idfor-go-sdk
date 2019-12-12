package fcuim

type Method struct {
	Name   string
	Method byte
}

// 注册全网唯一标识方法
// 例如注册以身份证为线下映射的标识方法、注册结婚证为线下映射标识的方法。
func RegisterMethod(name string) (*byte, error) {
	return nil, nil
}

// 获取已注册的场景协议
func Methods() ([]Method, error) {

	// 通过智能合约获取

	return []Method{}, nil
}
