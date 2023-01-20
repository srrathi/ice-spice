const analysisData = require("./data/analysis-config.json")
const res = []

analysisData.forEach(station=>{
    station.forEach(item=>{
        var str = `('${item["Analysis-basis"]}', '${item["Beginning-ice"]}', '${item["End-ice"]}', ${parseFloat(item["Ice-days"])}, ${parseFloat(item["Fl-total"])}, ${parseFloat(item["Vol-sum"])}, '${item.Country}', '${item["Country-code"]}', '${item["Station-name"]}', '${item["Station-code"]}')`
        res.push(str)
    })
})

const result = res.join(", ")
console.log(result)

const stationsData = require("./data/stations-config.json")
const data = []
let count = 1

stationsData.forEach(station => {
    station.forEach((item, index) => {
        var str = `('${item.Beginning}', '${item.End}', ${parseInt(item["Season-length"])}, ${parseFloat(item["Ice-days"])}, ${parseFloat(item["Fl-total"])}, ${parseFloat(item["Vol-sum"])}, '${item.Country}', ${item["Max-thickness"] === "NaN" ? null : item["Max-thickness"]}, ${parseInt(item["Max-thickness2"])}, '${item["Country-code"]}', '${item["Station-name"]}', '${item["Station-code"]}', ${count})`
        count++
        data.push(str)
    })
})

const response = data.join(", ")
console.log(response)

