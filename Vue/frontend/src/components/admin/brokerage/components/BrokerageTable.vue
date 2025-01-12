<template>
  <div>
    <v-card>
      <!-- Action Button to Open Dialog -->
      <div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">Charge Details</v-card-title>
        <v-btn color="primary" class="mr-5" @click="dialog = true">Add Charge</v-btn>
      </div>
      <v-card-text>
        <v-data-table
          :headers="headers"
          :items="charges"
          class="elevation-1"
          fixed-header
        >
          <!-- S.No. Column -->
          <template v-slot:body="{ items }">
            <tbody>
              <tr v-for="(charge, index) in items" :key="charge.id">
                <td>{{ index + 1 }}</td> <!-- S.No Column -->
                <td>{{ charge.chargeType }}</td>
                <td>{{ charge.chargePercentage }}</td>
                <td>{{ charge.effectiveDate }}</td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Dialog to Upload Charge Details -->
    <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
    >
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Upload Charge Details</v-card-title>
          <v-form @submit.prevent="submitChargeDetails">
            <v-card-text>
              <v-select
                label="Charge Type"
                v-model="newCharge.chargeType"
                :items="['STT', 'BROKERAGE']"
                required
              ></v-select>
              <v-text-field
                label="Charge Percentage (%)"
                v-model="newCharge.chargePercentage"
                type="number"
                required
              ></v-text-field>
              <v-text-field
                label="Effective Date"
                v-model="newCharge.effectiveDate"
                type="date"
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
      newCharge: {
        chargeType: '',
        chargePercentage: null,
        effectiveDate: '',
      },
      headers: [
        { text: 'S.No.', value: 'sno', sortable: false, align: 'start' },
        { text: 'Charge Type', value: 'chargeType' },
        { text: 'Charge Percentage (%)', value: 'chargePercentage' },
        { text: 'Effective Date', value: 'effectiveDate' },
      ],
      charges: [
        { id: 1, chargeType: 'STT', chargePercentage: 0.1, effectiveDate: '2025-01-01' },
        { id: 2, chargeType: 'BROKERAGE', chargePercentage: 0.5, effectiveDate: '2025-01-10' },
      ],
    };
  },
  methods: {
    submitChargeDetails() {
      this.charges.push({ ...this.newCharge, id: Date.now() });
      this.dialog = false;
      this.newCharge = { chargeType: '', chargePercentage: null, effectiveDate: '' };
    },
  },
};
</script>

<style scoped>
.v-data-table {
  width: 100%;
}
</style>
