package memory

import (
	"math"
)

// LoadCelData - ряд значений моделирующий рост момента при свинчивании трубы
// каждое следующее значение, значение графика на следующем шагу
type LoadCelData struct {
	torques  []int64
	position int
}

// New() - конструктор
// создаю значения
func New() *LoadCelData {
	torquesArray := make([]int64, 1500)
	for i := range torquesArray {
		if i == 0 {
			torquesArray[i] = 1
		} else if i < 1000 {
			torquesArray[i] = int64(i)
		} else if i < 1400 {
			torquesArray[i] = torquesArray[i-1] + int64(math.Round(float64(torquesArray[i-1])*0.0035))
		} else {
			torquesArray[i] = torquesArray[i-1] + int64(math.Round(float64(torquesArray[i-1])*0.01))
		}
	}

	return &LoadCelData{
		torques:  torquesArray,
		position: 0,
	}
}

func (l *LoadCelData) GetNext() int64 {
	if l.position < len(l.torques) {
		el := l.torques[l.position]
		l.position++
		return el
	} else {
		l.position = 0
		return 0
	}
}

func (l *LoadCelData) Reset() {
	l.position = 0
}
