<template>
  <div class="register">
    <BackgroundImage>
      <RegisterComponent />
    </BackgroundImage>
  </div>
</template>

<script>
import BackgroundImage from '@/components/Home/BackgroundImage.vue';
import RegisterComponent from '@/components/Home/RegisterComponent.vue';

export default {
  name: 'Register',
  components: {
    BackgroundImage,
    RegisterComponent,
  },
};
</script>

<style>
.register h1 {
  font-size: 3rem;
}
</style>
