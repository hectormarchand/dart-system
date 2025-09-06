import { createRouter, createWebHistory } from "vue-router"
import DartHomeView from "@/home/dart-home.view.vue"
import CreateGameView from "@/create-game/create-game.view.vue"
import CameraCalibrationView from "@/camera-calibration/camera-calibration.view.vue";
import DartGameView from "@/game/dart-game.view.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: DartHomeView
    },
    {
      path: "/create-game",
      name: "create-game",
      component: CreateGameView
    },
    {
      path: "/game",
      name: "game",
      component: DartGameView
    },
    {
      path: "/camera-calibration",
      name: "camera-calibration",
      component: CameraCalibrationView
    }
  ],
})

export default router;
