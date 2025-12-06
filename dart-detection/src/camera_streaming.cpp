//
// Created by hector on 22/07/25.
//

#include "../include/camera_streaming.hpp"

#include <iostream>
#include <opencv2/imgcodecs.hpp>
#include <curl/curl.h>

#include "utils.hpp"

size_t write_callback(char *ptr, size_t size, size_t nmemb, void *userdata) {
    return size*nmemb;
}

void CameraStreaming::stream_frames(const cv::Mat &left_frame, const cv::Mat &front_frame, const cv::Mat &right_frame) {
    std::vector<uchar> left_buf;
    std::vector<uchar> front_buf;
    std::vector<uchar> right_buf;

    // Encode frames to JPEG
    if (!cv::imencode(".jpg", left_frame, left_buf)
        || !cv::imencode(".jpg", front_frame, front_buf)
        || !cv::imencode(".jpg", right_frame, right_buf)) {
        std::cerr << "Failed to encode one of the frames." << std::endl;
        return;
    }

    CURL *curl = curl_easy_init();
    if (!curl) {
        std::cerr << "Failed to initialize CURL." << std::endl;
        return;
    }

    curl_mime *form = curl_mime_init(curl);
    curl_mimepart *part;

    // Left frame
    part = curl_mime_addpart(form);
    curl_mime_name(part, "left");
    curl_mime_filename(part, "left.jpg");
    curl_mime_data(part, reinterpret_cast<const char *>(left_buf.data()), left_buf.size());

    // Front frame
    part = curl_mime_addpart(form);
    curl_mime_name(part, "front");
    curl_mime_filename(part, "front.jpg");
    curl_mime_data(part, reinterpret_cast<const char *>(front_buf.data()), front_buf.size());

    // Front frame
    part = curl_mime_addpart(form);
    curl_mime_name(part, "right");
    curl_mime_filename(part, "right.jpg");
    curl_mime_data(part, reinterpret_cast<const char *>(right_buf.data()), right_buf.size());

    // Set CURL options
    std::string backend_url = utils::get_env_variable("BACKEND_URL");

    if (backend_url.empty()) {
        throw std::runtime_error("BACKEND_URL not set");
    }

    curl_easy_setopt(curl, CURLOPT_URL, backend_url.append("/api/stream/cameras").c_str());
    curl_easy_setopt(curl, CURLOPT_MIMEPOST, form);
    curl_easy_setopt(curl, CURLOPT_TIMEOUT, 5L);
    // Disable libcurl logs
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, nullptr);

    CURLcode res = curl_easy_perform(curl);
    if (res != CURLE_OK) {
        std::cerr << "CURL error: " << curl_easy_strerror(res) << std::endl;
    }

    curl_mime_free(form);
    curl_easy_cleanup(curl);
}
