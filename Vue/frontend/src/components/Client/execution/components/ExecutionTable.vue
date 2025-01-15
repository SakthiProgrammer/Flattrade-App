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

    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Trade</v-card-title>
          <v-form>
            <v-card-text>
              <v-text-field label="Stock Name" readonly :value="trade.stock_name" required></v-text-field>
              <v-text-field label="Stock Price" :value="trade.stock_price" readonly required></v-text-field>
              <v-text-field label="Segment" :value="trade.segment" readonly required></v-text-field>
              <v-text-field label="Type" :value="trade.trade_type" readonly required></v-text-field>
              <v-text-field label="Quantity" type="number" min="1" v-model="trade.quantity" required></v-text-field>
              <v-text-field label="Brokerage Fees" :value="trade.brokerageFees" v-model="brokerageFees" readonly
                type="number" required></v-text-field>
              <v-text-field label="Secure Transaction Fees" :value="transactionFees" readonly required></v-text-field>
              <div class="d-flex justify-end ma-4">
                <h2 class="mr-6">Total : </h2>
                <h2> {{ totalPrice }} </h2>
                <!-- {{ trade }} -->
              </div>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="dialog = false">Cancel</v-btn>
              <v-btn color="primary">Save</v-btn>
              <!-- @click="addTrade" -->
            </v-card-actions>
          </v-form>
        </v-card>
      </template>
    </v-dialog>

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
      stocks: [],
      dialog: false,
      /* newTrade: {
        client_id: '',
        stock_id: '',
        price: '',
        quantity: '',
        total_price: '',
        type: '',
      } */
      quantity: 1,
      brokerageFees: 0.2,
      transactionFees: 0.1,
      totalPrice: 0,
      trade: {},
      newTrade: {},
    };
  },
  methods: {
    buyStock(stock) {
      this.dialog = true
      let trade = {
        stock_name: stock.stock_name,
        stock_price: stock.stock_price,
        segment: stock.segment,
        quantity: 1,
        trade_type: "BUY",
      };
      this.trade = trade
      this.updateTotalPrice();
    },


    updateTotalPrice() {
      /*  Base Price: 5 * 100 = 500
       Brokerage Fee: 500 * (0.2 / 100) = 1
       Transaction Fee: 500 * (0.1 / 100) = 0.5
       Total Price: 500 + 1 + 0.5 = 501.5 */
      if (this.trade && this.trade.quantity && this.trade.stock_price && this.brokerageFees !== undefined && this.transactionFees !== undefined) {

        let basePrice = this.trade.quantity * this.trade.stock_price;

        let brokerageAmount = basePrice * (this.brokerageFees / 100);
        let transactionAmount = basePrice * (this.transactionFees / 100);

        this.totalPrice = basePrice + brokerageAmount + transactionAmount;
      } else {
        this.totalPrice = 0;
      }
    },
    sellStock(stock) {
      // alert(`You clicked SELL for ${stock.name}`);
      this.dialog = true
      let trade = {
        stock_name: stock.stock_name,
        stock_price: stock.stock_price,
        segment: stock.segment,
        quantity: 1,
        trade_type: "SELL",
        // total_price: trade.quantity * trade.stock_price,
      };
      this.trade = trade
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
  },
  computed() {

  },
  watch: {
    'trade.quantity': function () {
      this.updateTotalPrice();
    },
    // Watch for changes in brokerage fees and update total price
    'brokerageFees': function () {
      this.updateTotalPrice(); // Recalculate total price when brokerage fees change
    },

    // Watch for changes in transaction fees and update total price
    'transactionFees': function () {
      this.updateTotalPrice(); // Recalculate total price when transaction fees change
    },
  },
};
</script>