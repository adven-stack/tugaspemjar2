# README

## Anggota Kelompok  
- Nama: Advenat Desta Yolantino 
  NIM: 223400017
- Nama: Fulgensius Olivier Dorman
  NIM: 223400018

## Deskripsi Singkat Aplikasi  
Aplikasi ini dibuat untuk membuat simulasi jaringan komputer (TCP chat dan UDP notify) seperti yang terdapat dalam folder `tcp_chat` dan `udp_notify` pada repository. 
Aplikasi memungkinkan komunikasi real-time antar node jaringan dengan protokol berbeda.

## Cara Menjalankan Aplikasi  
### Setup  
1. Pastikan Go bahasa pemrograman sudah terinstall di sistem Anda (versi minimal Go 1.x).  
2. Clone repository ke lokal:  
   ```bash  
   git clone https://github.com/adven-stack/tugaspemjar2.git  
   cd tugaspemjar2  
3. Masuk ke folder modul yang ingin dijalankan (tcp_chat atau udp_notify):
   cd tcp_chat  atau cd udp_notify  
Run

1. Jalankan modul TCP chat: go run main.go  
2. Untuk modul UDP notify, lakukan hal yang sama di folder udp_notify: cd ../udp_notify  
go run main.go  

Cuplikan Tampilan / Contoh Interaksi

Tambahkan screenshot ke folder docs/ atau images/ lalu referensikan di sini.

Contoh interaksi di terminal:
Client: Hallo server, ini client A.
Server: Halo client A, pesan diterima.

Node A (UDP) : Broadcast notifikasi ke Node B.
Node B (UDP) : Terima notifikasi dari Node A.
