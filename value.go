package bolt

var _ ValueReceiver = (*Value)(nil)
var _ ValueProvider = (*Value)(nil)

type Value struct {
	key   Key
	value []byte
}

func (v *Value) Key() Key {
	return v.key
}

func (v *Value) Value() ([]byte, error) {
	return v.value, nil
}

func (v *Value) SetValue(bytes []byte) error {
	v.value = bytes
	return nil
}
