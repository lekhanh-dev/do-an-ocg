import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  {
    path: "/cart",
    name: "Cart",
    component: () => import("../views/Cart.vue"),
  },
  {
    path: "/news",
    name: "News",
    component: () => import("../views/News.vue"),
  },
  {
    path: "/category/:categoryName/:categoryId/:categoryChildName/:categoryChildId",
    name: "Category",
    component: () => import("../views/Category.vue"),
    props: true,
  },
  {
    path: "/product-detail/:productName/:productId",
    name: "ProductDetail",
    component: () => import("../views/ProductDetail.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  scrollBehavior() {
    // if (savedPosition) {
    //   return savedPosition;
    // } else {
    // return { top: 0 };
    // }
    return { top: 0 };
  },
});

export default router;
