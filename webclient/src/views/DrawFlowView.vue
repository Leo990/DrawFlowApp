<template>
  <offnavbar/>
  <div class="container py-2">
    <div class="row">
      <div class="col-2">
        <span>Parameters</span>
        <ul class="list-group">
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="variable">
            <span>Variable</span>
          </li>
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="constant">
            <span>Constant</span>
          </li>
        </ul>
        <br>
        <span>Conditional</span>
        <ul class="list-group">
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="conditional">
            <span>Conditional</span>
          </li>
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="else">
            <span>Else Statement</span>
          </li>
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="if">
            <span>If Statement</span>
          </li>
        </ul>
        <br>
        <span>Loops</span>
        <ul class="list-group">
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="for">
            <span>For Loop</span>
          </li>
        </ul>
        <br>
        <span>Assignment</span>
        <ul class="list-group">
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="assign">
            <span>Assign</span>
          </li>
          <li class="list-group-item drag-drawflow" draggable="true" @dragstart="drag($event)" data-node="operator">
            <span>Operator</span>
          </li>
        </ul>
      </div>
      <div class="col-10">
        <div id="drawflow" @drop="drop($event)" @dragover.prevent @dragenter.prevent></div>
      </div>
      <div class="col-1">
        <button class="btn btn-sm btn-danger" @click="executeProgram()">Ejecutar programa</button>
        <button class="btn btn-sm btn-success" @click="createProgram()">Guardar programa</button>
      </div>
    </div>
  </div>
</template>

<script>
import Drawflow from 'drawflow'
import { h, getCurrentInstance, render } from 'vue'

import {executeProgram, storeProgram} from '../services/apiService'

import offCanvasNavBar from '../public/OffCanvasNavBar.vue'

import assign from '../components/ast/Assign.vue'
import block from '../components/ast/Block.vue'
import conditional from '../components/ast/Conditional.vue'
import constant from '../components/ast/Constant.vue'
import elseSt from '../components/ast/ElseBody.vue'
import forLoop from '../components/ast/ForLoop.vue'
import ifSt from '../components/ast/IfBody.vue'
import operator from '../components/ast/Operator.vue'
import variable from '../components/ast/Variable.vue'

import base from '../assets/example_original.json'

export default {
  mounted() {
    var id = document.getElementById("drawflow");
    const Vue = { version: 3, h, render };
    this.editor = new Drawflow(id, Vue);
    const internalInstance = getCurrentInstance();
    this.editor.value = new Drawflow(id, Vue, internalInstance.appContext.app._context);

    this.registerNodes()

    this.editor.start();
    this.editor.import(base)
  },
  methods: {
    drop(evt) {
      this.createNode(evt.dataTransfer.getData("typeId"), evt)
    },
    drag(evt) {
      evt.dataTransfer.dropEffect = "move";
      evt.dataTransfer.effectAllowed = "move";
      evt.dataTransfer.setData("typeId", evt.target.getAttribute("data-node"));
    },
    createNode(typeOfNode, evt) {
      switch (typeOfNode) {
        case "assign":
          this.editor.addNode("assign", 1, 2, evt.clientX, evt.clientY, "assign", {}, typeOfNode, "vue");
          break;
        case "block":
          this.editor.addNode("block", 0, 1, evt.clientX, evt.clientY, "block", {}, typeOfNode, "vue");
          break;
        case "conditional":
          this.editor.addNode("conditional", 1, 2, evt.clientX, evt.clientY, "conditional", { type: "" , variable :"", constant : ""}, typeOfNode, "vue");
          break;
        case "constant":
          this.editor.addNode("constant", 1, 0, evt.clientX, evt.clientY, "constant", { value: "" }, typeOfNode, "vue");
          break;
        case "else":
          this.editor.addNode("else", 1, 1, evt.clientX, evt.clientY, "else", {}, typeOfNode, "vue");
          break;
        case "for":
          this.editor.addNode("for", 1, 1, evt.clientX, evt.clientY, "for", {begin : "", end : ""}, typeOfNode, "vue");
          break;
        case "if":
          this.editor.addNode("if", 1, 1, evt.clientX, evt.clientY, "if", {}, typeOfNode, "vue");
          break;
        case "operator":
          this.editor.addNode("operator", 1, 2, evt.clientX, evt.clientY, "operator", { operator: "" }, typeOfNode, "vue");
          break;
        case "variable":
          this.editor.addNode("variable", 1, 0, evt.clientX, evt.clientY, "variable", { name: ""}, typeOfNode, "vue");
          break;
        default:
          break;
      }
    },
    exportProgram() {
      console.log(this.editor.export())
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
    executeProgram(){
      let data_json = this.editor.export()
      executeProgram(data_json)
    },
    createProgram(){
      let data_json = this.editor.export()
      storeProgram(data_json)
    }
  },
  components: { offnavbar : offCanvasNavBar }
}
</script>

<style>
#drawflow {
  width: 100%;
  height: 550px;
  border: 1px solid #535353;
}

:root {
  --dfBackgroundColor: rgba(239, 239, 239, 1);
  --dfBackgroundSize: 10px;
  --dfBackgroundImage: radial-gradient(rgba(0, 0, 0, 1) 1px, transparent 1px);

  --dfNodeType: flex;
  --dfNodeTypeFloat: none;
  --dfNodeBackgroundColor: rgba(0, 0, 255, 1);
  --dfNodeTextColor: rgba(255, 255, 255, 1);
  --dfNodeBorderSize: 0px;
  --dfNodeBorderColor: rgba(255, 255, 255, 1);
  --dfNodeBorderRadius: 5px;
  --dfNodeMinHeight: 40px;
  --dfNodeMinWidth: 160px;
  --dfNodePaddingTop: 15px;
  --dfNodePaddingBottom: 15px;
  --dfNodeBoxShadowHL: 0px;
  --dfNodeBoxShadowVL: 2px;
  --dfNodeBoxShadowBR: 15px;
  --dfNodeBoxShadowS: 2px;
  --dfNodeBoxShadowColor: rgba(0, 0, 0, 1);

  --dfNodeHoverBackgroundColor: rgba(0, 0, 123, 1);
  --dfNodeHoverTextColor: rgba(255, 255, 255, 1);
  --dfNodeHoverBorderSize: 0px;
  --dfNodeHoverBorderColor: rgba(0, 0, 125, 1);
  --dfNodeHoverBorderRadius: 5px;

  --dfNodeHoverBoxShadowHL: 0px;
  --dfNodeHoverBoxShadowVL: 2px;
  --dfNodeHoverBoxShadowBR: 15px;
  --dfNodeHoverBoxShadowS: 2px;
  --dfNodeHoverBoxShadowColor: rgba(0, 0, 0, 1);

  --dfNodeSelectedBackgroundColor: rgba(123, 0, 123, 1);
  --dfNodeSelectedTextColor: #ffffff;
  --dfNodeSelectedBorderSize: 0px;
  --dfNodeSelectedBorderColor: #000000;
  --dfNodeSelectedBorderRadius: 5px;

  --dfNodeSelectedBoxShadowHL: 0px;
  --dfNodeSelectedBoxShadowVL: 2px;
  --dfNodeSelectedBoxShadowBR: 15px;
  --dfNodeSelectedBoxShadowS: 2px;
  --dfNodeSelectedBoxShadowColor: rgba(0, 0, 0, 1);

  --dfInputBackgroundColor: rgba(255, 255, 255, 1);
  --dfInputBorderSize: 2px;
  --dfInputBorderColor: #000000;
  --dfInputBorderRadius: 50px;
  --dfInputLeft: -27px;
  --dfInputHeight: 20px;
  --dfInputWidth: 20px;

  --dfInputHoverBackgroundColor: rgba(0, 0, 0, 1);
  --dfInputHoverBorderSize: 2px;
  --dfInputHoverBorderColor: #000000;
  --dfInputHoverBorderRadius: 50px;

  --dfOutputBackgroundColor: #ffffff;
  --dfOutputBorderSize: 2px;
  --dfOutputBorderColor: #000000;
  --dfOutputBorderRadius: 50px;
  --dfOutputRight: -3px;
  --dfOutputHeight: 20px;
  --dfOutputWidth: 20px;

  --dfOutputHoverBackgroundColor: rgba(0, 0, 0, 1);
  --dfOutputHoverBorderSize: 2px;
  --dfOutputHoverBorderColor: #000000;
  --dfOutputHoverBorderRadius: 50px;

  --dfLineWidth: 3px;
  --dfLineColor: rgba(123, 123, 123, 1);
  --dfLineHoverColor: rgba(123, 123, 123, 1);
  --dfLineSelectedColor: rgba(75, 75, 75, 1);

  --dfRerouteBorderWidth: 0px;
  --dfRerouteBorderColor: rgba(255, 255, 255, 1);
  --dfRerouteBackgroundColor: rgba(0, 0, 255, 1);

  --dfRerouteHoverBorderWidth: 0px;
  --dfRerouteHoverBorderColor: rgba(255, 255, 255, 1);
  --dfRerouteHoverBackgroundColor: rgba(0, 0, 123, 1);

  --dfDeleteDisplay: block;
  --dfDeleteColor: rgba(0, 0, 0, 1);
  --dfDeleteBackgroundColor: rgba(255, 255, 255, 1);
  --dfDeleteBorderSize: 0px;
  --dfDeleteBorderColor: #ffffff;
  --dfDeleteBorderRadius: 25px;
  --dfDeleteTop: -15px;

  --dfDeleteHoverColor: rgba(255, 255, 255, 1);
  --dfDeleteHoverBackgroundColor: rgba(0, 0, 0, 1);
  --dfDeleteHoverBorderSize: 0px;
  --dfDeleteHoverBorderColor: #000000;
  --dfDeleteHoverBorderRadius: 50px;

}

#drawflow {
  background: var(--dfBackgroundColor);
  background-size: var(--dfBackgroundSize) var(--dfBackgroundSize);
  background-image: var(--dfBackgroundImage);
}

.drawflow .drawflow-node {
  display: var(--dfNodeType);
  background: var(--dfNodeBackgroundColor);
  color: var(--dfNodeTextColor);
  border: var(--dfNodeBorderSize) solid var(--dfNodeBorderColor);
  border-radius: var(--dfNodeBorderRadius);
  min-height: var(--dfNodeMinHeight);
  width: auto;
  min-width: var(--dfNodeMinWidth);
  padding-top: var(--dfNodePaddingTop);
  padding-bottom: var(--dfNodePaddingBottom);
  -webkit-box-shadow: var(--dfNodeBoxShadowHL) var(--dfNodeBoxShadowVL) var(--dfNodeBoxShadowBR) var(--dfNodeBoxShadowS) var(--dfNodeBoxShadowColor);
  box-shadow: var(--dfNodeBoxShadowHL) var(--dfNodeBoxShadowVL) var(--dfNodeBoxShadowBR) var(--dfNodeBoxShadowS) var(--dfNodeBoxShadowColor);
}

.drawflow .drawflow-node:hover {
  background: var(--dfNodeHoverBackgroundColor);
  color: var(--dfNodeHoverTextColor);
  border: var(--dfNodeHoverBorderSize) solid var(--dfNodeHoverBorderColor);
  border-radius: var(--dfNodeHoverBorderRadius);
  -webkit-box-shadow: var(--dfNodeHoverBoxShadowHL) var(--dfNodeHoverBoxShadowVL) var(--dfNodeHoverBoxShadowBR) var(--dfNodeHoverBoxShadowS) var(--dfNodeHoverBoxShadowColor);
  box-shadow: var(--dfNodeHoverBoxShadowHL) var(--dfNodeHoverBoxShadowVL) var(--dfNodeHoverBoxShadowBR) var(--dfNodeHoverBoxShadowS) var(--dfNodeHoverBoxShadowColor);
}

.drawflow .drawflow-node.selected {
  background: var(--dfNodeSelectedBackgroundColor);
  color: var(--dfNodeSelectedTextColor);
  border: var(--dfNodeSelectedBorderSize) solid var(--dfNodeSelectedBorderColor);
  border-radius: var(--dfNodeSelectedBorderRadius);
  -webkit-box-shadow: var(--dfNodeSelectedBoxShadowHL) var(--dfNodeSelectedBoxShadowVL) var(--dfNodeSelectedBoxShadowBR) var(--dfNodeSelectedBoxShadowS) var(--dfNodeSelectedBoxShadowColor);
  box-shadow: var(--dfNodeSelectedBoxShadowHL) var(--dfNodeSelectedBoxShadowVL) var(--dfNodeSelectedBoxShadowBR) var(--dfNodeSelectedBoxShadowS) var(--dfNodeSelectedBoxShadowColor);
}

.drawflow .drawflow-node .input {
  left: var(--dfInputLeft);
  background: var(--dfInputBackgroundColor);
  border: var(--dfInputBorderSize) solid var(--dfInputBorderColor);
  border-radius: var(--dfInputBorderRadius);
  height: var(--dfInputHeight);
  width: var(--dfInputWidth);
}

.drawflow .drawflow-node .input:hover {
  background: var(--dfInputHoverBackgroundColor);
  border: var(--dfInputHoverBorderSize) solid var(--dfInputHoverBorderColor);
  border-radius: var(--dfInputHoverBorderRadius);
}

.drawflow .drawflow-node .outputs {
  float: var(--dfNodeTypeFloat);
}

.drawflow .drawflow-node .output {
  right: var(--dfOutputRight);
  background: var(--dfOutputBackgroundColor);
  border: var(--dfOutputBorderSize) solid var(--dfOutputBorderColor);
  border-radius: var(--dfOutputBorderRadius);
  height: var(--dfOutputHeight);
  width: var(--dfOutputWidth);
}

.drawflow .drawflow-node .output:hover {
  background: var(--dfOutputHoverBackgroundColor);
  border: var(--dfOutputHoverBorderSize) solid var(--dfOutputHoverBorderColor);
  border-radius: var(--dfOutputHoverBorderRadius);
}

.drawflow .connection .main-path {
  stroke-width: var(--dfLineWidth);
  stroke: var(--dfLineColor);
}

.drawflow .connection .main-path:hover {
  stroke: var(--dfLineHoverColor);
}

.drawflow .connection .main-path.selected {
  stroke: var(--dfLineSelectedColor);
}

.drawflow .connection .point {
  stroke: var(--dfRerouteBorderColor);
  stroke-width: var(--dfRerouteBorderWidth);
  fill: var(--dfRerouteBackgroundColor);
}

.drawflow .connection .point:hover {
  stroke: var(--dfRerouteHoverBorderColor);
  stroke-width: var(--dfRerouteHoverBorderWidth);
  fill: var(--dfRerouteHoverBackgroundColor);
}

.drawflow-delete {
  display: var(--dfDeleteDisplay);
  color: var(--dfDeleteColor);
  background: var(--dfDeleteBackgroundColor);
  border: var(--dfDeleteBorderSize) solid var(--dfDeleteBorderColor);
  border-radius: var(--dfDeleteBorderRadius);
}

.parent-node .drawflow-delete {
  top: var(--dfDeleteTop);
}

.drawflow-delete:hover {
  color: var(--dfDeleteHoverColor);
  background: var(--dfDeleteHoverBackgroundColor);
  border: var(--dfDeleteHoverBorderSize) solid var(--dfDeleteHoverBorderColor);
  border-radius: var(--dfDeleteHoverBorderRadius);
}
</style>