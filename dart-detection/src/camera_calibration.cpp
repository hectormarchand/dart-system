//
// Created by hector on 12/09/2025.
//

#include "include/camera_calibration.h"

#include <iostream>
#include <opencv2/imgproc.hpp>


// VERSION QUI MARCHE BIEN !!!!
// Du moins, très bonne détection de la cible
cv::RotatedRect CameraCalibration::calibrate(cv::Mat &src) {
    cv::Mat hsv, maskRed1, maskRed2, maskRed, maskGreen, mask, redAndGreen, result;

    // Convert to HSV color space
    cv::cvtColor(src, hsv, cv::COLOR_BGR2HSV);

    // --- Red color ranges (wraps around 0 and 180 hue)
    cv::inRange(hsv, cv::Scalar(0, 80, 50), cv::Scalar(10, 255, 255), maskRed1);
    cv::inRange(hsv, cv::Scalar(170, 80, 50), cv::Scalar(180, 255, 255), maskRed2);
    maskRed = maskRed1 | maskRed2;

    int regionMinSize = 100; // Define the min size of a region of the mask in px (to remove red text)
    removeSmallRegions(maskRed, regionMinSize);

    // --- Green color range
    cv::inRange(hsv, cv::Scalar(35, 60, 40), cv::Scalar(90, 255, 255), maskGreen);

    // Combine both masks (red + green)
    mask = maskRed | maskGreen;

    cv::dilate(mask, mask, cv::Mat::ones(6, 6, CV_8U));

    cv::RotatedRect resultEllipse;

    // Find contours on the mask (not the color image!)
    std::vector<std::vector<cv::Point>> contours;
    cv::findContours(mask, contours, cv::RETR_TREE, cv::CHAIN_APPROX_SIMPLE);

    if (!contours.empty()) {
        // Find the largest contour
        auto largest = *std::max_element(contours.begin(), contours.end(),
            [](const auto& a, const auto& b) { return cv::contourArea(a) < cv::contourArea(b); });

        if (largest.size() >= 5) { // Need at least 5 points for an ellipse
            resultEllipse = cv::fitEllipse(largest);
            //cv::ellipse(result, ellipse, cv::Scalar(255, 255, 255), 2);
        }
    }

    return resultEllipse;
}

cv::Mat CameraCalibration::toCircleSpace(cv::RotatedRect &ellipse) {
    return {};
}

void CameraCalibration::removeSmallRegions(cv::Mat &mask, int minSize) {
    cv::Mat labels, stats, centroids;
    int numComponents = cv::connectedComponentsWithStats(mask, labels, stats, centroids);

    cv::Mat cleaned = cv::Mat::zeros(mask.size(), CV_8UC1);

    for (int i = 1; i < numComponents; ++i) { // skip background (label 0)
        int area = stats.at<int>(i, cv::CC_STAT_AREA);
        if (area >= minSize) {
            cleaned.setTo(255, labels == i);
        }
    }

    mask = cleaned;
}