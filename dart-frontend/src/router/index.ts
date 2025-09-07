import { createRouter, createWebHistory } from "vue-router"
import DartHomeView from '@/main/home/dart-home.view.vue'
import CreateGameView from "@/create-game/create-game.view.vue"
import CameraCalibrationView from "@/camera-calibration/camera-calibration.view.vue";
import DartGameView from '@/main/game/dart-game.view.vue'
import DartboardCommonView from '@/main/dartboard-common.view.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/dartboard/home',
    },
    {
      path: '/dartboard',
      component: DartboardCommonView,
      children: [
        {
          path: 'home',
          name: 'home',
          component: DartHomeView,
        },
        {
          path: 'game',
          name: 'game',
          component: DartGameView,
        },
      ],
    },
    {
      path: '/create-game',
      name: 'create-game',
      component: CreateGameView,
    },
    {
      path: '/camera-calibration',
      name: 'camera-calibration',
      component: CameraCalibrationView,
    },
  ],
})

export default router;
