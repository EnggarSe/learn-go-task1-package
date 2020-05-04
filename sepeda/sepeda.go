package sepeda

type Maju interface {
	Cepat(jarak float32) float32
	Lambat(jarak float32) float32
}

var waktuSepeda float32

type Sepeda struct {
	jumlahBan  int
	jumlahGigi int
	warna      string
}

func (s *Sepeda) WaktuTempuh(jarak float32) float32 {
	waktuSepeda = jarak * 2.5
	return waktuSepeda
}

func (m *Sepeda) Cepat(jarak float32) float32 {
	return waktuSepeda * 0.5
}

func (m *Sepeda) Lambat(jarak float32) float32 {
	return waktuSepeda * 2
}

func CreateSepeda(ban, gigi int, warna string) *Sepeda {
	return &Sepeda{
		jumlahBan:  ban,
		jumlahGigi: gigi,
		warna:      warna,
	}
}
