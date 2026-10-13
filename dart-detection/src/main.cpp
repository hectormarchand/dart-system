//
// Created by hector on 22/07/25.
//

#include <iostream>
#include <opencv2/opencv.hpp>

#include "camera_calibration.hpp"
#include "camera_streaming.hpp"
#include "dart_detector.hpp"
#include "utils.hpp"
#include "../include/camera_manager.hpp"


void camera_frames_iteration(CameraManager &camera_manager,
                             DartDetector &dart_detector,
                             CameraCalibration &left_camera_calibration,
                             bool dev_mode) {
    //cv::Mat left_frame = camera_manager.get_left_frame();
    cv::Mat left_frame = cv::imread("../test-camera-2.png");
    //cv::Mat front_frame = camera_manager.get_front_frame();
    //cv::Mat right_frame = camera_manager.get_right_frame();

    if (dev_mode) {
        cv::imshow("Left camera", left_frame);
        //cv::imshow("Front camera", front_frame);
        //cv::imshow("Right camera", right_frame);
    }

    //CameraStreaming::stream_frames(left_camera_calibration.calibrate(left_frame),
    //                               left_camera_calibration.calibrate(front_frame),
    //                               left_camera_calibration.calibrate(right_frame));

    // Check if a new dart has hit the board
    // if (dart_detector.has_dart_hit_board(front_frame)) {
    //
    //
    //     dart_detector.init_background(front_frame);
    // }
}

int main(int argc, char** argv )
{
    CameraManager camera_manager;
    DartDetector dart_detector;
    CameraCalibration left_camera_calibration;
    const bool dev_mode = utils::get_env_variable("DEV") == "true";

    // Start the loop to stream and display the frames and detect the darts
    for (;;) {
        camera_frames_iteration(camera_manager, dart_detector, left_camera_calibration, dev_mode);

        cv::waitKey(75);
    }

    return 0;
}
