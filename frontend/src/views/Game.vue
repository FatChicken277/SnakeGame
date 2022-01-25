<template>
  <div class="game">
    <v-container class="fill-height">
      <v-row aling="center">
        <v-col cols="8">
          <v-sheet
            class="game-sheet px-5"
            rounded
            color="secondary"
            elevation="5"
          >
            <GameComponent />
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import GameComponent from '@/components/Game/GameComponent.vue';

export default {
  name: 'Game',
  components: {
    GameComponent,
  },
};
</script>

<style>
.game {
  height: 100vh;
}
</style>
