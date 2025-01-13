<template>
  <div>
    <v-card>

      <div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">Charge Details</v-card-title>
        <v-btn color="primary" class="mr-5" @click="dialog = true">Add Charge</v-btn>
      </div>
      <v-card-text>
        <v-data-table :headers="headers" :items="charges" class="elevation-1" fixed-header>

          <template v-slot:body="{ items }">
            <tbody>
              <tr v-for="(charge, index) in items" :key="charge.id">
                <td>{{ index + 1 }}</td>
                <td>{{ charge.charge_type }}</td>
                <td>{{ charge.charge_percentage }}</td>
                <td>{{ charge.effective_date }}</td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Upload Charge Details</v-card-title>
          <v-form>
            <v-card-text>
              <v-select label="Charge Type" v-model="newCharge.chargeType" :items="['STT', 'BROKERAGE']"
                required></v-select>
              <v-text-field label="Charge Percentage (%)" v-model="newCharge.chargePercentage" type="number"
                required></v-text-field>
              <v-text-field label="Effective Date" v-model="newCharge.effectiveDate" type="date"
                required></v-text-field>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="dialog = false">Cancel</v-btn>
              <v-btn @click="AddCharge" color="primary">Save</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>

<script>
import EventService from '../../../../Services/EventService';

export default {
  data() {
    return {
      dialog: false,
      newCharge: {
        charge_type: '',
        charge_percentage: null,
        effective_date: '',
      },
      headers: [
        { text: 'S.No.', value: 'sno', sortable: false, align: 'start' },
        { text: 'Charge Type', value: 'charge_type' },
        { text: 'Charge Percentage (%)', value: 'charge_percentage' },
        { text: 'Effective Date', value: 'effective_date' },
      ],
      charges: [],
    };
  },
  methods: {
    submitChargeDetails() {
      this.charges.push({ ...this.newCharge, id: Date.now() });
      this.newCharge = { chargeType: '', chargePercentage: null, effectiveDate: '' };
    },
    AddCharge() {
      EventService.CreateCharge(this.newCharge)
        .then((res) => {
          if (res.data.status == "S") {
            this.charge.push(this.newCharge)
            this.dialog = false;
          } else {
            console.log("err", res.data.errMsg);
          }
        })
    }
  },
  beforeMount() {
    EventService.GetCharges()
      .then((res) => {
        if (res.data.status == "S") {
          this.charges = res.data.charge_details
        } else {
          console.log(res.data.errMsg);
        }
      }).catch((err) => console.log(err));
  }
};
</script>

<style scoped>
.v-data-table {
  width: 100%;
}
</style>
