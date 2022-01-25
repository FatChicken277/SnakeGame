<template>
  <v-container fluid class="pa-0">
    <v-img
      class="home-img"
      src=""
      lazy-src="../../assets/background.jpg"
    >
      <template v-slot:placeholder>
        <v-row
          class="fill-height ma-0 placeholder"
          align="center"
          justify="center"
        >
          <h1>Welcome to Snakee</h1>
          <router-link to="/login">Login</router-link>
        </v-row>
      </template>
    </v-img>
  </v-container>
</template>

<script>
export default {};
</script>

<style>
  .home-img {
    width: 100%;
    height: 32.9rem;
  }
  .placeholder {
    flex-direction: column;
  }
  .placeholder h1 {
    font-size: 5rem;
  }
  .placeholder a {
    font-size: 3rem;
    z-index: 10 !important;
  }
  .home-img .v-responsive__sizer {
    z-index: -1;
  }
  .home-img .v-responsive__content {
    z-index: -2;
  }
</style>
