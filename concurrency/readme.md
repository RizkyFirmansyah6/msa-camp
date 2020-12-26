#### Problem 4
##### Concurrency Task Worker
pada program ini terdapat 2 flag yaitu concurren limit dan output path

program ini terbagi menjadi 3 fungsi utama yaitu
1.Fungsi getJson

fungsi ini digunakan untuk mengambil json dari API dan meletakkanya pada struct

package yang digunakan yaitu net/http untuk mengambil data dari API, bytes untuk 
merubah response body menjadi bytes dan encoding/json untuk unmarshal json ke struct

2.Fungsi grouping

fungsi ini digunakan untuk mengelempokkan struct bedasarkan kabupaten kota

fungsi ini menggunakan map untuk pengelompokanya

3.Fungsi writeToCSV

fungsi ini digunakan untuk merubah struct menjadi csv dan membuat file sesuai dengan kabupaten kota

package yang digunakan yaitu github.com/mohae/struct2csv untuk merubah struct menjadi csv 
dan os untuk membuat file dengan lokasi output yang diinputkan 