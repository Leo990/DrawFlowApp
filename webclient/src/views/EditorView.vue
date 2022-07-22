<template >
    <div class="container nav-scroller bg-body shadow-sm">
        <nav class="nav" aria-label="Secondary navigation">
            <a class="nav-link active material-icons" aria-current="page" @click="storeProgram()">
                <span class="material-icons">
                    save
                </span>
            </a>
            <a class="nav-link" aria-current="page" @click="executeProgram()">
                <span class="material-icons">
                    play_arrow
                </span>
            </a>
            <form @submit.prevent class="d-flex position-relative">
                <span></span>
                <input v-model="program_name"
                    class="form-control form-control-sm h-50 position-relative translate-middle-y top-50" type="text"
                    placeholder="Program Name" aria-label="name">
            </form>
        </nav>
    </div>
    <div class="container py-2">
        <div class="row">
            <div class="col-2">
                <span>Parameters</span>
                <ul class="list-group">
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="variable">
                        <span>Variable</span>
                    </li>
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="constant">
                        <span>Constant</span>
                    </li>
                </ul>
                <br>
                <span>Conditional</span>
                <ul class="list-group">
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="conditional">
                        <span>Conditional</span>
                    </li>
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="else">
                        <span>Else Statement</span>
                    </li>
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="if">
                        <span>If Statement</span>
                    </li>
                </ul>
                <br>
                <span>Loops</span>
                <ul class="list-group">
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="for">
                        <span>For Loop</span>
                    </li>
                </ul>
                <br>
                <span>Assignment</span>
                <ul class="list-group">
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="assign">
                        <span>Assign</span>
                    </li>
                    <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)"
                        data-node="operator">
                        <span>Operator</span>
                    </li>
                </ul>
            </div>
            <div class="col-10">
                <div id="drawflow" @drop="drop($event)" @dragover.prevent @dragenter.prevent></div>
            </div>
            <div class="col-2">

            </div>
            <div class="col-10" v-show="hasExecute()">
                <pre class="language-py" data-prismjs-copy="Copy code!!!" ><code class="language-py">{{ programOut.code }}</code></pre>
            </div>
            <div class="col-2">

            </div>
            <div class="col-10" v-show="hasExecute()">
                <pre class="language-bash" v-if="programOut.error != ''"><code>{{ programOut.error }}</code></pre>
                <pre class="language-bash" v-else><code>{{ programOut.output }}</code></pre>
            </div>
        </div>
    </div>

</template>
<script>
import { executeProgram, storeProgram, showProgram, programToDrawFlow } from '../services/apiService'
import Drawflow from 'drawflow'
import { h, getCurrentInstance, render } from 'vue'

import assign from '../components/ast/Assign.vue'
import block from '../components/ast/Block.vue'
import conditional from '../components/ast/Conditional.vue'
import constant from '../components/ast/Constant.vue'
import elseSt from '../components/ast/ElseBody.vue'
import forLoop from '../components/ast/ForLoop.vue'
import ifSt from '../components/ast/IfBody.vue'
import operator from '../components/ast/Operator.vue'
import variable from '../components/ast/Variable.vue'

import base from '../assets/base.json'
export default {
    mounted() {
        var drawflowID = document.getElementById("drawflow");
        const Vue = { version: 3, h, render };
        this.editor = new Drawflow(drawflowID, Vue);
        const internalInstance = getCurrentInstance();
        this.editor.value = new Drawflow(drawflowID, Vue, internalInstance.appContext.app._context);

        this.registerNodes()

        this.editor.start();

        this.id == "" ? this.editor.import(base) : showProgram(this.id).then((res) => {
            let programRes = Object.assign({}, res.data["program"][0])
            this.program_name = programRes["program_name"];
            let drawflow = programToDrawFlow(programRes)
            this.editor.import(drawflow)
        })

    },
    methods: {
        hasExecute() {
            return this.programOut.code != ""
        }
        ,
        drop(evt) {
            this.createNode(evt.dataTransfer.getData("typeId"), evt)
        },
        drag(evt) {
            evt.dataTransfer.dropEffect = "move";
            evt.dataTransfer.effectAllowed = "move";
            evt.dataTransfer.setData("typeId", evt.target.getAttribute("data-node"));
        }
        ,
        createNode(typeOfNode, evt) {
            let pos_x = evt.clientX
            let pos_y = evt.clientY
            pos_x = pos_x * (this.editor.precanvas.clientWidth / (this.editor.precanvas.clientWidth * this.editor.zoom)) - (this.editor.precanvas.getBoundingClientRect().x * (this.editor.precanvas.clientWidth / (this.editor.precanvas.clientWidth * this.editor.zoom)));
            pos_y = pos_x * (this.editor.precanvas.clientHeight / (this.editor.precanvas.clientHeight * this.editor.zoom)) - (this.editor.precanvas.getBoundingClientRect().y * (this.editor.precanvas.clientHeight / (this.editor.precanvas.clientHeight * this.editor.zoom)));
            switch (typeOfNode) {
                case "assign":
                    this.editor.addNode("assign", 1, 2, pos_x, pos_y, "assign", {}, typeOfNode, "vue");
                    break;
                case "block":
                    this.editor.addNode("block", 0, 1, pos_x, pos_y, "block", {}, typeOfNode, "vue");
                    break;
                case "conditional":
                    this.editor.addNode("conditional", 1, 2, pos_x, pos_y, "conditional", { type: "", variable: "", constant: "" }, typeOfNode, "vue");
                    break;
                case "constant":
                    this.editor.addNode("constant", 1, 0, pos_x, pos_y, "constant", { value: "" }, typeOfNode, "vue");
                    break;
                case "else":
                    this.editor.addNode("else", 1, 1, pos_x, pos_y, "else", {}, typeOfNode, "vue");
                    break;
                case "for":
                    this.editor.addNode("for", 1, 1, pos_x, pos_y, "for", { begin: "", end: "" }, typeOfNode, "vue");
                    break;
                case "if":
                    this.editor.addNode("if", 1, 1, pos_x, pos_y, "if", {}, typeOfNode, "vue");
                    break;
                case "operator":
                    this.editor.addNode("operator", 1, 2, pos_x, pos_y, "operator", { operator: "" }, typeOfNode, "vue");
                    break;
                case "variable":
                    this.editor.addNode("variable", 1, 0, pos_x, pos_y, "variable", { name: "" }, typeOfNode, "vue");
                    break;
                default:
                    break;
            }
        },
        registerNodes() {
            this.editor.registerNode("operator", operator);
            this.editor.registerNode("conditional", conditional);
            this.editor.registerNode("constant", constant);
            this.editor.registerNode("variable", variable);
            this.editor.registerNode("block", block);
            this.editor.registerNode("assign", assign);
            this.editor.registerNode("if", ifSt);
            this.editor.registerNode("else", elseSt);
            this.editor.registerNode("for", forLoop);
        },
        async storeProgram() {
            let response = await storeProgram(this.exportProgram())
        },
        async executeProgram() {
            let response = await executeProgram(this.exportProgram())
            this.programOut = response.data
        },
        exportProgram() {
            let program = this.editor.export()["drawflow"]["Home"]
            program["program_name"] = this.program_name
            program["uid"] = this.id
            return program
        }
    },
    props: {
        id: String
    },
    data: () => ({
        program_name: "",
        programOut: {
            "code": "",
            "error": "",
            "output": ""
        }
    })
}
</script>

<style scoped>
html,
body {
    overflow-x: hidden;
    /* Prevent scroll on narrow devices */
}

body {
    padding-top: 56px;
}

@media (max-width: 991.98px) {
    .offcanvas-collapse {
        position: fixed;
        top: 56px;
        /* Height of navbar */
        bottom: 0;
        left: 100%;
        width: 100%;
        padding-right: 1rem;
        padding-left: 1rem;
        overflow-y: auto;
        visibility: hidden;
        background-color: #343a40;
        transition: transform .3s ease-in-out, visibility .3s ease-in-out;
    }

    .offcanvas-collapse.open {
        visibility: visible;
        transform: translateX(-100%);
    }
}

.nav-scroller .nav {
    color: rgba(255, 255, 255, .75);
}

.nav-scroller .nav-link {
    padding-top: .75rem;
    padding-bottom: .75rem;
    font-size: .875rem;
    color: #6c757d;
}

.nav-scroller .nav-link:hover {
    color: #007bff;
}

.nav-scroller .active {
    font-weight: 500;
    color: #343a40;
}

.bg-purple {
    background-color: #6f42c1;
}

.bd-placeholder-img {
    font-size: 1.125rem;
    text-anchor: middle;
    -webkit-user-select: none;
    -moz-user-select: none;
    user-select: none;
}

@media (min-width: 768px) {
    .bd-placeholder-img-lg {
        font-size: 3.5rem;
    }
}

.b-example-divider {
    height: 3rem;
    background-color: rgba(0, 0, 0, 0.1);
    border: solid rgba(0, 0, 0, 0.15);
    border-width: 1px 0;
    box-shadow: inset 0 0.5em 1.5em rgba(0, 0, 0, 0.1), inset 0 0.125em 0.5em rgba(0, 0, 0, 0.15);
}

.b-example-vr {
    flex-shrink: 0;
    width: 1.5rem;
    height: 100vh;
}

.bi {
    vertical-align: -0.125em;
    fill: currentColor;
}

.nav-scroller {
    position: relative;
    z-index: 2;
    height: 2.75rem;
    overflow-y: hidden;
}

.nav-scroller .nav {
    display: flex;
    flex-wrap: nowrap;
    padding-bottom: 1rem;
    margin-top: -1px;
    overflow-x: auto;
    text-align: center;
    white-space: nowrap;
    -webkit-overflow-scrolling: touch;
}
</style>