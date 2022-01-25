<template>
  <div class="home">
    <BackgroundImage>
      <h1>Welcome to Snakee</h1>
      <router-link to="/login">Login</router-link>
    </BackgroundImage>
  </div>
</template>

<script>
import BackgroundImage from '@/components/Home/BackgroundImage.vue';

export default {
  name: 'Home',
  components: {
    BackgroundImage,
  },
};
</script>
