//record-sdk.js
import Recorder from "./recorder";
export default class Record {
  startRecord (param) {
    let self = this;
    try {
      Recorder.get(rec => {
        if (rec.error) return param.error(rec.error);
        self.recorder = rec;
        self.recorder.start();
        param.success("Start the recording");
      })
    } catch (e) {
      param.error("Failed to start recording" + e);
    }
  }

  stopRecord (param) {
    let self = this;
    try {
      let blobData = self.recorder.getBlob();
      param.success(blobData);
    } catch (e) {
      param.error("Failed to end recording" + e);
    }
  }

  play (audio) {
    let self = this;
    try {
      self.recorder.play(audio);
    } catch (e) {
      console.error("Recording playback failure" + e);
    }
  }

  clear (audio) {
    let self = this;
    try {
      self.recorder.clear(audio);
    } catch (e) {
      console.error("Failed to clear recording" + e);
    }
  }
}