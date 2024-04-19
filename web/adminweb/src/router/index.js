import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginView from "@/views/LoginView.vue";
import AdminView from "@/views/AdminView.vue";
import Index from "@/components/admin/IndexView.vue";
import AddArticle from '@/components/article/addArticleView.vue';
import ArtList from '@/components/article/artListView.vue';
import CategoryList from '@/components/category/categoryListView.vue';
import UserList from '@/components/user/userListView.vue'
import Profile from '@/components/user/profileView'

Vue.use(VueRouter)

const routes = [
    {
        path: '/login',
        name: 'login',
        meta: {
            title: "登录页面"
        },
        component: LoginView
    },
    {
        path: '/',
        name: 'admin',
        component: AdminView,
        children: [
            {
                path: 'index',
                component: Index,
                meta: {
                    title: "后台管理页面"
                }
            },
            {
                path: 'addArticle',
                component: AddArticle,
                meta: {
                    title: "添加文章"
                }
            },
             {
                path: 'addArticle/:id',
                component: AddArticle,
                meta: {
                    title: "修改文章"
                }
            },
            {
                path: 'artList',
                component: ArtList,
                meta: {
                    title: "文章列表"
                }
            },
            {
                path: 'categoryList',
                component: CategoryList,
                meta: {
                    title: "分类列表"
                }
            },
            {
                path: 'userList',
                component: UserList,
                meta: {
                    title: "用户列表"
                }
            },
             {
                path: 'profile',
                component: Profile,
                meta: {
                    title: "个人设置"
                }
            }
        ]
    }
]

const router = new VueRouter({
    routes
})
router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
    const token = window.sessionStorage.getItem('token')
    if (to.path === '/login') return next()
    if (!token && to.path === '/') {
        next('/login')
    } else {
        next()
    }
})
export default router
