import { createRouter, createWebHistory } from 'vue-router'
import EditorView from '../views/EditorView.vue'
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
            path: '/editor/:id?',
            name: 'editor',
            component: EditorView,
            props: true
        },
    ]
})

export default router