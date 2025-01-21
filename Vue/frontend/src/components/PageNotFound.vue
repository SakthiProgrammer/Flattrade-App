<template>
  <div>
    <h6 class="text-center title pa-10">404 Page Not Found</h6>
  </div>
</template>

<script>
export default {
  methods: {
    goHome() {
      this.$router.push({ path: "/login" });
    },
  },
};
</script>
