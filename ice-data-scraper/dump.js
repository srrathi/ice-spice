const { Pool } = require('pg');
require("dotenv").config();
const analysisData = require("./data/analysis-config.json")

// postgresql://postgres:2hvRX47659hziAaJY5C2@containers-us-west-167.railway.app:6663/railway
const pool = new Pool({
    user: process.env.DB_USER,
    database: process.env.DB_DATABASE,
    password: process.env.DB_PASSWORD,
    port: process.env.PORT,
    host: process.env.DB_HOST,
})

console.log(process.env.DB_USER)

let counter = 0
async function insertData() {
    for (let i = 0; i < analysisData.length; i++) {
        for (let j = 0; j < analysisData[i].length; j++) {
            try {
                console.log("INSERTING ", analysisData[i][j]['Analysis-basis'],
                    analysisData[i][j]['Beginning-ice'],
                    analysisData[i][j]['End-ice'])
                const res = await pool.query(
                    "INSERT INTO analysis_data (analysis_basis, beginning, end, ice_days, fl_total, vol_sum, country, country_code, station_name, station_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
                    [analysisData[i][j]['Analysis-basis'],
                    analysisData[i][j]['Beginning-ice'],
                    analysisData[i][j]['End-ice'],
                    parseFloat(analysisData[i][j]['Ice-days']),
                    parseFloat(analysisData[i][j]['Fl-total']),
                    parseFloat(analysisData[i][j]['Vol-sum']),
                    analysisData[i][j]['Country'],
                    analysisData[i][j]['Country-code'],
                    analysisData[i][j]['Station-name'],
                    analysisData[i][j]['Station-code']
                    ]
                );
                if (res) counter++;
            } catch (error) {
                console.log(error)
            }
        }
    }
    // await analysisData.forEach(async (station) => {
    //     await station.forEach(async (row) => {
    //         try {
    //             const res = await pool.query(
    //                 "INSERT INTO analysis_data (analysis_basis, beginning, end, ice_days, fl_total, vol_sum, country, country_code, station_name, station_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
    //                 [row['Analysis-basis'],
    //                 row['Beginning-ice'],
    //                 row['End-ice'],
    //                 parseFloat(row['Ice-days']),
    //                 parseFloat(row['Fl-total']),
    //                 parseFloat(row['Vol-sum']),
    //                 row['Country'],
    //                 row['Country-code'],
    //                 row['Station-name'],
    //                 row['Station-code']
    //                 ]
    //             );
    //             if (res) counter++;
    //         } catch (error) {
    //             console.log(error)
    //         }
    //     })
    // })

    console.log(`SUCCESSFULLY INSERTED ${counter} ENTRIES`)

}

insertData();