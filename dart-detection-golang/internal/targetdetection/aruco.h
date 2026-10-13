#ifndef DART_DETECTION_ARUCO_H
#define DART_DETECTION_ARUCO_H

#ifdef __cplusplus
extern "C"
{
#endif

typedef void *Mat;

// ArucoMarkers is the flattened result of a marker detection pass: `count`
// markers, each with an integer ID and 4 (x, y) corners (top-left,
// top-right, bottom-right, bottom-left, per OpenCV's convention). ids has
// `count` entries, corners has `count * 8` entries. Both are heap
// allocated by Aruco_DetectMarkers and must be released via
// ArucoMarkers_Close.
typedef struct ArucoMarkers
{
    int count;
    int *ids;
    double *corners;
} ArucoMarkers;

ArucoMarkers Aruco_DetectMarkers(Mat mat, int dictionaryId);
void ArucoMarkers_Close(ArucoMarkers markers);

#ifdef __cplusplus
}
#endif

#endif // DART_DETECTION_ARUCO_H
