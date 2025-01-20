<template>
  <v-card>
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
              <v-data-table :headers="pendingHeaders" :items="pendingRecords" :items-per-page="5" class="elevation-1">

                <template v-slot:item.s_no="{ index }">
                  {{ index + 1 }}
                </template>

                <!-- Trade Type with v-chip -->
                <template v-slot:[`item.trade_type`]="{ item }">
                  <v-chip :color="item.trade_type === 'Buy' ? 'green' : 'red'" dark>
                    {{ item.trade_type }}
                  </v-chip>
                </template>

                <!-- Status with v-chip -->
                <template v-slot:[`item.status`]="{ item }">
                  <v-chip color="grey" dark>
                    {{ item.status }}
                  </v-chip>
                </template>
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
              <v-data-table :headers="pendingHeaders" :items="historyRecords" :items-per-page="5" class="elevation-1">
                <!-- Trade Type with v-chip -->
                <template v-slot:item.s_no="{ index }">
                  {{ index + 1 }}
                </template>
                <template v-slot:[`item.trade_type`]="{ item }">
                  <v-chip :color="item.trade_type === 'Buy' ? 'green' : 'red'" dark>
                    {{ item.trade_type }}
                  </v-chip>
                </template>

                <!-- Status with v-chip -->
                <template v-slot:[`item.status`]="{ item }">
                  <v-chip :color="item.status === 'Approved' ? 'green' : 'red'" dark>
                    {{ item.status }}
                  </v-chip>
                </template>
              </v-data-table>
            </v-col>
          </v-row>
        </v-container>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      activeTab: 0,
      // Headers for Pending and History Tables
      pendingHeaders: [
        { text: "S.No.", value: "s_no", sortable: false },
        { text: "Stock Name", value: "name" },
        { text: "Trade Type", value: "trade_type" },
        { text: "Trade Date", value: "trade_date" },
        { text: "Quantity", value: "quantity" },
        { text: "Back Officer Status", value: "back_officer_status" },
        { text: "Biller Status", value: "biller_status" },
        { text: "Approver", value: "approver_status" },
        { text: "Total", value: "total" },
        { text: "Status", value: "status" },
      ],
      // Data for Pending Records
      pendingRecords: [
        {
          name: "Stock A",
          trade_type: "Buy",
          trade_date: "2025-01-08",
          quantity: 100,
          back_officer_status: "Pending",
          biller_status: "Pending",
          approver_status: "Pending",
          total: "$5000",
          status: "Pending",
        },
        {
          name: "Stock B",
          trade_type: "Sell",
          trade_date: "2025-01-09",
          quantity: 50,
          back_officer_status: "Pending",
          biller_status: "Pending",
          approver_status: "Pending",
          total: "$2500",
          status: "Pending",
        },
      ],
      // Data for History Records
      historyRecords: [
        {
          name: "Stock C",
          trade_type: "Buy",
          trade_date: "2025-01-07",
          quantity: 80,
          back_officer_status: "Completed",
          biller_status: "Completed",
          approver_status: "Completed",
          total: "$4000",
          status: "Approved",
        },
        {
          name: "Stock D",
          trade_type: "Sell",
          trade_date: "2025-01-06",
          quantity: 40,
          back_officer_status: "Completed",
          biller_status: "Completed",
          approver_status: "Completed",
          total: "$2000",
          status: "Rejected",
        },
      ],
    };
  },
};
</script>