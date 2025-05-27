import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import RecordGestures from '@/views/RecordGestures.vue';
import SetupActions from '@/views/SetupActions.vue';
import CreateEditAction from '@/views/CreateEditAction.vue';
import MapGestureActions from '@/views/MapGestureActions.vue'
import ValidateGesture from '@/views/ValidateGesture.vue'
import SetupGestures from '@/views/SetupGestures.vue'
import CreateEditGesture from '@/views/CreateEditGesture.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/setup-actions',
      name: 'setup-actions',
      component: SetupActions,
    },
    {
      path: '/setup-gestures',
      name: 'setup-gestures',
      component: SetupGestures,
    },
    {
      path: '/setup-actions/edit/:id',
      name: 'edit-action',
      component: CreateEditAction,
    },
    {
      path: '/setup-actions/create',
      name: 'create-action',
      component: CreateEditAction,
    },
    {
      path: '/setup-gestures/edit/:id',
      name: 'edit-gesture',
      component: CreateEditGesture,
    },
    {
      path: '/setup-gestures/create',
      name: 'create-gesture',
      component: CreateEditGesture,
    },
    {
      path: '/record-gestures',
      name: 'record-gestures',
      component: RecordGestures,
    },
    {
      path: '/view-gestures',
      name: 'view-gestures',
      component: ValidateGesture
    },
    {
      path: '/map-gestures-to-actions',
      name: 'map-gestures-to-actions',
      component: MapGestureActions
    }
  ],
});

export default router;
