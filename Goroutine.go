package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine dihentikan")
			return
		default:
			fmt.Println("Goroutine berjalan...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Membuat context dengan fungsi cancel
	ctx, cancel := context.WithCancel(context.Background())

	// Menjalankan worker dalam Goroutine
	go worker(ctx)

	// Biarkan Goroutine berjalan selama 3 detik
	time.Sleep(3 * time.Second)

	// Mengirim sinyal berhenti ke Goroutine dengan memanggil cancel()
	cancel()

	// Menunggu 1 detik agar output "Goroutine dihentikan" dapat terlihat sebelum program selesai
	time.Sleep(1 * time.Second)
}
