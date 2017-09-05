package tempconv

import "fmt"

type Celsius float64
type Kelvin float64
type Fahrenheit float64
type Inch float64
type Meter float64

const (
	AbsoluteZeroC Celsius = - 273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (m Meter) String() string {
	return fmt.Sprint("%g m", m)
}

func (i Inch) String() string {
	return fmt.Sprint("%g In", i)
}

func InToMeter(i Inch) Meter {
	return 0.0254 * Meter(i)
}

func MeterToInch(m Meter) Inch {
	return 39.3700787 * Inch(m)
}

func (c Celsius) String() string {
	return fmt.Sprint("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprint("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprint("%gK", k)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
