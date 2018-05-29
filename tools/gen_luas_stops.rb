
require 'json'
require 'net/http'
require 'uri'

# WARNING
# Hacky script, temporary

STOP_INFO_URL = "https://data.smartdublin.ie/cgi-bin/rtpi/busstopinformation?stopid&format=json"
TARGET_GO_FILE = File.join(__dir__, '..', 'luas_stops_def.go')

def get_luas_stops(str)
  json = JSON.parse(str)
  stop_structs = []
  json["results"].select do |stop|
    stop['operators'] && stop['operators'].first['name'] == 'LUAS'
  end.each do |luas_stop|
    stop_name = /LUAS\s+(?<stop_name>.*)/.match(luas_stop['displaystopid'])[:stop_name]
    short_stop_name = nil
    short_stop_name = case stop_name
                      when "St. Stephen's Green"
                        "STS"
                      when "Central Park"
                        "CPK"
                      when "The Gallops"
                        "GAL"
                      when "Ballyogan Wood"
                        "GAW"
                      when "Carrickmines"
                        "CCK"
                      when "Cheeverstown"
                        "CVN"
                      when "George's Dock"
                        "GDK"
                      when "Mayor Square - NCI"
                        "MYS"
                      when "Spencer Dock"
                        "SDK"
                      when "The Point"
                        "TPT"
                      when "Trinity"
                        "TRI"
                      when "O'Connell - Upper"
                        "OUP"
                      when "O'Connell - GPO"
                        "OGP"
                      when "Broadstone - DIT"
                        "BRD"
                      end
    new_stop = {
      name: stop_name,
      name_abv: short_stop_name || stop_name[0..2].upcase,# luas_stop['shortname'] is always an empty string in this API (sigh)
      line: luas_stop['operators'].first['routes'].first.downcase,
      coordinates: [luas_stop['latitude'].to_f, luas_stop['longitude'].to_f]
    }
    stop_structs << <<~EOG
    &Stop{
      Name: "#{stop_name}",
      NameAbv: "#{short_stop_name || stop_name[0..2].upcase}",
      Line: "#{luas_stop['operators'].first['routes'].first.downcase}",
      Coordinates: []float64{#{luas_stop['latitude'].to_f}, #{luas_stop['longitude'].to_f}},
    },
    EOG
  end
  all_stop_structs = stop_structs.join
  return <<~EOG
  package luas

  type Stop struct {
    Name        string    `json:"name"`
    NameAbv     string    `json:"name_abv"`
    Line        string    `json:"line"`
    Coordinates []float64 `json:"coordinates"`
  }

  var allStops = []*Stop{
    #{all_stop_structs}
  }
  EOG
end

def http_get(url)
  uri = URI(url)
  Net::HTTP.get(uri)
end

luas_stops = get_luas_stops(http_get(STOP_INFO_URL))
File.write(TARGET_GO_FILE, luas_stops)
`go fmt #{TARGET_GO_FILE}`
