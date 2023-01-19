const puppeteer = require("puppeteer");
const { writeFile } = require('fs');

const path = './data/config.json';
const stationLinks = require("./data/stations.json");

var columnNames = [];
var stationData = [];

(async () => {
    const browser = await puppeteer.launch({ headless: true })

    for (let i = 0; i < stationLinks.length; i++) {
        const page = await browser.newPage();
        await page.goto(`https://www.bsis-ice.de/statistik/${stationLinks[i]}`)
        columnNames.splice(0, columnNames.length)

        const getStationData = await page.evaluate((columnNames = []) => {
            const stationName = document.querySelector("body").querySelector("h2").innerText;
            const trTags = document.querySelectorAll("tbody")[0]?.querySelectorAll("tr")
            const columnTags = trTags[0].querySelectorAll("th")

            columnTags.forEach((tag) => {
                columnNames.push(tag.innerText)
            })

            let DATA = []
            let data = {}
            for (let j = 1; j < trTags.length; j++) {
                data = {}
                const ct = trTags[j].querySelectorAll("td")
                for (let k = 0; k < columnNames.length - 1; k++) {
                    data = { ...data, [columnNames[k]]: ct[k]?.innerText }
                }
                let restrictedDays = []
                for (let k = columnNames.length - 1; k < ct.length; k++) {
                    restrictedDays.push(ct[k].innerText)
                }
                data = { ...data, stationName, [columnNames[columnNames.length - 1]]: restrictedDays }
                DATA.push(data)
            }
            return DATA
        })
        stationData.push(getStationData)
    }

    console.log(stationData.length)
    writeFile(path, JSON.stringify(stationData, null, 4), (error) => {
        if (error) {
            console.log('An error has occurred ', error);
            return;
        }
        console.log('Data written successfully to disk');
    });
    await browser.close();
})()