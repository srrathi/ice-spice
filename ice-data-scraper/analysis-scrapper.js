const puppeteer = require("puppeteer");
const { writeFile } = require('fs');

const path = './data/config2.json';
const stationLinks = require("./data/stations.json");

var analysisColumnNames = [];
var stationData = [];

(async () => {
    const browser = await puppeteer.launch({ headless: false })

    for (let i = 0; i < stationLinks.length; i++) {
        const page = await browser.newPage();
        await page.goto(`https://www.bsis-ice.de/statistik/${stationLinks[i]}`)
        analysisColumnNames.splice(0, analysisColumnNames.length)

        const getStationData = await page.evaluate((analysisColumnNames = []) => {
            const stationName = document.querySelector("body")?.querySelector("h2").innerText;
            const analysisTrTags = document.querySelectorAll("tbody")[1]?.querySelectorAll("tr")
            let analysisData = []

            if (analysisTrTags && analysisTrTags?.length > 0) {
                const analysisColumnTags = analysisTrTags[0].querySelectorAll("th")

                analysisColumnTags.forEach((tag) => {
                    analysisColumnNames.push(tag.innerText)
                })

                let d = {}
                for (let n = 1; n < analysisTrTags.length; n++) {
                    d = {}
                    const act = analysisTrTags[n]?.querySelectorAll("td")
                    for (let m = 0; m < analysisColumnNames.length; m++) {
                        d = { ...d, [analysisColumnNames[m]]: act[m]?.innerText }
                    }

                    d = { ...d, stationName }
                    analysisData.push(d)
                }
            }

            return analysisData
        })
        stationData.push(getStationData)
    }

    console.log(stationData)
    writeFile(path, JSON.stringify(stationData, null, 4), (error) => {
        if (error) {
            console.log('An error has occurred ', error);
            return;
        }
        console.log('Data written successfully to disk');
    });
    await browser.close();
})()