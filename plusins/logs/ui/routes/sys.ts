export default {
  path: '/sys/logs',
  name: 'SysLogs',
  component: () => import('@/views/logs/index.vue'),
  meta: { title: '日志管理', requiresAuth: true },
};
