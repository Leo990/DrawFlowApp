var data = require('./example_original.json')

var nodes = data["drawflow"]["Home"]["data"]
var nodeList = []
for (let clave in nodes) {
    inAndOut(nodes[clave])
    nodeList.push(nodes[clave])
}

console.log(nodeList);

function inAndOut(node) {
    let inputs = []
    let outputs = []

    let outputJson = node["outputs"]

    for (let clave in outputJson) {
        outputs.push(outputJson[clave])
    }

    let inputJson = node["inputs"]

    for (let clave in inputJson) {
        inputs.push(inputJson[clave])
    }

    node["inputs"] = inputs
    node["outputs"] = outputs
}