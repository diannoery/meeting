package main

// Buatlah sebuah aplikasi personal digital meeting, dengan kententuan
// 1. bisa mendaftarkan rencana meeting,dengan informasi tanggal,jam, tempat, client
// 2. Dapat melihat jadual meeting untuk hari ini, atau besok (fleksibel berdasarkan jumlah hari yang diinput)
// 3. Dapat membatalkan meeting, dengan memasukan informasi alasannya

import (
	"meeting/config"
	"meeting/main/master"
)

func main() {
	db := config.Connection()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
