# ICE SPICE - API to fetch Baltic Sea Ice data
![Ice-spice](public/logo.png?raw=true)

SIMPLE REST API to fetch Baltic Sea Ice Data recorded by different metrological stations in Finland

## Table of Contents

- [API's Schema](#apis-schema)
- [Generated Station Code](#generated-station-code)
- [Usage and Examples](#usage-and-examples)
- [Description of Data](#description-of-data)
- [Data Source](#data-source)
- [Technologies Used](#technologies-used)
- [Contribution](#contribution)
- [Contact](#contact)

# API's Schema

Currently the application supports below end points and no auth is required for them as of now.

1. Endpoint 1

- Request URL - https://ice-spice.onrender.com/api/station_data?stationCode=STATION_CODE
- Request Method - `GET`
- Request Query Param - `stationCode`
- Query Param Value - value for `stationCode` query param can be obtained from the below table where station code for each metrological station is given.

It gives all the ice data and the analysis calculated on it by that metrological station whose station code is added in the params. It can be used to fetch all the data recorded by particular metrological station over the past years.

2. Endpoint 2 

- Request URL - https://ice-spice.onrender.com/api/analysis_data?stationCode=STATION_CODE
- Request Method - `GET`
- Request Query Param - `stationCode`
- Query Param Value - value for `stationCode` query param can be obtained from the below table where station code for each metrological station is given.

It can be used to fetch only the analysis data obtained by particular metrological station over the past years of the ice data recorded by it.

# Generated Station Code
Station code generated for different metrological stations of Finland is given below.

STATION NAME | GENERATED STATION CODE
--- | ---
Röyttä – Etukari  |  STROET
Etukari – Ristinmatala  |  STETRI
Ajos – Ristinmatala  |  STAJRI
Ristinmatala – Kemi 2  |  STRIKE2
Kemi 2 – Kemi 1  |  STKE2KE1
Kemi 1, Seegebiet im SW  |  STKE1SEIMSW
Kemi 2 – Ulkokrunni – Virpiniemi  |  STKE2ULVI
Oulu, Hafen – Kattilankalla  |  STOUHAKA
Kattilankalla – Oulu 1  |  STKAOU1
Oulu 1, Seegebiet im SW  |  STOU1SEIMSW
Offene See N-lich Breite Marjaniemi  |  STOFSEN-BRMA
Raahe, Hafen – Heikinkari  |  STRAHAHE
Heikinkari – Raahe Leuchtturm  |  STHERALE
Raahe Leuchtturm – Nahkiainen  |  STRALENA
Breitengrad Marjaniemi – Ulkokalla, See  |  STBRMAULSE
Rahja, Hafen – Välimatala  |  STRAHAVA
Välimatala bis Linie Ulkokalla – Ykskivi  |  STVABILIULYK
Breitengrad Ulkokalla – Pietarsaari, See  |  STBRULPISE
Ykspihlaja – Repskär  |  STYKRE
Repskär – Kokkola Leuchtturm  |  STREKOLE
Kokkola Leuchtturm, See außerhalb  |  STKOLESEAU
Pietarsaari – Kallan  |  STPIKA
Kallan, Seegebiet außerhalb  |  STKASEAU
Breite Pietarsaari – Nordvalen im NE  |  STBRPINOIMNE
Nordvalen, Seegebiet im ENE  |  STNOSEIMEN
Nordvalen – Norrskär, See im W  |  STNONOSEIMW
Vaskiluoto – Ensten  |  STVAEN
Ensten – Vaasa Leuchtturm  |  STENVALE
Vaasa Leuchtturm – Norrskär  |  STVALENO
Norrskär, Seegebiet im SW  |  STNOSEIMSW
Kaskinen – Sälgrund  |  STKASA
Sälgrund, Seegebiet außerhalb  |  STSASEAU
Offene See N-lich Breite Yttergrund  |  STOFSEN-BRYT
Pori – Linie Pori Leuchtturm – Säppi  |  STPOLIPOLESA
Linie Pori Lt. – Säppi – See im W  |  STLIPOLTSASEIMW
Hohe See Länge Yttergrund u. Rauma  |  STHOSELAYTU.RA
Rauma, Hafen – Kylmäpihlaja  |  STRAHAKY
Kylmäpihlaja – Rauma Leuchtturm  |  STKYRALE
Rauma Leuchtturm, See im W  |  STRALESEIMW
Breitengrad Rauma, offene See im S  |  STBRRAOFSEIMS
Uusikaupunki, Hafen – Kirsta  |  STUUHAKI
Kirsta – Isokari  |  STKIIS
Isokari – Sandbäck  |  STISSA
Sandbäck, Seegebiet außerhalb  |  STSASEAU2
Sälskär, See im N  |  STSASEIMN
Märket, See im N  |  STMASEIMN
Märket, See im W  |  STMASEIMW
Märket, See im S  |  STMASEIMS
Maarianhamina – Marhällan  |  STMAMA
See außerhalb Nyhamn u. Marhällan  |  STSEAUNYUMA
Ålandsee, mittlerer Teil  |  STALMITE
Lågskär, See im S  |  STLASEIMS
Naantali und Turku – Rajakari  |  STNAUNTURA
Rajakari – Lövskär  |  STRALO
Lövskär – Korra  |  STLOKO
Korra – Isokari  |  STKOIS
Lövskär – Berghamn  |  STLOBE
Berghamn – Stora Sottunga  |  STBESTSO
Stora Sottunga – Ledskär  |  STSTSOLE
Rödhamn, Seegebiet  |  STROSE
Lövskär – Grisselborg  |  STLOGR
Grisselborg – Norparskär  |  STGRNO
Vidskär, Seegebiet  |  STVISE
Utö – Suomen Leijona  |  STUTSULE
Suomen Leijona, See im S  |  STSULESEIMS
Hanko, Hafen – Hanko 1  |  STHAHAHA1
Hanko 1, See im S  |  STHA1SEIMS
Hanko – Vitgrund  |  STHAVI
Vitgrund – Utö  |  STVIUT
Koverhar – Hästö Busö  |  STKOHÄBU
Hästö Busö – Ajax  |  STHABUAJ
Ajax, See im S  |  STAJSEIMS
Inkoo u. Kantvik – Porkkala See  |  STINUKAPOSE
Porkkala, Seegebiet  |  STPOSE
Porkkala Leuchtturm, See im S  |  STPOLESEIMS
Helsinki, Hafen – Harmaja  |  STHEHAHA
Harmaja – Helsinki Leuchtturm  |  STHAHELE
Helsinki Lt. – Porkkala Lt., See im S  |  STHELTPOLTSEIMS
Helsinki – Porkkala – Rönnskär, Fahrw.  |  STHEPOROFA
Vuosaari Hafen – Eestiluoto  |  STVUHAEE
Eestiluoto – Helsinki Leuchtturm  |  STEEHELE
Porvoo, Hafen – Varlax  |  STPOHAVA
Varlax – Porvoo Leuchtturm  |  STVAPOLE
Porvoo Leuchtturm – Kalbådagrund  |  STPOLEKA
Kalbådagrund – Helsinki Lt.  |  STKAHELT
Valko, Hafen – Täktarn  |  STVAHATA
Boistö – Glosholm, Schärenfahrwasser  |  STBOGLSC
Glosholm–Helsinki, Schärenfahrwasser  |  STGLSC
Kotka – Viikari  |  STKOVI
Viikari – Orrengrund  |  STVIOR
Orrengrund – Tiiskeri  |  STORTI
Tiiskeri – Kalbådagrund  |  STTIKA
Hamina – Suurmusta  |  STHASU
Suurmusta – Merikari  |  STSUME
Merikari – Kaunissaari  |  STMEKA

# Usage and Examples
Below are the example usage of API's

1. Endpoint 1
- Request URL - https://ice-spice.onrender.com/api/station_data?stationCode=STTIKA
- Request Method - `GET`

Output
```JSON
{
    "data": {
        "analysis-data-station": [
            {
                "analysis-basis": "1.quartile",
                "beginning": "29.Dec",
                "end": "19.Feb",
                "ice_days": 9,
                "fl_total": 4.6,
                "vol_sum": 20.8,
                "country": "FINLAND",
                "country_code": "FI",
                "station_name": "Station: Tiiskeri – Kalbådagrund",
                "station_code": "STTIKA"
            },
            ...
        ],
        "stations-data": [
            {
                "beginning": "29.Dec.2002",
                "end": "09.May.2003",
                "season_length": 132,
                "ice_days": 126,
                "fl_total": 114.6,
                "vol_sum": 3986.5,
                "max_thickness2": 50,
                "country": "FINLAND",
                "country_code": "FI",
                "station_name": "Station: Tiiskeri – Kalbådagrund",
                "station_code": "STTIKA",
                "restricted_days": [
                    126,
                    126,
                    126,
                    126,
                    126,
                    108,
                    0,
                    0,
                    0
                ]
            }
            ...
        ]
    },
    "message": "metro station data fetched successfully"
}
```

2. Endpoint 2
- Request URL - https://ice-spice.onrender.com/api/analysis_data?stationCode=STTIKA
- Request Method - `GET`

Output
```JSON
{
    "data": {
        "analysis-data-station": [
            {
                "analysis-basis": "1.quartile",
                "beginning": "29.Dec",
                "end": "19.Feb",
                "ice_days": 9,
                "fl_total": 4.6,
                "vol_sum": 20.8,
                "country": "FINLAND",
                "country_code": "FI",
                "station_name": "Station: Tiiskeri – Kalbådagrund",
                "station_code": "STTIKA"
            },
            ...
        ]
    },
    "message": "metro station analysis data fetched successfully"
}
```

# Description of Data
- The yearly statistics of a single station are contained in the response of the first API. 
- In each data object in `stations-data` array in response. It contains the `beginning` and the `end` of the ice season at that station (so first and last day with ice).
- `season_length` is the length of the season, just the days between first and last ice occurence.
- `ice_days` gives the number of days with ice, the maximum possible value is the length of the season, but mostly there are less days with ice, as the season mostly is not continuos. 
- `fl_total` is the area sum, that is the sum of all concentrations. 
- `vol_sum` is the volume sum, the total sum of concentration times thickness.
- `max_thickness` is the maximum ice thichness from measurements.
- `max_thickness2` is the maximum ice thichness from coded ice reports according to the Baltic sea ice code.
- `restricted_days` give the days with restrictions acc to the baltic sea ice code. Starting with minimum restriction=1 up, then minimum=2, etc.
- `analysis-data-station` gives the summary of all the data of a station. Although the statistic is based on checked data, errors still can occur. And although many stations are continuosly, not all start the same year and some are discontinued. Also some stations were not measured in time in between (for example due to illness or death of the ice observer).
- For more information visit [here.](https://www.bsis-ice.de/statistik/StatistikInfo.html)

# Data Source
The data is scraped from website of [Baltic Sea Ice Services](https://www.bsis-ice.de/technical_informations.shtml). They have listed baltic sea ice data for various countries on their website. It can be viewed [here.](https://www.bsis-ice.de/statistik/Stationindex.html) I've scraped data for Finland and stored it in PostgreSQL database for querying through API.

# Technologies Used
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)

# Contribution
Feel free to raise issues and pull requests for contribution.

# Contact
[![srrathi github](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/srrathi)
[![Sitaram Rathi Linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/sitaram-rathi-519152197/)
[![Sitaram Rathi email](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:srrathi2000@gmail.com)
[![Sitaram Rathi twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/SitaramRathi5)
