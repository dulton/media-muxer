#ifndef Muxer_h
#define Muxer_h

#include "capture.h"
#include <map>

class Muxer {
public:
  Muxer();
  virtual ~Muxer();
  bool addOutput(const std::string& uri,
                  enum AVCodecID videoCodec = AV_CODEC_ID_H264,
                  enum AVCodecID audioCodec = AV_CODEC_ID_AAC);
  void writeVideo();
  void writeAudio();
private:
  int mAudioTs, mVideoTs;
  int mVideoId, mAudioId;
  AVAudioFifo* mAudioFifo;
  std::unique_ptr<Capture> mVideo;
  std::unique_ptr<Capture> mAudio;
  std::map<std::string, AVFormatContext*> mOutputs;
};

#endif // Muxer_h