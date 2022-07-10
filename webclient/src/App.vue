<template>
  <appnavbar />
  <div class="container py-2">
    <div class="row">
      <div class="col-1">
        <div class="drag-drawflow" draggable="true" @dragstart="drag($event, { input: 1, output: 2 })" data-node="operator">
          <span>Operador Aritmetico</span>
        </div>
        <div class="drag-drawflow" draggable="true" @dragstart="drag($event, { input: 1, output: 0 })" data-node="variable">
          <span>Variable</span>
        </div>
        <div class="drag-drawflow" draggable="true" @dragstart="drag($event, { input: 0, output: 1 })" data-node="block">
          <span>Bloque Inicial</span>
        </div>
        <div class="drag-drawflow" draggable="true" @dragstart="drag($event, { input: 1, output: 2 })" data-node="assign">
          <span>Asignar</span>
        </div>
      </div>
      <div class="col-10">
        <div id="drawflow" @drop="drop($event)" @dragover.prevent @dragenter.prevent></div>
      </div>
      <div class="col-1">
      </div>
    </div>
  </div>
  <appfooter />
</template>
<script>
import Drawflow from 'drawflow'
import { h, getCurrentInstance, render } from 'vue'

import appnavbar from './public/AppNavbar.vue'
import appfooter from './public/AppFooter.vue'

import operator from './components/Operator.vue'
import variable from './components/Variable.vue'
import block from './components/Block.vue'
import assign from './components/Assign.vue'

export default {
  mounted() {
    var id = document.getElementById("drawflow");
    const Vue = { version: 3, h, render };
    this.editor = new Drawflow(id, Vue);
    const internalInstance = getCurrentInstance();
    this.editor.value = new Drawflow(id, Vue, internalInstance.appContext.app._context);

    this.editor.registerNode("operator", operator);
    this.editor.registerNode("variable", variable);
    this.editor.registerNode("block", block);
    this.editor.registerNode("assign", assign);

    this.editor.start();
  },
  methods: {
    drop(evt) {
      var data = JSON.parse(evt.dataTransfer.getData("itemID"))
      this.editor.addNode("github", data.input, data.output, evt.clientX, evt.clientY, "github", {}, data.value, "vue");
    },
    drag(evt, receivedData) {
      evt.dataTransfer.dropEffect = "move";
      evt.dataTransfer.effectAllowed = "move";
      var sendedData = JSON.stringify({
        input: receivedData.input,
        output: receivedData.output,
        value: evt.target.getAttribute("data-node") 
        })
      evt.dataTransfer.setData("itemID", sendedData);
    }
  },
  components: { appnavbar: appnavbar, appfooter: appfooter }
}
</script>
<style>
#drawflow {
  width: 100%;
  height: 550px;
  border: 1px solid #26E07F;
}
</style>