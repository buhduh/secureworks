package main
import(
  "fmt"
  //maxminddb "github.com/oschwald/geoip2-golang"
  maxminddb "github.com/oschwald/maxminddb-golang"
  "log"
  "net"
)

/*
map[
  country:map[geoname_id:2635167 is_in_european_union:true iso_code:GB names:map[pt-BR:Reino Unido ru:Великобритания zh-CN:英国 de:Vereinigtes Königreich en:United Kingdom es:Reino Unido fr:Royaume-Uni ja:イギリス]] 
  location:map[
    accuracy_radius:100 
    latitude:52.1398 
    longitude:-0.2399 
    time_zone:Europe/London] postal:map[code:SG19] registered_country:map[iso_code:GB names:map[ja:イギリス pt-BR:Reino Unido ru:Великобритания zh-CN:英国 de:Vereinigtes Königreich en:United Kingdom es:Reino Unido fr:Royaume-Uni] geoname_id:2635167 is_in_european_union:true] subdivisions:[map[geoname_id:6269131 iso_code:ENG names:map[fr:Angleterre ja:イングランド pt-BR:Inglaterra ru:Англия zh-CN:英格兰 de:England en:England es:Inglaterra]] map[geoname_id:2653940 iso_code:CAM names:map[en:Cambridgeshire ja:ケンブリッジ ru:Кембридж]]] city:map[geoname_id:2648899 names:map[en:Gamlingay]] continent:map[code:EU geoname_id:6255148 names:map[en:Europe es:Europa fr:Europe ja:ヨーロッパ pt-BR:Europa ru:Европа zh-CN:欧洲 de:Europa]]]

*/


func main() {
  db, err := maxminddb.Open("../ips/GeoLite2-City_20200204/GeoLite2-City.mmdb")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  ip := net.ParseIP("81.2.69.142")
  var record struct {
      Location struct {
          Latitude float32 `maxminddb:"latitude"`
          Longitude float32 `maxminddb:"longitude"`
          Radius   int `maxminddb:"accuracy_radius"`
      } `maxminddb:"location"`
  }
  err = db.Lookup(ip, &record)
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("latitude: %f\nlogitude: %f\nradius: %d\n", record.Location.Latitude, record.Location.Longitude, record.Location.Radius)
}
