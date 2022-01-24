<template>
  <div class="notFound">
    <v-container fill-height>
      <v-row>
        <v-col cols="6">
          <v-img src="../assets/notfound.png" />
        </v-col>
        <v-col cols="6" align="center">
          <div class="notFound-text py-16 mt-5">
            <h1>404</h1>
            <h2>Not Found</h2>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>

export default {
  name: 'NotFound',
  components: {
  },
};
</script>

<style>
  .notFound {
    background-color: black;
    height: 100%;
  }
  .notFound-text {
    width: 70%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    margin-left: 0;
  }
  .notFound-text h1 {
    font-size: 15rem;
    line-height: 13rem;
  }
  .notFound-text h2 {
    font-size: 5rem;
    line-height: 5rem;
  }
</style>
