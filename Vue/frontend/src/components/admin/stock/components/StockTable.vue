<template>
  <div>
    <v-card>
      <div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">Stock Details</v-card-title>
        <v-btn color="primary" class="mr-5" @click="dialog = true">Add Stock</v-btn>
      </div>
      <v-card-text>
        <v-data-table :headers="headers" :items="stocks" class="elevation-1" fixed-header>
          <template v-slot:body="{ items }">
            <tbody>
              <tr v-for="(stock, index) in items" :key="stock.id">
                <td>{{ index + 1 }}</td>
                <td>{{ stock.stock_name }}</td>
                <td>{{ stock.stock_price }}</td>
                <td>{{ stock.segment }}</td>
                <td>{{ stock.isin }}</td>
                <td>
                  <v-btn color="red" small @click="editStock(stock)">
                    <v-icon>mdi-pencil</v-icon>
                  </v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Upload Stock Details</v-card-title>
          <v-form>
            <v-card-text>
              <v-text-field label="Stock Name" v-model="newStock.stock_name" required></v-text-field>
              <v-text-field label="Stock Price" v-model="newStock.stock_price" type="number" required></v-text-field>
              <v-select label="Segment" v-model="newStock.segment" :items="['NSE', 'BSE']" required></v-select>
              <v-text-field label="ISIN" v-model="newStock.isin" required></v-text-field>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="cancelDialog">Cancel</v-btn>
              <v-btn @click="AddStock" color="primary">Save</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
        <!-- {{ newStock }} -->
      </template>
    </v-dialog>
  </div>
</template>

<script>
import EventService from "../../../../Services/EventService";

export default {
  data() {
    return {
      dialog: false,
      edit: false,
      newStock: {
        id: "",
        stock_name: "",
        stock_price: null,
        segment: "",
        isin: "",
      },
      headers: [
        { text: "S.No.", value: "sno", sortable: false, align: "start" },
        { text: "Stock Name", value: "stock_name" },
        { text: "Price", value: "stock_price" },
        { text: "Segment", value: "segment" },
        { text: "ISIN", value: "isin" },
        { text: "Edit", value: "edit", sortable: false, align: "center" },
      ],
      stocks: [],
    };
  },
  methods: {
    submitStockDetails() {
      this.stocks.push({ ...this.newStock, id: Date.now() });
      this.dialog = false;
      this.newStock = { name: "", price: null, segment: "", isin: "" };
    },
    AddStock() {
      this.newStock.stock_price = parseFloat(this.newStock.stock_price);
      if (this.edit) {
        // alert("edit");
        EventService.UpdateStock(this.newStock)
          .then((res) => {
            if (res.data.status == "S") {
              console.log("Succefully Edit");
            } else {
              console.log(res.data.errMsg);
            }
          }).catch((err) => console.log(err)
          )
        this.edit = false
        this.stocks.find((item) => {
          if (this.newStock.id === item.id) {
            item.stock_name = this.newStock.stock_name;
            item.stock_price = this.newStock.stock_price;
            item.isin = this.newStock.isin;
            item.segment = this.newStock.segment;
          }
        });
      } else {
        EventService.CreateStock(this.newStock)
          .then((res) => {
            if (res.data.status == "S") {
              console.log("Stok Created Successfully");
              this.stocks.push(this.newStock);
            } else {
              console.log("error ", res.data.errMsg);
            }
          })
          .catch((err) => console.log(err));
      }
      this.dialog = false;
      this.newStock = { stock_name: "", stock_price: null, segment: "", isin: "" };

    },
    cancelDialog() {
      this.dialog = false;
      this.newStock = { stock_name: "", stock_price: null, segment: "", isin: "" };
    },
    editStock(stock) {
      this.edit = true;
      this.newStock.id = stock.id;
      this.newStock.stock_name = stock.stock_name;
      this.newStock.segment = stock.segment;
      this.newStock.stock_price = stock.stock_price;
      this.newStock.isin = stock.isin;
      this.dialog = true;
    },

  },
  beforeMount() {
    EventService.GetStocks()
      .then((res) => {
        if (res.data.status == "S") {
          this.stocks = res.data.stock_details;
        } else {
          console.log(res.data.errMsg);
        }
      })
      .catch((err) => console.log(err));
  },
};
</script>

<style scoped>
.v-data-table {
  width: 100%;
}
</style>
