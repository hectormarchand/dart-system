//
// Created by hector on 12/09/2025.
//

#include "include/camera_calibration.hpp"

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
cv::Matx33f CameraCalibration::toCircleSpace(const cv::RotatedRect& ellipse) {
    // OpenCV stores angle in degrees
    float angle = ellipse.angle * static_cast<float>(CV_PI) / 180.0f;

    float x = ellipse.center.x;
    float y = ellipse.center.y;

    // RotatedRect::size contains full width/height (diameters)
    float a = ellipse.size.width * 0.5f;   // semi-major
    float b = ellipse.size.height * 0.5f;  // semi-minor

    float c = std::cos(angle);
    float s = std::sin(angle);

    // Rotation matrices
    cv::Matx33f R1(
         c,  s, 0,
        -s,  c, 0,
         0,  0, 1
    );

    cv::Matx33f R2(
         c, -s, 0,
         s,  c, 0,
         0,  0, 1
    );

    // Translation to origin
    cv::Matx33f T1(
        1, 0, -x,
        0, 1, -y,
        0, 0,  1
    );

    // Translation back
    cv::Matx33f T2(
        1, 0, x,
        0, 1, y,
        0, 0, 1
    );

    // Scale Y so ellipse becomes circle
    float scale = (b != 0.0f) ? (a / b) : 1.0f;

    cv::Matx33f D(
        1,     0, 0,
        0, scale, 0,
        0,     0, 1
    );

    // Final transform: T2 * R2 * D * R1 * T1
    cv::Matx33f M = T2 * R2 * D * R1 * T1;

    return M;
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

cv::Point2f CameraCalibration::applyTransform(const cv::Matx33f& M, const cv::Point2f& p) {
    const cv::Vec3f v(p.x, p.y, 1.0f);
    cv::Vec3f r = M * v;
    return {r[0], r[1]};
}