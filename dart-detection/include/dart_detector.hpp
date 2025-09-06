//
// Created by hector on 22/07/25.
//

#pragma once
#include <opencv2/core/mat.hpp>

class DartDetector {
private:
    // Background frame took on the front camera, it changes every dart throw to know when a new dart appears on the board
    cv::Mat _background;

public:
    DartDetector() = default;
    ~DartDetector() = default;

    bool has_dart_hit_board(const cv::Mat &front_frame) const;

    void init_background(const cv::Mat &front_frame) const;
};
