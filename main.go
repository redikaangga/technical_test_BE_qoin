package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struktur Player merepresentasikan pemain dalam permainan dadu
type Player struct {
	ID    int   // ID pemain
	Dice  []int // Daftar dadu yang dimiliki oleh pemain
	Score int   // Skor pemain berdasarkan jumlah dadu bernilai 6
}

// Fungsi rollDice mengembalikan slice berisi hasil lemparan sejumlah dadu
func rollDice(r *rand.Rand, numDice int) []int {
	results := make([]int, numDice) // Buat slice untuk menyimpan hasil lemparan
	for i := 0; i < numDice; i++ {
		// Hasil lemparan antara 1 hingga 6
		results[i] = r.Intn(6) + 1
	}
	return results
}

// Fungsi calculateWinner menentukan siapa pemain dengan skor tertinggi
func calculateWinner(players []Player) ([]Player, int) {
	var winners []Player
	maxScore := 0

	// Cari skor tertinggi dan daftar pemenang
	for _, player := range players {
		if player.Score > maxScore {
			maxScore = player.Score
			winners = []Player{player} // Reset daftar pemenang
		} else if player.Score == maxScore {
			winners = append(winners, player) // Tambahkan ke daftar pemenang
		}
	}

	return winners, maxScore
}

// Fungsi finalResult menampilkan hasil akhir dari permainan
func finalResult(winners []Player, maxScore int, totalPlayers int) {
	// Notifikasi hasil
	if len(winners) == totalPlayers {
		fmt.Println("Semua pemain SERI!")
	} else if len(winners) > 1 {
		names := ""
		for i, winner := range winners {
			if i > 0 {
				names += " dan "
			}
			names += fmt.Sprintf("Pemain #%d", winner.ID)
		}
		fmt.Printf("%s SERI!\n", names)
	} else if len(winners) == 1 {
		fmt.Printf("Pemain #%d MENANG dengan %d poin!\n", winners[0].ID, winners[0].Score)
	}
}

func main() {
	var N, M, maxRounds int

	// Ambil input dari pengguna untuk jumlah pemain, jumlah dadu, dan jumlah maksimal giliran
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu per pemain: ")
	fmt.Scan(&M)
	fmt.Print("Masukkan jumlah giliran maksimal: ")
	fmt.Scan(&maxRounds)

	// Inisialisasi generator angka acak dengan seed dari waktu saat ini
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Inisialisasi pemain-pemain dengan dadu mereka
	players := make([]Player, N)
	for i := range players {
		players[i] = Player{ID: i + 1, Dice: rollDice(r, M)}
	}

	activePlayers := N // Jumlah pemain yang masih memiliki dadu
	round := 1         // Awal dari giliran permainan

	// Permainan berlangsung selama masih ada lebih dari 1 pemain aktif atau sampai gilirannya habis
	for activePlayers > 1 && round <= maxRounds {
		fmt.Printf("Giliran %d lempar dadu:\n", round)
		for i := 0; i < N; i++ {
			// Hanya tampilkan pemain yang masih punya dadu
			if len(players[i].Dice) == 0 {
				continue
			}

			// Tampilkan status pemain
			fmt.Printf("Pemain #%d (Skor: %d): %v\n", players[i].ID, players[i].Score, players[i].Dice)
		}

		// Slice baru untuk menyimpan dadu yang tersisa setelah evaluasi
		newDice := make([][]int, N)
		for i := 0; i < N; i++ {
			// Lewati pemain yang tidak punya dadu
			if len(players[i].Dice) == 0 {
				continue
			}

			// Evaluasi setiap nilai dadu
			for _, die := range players[i].Dice {
				switch die {
				case 6:
					// Tambah skor jika dadu bernilai 6
					players[i].Score++
				case 1:
					// Jika dadu bernilai 1, operkan dadu tersebut ke pemain berikutnya
					nextPlayer := (i + 1) % N
					players[nextPlayer].Dice = append(players[nextPlayer].Dice, 1)
				default:
					// Simpan dadu lainnya
					newDice[i] = append(newDice[i], die)
				}
			}
		}

		// Update dadu yang tersisa untuk tiap pemain
		for i := 0; i < N; i++ {
			players[i].Dice = newDice[i]
		}

		fmt.Println("Setelah evaluasi:")
		for i := 0; i < N; i++ {
			if len(players[i].Dice) == 0 {
				continue
			}
			fmt.Printf("Pemain #%d (Skor: %d): %v\n", players[i].ID, players[i].Score, players[i].Dice)
		}

		// Hapus pemain yang tidak punya dadu
		for i := 0; i < N; i++ {
			if len(players[i].Dice) == 0 {
				activePlayers--
			}
		}

		round++
	}

	// Panggil fungsi untuk menentukan pemenang
	winners, maxScore := calculateWinner(players)

	// Panggil fungsi untuk menampilkan hasil akhir
	finalResult(winners, maxScore, N)
}
