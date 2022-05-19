import http from "../http-common";
import store from '../store'
class UploadFilesService {
  upload(file, onUploadProgress) {
    let formData = new FormData();
    formData.append("audioFile", file);
    formData.append("addressInfo", store.state.cityCode);
    for (var [a, b] of formData.entries()) {
        console.log(a, b, '--------------');
      }
    return http.post("/Post/query/addQuery", formData, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      },
      onUploadProgress
    });
  }
}
export default new UploadFilesService();
// https://www.bezkoder.com/vuetify-file-upload/