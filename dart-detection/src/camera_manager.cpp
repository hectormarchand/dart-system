//
// Created by hector on 22/07/25.
//

#include "../include/camera_manager.hpp"
#include <iostream>
#include <opencv2/highgui.hpp>

CameraManager::CameraManager() {
    _left_cam = cv::VideoCapture("/dev/video-dart-left");
    _front_cam = cv::VideoCapture("/dev/video-dart-front");
    _right_cam = cv::VideoCapture("/dev/video-dart-right");

    if (!_left_cam.isOpened()) {
        throw std::runtime_error("Failed to open left camera");
    }

    if (!_front_cam.isOpened()) {
        throw std::runtime_error("Failed to open front camera");
    }

    if (!_right_cam.isOpened()) {
        throw std::runtime_error("Failed to open right camera");
    }
}

CameraManager::~CameraManager() {
    _left_cam.release();
    _front_cam.release();
    _right_cam.release();
}


cv::Mat CameraManager::read_frame(cv::VideoCapture &cap) const {
    cv::Mat frame;
    cap.read(frame);
    return frame;
}


cv::Mat CameraManager::get_left_frame() {
    return read_frame(_left_cam);
}

cv::Mat CameraManager::get_front_frame() {
    return read_frame(_front_cam);
}

cv::Mat CameraManager::get_right_frame() {
    return read_frame(_right_cam);
}
