package main

import "fmt"

// -----------------------------ini fungsi celcius jar------------------
func celcius(temperature float64) (float64, float64) {
	kelvin := temperature + 273.15
	fahrenheit := (temperature * 9 / 5) + 32
	return kelvin, fahrenheit
}

// -----------------------------kalo ini fungsi kelvin------------------
func kelvin(temperature float64) (float64, float64) {
	celcius := (temperature - 273.15)
	fahrenheit := (temperature-273.15)*9/5 + 32
	return celcius, fahrenheit
}

// -----------------------------kalo ini fungsi Fahrenheit------------------
func fahrenheit(temperature float64) (float64, float64) {
	celcius := (temperature - 32) * 5 / 9
	kelvin := (temperature-32)*5/9 + 273.15
	return celcius, kelvin
}

func main() {
	for {
		var opr string
		var temp float64
		fmt.Println("pilih opsi")
		fmt.Println("1. Celcius")
		fmt.Println("2. Kelvin")
		fmt.Println("3. Fahrenheit")

		fmt.Println("0. keluar program")

		fmt.Print("silahkan masukkan opsi operator : ")
		fmt.Scan(&opr)

		if opr == "0" {
			break
		}

		switch opr {
		// -----------------------------ini make fungsi celcius------------------
		case "1":
			fmt.Print("Masukkan temperatur suhu Celcius : ")
			fmt.Scan(&temp)

			k, f := celcius(temp)

			fmt.Printf("Kelvin = %.2f\n", k)
			fmt.Printf("Fahrenheit = %.2f\n", f)

		// -----------------------------ini make fungsi kelvin------------------
		case "2":
			fmt.Print("Masukkan temperatur suhu Kelvin : ")
			fmt.Scan(&temp)

			c, f := kelvin(temp)

			fmt.Printf("Celcius = %.2f\n", c)
			fmt.Printf("Fahrenheit = %.2f\n", f)

		//  -----------------------------ini make fungsi fahrenheit------------------
		case "3":
			fmt.Print("Masukkan temperatur suhu kelvin : ")
			fmt.Scan(&temp)

			c, k := fahrenheit(temp)

			fmt.Printf("celcius = %.2f\n", c)
			fmt.Printf("kelvin = %.2f\n", k)

		// -----------------------------ERROR HANDLING NIH LU URUS------------------
		default:

		}
		// ini buat anuan biar rapih aja
		fmt.Println("tekan enter untuk lanjut...")
		fmt.Scanln()

	}

}

// NOTE!
// jar lu urus error handling, sama kalo bisa rapihin outputnya di terminal biar lebih enak diliat
