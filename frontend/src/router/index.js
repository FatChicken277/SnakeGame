import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import Register from '../views/Register.vue';
import Game from '../views/Game.vue';
import Header from '../layout/Header.vue';
import NotFound from '../views/NotFound.vue';
import Footer from '../layout/Footer.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    components: {
      header: Header,
      default: Home,
      footer: Footer,
    },
    meta: {
      requiresVisitor: true,
    },
  },
  {
    path: '/login',
    name: 'Login',
    components: {
      header: Header,
      default: Login,
      footer: Footer,
    },
    meta: {
      requiresVisitor: true,
    },
  },
  {
    path: '/register',
    name: 'Register',
    components: {
      header: Header,
      default: Register,
      footer: Footer,
    },
    meta: {
      requiresVisitor: true,
    },
  },
  {
    path: '/game',
    name: 'Game',
    components: {
      header: Header,
      default: Game,
      footer: Footer,
    },
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: '*',
    name: '404',
    components: {
      header: Header,
      default: NotFound,
    },
  },
];

const router = new VueRouter({
  routes,
});

export default router;
