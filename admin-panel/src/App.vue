<template>
  <div id="app">
    <input type="text" placeholder="Username" id="username" />
    <br>
    <input type="password" placeholder="Password" id="password" />

  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

export default class App extends Vue {}
</script>
<style>
  div {
    width: 100px;
    height: 100px;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
  
    
  }
  
</style>
