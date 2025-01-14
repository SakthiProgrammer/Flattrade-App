<template>
  <div>
    <Header />
    <div class="mx-7 my-8">
      <ApproverHomePage />
    </div>
  </div>
</template>

<script>
import Header from "../header/header.vue";
import ApproverHomePage from "./components/ApproverHomePage.vue";
export default {
  components: {
    ApproverHomePage,
    Header,
  },
};
</script>
