import axios from 'axios'
const serveSrc = "http://42.192.21.119:8080/";

export const getPercentageDetectionResults = () => axios({
  method: "get",
  url: serveSrc + "Get/query/getPercentageDetectionResults",
});
export const getNumberDetectionByRegion = () => axios({
  method: "get",
  url: serveSrc + "Get/query/getNumberDetectionByRegion",
});
export const getAll = () => axios({
  method: "get",
  url: serveSrc + "Get/query/getAll",
});