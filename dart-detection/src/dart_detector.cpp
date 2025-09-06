//
// Created by hector on 22/07/25.
//

#include "../include/dart_detector.hpp"

#include <opencv2/imgproc.hpp>

bool DartDetector::has_dart_hit_board(const cv::Mat &front_frame) const {
    cv::Mat front_frame_to_compare, difference;
    cv::cvtColor(front_frame, front_frame_to_compare, cv::COLOR_BGR2RGB);
    cv::GaussianBlur(front_frame_to_compare, front_frame_to_compare, cv::Size(21, 21), 0);

    cv::absdiff(_background, front_frame_to_compare, difference);
    const int non_zero_difference = cv::countNonZero(difference);

    constexpr int THRESHOLD = 200;
    return non_zero_difference >= THRESHOLD;
}

void DartDetector::init_background(const cv::Mat &front_frame) const {
    cv::Mat background(front_frame);
    cv::cvtColor(background, background, cv::COLOR_BGR2GRAY);
    cv::GaussianBlur(background, _background, cv::Size(21, 21), 0);
}

