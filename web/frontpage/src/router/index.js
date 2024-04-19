import Vue from 'vue'
import VueRouter from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import ArtList from '@/components/ArticleListView'
import Detail from '@/components/DetailView'

Vue.use(VueRouter)

const routes = [
    // {
    //     path: '/',
    //     name: 'home',
    //     component: HomeView,
    //     children: [
    //         {
    //             path: '/',
    //             component: ArtList,
    //             meta: {
    //                 title: "GinBlog"
    //             }
    //         }
    //     ]
    // },
    {
        path: '/',
        component: ArtList,
        meta: {
            title: '欢迎来到GinBlog'
        }
    },
    {
        path: '/detail/:id',
        component: Detail,
        meta: {
            title: '文章详情页',
        },
        props: true
    },
]

const router = new VueRouter({
    routes
})

router.beforeEach((to, form, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

export default router
