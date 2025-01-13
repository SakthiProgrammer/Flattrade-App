<template>
  <v-container>
    {{ currentComponent }}
    <app-header @handleComponent="showComponent" />
    <div v-if="currentComponent === 'Home'">
      <client-table />
    </div>
    <div v-if="currentComponent === 'Brokerage'">
      <brokerage-table />
    </div>
    <div v-if="currentComponent === 'Bank'">
      <bank-table />
    </div>
    <div v-if="currentComponent === 'Stock'">
      <stock-table />
    </div>
    <div v-if="currentComponent === 'User'">
      <user-table />
    </div>
  </v-container>
</template>

<script>
import Header from "../header/header.vue";
import BankTable from "./bank/components/BankTable.vue";
import BrokerageTable from "./brokerage/components/BrokerageTable.vue";
import ClientTable from "./client/clientTable.vue";
import StockTable from "./stock/components/StockTable.vue";
import UserTable from "./user/UserTable.vue";

export default {
  data() {
    return {
      currentComponent: "Home",
    };
  },
  components: {
    "app-header": Header,
    "bank-table": BankTable,
    "brokerage-table": BrokerageTable,
    "client-table": ClientTable,
    "stock-table": StockTable,
    "user-table": UserTable,
  },
  mounted() {

    const savedComponent = localStorage.getItem("currentComponent");
    if (savedComponent) {
      this.currentComponent = savedComponent;
    }
  },
  methods: {
    showComponent(component) {
      this.currentComponent = component;
      localStorage.setItem("currentComponent", component);
    },
  },
};
</script>