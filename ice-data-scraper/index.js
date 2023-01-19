const { writeFile } = require("fs")
const data = require("./data/config2.json");
const path = "./data/analysis-config.json"

const updatedData = data.map(carr => {
    let updatedArr = []
    carr.forEach(item => {
        let name = item["Station-name"]
        let n = {}
        if (name && name.length > 0) {
            let tokens = name.split(" ")
            let code = ""
            tokens.forEach(token => {
                if (token !== "," && token !== "–") {
                    code += token[0].toUpperCase()
                    if (token.length > 1) {
                        code += token[1].toUpperCase()
                    }
                }
            })
            n = { ...item, "Station-code": code, "Country": "FINLAND", "Country-code": "FI" }
        } else {
            n = {
                ...item, "Station-code": "", "Station-name": "", "Country": "FINLAND", "Country-code": "FI",
            }
        }
        updatedArr.push(n)
    })
    return updatedArr
})

console.log(updatedData.length)

writeFile(path, JSON.stringify(updatedData, null, 4), (error) => {
    if (error) {
        console.log('An error has occurred ', error);
        return;
    }
    console.log('Data written successfully to disk');
});