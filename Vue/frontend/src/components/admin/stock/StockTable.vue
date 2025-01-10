<template>
    <div>
        <v-card>
         <!-- Action Button to Open Dialog -->
<div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">Stock Details</v-card-title>
      <v-btn color="primary" class="mr-5" @click="dialog = true">Add Stock</v-btn>
    </div>
        <v-card-text>
          <v-data-table
            :headers="headers"
            :items="stocks"
            class="elevation-1"
            fixed-header
          >
            <!-- Stock Details -->
            <!-- <template v-slot:body="{ items }">
              <tr v-for="stock in items" :key="stock.id">
                <td>{{ stock.name }}</td>
                <td>{{ stock.price }}</td>
                <td>{{ stock.segment }}</td>
                <td>{{ stock.isin }}</td>
              </tr>
            </template> -->
          </v-data-table>
        </v-card-text>
      </v-card>
  
     
  
      <!-- Dialog to Upload Stock Details -->
      <v-dialog
        v-model="dialog"
        persistent
        max-width="600px"
      >
        <template v-slot:default>
          <v-card>
            <v-card-title class="text-h5">Upload Stock Details</v-card-title>
            <v-form @submit.prevent="submitStockDetails">
              <v-card-text>
                <v-text-field
                  label="Stock Name"
                  v-model="newStock.name"
                  required
                ></v-text-field>
                <v-text-field
                  label="Stock Price"
                  v-model="newStock.price"
                  type="number"
                  required
                ></v-text-field>
                <v-select
                  label="Segment"
                  v-model="newStock.segment"
                  :items="['NSE', 'BSE']"
                  required
                ></v-select>
                <v-text-field
                  label="ISIN"
                  v-model="newStock.isin"
                  required
                ></v-text-field>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn text @click="dialog = false">Cancel</v-btn>
                <v-btn type="submit" color="primary">Save</v-btn>
              </v-card-actions>
            </v-form>
          </v-card>
        </template>
      </v-dialog>
    </div>
</template>

<script>
export default {
    data() {
      return {
        dialog: false,
        newStock: {
          name: '',
          price: null,
          segment: '',
          isin: '',
        },
        headers: [
          { text: 'Stock Name', value: 'name' },
          { text: 'Price', value: 'price' },
          { text: 'Segment', value: 'segment' },
          { text: 'ISIN', value: 'isin' },
        ],
        stocks: [
          { id: 1, name: 'Stock A', price: 100, segment: 'NSE', isin: 'ABCD12N' },
          { id: 2, name: 'Stock B', price: 200, segment: 'BSE', isin: 'XYZ123' },
        ],
      };
    },
    methods: {
      submitStockDetails() {
        this.stocks.push({ ...this.newStock, id: Date.now() });
        this.dialog = false;
        this.newStock = { name: '', price: null, segment: '', isin: '' };
      },
    },
    
}
</script>