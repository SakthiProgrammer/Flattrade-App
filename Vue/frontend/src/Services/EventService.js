import axios from "axios";

const baseApiClient = axios.create({
  baseURL: `http://localhost:29091`,
  // withCredentials: false,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

export default {
  LoginsApi(body, headers) {
    return baseApiClient.post("/login", body, headers);
  },
  GetBankDetails() {
    return baseApiClient.get("/getbanks");
  },
  CreateBank(Bank) {
    return baseApiClient.post("/createbank", Bank);
  },
  UpdateBank(Bank) {
    return baseApiClient.put("/updatebank", Bank);
  },
  CreateClient(ClientData) {
    return baseApiClient.post("/registerclient", ClientData);
  },
  LoginClient(ClientData) {
    const head = {
      headers: {
        ROLE: "c",
      },
    };
    return baseApiClient.post("/login", ClientData, head);
  },
  LoginAdmin(AdminData, Roldhead) {
    const head = {
      headers: Roldhead,
    };
    console.log(head);
    return baseApiClient.post("/login", AdminData, head);
  },
  LoginUser(UserData, Rolehead) {
    const head = {
      headers: Rolehead,
    };
    console.log(head);
    return baseApiClient.post("/login", UserData, head);
  },
  GetClientDetails() {
    return baseApiClient.get("/getclients");
  },
  GetClientById(head) {
    const header = {
      headers: {
        CLIENTID: head,
      },
    };
    return baseApiClient.get("/getclienttrades", header);
  },
  // GetCliendAndTrade(head) {
  //   const header = {
  //     headers: {
  //       ID: head,
  //     },
  //   };
  //   return baseApiClient.get("/getclients", header);
  // },
  ApproveClient(ClientData) {
    return baseApiClient.put("/updateclient", ClientData);
  },
  RejectClient(ClientData) {
    return baseApiClient.put("/updateclient", ClientData);
  },
  GetStocks() {
    return baseApiClient.get("/getstocks");
  },
  CreateStock(StockData) {
    return baseApiClient.post("/createstock", StockData);
  },
  UpdateStock(stock) {
    return baseApiClient.put("/updatestock", stock);
  },
  GetCharges() {
    return baseApiClient.get("/getcharges");
  },
  CreateCharge(ChargeData) {
    return baseApiClient.post("/createcharge", ChargeData);
  },
  UpdateCharge(ChargeData) {
    return baseApiClient.put("/updatecharge", ChargeData);
  },
  GetUsers() {
    return baseApiClient.get("/getusers");
  },
  CreateUser(User) {
    return baseApiClient.post("/createuser", User);
  },
  UpdateUser(User) {
    return baseApiClient.put("/updateuser", User);
  },
};
