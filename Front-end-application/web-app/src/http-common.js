import axios from "axios";
export default axios.create({
  baseURL: "http://42.192.21.119:8080",
  headers: {
    "Content-type": "application/json"
  }
});