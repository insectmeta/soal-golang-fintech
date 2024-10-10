package main

import (
	"fmt"
)

/*
Assignment C: Fintech
1. Fitur Exchange dan Kurs.
User mau menukarkan uang dari mata uang IDR & SGD ke USD.
Diberikan:
- Jumlah uang yg ingin ditukar (exchange). IDR & SGD
- Kurs USD ke IDR
- Kurs USD ke SGD
Hitunglah jumlah USD yg dimiliki user setelah ditukar!

Contoh:
jumlah IDR: 120,000 (seratus dua puluh ribu)
jumlah SGD: 29
Kurs_USD_IDR: 15,000 (1 USD = 15000 IDR)
Kurs_USD_SGD: 1.5 (1 USD = 1.5 SGD)
Jawaban: 27.33 USD
Penjelasan:
IDR to USD = (120,000 / 15,000) = 8
SGD to USD = (29 / 1.5) = 19.333333
total = 8 + 19.333333 = 27.333333 USD
*/

// exchangeToUSD, output total in USD
func exchangeToUSD(amountIDR float64, exchangeRateUSD_IDR float64, amountSGD float64, exchangeRateUSD_SGD float64) float64 {
	// write code here
	if amountIDR > 0 && amountSGD > 0 && exchangeRateUSD_IDR > 0 && exchangeRateUSD_SGD > 0 {
		return (amountIDR / exchangeRateUSD_IDR) + (amountSGD / exchangeRateUSD_SGD)
	}
	return 0
}

/*
2. Laporan transaksi terbesar & terkecil.
Setiap bulan Fintech ini mau memberi laporan pengeluaran bulanan kepada user2nya via email.
Di laporan tersebut ada rata-rata pengeluaran, transaksi terbesar & transaksi terkecil selama bulan ini.
Diberikan list pengeluaran user selama 1 bulan, carilah transaksi terkecil dan terbesar!
Contoh: [50000, 30000, 80000, 45000, 72000, 58000, 65000]

Jawaban: 57142.857143 30000 80000
Penjelasan:
Rata-rata = sum(expenseList) / 7 = 400000 / 7 = 57142,857143
Minimum = 30000 (ada di index 1)
Maksimum = 80000 (ada di index 2)
*/

// minAndMaxTransaction, output 3 values, mean, min & max transaction
func minAndMaxTransaction(expenseList []int) (avg float64, min int, max int) {
	// write code here
	if len(expenseList) == 0 {
		return 0, 0, 0
	}

	min, max = expenseList[0], expenseList[0]
	total := 0

	for _, expense := range expenseList {
		if expense < min {
			min = expense
		}
		if expense > max {
			max = expense
		}
		total += expense
	}

	avg = float64(total) / float64(len(expenseList))
	return avg, min, max
	// return 50, 1, 99 // contoh jika rata2 = 50, minimum = 1, dan maksimum = 99
}

func main() {

	fmt.Println(exchangeToUSD(120000, 15000, 29, 1.5)) // contoh soal. output 27.333333
	fmt.Println(exchangeToUSD(120000, 16000, 29, 1.5)) // output
	fmt.Println(exchangeToUSD(120000, 15000, 25, 1.5)) // output

	expenseList1 := []int{50000, 30000, 80000, 45000, 72000, 58000, 65000}
	fmt.Println(minAndMaxTransaction(expenseList1)) // contoh soal. output 57142,857143 30000 80000

	expenseList2 := []int{80000, 60000, 30000, 45000, 72000, 58000, 20000, 15000}
	fmt.Println(minAndMaxTransaction(expenseList2)) // output 3

	expenseList3 := []int{50000, 60000, 10000, 75000, 10000, 75000, 65000, 45000, 11000}
	fmt.Println(minAndMaxTransaction(expenseList3)) // output 6
}
