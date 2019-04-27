//package tempconv make celsius to fahrenheit or fahtenheit to celsius
package tempconv

//Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//Fahrenheit ti Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
