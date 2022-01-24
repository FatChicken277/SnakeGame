import Vue from 'vue';
import VueRouter from 'vue-router';
import Header from '../layout/Header.vue';
import Home from '../views/Home.vue';
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
];

const router = new VueRouter({
  routes,
});

export default router;
