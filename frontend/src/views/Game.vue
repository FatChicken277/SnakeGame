<template>
  <div class="game">
    <v-container class="fill-height">
      <v-row justify="center" style="margin-top: 15vh;">
        <v-col cols="8" xs="12" sm="12" md="12" lg="8">
          <v-sheet
            class="game-sheet px-5"
            rounded
            color="secondary"
            elevation="5"
          >
            <GameComponent />
          </v-sheet>
        </v-col>
        <v-spacer />
        <v-col cols="3" xs="10" sm="10" md="10" lg="3">
          <LeaderboardComponent />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import GameComponent from '@/components/Game/GameComponent.vue';
import LeaderboardComponent from '@/components/Game/LeaderboardComponent.vue';

export default {
  name: 'Game',
  components: {
    GameComponent,
    LeaderboardComponent,
  },
};
</script>

<style>
.game {
  min-height: 100vh;
}
</style>
