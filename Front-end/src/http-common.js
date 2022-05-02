import axios from "axios";
export default axios.create({
  baseURL: "42.192.21.119:8080/",
  headers: {
    "Content-type": "application/json"
  }
});