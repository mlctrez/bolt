package bolt

var _ ValueReceiver = (*Value)(nil)
var _ ValueProvider = (*Value)(nil)

type Value struct {
	K Key
	V []byte
}

func (v *Value) Key() Key {
	return v.K
}

func (v *Value) Value() ([]byte, error) {
	return v.V, nil
}

func (v *Value) SetValue(bytes []byte) error {
	v.V = bytes
	return nil
}
