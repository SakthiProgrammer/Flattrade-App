<template>
  <div>
    <Header @handleComponent="showComponent" />
    <div v-if="currentComponent == 'Home' || currentComponent == null">
      <ClientHomePage />
    </div>
    <div v-else-if="currentComponent == 'Execution'">
      <execution />
    </div>
  </div>

</template>


<script>
import Header from '../header/header.vue';
import Execution from './execution/Execution.vue';
import ClientHomePage from './home/ClientHomePage.vue';
export default {
  data() {
    return {
      currentComponent: "Home",
    }
  },
  components: {
    ClientHomePage,
    Header,
    Execution
  },
  mounted() {
    this.currentComponent = localStorage.getItem("currentComponent");
    // console.log(this.currentComponent);

  },
  methods: {
    showComponent(component) {
      localStorage.setItem("currentComponent", component)
      this.currentComponent = component
    }
  }
}
</script>