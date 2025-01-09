<template>
  <v-app>

    <app-header />

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import Header from './components/header/header.vue';

export default {
  name: "App",

  data: () => ({
    //
  }),
  components: {
    "app-header": Header
  }
};
</script>
