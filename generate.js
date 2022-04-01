'use strict'

const fs = require('fs');
const { execSync } = require('child_process');

// list all folders for target package
const folders = fs.readdirSync('./node_modules/@apex-protocol/core/artifacts/contracts/core');
for (const foldersKey in folders) {
    const folder = folders[foldersKey];
    if (folder === "interfaces") {
        continue;
    }
    const solJson = require(`@apex-protocol/core/artifacts/contracts/core/${folder}/${folder.split('.')[0]}.json`);
    const contractName = solJson['contractName'];
    const abiJson = solJson['abi'];
    // write to file
    fs.writeFileSync(`./abis/${contractName}.json`, JSON.stringify(abiJson));
    // run abigen
    execSync(`abigen --abi=abis/${contractName}.json --pkg=contract --type=${contractName} --out=contract/${contractName}.go`)
}