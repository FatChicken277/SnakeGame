<template>
  <v-sheet
    class="leaderboard my-9 px-5 pt-1"
    rounded
    color="secondary"
    elevation="4"
    align="center"
  >
    <h1>Leaderboard</h1>
    <div class="leaderboard-scroll">
      <v-expansion-panels v-for="(player, index) in leaderboard" v-bind:key="player.username">
        <v-expansion-panel class="mb-2">
          <v-expansion-panel-header>
            <v-img v-if="index < 3"
            class="shrink mr-4"
            width="25"
            height="25"
            :src="require(`../../assets/top${index + 1}.png`)"
          />
            {{ player.username }}
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            Maximum score: {{ player.max_score }}
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>
  </v-sheet>
</template>

<script>
export default {
  data: () => ({
    leaderboard: [],
  }),
  methods: {
    getLeaderboard() {
      this.$store.dispatch('getLeaderboard')
        .then((response) => {
          this.leaderboard = response;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  created() {
    this.getLeaderboard();
  },
};
</script>

<style>
  .leaderboard h1 {
    text-shadow: 4px 4px black;
  }
  .player {
    justify-content: space-between;
  }
  .leaderboard h1 {
    margin: 2rem 2rem;
  }
  .leaderboard {
    min-height: 35rem;
  }
  .v-expansion-panel-header, .v-expansion-panel-content {
    background-color: #525252;
  }
  .leaderboard-scroll {
    overflow-y: scroll;
    max-height: 20rem;
  }
  .leaderboard-scroll::-webkit-scrollbar {
    width: 2px;
  }
</style>
