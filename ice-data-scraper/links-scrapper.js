const puppeteer = require("puppeteer");

(async () => {
    const browser = await puppeteer.launch({ headless: false })
    const page = await browser.newPage();
    await page.goto("https://www.bsis-ice.de/statistik/Stationindex.html")

    const grabStations = await page.evaluate(() => {
        const stationTags = document.querySelector("body").querySelectorAll("*");
        const stations = []

        let startCollecting = false
        stationTags.forEach(tag => {
            if (tag.getAttribute("id") === "FI") {
                startCollecting = true
            }
            else if (tag.getAttribute("id") === "SE"){
                startCollecting = false
            }
            else if (startCollecting && tag.classList.contains("w3-margin-bottom")) {
                stations.push(tag.querySelector("a").getAttribute("href"))
            }
        })

        return stations
    })

    console.log(grabStations)
    await browser.close();
})()