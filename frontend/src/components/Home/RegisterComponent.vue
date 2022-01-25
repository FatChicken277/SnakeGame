<template>
  <v-sheet
    color="secondary"
    dark
    elevation="24"
    rounded
    width="350"
    align="center"
  >
    <v-alert
      v-if="alert.status"
      dense
      :color="alert.type === 'error' ? '#cc4949' : '#5beb34'"
    >
    {{ alert.message }}
    </v-alert>
    <v-col cols="9">
      <div class="register py-7">
        <h1>Sign up</h1>
        <div class="register-inputs mt-4">
          <v-form ref="form" v-model="valid">
            <v-text-field
              v-model="username"
              :rules="usernameRules"
              label="Username"
            ></v-text-field>
            <v-text-field
              v-model="password"
              type="password"
              :rules="passwordRules"
              label="Password"
            ></v-text-field>
            <v-text-field
              v-model="passwordConfirm"
              type="password"
              :rules="[passwordConfirmRequired, passwordMatch]"
              label="Password confirm"
            ></v-text-field>
          </v-form>
        </div>
        <v-btn
          block
          color="primary"
          outlined
          class="mt-7"
          :disabled="!valid"
          @click="register"
        >
        Register
        </v-btn>
      </div>
    </v-col>
  </v-sheet>
</template>

<script>
export default {
  name: 'Register',
  data: () => ({
    valid: false,
    username: '',
    usernameRules: [
      (v) => !!v || 'Username is required',
      (v) => v.length < 30 || 'Username must be less than 30 characters',
    ],
    password: '',
    passwordRules: [
      (v) => !!v || 'Password is required',
      (v) => v.length > 6 || 'Password must be great than 6 characters',
    ],
    passwordConfirm: '',
    passwordConfirmRequired: (v) => !!v || 'Password confirmation is required',
    alert: {
      status: false,
      type: '',
      message: '',
    },
  }),
  computed: {
    passwordMatch() {
      return this.password === this.passwordConfirm || 'Passwords must match';
    },
  },
  methods: {
    async register() {
      await this.$store
        .dispatch('register', {
          username: this.username,
          password: this.password,
          passwordConfirm: this.passwordConfirm,
        })
        .then(() => {
          this.$router.push('/login');
        })
        .catch((error) => {
          this.alert = {
            status: true,
            type: 'error',
            message: error.response.data.message.toUpperCase(),
          };
          setTimeout(() => {
            this.alert.status = false;
          }, 2500);
        });

      this.password = '';
      this.passwordConfirm = '';
    },
  },
};
</script>
