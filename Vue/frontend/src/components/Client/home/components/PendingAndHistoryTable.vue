<template>
  <v-card elevation="4">
    <v-tabs v-model="activeTab" color="deep-purple accent-4" grow>
      <v-tab class="pa-3">Pending</v-tab>
      <v-tab class="pa-3">History</v-tab>
    </v-tabs>

    <v-tabs-items v-model="activeTab">
      <!-- Pending Tab -->
      <v-tab-item>
        <v-container fluid>
          <v-row>
            <v-col cols="12">
              <v-data-table :headers="pendingHeaders" :items="TradeData.pending" :items-per-page="5"
                class="elevation-1">
                <template v-slot:item.s_no="{ index }">
                  {{ index + 1 }}
                </template>

                <!-- Trade Type with v-chip -->
                <template v-slot:[`item.trade_type`]="{ item }">
                  <v-chip small :color="item.trade_type === 'BUY' ? 'green' : 'red'" dark>
                    {{ item.trade_type }}
                  </v-chip>
                </template>

                <!-- Status with v-chip -->

              </v-data-table>
            </v-col>
          </v-row>
        </v-container>
      </v-tab-item>

      <!-- History Tab -->
      <v-tab-item>
        <v-container fluid>
          <v-row>
            <v-col cols="12">

              <v-data-table :headers="pendingHeaders" :items="TradeData.history" :items-per-page="5"
                class="elevation-1">

                <template v-slot:item.s_no="{ index }">
                  {{ index + 1 }}
                </template>
                <template v-slot:[`item.trade_type`]="{ item }">
                  <v-chip small :color="item.trade_type === 'Buy' ? 'green' : 'red'" dark>
                    {{ item.trade_type }}
                  </v-chip>
                </template>


              </v-data-table>
            </v-col>
            <!-- {{ pending }} -->
            <!-- {{ TradeData }} -->
          </v-row>
        </v-container>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
export default {

  props: {
    TradeData: {}
  },
  data() {
    return {
      activeTab: 0,
      // Headers for Pending and History Tables
      // pending: this.TradeData.pending,
      pendingHeaders: [
        { text: "S.No.", value: "s_no", sortable: false },
        { text: "Stock Name", value: "stock_detail.stock_name" },
        { text: "Stock Price", value: "stock_detail.stock_price" },
        { text: "Trade Type", value: "trade_type" },
        { text: "Trade Date", value: "trade_date" },
        { text: "Quantity", value: "quantity" },
        { text: "Back Officer Status", value: "back_officer_approval" },
        { text: "Biller Status", value: "biller_approval" },
        { text: "Approver", value: "approver_approval" },
        { text: "Total Price", value: "total_price" },
        // { text: "Status", value: "status" },
      ],
    };

  },
};
</script>
