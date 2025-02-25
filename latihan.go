package main 
import "fmt"

func main(){
    var nama, nama2 string
    var skor, i, a, b, j, jumlahSkorA, jumlahSkorB int
    
    fmt.Scan(&nama, &skor)
    i = 1
    for i <= 8{
        if skor < 301{
            a++
            jumlahSkorA = jumlahSkorA + skor 
        }
        fmt.Scan(&skor)
        i++
    }
    fmt.Scan(&nama2, &skor)
    j = 1
    for j <= 8{
        if skor < 301{
            b++
            jumlahSkorB = jumlahSkorB + skor
        }
        fmt.Scan(&skor)
        j++
    }
    if a > b {
        fmt.Print(nama," ", a, " ", jumlahSkorA)
    }else if a < b{
        fmt.Print(nama2," ", b, " ", jumlahSkorB)
    }else{
        fmt.Print("Seri")
    }
}