import axios from 'axios'

const BASE_URL = 'http://localhost:8082'

var executeProgram = (program) => {
    let data = castProgram(program)
    var dataJson
        /* console.log(data);
        axios.post(`${BASE_URL}/programs/execute`, data).then(function(res) {
                if (res.status == 200) {
                    dataJson = res.data;
                }
                console.log(res);
            })
            .catch(function(err) {
                dataJson = err
            })
        return dataJson */

}

var storeProgram = (program) => {
    let data = castProgram(program)
        /* var dataJson
        axios.post(`${BASE_URL}/programs/save`, data).then(function(res) {
                if (res.status == 200) {
                    dataJson = "res.data";
                }
            })
            .catch(function(err) {
                dataJson = err
            })

        return dataJson */

}

function castProgram(program) {
    var nodes = program["drawflow"]["Home"]["data"]
    var nodeList = []
    for (let clave in nodes) {
        inAndOut(nodes[clave])
        nodeList.push(nodes[clave])
    }

    return { "nodes": nodeList }
}

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

export {
    executeProgram,
    storeProgram
}