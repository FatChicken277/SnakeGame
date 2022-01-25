<template>
  <v-row>
    <v-col class="12">
      <div :id="containerId" v-if="loaded" />
      <div class="loader" v-else>
        <v-progress-linear
          :size="70"
          :width="7"
          color="primary"
          indeterminate
        ></v-progress-linear>
      </div>
    </v-col>
  </v-row>
</template>

<script>
import { launch } from '@/game/game';

export default {
  name: 'Game',
  data() {
    return {
      loaded: false,
      gameInstance: null,
      containerId: 'game-container',
    };
  },
  async mounted() {
    this.loaded = true;
    this.$nextTick(() => {
      this.gameInstance = launch(this.containerId);
    });
  },
  destroyed() {
    this.gameInstance.destroy(false);
  },
};
</script>

<style>
.loader {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
