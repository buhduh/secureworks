package main
import(
  "fmt"
  //maxminddb "github.com/oschwald/geoip2-golang"
  maxminddb "github.com/oschwald/maxminddb-golang"
  "log"
  "net"
)

asdfasdf


func main() {
  db, err := maxminddb.Open("/root/ips/GeoLite2-City.mmdb")
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()

  ip := net.ParseIP("81.2.69.142")

  var record interface{}
  err = db.Lookup(ip, &record)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%v", record)
}
