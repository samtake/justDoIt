import Vue from "vue";
import Router from "vue-router";
import findLast from "lodash/findLast";
import { notification } from "ant-design-vue";
import NProgress from "nprogress";
import "nprogress/nprogress.css";
import NotFound from "./views/404";
import Forbidden from "./views/403";
import { check, isLogin } from "./utils/auth";

Vue.use(Router);

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/user",
      hideInMenu: true,
      component: () =>
        import(/* webpackChunkName: "layout" */ "./layouts/UserLayout"),
      children: [
        {
          path: "/user",
          redirect: "/user/login"
        },
        {
          path: "/user/login",
          name: "login",
          component: () =>
            import(/* webpackChunkName: "user" */ "./views/User/Login")
        },
        {
          path: "/user/register",
          name: "register",
          component: () =>
            import(/* webpackChunkName: "user" */ "./views/User/Register")
        }
      ]
    },
    {
      path: "/",
      meta: { authority: ["user", "admin"] },
      component: () =>
        import(/* webpackChunkName: "layout" */ "./layouts/BasicLayout"),
      children: [
        // home
        {
          path: "/home",
          name: "home",
          component: { render: h => h("router-view") },
          meta: { icon: "form", title: "首页", authority: ["admin"] },
          children: [
            {
              path: "/home/basic-home",
              name: "basichome",
              meta: { title: "基础表单" },
              component: () =>
                import(/* webpackChunkName: "form" */ "./views/Home/BasicHome")
            },
            {
              path: "/form/step-home",
              name: "stephome",
              hideChildrenInMenu: true,
              meta: { title: "分布表单" },
              component: () =>
                import(/* webpackChunkName: "form" */ "./views/Home/StepHome"),
              children: [
                {
                  path: "/form/step-home",
                  redirect: "/form/step-home/info"
                },
                {
                  path: "/form/step-home/info",
                  name: "info",
                  component: () =>
                    import(/* webpackChunkName: "form" */ "./views/Home/StepHome/BannerList")
                }
              ]
            }
          ]
        }
      ]
    },
    {
      path: "/403",
      name: "403",
      hideInMenu: true,
      component: Forbidden
    },
    {
      path: "*",
      name: "404",
      hideInMenu: true,
      component: NotFound
    }
  ]
});

router.beforeEach((to, from, next) => {
  if (to.path !== from.path) {
    NProgress.start();
  }
  const record = findLast(to.matched, record => record.meta.authority);
  if (record && !check(record.meta.authority)) {
    if (!isLogin() && to.path !== "/user/login") {
      next({
        path: "/user/login"
      });
    } else if (to.path !== "/403") {
      notification.error({
        message: "403",
        description: "你没有权限访问，请联系管理员咨询。"
      });
      next({
        path: "/403"
      });
    }
    NProgress.done();
  }

  next();
});

router.afterEach(() => {
  NProgress.done();
});

export default router;
