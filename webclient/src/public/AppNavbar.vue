<template>
  <div>
    <nav class="navbar navbar-expand-lg bg-light">
      <div class="container-fluid">
        <img width="16" height="16"
          src="https://cdn.svgporn.com/logos/tensorflow.svg?response-content-disposition=attachment%3Bfilename%3Dtensorflow.svg"
          alt="DrawFlow">&nbsp;
        <RouterLink class="navbar-brand" to="/">DrawFlowApp</RouterLink>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
        </div>
      </div>
    </nav>
  </div>
</template>
<script setup>
import { RouterLink } from 'vue-router';
</script>