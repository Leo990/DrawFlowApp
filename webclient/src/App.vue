<template>
  <appnavbar/>
  <RouterView/>
  <appfooter/>
</template>
<script setup>
import { RouterView } from 'vue-router'
import appnavbar from './public/AppNavbar.vue'
import appfooter from './public/AppFooter.vue'
</script>
<style>
  
</style>