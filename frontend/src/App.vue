<template>
  <v-app>
    <v-main>
      <router-view name="header"/>
       <router-view/>
      <router-view name="footer"/>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  components: {
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
