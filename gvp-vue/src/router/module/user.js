const userRouter = [{
  path: '/register',
  name: 'register',
  component: () => import('@/views/register/UserRegister.vue'),
},
{
  path: '/login',
  name: 'login',
  component: () => import('@/views/login/UserLogin.vue'),
},
{
  path: '/profile',
  name: 'profile',
  meta: {
    auth: true,
  },
  component: () => import('@/views/profile/UserProfile.vue'),
},
];
export default userRouter;
