<template>
  <v-app-bar
      app
      color="secondary"
      dark
      class="px-9"
  >
    <div class="d-flex align-center">
      <v-img
          alt="Snake Logo"
          class="shrink mr-2"
          src="../assets/logo.png"
          transition="scale-transition"
          width="50"
      />
      <h1 style="text-shadow: 4px 4px black;">Snakee</h1>
    </div>

    <v-spacer></v-spacer>

    <router-link to="/register" v-if="!loggedIn">
      <v-btn
        color="primary"
        outlined
        class="mr-4"
      >
        Register
      </v-btn>
    </router-link>
    <router-link to="/login" v-if="!loggedIn">
      <v-btn
        color="primary"
        style="color: #1E1E1E;"
      >
        Login
      </v-btn>
    </router-link>
    <v-btn
      v-if="loggedIn"
      color="primary"
      outlined
      @click="signOut"
    >
      Logout
    </v-btn>
  </v-app-bar>
</template>

<script>
export default {
  name: 'Header',
  computed: {
    loggedIn() {
      return this.$store.getters.loggedIn;
    },
  },
  methods: {
    async signOut() {
      await this.$store.dispatch('signOut');
      this.$router.push('/');
    },
  },
};
</script>

<style>
</style>
