<template>
    <v-container>
      <v-card>
        <v-card-title class="text-h5">Stock Details</v-card-title>
        <v-card-text>
          <v-data-table
            :headers="headers"
            :items="stocks"
            class="elevation-1"
            item-value="name"
          >
            <template v-slot:item.segment="{ item }">
              <v-chip
                :color="item.segment === 'NSE' ? 'blue lighten-3' : 'green lighten-3'"
                class="white--text"
              >
                {{ item.segment }}
              </v-chip>
            </template>
  
            <template v-slot:item.action="{ item }">
              <v-btn
                color="primary"
                small
                @click="buyStock(item)"
              >
                BUY
              </v-btn>
              <v-btn
                color="red darken-1"
                small
                class="ml-2"
                @click="sellStock(item)"
              >
                SELL
              </v-btn>
            </template>
          </v-data-table>
        </v-card-text>
      </v-card>
    </v-container>
  </template>
  
  <script>
  export default {
    data() {
      return {
        headers: [
          { text: "Stock Name", value: "name" },
          { text: "Price", value: "price" },
          { text: "Segment", value: "segment" },
          { text: "ISIN", value: "isin" },
          { text: "Action", value: "action", sortable: false },
        ],
        stocks: [
          { name: "Stock A", price: 120.5, segment: "NSE", isin: "ABCD12N" },
          { name: "Stock B", price: 200.75, segment: "BSE", isin: "WXYZ34B" },
          { name: "Stock C", price: 150.3, segment: "NSE", isin: "LMNO56P" },
          { name: "Stock D", price: 98.4, segment: "BSE", isin: "QRST78Q" },
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
  };
  </script>
  