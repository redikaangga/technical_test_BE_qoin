package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	ID    int   // ID pemain
	Dice  []int // Dadu yang dimiliki pemain
	Score int   // Skor yang dimiliki pemain
}

// Fungsi untuk melempar dadu
func rollDice(r *rand.Rand, numDice int) []int {
	results := make([]int, numDice)
	for i := 0; i < numDice; i++ {
		// Setiap dadu bernilai acak antara 1 - 6
		results[i] = r.Intn(6) + 1
	}
	return results
}

func main() {
	var N, M, maxRounds int

	// Input jumlah pemain, dadu, dan maksimal giliran(agar looping tidak terlalu banyak)
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu per pemain: ")
	fmt.Scan(&M)
	fmt.Print("Masukkan jumlah giliran maksimal: ")
	fmt.Scan(&maxRounds)

	// generate angka acak berdasarkan waktu saat ini
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	players := make([]Player, N)
	for i := range players {
		// Setiap pemain mendapat ID dan dadu awal
		players[i] = Player{ID: i + 1, Dice: rollDice(r, M)}
	}

	activePlayers := N // Jumlah pemain yang masih aktif (punya dadu)
	round := 1         // Inisialisasi giliran pertama

	// Mulai permainan, berhenti jika hanya tersisa 1 pemain atau maksimal giliran
	for activePlayers > 1 && round <= maxRounds {
		fmt.Printf("Giliran %d lempar dadu:\n", round)
		for i := 0; i < N; i++ {
			// Jika pemain tidak punya dadu lagi, lewati
			if len(players[i].Dice) == 0 {
				continue
			}

			// Tampilkan status pemain
			fmt.Printf("Pemain #%d (%d): %v\n", players[i].ID, players[i].Score, players[i].Dice)
		}

		// Mengolah hasil lemparan dadu
		newDice := make([][]int, N) // Dadu baru setelah evaluasi
		for i := 0; i < N; i++ {
			if len(players[i].Dice) == 0 {
				continue
			}

			// Cek setiap dadu pemain
			for _, die := range players[i].Dice {
				switch die {
				case 6:
					// Jika dadu bernilai 6, pemain mendapat poin
					players[i].Score++
				case 1:
					// Jika dadu bernilai 1, dadu diberikan ke pemain berikutnya
					nextPlayer := (i + 1) % N
					players[nextPlayer].Dice = append(players[nextPlayer].Dice, 1)
				default:
					// Nilai selain 1 dan 6 tetap dipegang oleh pemain
					newDice[i] = append(newDice[i], die)
				}
			}
		}

		// Update dadu pemain dengan hasil cek diatas
		for i := 0; i < N; i++ {
			players[i].Dice = newDice[i]
		}

		// Tampilkan hasil setelah evaluasi
		fmt.Println("Setelah evaluasi:")
		for i := 0; i < N; i++ {
			if len(players[i].Dice) == 0 {
				continue
			}
			fmt.Printf("Pemain #%d (%d): %v\n", players[i].ID, players[i].Score, players[i].Dice)
		}

		// Cek apakah ada pemain yang sudah kehabisan dadu
		for i := 0; i < N; i++ {
			if len(players[i].Dice) == 0 {
				activePlayers-- // Kurangi jumlah pemain aktif
			}
		}

		round++ // Lanjut ke ronde berikutnya
	}

	// Menentukan pemenang berdasarkan skor tertinggi
	var winners []Player
	maxScore := 0

	for _, player := range players {
		if player.Score > maxScore {
			// Jika ada skor lebih tinggi, reset daftar pemenang
			maxScore = player.Score
			winners = []Player{player}
		} else if player.Score == maxScore {
			// Jika skor sama, tambahkan ke daftar pemenang
			winners = append(winners, player)
		}
	}

	// Notifikasi hasil
	if len(winners) == N {
		fmt.Println("Semua SERI") // Jika semua pemain punya skor yang sama
	} else if len(winners) > 1 {
		// Jika lebih dari satu pemenang, tampilkan mereka sebagai seri
		names := ""
		for i, winner := range winners {
			if i > 0 {
				names += " dan "
			}
			names += fmt.Sprintf("Pemain #%d", winner.ID)
		}
		fmt.Printf("%s SERI\n", names)
	} else if len(winners) == 1 {
		// Jika hanya satu pemenang, tampilkan pemenangnya
		fmt.Printf("Pemain #%d MENANG dengan %d poin!\n", winners[0].ID, winners[0].Score)
	}
}
