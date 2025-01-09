import axios from "axios";

const baseApiClient = axios.create({
  baseURL: `http://localhost:29091`,
  withCredentials: false,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

export default {
  LoginsApi(body, headers) {
    return baseApiClient.post("/login", body, headers);
  },
};
