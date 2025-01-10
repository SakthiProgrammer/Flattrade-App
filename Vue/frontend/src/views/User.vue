<template>
    <v-container class="mt-8">
      <user-table/>
    </v-container>
  </template>
  
  <script>
import UserTable from '../components/admin/user/UserTable.vue';


  export default {
    components:{
        UserTable
    }
  };
  </script>
  