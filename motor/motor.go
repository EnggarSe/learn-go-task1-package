package motor

type Maju interface {
	Cepat(jarak float32) float32
	Lambat(jarak float32) float32
}

type Motor struct {
	jumlahBan  int
	jumlahGigi int
	warna      string
}

var waktuMotor float32

func (m *Motor) WaktuTempuh(jarak float32) float32 {
	waktuMotor = jarak * 0.5
	return waktuMotor
}

func (m *Motor) Cepat(jarak float32) float32 {
	return waktuMotor * 0.5
}
func (m *Motor) Lambat(jarak float32) float32 {
	return waktuMotor * 2

}

func CreateMotor(ban, gigi int, warna string) *Motor {
	return &Motor{
		jumlahBan:  ban,
		jumlahGigi: gigi,
		warna:      warna,
	}
}
