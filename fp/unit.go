package fp

type Unit struct {
}

var instance Unit = Unit{}

func GetUnit() Unit {
	return instance
}
