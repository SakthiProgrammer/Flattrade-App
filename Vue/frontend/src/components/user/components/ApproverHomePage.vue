<template>
  <v-card elevation="2">
    <v-card-title class="text-h5">Trade Records</v-card-title>
    <v-card-text>
      <v-data-table
        :headers="headers"
        :items="tradeRecords"
        class="elevation-1"
        fixed-header
        style="overflow-x: auto"
      >
        <template v-slot:item.s_no="{ index }">
          {{ index + 1 }}
        </template>

        <template v-slot:item.trade_date="{ item }">
          {{ new Date(item.trade_date).toLocaleDateString() }}
        </template>

        <template v-slot:item.action="{ item }">
          <div class="d-flex">
            <v-btn color="green" small @click="approveTrade(item)">
              Approve
            </v-btn>
            <v-btn color="red" small class="ml-2" @click="rejectTrade(item)">
              Reject
            </v-btn>
          </div>
        </template>
      </v-data-table>
    </v-card-text>
  </v-card>
</template>
<script>
//  import EventService from "../../../Services/EventService";
export default {
  data() {
    return {
      headers: [
        { text: "S.No.", value: "s_no", align: "start", sortable: false },
        { text: "Trade Type", value: "trade_type" },
        { text: "Quantity", value: "quantity" },
        { text: "Trade Price", value: "trade_price" },
        { text: "Trade Date", value: "trade_date" },
        {
          text: "Approval Status (Back Officer)",
          value: "back_officer_approval_status",
        },
        { text: "Approval Status (Biller)", value: "biller_approval_status" },
        {
          text: "Approval Status (Approver)",
          value: "approver_approval_status",
        },
        { text: "Client Name", value: "client_name" },
        { text: "Stock Name", value: "stock_name" },
        { text: "Action", value: "action", sortable: false },
      ],
      tradeRecords: [
        {
          id: 1,
          trade_type: "Buy",
          quantity: 100,
          trade_price: 150.75,
          trade_date: "2024-12-01T10:30:00Z",
          back_officer_approval_status: "Pending",
          biller_approval_status: "Pending",
          approver_approval_status: "Pending",
          client_name: "John Doe",
          stock_name: "ABC Corporation",
        },
        {
          id: 2,
          trade_type: "Sell",
          quantity: 50,
          trade_price: 200.25,
          trade_date: "2024-12-02T14:45:00Z",
          back_officer_approval_status: "Approved",
          biller_approval_status: "Pending",
          approver_approval_status: "Pending",
          client_name: "Jane Smith",
          stock_name: "XYZ Industries",
        },
        {
          id: 3,
          trade_type: "Buy",
          quantity: 75,
          trade_price: 300.1,
          trade_date: "2024-12-03T09:15:00Z",
          back_officer_approval_status: "Rejected",
          biller_approval_status: "Approved",
          approver_approval_status: "Pending",
          client_name: "Alice Brown",
          stock_name: "PQR Holdings",
        },
      ],
    };
  },
};
//   },
/*
  methods: {
    approveTrade(trade) {
      trade.back_officer_approval_status = "Approved";
      EventService.ApproveTrade(trade)
        .then((res) => {
          if (res.data.status === "S") {
            console.log("Trade Approved");
            this.removeTrade(trade.id);
          } else {
            console.error("Error:", res.data.errMsg);
          }
        })
        .catch((err) => console.error(err));
    },
    rejectTrade(trade) {
      trade.back_officer_approval_status = "Rejected";
      EventService.RejectTrade(trade)
        .then((res) => {
          if (res.data.status === "S") {
            console.log("Trade Rejected");
            this.removeTrade(trade.id);
          } else {
            console.error("Error:", res.data.errMsg);
          }
        })
        .catch((err) => console.error(err));
    },
    removeTrade(tradeId) {
      this.tradeRecords = this.tradeRecords.filter(
        (trade) => trade.id !== tradeId
      );
    },
  },
  beforeMount() {
    EventService.GetTradeDetails()
      .then((res) => {
        if (res.data.status === "S") {
          this.tradeRecords = res.data.trade_details;
        } else {
          console.error(res.data.errMsg);
        }
      })
      .catch((err) => console.error(err));
  },
};
*/
</script>
