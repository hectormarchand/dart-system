// Minimal, self-contained ArUco wrapper. We intentionally avoid
// gocv.io/x/gocv/contrib: cgo compiles every .cpp file in a package
// together, and that package's xfeatures2d.cpp requires
// opencv2/xfeatures2d.hpp, a non-free contrib module Ubuntu's
// libopencv-contrib-dev does not ship — unrelated to aruco, but enough to
// break the whole build. This file only depends on opencv2/aruco.hpp,
// which is present.
#include "aruco.h"

#include <cstdlib>
#include <opencv2/aruco.hpp>
#include <opencv2/core.hpp>

extern "C"
{

    ArucoMarkers Aruco_DetectMarkers(Mat mat, int dictionaryId)
    {
        cv::Mat *image = static_cast<cv::Mat *>(mat);

        cv::Ptr<cv::aruco::Dictionary> dictionary = cv::aruco::getPredefinedDictionary(dictionaryId);
        cv::Ptr<cv::aruco::DetectorParameters> parameters = cv::aruco::DetectorParameters::create();

        std::vector<std::vector<cv::Point2f>> corners, rejected;
        std::vector<int> ids;
        cv::aruco::detectMarkers(*image, dictionary, corners, ids, parameters, rejected);

        ArucoMarkers result;
        result.count = static_cast<int>(ids.size());
        result.ids = nullptr;
        result.corners = nullptr;

        if (result.count == 0)
        {
            return result;
        }

        result.ids = static_cast<int *>(malloc(sizeof(int) * result.count));
        result.corners = static_cast<double *>(malloc(sizeof(double) * result.count * 8));

        for (int i = 0; i < result.count; i++)
        {
            result.ids[i] = ids[i];
            for (int c = 0; c < 4; c++)
            {
                result.corners[i * 8 + c * 2 + 0] = corners[i][c].x;
                result.corners[i * 8 + c * 2 + 1] = corners[i][c].y;
            }
        }

        return result;
    }

    void ArucoMarkers_Close(ArucoMarkers markers)
    {
        free(markers.ids);
        free(markers.corners);
    }
}
