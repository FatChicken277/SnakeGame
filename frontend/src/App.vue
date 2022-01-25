<template>
  <v-app>
    <v-main class="pt-0">
      <router-view name="header"/>
        <fade-transition
          origin="center"
          mode="out-in"
          :duration="250"
        >
          <router-view/>
        </fade-transition>
      <router-view name="footer"/>
    </v-main>
  </v-app>
</template>

<script>
import { FadeTransition } from 'vue2-transitions';

export default {
  name: 'App',
  components: {
    FadeTransition,
  },
  data: () => ({
    //
  }),
};
</script>

<style>
  .v-main {
    font-family: 'VT323', monospace;
  }
    /* width */
  ::-webkit-scrollbar {
    width: 5px;
  }
  ::-webkit-scrollbar-track {
    box-shadow: inset 0 0 5px #1E1E1E;
    border-radius: 2px;
  }
  /* Handle */
  ::-webkit-scrollbar-thumb {
    background: #6ddb18;
    border-radius: 10px;
  }
  /* Handle on hover */
  ::-webkit-scrollbar-thumb:hover {
    background: #6dad00;
  }
  #app {
    background-color: var(--v-background-base);
  }
</style>
