//
// Created by hector on 22/07/25.
//

#pragma once

#include <opencv2/core/mat.hpp>

namespace  CameraStreaming {
    void stream_frames(const cv::Mat &left_frame, const cv::Mat &front_frame, const cv::Mat &right_frame);
};
