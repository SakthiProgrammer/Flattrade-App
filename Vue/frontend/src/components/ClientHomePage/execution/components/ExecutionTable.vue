<template>
  <v-container>
    <v-card>
      <v-card-title class="text-h5">Stock Details</v-card-title>
      <v-card-text>
        <v-data-table :headers="headers" :items="stocks" class="elevation-1" item-value="name">
          <template v-slot:item.segment="{ item }">
            <v-chip :color="item.segment === 'NSE' ? 'blue lighten-3' : 'green lighten-3'" class="white--text">
              {{ item.segment }}
            </v-chip>
          </template>

          <template v-slot:item.action="{ item }">
            <v-btn color="primary" small @click="buyStock(item)">
              BUY
            </v-btn>
            <v-btn color="red darken-1" small class="ml-2" @click="sellStock(item)">
              SELL
            </v-btn>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

  </v-container>
</template>

<script>
import EventService from '../../../../Services/EventService';

export default {
  data() {
    return {
      headers: [
        { text: "Stock Name", value: "stock_name" },
        { text: "Price", value: "stock_price" },
        { text: "Segment", value: "segment" },
        { text: "ISIN", value: "isin" },
        { text: "Action", value: "action", sortable: false },
      ],
      stocks: [
      ],
    };
  },
  methods: {
    buyStock(stock) {
      alert(`You clicked BUY for ${stock.name}`);
    },
    sellStock(stock) {
      alert(`You clicked SELL for ${stock.name}`);
    },
  },
  beforeMount() {
    EventService.GetStocks()
      .then((res) => {
        if (res.data.status == "S") {
          this.stocks = res.data.stock_details
        } else {
          console.log(res.data.errMsg);
        }
      })
  }
};
</script>