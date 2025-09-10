package multiplier

import (
	"math"
	"math/rand"
	"time"
)

type MultiplierService struct {
	Alpha float64
}

func NewMultiplierService(alpha float64) *MultiplierService {
	rand.Seed(time.Now().UnixNano())
	return &MultiplierService{
		Alpha: alpha,
	}
}

// Генерирует число по распределению Парето  
func (service *MultiplierService) GenerateMultiplierPareto(xmin, xmax float64) float64 {
	u := rand.Float64()
	// Формула: m = xmin / u^(1/alpha), где u ~ Uniform(0,1)
	m := xmin / math.Pow(u, 1.0/service.Alpha) // Сама формула - Обратная CDF (функция распределения) распределения Парето

	if m > xmax {  // Если мультипликатор получился больше 10_000 -> присваиваем верхнюю границу т.е 10_000
		m = xmax
	}
	return m
}

// func CheckExpectedRTP(alpha float64) float64 {
// 	const N = 100000
// 	var totalsurv float64

// 	for i := 0; i < N; i++ {
// 		x := 1.0 + rand.Float64()*9999.0
// 		totalsurv += math.Pow(1.0/x, alpha)
// 	}
// 	return totalsurv / N
// }

// func FindAlpha(target, low, high, ac float64) float64 {
// 	for high-low > ac {
// 		mid := (low + high) / 2
// 		CheckRTP := CheckExpectedRTP(mid)

// 		if CheckRTP > target {
// 			low = mid
// 		} else {
// 			high = mid
// 		}
// 	}
// 	return (low + high) / 2
// }


// Ранжированное получение альфы взависимости от -rtp 
func GetAlphaRTP(rtp float64) float64 {
    switch {
    case rtp > 0.95:
        return 0.10
    case rtp > 0.90:
        return 0.20
    case rtp > 0.85:
        return 0.30
    case rtp > 0.80:
        return 0.40
    case rtp > 0.75:
        return 0.50
    case rtp > 0.70:
        return 0.60
    case rtp > 0.65:
        return 0.70
    case rtp > 0.60:
        return 0.80
    case rtp > 0.55:
        return 0.90
    case rtp > 0.50:
        return 1.00
    case rtp > 0.40:
        return 1.30
    case rtp > 0.30:
        return 1.70
    case rtp > 0.20:
        return 2.20
    case rtp > 0.10:
        return 3.00
    default:
        return 4.00
    }
}
