import { createRouter, createWebHistory } from 'vue-router'
import DrawFlowView from '../views/DrawFlowView.vue'
import HomeView from '../views/HomeView.vue'
import NotFoundView from '../views/NotFoundView.vue'

const router = createRouter({
    history: createWebHistory(
        import.meta.env.BASE_URL),
    routes: [{
            path: '/:pathMatch(.*)*',
            name: '404',
            component: NotFoundView
        },
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/draw-flow',
            name: 'drawflow',
            component: DrawFlowView
        },
    ]
})

export default router