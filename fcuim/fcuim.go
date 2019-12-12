package fcuim

const Idfor = "idfor"
const Fcuim = "fcuim"

const (
	IdentifierID = byte(0x0) // 身份证号码
)

// 验证一个全网唯一标识是否是有效的
func IsFcuimValid(fcuim []byte) bool {
	if len(fcuim) < 14 {
		return false
	}

	m := fcuim[13]

	if m != IdentifierID {
		return false
	}

	return true
}
