import axios from "axios";
export default axios.create({
  baseURL: "http://aruiplex.com:8080",
  headers: {
    "Content-type": "application/json"
  }
});