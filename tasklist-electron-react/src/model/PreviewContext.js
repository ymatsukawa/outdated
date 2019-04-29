import { Record } from 'immutable';
import ListPreviewStrategy from './previewStrategies/ListPreviewStrategy';

const PreviewContextModel = new Record ({
  streamType: null,
  stream: null,
  strategy: null
});

class PreviewContext extends PreviewContextModel {
  constructor(streamType, stream) {
    var strategy = null;
    switch(streamType) {
      case "list":
        strategy = new ListPreviewStrategy(stream);
        break;
      default:
        throw new Error('argument error at Preview Model');
    }

    super({
      streamType: streamType,
      stream: stream,
      strategy: strategy
    });
  }

  getPreviewText() {
    return this.strategy.getPreviewText();
  }
}

export default PreviewContext;