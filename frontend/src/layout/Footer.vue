<template>
  <v-footer
      color="secondary"
      dark
      bottom
      height="70"
      class="px-9"
  >
    <v-col cols="10">
      <div class="about">
          Created by FatChicken277 and duvanjm - 2022
      </div>
    </v-col>
    <v-col cols="2">
      <div class="social-media">
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://github.com/FatChicken277"
        >
          <v-icon>mdi-github</v-icon>
        </v-btn>
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://www.linkedin.com/in/alejandro-ramirez-ciceros/"
        >
          <v-icon>mdi-linkedin</v-icon>
        </v-btn>

        <v-divider vertical/>
        
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://github.com/duvanjm"
        >
          <v-icon>mdi-github</v-icon>
        </v-btn>
        <v-btn
          color="primary"
          fab
          x-small
          dark
          href="https://www.linkedin.com/in/duvanjm3/"
        >
          <v-icon>mdi-linkedin</v-icon>
        </v-btn>
      </div>
    </v-col>
  </v-footer>
</template>

<script>
export default {
  name: 'Footer',
};
</script>

<style>
  .social-media {
    display: flex;
    padding-left: 3rem;
    justify-content: space-between;
  }
  .social-media .v-btn .v-icon {
    color: #1E1E1E;
  }
</style>
