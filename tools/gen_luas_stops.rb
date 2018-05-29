
require 'json'
require 'net/http'
require 'uri'

# WARNING
# Hacky script, temporary

STOP_INFO_URL = "https://data.smartdublin.ie/cgi-bin/rtpi/busstopinformation?stopid&format=json"
TARGET_FILE = File.join(__dir__, '..', 'configs', 'luas_stops.json')

def get_luas_stops(str)
  json = JSON.parse(str)
  luas_stops = json["results"].select do |stop|
    stop['operators'] && stop['operators'].first['name'] == 'LUAS'
  end.map do |luas_stop|
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
    new_stop
  end
  luas_stops
end

def http_get(url)
  uri = URI(url)
  Net::HTTP.get(uri)
end

luas_stops = get_luas_stops(http_get(STOP_INFO_URL))
File.write(TARGET_FILE, JSON.pretty_generate(luas_stops))
