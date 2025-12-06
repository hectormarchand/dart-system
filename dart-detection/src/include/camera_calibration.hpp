//
// Created by hector on 11/09/2025.
//

#pragma once
#include <opencv2/core/mat.hpp>

class CameraCalibration {
private:
    // matrix ellipse->circle and its inverse


    // autre chose
    void removeSmallRegions(cv::Mat &mask, int minSize);

public:
    CameraCalibration() = default;
    ~CameraCalibration() = default;

    /**
     * Permet de trouver l'ellipse correspondant à la cible
     * Doit être appelé tous les x secondes pour recalibrer (ne recalibrer que si l'ellipse trouvée est plus grande que l'ancienne)
     * @param src l'image source
     * @return l'ellipse correspondante
     */
    cv::RotatedRect calibrate(cv::Mat& src);

    /**
     * Permet de convertir l'ellipse trouvée dans des coordonées de cercle pour faciliter l'estimation des points
     * Doit être appelé après this::calibrate
     * @param image l'ellipse trouvée dans this::calibrate
     * @return
     */
    cv::Mat toCircleSpace(cv::RotatedRect& ellipse);
};
