//
// Created by hector on 22/07/25.
//

#pragma once

#include <opencv2/videoio.hpp>

class CameraManager {
private:
    cv::VideoCapture _left_cam;
    cv::VideoCapture _front_cam;
    cv::VideoCapture _right_cam;

    cv::Mat read_frame(cv::VideoCapture& cap) const;

public:
    CameraManager();
    ~CameraManager();

    cv::Mat get_left_frame();
    cv::Mat get_front_frame();
    cv::Mat get_right_frame();
};

