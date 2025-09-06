//
// Created by hector on 22/07/25.
//

#include <iostream>
#include <opencv2/opencv.hpp>

#include "camera_streaming.h"
#include "dart_detector.hpp"
#include "utils.hpp"
#include "../include/camera_manager.hpp"


void camera_frames_iteration(CameraManager &camera_manager,
                             DartDetector &dart_detector,
                             bool dev_mode) {
    cv::Mat left_frame = camera_manager.get_left_frame();
    cv::Mat front_frame = camera_manager.get_front_frame();
    cv::Mat right_frame = camera_manager.get_right_frame();

    if (dev_mode) {
        cv::imshow("Left camera", left_frame);
        cv::imshow("Front camera", front_frame);
        cv::imshow("Right camera", right_frame);
    }

    CameraStreaming::stream_frames(left_frame, front_frame, right_frame);

    // Check if a new dart has hit the board
    // if (dart_detector.has_dart_hit_board(front_frame)) {
    //
    //
    //     dart_detector.init_background(front_frame);
    // }
}

int main(int argc, char** argv )
{
    std::cout << "Hello World!" << std::endl;

    // cv::VideoCapture cap;
    // cap.open("/dev/video-dart-left");
    //
    // if (!cap.isOpened()) {
    //     std::cerr << "ERROR! Unable to open camera\n";
    //     return -1;
    // }
    //
    // cv::Mat frame;
    // for (;;)
    // {
    //     // wait for a new frame from camera and store it into 'frame'
    //     cap.read(frame);
    //     // check if we succeeded
    //     if (frame.empty()) {
    //         std::cerr << "ERROR! blank frame grabbed\n";
    //         break;
    //     }
    //     // show live and wait for a key with timeout long enough to show images
    //     cv::imshow("Live", frame);
    //     if (cv::waitKey(5) >= 0)
    //         break;
    // }

    CameraManager camera_manager;
    DartDetector dart_detector;
    const bool dev_mode = utils::get_env_variable("DEV") == "true";

    // Start the loop to stream and display the frames and detect the darts
    for (;;) {
        camera_frames_iteration(camera_manager, dart_detector, dev_mode);

        cv::waitKey(75);
    }

    return 0;
}
