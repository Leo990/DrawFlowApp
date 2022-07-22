import axios from "axios";

const BASE_URL = "http://localhost:8082";

var executeProgram = (program) => {
    let data = drawFlowToProgram(program);
    return axios.post(`${BASE_URL}/programs/execute`, data)
};

var indexProgram = () => {
    return axios.get(`${BASE_URL}/programs`);
};

var showProgram = (id) => {
    return axios.get(`${BASE_URL}/programs/${id}`);
};

var storeProgram = (program) => {
    let data = drawFlowToProgram(program);

    return axios.post(`${BASE_URL}/programs/save`, data);
};

function drawFlowToProgram(drawflow) {
    let nodes = drawflow["data"];
    let nodeList = [];
    for (let clave in nodes) {
        inAndOut(nodes[clave]);
        nodeList.push(nodes[clave]);
    }

    let castedProgram = {
        program_name: drawflow.program_name,
        nodes: nodeList,
    };

    if (drawflow["uid"] != "") {
        castedProgram["uid"] = drawflow["uid"];
    }
    return castedProgram;
}

var programToDrawFlow = (program) => {
    let drawflow = {
        drawflow: {
            Home: {
                data: {},
            },
        },
    };
    program["nodes"].forEach((node) => {
        inAndOutReverse(node);
        drawflow.drawflow.Home.data[`${node.id}`] = node;
    });
    return drawflow;
};

function inAndOutReverse(node) {
    let inputs = {};
    let outputs = {};

    if (node["data"] == undefined) {
        node["data"] = {}
    }

    if (node["inputs"] != undefined) {
        node["inputs"].forEach((input, index) => {
            inputs[`input_${index + 1}`] = input;
        });
    }

    if (node["outputs"] != undefined) {
        node["outputs"].forEach((output, index) => {
            outputs[`output_${index + 1}`] = output;
        });
    }
    node["inputs"] = inputs;
    node["outputs"] = outputs;
}

function inAndOut(node) {
    let inputs = [];
    let outputs = [];

    let outputJson = node["outputs"];

    for (let clave in outputJson) {
        outputs.push(outputJson[clave]);
    }

    let inputJson = node["inputs"];

    for (let clave in inputJson) {
        inputs.push(inputJson[clave]);
    }
    node["inputs"] = inputs;
    node["outputs"] = outputs;
}

export {
    executeProgram,
    storeProgram,
    indexProgram,
    showProgram,
    programToDrawFlow,
};