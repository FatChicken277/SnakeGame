<template>
  <div class="login">
    <BackgroundImage>
      <LoginComponent />
    </BackgroundImage>
  </div>
</template>

<script>
import LoginComponent from '@/components/Home/LoginComponent.vue';
import BackgroundImage from '@/components/Home/BackgroundImage.vue';

export default {
  name: 'Login',
  components: {
    LoginComponent,
    BackgroundImage,
  },
};
</script>

<style>
.login h1 {
  font-size: 3rem;
}
</style>
